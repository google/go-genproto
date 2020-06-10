// Copyright 2019 Google LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.2
// source: google/cloud/tasks/v2beta2/queue.proto

package tasks

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// State of the queue.
type Queue_State int32

const (
	// Unspecified state.
	Queue_STATE_UNSPECIFIED Queue_State = 0
	// The queue is running. Tasks can be dispatched.
	//
	// If the queue was created using Cloud Tasks and the queue has
	// had no activity (method calls or task dispatches) for 30 days,
	// the queue may take a few minutes to re-activate. Some method
	// calls may return [NOT_FOUND][google.rpc.Code.NOT_FOUND] and
	// tasks may not be dispatched for a few minutes until the queue
	// has been re-activated.
	Queue_RUNNING Queue_State = 1
	// Tasks are paused by the user. If the queue is paused then Cloud
	// Tasks will stop delivering tasks from it, but more tasks can
	// still be added to it by the user. When a pull queue is paused,
	// all [LeaseTasks][google.cloud.tasks.v2beta2.CloudTasks.LeaseTasks] calls will return a
	// [FAILED_PRECONDITION][google.rpc.Code.FAILED_PRECONDITION].
	Queue_PAUSED Queue_State = 2
	// The queue is disabled.
	//
	// A queue becomes `DISABLED` when
	// [queue.yaml](https://cloud.google.com/appengine/docs/python/config/queueref)
	// or
	// [queue.xml](https://cloud.google.com/appengine/docs/standard/java/config/queueref)
	// is uploaded which does not contain the queue. You cannot directly disable
	// a queue.
	//
	// When a queue is disabled, tasks can still be added to a queue
	// but the tasks are not dispatched and
	// [LeaseTasks][google.cloud.tasks.v2beta2.CloudTasks.LeaseTasks] calls return a
	// `FAILED_PRECONDITION` error.
	//
	// To permanently delete this queue and all of its tasks, call
	// [DeleteQueue][google.cloud.tasks.v2beta2.CloudTasks.DeleteQueue].
	Queue_DISABLED Queue_State = 3
)

// Enum value maps for Queue_State.
var (
	Queue_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "RUNNING",
		2: "PAUSED",
		3: "DISABLED",
	}
	Queue_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"RUNNING":           1,
		"PAUSED":            2,
		"DISABLED":          3,
	}
)

func (x Queue_State) Enum() *Queue_State {
	p := new(Queue_State)
	*p = x
	return p
}

func (x Queue_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Queue_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_tasks_v2beta2_queue_proto_enumTypes[0].Descriptor()
}

func (Queue_State) Type() protoreflect.EnumType {
	return &file_google_cloud_tasks_v2beta2_queue_proto_enumTypes[0]
}

func (x Queue_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Queue_State.Descriptor instead.
func (Queue_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_tasks_v2beta2_queue_proto_rawDescGZIP(), []int{0, 0}
}

// A queue is a container of related tasks. Queues are configured to manage
// how those tasks are dispatched. Configurable properties include rate limits,
// retry options, target types, and others.
type Queue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Caller-specified and required in [CreateQueue][google.cloud.tasks.v2beta2.CloudTasks.CreateQueue],
	// after which it becomes output only.
	//
	// The queue name.
	//
	// The queue name must have the following format:
	// `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID`
	//
	// * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//    hyphens (-), colons (:), or periods (.).
	//    For more information, see
	//    [Identifying
	//    projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)
	// * `LOCATION_ID` is the canonical ID for the queue's location.
	//    The list of available locations can be obtained by calling
	//    [ListLocations][google.cloud.location.Locations.ListLocations].
	//    For more information, see https://cloud.google.com/about/locations/.
	// * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or
	//   hyphens (-). The maximum length is 100 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Caller-specified and required in [CreateQueue][google.cloud.tasks.v2beta2.CloudTasks.CreateQueue][],
	// after which the queue config type becomes output only, though fields within
	// the config are mutable.
	//
	// The queue's target.
	//
	// The target applies to all tasks in the queue.
	//
	// Types that are assignable to TargetType:
	//	*Queue_AppEngineHttpTarget
	//	*Queue_PullTarget
	TargetType isQueue_TargetType `protobuf_oneof:"target_type"`
	// Rate limits for task dispatches.
	//
	// [rate_limits][google.cloud.tasks.v2beta2.Queue.rate_limits] and
	// [retry_config][google.cloud.tasks.v2beta2.Queue.retry_config] are related because they both
	// control task attempts however they control how tasks are
	// attempted in different ways:
	//
	// * [rate_limits][google.cloud.tasks.v2beta2.Queue.rate_limits] controls the total rate of
	//   dispatches from a queue (i.e. all traffic dispatched from the
	//   queue, regardless of whether the dispatch is from a first
	//   attempt or a retry).
	// * [retry_config][google.cloud.tasks.v2beta2.Queue.retry_config] controls what happens to
	//   particular a task after its first attempt fails. That is,
	//   [retry_config][google.cloud.tasks.v2beta2.Queue.retry_config] controls task retries (the
	//   second attempt, third attempt, etc).
	RateLimits *RateLimits `protobuf:"bytes,5,opt,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
	// Settings that determine the retry behavior.
	//
	// * For tasks created using Cloud Tasks: the queue-level retry settings
	//   apply to all tasks in the queue that were created using Cloud Tasks.
	//   Retry settings cannot be set on individual tasks.
	// * For tasks created using the App Engine SDK: the queue-level retry
	//   settings apply to all tasks in the queue which do not have retry settings
	//   explicitly set on the task and were created by the App Engine SDK. See
	//   [App Engine
	//   documentation](https://cloud.google.com/appengine/docs/standard/python/taskqueue/push/retrying-tasks).
	RetryConfig *RetryConfig `protobuf:"bytes,6,opt,name=retry_config,json=retryConfig,proto3" json:"retry_config,omitempty"`
	// Output only. The state of the queue.
	//
	// `state` can only be changed by called
	// [PauseQueue][google.cloud.tasks.v2beta2.CloudTasks.PauseQueue],
	// [ResumeQueue][google.cloud.tasks.v2beta2.CloudTasks.ResumeQueue], or uploading
	// [queue.yaml/xml](https://cloud.google.com/appengine/docs/python/config/queueref).
	// [UpdateQueue][google.cloud.tasks.v2beta2.CloudTasks.UpdateQueue] cannot be used to change `state`.
	State Queue_State `protobuf:"varint,7,opt,name=state,proto3,enum=google.cloud.tasks.v2beta2.Queue_State" json:"state,omitempty"`
	// Output only. The last time this queue was purged.
	//
	// All tasks that were [created][google.cloud.tasks.v2beta2.Task.create_time] before this time
	// were purged.
	//
	// A queue can be purged using [PurgeQueue][google.cloud.tasks.v2beta2.CloudTasks.PurgeQueue], the
	// [App Engine Task Queue SDK, or the Cloud
	// Console](https://cloud.google.com/appengine/docs/standard/python/taskqueue/push/deleting-tasks-and-queues#purging_all_tasks_from_a_queue).
	//
	// Purge time will be truncated to the nearest microsecond. Purge
	// time will be unset if the queue has never been purged.
	PurgeTime *timestamp.Timestamp `protobuf:"bytes,8,opt,name=purge_time,json=purgeTime,proto3" json:"purge_time,omitempty"`
}

func (x *Queue) Reset() {
	*x = Queue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_tasks_v2beta2_queue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Queue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Queue) ProtoMessage() {}

func (x *Queue) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_tasks_v2beta2_queue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Queue.ProtoReflect.Descriptor instead.
func (*Queue) Descriptor() ([]byte, []int) {
	return file_google_cloud_tasks_v2beta2_queue_proto_rawDescGZIP(), []int{0}
}

func (x *Queue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *Queue) GetTargetType() isQueue_TargetType {
	if m != nil {
		return m.TargetType
	}
	return nil
}

func (x *Queue) GetAppEngineHttpTarget() *AppEngineHttpTarget {
	if x, ok := x.GetTargetType().(*Queue_AppEngineHttpTarget); ok {
		return x.AppEngineHttpTarget
	}
	return nil
}

func (x *Queue) GetPullTarget() *PullTarget {
	if x, ok := x.GetTargetType().(*Queue_PullTarget); ok {
		return x.PullTarget
	}
	return nil
}

func (x *Queue) GetRateLimits() *RateLimits {
	if x != nil {
		return x.RateLimits
	}
	return nil
}

func (x *Queue) GetRetryConfig() *RetryConfig {
	if x != nil {
		return x.RetryConfig
	}
	return nil
}

func (x *Queue) GetState() Queue_State {
	if x != nil {
		return x.State
	}
	return Queue_STATE_UNSPECIFIED
}

func (x *Queue) GetPurgeTime() *timestamp.Timestamp {
	if x != nil {
		return x.PurgeTime
	}
	return nil
}

type isQueue_TargetType interface {
	isQueue_TargetType()
}

type Queue_AppEngineHttpTarget struct {
	// App Engine HTTP target.
	//
	// An App Engine queue is a queue that has an [AppEngineHttpTarget][google.cloud.tasks.v2beta2.AppEngineHttpTarget].
	AppEngineHttpTarget *AppEngineHttpTarget `protobuf:"bytes,3,opt,name=app_engine_http_target,json=appEngineHttpTarget,proto3,oneof"`
}

type Queue_PullTarget struct {
	// Pull target.
	//
	// A pull queue is a queue that has a [PullTarget][google.cloud.tasks.v2beta2.PullTarget].
	PullTarget *PullTarget `protobuf:"bytes,4,opt,name=pull_target,json=pullTarget,proto3,oneof"`
}

func (*Queue_AppEngineHttpTarget) isQueue_TargetType() {}

func (*Queue_PullTarget) isQueue_TargetType() {}

// Rate limits.
//
// This message determines the maximum rate that tasks can be dispatched by a
// queue, regardless of whether the dispatch is a first task attempt or a retry.
//
// Note: The debugging command, [RunTask][google.cloud.tasks.v2beta2.CloudTasks.RunTask], will run a task
// even if the queue has reached its [RateLimits][google.cloud.tasks.v2beta2.RateLimits].
type RateLimits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum rate at which tasks are dispatched from this queue.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// * For [App Engine queues][google.cloud.tasks.v2beta2.AppEngineHttpTarget], the maximum allowed value
	//   is 500.
	// * This field is output only   for [pull queues][google.cloud.tasks.v2beta2.PullTarget]. In addition to the
	//   `max_tasks_dispatched_per_second` limit, a maximum of 10 QPS of
	//   [LeaseTasks][google.cloud.tasks.v2beta2.CloudTasks.LeaseTasks] requests are allowed per pull queue.
	//
	//
	// This field has the same meaning as
	// [rate in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#rate).
	MaxTasksDispatchedPerSecond float64 `protobuf:"fixed64,1,opt,name=max_tasks_dispatched_per_second,json=maxTasksDispatchedPerSecond,proto3" json:"max_tasks_dispatched_per_second,omitempty"`
	// Output only. The max burst size.
	//
	// Max burst size limits how fast tasks in queue are processed when
	// many tasks are in the queue and the rate is high. This field
	// allows the queue to have a high rate so processing starts shortly
	// after a task is enqueued, but still limits resource usage when
	// many tasks are enqueued in a short period of time.
	//
	// The [token bucket](https://wikipedia.org/wiki/Token_Bucket)
	// algorithm is used to control the rate of task dispatches. Each
	// queue has a token bucket that holds tokens, up to the maximum
	// specified by `max_burst_size`. Each time a task is dispatched, a
	// token is removed from the bucket. Tasks will be dispatched until
	// the queue's bucket runs out of tokens. The bucket will be
	// continuously refilled with new tokens based on
	// [max_tasks_dispatched_per_second][google.cloud.tasks.v2beta2.RateLimits.max_tasks_dispatched_per_second].
	//
	// Cloud Tasks will pick the value of `max_burst_size` based on the
	// value of
	// [max_tasks_dispatched_per_second][google.cloud.tasks.v2beta2.RateLimits.max_tasks_dispatched_per_second].
	//
	// For App Engine queues that were created or updated using
	// `queue.yaml/xml`, `max_burst_size` is equal to
	// [bucket_size](https://cloud.google.com/appengine/docs/standard/python/config/queueref#bucket_size).
	// Since `max_burst_size` is output only, if
	// [UpdateQueue][google.cloud.tasks.v2beta2.CloudTasks.UpdateQueue] is called on a queue
	// created by `queue.yaml/xml`, `max_burst_size` will be reset based
	// on the value of
	// [max_tasks_dispatched_per_second][google.cloud.tasks.v2beta2.RateLimits.max_tasks_dispatched_per_second],
	// regardless of whether
	// [max_tasks_dispatched_per_second][google.cloud.tasks.v2beta2.RateLimits.max_tasks_dispatched_per_second]
	// is updated.
	//
	MaxBurstSize int32 `protobuf:"varint,2,opt,name=max_burst_size,json=maxBurstSize,proto3" json:"max_burst_size,omitempty"`
	// The maximum number of concurrent tasks that Cloud Tasks allows
	// to be dispatched for this queue. After this threshold has been
	// reached, Cloud Tasks stops dispatching tasks until the number of
	// concurrent requests decreases.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	//
	// The maximum allowed value is 5,000.
	//
	// This field is output only for
	// [pull queues][google.cloud.tasks.v2beta2.PullTarget] and always -1, which indicates no limit. No other
	// queue types can have `max_concurrent_tasks` set to -1.
	//
	//
	// This field has the same meaning as
	// [max_concurrent_requests in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#max_concurrent_requests).
	MaxConcurrentTasks int32 `protobuf:"varint,3,opt,name=max_concurrent_tasks,json=maxConcurrentTasks,proto3" json:"max_concurrent_tasks,omitempty"`
}

func (x *RateLimits) Reset() {
	*x = RateLimits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_tasks_v2beta2_queue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimits) ProtoMessage() {}

func (x *RateLimits) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_tasks_v2beta2_queue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimits.ProtoReflect.Descriptor instead.
func (*RateLimits) Descriptor() ([]byte, []int) {
	return file_google_cloud_tasks_v2beta2_queue_proto_rawDescGZIP(), []int{1}
}

func (x *RateLimits) GetMaxTasksDispatchedPerSecond() float64 {
	if x != nil {
		return x.MaxTasksDispatchedPerSecond
	}
	return 0
}

func (x *RateLimits) GetMaxBurstSize() int32 {
	if x != nil {
		return x.MaxBurstSize
	}
	return 0
}

func (x *RateLimits) GetMaxConcurrentTasks() int32 {
	if x != nil {
		return x.MaxConcurrentTasks
	}
	return 0
}

// Retry config.
//
// These settings determine how a failed task attempt is retried.
type RetryConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of attempts per task.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	//
	//
	// This field has the same meaning as
	// [task_retry_limit in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	//
	// Types that are assignable to NumAttempts:
	//	*RetryConfig_MaxAttempts
	//	*RetryConfig_UnlimitedAttempts
	NumAttempts isRetryConfig_NumAttempts `protobuf_oneof:"num_attempts"`
	// If positive, `max_retry_duration` specifies the time limit for
	// retrying a failed task, measured from when the task was first
	// attempted. Once `max_retry_duration` time has passed *and* the
	// task has been attempted [max_attempts][google.cloud.tasks.v2beta2.RetryConfig.max_attempts]
	// times, no further attempts will be made and the task will be
	// deleted.
	//
	// If zero, then the task age is unlimited.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// This field is output only for [pull queues][google.cloud.tasks.v2beta2.PullTarget].
	//
	//
	// `max_retry_duration` will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [task_age_limit in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxRetryDuration *duration.Duration `protobuf:"bytes,3,opt,name=max_retry_duration,json=maxRetryDuration,proto3" json:"max_retry_duration,omitempty"`
	// A task will be [scheduled][google.cloud.tasks.v2beta2.Task.schedule_time] for retry between
	// [min_backoff][google.cloud.tasks.v2beta2.RetryConfig.min_backoff] and
	// [max_backoff][google.cloud.tasks.v2beta2.RetryConfig.max_backoff] duration after it fails,
	// if the queue's [RetryConfig][google.cloud.tasks.v2beta2.RetryConfig] specifies that the task should be
	// retried.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// This field is output only for [pull queues][google.cloud.tasks.v2beta2.PullTarget].
	//
	//
	// `min_backoff` will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [min_backoff_seconds in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MinBackoff *duration.Duration `protobuf:"bytes,4,opt,name=min_backoff,json=minBackoff,proto3" json:"min_backoff,omitempty"`
	// A task will be [scheduled][google.cloud.tasks.v2beta2.Task.schedule_time] for retry between
	// [min_backoff][google.cloud.tasks.v2beta2.RetryConfig.min_backoff] and
	// [max_backoff][google.cloud.tasks.v2beta2.RetryConfig.max_backoff] duration after it fails,
	// if the queue's [RetryConfig][google.cloud.tasks.v2beta2.RetryConfig] specifies that the task should be
	// retried.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// This field is output only for [pull queues][google.cloud.tasks.v2beta2.PullTarget].
	//
	//
	// `max_backoff` will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [max_backoff_seconds in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxBackoff *duration.Duration `protobuf:"bytes,5,opt,name=max_backoff,json=maxBackoff,proto3" json:"max_backoff,omitempty"`
	// The time between retries will double `max_doublings` times.
	//
	// A task's retry interval starts at
	// [min_backoff][google.cloud.tasks.v2beta2.RetryConfig.min_backoff], then doubles
	// `max_doublings` times, then increases linearly, and finally
	// retries retries at intervals of
	// [max_backoff][google.cloud.tasks.v2beta2.RetryConfig.max_backoff] up to
	// [max_attempts][google.cloud.tasks.v2beta2.RetryConfig.max_attempts] times.
	//
	// For example, if [min_backoff][google.cloud.tasks.v2beta2.RetryConfig.min_backoff] is 10s,
	// [max_backoff][google.cloud.tasks.v2beta2.RetryConfig.max_backoff] is 300s, and
	// `max_doublings` is 3, then the a task will first be retried in
	// 10s. The retry interval will double three times, and then
	// increase linearly by 2^3 * 10s.  Finally, the task will retry at
	// intervals of [max_backoff][google.cloud.tasks.v2beta2.RetryConfig.max_backoff] until the
	// task has been attempted [max_attempts][google.cloud.tasks.v2beta2.RetryConfig.max_attempts]
	// times. Thus, the requests will retry at 10s, 20s, 40s, 80s, 160s,
	// 240s, 300s, 300s, ....
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// This field is output only for [pull queues][google.cloud.tasks.v2beta2.PullTarget].
	//
	//
	// This field has the same meaning as
	// [max_doublings in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxDoublings int32 `protobuf:"varint,6,opt,name=max_doublings,json=maxDoublings,proto3" json:"max_doublings,omitempty"`
}

func (x *RetryConfig) Reset() {
	*x = RetryConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_tasks_v2beta2_queue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryConfig) ProtoMessage() {}

func (x *RetryConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_tasks_v2beta2_queue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryConfig.ProtoReflect.Descriptor instead.
func (*RetryConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_tasks_v2beta2_queue_proto_rawDescGZIP(), []int{2}
}

func (m *RetryConfig) GetNumAttempts() isRetryConfig_NumAttempts {
	if m != nil {
		return m.NumAttempts
	}
	return nil
}

func (x *RetryConfig) GetMaxAttempts() int32 {
	if x, ok := x.GetNumAttempts().(*RetryConfig_MaxAttempts); ok {
		return x.MaxAttempts
	}
	return 0
}

func (x *RetryConfig) GetUnlimitedAttempts() bool {
	if x, ok := x.GetNumAttempts().(*RetryConfig_UnlimitedAttempts); ok {
		return x.UnlimitedAttempts
	}
	return false
}

func (x *RetryConfig) GetMaxRetryDuration() *duration.Duration {
	if x != nil {
		return x.MaxRetryDuration
	}
	return nil
}

func (x *RetryConfig) GetMinBackoff() *duration.Duration {
	if x != nil {
		return x.MinBackoff
	}
	return nil
}

func (x *RetryConfig) GetMaxBackoff() *duration.Duration {
	if x != nil {
		return x.MaxBackoff
	}
	return nil
}

func (x *RetryConfig) GetMaxDoublings() int32 {
	if x != nil {
		return x.MaxDoublings
	}
	return 0
}

type isRetryConfig_NumAttempts interface {
	isRetryConfig_NumAttempts()
}

type RetryConfig_MaxAttempts struct {
	// The maximum number of attempts for a task.
	//
	// Cloud Tasks will attempt the task `max_attempts` times (that
	// is, if the first attempt fails, then there will be
	// `max_attempts - 1` retries).  Must be > 0.
	MaxAttempts int32 `protobuf:"varint,1,opt,name=max_attempts,json=maxAttempts,proto3,oneof"`
}

type RetryConfig_UnlimitedAttempts struct {
	// If true, then the number of attempts is unlimited.
	UnlimitedAttempts bool `protobuf:"varint,2,opt,name=unlimited_attempts,json=unlimitedAttempts,proto3,oneof"`
}

func (*RetryConfig_MaxAttempts) isRetryConfig_NumAttempts() {}

func (*RetryConfig_UnlimitedAttempts) isRetryConfig_NumAttempts() {}

var File_google_cloud_tasks_v2beta2_queue_proto protoreflect.FileDescriptor

var file_google_cloud_tasks_v2beta2_queue_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x05, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x66, 0x0a, 0x16, 0x61, 0x70, 0x70, 0x5f, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x48, 0x74, 0x74, 0x70,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48, 0x00, 0x52, 0x13, 0x61, 0x70, 0x70, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x48, 0x74, 0x74, 0x70, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x49, 0x0a,
	0x0b, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e,
	0x50, 0x75, 0x6c, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x75,
	0x6c, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x47, 0x0a, 0x0b, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x52, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x73, 0x12, 0x4a, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x32, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3d, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x70, 0x75, 0x72, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x70, 0x75,
	0x72, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x3a, 0x5c,
	0xea, 0x41, 0x59, 0x0a, 0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x12, 0x36, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x7d, 0x42, 0x0d, 0x0a, 0x0b,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xaa, 0x01, 0x0a, 0x0a,
	0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x44, 0x0a, 0x1f, 0x6d, 0x61,
	0x78, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x1b, 0x6d, 0x61, 0x78, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x44, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x75, 0x72, 0x73, 0x74, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x42, 0x75, 0x72,
	0x73, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f,
	0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x22, 0xd9, 0x02, 0x0a, 0x0b, 0x52, 0x65, 0x74,
	0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f,
	0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x0b, 0x6d, 0x61, 0x78, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x12, 0x2f, 0x0a,
	0x12, 0x75, 0x6e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x11, 0x75, 0x6e, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x12, 0x47,
	0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x72, 0x79, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x5f, 0x62,
	0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6d, 0x69, 0x6e, 0x42, 0x61, 0x63, 0x6b,
	0x6f, 0x66, 0x66, 0x12, 0x3a, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x6f,
	0x66, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x12,
	0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x69, 0x6e, 0x67, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x5f, 0x61, 0x74, 0x74, 0x65,
	0x6d, 0x70, 0x74, 0x73, 0x42, 0x6f, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x42, 0x0a, 0x51, 0x75, 0x65, 0x75, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x3b,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_tasks_v2beta2_queue_proto_rawDescOnce sync.Once
	file_google_cloud_tasks_v2beta2_queue_proto_rawDescData = file_google_cloud_tasks_v2beta2_queue_proto_rawDesc
)

func file_google_cloud_tasks_v2beta2_queue_proto_rawDescGZIP() []byte {
	file_google_cloud_tasks_v2beta2_queue_proto_rawDescOnce.Do(func() {
		file_google_cloud_tasks_v2beta2_queue_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_tasks_v2beta2_queue_proto_rawDescData)
	})
	return file_google_cloud_tasks_v2beta2_queue_proto_rawDescData
}

var file_google_cloud_tasks_v2beta2_queue_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_tasks_v2beta2_queue_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_tasks_v2beta2_queue_proto_goTypes = []interface{}{
	(Queue_State)(0),            // 0: google.cloud.tasks.v2beta2.Queue.State
	(*Queue)(nil),               // 1: google.cloud.tasks.v2beta2.Queue
	(*RateLimits)(nil),          // 2: google.cloud.tasks.v2beta2.RateLimits
	(*RetryConfig)(nil),         // 3: google.cloud.tasks.v2beta2.RetryConfig
	(*AppEngineHttpTarget)(nil), // 4: google.cloud.tasks.v2beta2.AppEngineHttpTarget
	(*PullTarget)(nil),          // 5: google.cloud.tasks.v2beta2.PullTarget
	(*timestamp.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*duration.Duration)(nil),   // 7: google.protobuf.Duration
}
var file_google_cloud_tasks_v2beta2_queue_proto_depIdxs = []int32{
	4, // 0: google.cloud.tasks.v2beta2.Queue.app_engine_http_target:type_name -> google.cloud.tasks.v2beta2.AppEngineHttpTarget
	5, // 1: google.cloud.tasks.v2beta2.Queue.pull_target:type_name -> google.cloud.tasks.v2beta2.PullTarget
	2, // 2: google.cloud.tasks.v2beta2.Queue.rate_limits:type_name -> google.cloud.tasks.v2beta2.RateLimits
	3, // 3: google.cloud.tasks.v2beta2.Queue.retry_config:type_name -> google.cloud.tasks.v2beta2.RetryConfig
	0, // 4: google.cloud.tasks.v2beta2.Queue.state:type_name -> google.cloud.tasks.v2beta2.Queue.State
	6, // 5: google.cloud.tasks.v2beta2.Queue.purge_time:type_name -> google.protobuf.Timestamp
	7, // 6: google.cloud.tasks.v2beta2.RetryConfig.max_retry_duration:type_name -> google.protobuf.Duration
	7, // 7: google.cloud.tasks.v2beta2.RetryConfig.min_backoff:type_name -> google.protobuf.Duration
	7, // 8: google.cloud.tasks.v2beta2.RetryConfig.max_backoff:type_name -> google.protobuf.Duration
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_cloud_tasks_v2beta2_queue_proto_init() }
func file_google_cloud_tasks_v2beta2_queue_proto_init() {
	if File_google_cloud_tasks_v2beta2_queue_proto != nil {
		return
	}
	file_google_cloud_tasks_v2beta2_target_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_tasks_v2beta2_queue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Queue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_tasks_v2beta2_queue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimits); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_tasks_v2beta2_queue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_cloud_tasks_v2beta2_queue_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Queue_AppEngineHttpTarget)(nil),
		(*Queue_PullTarget)(nil),
	}
	file_google_cloud_tasks_v2beta2_queue_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*RetryConfig_MaxAttempts)(nil),
		(*RetryConfig_UnlimitedAttempts)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_tasks_v2beta2_queue_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_tasks_v2beta2_queue_proto_goTypes,
		DependencyIndexes: file_google_cloud_tasks_v2beta2_queue_proto_depIdxs,
		EnumInfos:         file_google_cloud_tasks_v2beta2_queue_proto_enumTypes,
		MessageInfos:      file_google_cloud_tasks_v2beta2_queue_proto_msgTypes,
	}.Build()
	File_google_cloud_tasks_v2beta2_queue_proto = out.File
	file_google_cloud_tasks_v2beta2_queue_proto_rawDesc = nil
	file_google_cloud_tasks_v2beta2_queue_proto_goTypes = nil
	file_google_cloud_tasks_v2beta2_queue_proto_depIdxs = nil
}
