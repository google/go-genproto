// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/ad_group_label_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
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

// Request message for [AdGroupLabelService.GetAdGroupLabel][google.ads.googleads.v3.services.AdGroupLabelService.GetAdGroupLabel].
type GetAdGroupLabelRequest struct {
	// Required. The resource name of the ad group label to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupLabelRequest) Reset()         { *m = GetAdGroupLabelRequest{} }
func (m *GetAdGroupLabelRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupLabelRequest) ProtoMessage()    {}
func (*GetAdGroupLabelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7054d11c8a405fc3, []int{0}
}

func (m *GetAdGroupLabelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupLabelRequest.Unmarshal(m, b)
}
func (m *GetAdGroupLabelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupLabelRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupLabelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupLabelRequest.Merge(m, src)
}
func (m *GetAdGroupLabelRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupLabelRequest.Size(m)
}
func (m *GetAdGroupLabelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupLabelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupLabelRequest proto.InternalMessageInfo

func (m *GetAdGroupLabelRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AdGroupLabelService.MutateAdGroupLabels][google.ads.googleads.v3.services.AdGroupLabelService.MutateAdGroupLabels].
type MutateAdGroupLabelsRequest struct {
	// Required. ID of the customer whose ad group labels are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on ad group labels.
	Operations []*AdGroupLabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateAdGroupLabelsRequest) Reset()         { *m = MutateAdGroupLabelsRequest{} }
func (m *MutateAdGroupLabelsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupLabelsRequest) ProtoMessage()    {}
func (*MutateAdGroupLabelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7054d11c8a405fc3, []int{1}
}

func (m *MutateAdGroupLabelsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupLabelsRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupLabelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupLabelsRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupLabelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupLabelsRequest.Merge(m, src)
}
func (m *MutateAdGroupLabelsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupLabelsRequest.Size(m)
}
func (m *MutateAdGroupLabelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupLabelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupLabelsRequest proto.InternalMessageInfo

func (m *MutateAdGroupLabelsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupLabelsRequest) GetOperations() []*AdGroupLabelOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateAdGroupLabelsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateAdGroupLabelsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on an ad group label.
type AdGroupLabelOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupLabelOperation_Create
	//	*AdGroupLabelOperation_Remove
	Operation            isAdGroupLabelOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *AdGroupLabelOperation) Reset()         { *m = AdGroupLabelOperation{} }
func (m *AdGroupLabelOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupLabelOperation) ProtoMessage()    {}
func (*AdGroupLabelOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_7054d11c8a405fc3, []int{2}
}

func (m *AdGroupLabelOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupLabelOperation.Unmarshal(m, b)
}
func (m *AdGroupLabelOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupLabelOperation.Marshal(b, m, deterministic)
}
func (m *AdGroupLabelOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupLabelOperation.Merge(m, src)
}
func (m *AdGroupLabelOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupLabelOperation.Size(m)
}
func (m *AdGroupLabelOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupLabelOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupLabelOperation proto.InternalMessageInfo

type isAdGroupLabelOperation_Operation interface {
	isAdGroupLabelOperation_Operation()
}

type AdGroupLabelOperation_Create struct {
	Create *resources.AdGroupLabel `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupLabelOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*AdGroupLabelOperation_Create) isAdGroupLabelOperation_Operation() {}

func (*AdGroupLabelOperation_Remove) isAdGroupLabelOperation_Operation() {}

func (m *AdGroupLabelOperation) GetOperation() isAdGroupLabelOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupLabelOperation) GetCreate() *resources.AdGroupLabel {
	if x, ok := m.GetOperation().(*AdGroupLabelOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupLabelOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupLabelOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupLabelOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupLabelOperation_Create)(nil),
		(*AdGroupLabelOperation_Remove)(nil),
	}
}

// Response message for an ad group labels mutate.
type MutateAdGroupLabelsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateAdGroupLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MutateAdGroupLabelsResponse) Reset()         { *m = MutateAdGroupLabelsResponse{} }
func (m *MutateAdGroupLabelsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupLabelsResponse) ProtoMessage()    {}
func (*MutateAdGroupLabelsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7054d11c8a405fc3, []int{3}
}

func (m *MutateAdGroupLabelsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupLabelsResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupLabelsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupLabelsResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupLabelsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupLabelsResponse.Merge(m, src)
}
func (m *MutateAdGroupLabelsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupLabelsResponse.Size(m)
}
func (m *MutateAdGroupLabelsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupLabelsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupLabelsResponse proto.InternalMessageInfo

func (m *MutateAdGroupLabelsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateAdGroupLabelsResponse) GetResults() []*MutateAdGroupLabelResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for an ad group label mutate.
type MutateAdGroupLabelResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupLabelResult) Reset()         { *m = MutateAdGroupLabelResult{} }
func (m *MutateAdGroupLabelResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupLabelResult) ProtoMessage()    {}
func (*MutateAdGroupLabelResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_7054d11c8a405fc3, []int{4}
}

func (m *MutateAdGroupLabelResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupLabelResult.Unmarshal(m, b)
}
func (m *MutateAdGroupLabelResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupLabelResult.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupLabelResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupLabelResult.Merge(m, src)
}
func (m *MutateAdGroupLabelResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupLabelResult.Size(m)
}
func (m *MutateAdGroupLabelResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupLabelResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupLabelResult proto.InternalMessageInfo

func (m *MutateAdGroupLabelResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupLabelRequest)(nil), "google.ads.googleads.v3.services.GetAdGroupLabelRequest")
	proto.RegisterType((*MutateAdGroupLabelsRequest)(nil), "google.ads.googleads.v3.services.MutateAdGroupLabelsRequest")
	proto.RegisterType((*AdGroupLabelOperation)(nil), "google.ads.googleads.v3.services.AdGroupLabelOperation")
	proto.RegisterType((*MutateAdGroupLabelsResponse)(nil), "google.ads.googleads.v3.services.MutateAdGroupLabelsResponse")
	proto.RegisterType((*MutateAdGroupLabelResult)(nil), "google.ads.googleads.v3.services.MutateAdGroupLabelResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/ad_group_label_service.proto", fileDescriptor_7054d11c8a405fc3)
}

var fileDescriptor_7054d11c8a405fc3 = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0x77, 0x37, 0x52, 0xed, 0xa4, 0xb5, 0x30, 0xa5, 0x75, 0x49, 0x05, 0xc3, 0x5a, 0x30, 0x04,
	0xd9, 0xc5, 0x04, 0xac, 0xac, 0x06, 0xd9, 0x80, 0x4d, 0x0b, 0x6a, 0xcb, 0x56, 0x72, 0x90, 0xc0,
	0x32, 0xdd, 0x9d, 0xae, 0x0b, 0xbb, 0x3b, 0xeb, 0xcc, 0x6c, 0xa0, 0x94, 0x5e, 0xbc, 0x79, 0xf6,
	0x1b, 0x78, 0xf4, 0x2b, 0x78, 0xf6, 0xd2, 0xab, 0xb7, 0x9e, 0x3c, 0x08, 0x42, 0x0f, 0x7e, 0x05,
	0x65, 0xff, 0x4c, 0xb2, 0x69, 0x13, 0x82, 0xbd, 0xbd, 0xbc, 0x3f, 0xbf, 0xf7, 0xfb, 0xbd, 0xf7,
	0x66, 0x03, 0x3a, 0x1e, 0x21, 0x5e, 0x80, 0x75, 0xe4, 0x32, 0x3d, 0x37, 0x53, 0x6b, 0xd8, 0xd6,
	0x19, 0xa6, 0x43, 0xdf, 0xc1, 0x4c, 0x47, 0xae, 0xed, 0x51, 0x92, 0xc4, 0x76, 0x80, 0x0e, 0x71,
	0x60, 0x17, 0x7e, 0x2d, 0xa6, 0x84, 0x13, 0x58, 0xcf, 0x6b, 0x34, 0xe4, 0x32, 0x6d, 0x54, 0xae,
	0x0d, 0xdb, 0x9a, 0x28, 0xaf, 0x3d, 0x99, 0xd5, 0x80, 0x62, 0x46, 0x12, 0x7a, 0xb5, 0x43, 0x8e,
	0x5c, 0xbb, 0x27, 0xea, 0x62, 0x5f, 0x47, 0x51, 0x44, 0x38, 0xe2, 0x3e, 0x89, 0x58, 0x11, 0xbd,
	0x5b, 0x8a, 0x3a, 0x81, 0x8f, 0x23, 0x5e, 0x04, 0xee, 0x97, 0x02, 0x47, 0x3e, 0x0e, 0x5c, 0xfb,
	0x10, 0xbf, 0x47, 0x43, 0x9f, 0xd0, 0x4b, 0x95, 0x34, 0x76, 0x74, 0xc6, 0x11, 0x4f, 0x0a, 0x48,
	0xb5, 0x0b, 0xd6, 0x7b, 0x98, 0x9b, 0x6e, 0x2f, 0xa5, 0xf2, 0x2a, 0x65, 0x62, 0xe1, 0x0f, 0x09,
	0x66, 0x1c, 0x36, 0xc0, 0xb2, 0x20, 0x6b, 0x47, 0x28, 0xc4, 0x8a, 0x54, 0x97, 0x1a, 0x8b, 0xdd,
	0xca, 0x4f, 0x53, 0xb6, 0x96, 0x44, 0xe4, 0x0d, 0x0a, 0xb1, 0xfa, 0x47, 0x02, 0xb5, 0xd7, 0x09,
	0x47, 0x1c, 0x97, 0x71, 0x98, 0x00, 0xda, 0x04, 0x55, 0x27, 0x61, 0x9c, 0x84, 0x98, 0xda, 0xbe,
	0x5b, 0x86, 0x01, 0xc2, 0xbf, 0xeb, 0xc2, 0x01, 0x00, 0x24, 0xc6, 0x34, 0xd7, 0xab, 0xc8, 0xf5,
	0x4a, 0xa3, 0xda, 0xda, 0xd2, 0xe6, 0x0d, 0x5a, 0x2b, 0x77, 0xdc, 0x13, 0xf5, 0x05, 0xfa, 0x18,
	0x0f, 0x3e, 0x04, 0x2b, 0x31, 0xa2, 0xdc, 0x47, 0x81, 0x7d, 0x84, 0xfc, 0x20, 0xa1, 0x58, 0xa9,
	0xd4, 0xa5, 0xc6, 0x6d, 0xeb, 0x4e, 0xe1, 0xde, 0xce, 0xbd, 0xf0, 0x01, 0x58, 0x1e, 0xa2, 0xc0,
	0x77, 0x11, 0xc7, 0x36, 0x89, 0x82, 0x63, 0xe5, 0x66, 0x96, 0xb6, 0x24, 0x9c, 0x7b, 0x51, 0x70,
	0xac, 0x7e, 0x92, 0xc0, 0xda, 0xd4, 0xc6, 0x70, 0x17, 0x2c, 0x38, 0x14, 0x23, 0x9e, 0x4f, 0xab,
	0xda, 0xd2, 0x67, 0x2a, 0x18, 0x1d, 0xc2, 0x84, 0x84, 0x9d, 0x1b, 0x56, 0x01, 0x00, 0x15, 0xb0,
	0x40, 0x71, 0x48, 0x86, 0x58, 0x91, 0xd3, 0x89, 0xa5, 0x91, 0xfc, 0x77, 0xb7, 0x0a, 0x16, 0x47,
	0xd2, 0xd4, 0x6f, 0x12, 0xd8, 0x98, 0x3a, 0x7c, 0x16, 0x93, 0x88, 0x61, 0xb8, 0x0d, 0xd6, 0x2e,
	0x29, 0xb7, 0x31, 0xa5, 0x84, 0x66, 0xfa, 0xab, 0x2d, 0x28, 0x08, 0xd2, 0xd8, 0xd1, 0x0e, 0xb2,
	0xcb, 0xb0, 0x56, 0x27, 0x67, 0xf2, 0x32, 0x4d, 0x87, 0x6f, 0xc1, 0x2d, 0x8a, 0x59, 0x12, 0x70,
	0xb1, 0x1c, 0x63, 0xfe, 0x72, 0xae, 0xf2, 0xb2, 0x32, 0x08, 0x4b, 0x40, 0xa9, 0x2f, 0x80, 0x32,
	0x2b, 0x29, 0x5d, 0xc5, 0x94, 0x03, 0x9c, 0xbc, 0xbd, 0xd6, 0xef, 0x0a, 0x58, 0x2d, 0xd7, 0x1e,
	0xe4, 0xbd, 0xe1, 0x77, 0x09, 0xac, 0x5c, 0x3a, 0x6c, 0xf8, 0x74, 0x3e, 0xe3, 0xe9, 0x6f, 0xa1,
	0xf6, 0xbf, 0x6b, 0x54, 0x7b, 0xe7, 0xe6, 0x24, 0xf9, 0x8f, 0x3f, 0x7e, 0x7d, 0x96, 0x1f, 0x43,
	0x3d, 0xfd, 0x06, 0x9c, 0x4c, 0x44, 0x3a, 0xe2, 0x11, 0x30, 0xbd, 0xa9, 0xa3, 0xf2, 0x0e, 0xf5,
	0xe6, 0x29, 0xbc, 0x90, 0xc0, 0xea, 0x94, 0xf5, 0xc2, 0xe7, 0xd7, 0x99, 0xbe, 0x78, 0x92, 0xb5,
	0xce, 0x35, 0xab, 0xf3, 0x9b, 0x52, 0xfb, 0xe7, 0xe6, 0x7a, 0xe9, 0x49, 0x3f, 0x1a, 0x3f, 0xb4,
	0x4c, 0xe6, 0x96, 0xda, 0x4a, 0x65, 0x8e, 0x75, 0x9d, 0x94, 0x92, 0x3b, 0xcd, 0xd3, 0x49, 0x95,
	0x46, 0x98, 0x75, 0x32, 0xa4, 0x66, 0x6d, 0xe3, 0xcc, 0x54, 0xc6, 0x6c, 0x0a, 0x2b, 0xf6, 0x99,
	0xe6, 0x90, 0xb0, 0xfb, 0x57, 0x02, 0x9b, 0x0e, 0x09, 0xe7, 0x32, 0xef, 0x2a, 0x53, 0xee, 0x61,
	0x3f, 0xfd, 0xd8, 0xed, 0x4b, 0xef, 0x76, 0x8a, 0x6a, 0x8f, 0x04, 0x28, 0xf2, 0x34, 0x42, 0x3d,
	0xdd, 0xc3, 0x51, 0xf6, 0x29, 0xd4, 0xc7, 0xfd, 0x66, 0xff, 0x2f, 0x3c, 0x13, 0xc6, 0x17, 0xb9,
	0xd2, 0x33, 0xcd, 0xaf, 0x72, 0xbd, 0x97, 0x03, 0x9a, 0x2e, 0xd3, 0x72, 0x33, 0xb5, 0xfa, 0x6d,
	0xad, 0x68, 0xcc, 0xce, 0x44, 0xca, 0xc0, 0x74, 0xd9, 0x60, 0x94, 0x32, 0xe8, 0xb7, 0x07, 0x22,
	0xe5, 0x42, 0xde, 0xcc, 0xfd, 0x86, 0x61, 0xba, 0xcc, 0x30, 0x46, 0x49, 0x86, 0xd1, 0x6f, 0x1b,
	0x86, 0x48, 0x3b, 0x5c, 0xc8, 0x78, 0xb6, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xf4, 0xfc,
	0x5b, 0xbe, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupLabelServiceClient is the client API for AdGroupLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupLabelServiceClient interface {
	// Returns the requested ad group label in full detail.
	GetAdGroupLabel(ctx context.Context, in *GetAdGroupLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupLabel, error)
	// Creates and removes ad group labels.
	// Operation statuses are returned.
	MutateAdGroupLabels(ctx context.Context, in *MutateAdGroupLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupLabelsResponse, error)
}

type adGroupLabelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupLabelServiceClient(cc grpc.ClientConnInterface) AdGroupLabelServiceClient {
	return &adGroupLabelServiceClient{cc}
}

func (c *adGroupLabelServiceClient) GetAdGroupLabel(ctx context.Context, in *GetAdGroupLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupLabel, error) {
	out := new(resources.AdGroupLabel)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.AdGroupLabelService/GetAdGroupLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupLabelServiceClient) MutateAdGroupLabels(ctx context.Context, in *MutateAdGroupLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupLabelsResponse, error) {
	out := new(MutateAdGroupLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.AdGroupLabelService/MutateAdGroupLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupLabelServiceServer is the server API for AdGroupLabelService service.
type AdGroupLabelServiceServer interface {
	// Returns the requested ad group label in full detail.
	GetAdGroupLabel(context.Context, *GetAdGroupLabelRequest) (*resources.AdGroupLabel, error)
	// Creates and removes ad group labels.
	// Operation statuses are returned.
	MutateAdGroupLabels(context.Context, *MutateAdGroupLabelsRequest) (*MutateAdGroupLabelsResponse, error)
}

// UnimplementedAdGroupLabelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupLabelServiceServer struct {
}

func (*UnimplementedAdGroupLabelServiceServer) GetAdGroupLabel(ctx context.Context, req *GetAdGroupLabelRequest) (*resources.AdGroupLabel, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetAdGroupLabel not implemented")
}
func (*UnimplementedAdGroupLabelServiceServer) MutateAdGroupLabels(ctx context.Context, req *MutateAdGroupLabelsRequest) (*MutateAdGroupLabelsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateAdGroupLabels not implemented")
}

func RegisterAdGroupLabelServiceServer(s *grpc.Server, srv AdGroupLabelServiceServer) {
	s.RegisterService(&_AdGroupLabelService_serviceDesc, srv)
}

func _AdGroupLabelService_GetAdGroupLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupLabelServiceServer).GetAdGroupLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.AdGroupLabelService/GetAdGroupLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupLabelServiceServer).GetAdGroupLabel(ctx, req.(*GetAdGroupLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupLabelService_MutateAdGroupLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupLabelServiceServer).MutateAdGroupLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.AdGroupLabelService/MutateAdGroupLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupLabelServiceServer).MutateAdGroupLabels(ctx, req.(*MutateAdGroupLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupLabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.AdGroupLabelService",
	HandlerType: (*AdGroupLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupLabel",
			Handler:    _AdGroupLabelService_GetAdGroupLabel_Handler,
		},
		{
			MethodName: "MutateAdGroupLabels",
			Handler:    _AdGroupLabelService_MutateAdGroupLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/ad_group_label_service.proto",
}
