/*
Copyright 2019 Google LLC

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
	"knative.dev/pkg/apis"

	duckv1alpha1 "github.com/google/knative-gcp/pkg/apis/duck/v1alpha1"
)

// GetCondition returns the condition currently associated with the given type, or nil.
func (s *CloudStorageSourceStatus) GetCondition(t apis.ConditionType) *apis.Condition {
	return storageCondSet.Manage(s).GetCondition(t)
}

// GetTopLevelCondition returns the top level condition.
func (s *CloudStorageSourceStatus) GetTopLevelCondition() *apis.Condition {
	return storageCondSet.Manage(s).GetTopLevelCondition()
}

// IsReady returns true if the resource is ready overall.
func (s *CloudStorageSourceStatus) IsReady() bool {
	return storageCondSet.Manage(s).IsHappy()
}

// InitializeConditions sets relevant unset conditions to Unknown state.
func (s *CloudStorageSourceStatus) InitializeConditions() {
	storageCondSet.Manage(s).InitializeConditions()
}

// MarkPullSubscriptionFailed sets the condition that the status of underlying PullSubscription
// is False and why.
func (s *CloudStorageSourceStatus) MarkPullSubscriptionFailed(reason, messageFormat string, messageA ...interface{}) {
	storageCondSet.Manage(s).MarkFalse(duckv1alpha1.PullSubscriptionReady, reason, messageFormat, messageA...)
}

// MarkPullSubscriptionUnknown sets the condition that the status of underlying PullSubscription
// is Unknown and why.
func (s *CloudStorageSourceStatus) MarkPullSubscriptionUnknown(reason, messageFormat string, messageA ...interface{}) {
	storageCondSet.Manage(s).MarkUnknown(duckv1alpha1.PullSubscriptionReady, reason, messageFormat, messageA...)
}

// MarkPullSubscriptionReady sets the condition that the underlying PullSubscription is ready.
func (s *CloudStorageSourceStatus) MarkPullSubscriptionReady() {
	storageCondSet.Manage(s).MarkTrue(duckv1alpha1.PullSubscriptionReady)
}

// MarkTopicFailed sets the condition that the status of PubSub topic is False why.
func (s *CloudStorageSourceStatus) MarkTopicFailed(reason, messageFormat string, messageA ...interface{}) {
	storageCondSet.Manage(s).MarkFalse(duckv1alpha1.TopicReady, reason, messageFormat, messageA...)
}

// MarkTopicUnknown sets the condition that the status of PubSub topic is Unknown why.
func (s *CloudStorageSourceStatus) MarkTopicUnknown(reason, messageFormat string, messageA ...interface{}) {
	storageCondSet.Manage(s).MarkUnknown(duckv1alpha1.TopicReady, reason, messageFormat, messageA...)
}

// MarkTopicReady sets the condition that the underlying PubSub topic was created successfully.
func (s *CloudStorageSourceStatus) MarkTopicReady() {
	storageCondSet.Manage(s).MarkTrue(duckv1alpha1.TopicReady)
}

// MarkNotificationNotReady sets the condition that the GCS has not been configured
// to send Notifications and why.
func (s *CloudStorageSourceStatus) MarkNotificationNotReady(reason, messageFormat string, messageA ...interface{}) {
	storageCondSet.Manage(s).MarkFalse(NotificationReady, reason, messageFormat, messageA...)
}

func (s *CloudStorageSourceStatus) MarkNotificationReady(notificationID string) {
	s.NotificationID = notificationID
	storageCondSet.Manage(s).MarkTrue(NotificationReady)
}
