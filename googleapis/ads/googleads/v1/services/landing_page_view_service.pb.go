// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/landing_page_view_service.proto

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

// Request message for [LandingPageViewService.GetLandingPageView][google.ads.googleads.v1.services.LandingPageViewService.GetLandingPageView].
type GetLandingPageViewRequest struct {
	// The resource name of the landing page view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLandingPageViewRequest) Reset()         { *m = GetLandingPageViewRequest{} }
func (m *GetLandingPageViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetLandingPageViewRequest) ProtoMessage()    {}
func (*GetLandingPageViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_landing_page_view_service_969504292394555a, []int{0}
}
func (m *GetLandingPageViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLandingPageViewRequest.Unmarshal(m, b)
}
func (m *GetLandingPageViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLandingPageViewRequest.Marshal(b, m, deterministic)
}
func (dst *GetLandingPageViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLandingPageViewRequest.Merge(dst, src)
}
func (m *GetLandingPageViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetLandingPageViewRequest.Size(m)
}
func (m *GetLandingPageViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLandingPageViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLandingPageViewRequest proto.InternalMessageInfo

func (m *GetLandingPageViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetLandingPageViewRequest)(nil), "google.ads.googleads.v1.services.GetLandingPageViewRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LandingPageViewServiceClient is the client API for LandingPageViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LandingPageViewServiceClient interface {
	// Returns the requested landing page view in full detail.
	GetLandingPageView(ctx context.Context, in *GetLandingPageViewRequest, opts ...grpc.CallOption) (*resources.LandingPageView, error)
}

type landingPageViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewLandingPageViewServiceClient(cc *grpc.ClientConn) LandingPageViewServiceClient {
	return &landingPageViewServiceClient{cc}
}

func (c *landingPageViewServiceClient) GetLandingPageView(ctx context.Context, in *GetLandingPageViewRequest, opts ...grpc.CallOption) (*resources.LandingPageView, error) {
	out := new(resources.LandingPageView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.LandingPageViewService/GetLandingPageView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LandingPageViewServiceServer is the server API for LandingPageViewService service.
type LandingPageViewServiceServer interface {
	// Returns the requested landing page view in full detail.
	GetLandingPageView(context.Context, *GetLandingPageViewRequest) (*resources.LandingPageView, error)
}

func RegisterLandingPageViewServiceServer(s *grpc.Server, srv LandingPageViewServiceServer) {
	s.RegisterService(&_LandingPageViewService_serviceDesc, srv)
}

func _LandingPageViewService_GetLandingPageView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLandingPageViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LandingPageViewServiceServer).GetLandingPageView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.LandingPageViewService/GetLandingPageView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LandingPageViewServiceServer).GetLandingPageView(ctx, req.(*GetLandingPageViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LandingPageViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.LandingPageViewService",
	HandlerType: (*LandingPageViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLandingPageView",
			Handler:    _LandingPageViewService_GetLandingPageView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/landing_page_view_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/landing_page_view_service.proto", fileDescriptor_landing_page_view_service_969504292394555a)
}

var fileDescriptor_landing_page_view_service_969504292394555a = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xbf, 0x6a, 0xdb, 0x40,
	0x18, 0x47, 0x2a, 0x14, 0x2a, 0xda, 0x45, 0x43, 0x69, 0xdd, 0x52, 0x8c, 0xeb, 0xa1, 0x78, 0xb8,
	0x43, 0x6e, 0x97, 0x9e, 0x09, 0x58, 0x5e, 0x9c, 0x21, 0x04, 0xe3, 0x80, 0x86, 0x20, 0x10, 0x67,
	0xeb, 0xcb, 0x21, 0xb0, 0xee, 0x94, 0x3b, 0x49, 0x1e, 0x42, 0x16, 0xbf, 0x42, 0xde, 0x20, 0x63,
	0xde, 0x22, 0x6b, 0xd6, 0xac, 0x19, 0x33, 0xe5, 0x29, 0x82, 0x7c, 0x3a, 0x41, 0x1c, 0x0b, 0x6f,
	0x3f, 0xdd, 0xf7, 0xfb, 0xf3, 0xdd, 0xef, 0xe4, 0x8c, 0x99, 0x10, 0x6c, 0x05, 0x98, 0xc6, 0x0a,
	0x6b, 0x58, 0xa1, 0xd2, 0xc3, 0x0a, 0x64, 0x99, 0x2c, 0x41, 0xe1, 0x15, 0xe5, 0x71, 0xc2, 0x59,
	0x94, 0x51, 0x06, 0x51, 0x99, 0xc0, 0x3a, 0xaa, 0x47, 0x28, 0x93, 0x22, 0x17, 0x6e, 0x57, 0xcb,
	0x10, 0x8d, 0x15, 0x6a, 0x1c, 0x50, 0xe9, 0x21, 0xe3, 0xd0, 0xf9, 0xdf, 0x96, 0x21, 0x41, 0x89,
	0x42, 0xee, 0x0d, 0xd1, 0xe6, 0x9d, 0x9f, 0x46, 0x9a, 0x25, 0x98, 0x72, 0x2e, 0x72, 0x9a, 0x27,
	0x82, 0xab, 0x7a, 0xfa, 0xab, 0x9e, 0x6e, 0xbf, 0x16, 0xc5, 0x05, 0x5e, 0x4b, 0x9a, 0x65, 0x20,
	0xeb, 0x79, 0x6f, 0xec, 0x7c, 0x9f, 0x42, 0x7e, 0xa2, 0xbd, 0x67, 0x94, 0x41, 0x90, 0xc0, 0x7a,
	0x0e, 0x97, 0x05, 0xa8, 0xdc, 0xfd, 0xed, 0x7c, 0x31, 0xf9, 0x11, 0xa7, 0x29, 0x7c, 0xb3, 0xba,
	0xd6, 0x9f, 0x4f, 0xf3, 0xcf, 0xe6, 0xf0, 0x94, 0xa6, 0x30, 0x7c, 0xb2, 0x9c, 0xaf, 0x3b, 0xfa,
	0x33, 0x7d, 0x2d, 0xf7, 0xde, 0x72, 0xdc, 0xf7, 0xee, 0xee, 0x08, 0x1d, 0xea, 0x03, 0xb5, 0xee,
	0xd4, 0x19, 0xb6, 0x8a, 0x9b, 0xaa, 0xd0, 0x8e, 0xb4, 0x47, 0x36, 0x8f, 0xcf, 0x37, 0xf6, 0x3f,
	0x77, 0x58, 0x35, 0x7a, 0xf5, 0xe6, 0x4a, 0x47, 0xcb, 0x42, 0xe5, 0x22, 0x05, 0xa9, 0xf0, 0xc0,
	0x54, 0x6c, 0x74, 0x0a, 0x0f, 0xae, 0x27, 0x1b, 0xdb, 0xe9, 0x2f, 0x45, 0x7a, 0x70, 0xe5, 0xc9,
	0x8f, 0xfd, 0x25, 0xcc, 0xaa, 0x9a, 0x67, 0xd6, 0xf9, 0x71, 0x6d, 0xc0, 0xc4, 0x8a, 0x72, 0x86,
	0x84, 0x64, 0x98, 0x01, 0xdf, 0x3e, 0x82, 0x79, 0xf1, 0x2c, 0x51, 0xed, 0x3f, 0xd9, 0xc8, 0x80,
	0x5b, 0xfb, 0xc3, 0xd4, 0xf7, 0xef, 0xec, 0xee, 0x54, 0x1b, 0xfa, 0xb1, 0x42, 0x1a, 0x56, 0x28,
	0xf0, 0x50, 0x1d, 0xac, 0x1e, 0x0c, 0x25, 0xf4, 0x63, 0x15, 0x36, 0x94, 0x30, 0xf0, 0x42, 0x43,
	0x79, 0xb1, 0xfb, 0xfa, 0x9c, 0x10, 0x3f, 0x56, 0x84, 0x34, 0x24, 0x42, 0x02, 0x8f, 0x10, 0x43,
	0x5b, 0x7c, 0xdc, 0xee, 0xf9, 0xf7, 0x35, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xa1, 0x45, 0x50, 0x0b,
	0x03, 0x00, 0x00,
}
