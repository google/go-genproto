// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/customer_client_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [CustomerClientService.GetCustomerClient][google.ads.googleads.v2.services.CustomerClientService.GetCustomerClient].
type GetCustomerClientRequest struct {
	// Required. The resource name of the client to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerClientRequest) Reset()         { *m = GetCustomerClientRequest{} }
func (m *GetCustomerClientRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerClientRequest) ProtoMessage()    {}
func (*GetCustomerClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfd3ae0efb37b5a5, []int{0}
}

func (m *GetCustomerClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerClientRequest.Unmarshal(m, b)
}
func (m *GetCustomerClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerClientRequest.Marshal(b, m, deterministic)
}
func (m *GetCustomerClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerClientRequest.Merge(m, src)
}
func (m *GetCustomerClientRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerClientRequest.Size(m)
}
func (m *GetCustomerClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerClientRequest proto.InternalMessageInfo

func (m *GetCustomerClientRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCustomerClientRequest)(nil), "google.ads.googleads.v2.services.GetCustomerClientRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/customer_client_service.proto", fileDescriptor_dfd3ae0efb37b5a5)
}

var fileDescriptor_dfd3ae0efb37b5a5 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x25, 0x29, 0x3c, 0x78, 0xe1, 0xbd, 0xc5, 0x0b, 0x3c, 0x2c, 0x51, 0xb0, 0x94, 0x2e, 0x4a,
	0x17, 0x33, 0x34, 0x5d, 0x08, 0x23, 0x0a, 0xd3, 0x0a, 0xd5, 0x8d, 0x94, 0x0a, 0x5d, 0x48, 0x20,
	0x4c, 0x93, 0x31, 0x06, 0x92, 0x4c, 0xcd, 0xa4, 0xd9, 0x88, 0x0b, 0xc5, 0x3f, 0xf0, 0x0f, 0x5c,
	0xfa, 0x29, 0x05, 0x57, 0xee, 0x5c, 0xb9, 0x70, 0xe5, 0x57, 0x48, 0x3a, 0x99, 0xb4, 0xc1, 0x86,
	0xee, 0x0e, 0x39, 0xe7, 0xdc, 0x33, 0xf7, 0xdc, 0x68, 0xc7, 0x1e, 0x63, 0x5e, 0x40, 0x21, 0x71,
	0x39, 0x14, 0x30, 0x43, 0xa9, 0x09, 0x39, 0x8d, 0x53, 0xdf, 0xa1, 0x1c, 0x3a, 0x73, 0x9e, 0xb0,
	0x90, 0xc6, 0xb6, 0x13, 0xf8, 0x34, 0x4a, 0xec, 0x9c, 0x00, 0xb3, 0x98, 0x25, 0x4c, 0x6f, 0x08,
	0x13, 0x20, 0x2e, 0x07, 0x85, 0x1f, 0xa4, 0x26, 0x90, 0x7e, 0xe3, 0xa0, 0x2a, 0x21, 0xa6, 0x9c,
	0xcd, 0xe3, 0x0d, 0x11, 0x62, 0xb4, 0xb1, 0x27, 0x8d, 0x33, 0x1f, 0x92, 0x28, 0x62, 0x09, 0x49,
	0x7c, 0x16, 0xf1, 0x9c, 0xdd, 0x59, 0x63, 0x4b, 0xb6, 0xfd, 0x35, 0xe2, 0xca, 0xa7, 0x81, 0x6b,
	0x4f, 0xe9, 0x35, 0x49, 0x7d, 0x16, 0x0b, 0x41, 0xf3, 0x44, 0xab, 0x0f, 0x69, 0x32, 0xc8, 0x33,
	0x07, 0x4b, 0xef, 0x98, 0xde, 0xcc, 0x29, 0x4f, 0xf4, 0xb6, 0xf6, 0x57, 0x3e, 0xcb, 0x8e, 0x48,
	0x48, 0xeb, 0x4a, 0x43, 0x69, 0xff, 0xee, 0xd7, 0x3e, 0xb0, 0x3a, 0xfe, 0x23, 0x99, 0x73, 0x12,
	0x52, 0xf3, 0x51, 0xd5, 0xfe, 0x97, 0x67, 0x5c, 0x88, 0x8d, 0xf5, 0x57, 0x45, 0xfb, 0xf7, 0x23,
	0x40, 0x47, 0x60, 0x5b, 0x53, 0xa0, 0xea, 0x55, 0x46, 0xb7, 0xd2, 0x5b, 0x74, 0x08, 0xca, 0xce,
	0xe6, 0xd9, 0x3b, 0x2e, 0x6f, 0xf2, 0xf0, 0xf6, 0xf9, 0xa4, 0xf6, 0xf4, 0x6e, 0xd6, 0xfc, 0x6d,
	0x89, 0x39, 0x92, 0xf5, 0x73, 0xd8, 0x29, 0x4e, 0x21, 0xc6, 0x70, 0xd8, 0xb9, 0x33, 0x76, 0x17,
	0xb8, 0xbe, 0x0a, 0xcd, 0xd1, 0xcc, 0xe7, 0xc0, 0x61, 0x61, 0xff, 0x5e, 0xd5, 0x5a, 0x0e, 0x0b,
	0xb7, 0x2e, 0xd7, 0x37, 0x36, 0x96, 0x35, 0xca, 0x2e, 0x32, 0x52, 0x2e, 0x4f, 0x73, 0xbf, 0xc7,
	0x02, 0x12, 0x79, 0x80, 0xc5, 0x1e, 0xf4, 0x68, 0xb4, 0xbc, 0x17, 0x5c, 0x25, 0x56, 0xff, 0xa5,
	0x87, 0x12, 0x3c, 0xab, 0xb5, 0x21, 0xc6, 0x2f, 0x6a, 0x63, 0x28, 0x06, 0x62, 0x97, 0x03, 0x01,
	0x33, 0x34, 0x31, 0x41, 0x1e, 0xcc, 0x17, 0x52, 0x62, 0x61, 0x97, 0x5b, 0x85, 0xc4, 0x9a, 0x98,
	0x96, 0x94, 0x7c, 0xa9, 0x2d, 0xf1, 0x1d, 0x21, 0xec, 0x72, 0x84, 0x0a, 0x11, 0x42, 0x13, 0x13,
	0x21, 0x29, 0x9b, 0xfe, 0x5a, 0xbe, 0xb3, 0xf7, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x0f, 0xbf,
	0xe8, 0x4c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerClientServiceClient is the client API for CustomerClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClientServiceClient interface {
	// Returns the requested client in full detail.
	GetCustomerClient(ctx context.Context, in *GetCustomerClientRequest, opts ...grpc.CallOption) (*resources.CustomerClient, error)
}

type customerClientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerClientServiceClient(cc grpc.ClientConnInterface) CustomerClientServiceClient {
	return &customerClientServiceClient{cc}
}

func (c *customerClientServiceClient) GetCustomerClient(ctx context.Context, in *GetCustomerClientRequest, opts ...grpc.CallOption) (*resources.CustomerClient, error) {
	out := new(resources.CustomerClient)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CustomerClientService/GetCustomerClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerClientServiceServer is the server API for CustomerClientService service.
type CustomerClientServiceServer interface {
	// Returns the requested client in full detail.
	GetCustomerClient(context.Context, *GetCustomerClientRequest) (*resources.CustomerClient, error)
}

// UnimplementedCustomerClientServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerClientServiceServer struct {
}

func (*UnimplementedCustomerClientServiceServer) GetCustomerClient(ctx context.Context, req *GetCustomerClientRequest) (*resources.CustomerClient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerClient not implemented")
}

func RegisterCustomerClientServiceServer(s *grpc.Server, srv CustomerClientServiceServer) {
	s.RegisterService(&_CustomerClientService_serviceDesc, srv)
}

func _CustomerClientService_GetCustomerClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerClientServiceServer).GetCustomerClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CustomerClientService/GetCustomerClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerClientServiceServer).GetCustomerClient(ctx, req.(*GetCustomerClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.CustomerClientService",
	HandlerType: (*CustomerClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerClient",
			Handler:    _CustomerClientService_GetCustomerClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/customer_client_service.proto",
}
