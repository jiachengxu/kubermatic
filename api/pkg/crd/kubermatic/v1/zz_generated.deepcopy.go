// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSCloudSpec) DeepCopyInto(out *AWSCloudSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSCloudSpec.
func (in *AWSCloudSpec) DeepCopy() *AWSCloudSpec {
	if in == nil {
		return nil
	}
	out := new(AWSCloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addon) DeepCopyInto(out *Addon) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addon.
func (in *Addon) DeepCopy() *Addon {
	if in == nil {
		return nil
	}
	out := new(Addon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Addon) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonList) DeepCopyInto(out *AddonList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Addon, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonList.
func (in *AddonList) DeepCopy() *AddonList {
	if in == nil {
		return nil
	}
	out := new(AddonList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddonList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonSpec) DeepCopyInto(out *AddonSpec) {
	*out = *in
	out.Cluster = in.Cluster
	in.Variables.DeepCopyInto(&out.Variables)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonSpec.
func (in *AddonSpec) DeepCopy() *AddonSpec {
	if in == nil {
		return nil
	}
	out := new(AddonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureCloudSpec) DeepCopyInto(out *AzureCloudSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureCloudSpec.
func (in *AzureCloudSpec) DeepCopy() *AzureCloudSpec {
	if in == nil {
		return nil
	}
	out := new(AzureCloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BringYourOwnCloudSpec) DeepCopyInto(out *BringYourOwnCloudSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BringYourOwnCloudSpec.
func (in *BringYourOwnCloudSpec) DeepCopy() *BringYourOwnCloudSpec {
	if in == nil {
		return nil
	}
	out := new(BringYourOwnCloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Bytes) DeepCopyInto(out *Bytes) {
	{
		in := &in
		*out = make(Bytes, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bytes.
func (in Bytes) DeepCopy() Bytes {
	if in == nil {
		return nil
	}
	out := new(Bytes)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSpec) DeepCopyInto(out *CloudSpec) {
	*out = *in
	if in.Fake != nil {
		in, out := &in.Fake, &out.Fake
		*out = new(FakeCloudSpec)
		**out = **in
	}
	if in.Digitalocean != nil {
		in, out := &in.Digitalocean, &out.Digitalocean
		*out = new(DigitaloceanCloudSpec)
		**out = **in
	}
	if in.BringYourOwn != nil {
		in, out := &in.BringYourOwn, &out.BringYourOwn
		*out = new(BringYourOwnCloudSpec)
		**out = **in
	}
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSCloudSpec)
		**out = **in
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(AzureCloudSpec)
		**out = **in
	}
	if in.Openstack != nil {
		in, out := &in.Openstack, &out.Openstack
		*out = new(OpenstackCloudSpec)
		**out = **in
	}
	if in.Packet != nil {
		in, out := &in.Packet, &out.Packet
		*out = new(PacketCloudSpec)
		**out = **in
	}
	if in.Hetzner != nil {
		in, out := &in.Hetzner, &out.Hetzner
		*out = new(HetznerCloudSpec)
		**out = **in
	}
	if in.VSphere != nil {
		in, out := &in.VSphere, &out.VSphere
		*out = new(VSphereCloudSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSpec.
func (in *CloudSpec) DeepCopy() *CloudSpec {
	if in == nil {
		return nil
	}
	out := new(CloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Address = in.Address
	in.Status.DeepCopyInto(&out.Status)
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

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAddress) DeepCopyInto(out *ClusterAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAddress.
func (in *ClusterAddress) DeepCopy() *ClusterAddress {
	if in == nil {
		return nil
	}
	out := new(ClusterAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCondition) DeepCopyInto(out *ClusterCondition) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCondition.
func (in *ClusterCondition) DeepCopy() *ClusterCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterHealth) DeepCopyInto(out *ClusterHealth) {
	*out = *in
	out.ClusterHealthStatus = in.ClusterHealthStatus
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterHealth.
func (in *ClusterHealth) DeepCopy() *ClusterHealth {
	if in == nil {
		return nil
	}
	out := new(ClusterHealth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterHealthStatus) DeepCopyInto(out *ClusterHealthStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterHealthStatus.
func (in *ClusterHealthStatus) DeepCopy() *ClusterHealthStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterHealthStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNetworkingConfig) DeepCopyInto(out *ClusterNetworkingConfig) {
	*out = *in
	in.Services.DeepCopyInto(&out.Services)
	in.Pods.DeepCopyInto(&out.Pods)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNetworkingConfig.
func (in *ClusterNetworkingConfig) DeepCopy() *ClusterNetworkingConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterNetworkingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.Cloud.DeepCopyInto(&out.Cloud)
	in.ClusterNetwork.DeepCopyInto(&out.ClusterNetwork)
	if in.MachineNetworks != nil {
		in, out := &in.MachineNetworks, &out.MachineNetworks
		*out = make([]MachineNetworkingConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Version = in.Version.DeepCopy()
	in.ComponentsOverride.DeepCopyInto(&out.ComponentsOverride)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.LastUpdated.DeepCopyInto(&out.LastUpdated)
	in.Health.DeepCopyInto(&out.Health)
	if in.RootCA != nil {
		in, out := &in.RootCA, &out.RootCA
		*out = new(KeyCert)
		(*in).DeepCopyInto(*out)
	}
	if in.ApiserverCert != nil {
		in, out := &in.ApiserverCert, &out.ApiserverCert
		*out = new(KeyCert)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeletCert != nil {
		in, out := &in.KubeletCert, &out.KubeletCert
		*out = new(KeyCert)
		(*in).DeepCopyInto(*out)
	}
	if in.ApiserverSSHKey != nil {
		in, out := &in.ApiserverSSHKey, &out.ApiserverSSHKey
		*out = new(RSAKeys)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccountKey != nil {
		in, out := &in.ServiceAccountKey, &out.ServiceAccountKey
		*out = make(Bytes, len(*in))
		copy(*out, *in)
	}
	if in.ErrorReason != nil {
		in, out := &in.ErrorReason, &out.ErrorReason
		*out = new(ClusterStatusError)
		**out = **in
	}
	if in.ErrorMessage != nil {
		in, out := &in.ErrorMessage, &out.ErrorMessage
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
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
func (in *ComponentSettings) DeepCopyInto(out *ComponentSettings) {
	*out = *in
	in.Apiserver.DeepCopyInto(&out.Apiserver)
	in.ControllerManager.DeepCopyInto(&out.ControllerManager)
	in.Scheduler.DeepCopyInto(&out.Scheduler)
	in.Etcd.DeepCopyInto(&out.Etcd)
	in.Prometheus.DeepCopyInto(&out.Prometheus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentSettings.
func (in *ComponentSettings) DeepCopy() *ComponentSettings {
	if in == nil {
		return nil
	}
	out := new(ComponentSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentSettings) DeepCopyInto(out *DeploymentSettings) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentSettings.
func (in *DeploymentSettings) DeepCopy() *DeploymentSettings {
	if in == nil {
		return nil
	}
	out := new(DeploymentSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DigitaloceanCloudSpec) DeepCopyInto(out *DigitaloceanCloudSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DigitaloceanCloudSpec.
func (in *DigitaloceanCloudSpec) DeepCopy() *DigitaloceanCloudSpec {
	if in == nil {
		return nil
	}
	out := new(DigitaloceanCloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FakeCloudSpec) DeepCopyInto(out *FakeCloudSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FakeCloudSpec.
func (in *FakeCloudSpec) DeepCopy() *FakeCloudSpec {
	if in == nil {
		return nil
	}
	out := new(FakeCloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerCloudSpec) DeepCopyInto(out *HetznerCloudSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerCloudSpec.
func (in *HetznerCloudSpec) DeepCopy() *HetznerCloudSpec {
	if in == nil {
		return nil
	}
	out := new(HetznerCloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyCert) DeepCopyInto(out *KeyCert) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = make(Bytes, len(*in))
		copy(*out, *in)
	}
	if in.Cert != nil {
		in, out := &in.Cert, &out.Cert
		*out = make(Bytes, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyCert.
func (in *KeyCert) DeepCopy() *KeyCert {
	if in == nil {
		return nil
	}
	out := new(KeyCert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineNetworkingConfig) DeepCopyInto(out *MachineNetworkingConfig) {
	*out = *in
	if in.DNSServers != nil {
		in, out := &in.DNSServers, &out.DNSServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineNetworkingConfig.
func (in *MachineNetworkingConfig) DeepCopy() *MachineNetworkingConfig {
	if in == nil {
		return nil
	}
	out := new(MachineNetworkingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRanges) DeepCopyInto(out *NetworkRanges) {
	*out = *in
	if in.CIDRBlocks != nil {
		in, out := &in.CIDRBlocks, &out.CIDRBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRanges.
func (in *NetworkRanges) DeepCopy() *NetworkRanges {
	if in == nil {
		return nil
	}
	out := new(NetworkRanges)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackCloudSpec) DeepCopyInto(out *OpenstackCloudSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackCloudSpec.
func (in *OpenstackCloudSpec) DeepCopy() *OpenstackCloudSpec {
	if in == nil {
		return nil
	}
	out := new(OpenstackCloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PacketCloudSpec) DeepCopyInto(out *PacketCloudSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PacketCloudSpec.
func (in *PacketCloudSpec) DeepCopy() *PacketCloudSpec {
	if in == nil {
		return nil
	}
	out := new(PacketCloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Project) DeepCopyInto(out *Project) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Project.
func (in *Project) DeepCopy() *Project {
	if in == nil {
		return nil
	}
	out := new(Project)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Project) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectGroup) DeepCopyInto(out *ProjectGroup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectGroup.
func (in *ProjectGroup) DeepCopy() *ProjectGroup {
	if in == nil {
		return nil
	}
	out := new(ProjectGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectList) DeepCopyInto(out *ProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Project, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectList.
func (in *ProjectList) DeepCopy() *ProjectList {
	if in == nil {
		return nil
	}
	out := new(ProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSpec) DeepCopyInto(out *ProjectSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSpec.
func (in *ProjectSpec) DeepCopy() *ProjectSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectStatus) DeepCopyInto(out *ProjectStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectStatus.
func (in *ProjectStatus) DeepCopy() *ProjectStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RSAKeys) DeepCopyInto(out *RSAKeys) {
	*out = *in
	if in.PrivateKey != nil {
		in, out := &in.PrivateKey, &out.PrivateKey
		*out = make(Bytes, len(*in))
		copy(*out, *in)
	}
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = make(Bytes, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RSAKeys.
func (in *RSAKeys) DeepCopy() *RSAKeys {
	if in == nil {
		return nil
	}
	out := new(RSAKeys)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHKeySpec) DeepCopyInto(out *SSHKeySpec) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHKeySpec.
func (in *SSHKeySpec) DeepCopy() *SSHKeySpec {
	if in == nil {
		return nil
	}
	out := new(SSHKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetSettings) DeepCopyInto(out *StatefulSetSettings) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetSettings.
func (in *StatefulSetSettings) DeepCopy() *StatefulSetSettings {
	if in == nil {
		return nil
	}
	out := new(StatefulSetSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *User) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserList) DeepCopyInto(out *UserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]User, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserList.
func (in *UserList) DeepCopy() *UserList {
	if in == nil {
		return nil
	}
	out := new(UserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserProjectBinding) DeepCopyInto(out *UserProjectBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserProjectBinding.
func (in *UserProjectBinding) DeepCopy() *UserProjectBinding {
	if in == nil {
		return nil
	}
	out := new(UserProjectBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserProjectBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserProjectBindingList) DeepCopyInto(out *UserProjectBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserProjectBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserProjectBindingList.
func (in *UserProjectBindingList) DeepCopy() *UserProjectBindingList {
	if in == nil {
		return nil
	}
	out := new(UserProjectBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserProjectBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserProjectBindingSpec) DeepCopyInto(out *UserProjectBindingSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserProjectBindingSpec.
func (in *UserProjectBindingSpec) DeepCopy() *UserProjectBindingSpec {
	if in == nil {
		return nil
	}
	out := new(UserProjectBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSSHKey) DeepCopyInto(out *UserSSHKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSSHKey.
func (in *UserSSHKey) DeepCopy() *UserSSHKey {
	if in == nil {
		return nil
	}
	out := new(UserSSHKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserSSHKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSSHKeyList) DeepCopyInto(out *UserSSHKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserSSHKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSSHKeyList.
func (in *UserSSHKeyList) DeepCopy() *UserSSHKeyList {
	if in == nil {
		return nil
	}
	out := new(UserSSHKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserSSHKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpec) DeepCopyInto(out *UserSpec) {
	*out = *in
	if in.Projects != nil {
		in, out := &in.Projects, &out.Projects
		*out = make([]ProjectGroup, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpec.
func (in *UserSpec) DeepCopy() *UserSpec {
	if in == nil {
		return nil
	}
	out := new(UserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereCloudSpec) DeepCopyInto(out *VSphereCloudSpec) {
	*out = *in
	out.InfraManagementUser = in.InfraManagementUser
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereCloudSpec.
func (in *VSphereCloudSpec) DeepCopy() *VSphereCloudSpec {
	if in == nil {
		return nil
	}
	out := new(VSphereCloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereCredentials) DeepCopyInto(out *VSphereCredentials) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereCredentials.
func (in *VSphereCredentials) DeepCopy() *VSphereCredentials {
	if in == nil {
		return nil
	}
	out := new(VSphereCredentials)
	in.DeepCopyInto(out)
	return out
}
