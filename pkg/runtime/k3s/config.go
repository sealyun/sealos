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
	"bufio"
	"bytes"
	"io"
	"path/filepath"

	"github.com/imdario/mergo"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"
	netutils "k8s.io/utils/net"
	"sigs.k8s.io/yaml"

	"github.com/labring/sealos/pkg/constants"
	"github.com/labring/sealos/pkg/utils/logger"
)

var defaultMergeOpts = []func(*mergo.Config){
	mergo.WithOverride,
}

func defaultingServerConfig(c *Config) {
	c.BindAddress = "0.0.0.0"
	c.HTTPSPort = 6443
	c.ClusterCIDR = []string{"10.42.0.0/16"}
	c.ServiceCIDR = []string{"10.43.0.0/16"}
	c.ClusterDomain = constants.DefaultDNSDomain
	c.DisableCCM = true
	c.DisableHelmController = true
	c.DisableNPC = true
	c.EgressSelectorMode = "agent"
	c.FlannelBackend = "vxlan"
	c.KubeConfigMode = "0644"
	c.PreferBundledBin = true
	c.Disable = []string{"servicelb", "traefik", "local-storage", "metrics-server"}
	defaultingAgentConfig(c)
}

func defaultingAgentConfig(c *Config) {
	if c.AgentConfig == nil {
		c.AgentConfig = &AgentConfig{}
	}
	c.AgentConfig.ExtraKubeProxyArgs = []string{}
	c.AgentConfig.ExtraKubeletArgs = []string{}
	c.AgentConfig.PauseImage = "docker.io/rancher/pause:3.1"
	c.AgentConfig.PrivateRegistry = "/etc/rancher/k3s/registries.yaml"
	c.AgentConfig.Labels = []string{"provisioner=k3s"}
}

type callback func(*Config)

func setClusterInit(c *Config) {
	c.ClusterInit = true
}

func (k *K3s) merge(c *Config) {
	if k.config == nil {
		return
	}
	if err := mergo.Merge(c, k.config, defaultMergeOpts...); err != nil {
		logger.Error("failed to merge config: %v", err)
	}
}

func (k *K3s) overrideServerConfig(c *Config) {
	c.TokenFile = filepath.Join(k.pathResolver.ConfigsPath(), "token")
	c.AgentTokenFile = filepath.Join(k.pathResolver.ConfigsPath(), "agent-token")
	c.TLSSan = append(c.TLSSan, constants.DefaultAPIServerDomain)

	if len(c.ClusterDNS) == 0 && len(c.ServiceCIDR) > 0 {
		svcSubnetCIDR, err := netutils.ParseCIDRs(c.ServiceCIDR)
		if err == nil {
			clusterDns, err := netutils.GetIndexedIP(svcSubnetCIDR[0], 10)
			if err == nil {
				c.ClusterDNS = []string{clusterDns.String()}
			}
		}
	}
}

func (k *K3s) overrideAgentConfig(c *Config) {
	c.AgentTokenFile = filepath.Join(k.pathResolver.ConfigsPath(), "agent-token")
}

func (k *K3s) getInitConfig(callbacks ...callback) (*Config, error) {
	cfg := Config{}
	for i := range callbacks {
		callbacks[i](&cfg)
	}
	return &cfg, nil
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (c *Config) getContainerRuntimeEndpoint() string {
	if c.AgentConfig.Docker {
		return "unix:///run/k3s/cri-dockerd/cri-dockerd.sock"
	} else if len(c.AgentConfig.ContainerRuntimeEndpoint) == 0 {
		return "unix:///run/k3s/containerd/containerd.sock"
	}
	return c.AgentConfig.ContainerRuntimeEndpoint
}

// ParseConfig return nil if data structure is not matched
func ParseConfig(data []byte) (*Config, error) {
	d := yamlutil.NewYAMLReader(bufio.NewReader(bytes.NewReader(data)))
	for {
		b, err := d.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		cfg, err := parseConfig(b)
		if err != nil {
			return nil, err
		}
		if cfg != nil {
			return cfg, nil
		}
	}
	return nil, nil
}

func parseConfig(data []byte) (*Config, error) {
	var c Config
	if err := yaml.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	out, err := yaml.Marshal(&c)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := yaml.Unmarshal(out, &m); err != nil {
		return nil, err
	}
	if len(m) == 0 {
		return nil, nil
	}
	return &c, nil
}
