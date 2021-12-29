// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessChannels) DeepCopyInto(out *AccessChannels) {
	*out = *in
	out.SSH = in.SSH
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessChannels.
func (in *AccessChannels) DeepCopy() *AccessChannels {
	if in == nil {
		return nil
	}
	out := new(AccessChannels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	if in.RegionIDs != nil {
		in, out := &in.RegionIDs, &out.RegionIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ZoneIDs != nil {
		in, out := &in.ZoneIDs, &out.ZoneIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.AccessChannels = in.AccessChannels
	in.Metadata.DeepCopyInto(&out.Metadata)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMeta) DeepCopyInto(out *ClusterMeta) {
	*out = *in
	in.Network.DeepCopyInto(&out.Network)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMeta.
func (in *ClusterMeta) DeepCopy() *ClusterMeta {
	if in == nil {
		return nil
	}
	out := new(ClusterMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNetworkMeta) DeepCopyInto(out *ClusterNetworkMeta) {
	*out = *in
	if in.ExportPorts != nil {
		in, out := &in.ExportPorts, &out.ExportPorts
		*out = make([]ExportPort, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNetworkMeta.
func (in *ClusterNetworkMeta) DeepCopy() *ClusterNetworkMeta {
	if in == nil {
		return nil
	}
	out := new(ClusterNetworkMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Credential) DeepCopyInto(out *Credential) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Credential.
func (in *Credential) DeepCopy() *Credential {
	if in == nil {
		return nil
	}
	out := new(Credential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disk) DeepCopyInto(out *Disk) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disk.
func (in *Disk) DeepCopy() *Disk {
	if in == nil {
		return nil
	}
	out := new(Disk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportPort) DeepCopyInto(out *ExportPort) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportPort.
func (in *ExportPort) DeepCopy() *ExportPort {
	if in == nil {
		return nil
	}
	out := new(ExportPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Host) DeepCopyInto(out *Host) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make([]Disk, len(*in))
		copy(*out, *in)
	}
	out.OS = in.OS
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Host.
func (in *Host) DeepCopy() *Host {
	if in == nil {
		return nil
	}
	out := new(Host)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostStatus) DeepCopyInto(out *HostStatus) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostStatus.
func (in *HostStatus) DeepCopy() *HostStatus {
	if in == nil {
		return nil
	}
	out := new(HostStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Infra) DeepCopyInto(out *Infra) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Infra.
func (in *Infra) DeepCopy() *Infra {
	if in == nil {
		return nil
	}
	out := new(Infra)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Infra) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraList) DeepCopyInto(out *InfraList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Infra, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraList.
func (in *InfraList) DeepCopy() *InfraList {
	if in == nil {
		return nil
	}
	out := new(InfraList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfraList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraSpec) DeepCopyInto(out *InfraSpec) {
	*out = *in
	out.Credential = in.Credential
	in.Cluster.DeepCopyInto(&out.Cluster)
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]Host, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraSpec.
func (in *InfraSpec) DeepCopy() *InfraSpec {
	if in == nil {
		return nil
	}
	out := new(InfraSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraStatus) DeepCopyInto(out *InfraStatus) {
	*out = *in
	in.Cluster.DeepCopyInto(&out.Cluster)
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]HostStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraStatus.
func (in *InfraStatus) DeepCopy() *InfraStatus {
	if in == nil {
		return nil
	}
	out := new(InfraStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OS) DeepCopyInto(out *OS) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OS.
func (in *OS) DeepCopy() *OS {
	if in == nil {
		return nil
	}
	out := new(OS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSH) DeepCopyInto(out *SSH) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSH.
func (in *SSH) DeepCopy() *SSH {
	if in == nil {
		return nil
	}
	out := new(SSH)
	in.DeepCopyInto(out)
	return out
}
