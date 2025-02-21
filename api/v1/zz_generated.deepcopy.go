// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentConfig) DeepCopyInto(out *AgentConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentConfig.
func (in *AgentConfig) DeepCopy() *AgentConfig {
	if in == nil {
		return nil
	}
	out := new(AgentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentConfigList) DeepCopyInto(out *AgentConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AgentConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentConfigList.
func (in *AgentConfigList) DeepCopy() *AgentConfigList {
	if in == nil {
		return nil
	}
	out := new(AgentConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentConfigSpec) DeepCopyInto(out *AgentConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentConfigSpec.
func (in *AgentConfigSpec) DeepCopy() *AgentConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AgentConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Installation) DeepCopyInto(out *Installation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Installation.
func (in *Installation) DeepCopy() *Installation {
	if in == nil {
		return nil
	}
	out := new(Installation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Installation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationList) DeepCopyInto(out *InstallationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Installation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationList.
func (in *InstallationList) DeepCopy() *InstallationList {
	if in == nil {
		return nil
	}
	out := new(InstallationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstallationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationSpec) DeepCopyInto(out *InstallationSpec) {
	*out = *in
	out.AgentConfig = in.AgentConfig
	out.PorterConfig = in.PorterConfig
	out.Bundle = in.Bundle
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Parameters.DeepCopyInto(&out.Parameters)
	if in.CredentialSets != nil {
		in, out := &in.CredentialSets, &out.CredentialSets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ParameterSets != nil {
		in, out := &in.ParameterSets, &out.ParameterSets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationSpec.
func (in *InstallationSpec) DeepCopy() *InstallationSpec {
	if in == nil {
		return nil
	}
	out := new(InstallationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationStatus) DeepCopyInto(out *InstallationStatus) {
	*out = *in
	out.ActiveJob = in.ActiveJob
	out.LastJob = in.LastJob
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationStatus.
func (in *InstallationStatus) DeepCopy() *InstallationStatus {
	if in == nil {
		return nil
	}
	out := new(InstallationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIReferenceParts) DeepCopyInto(out *OCIReferenceParts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIReferenceParts.
func (in *OCIReferenceParts) DeepCopy() *OCIReferenceParts {
	if in == nil {
		return nil
	}
	out := new(OCIReferenceParts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginConfig) DeepCopyInto(out *PluginConfig) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginConfig.
func (in *PluginConfig) DeepCopy() *PluginConfig {
	if in == nil {
		return nil
	}
	out := new(PluginConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PorterConfig) DeepCopyInto(out *PorterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PorterConfig.
func (in *PorterConfig) DeepCopy() *PorterConfig {
	if in == nil {
		return nil
	}
	out := new(PorterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PorterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PorterConfigList) DeepCopyInto(out *PorterConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PorterConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PorterConfigList.
func (in *PorterConfigList) DeepCopy() *PorterConfigList {
	if in == nil {
		return nil
	}
	out := new(PorterConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PorterConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PorterConfigSpec) DeepCopyInto(out *PorterConfigSpec) {
	*out = *in
	if in.Debug != nil {
		in, out := &in.Debug, &out.Debug
		*out = new(bool)
		**out = **in
	}
	if in.DebugPlugins != nil {
		in, out := &in.DebugPlugins, &out.DebugPlugins
		*out = new(bool)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Experimental != nil {
		in, out := &in.Experimental, &out.Experimental
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BuildDriver != nil {
		in, out := &in.BuildDriver, &out.BuildDriver
		*out = new(string)
		**out = **in
	}
	if in.DefaultStorage != nil {
		in, out := &in.DefaultStorage, &out.DefaultStorage
		*out = new(string)
		**out = **in
	}
	if in.DefaultSecrets != nil {
		in, out := &in.DefaultSecrets, &out.DefaultSecrets
		*out = new(string)
		**out = **in
	}
	if in.DefaultStoragePlugin != nil {
		in, out := &in.DefaultStoragePlugin, &out.DefaultStoragePlugin
		*out = new(string)
		**out = **in
	}
	if in.DefaultSecretsPlugin != nil {
		in, out := &in.DefaultSecretsPlugin, &out.DefaultSecretsPlugin
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make([]StorageConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]SecretsConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PorterConfigSpec.
func (in *PorterConfigSpec) DeepCopy() *PorterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(PorterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretsConfig) DeepCopyInto(out *SecretsConfig) {
	*out = *in
	in.PluginConfig.DeepCopyInto(&out.PluginConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretsConfig.
func (in *SecretsConfig) DeepCopy() *SecretsConfig {
	if in == nil {
		return nil
	}
	out := new(SecretsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageConfig) DeepCopyInto(out *StorageConfig) {
	*out = *in
	in.PluginConfig.DeepCopyInto(&out.PluginConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageConfig.
func (in *StorageConfig) DeepCopy() *StorageConfig {
	if in == nil {
		return nil
	}
	out := new(StorageConfig)
	in.DeepCopyInto(out)
	return out
}
