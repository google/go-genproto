// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/feed_mapping_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for [FeedMappingService.GetFeedMapping][google.ads.googleads.v2.services.FeedMappingService.GetFeedMapping].
type GetFeedMappingRequest struct {
	// Required. The resource name of the feed mapping to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedMappingRequest) Reset()         { *m = GetFeedMappingRequest{} }
func (m *GetFeedMappingRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedMappingRequest) ProtoMessage()    {}
func (*GetFeedMappingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_378f0144add6067b, []int{0}
}

func (m *GetFeedMappingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedMappingRequest.Unmarshal(m, b)
}
func (m *GetFeedMappingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedMappingRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedMappingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedMappingRequest.Merge(m, src)
}
func (m *GetFeedMappingRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedMappingRequest.Size(m)
}
func (m *GetFeedMappingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedMappingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedMappingRequest proto.InternalMessageInfo

func (m *GetFeedMappingRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [FeedMappingService.MutateFeedMappings][google.ads.googleads.v2.services.FeedMappingService.MutateFeedMappings].
type MutateFeedMappingsRequest struct {
	// Required. The ID of the customer whose feed mappings are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual feed mappings.
	Operations []*FeedMappingOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFeedMappingsRequest) Reset()         { *m = MutateFeedMappingsRequest{} }
func (m *MutateFeedMappingsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateFeedMappingsRequest) ProtoMessage()    {}
func (*MutateFeedMappingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_378f0144add6067b, []int{1}
}

func (m *MutateFeedMappingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedMappingsRequest.Unmarshal(m, b)
}
func (m *MutateFeedMappingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedMappingsRequest.Marshal(b, m, deterministic)
}
func (m *MutateFeedMappingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedMappingsRequest.Merge(m, src)
}
func (m *MutateFeedMappingsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateFeedMappingsRequest.Size(m)
}
func (m *MutateFeedMappingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedMappingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedMappingsRequest proto.InternalMessageInfo

func (m *MutateFeedMappingsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateFeedMappingsRequest) GetOperations() []*FeedMappingOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateFeedMappingsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateFeedMappingsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on a feed mapping.
type FeedMappingOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*FeedMappingOperation_Create
	//	*FeedMappingOperation_Remove
	Operation            isFeedMappingOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *FeedMappingOperation) Reset()         { *m = FeedMappingOperation{} }
func (m *FeedMappingOperation) String() string { return proto.CompactTextString(m) }
func (*FeedMappingOperation) ProtoMessage()    {}
func (*FeedMappingOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_378f0144add6067b, []int{2}
}

func (m *FeedMappingOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedMappingOperation.Unmarshal(m, b)
}
func (m *FeedMappingOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedMappingOperation.Marshal(b, m, deterministic)
}
func (m *FeedMappingOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedMappingOperation.Merge(m, src)
}
func (m *FeedMappingOperation) XXX_Size() int {
	return xxx_messageInfo_FeedMappingOperation.Size(m)
}
func (m *FeedMappingOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedMappingOperation.DiscardUnknown(m)
}

var xxx_messageInfo_FeedMappingOperation proto.InternalMessageInfo

type isFeedMappingOperation_Operation interface {
	isFeedMappingOperation_Operation()
}

type FeedMappingOperation_Create struct {
	Create *resources.FeedMapping `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type FeedMappingOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*FeedMappingOperation_Create) isFeedMappingOperation_Operation() {}

func (*FeedMappingOperation_Remove) isFeedMappingOperation_Operation() {}

func (m *FeedMappingOperation) GetOperation() isFeedMappingOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *FeedMappingOperation) GetCreate() *resources.FeedMapping {
	if x, ok := m.GetOperation().(*FeedMappingOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *FeedMappingOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*FeedMappingOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FeedMappingOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FeedMappingOperation_Create)(nil),
		(*FeedMappingOperation_Remove)(nil),
	}
}

// Response message for a feed mapping mutate.
type MutateFeedMappingsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateFeedMappingResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MutateFeedMappingsResponse) Reset()         { *m = MutateFeedMappingsResponse{} }
func (m *MutateFeedMappingsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateFeedMappingsResponse) ProtoMessage()    {}
func (*MutateFeedMappingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_378f0144add6067b, []int{3}
}

func (m *MutateFeedMappingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedMappingsResponse.Unmarshal(m, b)
}
func (m *MutateFeedMappingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedMappingsResponse.Marshal(b, m, deterministic)
}
func (m *MutateFeedMappingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedMappingsResponse.Merge(m, src)
}
func (m *MutateFeedMappingsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateFeedMappingsResponse.Size(m)
}
func (m *MutateFeedMappingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedMappingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedMappingsResponse proto.InternalMessageInfo

func (m *MutateFeedMappingsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateFeedMappingsResponse) GetResults() []*MutateFeedMappingResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the feed mapping mutate.
type MutateFeedMappingResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFeedMappingResult) Reset()         { *m = MutateFeedMappingResult{} }
func (m *MutateFeedMappingResult) String() string { return proto.CompactTextString(m) }
func (*MutateFeedMappingResult) ProtoMessage()    {}
func (*MutateFeedMappingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_378f0144add6067b, []int{4}
}

func (m *MutateFeedMappingResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedMappingResult.Unmarshal(m, b)
}
func (m *MutateFeedMappingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedMappingResult.Marshal(b, m, deterministic)
}
func (m *MutateFeedMappingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedMappingResult.Merge(m, src)
}
func (m *MutateFeedMappingResult) XXX_Size() int {
	return xxx_messageInfo_MutateFeedMappingResult.Size(m)
}
func (m *MutateFeedMappingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedMappingResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedMappingResult proto.InternalMessageInfo

func (m *MutateFeedMappingResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFeedMappingRequest)(nil), "google.ads.googleads.v2.services.GetFeedMappingRequest")
	proto.RegisterType((*MutateFeedMappingsRequest)(nil), "google.ads.googleads.v2.services.MutateFeedMappingsRequest")
	proto.RegisterType((*FeedMappingOperation)(nil), "google.ads.googleads.v2.services.FeedMappingOperation")
	proto.RegisterType((*MutateFeedMappingsResponse)(nil), "google.ads.googleads.v2.services.MutateFeedMappingsResponse")
	proto.RegisterType((*MutateFeedMappingResult)(nil), "google.ads.googleads.v2.services.MutateFeedMappingResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/feed_mapping_service.proto", fileDescriptor_378f0144add6067b)
}

var fileDescriptor_378f0144add6067b = []byte{
	// 707 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x09, 0x2a, 0x74, 0xd3, 0x16, 0x69, 0xa1, 0x34, 0x04, 0x24, 0x22, 0x53, 0x89, 0x28,
	0x42, 0x6b, 0x30, 0xa8, 0x08, 0x97, 0x22, 0x39, 0x12, 0x69, 0x39, 0x94, 0x56, 0xa9, 0xe8, 0x01,
	0x22, 0x59, 0x5b, 0x7b, 0x1b, 0x2c, 0xd9, 0x5e, 0xb3, 0xbb, 0x8e, 0x54, 0x55, 0xbd, 0x70, 0xe1,
	0x03, 0xf8, 0x03, 0x8e, 0xfc, 0x01, 0x47, 0xae, 0xbd, 0x72, 0xeb, 0x01, 0x71, 0x40, 0x1c, 0x10,
	0xdf, 0x80, 0x90, 0xbd, 0xde, 0xc4, 0x69, 0x13, 0x45, 0xf4, 0x36, 0x99, 0x99, 0xf7, 0xe6, 0xcd,
	0xce, 0x8c, 0x03, 0x56, 0x7b, 0x94, 0xf6, 0x02, 0x62, 0x60, 0x8f, 0x1b, 0xd2, 0x4c, 0xad, 0xbe,
	0x69, 0x70, 0xc2, 0xfa, 0xbe, 0x4b, 0xb8, 0xb1, 0x4f, 0x88, 0xe7, 0x84, 0x38, 0x8e, 0xfd, 0xa8,
	0xe7, 0xe4, 0x5e, 0x14, 0x33, 0x2a, 0x28, 0xac, 0x4b, 0x04, 0xc2, 0x1e, 0x47, 0x03, 0x30, 0xea,
	0x9b, 0x48, 0x81, 0x6b, 0x8f, 0x26, 0xd1, 0x33, 0xc2, 0x69, 0xc2, 0x4e, 0xf3, 0x4b, 0xde, 0xda,
	0x2d, 0x85, 0x8a, 0x7d, 0x03, 0x47, 0x11, 0x15, 0x58, 0xf8, 0x34, 0xe2, 0x79, 0x74, 0xa9, 0x10,
	0x75, 0x03, 0x9f, 0x44, 0x22, 0x0f, 0xdc, 0x2e, 0x04, 0xf6, 0x7d, 0x12, 0x78, 0xce, 0x1e, 0x79,
	0x8b, 0xfb, 0x3e, 0x65, 0xa7, 0x90, 0x2c, 0x76, 0x0d, 0x2e, 0xb0, 0x48, 0x72, 0x4a, 0xdd, 0x06,
	0x8b, 0xeb, 0x44, 0xb4, 0x09, 0xf1, 0x36, 0xa5, 0x90, 0x0e, 0x79, 0x97, 0x10, 0x2e, 0x60, 0x03,
	0xcc, 0x2b, 0xa5, 0x4e, 0x84, 0x43, 0x52, 0xd5, 0xea, 0x5a, 0x63, 0xb6, 0x55, 0xfe, 0x61, 0x97,
	0x3a, 0x73, 0x2a, 0xf2, 0x12, 0x87, 0x44, 0xff, 0xa3, 0x81, 0x1b, 0x9b, 0x89, 0xc0, 0x82, 0x14,
	0x68, 0xb8, 0xe2, 0x59, 0x06, 0x15, 0x37, 0xe1, 0x82, 0x86, 0x84, 0x39, 0xbe, 0x57, 0x64, 0x01,
	0xca, 0xff, 0xc2, 0x83, 0x6f, 0x00, 0xa0, 0x31, 0x61, 0xb2, 0xdb, 0x6a, 0xa9, 0x5e, 0x6e, 0x54,
	0xcc, 0x15, 0x34, 0xed, 0x91, 0x51, 0xa1, 0xe0, 0x96, 0x82, 0xe7, 0xe4, 0x43, 0x3a, 0x78, 0x17,
	0x5c, 0x89, 0x31, 0x13, 0x3e, 0x0e, 0x9c, 0x7d, 0xec, 0x07, 0x09, 0x23, 0xd5, 0x72, 0x5d, 0x6b,
	0x5c, 0xee, 0x2c, 0xe4, 0xee, 0xb6, 0xf4, 0xc2, 0x3b, 0x60, 0xbe, 0x8f, 0x03, 0xdf, 0xc3, 0x82,
	0x38, 0x34, 0x0a, 0x0e, 0xaa, 0x17, 0xb3, 0xb4, 0x39, 0xe5, 0xdc, 0x8a, 0x82, 0x03, 0xfd, 0x83,
	0x06, 0xae, 0x8d, 0xab, 0x0b, 0x37, 0xc0, 0x8c, 0xcb, 0x08, 0x16, 0xf2, 0xa9, 0x2a, 0x26, 0x9a,
	0xa8, 0x7f, 0xb0, 0x02, 0xc5, 0x06, 0x36, 0x2e, 0x74, 0x72, 0x3c, 0xac, 0x82, 0x19, 0x46, 0x42,
	0xda, 0x97, 0x3a, 0x67, 0xd3, 0x88, 0xfc, 0xdd, 0xaa, 0x80, 0xd9, 0x41, 0x63, 0xfa, 0x17, 0x0d,
	0xd4, 0xc6, 0x3d, 0x3c, 0x8f, 0x69, 0xc4, 0x09, 0x6c, 0x83, 0xc5, 0x53, 0x6d, 0x3b, 0x84, 0x31,
	0xca, 0x32, 0xd2, 0x8a, 0x09, 0x95, 0x3c, 0x16, 0xbb, 0x68, 0x27, 0xdb, 0x89, 0xce, 0xd5, 0xd1,
	0x07, 0x79, 0x9e, 0xa6, 0xc3, 0x1d, 0x70, 0x89, 0x11, 0x9e, 0x04, 0x42, 0x0d, 0xe6, 0xc9, 0xf4,
	0xc1, 0x9c, 0x91, 0xd5, 0xc9, 0x18, 0x3a, 0x8a, 0x49, 0x7f, 0x06, 0x96, 0x26, 0xe4, 0xa4, 0x53,
	0x18, 0xb3, 0x79, 0xa3, 0x4b, 0x67, 0x7e, 0x2f, 0x03, 0x58, 0x80, 0xee, 0xc8, 0xc2, 0xf0, 0xab,
	0x06, 0x16, 0x46, 0xf7, 0x19, 0x3e, 0x9e, 0xae, 0x76, 0xec, 0x05, 0xd4, 0xfe, 0x73, 0x7e, 0x7a,
	0xfb, 0xc4, 0x1e, 0x15, 0xfe, 0xfe, 0xdb, 0xcf, 0x8f, 0xa5, 0xfb, 0x10, 0xa5, 0x57, 0x7f, 0x38,
	0x12, 0x59, 0x53, 0xab, 0xcf, 0x8d, 0x66, 0xf6, 0x19, 0x50, 0xc3, 0x33, 0x9a, 0x47, 0xf0, 0x97,
	0x06, 0xe0, 0xd9, 0xb1, 0xc2, 0xd5, 0x73, 0xbc, 0xba, 0xba, 0xc2, 0xda, 0xd3, 0xf3, 0x81, 0xe5,
	0x26, 0xe9, 0xaf, 0x4e, 0xec, 0xeb, 0x85, 0x23, 0xbe, 0x37, 0xbc, 0xad, 0xac, 0xc5, 0x15, 0xfd,
	0x41, 0xda, 0xe2, 0xb0, 0xa7, 0xc3, 0x42, 0xf2, 0x5a, 0xf3, 0x68, 0xa4, 0x43, 0x2b, 0xcc, 0xea,
	0x58, 0x5a, 0xb3, 0x76, 0xf3, 0xd8, 0xae, 0x0e, 0xb5, 0xe4, 0x56, 0xec, 0x73, 0xe4, 0xd2, 0xb0,
	0xf5, 0x57, 0x03, 0xcb, 0x2e, 0x0d, 0xa7, 0xea, 0x6e, 0x2d, 0x9d, 0x5d, 0x83, 0xed, 0xf4, 0xd3,
	0xb6, 0xad, 0xbd, 0xde, 0xc8, 0xc1, 0x3d, 0x1a, 0xe0, 0xa8, 0x87, 0x28, 0xeb, 0x19, 0x3d, 0x12,
	0x65, 0x1f, 0x3e, 0x63, 0x58, 0x6e, 0xf2, 0x3f, 0xc0, 0xaa, 0x32, 0x3e, 0x95, 0xca, 0xeb, 0xb6,
	0xfd, 0xb9, 0x54, 0x5f, 0x97, 0x84, 0xb6, 0xc7, 0x91, 0x34, 0x53, 0x6b, 0xd7, 0x44, 0x79, 0x61,
	0x7e, 0xac, 0x52, 0xba, 0xb6, 0xc7, 0xbb, 0x83, 0x94, 0xee, 0xae, 0xd9, 0x55, 0x29, 0xbf, 0x4b,
	0xcb, 0xd2, 0x6f, 0x59, 0xb6, 0xc7, 0x2d, 0x6b, 0x90, 0x64, 0x59, 0xbb, 0xa6, 0x65, 0xa9, 0xb4,
	0xbd, 0x99, 0x4c, 0xe7, 0xc3, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xd1, 0x7a, 0x87, 0xa8,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FeedMappingServiceClient is the client API for FeedMappingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedMappingServiceClient interface {
	// Returns the requested feed mapping in full detail.
	GetFeedMapping(ctx context.Context, in *GetFeedMappingRequest, opts ...grpc.CallOption) (*resources.FeedMapping, error)
	// Creates or removes feed mappings. Operation statuses are
	// returned.
	MutateFeedMappings(ctx context.Context, in *MutateFeedMappingsRequest, opts ...grpc.CallOption) (*MutateFeedMappingsResponse, error)
}

type feedMappingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedMappingServiceClient(cc grpc.ClientConnInterface) FeedMappingServiceClient {
	return &feedMappingServiceClient{cc}
}

func (c *feedMappingServiceClient) GetFeedMapping(ctx context.Context, in *GetFeedMappingRequest, opts ...grpc.CallOption) (*resources.FeedMapping, error) {
	out := new(resources.FeedMapping)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.FeedMappingService/GetFeedMapping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedMappingServiceClient) MutateFeedMappings(ctx context.Context, in *MutateFeedMappingsRequest, opts ...grpc.CallOption) (*MutateFeedMappingsResponse, error) {
	out := new(MutateFeedMappingsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.FeedMappingService/MutateFeedMappings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedMappingServiceServer is the server API for FeedMappingService service.
type FeedMappingServiceServer interface {
	// Returns the requested feed mapping in full detail.
	GetFeedMapping(context.Context, *GetFeedMappingRequest) (*resources.FeedMapping, error)
	// Creates or removes feed mappings. Operation statuses are
	// returned.
	MutateFeedMappings(context.Context, *MutateFeedMappingsRequest) (*MutateFeedMappingsResponse, error)
}

// UnimplementedFeedMappingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFeedMappingServiceServer struct {
}

func (*UnimplementedFeedMappingServiceServer) GetFeedMapping(ctx context.Context, req *GetFeedMappingRequest) (*resources.FeedMapping, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetFeedMapping not implemented")
}
func (*UnimplementedFeedMappingServiceServer) MutateFeedMappings(ctx context.Context, req *MutateFeedMappingsRequest) (*MutateFeedMappingsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateFeedMappings not implemented")
}

func RegisterFeedMappingServiceServer(s *grpc.Server, srv FeedMappingServiceServer) {
	s.RegisterService(&_FeedMappingService_serviceDesc, srv)
}

func _FeedMappingService_GetFeedMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedMappingServiceServer).GetFeedMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.FeedMappingService/GetFeedMapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedMappingServiceServer).GetFeedMapping(ctx, req.(*GetFeedMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedMappingService_MutateFeedMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedMappingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedMappingServiceServer).MutateFeedMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.FeedMappingService/MutateFeedMappings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedMappingServiceServer).MutateFeedMappings(ctx, req.(*MutateFeedMappingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedMappingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.FeedMappingService",
	HandlerType: (*FeedMappingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeedMapping",
			Handler:    _FeedMappingService_GetFeedMapping_Handler,
		},
		{
			MethodName: "MutateFeedMappings",
			Handler:    _FeedMappingService_MutateFeedMappings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/feed_mapping_service.proto",
}
