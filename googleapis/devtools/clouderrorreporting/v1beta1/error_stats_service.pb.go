// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/devtools/clouderrorreporting/v1beta1/error_stats_service.proto
// DO NOT EDIT!

package google_devtools_clouderrorreporting_v1beta1 // import "google.golang.org/genproto/googleapis/devtools/clouderrorreporting/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_protobuf2 "github.com/golang/protobuf/ptypes/duration"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Specifies how the time periods of error group counts are aligned.
type TimedCountAlignment int32

const (
	// No alignment specified.
	TimedCountAlignment_ERROR_COUNT_ALIGNMENT_UNSPECIFIED TimedCountAlignment = 0
	// The time periods shall be consecutive, have width equal to the
	// requested duration, and be aligned at the `alignment_time` provided in
	// the request.
	// The `alignment_time` does not have to be inside the query period but
	// even if it is outside, only time periods are returned which overlap
	// with the query period.
	// A rounded alignment will typically result in a
	// different size of the first or the last time period.
	TimedCountAlignment_ALIGNMENT_EQUAL_ROUNDED TimedCountAlignment = 1
	// The time periods shall be consecutive, have width equal to the
	// requested duration, and be aligned at the end of the requested time
	// period. This can result in a different size of the
	// first time period.
	TimedCountAlignment_ALIGNMENT_EQUAL_AT_END TimedCountAlignment = 2
)

var TimedCountAlignment_name = map[int32]string{
	0: "ERROR_COUNT_ALIGNMENT_UNSPECIFIED",
	1: "ALIGNMENT_EQUAL_ROUNDED",
	2: "ALIGNMENT_EQUAL_AT_END",
}
var TimedCountAlignment_value = map[string]int32{
	"ERROR_COUNT_ALIGNMENT_UNSPECIFIED": 0,
	"ALIGNMENT_EQUAL_ROUNDED":           1,
	"ALIGNMENT_EQUAL_AT_END":            2,
}

func (x TimedCountAlignment) String() string {
	return proto.EnumName(TimedCountAlignment_name, int32(x))
}
func (TimedCountAlignment) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// A sorting order of error groups.
type ErrorGroupOrder int32

const (
	// No group order specified.
	ErrorGroupOrder_GROUP_ORDER_UNSPECIFIED ErrorGroupOrder = 0
	// Total count of errors in the given time window in descending order.
	ErrorGroupOrder_COUNT_DESC ErrorGroupOrder = 1
	// Timestamp when the group was last seen in the given time window
	// in descending order.
	ErrorGroupOrder_LAST_SEEN_DESC ErrorGroupOrder = 2
	// Timestamp when the group was created in descending order.
	ErrorGroupOrder_CREATED_DESC ErrorGroupOrder = 3
	// Number of affected users in the given time window in descending order.
	ErrorGroupOrder_AFFECTED_USERS_DESC ErrorGroupOrder = 4
)

var ErrorGroupOrder_name = map[int32]string{
	0: "GROUP_ORDER_UNSPECIFIED",
	1: "COUNT_DESC",
	2: "LAST_SEEN_DESC",
	3: "CREATED_DESC",
	4: "AFFECTED_USERS_DESC",
}
var ErrorGroupOrder_value = map[string]int32{
	"GROUP_ORDER_UNSPECIFIED": 0,
	"COUNT_DESC":              1,
	"LAST_SEEN_DESC":          2,
	"CREATED_DESC":            3,
	"AFFECTED_USERS_DESC":     4,
}

func (x ErrorGroupOrder) String() string {
	return proto.EnumName(ErrorGroupOrder_name, int32(x))
}
func (ErrorGroupOrder) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

// The supported time ranges.
type QueryTimeRange_Period int32

const (
	// Do not use.
	QueryTimeRange_PERIOD_UNSPECIFIED QueryTimeRange_Period = 0
	// Retrieve data for the last hour.
	// Recommended minimum timed count duration: 1 min.
	QueryTimeRange_PERIOD_1_HOUR QueryTimeRange_Period = 1
	// Retrieve data for the last 6 hours.
	// Recommended minimum timed count duration: 10 min.
	QueryTimeRange_PERIOD_6_HOURS QueryTimeRange_Period = 2
	// Retrieve data for the last day.
	// Recommended minimum timed count duration: 1 hour.
	QueryTimeRange_PERIOD_1_DAY QueryTimeRange_Period = 3
	// Retrieve data for the last week.
	// Recommended minimum timed count duration: 6 hours.
	QueryTimeRange_PERIOD_1_WEEK QueryTimeRange_Period = 4
	// Retrieve data for the last 30 days.
	// Recommended minimum timed count duration: 1 day.
	QueryTimeRange_PERIOD_30_DAYS QueryTimeRange_Period = 5
)

var QueryTimeRange_Period_name = map[int32]string{
	0: "PERIOD_UNSPECIFIED",
	1: "PERIOD_1_HOUR",
	2: "PERIOD_6_HOURS",
	3: "PERIOD_1_DAY",
	4: "PERIOD_1_WEEK",
	5: "PERIOD_30_DAYS",
}
var QueryTimeRange_Period_value = map[string]int32{
	"PERIOD_UNSPECIFIED": 0,
	"PERIOD_1_HOUR":      1,
	"PERIOD_6_HOURS":     2,
	"PERIOD_1_DAY":       3,
	"PERIOD_1_WEEK":      4,
	"PERIOD_30_DAYS":     5,
}

func (x QueryTimeRange_Period) String() string {
	return proto.EnumName(QueryTimeRange_Period_name, int32(x))
}
func (QueryTimeRange_Period) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{6, 0} }

// Specifies a set of `ErrorGroupStats` to return.
type ListGroupStatsRequest struct {
	// [Required] The resource name of the Google Cloud Platform project. Written
	// as <code>projects/</code> plus the
	// <a href="https://support.google.com/cloud/answer/6158840">Google Cloud
	// Platform project ID</a>.
	//
	// Example: <code>projects/my-project-123</code>.
	ProjectName string `protobuf:"bytes,1,opt,name=project_name,json=projectName" json:"project_name,omitempty"`
	// [Optional] List all <code>ErrorGroupStats</code> with these IDs.
	// If not specified, all error group stats with a non-zero error count
	// for the given selection criteria are returned.
	GroupId []string `protobuf:"bytes,2,rep,name=group_id,json=groupId" json:"group_id,omitempty"`
	// [Optional] List only <code>ErrorGroupStats</code> which belong to a service
	// context that matches the filter.
	// Data for all service contexts is returned if this field is not specified.
	ServiceFilter *ServiceContextFilter `protobuf:"bytes,3,opt,name=service_filter,json=serviceFilter" json:"service_filter,omitempty"`
	// [Required] List data for the given time range.
	// The service is tuned for retrieving data up to (approximately) 'now'.
	// Retrieving data for arbitrary time periods in the past can result in
	// higher response times or in returning incomplete results.
	TimeRange *QueryTimeRange `protobuf:"bytes,5,opt,name=time_range,json=timeRange" json:"time_range,omitempty"`
	// [Optional] The preferred duration for a single returned `TimedCount`.
	// If not set, no timed counts are returned.
	TimedCountDuration *google_protobuf2.Duration `protobuf:"bytes,6,opt,name=timed_count_duration,json=timedCountDuration" json:"timed_count_duration,omitempty"`
	// [Optional] The alignment of the timed counts to be returned.
	// Default is `ALIGNMENT_EQUAL_AT_END`.
	Alignment TimedCountAlignment `protobuf:"varint,7,opt,name=alignment,enum=google.devtools.clouderrorreporting.v1beta1.TimedCountAlignment" json:"alignment,omitempty"`
	// [Optional] Time where the timed counts shall be aligned if rounded
	// alignment is chosen. Default is 00:00 UTC.
	AlignmentTime *google_protobuf1.Timestamp `protobuf:"bytes,8,opt,name=alignment_time,json=alignmentTime" json:"alignment_time,omitempty"`
	// [Optional] The sort order in which the results are returned.
	// Default is `COUNT_DESC`.
	Order ErrorGroupOrder `protobuf:"varint,9,opt,name=order,enum=google.devtools.clouderrorreporting.v1beta1.ErrorGroupOrder" json:"order,omitempty"`
	// [Optional] The maximum number of results to return per response.
	// Default is 20.
	PageSize int32 `protobuf:"varint,11,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// [Optional] A `next_page_token` provided by a previous response. To view
	// additional results, pass this token along with the identical query
	// parameters as the first request.
	PageToken string `protobuf:"bytes,12,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListGroupStatsRequest) Reset()                    { *m = ListGroupStatsRequest{} }
func (m *ListGroupStatsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListGroupStatsRequest) ProtoMessage()               {}
func (*ListGroupStatsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ListGroupStatsRequest) GetServiceFilter() *ServiceContextFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *ListGroupStatsRequest) GetTimeRange() *QueryTimeRange {
	if m != nil {
		return m.TimeRange
	}
	return nil
}

func (m *ListGroupStatsRequest) GetTimedCountDuration() *google_protobuf2.Duration {
	if m != nil {
		return m.TimedCountDuration
	}
	return nil
}

func (m *ListGroupStatsRequest) GetAlignmentTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.AlignmentTime
	}
	return nil
}

// Contains a set of requested error group stats.
type ListGroupStatsResponse struct {
	// The error group stats which match the given request.
	ErrorGroupStats []*ErrorGroupStats `protobuf:"bytes,1,rep,name=error_group_stats,json=errorGroupStats" json:"error_group_stats,omitempty"`
	// If non-empty, more results are available.
	// Pass this token, along with the same query parameters as the first
	// request, to view the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListGroupStatsResponse) Reset()                    { *m = ListGroupStatsResponse{} }
func (m *ListGroupStatsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListGroupStatsResponse) ProtoMessage()               {}
func (*ListGroupStatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *ListGroupStatsResponse) GetErrorGroupStats() []*ErrorGroupStats {
	if m != nil {
		return m.ErrorGroupStats
	}
	return nil
}

// Data extracted for a specific group based on certain selection criteria,
// such as a given time period and/or service filter.
type ErrorGroupStats struct {
	// Group data that is independent of the selection criteria.
	Group *ErrorGroup `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	// Approximate total number of events in the given group that match
	// the selection criteria.
	Count int64 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	// Approximate number of affected users in the given group that
	// match the selection criteria.
	// Users are distinguished by data in the `ErrorContext` of the
	// individual error events, such as their login name or their remote
	// IP address in case of HTTP requests.
	// The number of affected users can be zero even if the number of
	// errors is non-zero if no data was provided from which the
	// affected user could be deduced.
	// Users are counted based on data in the request
	// context that was provided in the error report. If more users are
	// implicitly affected, such as due to a crash of the whole service,
	// this is not reflected here.
	AffectedUsersCount int64 `protobuf:"varint,3,opt,name=affected_users_count,json=affectedUsersCount" json:"affected_users_count,omitempty"`
	// Approximate number of occurrences over time.
	// Timed counts returned by ListGroups are guaranteed to be:
	//
	// - Inside the requested time interval
	// - Non-overlapping, and
	// - Ordered by ascending time.
	TimedCounts []*TimedCount `protobuf:"bytes,4,rep,name=timed_counts,json=timedCounts" json:"timed_counts,omitempty"`
	// Approximate first occurrence that was seen for this group and
	// which matches the given selection criteria.
	FirstSeenTime *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=first_seen_time,json=firstSeenTime" json:"first_seen_time,omitempty"`
	// Approximate last occurrence that was seen for this group
	// and which matches the given selection criteria.
	LastSeenTime *google_protobuf1.Timestamp `protobuf:"bytes,6,opt,name=last_seen_time,json=lastSeenTime" json:"last_seen_time,omitempty"`
	// Service contexts with a non-zero error count for the given selection
	// criteria. This list can be truncated if multiple services are affected.
	// Refer to `num_affected_services` for the total count.
	AffectedServices []*ServiceContext `protobuf:"bytes,7,rep,name=affected_services,json=affectedServices" json:"affected_services,omitempty"`
	// The total number of services with a non-zero error count for the given
	// selection criteria.
	NumAffectedServices int32 `protobuf:"varint,8,opt,name=num_affected_services,json=numAffectedServices" json:"num_affected_services,omitempty"`
	// An arbitrary event that is chosen as representative for the whole group.
	// The representative event is intended to be used as a quick preview for
	// the whole group. Events in the group are usually sufficiently similar
	// to each other such that showing an arbitrary representative provides
	// insight into the characteristics of the group as a whole.
	Representative *ErrorEvent `protobuf:"bytes,9,opt,name=representative" json:"representative,omitempty"`
}

func (m *ErrorGroupStats) Reset()                    { *m = ErrorGroupStats{} }
func (m *ErrorGroupStats) String() string            { return proto.CompactTextString(m) }
func (*ErrorGroupStats) ProtoMessage()               {}
func (*ErrorGroupStats) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ErrorGroupStats) GetGroup() *ErrorGroup {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ErrorGroupStats) GetTimedCounts() []*TimedCount {
	if m != nil {
		return m.TimedCounts
	}
	return nil
}

func (m *ErrorGroupStats) GetFirstSeenTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.FirstSeenTime
	}
	return nil
}

func (m *ErrorGroupStats) GetLastSeenTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LastSeenTime
	}
	return nil
}

func (m *ErrorGroupStats) GetAffectedServices() []*ServiceContext {
	if m != nil {
		return m.AffectedServices
	}
	return nil
}

func (m *ErrorGroupStats) GetRepresentative() *ErrorEvent {
	if m != nil {
		return m.Representative
	}
	return nil
}

// The number of errors in a given time period.
// All numbers are approximate since the error events are sampled
// before counting them.
type TimedCount struct {
	// Approximate number of occurrences in the given time period.
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	// Start of the time period to which `count` refers (included).
	StartTime *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// End of the time period to which `count` refers (excluded).
	EndTime *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
}

func (m *TimedCount) Reset()                    { *m = TimedCount{} }
func (m *TimedCount) String() string            { return proto.CompactTextString(m) }
func (*TimedCount) ProtoMessage()               {}
func (*TimedCount) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *TimedCount) GetStartTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TimedCount) GetEndTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// Specifies a set of error events to return.
type ListEventsRequest struct {
	// [Required] The resource name of the Google Cloud Platform project. Written
	// as `projects/` plus the
	// [Google Cloud Platform project ID](https://support.google.com/cloud/answer/6158840).
	// Example: `projects/my-project-123`.
	ProjectName string `protobuf:"bytes,1,opt,name=project_name,json=projectName" json:"project_name,omitempty"`
	// [Required] The group for which events shall be returned.
	GroupId string `protobuf:"bytes,2,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	// [Optional] List only ErrorGroups which belong to a service context that
	// matches the filter.
	// Data for all service contexts is returned if this field is not specified.
	ServiceFilter *ServiceContextFilter `protobuf:"bytes,3,opt,name=service_filter,json=serviceFilter" json:"service_filter,omitempty"`
	// [Optional] List only data for the given time range.
	// The service is tuned for retrieving data up to (approximately) 'now'.
	// Retrieving data for arbitrary time periods in the past can result in
	// higher response times or in returning incomplete results.
	// Data for the last hour until now is returned if not specified.
	TimeRange *QueryTimeRange `protobuf:"bytes,4,opt,name=time_range,json=timeRange" json:"time_range,omitempty"`
	// [Optional] The maximum number of results to return per response.
	PageSize int32 `protobuf:"varint,6,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// [Optional] A `next_page_token` provided by a previous response.
	PageToken string `protobuf:"bytes,7,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListEventsRequest) Reset()                    { *m = ListEventsRequest{} }
func (m *ListEventsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListEventsRequest) ProtoMessage()               {}
func (*ListEventsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *ListEventsRequest) GetServiceFilter() *ServiceContextFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *ListEventsRequest) GetTimeRange() *QueryTimeRange {
	if m != nil {
		return m.TimeRange
	}
	return nil
}

// Contains a set of requested error events.
type ListEventsResponse struct {
	// The error events which match the given request.
	ErrorEvents []*ErrorEvent `protobuf:"bytes,1,rep,name=error_events,json=errorEvents" json:"error_events,omitempty"`
	// If non-empty, more results are available.
	// Pass this token, along with the same query parameters as the first
	// request, to view the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListEventsResponse) Reset()                    { *m = ListEventsResponse{} }
func (m *ListEventsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListEventsResponse) ProtoMessage()               {}
func (*ListEventsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *ListEventsResponse) GetErrorEvents() []*ErrorEvent {
	if m != nil {
		return m.ErrorEvents
	}
	return nil
}

// Requests might be rejected or the resulting timed count durations might be
// adjusted for lower durations.
type QueryTimeRange struct {
	// Restricts the query to the specified time range.
	Period QueryTimeRange_Period `protobuf:"varint,1,opt,name=period,enum=google.devtools.clouderrorreporting.v1beta1.QueryTimeRange_Period" json:"period,omitempty"`
}

func (m *QueryTimeRange) Reset()                    { *m = QueryTimeRange{} }
func (m *QueryTimeRange) String() string            { return proto.CompactTextString(m) }
func (*QueryTimeRange) ProtoMessage()               {}
func (*QueryTimeRange) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

// Specifies criteria for filtering a subset of service contexts.
// The fields in the filter correspond to the fields in `ServiceContext`.
// Only exact, case-sensitive matches are supported.
// If a field is unset or empty, it matches arbitrary values.
type ServiceContextFilter struct {
	// [Optional] The exact value to match against
	// [`ServiceContext.service`](/error-reporting/reference/rest/v1beta1/ServiceContext#FIELDS.service).
	Service string `protobuf:"bytes,2,opt,name=service" json:"service,omitempty"`
	// [Optional] The exact value to match against
	// [`ServiceContext.version`](/error-reporting/reference/rest/v1beta1/ServiceContext#FIELDS.version).
	Version string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
}

func (m *ServiceContextFilter) Reset()                    { *m = ServiceContextFilter{} }
func (m *ServiceContextFilter) String() string            { return proto.CompactTextString(m) }
func (*ServiceContextFilter) ProtoMessage()               {}
func (*ServiceContextFilter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

// Deletes all events in the project.
type DeleteEventsRequest struct {
	// [Required] The resource name of the Google Cloud Platform project. Written
	// as `projects/` plus the
	// [Google Cloud Platform project ID](https://support.google.com/cloud/answer/6158840).
	// Example: `projects/my-project-123`.
	ProjectName string `protobuf:"bytes,1,opt,name=project_name,json=projectName" json:"project_name,omitempty"`
}

func (m *DeleteEventsRequest) Reset()                    { *m = DeleteEventsRequest{} }
func (m *DeleteEventsRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteEventsRequest) ProtoMessage()               {}
func (*DeleteEventsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

// Response message for deleting error events.
type DeleteEventsResponse struct {
}

func (m *DeleteEventsResponse) Reset()                    { *m = DeleteEventsResponse{} }
func (m *DeleteEventsResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteEventsResponse) ProtoMessage()               {}
func (*DeleteEventsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func init() {
	proto.RegisterType((*ListGroupStatsRequest)(nil), "google.devtools.clouderrorreporting.v1beta1.ListGroupStatsRequest")
	proto.RegisterType((*ListGroupStatsResponse)(nil), "google.devtools.clouderrorreporting.v1beta1.ListGroupStatsResponse")
	proto.RegisterType((*ErrorGroupStats)(nil), "google.devtools.clouderrorreporting.v1beta1.ErrorGroupStats")
	proto.RegisterType((*TimedCount)(nil), "google.devtools.clouderrorreporting.v1beta1.TimedCount")
	proto.RegisterType((*ListEventsRequest)(nil), "google.devtools.clouderrorreporting.v1beta1.ListEventsRequest")
	proto.RegisterType((*ListEventsResponse)(nil), "google.devtools.clouderrorreporting.v1beta1.ListEventsResponse")
	proto.RegisterType((*QueryTimeRange)(nil), "google.devtools.clouderrorreporting.v1beta1.QueryTimeRange")
	proto.RegisterType((*ServiceContextFilter)(nil), "google.devtools.clouderrorreporting.v1beta1.ServiceContextFilter")
	proto.RegisterType((*DeleteEventsRequest)(nil), "google.devtools.clouderrorreporting.v1beta1.DeleteEventsRequest")
	proto.RegisterType((*DeleteEventsResponse)(nil), "google.devtools.clouderrorreporting.v1beta1.DeleteEventsResponse")
	proto.RegisterEnum("google.devtools.clouderrorreporting.v1beta1.TimedCountAlignment", TimedCountAlignment_name, TimedCountAlignment_value)
	proto.RegisterEnum("google.devtools.clouderrorreporting.v1beta1.ErrorGroupOrder", ErrorGroupOrder_name, ErrorGroupOrder_value)
	proto.RegisterEnum("google.devtools.clouderrorreporting.v1beta1.QueryTimeRange_Period", QueryTimeRange_Period_name, QueryTimeRange_Period_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ErrorStatsService service

type ErrorStatsServiceClient interface {
	// Lists the specified groups.
	ListGroupStats(ctx context.Context, in *ListGroupStatsRequest, opts ...grpc.CallOption) (*ListGroupStatsResponse, error)
	// Lists the specified events.
	ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error)
	// Deletes all error events of a given project.
	DeleteEvents(ctx context.Context, in *DeleteEventsRequest, opts ...grpc.CallOption) (*DeleteEventsResponse, error)
}

type errorStatsServiceClient struct {
	cc *grpc.ClientConn
}

func NewErrorStatsServiceClient(cc *grpc.ClientConn) ErrorStatsServiceClient {
	return &errorStatsServiceClient{cc}
}

func (c *errorStatsServiceClient) ListGroupStats(ctx context.Context, in *ListGroupStatsRequest, opts ...grpc.CallOption) (*ListGroupStatsResponse, error) {
	out := new(ListGroupStatsResponse)
	err := grpc.Invoke(ctx, "/google.devtools.clouderrorreporting.v1beta1.ErrorStatsService/ListGroupStats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *errorStatsServiceClient) ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error) {
	out := new(ListEventsResponse)
	err := grpc.Invoke(ctx, "/google.devtools.clouderrorreporting.v1beta1.ErrorStatsService/ListEvents", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *errorStatsServiceClient) DeleteEvents(ctx context.Context, in *DeleteEventsRequest, opts ...grpc.CallOption) (*DeleteEventsResponse, error) {
	out := new(DeleteEventsResponse)
	err := grpc.Invoke(ctx, "/google.devtools.clouderrorreporting.v1beta1.ErrorStatsService/DeleteEvents", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ErrorStatsService service

type ErrorStatsServiceServer interface {
	// Lists the specified groups.
	ListGroupStats(context.Context, *ListGroupStatsRequest) (*ListGroupStatsResponse, error)
	// Lists the specified events.
	ListEvents(context.Context, *ListEventsRequest) (*ListEventsResponse, error)
	// Deletes all error events of a given project.
	DeleteEvents(context.Context, *DeleteEventsRequest) (*DeleteEventsResponse, error)
}

func RegisterErrorStatsServiceServer(s *grpc.Server, srv ErrorStatsServiceServer) {
	s.RegisterService(&_ErrorStatsService_serviceDesc, srv)
}

func _ErrorStatsService_ListGroupStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorStatsServiceServer).ListGroupStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouderrorreporting.v1beta1.ErrorStatsService/ListGroupStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorStatsServiceServer).ListGroupStats(ctx, req.(*ListGroupStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ErrorStatsService_ListEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorStatsServiceServer).ListEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouderrorreporting.v1beta1.ErrorStatsService/ListEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorStatsServiceServer).ListEvents(ctx, req.(*ListEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ErrorStatsService_DeleteEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorStatsServiceServer).DeleteEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouderrorreporting.v1beta1.ErrorStatsService/DeleteEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorStatsServiceServer).DeleteEvents(ctx, req.(*DeleteEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ErrorStatsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.clouderrorreporting.v1beta1.ErrorStatsService",
	HandlerType: (*ErrorStatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListGroupStats",
			Handler:    _ErrorStatsService_ListGroupStats_Handler,
		},
		{
			MethodName: "ListEvents",
			Handler:    _ErrorStatsService_ListEvents_Handler,
		},
		{
			MethodName: "DeleteEvents",
			Handler:    _ErrorStatsService_DeleteEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor2,
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/devtools/clouderrorreporting/v1beta1/error_stats_service.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 1270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x57, 0xdf, 0x6e, 0x1b, 0xc5,
	0x17, 0xfe, 0xad, 0x1d, 0x27, 0xf1, 0xb1, 0xe3, 0x38, 0x93, 0x34, 0xdd, 0xba, 0xfa, 0x49, 0xa9,
	0x25, 0x50, 0x48, 0x85, 0xb7, 0x49, 0x55, 0x4a, 0x55, 0xfe, 0xd4, 0xb1, 0x37, 0x25, 0x34, 0xb5,
	0xdd, 0xb1, 0xad, 0x8a, 0x5c, 0xb0, 0xda, 0xd8, 0x27, 0x9b, 0x05, 0x7b, 0x77, 0xd9, 0x19, 0x47,
	0x6d, 0x51, 0x25, 0xc4, 0x1d, 0xd7, 0x70, 0x81, 0xc4, 0x1b, 0xf0, 0x14, 0x5c, 0x71, 0x8d, 0xd4,
	0x57, 0xe0, 0x1e, 0x89, 0x27, 0x40, 0x33, 0xb3, 0xfe, 0x9b, 0x88, 0xc4, 0x4e, 0x2f, 0xb8, 0xdb,
	0x39, 0x67, 0xce, 0x77, 0xbe, 0x39, 0xf3, 0xcd, 0x99, 0x59, 0x68, 0x3b, 0xbe, 0xef, 0x74, 0xb0,
	0xe0, 0xf8, 0x1d, 0xdb, 0x73, 0x0a, 0x7e, 0xe8, 0x18, 0x0e, 0x7a, 0x41, 0xe8, 0x73, 0xdf, 0x50,
	0x2e, 0x3b, 0x70, 0x99, 0xd1, 0xc6, 0x53, 0xee, 0xfb, 0x1d, 0x66, 0xb4, 0x3a, 0x7e, 0xaf, 0x8d,
	0x61, 0xe8, 0x87, 0x21, 0x06, 0x7e, 0xc8, 0x5d, 0xcf, 0x31, 0x4e, 0xb7, 0x8f, 0x90, 0xdb, 0xdb,
	0x86, 0x34, 0x5b, 0x8c, 0xdb, 0x9c, 0x59, 0x0c, 0xc3, 0x53, 0xb7, 0x85, 0x05, 0x89, 0x44, 0x6e,
	0x47, 0x59, 0xfa, 0x30, 0x85, 0x73, 0x60, 0x0a, 0x11, 0x4c, 0x6e, 0xff, 0x72, 0x94, 0xec, 0xc0,
	0x35, 0xa2, 0x2c, 0x2d, 0xdf, 0x3b, 0x76, 0x1d, 0xc3, 0xf6, 0x3c, 0x9f, 0xdb, 0xdc, 0xf5, 0x3d,
	0xa6, 0xf2, 0xe6, 0x9e, 0xbf, 0xc5, 0xd5, 0xb5, 0xfc, 0x6e, 0xd7, 0xf7, 0x22, 0xe0, 0x07, 0x8e,
	0xcb, 0x4f, 0x7a, 0x47, 0x85, 0x96, 0xdf, 0x35, 0x14, 0xb8, 0x21, 0x1d, 0x47, 0xbd, 0x63, 0x23,
	0xe0, 0x2f, 0x03, 0x64, 0x46, 0xbb, 0x17, 0x4a, 0x2e, 0x83, 0x8f, 0x28, 0xf4, 0xe1, 0xc5, 0xa1,
	0xdc, 0xed, 0x22, 0xe3, 0x76, 0x37, 0x18, 0x7e, 0xa9, 0xe0, 0xfc, 0x2f, 0x09, 0xb8, 0x76, 0xe0,
	0x32, 0xfe, 0x38, 0xf4, 0x7b, 0x41, 0x5d, 0x54, 0x9a, 0xe2, 0x37, 0x3d, 0x64, 0x9c, 0xdc, 0x82,
	0x74, 0x10, 0xfa, 0x5f, 0x61, 0x8b, 0x5b, 0x9e, 0xdd, 0x45, 0x5d, 0xdb, 0xd0, 0x36, 0x93, 0x34,
	0x15, 0xd9, 0x2a, 0x76, 0x17, 0xc9, 0x0d, 0x58, 0x74, 0x44, 0x9c, 0xe5, 0xb6, 0xf5, 0xd8, 0x46,
	0x7c, 0x33, 0x49, 0x17, 0xe4, 0x78, 0xbf, 0x4d, 0x4e, 0x20, 0x13, 0xd5, 0xd2, 0x3a, 0x76, 0x3b,
	0x1c, 0x43, 0x3d, 0xbe, 0xa1, 0x6d, 0xa6, 0x76, 0x8a, 0x85, 0x29, 0x76, 0xae, 0x50, 0x57, 0x10,
	0x25, 0xdf, 0xe3, 0xf8, 0x82, 0xef, 0x49, 0x20, 0xba, 0x14, 0x01, 0xab, 0x21, 0x39, 0x04, 0x10,
	0x8b, 0xb2, 0x42, 0xdb, 0x73, 0x50, 0x4f, 0xc8, 0x2c, 0x0f, 0xa7, 0xca, 0xf2, 0xac, 0x87, 0xe1,
	0xcb, 0x86, 0xdb, 0x45, 0x2a, 0x20, 0x68, 0x92, 0xf7, 0x3f, 0xc9, 0x13, 0x58, 0x13, 0x83, 0xb6,
	0xd5, 0xf2, 0x7b, 0x1e, 0xb7, 0xfa, 0x85, 0xd7, 0xe7, 0x65, 0x96, 0x1b, 0xfd, 0x2c, 0xfd, 0x72,
	0x17, 0xca, 0xd1, 0x04, 0x4a, 0x64, 0x58, 0x49, 0x44, 0xf5, 0x6d, 0xe4, 0x4b, 0x48, 0xda, 0x1d,
	0xd7, 0xf1, 0xba, 0xe8, 0x71, 0x7d, 0x61, 0x43, 0xdb, 0xcc, 0xec, 0x3c, 0x9a, 0x8a, 0x67, 0x63,
	0x80, 0x59, 0xec, 0xe3, 0xd0, 0x21, 0x24, 0x29, 0x42, 0x66, 0x30, 0xb0, 0x44, 0x7e, 0x7d, 0x51,
	0xd2, 0xcc, 0x9d, 0xa1, 0xd9, 0xe8, 0x8b, 0x80, 0x2e, 0x0d, 0x22, 0x84, 0x8d, 0x50, 0x48, 0xf8,
	0x61, 0x1b, 0x43, 0x3d, 0x29, 0xe9, 0x7d, 0x34, 0x15, 0x3d, 0x53, 0x98, 0xa5, 0x8e, 0xaa, 0x02,
	0x83, 0x2a, 0x28, 0x72, 0x13, 0x92, 0x81, 0xed, 0xa0, 0xc5, 0xdc, 0x57, 0xa8, 0xa7, 0x36, 0xb4,
	0xcd, 0x04, 0x5d, 0x14, 0x86, 0xba, 0xfb, 0x0a, 0xc9, 0xff, 0x01, 0xa4, 0x93, 0xfb, 0x5f, 0xa3,
	0xa7, 0xa7, 0xa5, 0xc4, 0xe4, 0xf4, 0x86, 0x30, 0xe4, 0x7f, 0xd5, 0x60, 0x7d, 0x52, 0x9d, 0x2c,
	0xf0, 0x3d, 0x86, 0xe4, 0x04, 0x56, 0x54, 0x7b, 0x50, 0x0a, 0x94, 0x4d, 0x42, 0xd7, 0x36, 0xe2,
	0x9b, 0xa9, 0x99, 0x69, 0xab, 0x04, 0xcb, 0x38, 0x6e, 0x20, 0xef, 0xc2, 0xb2, 0x87, 0x2f, 0xb8,
	0x35, 0x42, 0x34, 0x26, 0x89, 0x2e, 0x09, 0x73, 0x6d, 0x40, 0xf6, 0x87, 0x04, 0x2c, 0x4f, 0x80,
	0x91, 0xa7, 0x90, 0x90, 0xfc, 0xe4, 0xe9, 0x49, 0xed, 0xdc, 0x9f, 0x91, 0x19, 0x55, 0x28, 0x64,
	0x0d, 0x12, 0x52, 0x89, 0x92, 0x40, 0x9c, 0xaa, 0x01, 0xb9, 0x03, 0x6b, 0xf6, 0xf1, 0x31, 0xb6,
	0x38, 0xb6, 0xad, 0x1e, 0xc3, 0x90, 0x29, 0xb9, 0xca, 0x13, 0x17, 0xa7, 0xa4, 0xef, 0x6b, 0x0a,
	0x97, 0x94, 0x0f, 0x39, 0x84, 0xf4, 0x88, 0xae, 0x99, 0x3e, 0x27, 0xeb, 0x76, 0x7f, 0x46, 0x35,
	0xd2, 0xd4, 0x50, 0xed, 0x8c, 0xec, 0xc2, 0xf2, 0xb1, 0x1b, 0x32, 0x6e, 0x31, 0x44, 0x4f, 0xe9,
	0x30, 0x71, 0xb1, 0x0e, 0x65, 0x48, 0x1d, 0xd1, 0x93, 0x3a, 0x7c, 0x04, 0x99, 0x8e, 0x3d, 0x06,
	0x31, 0x7f, 0x21, 0x44, 0x5a, 0x44, 0x0c, 0x10, 0x4e, 0x60, 0x65, 0x50, 0x93, 0xa8, 0x5f, 0x30,
	0x7d, 0x41, 0x2e, 0xf3, 0xe1, 0x15, 0x5a, 0x10, 0xcd, 0xf6, 0x51, 0x23, 0x3b, 0x23, 0x3b, 0x70,
	0xcd, 0xeb, 0x75, 0xad, 0xb3, 0xd9, 0x16, 0xa5, 0xd6, 0x57, 0xbd, 0x5e, 0xb7, 0x38, 0x19, 0x63,
	0x41, 0x26, 0xc4, 0x20, 0x44, 0x86, 0x9e, 0xb8, 0x5f, 0x4e, 0x51, 0x1e, 0xb8, 0x99, 0xf4, 0x61,
	0x9e, 0x8a, 0x36, 0x30, 0x01, 0x97, 0xff, 0x49, 0x03, 0x18, 0x6e, 0xd0, 0x50, 0x37, 0xda, 0xa8,
	0x6e, 0x1e, 0x00, 0x30, 0x6e, 0x87, 0x51, 0xb3, 0x88, 0x5d, 0x58, 0xe1, 0xa4, 0x9c, 0x2d, 0xcb,
	0x7b, 0x0f, 0x16, 0xd1, 0x6b, 0xab, 0xc0, 0xf8, 0x85, 0x81, 0x0b, 0xe8, 0xb5, 0xc5, 0x28, 0xff,
	0x26, 0x06, 0x2b, 0xe2, 0x3c, 0x4b, 0xd2, 0xb3, 0xdf, 0x34, 0xda, 0x7f, 0xe1, 0xa6, 0x99, 0x7b,
	0xab, 0x37, 0xcd, 0x58, 0x97, 0x9c, 0xff, 0xd7, 0x2e, 0xb9, 0x30, 0xd9, 0x25, 0x7f, 0xd6, 0x80,
	0x8c, 0x56, 0x35, 0xea, 0x90, 0x87, 0x90, 0x56, 0x1d, 0x12, 0xa5, 0x3d, 0x6a, 0x8e, 0x33, 0x4b,
	0x2c, 0x85, 0x83, 0xef, 0xcb, 0xf7, 0xc4, 0xbf, 0x34, 0xc8, 0x8c, 0x2f, 0x9a, 0x1c, 0xc2, 0x7c,
	0x80, 0xa1, 0xeb, 0xb7, 0xe5, 0x3e, 0x67, 0x76, 0x76, 0xaf, 0x50, 0xc1, 0x42, 0x4d, 0x22, 0xd1,
	0x08, 0x31, 0xff, 0x9d, 0x06, 0xf3, 0xca, 0x44, 0xd6, 0x81, 0xd4, 0x4c, 0xba, 0x5f, 0x2d, 0x5b,
	0xcd, 0x4a, 0xbd, 0x66, 0x96, 0xf6, 0xf7, 0xf6, 0xcd, 0x72, 0xf6, 0x7f, 0x64, 0x05, 0x96, 0x22,
	0xfb, 0xb6, 0xf5, 0x59, 0xb5, 0x49, 0xb3, 0x1a, 0x21, 0x90, 0x89, 0x4c, 0x1f, 0x48, 0x53, 0x3d,
	0x1b, 0x23, 0x59, 0x48, 0x0f, 0xa6, 0x95, 0x8b, 0x5f, 0x64, 0xe3, 0x63, 0x81, 0xcf, 0x4d, 0xf3,
	0x49, 0x76, 0x6e, 0x24, 0xf0, 0xee, 0x1d, 0x31, 0xab, 0x9e, 0x4d, 0xe4, 0x3f, 0x87, 0xb5, 0xf3,
	0xb4, 0x44, 0x74, 0x58, 0x88, 0xd4, 0xd4, 0x17, 0x70, 0x34, 0x14, 0x9e, 0x53, 0x0c, 0x99, 0x78,
	0x57, 0xc4, 0x95, 0x27, 0x1a, 0xe6, 0x3f, 0x84, 0xd5, 0x32, 0x76, 0x90, 0xe3, 0xb4, 0xe7, 0x25,
	0xbf, 0x0e, 0x6b, 0xe3, 0x91, 0x4a, 0x13, 0x5b, 0x3d, 0x58, 0x3d, 0xe7, 0x15, 0x41, 0xde, 0x81,
	0x5b, 0x26, 0xa5, 0x55, 0x6a, 0x95, 0xaa, 0xcd, 0x4a, 0xc3, 0x2a, 0x1e, 0xec, 0x3f, 0xae, 0x3c,
	0x35, 0x2b, 0x8d, 0x89, 0xda, 0xdd, 0x84, 0xeb, 0x43, 0x97, 0xf9, 0xac, 0x59, 0x3c, 0xb0, 0x68,
	0xb5, 0x59, 0x29, 0x9b, 0xe5, 0xac, 0x46, 0x72, 0xb0, 0x3e, 0xe9, 0x2c, 0x36, 0x2c, 0xb3, 0x52,
	0xce, 0xc6, 0xb6, 0x5e, 0x8f, 0xde, 0x8c, 0xd5, 0xe8, 0x59, 0x70, 0xfd, 0x31, 0xad, 0x36, 0x6b,
	0x56, 0x95, 0x96, 0x4d, 0x3a, 0x91, 0x28, 0x03, 0xa0, 0x98, 0x94, 0xcd, 0x7a, 0x49, 0xed, 0xd0,
	0x41, 0xb1, 0xde, 0xb0, 0xea, 0xa6, 0x59, 0x51, 0x36, 0xb9, 0x43, 0x25, 0x6a, 0x16, 0x1b, 0x66,
	0x59, 0x59, 0xe2, 0xe4, 0x3a, 0xac, 0x16, 0xf7, 0xf6, 0xcc, 0x92, 0x30, 0x35, 0xeb, 0x26, 0xad,
	0x2b, 0xc7, 0xdc, 0xce, 0xdf, 0x73, 0xb0, 0x22, 0xf3, 0xcb, 0x4b, 0x39, 0xda, 0x1e, 0xf2, 0x87,
	0x06, 0x99, 0xf1, 0xc7, 0x05, 0x99, 0x4e, 0x8b, 0xe7, 0xbe, 0x9b, 0x73, 0xa5, 0x2b, 0x61, 0xa8,
	0x7d, 0xca, 0xdf, 0xfb, 0xfe, 0xcd, 0x9f, 0x3f, 0xc6, 0x0c, 0xf2, 0xfe, 0xe0, 0x6f, 0xe1, 0xdb,
	0xd1, 0x2d, 0xff, 0x38, 0x1a, 0x30, 0x63, 0xeb, 0xb5, 0xe1, 0x0c, 0xf9, 0xff, 0xa6, 0x01, 0x0c,
	0x3b, 0x01, 0xf9, 0x64, 0x6a, 0x2a, 0x63, 0x42, 0xcb, 0x7d, 0x3a, 0x73, 0x7c, 0xb4, 0x8c, 0x6d,
	0xb9, 0x8c, 0xdb, 0xe4, 0xbd, 0x4b, 0x2c, 0x43, 0x75, 0x29, 0xf2, 0xbb, 0x06, 0xe9, 0x51, 0xe9,
	0x92, 0xe9, 0xde, 0xc8, 0xe7, 0x9c, 0x97, 0x5c, 0xf1, 0x0a, 0x08, 0xe3, 0x0b, 0xd9, 0xba, 0xfc,
	0x42, 0x76, 0x9f, 0x80, 0xf8, 0xc3, 0x9b, 0x26, 0xf5, 0xee, 0xfa, 0x19, 0x91, 0xd6, 0xc4, 0x65,
	0x5a, 0xd3, 0x8e, 0xe6, 0xe5, 0xad, 0x7a, 0xf7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x82,
	0xf6, 0x9d, 0x5e, 0x0f, 0x00, 0x00,
}
