// +build !ignore_autogenerated

/*


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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CalicoNetworking) DeepCopyInto(out *CalicoNetworking) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CalicoNetworking.
func (in *CalicoNetworking) DeepCopy() *CalicoNetworking {
	if in == nil {
		return nil
	}
	out := new(CalicoNetworking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CalicoNetworking) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CalicoNetworkingList) DeepCopyInto(out *CalicoNetworkingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CalicoNetworking, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CalicoNetworkingList.
func (in *CalicoNetworkingList) DeepCopy() *CalicoNetworkingList {
	if in == nil {
		return nil
	}
	out := new(CalicoNetworkingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CalicoNetworkingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CalicoNetworkingSpec) DeepCopyInto(out *CalicoNetworkingSpec) {
	*out = *in
	out.CommonSpec = in.CommonSpec
	in.PatchSpec.DeepCopyInto(&out.PatchSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CalicoNetworkingSpec.
func (in *CalicoNetworkingSpec) DeepCopy() *CalicoNetworkingSpec {
	if in == nil {
		return nil
	}
	out := new(CalicoNetworkingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CalicoNetworkingStatus) DeepCopyInto(out *CalicoNetworkingStatus) {
	*out = *in
	in.CommonStatus.DeepCopyInto(&out.CommonStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CalicoNetworkingStatus.
func (in *CalicoNetworkingStatus) DeepCopy() *CalicoNetworkingStatus {
	if in == nil {
		return nil
	}
	out := new(CalicoNetworkingStatus)
	in.DeepCopyInto(out)
	return out
}
