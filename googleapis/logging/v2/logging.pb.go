// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/logging/v2/logging.proto
// DO NOT EDIT!

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_api3 "google.golang.org/genproto/googleapis/api/monitoredres"
import _ "github.com/golang/protobuf/ptypes/duration"
import google_protobuf4 "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/rpc/status"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The parameters to DeleteLog.
type DeleteLogRequest struct {
	// Required. The resource name of the log to delete.  Example:
	// `"projects/my-project/logs/syslog"`.
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName" json:"log_name,omitempty"`
}

func (m *DeleteLogRequest) Reset()                    { *m = DeleteLogRequest{} }
func (m *DeleteLogRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteLogRequest) ProtoMessage()               {}
func (*DeleteLogRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// The parameters to WriteLogEntries.
type WriteLogEntriesRequest struct {
	// Optional. A default log resource name for those log entries in `entries`
	// that do not specify their own `logName`.  Example:
	// `"projects/my-project/logs/syslog"`.  See
	// [LogEntry][google.logging.v2.LogEntry].
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName" json:"log_name,omitempty"`
	// Optional. A default monitored resource for those log entries in `entries`
	// that do not specify their own `resource`.
	Resource *google_api3.MonitoredResource `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
	// Optional. User-defined `key:value` items that are added to
	// the `labels` field of each log entry in `entries`, except when a log
	// entry specifies its own `key:value` item with the same key.
	// Example: `{ "size": "large", "color":"red" }`
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Required. The log entries to write. The log entries must have values for
	// all required fields.
	//
	// To improve throughput and to avoid exceeding the quota limit for calls
	// to `entries.write`, use this field to write multiple log entries at once
	// rather than  // calling this method for each log entry.
	Entries []*LogEntry `protobuf:"bytes,4,rep,name=entries" json:"entries,omitempty"`
	// Optional. Whether valid entries should be written even if some other
	// entries fail due to INVALID_ARGUMENT or PERMISSION_DENIED errors. If any
	// entry is not written, the response status will be the error associated
	// with one of the failed entries and include error details in the form of
	// WriteLogEntriesPartialErrors.
	PartialSuccess bool `protobuf:"varint,5,opt,name=partial_success,json=partialSuccess" json:"partial_success,omitempty"`
}

func (m *WriteLogEntriesRequest) Reset()                    { *m = WriteLogEntriesRequest{} }
func (m *WriteLogEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteLogEntriesRequest) ProtoMessage()               {}
func (*WriteLogEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *WriteLogEntriesRequest) GetResource() *google_api3.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *WriteLogEntriesRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *WriteLogEntriesRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// Result returned from WriteLogEntries.
// empty
type WriteLogEntriesResponse struct {
}

func (m *WriteLogEntriesResponse) Reset()                    { *m = WriteLogEntriesResponse{} }
func (m *WriteLogEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*WriteLogEntriesResponse) ProtoMessage()               {}
func (*WriteLogEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

// The parameters to `ListLogEntries`.
type ListLogEntriesRequest struct {
	// Required. One or more project IDs or project numbers from which to retrieve
	// log entries.  Examples of a project ID: `"my-project-1A"`, `"1234567890"`.
	ProjectIds []string `protobuf:"bytes,1,rep,name=project_ids,json=projectIds" json:"project_ids,omitempty"`
	// Optional. An [advanced logs filter](/logging/docs/view/advanced_filters).
	// The filter is compared against all log entries in the projects specified by
	// `projectIds`.  Only entries that match the filter are retrieved.  An empty
	// filter matches all log entries.
	Filter string `protobuf:"bytes,2,opt,name=filter" json:"filter,omitempty"`
	// Optional. How the results should be sorted.  Presently, the only permitted
	// values are `"timestamp asc"` (default) and `"timestamp desc"`. The first
	// option returns entries in order of increasing values of
	// `LogEntry.timestamp` (oldest first), and the second option returns entries
	// in order of decreasing timestamps (newest first).  Entries with equal
	// timestamps are returned in order of `LogEntry.insertId`.
	OrderBy string `protobuf:"bytes,3,opt,name=order_by,json=orderBy" json:"order_by,omitempty"`
	// Optional. The maximum number of results to return from this request.
	// You must check for presence of `nextPageToken` to determine if additional
	// results are available, which you can retrieve by passing the
	// `nextPageToken` value as the `pageToken` parameter in the next request.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// Optional. If the `pageToken` parameter is supplied, then the next page of
	// results is retrieved.  The `pageToken` parameter must be set to the value
	// of the `nextPageToken` from the previous response.
	// The values of `projectIds`, `filter`, and `orderBy` must be the same
	// as in the previous request.
	PageToken string `protobuf:"bytes,5,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListLogEntriesRequest) Reset()                    { *m = ListLogEntriesRequest{} }
func (m *ListLogEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLogEntriesRequest) ProtoMessage()               {}
func (*ListLogEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

// Result returned from `ListLogEntries`.
type ListLogEntriesResponse struct {
	// A list of log entries.
	Entries []*LogEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	// If there are more results than were returned, then `nextPageToken` is
	// included in the response.  To get the next set of results, call this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListLogEntriesResponse) Reset()                    { *m = ListLogEntriesResponse{} }
func (m *ListLogEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListLogEntriesResponse) ProtoMessage()               {}
func (*ListLogEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *ListLogEntriesResponse) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// The parameters to ListMonitoredResourceDescriptors
type ListMonitoredResourceDescriptorsRequest struct {
	// Optional. The maximum number of results to return from this request.
	// You must check for presence of `nextPageToken` to determine if additional
	// results are available, which you can retrieve by passing the
	// `nextPageToken` value as the `pageToken` parameter in the next request.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// Optional. If the `pageToken` parameter is supplied, then the next page of
	// results is retrieved.  The `pageToken` parameter must be set to the value
	// of the `nextPageToken` from the previous response.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListMonitoredResourceDescriptorsRequest) Reset() {
	*m = ListMonitoredResourceDescriptorsRequest{}
}
func (m *ListMonitoredResourceDescriptorsRequest) String() string { return proto.CompactTextString(m) }
func (*ListMonitoredResourceDescriptorsRequest) ProtoMessage()    {}
func (*ListMonitoredResourceDescriptorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{5}
}

// Result returned from ListMonitoredResourceDescriptors.
type ListMonitoredResourceDescriptorsResponse struct {
	// A list of resource descriptors.
	ResourceDescriptors []*google_api3.MonitoredResourceDescriptor `protobuf:"bytes,1,rep,name=resource_descriptors,json=resourceDescriptors" json:"resource_descriptors,omitempty"`
	// If there are more results than were returned, then `nextPageToken` is
	// included in the response.  To get the next set of results, call this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListMonitoredResourceDescriptorsResponse) Reset() {
	*m = ListMonitoredResourceDescriptorsResponse{}
}
func (m *ListMonitoredResourceDescriptorsResponse) String() string { return proto.CompactTextString(m) }
func (*ListMonitoredResourceDescriptorsResponse) ProtoMessage()    {}
func (*ListMonitoredResourceDescriptorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{6}
}

func (m *ListMonitoredResourceDescriptorsResponse) GetResourceDescriptors() []*google_api3.MonitoredResourceDescriptor {
	if m != nil {
		return m.ResourceDescriptors
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteLogRequest)(nil), "google.logging.v2.DeleteLogRequest")
	proto.RegisterType((*WriteLogEntriesRequest)(nil), "google.logging.v2.WriteLogEntriesRequest")
	proto.RegisterType((*WriteLogEntriesResponse)(nil), "google.logging.v2.WriteLogEntriesResponse")
	proto.RegisterType((*ListLogEntriesRequest)(nil), "google.logging.v2.ListLogEntriesRequest")
	proto.RegisterType((*ListLogEntriesResponse)(nil), "google.logging.v2.ListLogEntriesResponse")
	proto.RegisterType((*ListMonitoredResourceDescriptorsRequest)(nil), "google.logging.v2.ListMonitoredResourceDescriptorsRequest")
	proto.RegisterType((*ListMonitoredResourceDescriptorsResponse)(nil), "google.logging.v2.ListMonitoredResourceDescriptorsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for LoggingServiceV2 service

type LoggingServiceV2Client interface {
	// Deletes a log and all its log entries.
	// The log will reappear if it receives new entries.
	//
	DeleteLog(ctx context.Context, in *DeleteLogRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error)
	// Writes log entries to Stackdriver Logging.  All log entries are
	// written by this method.
	//
	WriteLogEntries(ctx context.Context, in *WriteLogEntriesRequest, opts ...grpc.CallOption) (*WriteLogEntriesResponse, error)
	// Lists log entries.  Use this method to retrieve log entries from Cloud
	// Logging.  For ways to export log entries, see
	// [Exporting Logs](/logging/docs/export).
	//
	ListLogEntries(ctx context.Context, in *ListLogEntriesRequest, opts ...grpc.CallOption) (*ListLogEntriesResponse, error)
	// Lists the monitored resource descriptors used by Stackdriver Logging.
	ListMonitoredResourceDescriptors(ctx context.Context, in *ListMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListMonitoredResourceDescriptorsResponse, error)
}

type loggingServiceV2Client struct {
	cc *grpc.ClientConn
}

func NewLoggingServiceV2Client(cc *grpc.ClientConn) LoggingServiceV2Client {
	return &loggingServiceV2Client{cc}
}

func (c *loggingServiceV2Client) DeleteLog(ctx context.Context, in *DeleteLogRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error) {
	out := new(google_protobuf4.Empty)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/DeleteLog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) WriteLogEntries(ctx context.Context, in *WriteLogEntriesRequest, opts ...grpc.CallOption) (*WriteLogEntriesResponse, error) {
	out := new(WriteLogEntriesResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/WriteLogEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) ListLogEntries(ctx context.Context, in *ListLogEntriesRequest, opts ...grpc.CallOption) (*ListLogEntriesResponse, error) {
	out := new(ListLogEntriesResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/ListLogEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) ListMonitoredResourceDescriptors(ctx context.Context, in *ListMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListMonitoredResourceDescriptorsResponse, error) {
	out := new(ListMonitoredResourceDescriptorsResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/ListMonitoredResourceDescriptors", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoggingServiceV2 service

type LoggingServiceV2Server interface {
	// Deletes a log and all its log entries.
	// The log will reappear if it receives new entries.
	//
	DeleteLog(context.Context, *DeleteLogRequest) (*google_protobuf4.Empty, error)
	// Writes log entries to Stackdriver Logging.  All log entries are
	// written by this method.
	//
	WriteLogEntries(context.Context, *WriteLogEntriesRequest) (*WriteLogEntriesResponse, error)
	// Lists log entries.  Use this method to retrieve log entries from Cloud
	// Logging.  For ways to export log entries, see
	// [Exporting Logs](/logging/docs/export).
	//
	ListLogEntries(context.Context, *ListLogEntriesRequest) (*ListLogEntriesResponse, error)
	// Lists the monitored resource descriptors used by Stackdriver Logging.
	ListMonitoredResourceDescriptors(context.Context, *ListMonitoredResourceDescriptorsRequest) (*ListMonitoredResourceDescriptorsResponse, error)
}

func RegisterLoggingServiceV2Server(s *grpc.Server, srv LoggingServiceV2Server) {
	s.RegisterService(&_LoggingServiceV2_serviceDesc, srv)
}

func _LoggingServiceV2_DeleteLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).DeleteLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/DeleteLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).DeleteLog(ctx, req.(*DeleteLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_WriteLogEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteLogEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).WriteLogEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/WriteLogEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).WriteLogEntries(ctx, req.(*WriteLogEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_ListLogEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLogEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).ListLogEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/ListLogEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).ListLogEntries(ctx, req.(*ListLogEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_ListMonitoredResourceDescriptors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMonitoredResourceDescriptorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).ListMonitoredResourceDescriptors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/ListMonitoredResourceDescriptors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).ListMonitoredResourceDescriptors(ctx, req.(*ListMonitoredResourceDescriptorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoggingServiceV2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.logging.v2.LoggingServiceV2",
	HandlerType: (*LoggingServiceV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteLog",
			Handler:    _LoggingServiceV2_DeleteLog_Handler,
		},
		{
			MethodName: "WriteLogEntries",
			Handler:    _LoggingServiceV2_WriteLogEntries_Handler,
		},
		{
			MethodName: "ListLogEntries",
			Handler:    _LoggingServiceV2_ListLogEntries_Handler,
		},
		{
			MethodName: "ListMonitoredResourceDescriptors",
			Handler:    _LoggingServiceV2_ListMonitoredResourceDescriptors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor3,
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/logging/v2/logging.proto", fileDescriptor3)
}

var fileDescriptor3 = []byte{
	// 828 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x06, 0x25, 0xff, 0x48, 0xeb, 0xd6, 0x76, 0xb7, 0xb6, 0x2a, 0xd3, 0x30, 0xaa, 0xb2, 0x68,
	0x25, 0x0b, 0x30, 0xd9, 0xca, 0x30, 0x50, 0xcb, 0x70, 0x51, 0x18, 0xf6, 0xc1, 0x80, 0xdc, 0x1a,
	0x74, 0xd1, 0x02, 0x46, 0x00, 0x82, 0xa2, 0xd6, 0xcc, 0xc6, 0x14, 0x97, 0xd9, 0x5d, 0xca, 0x51,
	0x82, 0x5c, 0x72, 0xc9, 0x21, 0xc7, 0x3c, 0x44, 0x6e, 0x79, 0x90, 0x5c, 0x73, 0xc9, 0x03, 0xe4,
	0x01, 0x72, 0xcc, 0x31, 0x4b, 0x72, 0x29, 0xd9, 0x92, 0x62, 0xd1, 0xbe, 0x48, 0x3b, 0xb3, 0xf3,
	0xcd, 0xcf, 0x37, 0xb3, 0x43, 0xf0, 0xa7, 0x4b, 0x88, 0xeb, 0x21, 0xdd, 0x25, 0x9e, 0xed, 0xbb,
	0x3a, 0xa1, 0xae, 0xe1, 0x22, 0x3f, 0xa0, 0x84, 0x13, 0x23, 0xb9, 0xb2, 0x03, 0xcc, 0x0c, 0x8f,
	0xb8, 0x2e, 0xf6, 0x5d, 0xa3, 0xd7, 0x48, 0x8f, 0x7a, 0x6c, 0x03, 0xbf, 0x93, 0xf8, 0x54, 0xdb,
	0x6b, 0xa8, 0xc7, 0xd9, 0x5c, 0x8a, 0x1f, 0x83, 0x21, 0xda, 0xc3, 0x0e, 0x72, 0x88, 0x7f, 0x81,
	0x5d, 0xc3, 0xf6, 0x7d, 0xc2, 0x6d, 0x8e, 0x89, 0xcf, 0x12, 0xef, 0xea, 0x3f, 0xd9, 0x5d, 0x75,
	0x89, 0x8f, 0x39, 0xa1, 0xa8, 0x43, 0x11, 0x1b, 0x0a, 0x96, 0x90, 0x48, 0x48, 0x1d, 0x24, 0x1d,
	0xfe, 0x75, 0x9f, 0x72, 0x2d, 0xe4, 0x73, 0xda, 0x97, 0x1e, 0x76, 0x5d, 0xcc, 0x1f, 0x86, 0x6d,
	0xdd, 0x21, 0x5d, 0x23, 0xf1, 0x62, 0xc4, 0x17, 0xed, 0xf0, 0xc2, 0x08, 0x78, 0x3f, 0x10, 0xd1,
	0x3b, 0x21, 0x8d, 0xab, 0x18, 0x1c, 0x24, 0x74, 0x7b, 0x3a, 0x14, 0x75, 0xc5, 0x21, 0xf9, 0x95,
	0xa0, 0xbd, 0xe9, 0x20, 0x8e, 0xbb, 0x88, 0x71, 0xbb, 0x1b, 0x0c, 0x4f, 0x12, 0xbc, 0x9f, 0xad,
	0x5c, 0x1a, 0x38, 0x86, 0x80, 0xf1, 0x90, 0xc9, 0xbf, 0x04, 0xae, 0x6d, 0x81, 0xe5, 0x43, 0xe4,
	0x21, 0x8e, 0x5a, 0xc4, 0x35, 0xd1, 0xe3, 0x50, 0xf8, 0x86, 0x6b, 0xa0, 0x10, 0x51, 0xe2, 0xdb,
	0x5d, 0x54, 0x56, 0x2a, 0x4a, 0xad, 0x68, 0xce, 0x0b, 0xf9, 0x6f, 0x21, 0x6a, 0x1f, 0x72, 0xa0,
	0xf4, 0x3f, 0xc5, 0xb1, 0xf9, 0x91, 0xa0, 0x0c, 0x23, 0x36, 0x1d, 0x05, 0x77, 0x41, 0x21, 0x6d,
	0x52, 0x39, 0x27, 0xae, 0x16, 0x1a, 0x1b, 0xba, 0x4c, 0x5b, 0x24, 0xa7, 0x9f, 0xa4, 0xad, 0x34,
	0xa5, 0x91, 0x39, 0x30, 0x87, 0x27, 0x60, 0xce, 0xb3, 0xdb, 0xc8, 0x63, 0xe5, 0x7c, 0x25, 0x2f,
	0x80, 0x3b, 0xfa, 0xd8, 0x34, 0xea, 0x93, 0x13, 0xd2, 0x5b, 0x31, 0x2e, 0x52, 0xf6, 0x4d, 0xe9,
	0x04, 0xee, 0x80, 0x79, 0x94, 0x58, 0x95, 0x67, 0x62, 0x7f, 0xeb, 0x13, 0xfc, 0x49, 0x57, 0x7d,
	0x33, 0xb5, 0x85, 0x55, 0xb0, 0x14, 0xd8, 0x94, 0x63, 0xdb, 0xb3, 0x58, 0xe8, 0x38, 0x88, 0xb1,
	0xf2, 0xac, 0xa8, 0xa3, 0x60, 0x2e, 0x4a, 0xf5, 0x59, 0xa2, 0x55, 0x77, 0xc1, 0xc2, 0xb5, 0xb0,
	0x70, 0x19, 0xe4, 0x2f, 0x51, 0x5f, 0xd2, 0x11, 0x1d, 0xe1, 0x0a, 0x98, 0xed, 0xd9, 0x5e, 0x98,
	0xf0, 0x50, 0x34, 0x13, 0xa1, 0x99, 0xfb, 0x43, 0xd1, 0xd6, 0xc0, 0x0f, 0x63, 0x85, 0xb0, 0x40,
	0x3c, 0x14, 0xa4, 0xbd, 0x51, 0xc0, 0x6a, 0x0b, 0x33, 0x3e, 0x4e, 0xfa, 0x8f, 0x60, 0x41, 0xf4,
	0xf1, 0x11, 0x72, 0xb8, 0x85, 0x3b, 0x4c, 0x04, 0xca, 0x0b, 0xa7, 0x40, 0xaa, 0x8e, 0x3b, 0x0c,
	0x96, 0xc0, 0xdc, 0x05, 0xf6, 0x38, 0xa2, 0x32, 0xa0, 0x94, 0xa2, 0x6e, 0x11, 0xda, 0x41, 0xd4,
	0x6a, 0xf7, 0x05, 0xb3, 0x71, 0xb7, 0x62, 0xf9, 0xa0, 0x0f, 0xd7, 0x41, 0x31, 0xb0, 0x5d, 0x64,
	0x31, 0xfc, 0x14, 0x09, 0x96, 0x94, 0xda, 0xac, 0x59, 0x88, 0x14, 0x67, 0x42, 0x86, 0x1b, 0x00,
	0xc4, 0x97, 0x9c, 0x5c, 0x22, 0x3f, 0x26, 0xa1, 0x68, 0xc6, 0xe6, 0xff, 0x46, 0x0a, 0xed, 0x0a,
	0x94, 0x46, 0x13, 0x4d, 0x6a, 0xb8, 0xce, 0xbc, 0x72, 0x07, 0xe6, 0x7f, 0x05, 0x4b, 0x3e, 0x7a,
	0xc2, 0xad, 0x6b, 0x41, 0x93, 0x42, 0xbe, 0x8d, 0xd4, 0xa7, 0x83, 0xc0, 0x08, 0x54, 0xa3, 0xc0,
	0x63, 0xa3, 0x74, 0x88, 0x98, 0x43, 0x71, 0x20, 0x74, 0x03, 0xce, 0x6e, 0xd4, 0xa7, 0xdc, 0x5a,
	0x5f, 0x6e, 0xb4, 0xbe, 0xb7, 0x0a, 0xa8, 0x4d, 0x8f, 0x23, 0x4b, 0x3e, 0x07, 0x2b, 0xe9, 0x1c,
	0x5b, 0x9d, 0xe1, 0xbd, 0xac, 0xbf, 0x7a, 0xeb, 0x13, 0x18, 0xfa, 0x33, 0xbf, 0xa7, 0xe3, 0x31,
	0xb2, 0xf2, 0xd2, 0xf8, 0x34, 0x03, 0x96, 0x5b, 0x09, 0xc1, 0x67, 0xc9, 0x26, 0xfe, 0xaf, 0x01,
	0xaf, 0x40, 0x71, 0xf0, 0xe8, 0xe1, 0xcf, 0x13, 0xfa, 0x30, 0xba, 0x12, 0xd4, 0x52, 0x6a, 0x94,
	0x2e, 0x26, 0xfd, 0x28, 0x5a, 0x60, 0xda, 0xd6, 0x8b, 0xf7, 0x1f, 0x5f, 0xe7, 0xaa, 0xf5, 0x5f,
	0xc4, 0x1e, 0x6d, 0x23, 0x6e, 0xff, 0x6e, 0x3c, 0x4b, 0x97, 0xc0, 0xbe, 0x9c, 0x42, 0x66, 0xd4,
	0xa3, 0x0d, 0x2b, 0xfe, 0x9e, 0xc3, 0x57, 0x0a, 0x58, 0x1a, 0x19, 0x72, 0xb8, 0x99, 0xf9, 0x45,
	0xab, 0xf5, 0x2c, 0xa6, 0xf2, 0xcd, 0xfc, 0x14, 0x67, 0xb6, 0xae, 0x95, 0x06, 0x99, 0xc9, 0x91,
	0x6a, 0x5e, 0x45, 0x88, 0xa6, 0x52, 0x87, 0x2f, 0x15, 0xb0, 0x78, 0x73, 0x5a, 0x61, 0x6d, 0xd2,
	0x50, 0x4e, 0x7a, 0x79, 0xea, 0x66, 0x06, 0x4b, 0x99, 0x4a, 0x25, 0x4e, 0x45, 0xd5, 0x56, 0xc7,
	0x52, 0xf1, 0x04, 0x20, 0xca, 0xe4, 0x9d, 0x02, 0x2a, 0xd3, 0xc6, 0x0a, 0x36, 0xbf, 0x12, 0x31,
	0xc3, 0xcc, 0xab, 0x7b, 0xf7, 0xc2, 0xca, 0xfc, 0x65, 0x93, 0xe1, 0xb0, 0xc9, 0xdd, 0x5b, 0x60,
	0x07, 0x0f, 0xc0, 0xaa, 0xf8, 0x92, 0x8d, 0x07, 0x3c, 0xf8, 0x46, 0x0e, 0xe2, 0x69, 0x34, 0x43,
	0xa7, 0xca, 0xf9, 0x6f, 0x77, 0xfd, 0x52, 0x7f, 0x56, 0x94, 0xf6, 0x5c, 0x7c, 0xbf, 0xfd, 0x25,
	0x00, 0x00, 0xff, 0xff, 0x5b, 0x63, 0x2c, 0x2c, 0xd6, 0x08, 0x00, 0x00,
}
