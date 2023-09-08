// Copyright © 2023 sealos.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k3s

import (
	"fmt"
	"path/filepath"

	"github.com/labring/sealos/pkg/constants"

	"github.com/labring/sealos/pkg/utils/file"
	"github.com/labring/sealos/pkg/utils/logger"
	"github.com/labring/sealos/pkg/utils/rand"
	"github.com/labring/sealos/pkg/utils/yaml"
)

func (k *K3s) initMaster0() error {
	master0 := k.cluster.GetMaster0IPAndPort()
	return k.runPipelines("init master0",
		k.generateAndSendCerts,
		func() error { return k.generateAndSendTokenFiles(master0, "token", "agent-token") },
		k.generateAndSendInitConfig,
	)
}

func (k *K3s) joinMasters(masters []string) error {
	_, err := k.writeJoinConfigWithCallbacks(serverMode)
	if err != nil {
		return err
	}
	for _, master := range masters {
		if err = k.joinMaster(master); err != nil {
			return err
		}
	}
	return nil
}

func (k *K3s) writeJoinConfigWithCallbacks(runMode string, callbacks ...callback) (string, error) {
	defaultCallbacks := []callback{defaultingConfig, k.merge, k.sealosCfg, k.overrideCertSans}
	switch runMode {
	case serverMode:
		defaultCallbacks = append(defaultCallbacks, k.overrideServerConfig)
	case agentMode:
		defaultCallbacks = append(defaultCallbacks, k.overrideAgentConfig)
	}

	defaultCallbacks = append(defaultCallbacks,
		func(c *Config) *Config {
			c.ServerURL = fmt.Sprintf("https://%s:%d", constants.DefaultAPIServerDomain, c.HTTPSPort)
			return c
		},
	)
	raw, err := k.getRawInitConfig(
		append(defaultCallbacks, callbacks...)...,
	)
	if err != nil {
		return "", err
	}
	var filename string
	switch runMode {
	case serverMode:
		filename = defaultJoinMastersFilename
	case agentMode:
		filename = defaultJoinNodesFilename
	}
	path := filepath.Join(k.pathResolver.EtcPath(), filename)
	return path, file.WriteFile(path, raw)
}

func (k *K3s) joinMaster(master string) error {
	return k.runPipelines(fmt.Sprintf("join master %s", master),
		func() error {
			// the rest masters are also running in agent mode, so agent-token file is needed.
			return k.generateAndSendTokenFiles(master, "token", "agent-token")
		},
		func() error {
			return k.sshClient.Copy(master, filepath.Join(k.pathResolver.EtcPath(), defaultJoinMastersFilename), defaultConfigPath)
		},
	)
}

func (k *K3s) joinNodes(nodes []string) error {
	if _, err := k.writeJoinConfigWithCallbacks(agentMode, removeServerFlagsInAgentConfig); err != nil {
		return err
	}
	for i := range nodes {
		if err := k.joinNode(nodes[i]); err != nil {
			return err
		}
	}
	return nil
}

func (k *K3s) joinNode(node string) error {
	return k.runPipelines(fmt.Sprintf("join node %s", node),
		func() error { return k.generateAndSendTokenFiles(node, "agent-token") },
		func() error {
			return k.sshClient.Copy(node, filepath.Join(k.pathResolver.EtcPath(), defaultJoinNodesFilename), defaultConfigPath)
		},
	)
}

func (k *K3s) generateAndSendCerts() error {
	logger.Debug("generate and send self-signed certificates")
	// TODO: use self-signed certificates
	return nil
}

func (k *K3s) generateRandomTokenFileIfNotExists(filename string) (string, error) {
	fp := filepath.Join(k.pathResolver.EtcPath(), filepath.Base(filename))
	if !file.IsExist(fp) {
		logger.Debug("token file %s not exists, create new one", fp)
		token, err := rand.CreateCertificateKey()
		if err != nil {
			return "", err
		}
		return fp, file.WriteFile(fp, []byte(token))
	}
	return fp, nil
}

func (k *K3s) generateAndSendTokenFiles(host string, filenames ...string) error {
	for _, filename := range filenames {
		src, err := k.generateRandomTokenFileIfNotExists(filename)
		if err != nil {
			return fmt.Errorf("generate token: %v", err)
		}
		dst := filepath.Join(k.pathResolver.ConfigsPath(), filename)
		if err = k.sshClient.Copy(host, src, dst); err != nil {
			return fmt.Errorf("copy token file: %v", err)
		}
	}
	return nil
}

func (k *K3s) getRawInitConfig(callbacks ...callback) ([]byte, error) {
	cfg, err := k.getInitConfig(callbacks...)
	if err != nil {
		return nil, err
	}
	return yaml.MarshalConfigs(cfg)
}

func (k *K3s) generateAndSendInitConfig() error {
	src := filepath.Join(k.pathResolver.EtcPath(), defaultInitFilename)
	defaultCallbacks := []callback{defaultingConfig, k.merge, k.sealosCfg, k.overrideCertSans, k.overrideServerConfig, setClusterInit}
	if !file.IsExist(src) {
		raw, err := k.getRawInitConfig(defaultCallbacks...)
		if err != nil {
			return err
		}
		if err = file.WriteFile(src, raw); err != nil {
			return err
		}
	}
	return k.sshClient.Copy(k.cluster.GetMaster0IPAndPort(), src, defaultConfigPath)
}
