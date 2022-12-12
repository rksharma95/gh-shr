//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright 2022 Authors of KubeArmor

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapabilitiesType) DeepCopyInto(out *CapabilitiesType) {
	*out = *in
	if in.MatchCapabilities != nil {
		in, out := &in.MatchCapabilities, &out.MatchCapabilities
		*out = make([]MatchCapabilitiesType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapabilitiesType.
func (in *CapabilitiesType) DeepCopy() *CapabilitiesType {
	if in == nil {
		return nil
	}
	out := new(CapabilitiesType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileDirectoryType) DeepCopyInto(out *FileDirectoryType) {
	*out = *in
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileDirectoryType.
func (in *FileDirectoryType) DeepCopy() *FileDirectoryType {
	if in == nil {
		return nil
	}
	out := new(FileDirectoryType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilePathType) DeepCopyInto(out *FilePathType) {
	*out = *in
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilePathType.
func (in *FilePathType) DeepCopy() *FilePathType {
	if in == nil {
		return nil
	}
	out := new(FilePathType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilePatternType) DeepCopyInto(out *FilePatternType) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilePatternType.
func (in *FilePatternType) DeepCopy() *FilePatternType {
	if in == nil {
		return nil
	}
	out := new(FilePatternType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileType) DeepCopyInto(out *FileType) {
	*out = *in
	if in.MatchPaths != nil {
		in, out := &in.MatchPaths, &out.MatchPaths
		*out = make([]FilePathType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchDirectories != nil {
		in, out := &in.MatchDirectories, &out.MatchDirectories
		*out = make([]FileDirectoryType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchPatterns != nil {
		in, out := &in.MatchPatterns, &out.MatchPatterns
		*out = make([]FilePatternType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileType.
func (in *FileType) DeepCopy() *FileType {
	if in == nil {
		return nil
	}
	out := new(FileType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostCapabilitiesType) DeepCopyInto(out *HostCapabilitiesType) {
	*out = *in
	if in.MatchCapabilities != nil {
		in, out := &in.MatchCapabilities, &out.MatchCapabilities
		*out = make([]MatchHostCapabilitiesType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostCapabilitiesType.
func (in *HostCapabilitiesType) DeepCopy() *HostCapabilitiesType {
	if in == nil {
		return nil
	}
	out := new(HostCapabilitiesType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostNetworkType) DeepCopyInto(out *HostNetworkType) {
	*out = *in
	if in.MatchProtocols != nil {
		in, out := &in.MatchProtocols, &out.MatchProtocols
		*out = make([]MatchHostNetworkProtocolType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostNetworkType.
func (in *HostNetworkType) DeepCopy() *HostNetworkType {
	if in == nil {
		return nil
	}
	out := new(HostNetworkType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorHostPolicy) DeepCopyInto(out *KubeArmorHostPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorHostPolicy.
func (in *KubeArmorHostPolicy) DeepCopy() *KubeArmorHostPolicy {
	if in == nil {
		return nil
	}
	out := new(KubeArmorHostPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeArmorHostPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorHostPolicyList) DeepCopyInto(out *KubeArmorHostPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeArmorHostPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorHostPolicyList.
func (in *KubeArmorHostPolicyList) DeepCopy() *KubeArmorHostPolicyList {
	if in == nil {
		return nil
	}
	out := new(KubeArmorHostPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeArmorHostPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorHostPolicySpec) DeepCopyInto(out *KubeArmorHostPolicySpec) {
	*out = *in
	in.NodeSelector.DeepCopyInto(&out.NodeSelector)
	in.Process.DeepCopyInto(&out.Process)
	in.File.DeepCopyInto(&out.File)
	in.Network.DeepCopyInto(&out.Network)
	in.Capabilities.DeepCopyInto(&out.Capabilities)
	in.Syscalls.DeepCopyInto(&out.Syscalls)
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorHostPolicySpec.
func (in *KubeArmorHostPolicySpec) DeepCopy() *KubeArmorHostPolicySpec {
	if in == nil {
		return nil
	}
	out := new(KubeArmorHostPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorHostPolicyStatus) DeepCopyInto(out *KubeArmorHostPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorHostPolicyStatus.
func (in *KubeArmorHostPolicyStatus) DeepCopy() *KubeArmorHostPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(KubeArmorHostPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorPolicy) DeepCopyInto(out *KubeArmorPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorPolicy.
func (in *KubeArmorPolicy) DeepCopy() *KubeArmorPolicy {
	if in == nil {
		return nil
	}
	out := new(KubeArmorPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeArmorPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorPolicyList) DeepCopyInto(out *KubeArmorPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeArmorPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorPolicyList.
func (in *KubeArmorPolicyList) DeepCopy() *KubeArmorPolicyList {
	if in == nil {
		return nil
	}
	out := new(KubeArmorPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeArmorPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorPolicySpec) DeepCopyInto(out *KubeArmorPolicySpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.Process.DeepCopyInto(&out.Process)
	in.File.DeepCopyInto(&out.File)
	in.Network.DeepCopyInto(&out.Network)
	in.Capabilities.DeepCopyInto(&out.Capabilities)
	in.Syscalls.DeepCopyInto(&out.Syscalls)
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorPolicySpec.
func (in *KubeArmorPolicySpec) DeepCopy() *KubeArmorPolicySpec {
	if in == nil {
		return nil
	}
	out := new(KubeArmorPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorPolicyStatus) DeepCopyInto(out *KubeArmorPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorPolicyStatus.
func (in *KubeArmorPolicyStatus) DeepCopy() *KubeArmorPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(KubeArmorPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchCapabilitiesType) DeepCopyInto(out *MatchCapabilitiesType) {
	*out = *in
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchCapabilitiesType.
func (in *MatchCapabilitiesType) DeepCopy() *MatchCapabilitiesType {
	if in == nil {
		return nil
	}
	out := new(MatchCapabilitiesType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchHostCapabilitiesType) DeepCopyInto(out *MatchHostCapabilitiesType) {
	*out = *in
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchHostCapabilitiesType.
func (in *MatchHostCapabilitiesType) DeepCopy() *MatchHostCapabilitiesType {
	if in == nil {
		return nil
	}
	out := new(MatchHostCapabilitiesType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchHostNetworkProtocolType) DeepCopyInto(out *MatchHostNetworkProtocolType) {
	*out = *in
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchHostNetworkProtocolType.
func (in *MatchHostNetworkProtocolType) DeepCopy() *MatchHostNetworkProtocolType {
	if in == nil {
		return nil
	}
	out := new(MatchHostNetworkProtocolType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchNetworkProtocolType) DeepCopyInto(out *MatchNetworkProtocolType) {
	*out = *in
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchNetworkProtocolType.
func (in *MatchNetworkProtocolType) DeepCopy() *MatchNetworkProtocolType {
	if in == nil {
		return nil
	}
	out := new(MatchNetworkProtocolType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchSourceType) DeepCopyInto(out *MatchSourceType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchSourceType.
func (in *MatchSourceType) DeepCopy() *MatchSourceType {
	if in == nil {
		return nil
	}
	out := new(MatchSourceType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchVolumeMountType) DeepCopyInto(out *MatchVolumeMountType) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchVolumeMountType.
func (in *MatchVolumeMountType) DeepCopy() *MatchVolumeMountType {
	if in == nil {
		return nil
	}
	out := new(MatchVolumeMountType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkType) DeepCopyInto(out *NetworkType) {
	*out = *in
	if in.MatchProtocols != nil {
		in, out := &in.MatchProtocols, &out.MatchProtocols
		*out = make([]MatchNetworkProtocolType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkType.
func (in *NetworkType) DeepCopy() *NetworkType {
	if in == nil {
		return nil
	}
	out := new(NetworkType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSelectorType) DeepCopyInto(out *NodeSelectorType) {
	*out = *in
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSelectorType.
func (in *NodeSelectorType) DeepCopy() *NodeSelectorType {
	if in == nil {
		return nil
	}
	out := new(NodeSelectorType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessDirectoryType) DeepCopyInto(out *ProcessDirectoryType) {
	*out = *in
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessDirectoryType.
func (in *ProcessDirectoryType) DeepCopy() *ProcessDirectoryType {
	if in == nil {
		return nil
	}
	out := new(ProcessDirectoryType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessPathType) DeepCopyInto(out *ProcessPathType) {
	*out = *in
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessPathType.
func (in *ProcessPathType) DeepCopy() *ProcessPathType {
	if in == nil {
		return nil
	}
	out := new(ProcessPathType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessPatternType) DeepCopyInto(out *ProcessPatternType) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessPatternType.
func (in *ProcessPatternType) DeepCopy() *ProcessPatternType {
	if in == nil {
		return nil
	}
	out := new(ProcessPatternType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessType) DeepCopyInto(out *ProcessType) {
	*out = *in
	if in.MatchPaths != nil {
		in, out := &in.MatchPaths, &out.MatchPaths
		*out = make([]ProcessPathType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchDirectories != nil {
		in, out := &in.MatchDirectories, &out.MatchDirectories
		*out = make([]ProcessDirectoryType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchPatterns != nil {
		in, out := &in.MatchPatterns, &out.MatchPatterns
		*out = make([]ProcessPatternType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessType.
func (in *ProcessType) DeepCopy() *ProcessType {
	if in == nil {
		return nil
	}
	out := new(ProcessType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SELinuxType) DeepCopyInto(out *SELinuxType) {
	*out = *in
	if in.MatchVolumeMounts != nil {
		in, out := &in.MatchVolumeMounts, &out.MatchVolumeMounts
		*out = make([]MatchVolumeMountType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SELinuxType.
func (in *SELinuxType) DeepCopy() *SELinuxType {
	if in == nil {
		return nil
	}
	out := new(SELinuxType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectorType) DeepCopyInto(out *SelectorType) {
	*out = *in
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectorType.
func (in *SelectorType) DeepCopy() *SelectorType {
	if in == nil {
		return nil
	}
	out := new(SelectorType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyscallFromSourceType) DeepCopyInto(out *SyscallFromSourceType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyscallFromSourceType.
func (in *SyscallFromSourceType) DeepCopy() *SyscallFromSourceType {
	if in == nil {
		return nil
	}
	out := new(SyscallFromSourceType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyscallMatchPathType) DeepCopyInto(out *SyscallMatchPathType) {
	*out = *in
	if in.Syscalls != nil {
		in, out := &in.Syscalls, &out.Syscalls
		*out = make([]syscall, len(*in))
		copy(*out, *in)
	}
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]SyscallFromSourceType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyscallMatchPathType.
func (in *SyscallMatchPathType) DeepCopy() *SyscallMatchPathType {
	if in == nil {
		return nil
	}
	out := new(SyscallMatchPathType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyscallMatchType) DeepCopyInto(out *SyscallMatchType) {
	*out = *in
	if in.Syscalls != nil {
		in, out := &in.Syscalls, &out.Syscalls
		*out = make([]syscall, len(*in))
		copy(*out, *in)
	}
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]SyscallFromSourceType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyscallMatchType.
func (in *SyscallMatchType) DeepCopy() *SyscallMatchType {
	if in == nil {
		return nil
	}
	out := new(SyscallMatchType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyscallsType) DeepCopyInto(out *SyscallsType) {
	*out = *in
	if in.MatchSyscalls != nil {
		in, out := &in.MatchSyscalls, &out.MatchSyscalls
		*out = make([]SyscallMatchType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchPaths != nil {
		in, out := &in.MatchPaths, &out.MatchPaths
		*out = make([]SyscallMatchPathType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyscallsType.
func (in *SyscallsType) DeepCopy() *SyscallsType {
	if in == nil {
		return nil
	}
	out := new(SyscallsType)
	in.DeepCopyInto(out)
	return out
}
