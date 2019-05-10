// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/expanded_landing_page_view_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v1/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for
// [ExpandedLandingPageViewService.GetExpandedLandingPageView][google.ads.googleads.v1.services.ExpandedLandingPageViewService.GetExpandedLandingPageView].
type GetExpandedLandingPageViewRequest struct {
	// The resource name of the expanded landing page view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetExpandedLandingPageViewRequest) Reset()         { *m = GetExpandedLandingPageViewRequest{} }
func (m *GetExpandedLandingPageViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetExpandedLandingPageViewRequest) ProtoMessage()    {}
func (*GetExpandedLandingPageViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_expanded_landing_page_view_service_c886b37cf2651c3b, []int{0}
}
func (m *GetExpandedLandingPageViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetExpandedLandingPageViewRequest.Unmarshal(m, b)
}
func (m *GetExpandedLandingPageViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetExpandedLandingPageViewRequest.Marshal(b, m, deterministic)
}
func (dst *GetExpandedLandingPageViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExpandedLandingPageViewRequest.Merge(dst, src)
}
func (m *GetExpandedLandingPageViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetExpandedLandingPageViewRequest.Size(m)
}
func (m *GetExpandedLandingPageViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExpandedLandingPageViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetExpandedLandingPageViewRequest proto.InternalMessageInfo

func (m *GetExpandedLandingPageViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetExpandedLandingPageViewRequest)(nil), "google.ads.googleads.v1.services.GetExpandedLandingPageViewRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExpandedLandingPageViewServiceClient is the client API for ExpandedLandingPageViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExpandedLandingPageViewServiceClient interface {
	// Returns the requested expanded landing page view in full detail.
	GetExpandedLandingPageView(ctx context.Context, in *GetExpandedLandingPageViewRequest, opts ...grpc.CallOption) (*resources.ExpandedLandingPageView, error)
}

type expandedLandingPageViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewExpandedLandingPageViewServiceClient(cc *grpc.ClientConn) ExpandedLandingPageViewServiceClient {
	return &expandedLandingPageViewServiceClient{cc}
}

func (c *expandedLandingPageViewServiceClient) GetExpandedLandingPageView(ctx context.Context, in *GetExpandedLandingPageViewRequest, opts ...grpc.CallOption) (*resources.ExpandedLandingPageView, error) {
	out := new(resources.ExpandedLandingPageView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ExpandedLandingPageViewService/GetExpandedLandingPageView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpandedLandingPageViewServiceServer is the server API for ExpandedLandingPageViewService service.
type ExpandedLandingPageViewServiceServer interface {
	// Returns the requested expanded landing page view in full detail.
	GetExpandedLandingPageView(context.Context, *GetExpandedLandingPageViewRequest) (*resources.ExpandedLandingPageView, error)
}

func RegisterExpandedLandingPageViewServiceServer(s *grpc.Server, srv ExpandedLandingPageViewServiceServer) {
	s.RegisterService(&_ExpandedLandingPageViewService_serviceDesc, srv)
}

func _ExpandedLandingPageViewService_GetExpandedLandingPageView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExpandedLandingPageViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpandedLandingPageViewServiceServer).GetExpandedLandingPageView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ExpandedLandingPageViewService/GetExpandedLandingPageView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpandedLandingPageViewServiceServer).GetExpandedLandingPageView(ctx, req.(*GetExpandedLandingPageViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExpandedLandingPageViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.ExpandedLandingPageViewService",
	HandlerType: (*ExpandedLandingPageViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExpandedLandingPageView",
			Handler:    _ExpandedLandingPageViewService_GetExpandedLandingPageView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/expanded_landing_page_view_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/expanded_landing_page_view_service.proto", fileDescriptor_expanded_landing_page_view_service_c886b37cf2651c3b)
}

var fileDescriptor_expanded_landing_page_view_service_c886b37cf2651c3b = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x8b, 0xdb, 0x30,
	0x14, 0xc6, 0x2e, 0x14, 0x6a, 0xda, 0xc5, 0x53, 0x31, 0x25, 0xa4, 0x49, 0x86, 0x92, 0x41, 0xc2,
	0xed, 0xa6, 0xb6, 0x83, 0x5d, 0x4a, 0x52, 0x28, 0x25, 0xa4, 0xe0, 0xa1, 0x18, 0x8c, 0x12, 0xbd,
	0x0a, 0x43, 0x2c, 0xb9, 0x92, 0xed, 0x14, 0x4a, 0x97, 0xd2, 0xe9, 0xd6, 0xfb, 0x07, 0x37, 0xde,
	0x4f, 0xb9, 0xf5, 0x7e, 0xc1, 0xc1, 0x4d, 0xf7, 0x2b, 0x0e, 0x47, 0x96, 0xe1, 0x0e, 0x9c, 0x6c,
	0x9f, 0xf4, 0x3e, 0x7d, 0xdf, 0x7b, 0xef, 0x93, 0xf7, 0x85, 0x4b, 0xc9, 0x77, 0x80, 0x29, 0xd3,
	0xd8, 0xc0, 0x16, 0x35, 0x21, 0xd6, 0xa0, 0x9a, 0x7c, 0x0b, 0x1a, 0xc3, 0xef, 0x92, 0x0a, 0x06,
	0x2c, 0xdb, 0x51, 0xc1, 0x72, 0xc1, 0xb3, 0x92, 0x72, 0xc8, 0x9a, 0x1c, 0xf6, 0x59, 0xc7, 0x41,
	0xa5, 0x92, 0x95, 0xf4, 0xc7, 0xe6, 0x3d, 0xa2, 0x4c, 0xa3, 0x5e, 0x0a, 0x35, 0x21, 0xb2, 0x52,
	0x41, 0x3c, 0x64, 0xa6, 0x40, 0xcb, 0x5a, 0x1d, 0x77, 0x33, 0x2e, 0xc1, 0x2b, 0xab, 0x51, 0xe6,
	0x98, 0x0a, 0x21, 0x2b, 0x5a, 0xe5, 0x52, 0xe8, 0xae, 0x3a, 0xea, 0xaa, 0x87, 0xd3, 0xa6, 0xfe,
	0x89, 0xf7, 0x8a, 0x96, 0x25, 0xa8, 0xae, 0x3e, 0x59, 0x7a, 0xaf, 0x17, 0x50, 0x7d, 0xee, 0x4c,
	0xbe, 0x1a, 0x8f, 0x15, 0xe5, 0x90, 0xe4, 0xb0, 0x5f, 0xc3, 0xaf, 0x1a, 0x74, 0xe5, 0x4f, 0xbd,
	0x17, 0xb6, 0xa1, 0x4c, 0xd0, 0x02, 0x5e, 0x3a, 0x63, 0xe7, 0xcd, 0xb3, 0xf5, 0x73, 0x7b, 0xf9,
	0x8d, 0x16, 0xf0, 0xf6, 0xbf, 0xeb, 0x8d, 0x06, 0x74, 0xbe, 0x9b, 0x79, 0xfd, 0x1b, 0xc7, 0x0b,
	0x86, 0xdd, 0xfc, 0x4f, 0xe8, 0xd4, 0xc2, 0xd0, 0xc9, 0x5e, 0x03, 0x32, 0x28, 0xd2, 0xef, 0x14,
	0x0d, 0x48, 0x4c, 0xe2, 0x7f, 0xd7, 0xb7, 0xe7, 0xee, 0x07, 0x9f, 0xb4, 0x11, 0xfc, 0x79, 0x30,
	0xf2, 0xc7, 0x6d, 0xad, 0x2b, 0x59, 0x80, 0xd2, 0x78, 0xde, 0x67, 0xf2, 0xe8, 0xbd, 0xc6, 0xf3,
	0xbf, 0xf1, 0x99, 0xeb, 0xcd, 0xb6, 0xb2, 0x38, 0x39, 0x4a, 0x3c, 0x3d, 0xbe, 0xac, 0x55, 0x1b,
	0xcf, 0xca, 0xf9, 0xb1, 0xec, 0x84, 0xb8, 0xdc, 0x51, 0xc1, 0x91, 0x54, 0x1c, 0x73, 0x10, 0x87,
	0xf0, 0xec, 0x97, 0x29, 0x73, 0x3d, 0xfc, 0x5d, 0xdf, 0x5b, 0x70, 0xe1, 0x3e, 0x59, 0x44, 0xd1,
	0xa5, 0x3b, 0x5e, 0x18, 0xc1, 0x88, 0x69, 0x64, 0x60, 0x8b, 0x92, 0x10, 0x75, 0xc6, 0xfa, 0xca,
	0x52, 0xd2, 0x88, 0xe9, 0xb4, 0xa7, 0xa4, 0x49, 0x98, 0x5a, 0xca, 0x9d, 0x3b, 0x33, 0xf7, 0x84,
	0x44, 0x4c, 0x13, 0xd2, 0x93, 0x08, 0x49, 0x42, 0x42, 0x2c, 0x6d, 0xf3, 0xf4, 0xd0, 0xe7, 0xbb,
	0xfb, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xc2, 0x24, 0x75, 0x55, 0x03, 0x00, 0x00,
}
