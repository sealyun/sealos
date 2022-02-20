// Copyright © 2021 sealos.
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

package ssh

import (
	"fmt"
	v2 "github.com/fanux/sealos/pkg/types/v1beta1"
	"github.com/fanux/sealos/pkg/utils/iputils"
	"github.com/fanux/sealos/pkg/utils/logger"
	"net"
	"sync"
	"time"
)

type Interface interface {
	// Copy is copy local files to remote host
	// scp -r /tmp root@192.168.0.2:/root/tmp => Copy("192.168.0.2","tmp","/root/tmp")
	// need check md5sum
	Copy(host, srcFilePath, dstFilePath string) error
	// Fetch is copy remote host files to localhost
	Fetch(host, srcFilePath, dstFilePath string) error
	// CmdAsync is exec command on remote host, and asynchronous return logs
	CmdAsync(host string, cmd ...string) error
	// Cmd is exec command on remote host, and return combined standard output and standard error
	Cmd(host, cmd string) ([]byte, error)
	// IsFileExist is check remote file exist or not
	IsFileExist(host, remoteFilePath string) bool
	//RemoteDirExist is remote file existence returns true, nil
	RemoteDirExist(host, remoteDirpath string) (bool, error)
	//CmdToString is exec command on remote host, and return spilt standard output and standard error
	CmdToString(host, cmd, spilt string) (string, error)
	Ping(host string) error
}

type SSH struct {
	isStdout     bool
	User         string
	Password     string
	PkFile       string
	PkPassword   string
	Timeout      *time.Duration
	LocalAddress *[]net.Addr
}

func NewSSHByCluster(cluster *v2.Cluster, isStdout bool) Interface {
	address, err := iputils.IsLocalHostAddrs()
	if err != nil {
		logger.Warn("failed to get local address, %v", err)
	}
	return &SSH{
		isStdout:     isStdout,
		User:         cluster.Spec.SSH.User,
		Password:     cluster.Spec.SSH.Passwd,
		PkFile:       cluster.Spec.SSH.Pk,
		PkPassword:   cluster.Spec.SSH.PkPasswd,
		LocalAddress: address,
	}
}

type Client struct {
	SSH  Interface
	Host string
}

func WaitSSHReady(ssh Interface, tryTimes int, hosts ...string) error {
	var err error
	var wg sync.WaitGroup
	for _, h := range hosts {
		wg.Add(1)
		go func(host string) {
			defer wg.Done()
			for i := 0; i < tryTimes; i++ {
				err = ssh.Ping(host)
				if err == nil {
					return
				}
				time.Sleep(time.Duration(i) * time.Second)
			}
			err = fmt.Errorf("wait for [%s] ssh ready timeout:  %v, ensure that the IP address or password is correct", host, err)
		}(h)
	}
	wg.Wait()
	return err
}
