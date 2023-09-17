/*
Copyright 2022 The Kubernetes Authors.

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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.25

// ClusterCIDR represents a single configuration for per-Node Pod CIDR
// allocations when the MultiCIDRRangeAllocator is enabled (see the config for
// kube-controller-manager).  A cluster may have any number of ClusterCIDR
// resources, all of which will be considered when allocating a CIDR for a
// Node.  A ClusterCIDR is eligible to be used for a given Node when the node
// selector matches the node in question and has free CIDRs to allocate.  In
// case of multiple matching ClusterCIDR resources, the allocator will attempt
// to break ties using internal heuristics, but any ClusterCIDR whose node
// selector matches the Node may be used.
type ClusterCIDR struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// spec is the desired state of the ClusterCIDR.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Spec ClusterCIDRSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// ClusterCIDRSpec defines the desired state of ClusterCIDR.
type ClusterCIDRSpec struct {
	// nodeSelector defines which nodes the config is applicable to.
	// An empty or nil nodeSelector selects all nodes.
	// This field is immutable.
	// +optional
	NodeSelector *v1.NodeSelector `json:"nodeSelector,omitempty" protobuf:"bytes,1,opt,name=nodeSelector"`

	// perNodeHostBits defines the number of host bits to be configured per node.
	// A subnet mask determines how much of the address is used for network bits
	// and host bits. For example an IPv4 address of 192.168.0.0/24, splits the
	// address into 24 bits for the network portion and 8 bits for the host portion.
	// To allocate 256 IPs, set this field to 8 (a /24 mask for IPv4 or a /120 for IPv6).
	// Minimum value is 4 (16 IPs).
	// This field is immutable.
	// +required
	PerNodeHostBits int32 `json:"perNodeHostBits" protobuf:"varint,2,opt,name=perNodeHostBits"`

	// ipv4 defines an IPv4 IP block in CIDR notation(e.g. "10.0.0.0/8").
	// At least one of ipv4 and ipv6 must be specified.
	// This field is immutable.
	// +optional
	IPv4 string `json:"ipv4" protobuf:"bytes,3,opt,name=ipv4"`

	// ipv6 defines an IPv6 IP block in CIDR notation(e.g. "2001:db8::/64").
	// At least one of ipv4 and ipv6 must be specified.
	// This field is immutable.
	// +optional
	IPv6 string `json:"ipv6" protobuf:"bytes,4,opt,name=ipv6"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.25

// ClusterCIDRList contains a list of ClusterCIDR.
type ClusterCIDRList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// items is the list of ClusterCIDRs.
	Items []ClusterCIDR `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.27

// IPAddress represents a single IP of a single IP Family. The object is designed to be used by APIs
// that operate on IP addresses. The object is used by the Service core API for allocation of IP addresses.
// An IP address can be represented in different formats, to guarantee the uniqueness of the IP,
// the name of the object is the IP address in canonical format, four decimal digits separated
// by dots suppressing leading zeros for IPv4 and the representation defined by RFC 5952 for IPv6.
// Valid: 192.168.1.5 or 2001:db8::1 or 2001:db8:aaaa:bbbb:cccc:dddd:eeee:1
// Invalid: 10.01.2.3 or 2001:db8:0:0:0::1
type IPAddress struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// spec is the desired state of the IPAddress.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Spec IPAddressSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// IPAddressSpec describe the attributes in an IP Address.
type IPAddressSpec struct {
	// ParentRef references the resource that an IPAddress is attached to.
	// An IPAddress must reference a parent object.
	// +required
	ParentRef *ParentReference `json:"parentRef,omitempty" protobuf:"bytes,1,opt,name=parentRef"`
}

// ParentReference describes a reference to a parent object.
type ParentReference struct {
	// Group is the group of the object being referenced.
	// +optional
	Group string `json:"group,omitempty" protobuf:"bytes,1,opt,name=group"`
	// Resource is the resource of the object being referenced.
	// +required
	Resource string `json:"resource,omitempty" protobuf:"bytes,2,opt,name=resource"`
	// Namespace is the namespace of the object being referenced.
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,3,opt,name=namespace"`
	// Name is the name of the object being referenced.
	// +required
	Name string `json:"name,omitempty" protobuf:"bytes,4,opt,name=name"`
	// UID is the uid of the object being referenced.
	// +optional
	UID types.UID `json:"uid,omitempty" protobuf:"bytes,5,opt,name=uid,casttype=k8s.io/apimachinery/pkg/types.UID"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.27

// IPAddressList contains a list of IPAddress.
type IPAddressList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// items is the list of IPAddresses.
	Items []IPAddress `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.27

// PodNetwork represents a logical network in the K8s Cluster.
// This logical network depends on the host networking setup on cluster nodes.
type PodNetwork struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the behavior of a PodNetwork.
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Spec PodNetworkSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Most recently observed status of the PodNetwork.
	// Populated by the system.
	// Read-only.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Status PodNetworkStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// PodNetworkSpec contains the specifications for podNetwork object
type PodNetworkSpec struct {

	// IPAM4 defines what is handling v4 IPAM for Pods attaching to
	// this PodNetwork. When specified, an IPv4 address is required
	// for the attachment.
	// When not specified, no IPv4 is expected to be present and
	// reported on the attachment to this PodNetwork.
	//
	// +optional
	IPAM4 *IPAMType `json:"ipam4,omitempty" protobuf:"bytes,1,opt,name=ipam4,casttype=IPAMType"`

	// IPAM6 defines what is handling v6 IPAM for Pods attaching to
	// this PodNetwork. When specified, an IPv6 address is required
	// for the attachment.
	// When not specified, no IPv6 is expected to be present and
	// reported on the attachment to this PodNetwork.
	//
	// +optional
	IPAM6 *IPAMType `json:"ipam6,omitempty" protobuf:"bytes,2,opt,name=ipam6,casttype=IPAMType"`

	// ParametersRef points to the vendor or implementation specific params for the
	// podNetwork.
	// +optional
	ParametersRefs []ParametersRef `json:"parametersRefs,omitempty" protobuf:"bytes,3,opt,name=parametersRefs"`

	// Provider specifies the provider implementing this PodNetwork.
	// +optional
	Provider string `json:"provider,omitempty" protobuf:"bytes,4,opt,name=provider"`
}

// IPAMType defines source of Pods IPAM handling.
// +enum
type IPAMType string

const (
	// External uses external mechanisms to define IPAM configuration.
	External IPAMType = "external"
	// Kubernetes uses a built-in mechanism to configure IPAM.
	// Based on KCM IPAM controller and ClusterCIDR.
	Kubernetes IPAMType = "kubernetes"
	// None option is used when no IP will be present and
	// reported on the attachment to this PodNetwork.
	NoneType IPAMType = "none"
)

// ParametersRef defines a custom resource containing additional parameters for the
// PodNetwork.
type ParametersRef struct {
	// Group is the group of the object being referenced.
	Group string `json:"group" protobuf:"bytes,1,opt,name=group"`
	// Kind is the resource of the object being referenced.
	Kind string `json:"kind" protobuf:"bytes,2,opt,name=kind"`
	// Namespace is the namespace of the object being referenced.
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,3,opt,name=namespace"`
	// Name is the name of the object being referenced.
	// +required
	Name string `json:"name" protobuf:"bytes,4,opt,name=name"`
}

// PodNetworkConditionType is the type for status conditions on
// a PodNetwork. This type should be used with the
// PodNetworkStatus.Conditions field.
type PodNetworkConditionType string

const (
	// PodNetworkConditionStatusReady represents that the PodNetwork object is
	// correct (validated) and all other conditions are set to “true”. This
	// condition will switch back to “false” if any of the other conditions are
	// “false”. This condition does not indicate readiness of specific PodNetwork
	// on a per Node-basis.
	PodNetworkConditionStatusReady PodNetworkConditionType = "Ready"

	// PodNetworkConditionStatusParamsReady represents that object specified in
	// the “parametersRef” field is ready for use. The owner of the specified
	// object is responsible for handling this condition. The “Ready” condition is
	// dependent on the value of this condition when the “parametersRef” field is
	// not empty. The available “reasons” for this condition are implementation
	// specific.
	PodNetworkConditionStatusParamsReady PodNetworkConditionType = "ParamsReady"
)

// PodNetworkConditionReason defines the set of reasons that explain why a
// particular PodNetwork condition type has been raised.
type PodNetworkConditionReason string

const (
	// PodNetworkConditionReasonParamsNotReady represents a reason where the
	// ParamsReady condition is not present or has “false” value. This can only
	// happen when the “parametersRef” field has a value.
	PodNetworkConditionReasonParamsNotReady PodNetworkConditionReason = "ParamsNotReady"

	// PodNetworkConditionReasonDeleteInProgress represents a reason where the
	// PodNetwork object's deletionTimestamp has value, which indicates it is
	// being removed, but is still in use by other Pods. No new attachment should
	// be added to this PodNetwork.
	PodNetworkConditionReasonDeleteInProgress PodNetworkConditionReason = "DeleteInProgress"
)

// PodNetworkStatus contains the status information related to the PodNetwork.
type PodNetworkStatus struct {
	// Conditions describe the current state of the PodNetwork.
	//
	// Known condition types are:
	//
	// * "Ready"
	// * "ParamsReady"
	// * "InUse"
	//
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=5
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,req,name=conditions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.27

// PodNetworkList contains a list of PodNetwork.
type PodNetworkList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// items is the list of PodNetworks.
	Items []PodNetwork `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.27

// PodNetworkAttachment provides optional pod-level configuration of PodNetwork.
type PodNetworkAttachment struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the behavior of a PodNetworkAttachment.
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Spec PodNetworkAttachmentSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Most recently observed status of the PodNetworkAttachment.
	// Populated by the system.
	// Read-only.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Status PodNetworkAttachmentStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// PodNetworkAttachmentSpec is the specification for the PodNetworkAttachment resource.
type PodNetworkAttachmentSpec struct {
	// PodNetworkName refers to a PodNetwork object that this PodNetworkAttachment is
	// connected to.
	// +required
	PodNetworkName string `json:"podNetworkName" protobuf:"bytes,1,req,name=podNetworkName"`

	// ParametersRefs points to the vendor or implementation specific parameters
	// object for the PodNetworkAttachment.
	// +optional
	ParametersRefs []ParametersRef `json:"parametersRefs,omitempty" protobuf:"bytes,2,opt,name=parametersRefs"`
}

// PodNetworkAttachmentStatus is the status for the PodNetworkAttachment resource.
type PodNetworkAttachmentStatus struct {
	// Conditions describe the current conditions of the PodNetworkAttachment.
	//
	//
	// Known condition types are:
	//
	// * "Ready"
	// * "ParamsReady"
	// * "InUse"
	//
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=5
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,req,name=conditions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.27

// PodNetworkAttachmentList contains a list of PodNetworkAttachment.
type PodNetworkAttachmentList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// items is the list of PodNetworkAttachments.
	Items []PodNetworkAttachment `json:"items" protobuf:"bytes,2,rep,name=items"`
}
