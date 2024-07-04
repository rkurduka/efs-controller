// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FileSystemSpec defines the desired state of FileSystem.
type FileSystemSpec struct {
	AvailabilityZoneName *string                                  `json:"availabilityZoneName,omitempty"`
	Backup               *bool                                    `json:"backup,omitempty"`
	BackupPolicy         *BackupPolicy                            `json:"backupPolicy,omitempty"`
	Encrypted            *bool                                    `json:"encrypted,omitempty"`
	FileSystemProtection *UpdateFileSystemProtectionInput         `json:"fileSystemProtection,omitempty"`
	KMSKeyID             *string                                  `json:"kmsKeyID,omitempty"`
	KMSKeyRef            *ackv1alpha1.AWSResourceReferenceWrapper `json:"kmsKeyRef,omitempty"`
	LifecyclePolicies    []*LifecyclePolicy                       `json:"lifecyclePolicies,omitempty"`
	PerformanceMode      *string                                  `json:"performanceMode,omitempty"`
	Policy               *string                                  `json:"policy,omitempty"`
	Tags                 []*Tag                                   `json:"tags,omitempty"`
	ThroughputMode       *string                                  `json:"throughputMode,omitempty"`
}

// FileSystemStatus defines the observed state of FileSystem
type FileSystemStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// +kubebuilder:validation:Optional
	AvailabilityZoneID *string `json:"availabilityZoneID,omitempty"`
	// +kubebuilder:validation:Optional
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
	// +kubebuilder:validation:Optional
	FileSystemID *string `json:"fileSystemID,omitempty"`
	// +kubebuilder:validation:Optional
	LifeCycleState *string `json:"lifeCycleState,omitempty"`
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty"`
	// +kubebuilder:validation:Optional
	NumberOfMountTargets *int64 `json:"numberOfMountTargets,omitempty"`
	// +kubebuilder:validation:Optional
	OwnerID *string `json:"ownerID,omitempty"`
	// +kubebuilder:validation:Optional
	SizeInBytes *FileSystemSize `json:"sizeInBytes,omitempty"`
}

// FileSystem is the Schema for the FileSystems API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="ID",type=string,priority=0,JSONPath=`.status.fileSystemID`
// +kubebuilder:printcolumn:name="ENCRYPTED",type=boolean,priority=0,JSONPath=`.spec.encrypted`
// +kubebuilder:printcolumn:name="PERFORMANCEMODE",type=string,priority=1,JSONPath=`.spec.performanceMode`
// +kubebuilder:printcolumn:name="THROUGHPUTMODE",type=string,priority=1,JSONPath=`.spec.throughputMode`
// +kubebuilder:printcolumn:name="PROVISIONEDTHROUGHPUT",type=string,priority=1,JSONPath=`.status.provisionedThroughputInMiBps`
// +kubebuilder:printcolumn:name="SIZE",type=integer,priority=0,JSONPath=`.status.sizeInBytes.value`
// +kubebuilder:printcolumn:name="MOUNTTARGETS",type=integer,priority=0,JSONPath=`.status.numberOfMountTargets`
// +kubebuilder:printcolumn:name="STATE",type=string,priority=0,JSONPath=`.status.lifeCycleState`
// +kubebuilder:printcolumn:name="Synced",type="string",priority=0,JSONPath=".status.conditions[?(@.type==\"ACK.ResourceSynced\")].status"
// +kubebuilder:printcolumn:name="Age",type="date",priority=0,JSONPath=".metadata.creationTimestamp"
type FileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FileSystemSpec   `json:"spec,omitempty"`
	Status            FileSystemStatus `json:"status,omitempty"`
}

// FileSystemList contains a list of FileSystem
// +kubebuilder:object:root=true
type FileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FileSystem `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FileSystem{}, &FileSystemList{})
}
