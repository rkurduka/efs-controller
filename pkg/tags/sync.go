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

package tags

import (
	"context"

	"github.com/aws-controllers-k8s/efs-controller/apis/v1alpha1"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"

	//svcsdk "github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/efs"
	svcsdktypes "github.com/aws/aws-sdk-go-v2/service/efs/types"
)

// Ideally, a part of this code needs to be generated, the other part
// needs to be implemented in a different repository (runtime or pkg).
//
// Few things to node:
// - Some AWS APIs support map[string]string for tags, while others support []*Tag.
// - Some AWS APIs have a limit on the number of tags that can be associated with a resource.
// - We can call a few different names and ways to tag and untag resources.
//   - Some allow to add/remove one tag at a time, while others allow to add/remove multiple tags at once.
//   - Some have a seperate API to list tags, while others return tags as part of the describe response.
//   - Even when the API model states that a Describe response will contain tags, the actual response may
//     not contain any tags. And users are expected to call a seperate ListTags API to get the tags.
//
// - Noting a few diffrent API names:
//   - TagResource, UntagResource, ListTagsForResource
//   - CreateTags, DeleteTags, ListTags
//   - AddTags, RemoveTags, ListTags

// Below are some abstractions that can be used to abstract the implementation details
// of tagging and untagging resources.

type metricsRecorder interface {
	RecordAPICall(opType string, opID string, err error)
}

// tagclient interface commented out for AWS-SDk-GO-V2

// type tagsClient interface {
// 	TagResourceWithContext(context.Context, *svcsdk.TagResourceInput, ...request.Option) (*svcsdk.TagResourceOutput, error)
// 	ListTagsForResourceWithContext(context.Context, *svcsdk.ListTagsForResourceInput, ...request.Option) (*svcsdk.ListTagsForResourceOutput, error)
// 	UntagResourceWithContext(context.Context, *svcsdk.UntagResourceInput, ...request.Option) (*svcsdk.UntagResourceOutput, error)
// }

// syncTags examines the Tags in the supplied Resource and calls the
// TagResource and UntagResource APIs to ensure that the set of
// associated Tags stays in sync with the Resource.Spec.Tags
func SyncTags(
	ctx context.Context,
	clientV2 *efs.Client,
	mr metricsRecorder,
	resourceID string,
	aTags []*v1alpha1.Tag,
	bTags []*v1alpha1.Tag,
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.syncTags")
	defer func() { exit(err) }()

	desiredTags := map[string]*string{}
	for _, t := range aTags {
		desiredTags[*t.Key] = t.Value
	}
	existingTags := map[string]*string{}
	for _, t := range bTags {
		existingTags[*t.Key] = t.Value
	}

	toAdd := map[string]*string{}

	// This is require for AWS-SDK-V2 - converting toDelete to []string from []*string
	// because TagKeys is []string in UntagResourceInput
	toDelete := []string{}

	for k, v := range desiredTags {
		if ev, found := existingTags[k]; !found || *ev != *v {
			toAdd[k] = v
		}
	}

	for k, _ := range existingTags {
		if _, found := desiredTags[k]; !found {
			deleteKey := k
			toDelete = append(toDelete, deleteKey)
		}
	}

	if len(toAdd) > 0 {
		for k, v := range toAdd {
			rlog.Debug("adding tag to resource", "key", k, "value", *v)
		}
		if err = addTags(
			ctx,
			*clientV2,
			mr,
			resourceID,
			toAdd,
		); err != nil {
			return err
		}
	}
	if len(toDelete) > 0 {
		for _, k := range toDelete {
			rlog.Debug("removing tag from resource", "key", k)
		}
		if err = removeTags(
			ctx,
			*clientV2,
			mr,
			resourceID,
			toDelete,
		); err != nil {
			return err
		}
	}

	return nil
}

// addTags adds the supplied Tags to the supplied resource
func addTags(
	ctx context.Context,
	clientv2 efs.Client,
	mr metricsRecorder,
	resourceID string,
	tags map[string]*string,
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.addTag")
	defer func() { exit(err) }()

	sdkTags := []svcsdktypes.Tag{}
	for k, v := range tags {
		sdkTags = append(sdkTags, svcsdktypes.Tag{
			Key:   &k,
			Value: v,
		})
	}

	input := &svcsdk.TagResourceInput{
		ResourceId: &resourceID,
		Tags:       sdkTags,
	}

	_, err = clientv2.TagResource(ctx, input)
	mr.RecordAPICall("UPDATE", "TagResource", err)
	return err
}

// removeTags removes the supplied Tags from the supplied resource
func removeTags(
	ctx context.Context,
	clientv2 efs.Client,
	mr metricsRecorder,
	resourceID string,
	tagKeys []string, // the set of tag keys to delete
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.removeTag")
	defer func() { exit(err) }()

	input := &svcsdk.UntagResourceInput{
		ResourceId: &resourceID,
		TagKeys:    tagKeys,
	}
	_, err = clientv2.UntagResource(ctx, input)
	mr.RecordAPICall("UPDATE", "UntagResource", err)
	return err
}
