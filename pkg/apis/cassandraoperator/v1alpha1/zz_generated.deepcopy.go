// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraAuth) DeepCopyInto(out *CassandraAuth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraAuth.
func (in *CassandraAuth) DeepCopy() *CassandraAuth {
	if in == nil {
		return nil
	}
	out := new(CassandraAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraBackup) DeepCopyInto(out *CassandraBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = make([]*CassandraBackupStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(CassandraBackupStatus)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraBackup.
func (in *CassandraBackup) DeepCopy() *CassandraBackup {
	if in == nil {
		return nil
	}
	out := new(CassandraBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraBackupList) DeepCopyInto(out *CassandraBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CassandraBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraBackupList.
func (in *CassandraBackupList) DeepCopy() *CassandraBackupList {
	if in == nil {
		return nil
	}
	out := new(CassandraBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraBackupSpec) DeepCopyInto(out *CassandraBackupSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraBackupSpec.
func (in *CassandraBackupSpec) DeepCopy() *CassandraBackupSpec {
	if in == nil {
		return nil
	}
	out := new(CassandraBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraBackupStatus) DeepCopyInto(out *CassandraBackupStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraBackupStatus.
func (in *CassandraBackupStatus) DeepCopy() *CassandraBackupStatus {
	if in == nil {
		return nil
	}
	out := new(CassandraBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraCluster) DeepCopyInto(out *CassandraCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraCluster.
func (in *CassandraCluster) DeepCopy() *CassandraCluster {
	if in == nil {
		return nil
	}
	out := new(CassandraCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraClusterList) DeepCopyInto(out *CassandraClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CassandraCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraClusterList.
func (in *CassandraClusterList) DeepCopy() *CassandraClusterList {
	if in == nil {
		return nil
	}
	out := new(CassandraClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraClusterSpec) DeepCopyInto(out *CassandraClusterSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraClusterSpec.
func (in *CassandraClusterSpec) DeepCopy() *CassandraClusterSpec {
	if in == nil {
		return nil
	}
	out := new(CassandraClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraClusterStatus) DeepCopyInto(out *CassandraClusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraClusterStatus.
func (in *CassandraClusterStatus) DeepCopy() *CassandraClusterStatus {
	if in == nil {
		return nil
	}
	out := new(CassandraClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraDataCenter) DeepCopyInto(out *CassandraDataCenter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraDataCenter.
func (in *CassandraDataCenter) DeepCopy() *CassandraDataCenter {
	if in == nil {
		return nil
	}
	out := new(CassandraDataCenter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraDataCenter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraDataCenterList) DeepCopyInto(out *CassandraDataCenterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CassandraDataCenter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraDataCenterList.
func (in *CassandraDataCenterList) DeepCopy() *CassandraDataCenterList {
	if in == nil {
		return nil
	}
	out := new(CassandraDataCenterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraDataCenterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraDataCenterSpec) DeepCopyInto(out *CassandraDataCenterSpec) {
	*out = *in
	if in.Racks != nil {
		in, out := &in.Racks, &out.Racks
		*out = make([]Rack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.UserSecretVolumeSource != nil {
		in, out := &in.UserSecretVolumeSource, &out.UserSecretVolumeSource
		*out = make([]*v1.SecretVolumeSource, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.SecretVolumeSource)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.UserConfigMapVolumeSource != nil {
		in, out := &in.UserConfigMapVolumeSource, &out.UserConfigMapVolumeSource
		*out = new(v1.ConfigMapVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.SidecarSecretVolumeSource != nil {
		in, out := &in.SidecarSecretVolumeSource, &out.SidecarSecretVolumeSource
		*out = new(v1.SecretVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.SidecarResources != nil {
		in, out := &in.SidecarResources, &out.SidecarResources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.DummyVolume != nil {
		in, out := &in.DummyVolume, &out.DummyVolume
		*out = new(v1.EmptyDirVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.DataVolumeClaimSpec != nil {
		in, out := &in.DataVolumeClaimSpec, &out.DataVolumeClaimSpec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.OperatorLabels != nil {
		in, out := &in.OperatorLabels, &out.OperatorLabels
		*out = new(OperatorLabels)
		(*in).DeepCopyInto(*out)
	}
	if in.OperatorAnnotations != nil {
		in, out := &in.OperatorAnnotations, &out.OperatorAnnotations
		*out = new(OperatorAnnotations)
		(*in).DeepCopyInto(*out)
	}
	if in.SidecarEnv != nil {
		in, out := &in.SidecarEnv, &out.SidecarEnv
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CassandraAuth != nil {
		in, out := &in.CassandraAuth, &out.CassandraAuth
		*out = new(CassandraAuth)
		**out = **in
	}
	if in.CassandraEnv != nil {
		in, out := &in.CassandraEnv, &out.CassandraEnv
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Restore != nil {
		in, out := &in.Restore, &out.Restore
		*out = new(Restore)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraDataCenterSpec.
func (in *CassandraDataCenterSpec) DeepCopy() *CassandraDataCenterSpec {
	if in == nil {
		return nil
	}
	out := new(CassandraDataCenterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraDataCenterStatus) DeepCopyInto(out *CassandraDataCenterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraDataCenterStatus.
func (in *CassandraDataCenterStatus) DeepCopy() *CassandraDataCenterStatus {
	if in == nil {
		return nil
	}
	out := new(CassandraDataCenterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorAnnotations) DeepCopyInto(out *OperatorAnnotations) {
	*out = *in
	if in.PrometheusService != nil {
		in, out := &in.PrometheusService, &out.PrometheusService
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodesService != nil {
		in, out := &in.NodesService, &out.NodesService
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SeedNodesService != nil {
		in, out := &in.SeedNodesService, &out.SeedNodesService
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StatefulSet != nil {
		in, out := &in.StatefulSet, &out.StatefulSet
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodTemplate != nil {
		in, out := &in.PodTemplate, &out.PodTemplate
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorAnnotations.
func (in *OperatorAnnotations) DeepCopy() *OperatorAnnotations {
	if in == nil {
		return nil
	}
	out := new(OperatorAnnotations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorLabels) DeepCopyInto(out *OperatorLabels) {
	*out = *in
	if in.PrometheusService != nil {
		in, out := &in.PrometheusService, &out.PrometheusService
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodesService != nil {
		in, out := &in.NodesService, &out.NodesService
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SeedNodesService != nil {
		in, out := &in.SeedNodesService, &out.SeedNodesService
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StatefulSet != nil {
		in, out := &in.StatefulSet, &out.StatefulSet
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodTemplate != nil {
		in, out := &in.PodTemplate, &out.PodTemplate
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorLabels.
func (in *OperatorLabels) DeepCopy() *OperatorLabels {
	if in == nil {
		return nil
	}
	out := new(OperatorLabels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rack) DeepCopyInto(out *Rack) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rack.
func (in *Rack) DeepCopy() *Rack {
	if in == nil {
		return nil
	}
	out := new(Rack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Restore) DeepCopyInto(out *Restore) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Restore.
func (in *Restore) DeepCopy() *Restore {
	if in == nil {
		return nil
	}
	out := new(Restore)
	in.DeepCopyInto(out)
	return out
}
