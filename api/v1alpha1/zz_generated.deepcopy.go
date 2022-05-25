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

package v1alpha1

import (
	"github.com/fluxcd/pkg/apis/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Change) DeepCopyInto(out *Change) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Change.
func (in *Change) DeepCopy() *Change {
	if in == nil {
		return nil
	}
	out := new(Change)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommandResult) DeepCopyInto(out *CommandResult) {
	*out = *in
	if in.NewObjects != nil {
		in, out := &in.NewObjects, &out.NewObjects
		*out = make([]ResourceRef, len(*in))
		copy(*out, *in)
	}
	if in.ChangedObjects != nil {
		in, out := &in.ChangedObjects, &out.ChangedObjects
		*out = make([]ResourceRef, len(*in))
		copy(*out, *in)
	}
	if in.HookObjects != nil {
		in, out := &in.HookObjects, &out.HookObjects
		*out = make([]ResourceRef, len(*in))
		copy(*out, *in)
	}
	if in.OrphanObjects != nil {
		in, out := &in.OrphanObjects, &out.OrphanObjects
		*out = make([]ResourceRef, len(*in))
		copy(*out, *in)
	}
	if in.DeletedObjects != nil {
		in, out := &in.DeletedObjects, &out.DeletedObjects
		*out = make([]ResourceRef, len(*in))
		copy(*out, *in)
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]DeploymentError, len(*in))
		copy(*out, *in)
	}
	if in.Warnings != nil {
		in, out := &in.Warnings, &out.Warnings
		*out = make([]DeploymentError, len(*in))
		copy(*out, *in)
	}
	if in.SeenImages != nil {
		in, out := &in.SeenImages, &out.SeenImages
		*out = make([]FixedImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommandResult.
func (in *CommandResult) DeepCopy() *CommandResult {
	if in == nil {
		return nil
	}
	out := new(CommandResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrossNamespaceSourceReference) DeepCopyInto(out *CrossNamespaceSourceReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrossNamespaceSourceReference.
func (in *CrossNamespaceSourceReference) DeepCopy() *CrossNamespaceSourceReference {
	if in == nil {
		return nil
	}
	out := new(CrossNamespaceSourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentError) DeepCopyInto(out *DeploymentError) {
	*out = *in
	out.Ref = in.Ref
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentError.
func (in *DeploymentError) DeepCopy() *DeploymentError {
	if in == nil {
		return nil
	}
	out := new(DeploymentError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FixedImage) DeepCopyInto(out *FixedImage) {
	*out = *in
	if in.DeployedImage != nil {
		in, out := &in.DeployedImage, &out.DeployedImage
		*out = new(string)
		**out = **in
	}
	if in.RegistryImage != nil {
		in, out := &in.RegistryImage, &out.RegistryImage
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(ResourceRef)
		**out = **in
	}
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(string)
		**out = **in
	}
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(string)
		**out = **in
	}
	if in.VersionFilter != nil {
		in, out := &in.VersionFilter, &out.VersionFilter
		*out = new(string)
		**out = **in
	}
	if in.DeployTags != nil {
		in, out := &in.DeployTags, &out.DeployTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DeploymentDir != nil {
		in, out := &in.DeploymentDir, &out.DeploymentDir
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FixedImage.
func (in *FixedImage) DeepCopy() *FixedImage {
	if in == nil {
		return nil
	}
	out := new(FixedImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KluctlDeployment) DeepCopyInto(out *KluctlDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KluctlDeployment.
func (in *KluctlDeployment) DeepCopy() *KluctlDeployment {
	if in == nil {
		return nil
	}
	out := new(KluctlDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KluctlDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KluctlDeploymentList) DeepCopyInto(out *KluctlDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KluctlDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KluctlDeploymentList.
func (in *KluctlDeploymentList) DeepCopy() *KluctlDeploymentList {
	if in == nil {
		return nil
	}
	out := new(KluctlDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KluctlDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KluctlDeploymentSpec) DeepCopyInto(out *KluctlDeploymentSpec) {
	*out = *in
	if in.DependsOn != nil {
		in, out := &in.DependsOn, &out.DependsOn
		*out = make([]meta.NamespacedObjectReference, len(*in))
		copy(*out, *in)
	}
	out.Interval = in.Interval
	if in.RetryInterval != nil {
		in, out := &in.RetryInterval, &out.RetryInterval
		*out = new(v1.Duration)
		**out = **in
	}
	out.SourceRef = in.SourceRef
	if in.RegistrySecrets != nil {
		in, out := &in.RegistrySecrets, &out.RegistrySecrets
		*out = make([]meta.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.KubeConfig != nil {
		in, out := &in.KubeConfig, &out.KubeConfig
		*out = new(KubeConfig)
		**out = **in
	}
	if in.RenameContexts != nil {
		in, out := &in.RenameContexts, &out.RenameContexts
		*out = make([]RenameContext, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]FixedImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IncludeTags != nil {
		in, out := &in.IncludeTags, &out.IncludeTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeTags != nil {
		in, out := &in.ExcludeTags, &out.ExcludeTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncludeDeploymentDirs != nil {
		in, out := &in.IncludeDeploymentDirs, &out.IncludeDeploymentDirs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeDeploymentDirs != nil {
		in, out := &in.ExcludeDeploymentDirs, &out.ExcludeDeploymentDirs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KluctlDeploymentSpec.
func (in *KluctlDeploymentSpec) DeepCopy() *KluctlDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(KluctlDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KluctlDeploymentStatus) DeepCopyInto(out *KluctlDeploymentStatus) {
	*out = *in
	out.ReconcileRequestStatus = in.ReconcileRequestStatus
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastAttemptedReconcile != nil {
		in, out := &in.LastAttemptedReconcile, &out.LastAttemptedReconcile
		*out = new(ReconcileAttempt)
		(*in).DeepCopyInto(*out)
	}
	if in.LastSuccessfulReconcile != nil {
		in, out := &in.LastSuccessfulReconcile, &out.LastSuccessfulReconcile
		*out = new(ReconcileAttempt)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KluctlDeploymentStatus.
func (in *KluctlDeploymentStatus) DeepCopy() *KluctlDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(KluctlDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeConfig) DeepCopyInto(out *KubeConfig) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeConfig.
func (in *KubeConfig) DeepCopy() *KubeConfig {
	if in == nil {
		return nil
	}
	out := new(KubeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReconcileAttempt) DeepCopyInto(out *ReconcileAttempt) {
	*out = *in
	in.AttemptedAt.DeepCopyInto(&out.AttemptedAt)
	if in.DeployResult != nil {
		in, out := &in.DeployResult, &out.DeployResult
		*out = new(CommandResult)
		(*in).DeepCopyInto(*out)
	}
	if in.PruneResult != nil {
		in, out := &in.PruneResult, &out.PruneResult
		*out = new(CommandResult)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReconcileAttempt.
func (in *ReconcileAttempt) DeepCopy() *ReconcileAttempt {
	if in == nil {
		return nil
	}
	out := new(ReconcileAttempt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RenameContext) DeepCopyInto(out *RenameContext) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RenameContext.
func (in *RenameContext) DeepCopy() *RenameContext {
	if in == nil {
		return nil
	}
	out := new(RenameContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRef) DeepCopyInto(out *ResourceRef) {
	*out = *in
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
func (in *ValidateResult) DeepCopyInto(out *ValidateResult) {
	*out = *in
	if in.Warnings != nil {
		in, out := &in.Warnings, &out.Warnings
		*out = make([]DeploymentError, len(*in))
		copy(*out, *in)
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]DeploymentError, len(*in))
		copy(*out, *in)
	}
	if in.Results != nil {
		in, out := &in.Results, &out.Results
		*out = make([]ValidateResultEntry, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidateResult.
func (in *ValidateResult) DeepCopy() *ValidateResult {
	if in == nil {
		return nil
	}
	out := new(ValidateResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidateResultEntry) DeepCopyInto(out *ValidateResultEntry) {
	*out = *in
	out.Ref = in.Ref
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidateResultEntry.
func (in *ValidateResultEntry) DeepCopy() *ValidateResultEntry {
	if in == nil {
		return nil
	}
	out := new(ValidateResultEntry)
	in.DeepCopyInto(out)
	return out
}
