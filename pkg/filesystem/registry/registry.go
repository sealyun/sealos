/*
Copyright 2022 fengxsong@outlook.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package registry

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/labring/sealos/pkg/constants"
	"github.com/labring/sealos/pkg/ssh"
	v2 "github.com/labring/sealos/pkg/types/v1beta1"
)

const (
	defaultUntarRegistry = "cd %s/%s ;if [ -d ../registry/compressed/ ]; then find ../registry/compressed/ -type f -exec file {} \\; | grep compressed | awk -F: '{print $1}' | while IFS='' read -r cpd; do tar -zxf \"$cpd\"  -C ../registry && rm -rf \"$cpd\" ; done; fi\n "
)

type Interface interface {
	MirrorTo(context.Context, ...string) error
}

type scp struct {
	root   string
	ssh    ssh.Interface
	mounts []v2.MountImage
}

func (s *scp) MirrorTo(ctx context.Context, hosts ...string) error {
	if ctx == nil {
		ctx = context.Background()
	}
	outerEg, _ := errgroup.WithContext(ctx)
	for i := range hosts {
		host := hosts[i]
		outerEg.Go(func() error {
			eg, _ := errgroup.WithContext(ctx)
			for j := range s.mounts {
				m := s.mounts[j]
				eg.Go(func() error {
					return ssh.CopyDir(s.ssh, host, m.MountPoint, s.root, constants.IsRegistryDir)
				})
			}
			if err := eg.Wait(); err != nil {
				return err
			}
			return s.ssh.CmdAsync(host, fmt.Sprintf(defaultUntarRegistry, s.root, constants.ScriptsDirName))
		})
	}
	return outerEg.Wait()
}

func New(root string, ssh ssh.Interface, mounts []v2.MountImage) Interface {
	return &scp{root, ssh, mounts}
}
