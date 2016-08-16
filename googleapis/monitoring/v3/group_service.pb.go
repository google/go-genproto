// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/monitoring/v3/group_service.proto
// DO NOT EDIT!

package google_monitoring_v3 // import "google.golang.org/genproto/googleapis/monitoring/v3"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_api4 "google.golang.org/genproto/googleapis/api/monitoredres"
import google_protobuf4 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The `ListGroup` request.
type ListGroupsRequest struct {
	// The project whose groups are to be listed. The format is
	// `"projects/{project_id_or_number}"`.
	Name string `protobuf:"bytes,7,opt,name=name" json:"name,omitempty"`
	// An optional filter consisting of a single group name.  The filters limit the
	// groups returned based on their parent-child relationship with the specified
	// group. If no filter is specified, all groups are returned.
	//
	// Types that are valid to be assigned to Filter:
	//	*ListGroupsRequest_ChildrenOfGroup
	//	*ListGroupsRequest_AncestorsOfGroup
	//	*ListGroupsRequest_DescendantsOfGroup
	Filter isListGroupsRequest_Filter `protobuf_oneof:"filter"`
	// A positive number that is the maximum number of results to return.
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// If this field is not empty then it must contain the `nextPageToken` value
	// returned by a previous call to this method.  Using this field causes the
	// method to return additional results from the previous method call.
	PageToken string `protobuf:"bytes,6,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListGroupsRequest) Reset()                    { *m = ListGroupsRequest{} }
func (m *ListGroupsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListGroupsRequest) ProtoMessage()               {}
func (*ListGroupsRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type isListGroupsRequest_Filter interface {
	isListGroupsRequest_Filter()
}

type ListGroupsRequest_ChildrenOfGroup struct {
	ChildrenOfGroup string `protobuf:"bytes,2,opt,name=children_of_group,json=childrenOfGroup,oneof"`
}
type ListGroupsRequest_AncestorsOfGroup struct {
	AncestorsOfGroup string `protobuf:"bytes,3,opt,name=ancestors_of_group,json=ancestorsOfGroup,oneof"`
}
type ListGroupsRequest_DescendantsOfGroup struct {
	DescendantsOfGroup string `protobuf:"bytes,4,opt,name=descendants_of_group,json=descendantsOfGroup,oneof"`
}

func (*ListGroupsRequest_ChildrenOfGroup) isListGroupsRequest_Filter()    {}
func (*ListGroupsRequest_AncestorsOfGroup) isListGroupsRequest_Filter()   {}
func (*ListGroupsRequest_DescendantsOfGroup) isListGroupsRequest_Filter() {}

func (m *ListGroupsRequest) GetFilter() isListGroupsRequest_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListGroupsRequest) GetChildrenOfGroup() string {
	if x, ok := m.GetFilter().(*ListGroupsRequest_ChildrenOfGroup); ok {
		return x.ChildrenOfGroup
	}
	return ""
}

func (m *ListGroupsRequest) GetAncestorsOfGroup() string {
	if x, ok := m.GetFilter().(*ListGroupsRequest_AncestorsOfGroup); ok {
		return x.AncestorsOfGroup
	}
	return ""
}

func (m *ListGroupsRequest) GetDescendantsOfGroup() string {
	if x, ok := m.GetFilter().(*ListGroupsRequest_DescendantsOfGroup); ok {
		return x.DescendantsOfGroup
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ListGroupsRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ListGroupsRequest_OneofMarshaler, _ListGroupsRequest_OneofUnmarshaler, _ListGroupsRequest_OneofSizer, []interface{}{
		(*ListGroupsRequest_ChildrenOfGroup)(nil),
		(*ListGroupsRequest_AncestorsOfGroup)(nil),
		(*ListGroupsRequest_DescendantsOfGroup)(nil),
	}
}

func _ListGroupsRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ListGroupsRequest)
	// filter
	switch x := m.Filter.(type) {
	case *ListGroupsRequest_ChildrenOfGroup:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ChildrenOfGroup)
	case *ListGroupsRequest_AncestorsOfGroup:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.AncestorsOfGroup)
	case *ListGroupsRequest_DescendantsOfGroup:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.DescendantsOfGroup)
	case nil:
	default:
		return fmt.Errorf("ListGroupsRequest.Filter has unexpected type %T", x)
	}
	return nil
}

func _ListGroupsRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ListGroupsRequest)
	switch tag {
	case 2: // filter.children_of_group
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Filter = &ListGroupsRequest_ChildrenOfGroup{x}
		return true, err
	case 3: // filter.ancestors_of_group
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Filter = &ListGroupsRequest_AncestorsOfGroup{x}
		return true, err
	case 4: // filter.descendants_of_group
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Filter = &ListGroupsRequest_DescendantsOfGroup{x}
		return true, err
	default:
		return false, nil
	}
}

func _ListGroupsRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ListGroupsRequest)
	// filter
	switch x := m.Filter.(type) {
	case *ListGroupsRequest_ChildrenOfGroup:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.ChildrenOfGroup)))
		n += len(x.ChildrenOfGroup)
	case *ListGroupsRequest_AncestorsOfGroup:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AncestorsOfGroup)))
		n += len(x.AncestorsOfGroup)
	case *ListGroupsRequest_DescendantsOfGroup:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.DescendantsOfGroup)))
		n += len(x.DescendantsOfGroup)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The `ListGroups` response.
type ListGroupsResponse struct {
	// The groups that match the specified filters.
	Group []*Group `protobuf:"bytes,1,rep,name=group" json:"group,omitempty"`
	// If there are more results than have been returned, then this field is set
	// to a non-empty value.  To see the additional results,
	// use that value as `pageToken` in the next call to this method.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListGroupsResponse) Reset()                    { *m = ListGroupsResponse{} }
func (m *ListGroupsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListGroupsResponse) ProtoMessage()               {}
func (*ListGroupsResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *ListGroupsResponse) GetGroup() []*Group {
	if m != nil {
		return m.Group
	}
	return nil
}

// The `GetGroup` request.
type GetGroupRequest struct {
	// The group to retrieve. The format is
	// `"projects/{project_id_or_number}/groups/{group_id}"`.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *GetGroupRequest) Reset()                    { *m = GetGroupRequest{} }
func (m *GetGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*GetGroupRequest) ProtoMessage()               {}
func (*GetGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

// The `CreateGroup` request.
type CreateGroupRequest struct {
	// The project in which to create the group. The format is
	// `"projects/{project_id_or_number}"`.
	Name string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	// A group definition. It is an error to define the `name` field because
	// the system assigns the name.
	Group *Group `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
	// If true, validate this request but do not create the group.
	ValidateOnly bool `protobuf:"varint,3,opt,name=validate_only,json=validateOnly" json:"validate_only,omitempty"`
}

func (m *CreateGroupRequest) Reset()                    { *m = CreateGroupRequest{} }
func (m *CreateGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateGroupRequest) ProtoMessage()               {}
func (*CreateGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *CreateGroupRequest) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

// The `UpdateGroup` request.
type UpdateGroupRequest struct {
	// The new definition of the group.  All fields of the existing group,
	// excepting `name`, are replaced with the corresponding fields of this group.
	Group *Group `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
	// If true, validate this request but do not update the existing group.
	ValidateOnly bool `protobuf:"varint,3,opt,name=validate_only,json=validateOnly" json:"validate_only,omitempty"`
}

func (m *UpdateGroupRequest) Reset()                    { *m = UpdateGroupRequest{} }
func (m *UpdateGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateGroupRequest) ProtoMessage()               {}
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *UpdateGroupRequest) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

// The `DeleteGroup` request. You can only delete a group if it has no children.
type DeleteGroupRequest struct {
	// The group to delete. The format is
	// `"projects/{project_id_or_number}/groups/{group_id}"`.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteGroupRequest) Reset()                    { *m = DeleteGroupRequest{} }
func (m *DeleteGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteGroupRequest) ProtoMessage()               {}
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

// The `ListGroupMembers` request.
type ListGroupMembersRequest struct {
	// The group whose members are listed. The format is
	// `"projects/{project_id_or_number}/groups/{group_id}"`.
	Name string `protobuf:"bytes,7,opt,name=name" json:"name,omitempty"`
	// A positive number that is the maximum number of results to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// If this field is not empty then it must contain the `nextPageToken` value
	// returned by a previous call to this method.  Using this field causes the
	// method to return additional results from the previous method call.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// An optional [list filter](/monitoring/api/learn_more#filtering) describing
	// the members to be returned.  The filter may reference the type, labels, and
	// metadata of monitored resources that comprise the group.
	// For example, to return only resources representing Compute Engine VM
	// instances, use this filter:
	//
	//     resource.type = "gce_instance"
	Filter string `protobuf:"bytes,5,opt,name=filter" json:"filter,omitempty"`
	// An optional time interval for which results should be returned. Only
	// members that were part of the group during the specified interval are
	// included in the response.  If no interval is provided then the group
	// membership over the last minute is returned.
	Interval *TimeInterval `protobuf:"bytes,6,opt,name=interval" json:"interval,omitempty"`
}

func (m *ListGroupMembersRequest) Reset()                    { *m = ListGroupMembersRequest{} }
func (m *ListGroupMembersRequest) String() string            { return proto.CompactTextString(m) }
func (*ListGroupMembersRequest) ProtoMessage()               {}
func (*ListGroupMembersRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *ListGroupMembersRequest) GetInterval() *TimeInterval {
	if m != nil {
		return m.Interval
	}
	return nil
}

// The `ListGroupMembers` response.
type ListGroupMembersResponse struct {
	// A set of monitored resources in the group.
	Members []*google_api4.MonitoredResource `protobuf:"bytes,1,rep,name=members" json:"members,omitempty"`
	// If there are more results than have been returned, then this field is
	// set to a non-empty value.  To see the additional results, use that value as
	// `pageToken` in the next call to this method.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
	// The total number of elements matching this request.
	TotalSize int32 `protobuf:"varint,3,opt,name=total_size,json=totalSize" json:"total_size,omitempty"`
}

func (m *ListGroupMembersResponse) Reset()                    { *m = ListGroupMembersResponse{} }
func (m *ListGroupMembersResponse) String() string            { return proto.CompactTextString(m) }
func (*ListGroupMembersResponse) ProtoMessage()               {}
func (*ListGroupMembersResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *ListGroupMembersResponse) GetMembers() []*google_api4.MonitoredResource {
	if m != nil {
		return m.Members
	}
	return nil
}

func init() {
	proto.RegisterType((*ListGroupsRequest)(nil), "google.monitoring.v3.ListGroupsRequest")
	proto.RegisterType((*ListGroupsResponse)(nil), "google.monitoring.v3.ListGroupsResponse")
	proto.RegisterType((*GetGroupRequest)(nil), "google.monitoring.v3.GetGroupRequest")
	proto.RegisterType((*CreateGroupRequest)(nil), "google.monitoring.v3.CreateGroupRequest")
	proto.RegisterType((*UpdateGroupRequest)(nil), "google.monitoring.v3.UpdateGroupRequest")
	proto.RegisterType((*DeleteGroupRequest)(nil), "google.monitoring.v3.DeleteGroupRequest")
	proto.RegisterType((*ListGroupMembersRequest)(nil), "google.monitoring.v3.ListGroupMembersRequest")
	proto.RegisterType((*ListGroupMembersResponse)(nil), "google.monitoring.v3.ListGroupMembersResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for GroupService service

type GroupServiceClient interface {
	// Lists the existing groups. The project ID in the URL path must refer
	// to a Stackdriver account.
	ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error)
	// Gets a single group. The project ID in the URL path must refer to a
	// Stackdriver account.
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*Group, error)
	// Creates a new group. The project ID in the URL path must refer to a
	// Stackdriver account.
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*Group, error)
	// Updates an existing group.
	// You can change any group attributes except `name`.
	// The project ID in the URL path must refer to a Stackdriver account.
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*Group, error)
	// Deletes an existing group. The project ID in the URL path must refer to a
	// Stackdriver account.
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error)
	// Lists the monitored resources that are members of a group. The project ID
	// in the URL path must refer to a Stackdriver account.
	ListGroupMembers(ctx context.Context, in *ListGroupMembersRequest, opts ...grpc.CallOption) (*ListGroupMembersResponse, error)
}

type groupServiceClient struct {
	cc *grpc.ClientConn
}

func NewGroupServiceClient(cc *grpc.ClientConn) GroupServiceClient {
	return &groupServiceClient{cc}
}

func (c *groupServiceClient) ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error) {
	out := new(ListGroupsResponse)
	err := grpc.Invoke(ctx, "/google.monitoring.v3.GroupService/ListGroups", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := grpc.Invoke(ctx, "/google.monitoring.v3.GroupService/GetGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := grpc.Invoke(ctx, "/google.monitoring.v3.GroupService/CreateGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := grpc.Invoke(ctx, "/google.monitoring.v3.GroupService/UpdateGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error) {
	out := new(google_protobuf4.Empty)
	err := grpc.Invoke(ctx, "/google.monitoring.v3.GroupService/DeleteGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) ListGroupMembers(ctx context.Context, in *ListGroupMembersRequest, opts ...grpc.CallOption) (*ListGroupMembersResponse, error) {
	out := new(ListGroupMembersResponse)
	err := grpc.Invoke(ctx, "/google.monitoring.v3.GroupService/ListGroupMembers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GroupService service

type GroupServiceServer interface {
	// Lists the existing groups. The project ID in the URL path must refer
	// to a Stackdriver account.
	ListGroups(context.Context, *ListGroupsRequest) (*ListGroupsResponse, error)
	// Gets a single group. The project ID in the URL path must refer to a
	// Stackdriver account.
	GetGroup(context.Context, *GetGroupRequest) (*Group, error)
	// Creates a new group. The project ID in the URL path must refer to a
	// Stackdriver account.
	CreateGroup(context.Context, *CreateGroupRequest) (*Group, error)
	// Updates an existing group.
	// You can change any group attributes except `name`.
	// The project ID in the URL path must refer to a Stackdriver account.
	UpdateGroup(context.Context, *UpdateGroupRequest) (*Group, error)
	// Deletes an existing group. The project ID in the URL path must refer to a
	// Stackdriver account.
	DeleteGroup(context.Context, *DeleteGroupRequest) (*google_protobuf4.Empty, error)
	// Lists the monitored resources that are members of a group. The project ID
	// in the URL path must refer to a Stackdriver account.
	ListGroupMembers(context.Context, *ListGroupMembersRequest) (*ListGroupMembersResponse, error)
}

func RegisterGroupServiceServer(s *grpc.Server, srv GroupServiceServer) {
	s.RegisterService(&_GroupService_serviceDesc, srv)
}

func _GroupService_ListGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).ListGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/ListGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).ListGroups(ctx, req.(*ListGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetGroup(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_ListGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).ListGroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/ListGroupMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).ListGroupMembers(ctx, req.(*ListGroupMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.monitoring.v3.GroupService",
	HandlerType: (*GroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListGroups",
			Handler:    _GroupService_ListGroups_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _GroupService_GetGroup_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _GroupService_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _GroupService_UpdateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _GroupService_DeleteGroup_Handler,
		},
		{
			MethodName: "ListGroupMembers",
			Handler:    _GroupService_ListGroupMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor4,
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/monitoring/v3/group_service.proto", fileDescriptor4)
}

var fileDescriptor4 = []byte{
	// 819 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x4e, 0xdb, 0x48,
	0x14, 0x5e, 0x93, 0x10, 0x92, 0x09, 0x08, 0x18, 0x21, 0x36, 0x32, 0x0b, 0x8a, 0xbc, 0xc0, 0x46,
	0xec, 0x62, 0x6b, 0x93, 0x8b, 0x95, 0x56, 0xda, 0x1f, 0xb1, 0xbb, 0x62, 0x91, 0x8a, 0x40, 0x86,
	0x5e, 0x47, 0x8e, 0x73, 0x62, 0xa6, 0xb5, 0x67, 0x5c, 0xcf, 0x24, 0x6d, 0xa8, 0x90, 0xda, 0x4a,
	0xbd, 0xe8, 0x75, 0xef, 0x7a, 0xd7, 0xe7, 0xe8, 0x55, 0x9f, 0xa1, 0xaf, 0xd0, 0xf7, 0x68, 0xe5,
	0xf1, 0x38, 0x31, 0xf9, 0x23, 0x42, 0xea, 0x4d, 0x14, 0xcf, 0x7c, 0x73, 0xbe, 0xef, 0x9c, 0xf9,
	0xce, 0x19, 0x74, 0xec, 0x31, 0xe6, 0xf9, 0x60, 0x7a, 0xcc, 0x77, 0xa8, 0x67, 0xb2, 0xc8, 0xb3,
	0x3c, 0xa0, 0x61, 0xc4, 0x04, 0xb3, 0x92, 0x2d, 0x27, 0x24, 0xdc, 0x0a, 0x18, 0x25, 0x82, 0x45,
	0x84, 0x7a, 0x56, 0xaf, 0x61, 0x79, 0x11, 0xeb, 0x86, 0x4d, 0x0e, 0x51, 0x8f, 0xb8, 0x60, 0x4a,
	0x30, 0xde, 0x50, 0x81, 0x86, 0x48, 0xb3, 0xd7, 0xd0, 0x4f, 0xe6, 0x0b, 0xef, 0x84, 0xc4, 0x52,
	0xe1, 0x5c, 0x46, 0x3b, 0xc4, 0xb3, 0x1c, 0x4a, 0x99, 0x70, 0x04, 0x61, 0x94, 0x27, 0x04, 0xfa,
	0xd9, 0xfc, 0xa1, 0x94, 0x06, 0x68, 0x47, 0xc0, 0x87, 0x1f, 0xcd, 0x08, 0x38, 0xeb, 0x46, 0xa9,
	0x62, 0xfd, 0xef, 0xfb, 0xa4, 0xee, 0xb2, 0x20, 0x60, 0x54, 0x45, 0xf8, 0xeb, 0xde, 0xc5, 0x53,
	0x01, 0x1a, 0x1e, 0x11, 0x57, 0xdd, 0x96, 0xe9, 0xb2, 0xc0, 0x4a, 0x82, 0x58, 0x72, 0xa3, 0xd5,
	0xed, 0x58, 0xa1, 0xe8, 0x87, 0xc0, 0x2d, 0x08, 0x42, 0xd1, 0x4f, 0x7e, 0x93, 0x43, 0xc6, 0x17,
	0x0d, 0xad, 0x3f, 0x20, 0x5c, 0x1c, 0xc7, 0x81, 0xb8, 0x0d, 0x4f, 0xba, 0xc0, 0x05, 0xc6, 0x28,
	0x4f, 0x9d, 0x00, 0x2a, 0x4b, 0x55, 0xad, 0x56, 0xb2, 0xe5, 0x7f, 0xfc, 0x0b, 0x5a, 0x77, 0xaf,
	0x88, 0xdf, 0x8e, 0x80, 0x36, 0x59, 0xa7, 0x29, 0x99, 0x2b, 0x0b, 0x31, 0xe0, 0xff, 0xef, 0xec,
	0xd5, 0x74, 0xeb, 0xac, 0x23, 0x23, 0x61, 0x13, 0x61, 0x87, 0xba, 0xc0, 0x05, 0x8b, 0xf8, 0x10,
	0x9e, 0x53, 0xf0, 0xb5, 0xc1, 0x5e, 0x8a, 0xaf, 0xa3, 0x8d, 0x36, 0x70, 0x17, 0x68, 0xdb, 0xa1,
	0x22, 0x73, 0x22, 0xaf, 0x4e, 0xe0, 0xcc, 0x6e, 0x7a, 0x66, 0x0b, 0x95, 0x42, 0xc7, 0x83, 0x26,
	0x27, 0xd7, 0x50, 0x59, 0xac, 0x6a, 0xb5, 0x45, 0xbb, 0x18, 0x2f, 0x5c, 0x90, 0x6b, 0xc0, 0xdb,
	0x08, 0xc9, 0x4d, 0xc1, 0x1e, 0x03, 0xad, 0x14, 0x64, 0x22, 0x12, 0x7e, 0x19, 0x2f, 0x1c, 0x15,
	0x51, 0xa1, 0x43, 0x7c, 0x01, 0x91, 0xc1, 0x10, 0xce, 0x16, 0x80, 0x87, 0x8c, 0x72, 0xc0, 0xbf,
	0xa2, 0xc5, 0x44, 0x80, 0x56, 0xcd, 0xd5, 0xca, 0xf5, 0x2d, 0x73, 0x92, 0x23, 0x4d, 0x79, 0xc8,
	0x4e, 0x90, 0x78, 0x1f, 0xad, 0x52, 0x78, 0x26, 0x9a, 0x19, 0x5a, 0x59, 0x1e, 0x7b, 0x25, 0x5e,
	0x3e, 0x4f, 0xa9, 0x8d, 0x3d, 0xb4, 0x7a, 0x0c, 0x09, 0xdf, 0x68, 0xbd, 0x73, 0xc3, 0x7a, 0x1b,
	0x2f, 0x34, 0x84, 0xff, 0x89, 0xc0, 0x11, 0x30, 0x11, 0x9a, 0xcf, 0x5c, 0xcd, 0x40, 0x6c, 0xcc,
	0x37, 0x9f, 0xd8, 0x1f, 0xd1, 0x4a, 0xcf, 0xf1, 0x49, 0xdb, 0x11, 0xd0, 0x64, 0xd4, 0xef, 0x4b,
	0xea, 0xa2, 0xbd, 0x9c, 0x2e, 0x9e, 0x51, 0xbf, 0x6f, 0xf8, 0x08, 0x3f, 0x0c, 0xdb, 0xa3, 0x0a,
	0xbe, 0x15, 0x5b, 0x0d, 0xe1, 0x7f, 0xc1, 0x87, 0x29, 0xf9, 0x66, 0x4b, 0xf3, 0x51, 0x43, 0xdf,
	0x0f, 0xee, 0xec, 0x14, 0x82, 0x16, 0x44, 0x33, 0xad, 0x7b, 0xcb, 0x28, 0xb9, 0x99, 0x46, 0xc9,
	0x8f, 0x18, 0x05, 0x6f, 0xa6, 0x46, 0x91, 0x0e, 0x2b, 0xd9, 0xea, 0x0b, 0xff, 0x89, 0x8a, 0x84,
	0x0a, 0x88, 0x7a, 0x8e, 0x2f, 0xdd, 0x55, 0xae, 0x1b, 0x93, 0x0b, 0x71, 0x49, 0x02, 0x38, 0x51,
	0x48, 0x7b, 0x70, 0xc6, 0x78, 0xa7, 0xa1, 0xca, 0x78, 0x0e, 0xca, 0x7d, 0xbf, 0xa1, 0xa5, 0x20,
	0x59, 0x52, 0xfe, 0xdb, 0x4e, 0x63, 0x3b, 0x21, 0x31, 0x4f, 0xd3, 0x21, 0x64, 0xab, 0x19, 0x64,
	0xa7, 0xe8, 0x79, 0x3d, 0x18, 0x27, 0x2d, 0x98, 0x70, 0xfc, 0x6c, 0x49, 0x4a, 0x72, 0x25, 0xae,
	0x49, 0xfd, 0x43, 0x01, 0x2d, 0x4b, 0x61, 0x17, 0xc9, 0x1c, 0xc5, 0xaf, 0x35, 0x84, 0x86, 0x5d,
	0x82, 0x7f, 0x9a, 0x9c, 0xea, 0xd8, 0x20, 0xd1, 0x6b, 0x77, 0x03, 0x93, 0x94, 0x8d, 0xdd, 0x57,
	0x9f, 0x3e, 0xbf, 0x5d, 0xd8, 0xc1, 0x3f, 0xc4, 0x63, 0xed, 0x79, 0x7c, 0x6d, 0x7f, 0x84, 0x11,
	0x7b, 0x04, 0xae, 0xe0, 0xd6, 0xc1, 0x4d, 0x32, 0xe8, 0x38, 0xee, 0xa1, 0x62, 0xda, 0x3b, 0x78,
	0x6f, 0x8a, 0xf1, 0x6e, 0xf7, 0x96, 0x3e, 0xcb, 0x9f, 0xc6, 0xbe, 0x64, 0xad, 0xe2, 0x9d, 0x49,
	0xac, 0x8a, 0xd4, 0x3a, 0xb8, 0xc1, 0x2f, 0x35, 0x54, 0xce, 0x34, 0x23, 0x9e, 0x92, 0xd7, 0x78,
	0xbf, 0xce, 0xa6, 0xff, 0x59, 0xd2, 0xef, 0x19, 0x33, 0x93, 0xfe, 0x5d, 0x35, 0xd1, 0x1b, 0x0d,
	0x95, 0x33, 0xed, 0x38, 0x4d, 0xc3, 0x78, 0xc7, 0xce, 0xd6, 0xd0, 0x90, 0x1a, 0x0e, 0xf5, 0x5d,
	0xa9, 0x21, 0x79, 0x50, 0xa6, 0x16, 0x22, 0xd5, 0xf2, 0x14, 0x95, 0x33, 0xbd, 0x3a, 0x4d, 0xca,
	0x78, 0x3b, 0xeb, 0x9b, 0x29, 0x32, 0x7d, 0x9a, 0xcc, 0xff, 0xe2, 0xd7, 0x28, 0xbd, 0x88, 0x83,
	0xbb, 0x2e, 0xe2, 0xbd, 0x86, 0xd6, 0x46, 0xdb, 0x06, 0x1f, 0xde, 0xe1, 0xb2, 0xdb, 0x23, 0x42,
	0x37, 0xe7, 0x85, 0x2b, 0x6b, 0x9a, 0x52, 0x5b, 0x0d, 0xef, 0xcf, 0xd6, 0x66, 0xa9, 0x26, 0x3c,
	0xb2, 0x50, 0xc5, 0x65, 0xc1, 0x44, 0x92, 0xa3, 0xf5, 0x6c, 0x5b, 0x9d, 0xc7, 0x35, 0x38, 0xd7,
	0x5a, 0x05, 0x59, 0x8c, 0xc6, 0xd7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x6f, 0x90, 0xee, 0x40,
	0x09, 0x00, 0x00,
}
