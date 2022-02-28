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

package args

import v2 "github.com/fanux/sealos/pkg/types/v1beta1"

type RunArgs struct {
	Masters         string
	Nodes           string
	User            string
	Password        string
	Port            string
	Pk              string
	PkPassword      string
	PodCidr         string
	SvcCidr         string
	APIServerDomain string
	VIP             string
	CertSANS        []string
	Interface       string
	IPIPFalse       bool
	MTU             string
	RegistryAddress string
	CRIData         string
	RegistryConfig  string
	RegistryData    string
	KubeadmfilePath string
	Amd64URI        string
	Arm64URI        string
	CTLAdm64URI     string
	CTLArm64URI     string
}

type run struct {
	cluster   *v2.Cluster
	configs   []v2.Config
	_packages []v2.Package
	hosts     []v2.ClusterHost
}
