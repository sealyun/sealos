//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingList) DeepCopyInto(out *BillingList) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingList.
func (in *BillingList) DeepCopy() *BillingList {
	if in == nil {
		return nil
	}
	out := new(BillingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Deduction) DeepCopyInto(out *Deduction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Deduction.
func (in *Deduction) DeepCopy() *Deduction {
	if in == nil {
		return nil
	}
	out := new(Deduction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Deduction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeductionList) DeepCopyInto(out *DeductionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Deduction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeductionList.
func (in *DeductionList) DeepCopy() *DeductionList {
	if in == nil {
		return nil
	}
	out := new(DeductionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeductionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeductionSpec) DeepCopyInto(out *DeductionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeductionSpec.
func (in *DeductionSpec) DeepCopy() *DeductionSpec {
	if in == nil {
		return nil
	}
	out := new(DeductionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeductionStatus) DeepCopyInto(out *DeductionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeductionStatus.
func (in *DeductionStatus) DeepCopy() *DeductionStatus {
	if in == nil {
		return nil
	}
	out := new(DeductionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionResourcesPrice) DeepCopyInto(out *ExtensionResourcesPrice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionResourcesPrice.
func (in *ExtensionResourcesPrice) DeepCopy() *ExtensionResourcesPrice {
	if in == nil {
		return nil
	}
	out := new(ExtensionResourcesPrice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExtensionResourcesPrice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionResourcesPriceList) DeepCopyInto(out *ExtensionResourcesPriceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExtensionResourcesPrice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionResourcesPriceList.
func (in *ExtensionResourcesPriceList) DeepCopy() *ExtensionResourcesPriceList {
	if in == nil {
		return nil
	}
	out := new(ExtensionResourcesPriceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExtensionResourcesPriceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionResourcesPriceSpec) DeepCopyInto(out *ExtensionResourcesPriceSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(map[corev1.ResourceName]ResourcePrice, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionResourcesPriceSpec.
func (in *ExtensionResourcesPriceSpec) DeepCopy() *ExtensionResourcesPriceSpec {
	if in == nil {
		return nil
	}
	out := new(ExtensionResourcesPriceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionResourcesPriceStatus) DeepCopyInto(out *ExtensionResourcesPriceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionResourcesPriceStatus.
func (in *ExtensionResourcesPriceStatus) DeepCopy() *ExtensionResourcesPriceStatus {
	if in == nil {
		return nil
	}
	out := new(ExtensionResourcesPriceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metering) DeepCopyInto(out *Metering) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metering.
func (in *Metering) DeepCopy() *Metering {
	if in == nil {
		return nil
	}
	out := new(Metering)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Metering) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeteringList) DeepCopyInto(out *MeteringList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Metering, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeteringList.
func (in *MeteringList) DeepCopy() *MeteringList {
	if in == nil {
		return nil
	}
	out := new(MeteringList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MeteringList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeteringQuota) DeepCopyInto(out *MeteringQuota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeteringQuota.
func (in *MeteringQuota) DeepCopy() *MeteringQuota {
	if in == nil {
		return nil
	}
	out := new(MeteringQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MeteringQuota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeteringQuotaList) DeepCopyInto(out *MeteringQuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MeteringQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeteringQuotaList.
func (in *MeteringQuotaList) DeepCopy() *MeteringQuotaList {
	if in == nil {
		return nil
	}
	out := new(MeteringQuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MeteringQuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeteringQuotaSpec) DeepCopyInto(out *MeteringQuotaSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(map[corev1.ResourceName]ResourceUsed, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeteringQuotaSpec.
func (in *MeteringQuotaSpec) DeepCopy() *MeteringQuotaSpec {
	if in == nil {
		return nil
	}
	out := new(MeteringQuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeteringQuotaStatus) DeepCopyInto(out *MeteringQuotaStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeteringQuotaStatus.
func (in *MeteringQuotaStatus) DeepCopy() *MeteringQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(MeteringQuotaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeteringSpec) DeepCopyInto(out *MeteringSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(map[corev1.ResourceName]ResourcePrice, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeteringSpec.
func (in *MeteringSpec) DeepCopy() *MeteringSpec {
	if in == nil {
		return nil
	}
	out := new(MeteringSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeteringStatus) DeepCopyInto(out *MeteringStatus) {
	*out = *in
	if in.BillingListM != nil {
		in, out := &in.BillingListM, &out.BillingListM
		*out = make([]BillingList, len(*in))
		copy(*out, *in)
	}
	if in.BillingListH != nil {
		in, out := &in.BillingListH, &out.BillingListH
		*out = make([]BillingList, len(*in))
		copy(*out, *in)
	}
	if in.BillingListD != nil {
		in, out := &in.BillingListD, &out.BillingListD
		*out = make([]BillingList, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeteringStatus.
func (in *MeteringStatus) DeepCopy() *MeteringStatus {
	if in == nil {
		return nil
	}
	out := new(MeteringStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodResourcePrice) DeepCopyInto(out *PodResourcePrice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodResourcePrice.
func (in *PodResourcePrice) DeepCopy() *PodResourcePrice {
	if in == nil {
		return nil
	}
	out := new(PodResourcePrice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodResourcePrice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodResourcePriceList) DeepCopyInto(out *PodResourcePriceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodResourcePrice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodResourcePriceList.
func (in *PodResourcePriceList) DeepCopy() *PodResourcePriceList {
	if in == nil {
		return nil
	}
	out := new(PodResourcePriceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodResourcePriceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodResourcePriceSpec) DeepCopyInto(out *PodResourcePriceSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(map[corev1.ResourceName]ResourcePrice, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodResourcePriceSpec.
func (in *PodResourcePriceSpec) DeepCopy() *PodResourcePriceSpec {
	if in == nil {
		return nil
	}
	out := new(PodResourcePriceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodResourcePriceStatus) DeepCopyInto(out *PodResourcePriceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodResourcePriceStatus.
func (in *PodResourcePriceStatus) DeepCopy() *PodResourcePriceStatus {
	if in == nil {
		return nil
	}
	out := new(PodResourcePriceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMsg) DeepCopyInto(out *ResourceMsg) {
	*out = *in
	if in.Used != nil {
		in, out := &in.Used, &out.Used
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMsg.
func (in *ResourceMsg) DeepCopy() *ResourceMsg {
	if in == nil {
		return nil
	}
	out := new(ResourceMsg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePrice) DeepCopyInto(out *ResourcePrice) {
	*out = *in
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePrice.
func (in *ResourcePrice) DeepCopy() *ResourcePrice {
	if in == nil {
		return nil
	}
	out := new(ResourcePrice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceUsed) DeepCopyInto(out *ResourceUsed) {
	*out = *in
	if in.Used != nil {
		in, out := &in.Used, &out.Used
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceUsed.
func (in *ResourceUsed) DeepCopy() *ResourceUsed {
	if in == nil {
		return nil
	}
	out := new(ResourceUsed)
	in.DeepCopyInto(out)
	return out
}
