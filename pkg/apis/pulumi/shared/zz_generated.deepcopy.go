//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package shared

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicAuth) DeepCopyInto(out *BasicAuth) {
	*out = *in
	in.UserName.DeepCopyInto(&out.UserName)
	in.Password.DeepCopyInto(&out.Password)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicAuth.
func (in *BasicAuth) DeepCopy() *BasicAuth {
	if in == nil {
		return nil
	}
	out := new(BasicAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvSelector) DeepCopyInto(out *EnvSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvSelector.
func (in *EnvSelector) DeepCopy() *EnvSelector {
	if in == nil {
		return nil
	}
	out := new(EnvSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FSSelector) DeepCopyInto(out *FSSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FSSelector.
func (in *FSSelector) DeepCopy() *FSSelector {
	if in == nil {
		return nil
	}
	out := new(FSSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitAuthConfig) DeepCopyInto(out *GitAuthConfig) {
	*out = *in
	if in.PersonalAccessToken != nil {
		in, out := &in.PersonalAccessToken, &out.PersonalAccessToken
		*out = new(ResourceRef)
		(*in).DeepCopyInto(*out)
	}
	if in.SSHAuth != nil {
		in, out := &in.SSHAuth, &out.SSHAuth
		*out = new(SSHAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		*out = new(BasicAuth)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitAuthConfig.
func (in *GitAuthConfig) DeepCopy() *GitAuthConfig {
	if in == nil {
		return nil
	}
	out := new(GitAuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LiteralRef) DeepCopyInto(out *LiteralRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LiteralRef.
func (in *LiteralRef) DeepCopy() *LiteralRef {
	if in == nil {
		return nil
	}
	out := new(LiteralRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRef) DeepCopyInto(out *ResourceRef) {
	*out = *in
	in.ResourceSelector.DeepCopyInto(&out.ResourceSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRef.
func (in *ResourceRef) DeepCopy() *ResourceRef {
	if in == nil {
		return nil
	}
	out := new(ResourceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSelector) DeepCopyInto(out *ResourceSelector) {
	*out = *in
	if in.FileSystem != nil {
		in, out := &in.FileSystem, &out.FileSystem
		*out = new(FSSelector)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = new(EnvSelector)
		**out = **in
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(SecretSelector)
		**out = **in
	}
	if in.LiteralRef != nil {
		in, out := &in.LiteralRef, &out.LiteralRef
		*out = new(LiteralRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSelector.
func (in *ResourceSelector) DeepCopy() *ResourceSelector {
	if in == nil {
		return nil
	}
	out := new(ResourceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHAuth) DeepCopyInto(out *SSHAuth) {
	*out = *in
	in.SSHPrivateKey.DeepCopyInto(&out.SSHPrivateKey)
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(ResourceRef)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHAuth.
func (in *SSHAuth) DeepCopy() *SSHAuth {
	if in == nil {
		return nil
	}
	out := new(SSHAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretSelector) DeepCopyInto(out *SecretSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretSelector.
func (in *SecretSelector) DeepCopy() *SecretSelector {
	if in == nil {
		return nil
	}
	out := new(SecretSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in StackOutputs) DeepCopyInto(out *StackOutputs) {
	{
		in := &in
		*out = make(StackOutputs, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackOutputs.
func (in StackOutputs) DeepCopy() StackOutputs {
	if in == nil {
		return nil
	}
	out := new(StackOutputs)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackSpec) DeepCopyInto(out *StackSpec) {
	*out = *in
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EnvRefs != nil {
		in, out := &in.EnvRefs, &out.EnvRefs
		*out = make(map[string]ResourceRef, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.SecretEnvs != nil {
		in, out := &in.SecretEnvs, &out.SecretEnvs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecretRefs != nil {
		in, out := &in.SecretRefs, &out.SecretRefs
		*out = make(map[string]ResourceRef, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.GitAuth != nil {
		in, out := &in.GitAuth, &out.GitAuth
		*out = new(GitAuthConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackSpec.
func (in *StackSpec) DeepCopy() *StackSpec {
	if in == nil {
		return nil
	}
	out := new(StackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackStatus) DeepCopyInto(out *StackStatus) {
	*out = *in
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make(StackOutputs, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.LastUpdate != nil {
		in, out := &in.LastUpdate, &out.LastUpdate
		*out = new(StackUpdateState)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackStatus.
func (in *StackStatus) DeepCopy() *StackStatus {
	if in == nil {
		return nil
	}
	out := new(StackStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackUpdateState) DeepCopyInto(out *StackUpdateState) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackUpdateState.
func (in *StackUpdateState) DeepCopy() *StackUpdateState {
	if in == nil {
		return nil
	}
	out := new(StackUpdateState)
	in.DeepCopyInto(out)
	return out
}