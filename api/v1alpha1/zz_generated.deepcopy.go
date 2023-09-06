//go:build !ignore_autogenerated

/*
Copyright 2023.

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
func (in *ConfigMapRef) DeepCopyInto(out *ConfigMapRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapRef.
func (in *ConfigMapRef) DeepCopy() *ConfigMapRef {
	if in == nil {
		return nil
	}
	out := new(ConfigMapRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBasedUpgrade) DeepCopyInto(out *ImageBasedUpgrade) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBasedUpgrade.
func (in *ImageBasedUpgrade) DeepCopy() *ImageBasedUpgrade {
	if in == nil {
		return nil
	}
	out := new(ImageBasedUpgrade)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageBasedUpgrade) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBasedUpgradeList) DeepCopyInto(out *ImageBasedUpgradeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageBasedUpgrade, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBasedUpgradeList.
func (in *ImageBasedUpgradeList) DeepCopy() *ImageBasedUpgradeList {
	if in == nil {
		return nil
	}
	out := new(ImageBasedUpgradeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageBasedUpgradeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBasedUpgradeSpec) DeepCopyInto(out *ImageBasedUpgradeSpec) {
	*out = *in
	out.SeedImageRef = in.SeedImageRef
	out.AdditionalImages = in.AdditionalImages
	out.OADPContent = in.OADPContent
	out.ExtraManifests = in.ExtraManifests
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBasedUpgradeSpec.
func (in *ImageBasedUpgradeSpec) DeepCopy() *ImageBasedUpgradeSpec {
	if in == nil {
		return nil
	}
	out := new(ImageBasedUpgradeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBasedUpgradeStatus) DeepCopyInto(out *ImageBasedUpgradeStatus) {
	*out = *in
	in.StartedAt.DeepCopyInto(&out.StartedAt)
	in.CompletedAt.DeepCopyInto(&out.CompletedAt)
	if in.StateRoots != nil {
		in, out := &in.StateRoots, &out.StateRoots
		*out = make([]StateRoot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBasedUpgradeStatus.
func (in *ImageBasedUpgradeStatus) DeepCopy() *ImageBasedUpgradeStatus {
	if in == nil {
		return nil
	}
	out := new(ImageBasedUpgradeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodState) DeepCopyInto(out *PodState) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodState.
func (in *PodState) DeepCopy() *PodState {
	if in == nil {
		return nil
	}
	out := new(PodState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeedImageRef) DeepCopyInto(out *SeedImageRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeedImageRef.
func (in *SeedImageRef) DeepCopy() *SeedImageRef {
	if in == nil {
		return nil
	}
	out := new(SeedImageRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateRoot) DeepCopyInto(out *StateRoot) {
	*out = *in
	if in.PodStates != nil {
		in, out := &in.PodStates, &out.PodStates
		*out = make([]PodState, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateRoot.
func (in *StateRoot) DeepCopy() *StateRoot {
	if in == nil {
		return nil
	}
	out := new(StateRoot)
	in.DeepCopyInto(out)
	return out
}
