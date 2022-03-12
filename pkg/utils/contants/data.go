/*
Copyright 2022 cuisongliu@qq.com.

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

package contants

import "path/filepath"

const (
	DefaultClusterRootfsDir = "/var/lib/sealos"
	DefaultKubeadmFileName  = "Kubeadmfile"
)

func LogPath() string {
	return filepath.Join(DefaultClusterRootfsDir, "log")
}
func DataPath() string {
	return filepath.Join(DefaultClusterRootfsDir, "data")
}
func ResourcePath() string {
	return filepath.Join(DefaultClusterRootfsDir, "resource")
}

func TempPath() string {
	return filepath.Join(DefaultClusterRootfsDir, "temp")
}

type Data interface {
	Homedir() string
	KubePath() string
	KubeEtcPath() string
	KubeScriptsPath() string
	KubeRegistryPath() string
	KubeKubeadmfile() string

	PkiPath() string
	PkiEtcdPath() string
	AdminFile() string
	EtcPath() string

	KubeCharsPath() string
	KubeManifestsPath() string
	KubeSealctlPath() string
}

type data struct {
	clusterName string
}

func (d *data) KubeSealctlPath() string {
	return filepath.Join(d.KubePath(), "opt", "sealctl")
}

func (d *data) KubeScriptsPath() string {
	return filepath.Join(d.KubePath(), "scripts")
}
func (d *data) KubeEtcPath() string {
	return filepath.Join(d.KubePath(), "etc")
}
func (d *data) KubeKubeadmfile() string {
	return filepath.Join(d.KubePath(), "etc", DefaultKubeadmFileName)
}

func (d *data) KubeRegistryPath() string {
	return filepath.Join(d.KubePath(), "registry")
}

func (d *data) KubeCharsPath() string {
	return filepath.Join(d.KubePath(), "charts")
}

func (d *data) KubeManifestsPath() string {
	return filepath.Join(d.KubePath(), "manifests")
}

func (d *data) EtcPath() string {
	return filepath.Join(d.Homedir(), "etc")
}
func (d *data) AdminFile() string {
	return filepath.Join(d.EtcPath(), "admin.conf")
}

func (d *data) PkiPath() string {
	return filepath.Join(d.Homedir(), "pki")
}

func (d *data) PkiEtcdPath() string {
	return filepath.Join(d.PkiPath(), "etcd")
}

func (d *data) KubePath() string {
	return filepath.Join(d.Homedir(), DataDirName)
}

func (d *data) Homedir() string {
	return filepath.Join(DefaultClusterRootfsDir, "data", d.clusterName)
}

func NewData(clusterName string) Data {
	return &data{
		clusterName: clusterName,
	}
}
