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

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	//_ = &aws.JSONValue{}
	_ = ackv1alpha1.AWSAccountID("")
)

type AccessPointAlreadyExists struct {
	AccessPointID *string `json:"accessPointID,omitempty"`
}

type AccessPointDescription struct {
	AccessPointARN *string        `json:"accessPointARN,omitempty"`
	AccessPointID  *string        `json:"accessPointID,omitempty"`
	FileSystemID   *string        `json:"fileSystemID,omitempty"`
	LifeCycleState *string        `json:"lifeCycleState,omitempty"`
	Name           *string        `json:"name,omitempty"`
	OwnerID        *string        `json:"ownerID,omitempty"`
	PosixUser      *PosixUser     `json:"posixUser,omitempty"`
	RootDirectory  *RootDirectory `json:"rootDirectory,omitempty"`
	Tags           []*Tag         `json:"tags,omitempty"`
}

type BackupPolicy struct {
	Status *string `json:"status,omitempty"`
}

type CreateAccessPointOutput struct {
	AccessPointARN *string        `json:"accessPointARN,omitempty"`
	AccessPointID  *string        `json:"accessPointID,omitempty"`
	FileSystemID   *string        `json:"fileSystemID,omitempty"`
	LifeCycleState *string        `json:"lifeCycleState,omitempty"`
	Name           *string        `json:"name,omitempty"`
	OwnerID        *string        `json:"ownerID,omitempty"`
	PosixUser      *PosixUser     `json:"posixUser,omitempty"`
	RootDirectory  *RootDirectory `json:"rootDirectory,omitempty"`
	Tags           []*Tag         `json:"tags,omitempty"`
}

type CreateAccessPointRequest struct {
	FileSystemID  *string        `json:"fileSystemID,omitempty"`
	PosixUser     *PosixUser     `json:"posixUser,omitempty"`
	RootDirectory *RootDirectory `json:"rootDirectory,omitempty"`
	Tags          []*Tag         `json:"tags,omitempty"`
}

type CreateFileSystemRequest struct {
	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty"`
	Backup               *bool   `json:"backup,omitempty"`
	Encrypted            *bool   `json:"encrypted,omitempty"`
	KMSKeyID             *string `json:"kmsKeyID,omitempty"`
	PerformanceMode      *string `json:"performanceMode,omitempty"`
	Tags                 []*Tag  `json:"tags,omitempty"`
	ThroughputMode       *string `json:"throughputMode,omitempty"`
}

type CreateMountTargetOutput struct {
	AvailabilityZoneID   *string `json:"availabilityZoneID,omitempty"`
	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty"`
	FileSystemID         *string `json:"fileSystemID,omitempty"`
	IPAddress            *string `json:"ipAddress,omitempty"`
	LifeCycleState       *string `json:"lifeCycleState,omitempty"`
	MountTargetID        *string `json:"mountTargetID,omitempty"`
	NetworkInterfaceID   *string `json:"networkInterfaceID,omitempty"`
	OwnerID              *string `json:"ownerID,omitempty"`
	SubnetID             *string `json:"subnetID,omitempty"`
	VPCID                *string `json:"vpcID,omitempty"`
}

type CreateMountTargetRequest struct {
	FileSystemID   *string   `json:"fileSystemID,omitempty"`
	IPAddress      *string   `json:"ipAddress,omitempty"`
	SecurityGroups []*string `json:"securityGroups,omitempty"`
	SubnetID       *string   `json:"subnetID,omitempty"`
}

type CreateReplicationConfigurationRequest struct {
	SourceFileSystemID *string `json:"sourceFileSystemID,omitempty"`
}

type CreateTagsRequest struct {
	FileSystemID *string `json:"fileSystemID,omitempty"`
	Tags         []*Tag  `json:"tags,omitempty"`
}

type CreationInfo struct {
	OwnerGID    *int64  `json:"ownerGID,omitempty"`
	OwnerUID    *int64  `json:"ownerUID,omitempty"`
	Permissions *string `json:"permissions,omitempty"`
}

type DeleteAccessPointRequest struct {
	AccessPointID *string `json:"accessPointID,omitempty"`
}

type DeleteFileSystemPolicyRequest struct {
	FileSystemID *string `json:"fileSystemID,omitempty"`
}

type DeleteFileSystemRequest struct {
	FileSystemID *string `json:"fileSystemID,omitempty"`
}

type DeleteMountTargetRequest struct {
	MountTargetID *string `json:"mountTargetID,omitempty"`
}

type DeleteReplicationConfigurationRequest struct {
	SourceFileSystemID *string `json:"sourceFileSystemID,omitempty"`
}

type DeleteTagsRequest struct {
	FileSystemID *string `json:"fileSystemID,omitempty"`
}

type DescribeAccessPointsRequest struct {
	AccessPointID *string `json:"accessPointID,omitempty"`
	FileSystemID  *string `json:"fileSystemID,omitempty"`
	MaxResults    *int64  `json:"maxResults,omitempty"`
	NextToken     *string `json:"nextToken,omitempty"`
}

type DescribeAccessPointsResponse struct {
	AccessPoints []*CreateAccessPointOutput `json:"accessPoints,omitempty"`
	NextToken    *string                    `json:"nextToken,omitempty"`
}

type DescribeAccountPreferencesRequest struct {
	MaxResults *int64  `json:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty"`
}

type DescribeAccountPreferencesResponse struct {
	NextToken *string `json:"nextToken,omitempty"`
}

type DescribeBackupPolicyRequest struct {
	FileSystemID *string `json:"fileSystemID,omitempty"`
}

type DescribeFileSystemPolicyRequest struct {
	FileSystemID *string `json:"fileSystemID,omitempty"`
}

type DescribeFileSystemsRequest struct {
	CreationToken *string `json:"creationToken,omitempty"`
	FileSystemID  *string `json:"fileSystemID,omitempty"`
	Marker        *string `json:"marker,omitempty"`
	MaxItems      *int64  `json:"maxItems,omitempty"`
}

type DescribeFileSystemsResponse struct {
	FileSystems []*FileSystemDescription `json:"fileSystems,omitempty"`
	Marker      *string                  `json:"marker,omitempty"`
	NextMarker  *string                  `json:"nextMarker,omitempty"`
}

type DescribeLifecycleConfigurationRequest struct {
	FileSystemID *string `json:"fileSystemID,omitempty"`
}

type DescribeMountTargetSecurityGroupsRequest struct {
	MountTargetID *string `json:"mountTargetID,omitempty"`
}

type DescribeMountTargetSecurityGroupsResponse struct {
	SecurityGroups []*string `json:"securityGroups,omitempty"`
}

type DescribeMountTargetsRequest struct {
	AccessPointID *string `json:"accessPointID,omitempty"`
	FileSystemID  *string `json:"fileSystemID,omitempty"`
	Marker        *string `json:"marker,omitempty"`
	MaxItems      *int64  `json:"maxItems,omitempty"`
	MountTargetID *string `json:"mountTargetID,omitempty"`
}

type DescribeMountTargetsResponse struct {
	Marker       *string                    `json:"marker,omitempty"`
	MountTargets []*CreateMountTargetOutput `json:"mountTargets,omitempty"`
	NextMarker   *string                    `json:"nextMarker,omitempty"`
}

type DescribeReplicationConfigurationsRequest struct {
	FileSystemID *string `json:"fileSystemID,omitempty"`
	MaxResults   *int64  `json:"maxResults,omitempty"`
	NextToken    *string `json:"nextToken,omitempty"`
}

type DescribeReplicationConfigurationsResponse struct {
	NextToken *string `json:"nextToken,omitempty"`
}

type DescribeTagsRequest struct {
	FileSystemID *string `json:"fileSystemID,omitempty"`
	Marker       *string `json:"marker,omitempty"`
	MaxItems     *int64  `json:"maxItems,omitempty"`
}

type DescribeTagsResponse struct {
	Marker     *string `json:"marker,omitempty"`
	NextMarker *string `json:"nextMarker,omitempty"`
	Tags       []*Tag  `json:"tags,omitempty"`
}

type Destination struct {
	FileSystemID            *string      `json:"fileSystemID,omitempty"`
	LastReplicatedTimestamp *metav1.Time `json:"lastReplicatedTimestamp,omitempty"`
}

type DestinationToCreate struct {
	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty"`
	FileSystemID         *string `json:"fileSystemID,omitempty"`
	KMSKeyID             *string `json:"kmsKeyID,omitempty"`
}

type FileSystemAlreadyExists struct {
	FileSystemID *string `json:"fileSystemID,omitempty"`
}

type FileSystemDescription struct {
	AvailabilityZoneID   *string                           `json:"availabilityZoneID,omitempty"`
	AvailabilityZoneName *string                           `json:"availabilityZoneName,omitempty"`
	CreationTime         *metav1.Time                      `json:"creationTime,omitempty"`
	Encrypted            *bool                             `json:"encrypted,omitempty"`
	FileSystemARN        *string                           `json:"fileSystemARN,omitempty"`
	FileSystemID         *string                           `json:"fileSystemID,omitempty"`
	FileSystemProtection *UpdateFileSystemProtectionOutput `json:"fileSystemProtection,omitempty"`
	KMSKeyID             *string                           `json:"kmsKeyID,omitempty"`
	LifeCycleState       *string                           `json:"lifeCycleState,omitempty"`
	Name                 *string                           `json:"name,omitempty"`
	NumberOfMountTargets *int64                            `json:"numberOfMountTargets,omitempty"`
	OwnerID              *string                           `json:"ownerID,omitempty"`
	PerformanceMode      *string                           `json:"performanceMode,omitempty"`
	SizeInBytes          *FileSystemSize                   `json:"sizeInBytes,omitempty"`
	Tags                 []*Tag                            `json:"tags,omitempty"`
	ThroughputMode       *string                           `json:"throughputMode,omitempty"`
}

type FileSystemProtectionDescription struct {
	ReplicationOverwriteProtection *string `json:"replicationOverwriteProtection,omitempty"`
}

type FileSystemSize struct {
	Timestamp       *metav1.Time `json:"timestamp,omitempty"`
	Value           *int64       `json:"value,omitempty"`
	ValueInArchive  *int64       `json:"valueInArchive,omitempty"`
	ValueInIA       *int64       `json:"valueInIA,omitempty"`
	ValueInStandard *int64       `json:"valueInStandard,omitempty"`
}

type LifecyclePolicy struct {
	TransitionToArchive             *string `json:"transitionToArchive,omitempty"`
	TransitionToIA                  *string `json:"transitionToIA,omitempty"`
	TransitionToPrimaryStorageClass *string `json:"transitionToPrimaryStorageClass,omitempty"`
}

type ListTagsForResourceRequest struct {
	MaxResults *int64  `json:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty"`
}

type ListTagsForResourceResponse struct {
	NextToken *string `json:"nextToken,omitempty"`
	Tags      []*Tag  `json:"tags,omitempty"`
}

type ModifyMountTargetSecurityGroupsRequest struct {
	MountTargetID  *string   `json:"mountTargetID,omitempty"`
	SecurityGroups []*string `json:"securityGroups,omitempty"`
}

type MountTargetDescription struct {
	AvailabilityZoneID   *string `json:"availabilityZoneID,omitempty"`
	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty"`
	FileSystemID         *string `json:"fileSystemID,omitempty"`
	IPAddress            *string `json:"ipAddress,omitempty"`
	LifeCycleState       *string `json:"lifeCycleState,omitempty"`
	MountTargetID        *string `json:"mountTargetID,omitempty"`
	NetworkInterfaceID   *string `json:"networkInterfaceID,omitempty"`
	OwnerID              *string `json:"ownerID,omitempty"`
	SubnetID             *string `json:"subnetID,omitempty"`
	VPCID                *string `json:"vpcID,omitempty"`
}

type PosixUser struct {
	GID           *int64   `json:"gid,omitempty"`
	SecondaryGIDs []*int64 `json:"secondaryGIDs,omitempty"`
	UID           *int64   `json:"uid,omitempty"`
}

type PutBackupPolicyRequest struct {
	BackupPolicy *BackupPolicy `json:"backupPolicy,omitempty"`
}

type PutFileSystemPolicyRequest struct {
	Policy *string `json:"policy,omitempty"`
}

type PutLifecycleConfigurationRequest struct {
	LifecyclePolicies []*LifecyclePolicy `json:"lifecyclePolicies,omitempty"`
}

type ReplicationConfigurationDescription struct {
	CreationTime                *metav1.Time `json:"creationTime,omitempty"`
	OriginalSourceFileSystemARN *string      `json:"originalSourceFileSystemARN,omitempty"`
	SourceFileSystemARN         *string      `json:"sourceFileSystemARN,omitempty"`
	SourceFileSystemID          *string      `json:"sourceFileSystemID,omitempty"`
}

type RootDirectory struct {
	CreationInfo *CreationInfo `json:"creationInfo,omitempty"`
	Path         *string       `json:"path,omitempty"`
}

type Tag struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type TagResourceRequest struct {
	Tags []*Tag `json:"tags,omitempty"`
}

type UpdateFileSystemProtectionInput struct {
	ReplicationOverwriteProtection *string `json:"replicationOverwriteProtection,omitempty"`
}

type UpdateFileSystemProtectionOutput struct {
	ReplicationOverwriteProtection *string `json:"replicationOverwriteProtection,omitempty"`
}

type UpdateFileSystemProtectionRequest struct {
	ReplicationOverwriteProtection *string `json:"replicationOverwriteProtection,omitempty"`
}

type UpdateFileSystemRequest struct {
	FileSystemID   *string `json:"fileSystemID,omitempty"`
	ThroughputMode *string `json:"throughputMode,omitempty"`
}
