// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/policy_finding_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible policy finding errors.
type PolicyFindingErrorEnum_PolicyFindingError int32

const (
	// Enum unspecified.
	PolicyFindingErrorEnum_UNSPECIFIED PolicyFindingErrorEnum_PolicyFindingError = 0
	// The received error code is not known in this version.
	PolicyFindingErrorEnum_UNKNOWN PolicyFindingErrorEnum_PolicyFindingError = 1
	// The resource has been disapproved since the policy summary includes
	// policy topics of type PROHIBITED.
	PolicyFindingErrorEnum_POLICY_FINDING PolicyFindingErrorEnum_PolicyFindingError = 2
	// The given policy topic does not exist.
	PolicyFindingErrorEnum_POLICY_TOPIC_NOT_FOUND PolicyFindingErrorEnum_PolicyFindingError = 3
)

var PolicyFindingErrorEnum_PolicyFindingError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "POLICY_FINDING",
	3: "POLICY_TOPIC_NOT_FOUND",
}

var PolicyFindingErrorEnum_PolicyFindingError_value = map[string]int32{
	"UNSPECIFIED":            0,
	"UNKNOWN":                1,
	"POLICY_FINDING":         2,
	"POLICY_TOPIC_NOT_FOUND": 3,
}

func (x PolicyFindingErrorEnum_PolicyFindingError) String() string {
	return proto.EnumName(PolicyFindingErrorEnum_PolicyFindingError_name, int32(x))
}

func (PolicyFindingErrorEnum_PolicyFindingError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f57e9417b2a9ec9c, []int{0, 0}
}

// Container for enum describing possible policy finding errors.
type PolicyFindingErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyFindingErrorEnum) Reset()         { *m = PolicyFindingErrorEnum{} }
func (m *PolicyFindingErrorEnum) String() string { return proto.CompactTextString(m) }
func (*PolicyFindingErrorEnum) ProtoMessage()    {}
func (*PolicyFindingErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_f57e9417b2a9ec9c, []int{0}
}

func (m *PolicyFindingErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyFindingErrorEnum.Unmarshal(m, b)
}
func (m *PolicyFindingErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyFindingErrorEnum.Marshal(b, m, deterministic)
}
func (m *PolicyFindingErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyFindingErrorEnum.Merge(m, src)
}
func (m *PolicyFindingErrorEnum) XXX_Size() int {
	return xxx_messageInfo_PolicyFindingErrorEnum.Size(m)
}
func (m *PolicyFindingErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyFindingErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyFindingErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.PolicyFindingErrorEnum_PolicyFindingError", PolicyFindingErrorEnum_PolicyFindingError_name, PolicyFindingErrorEnum_PolicyFindingError_value)
	proto.RegisterType((*PolicyFindingErrorEnum)(nil), "google.ads.googleads.v3.errors.PolicyFindingErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/policy_finding_error.proto", fileDescriptor_f57e9417b2a9ec9c)
}

var fileDescriptor_f57e9417b2a9ec9c = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xff, 0xeb, 0xe0, 0x2f, 0x64, 0xa0, 0x25, 0x87, 0x09, 0x43, 0x76, 0xe8, 0x07, 0x48,
	0x0f, 0x39, 0x19, 0x4f, 0xdd, 0xda, 0x8e, 0xa0, 0xa4, 0x01, 0xb7, 0x89, 0x52, 0x28, 0xdd, 0x32,
	0x43, 0x60, 0x4b, 0x4a, 0x33, 0x07, 0x82, 0x9f, 0xc6, 0xa3, 0x1f, 0xc5, 0x8f, 0xe2, 0xd1, 0x4f,
	0x20, 0x6d, 0xb6, 0x5d, 0x86, 0x9e, 0xfa, 0xf0, 0xf6, 0xf7, 0x3c, 0x79, 0xde, 0x17, 0x5c, 0x4b,
	0x63, 0xe4, 0x7a, 0x15, 0x96, 0xc2, 0x86, 0x4e, 0x36, 0x6a, 0x87, 0xc3, 0x55, 0x5d, 0x9b, 0xda,
	0x86, 0x95, 0x59, 0xab, 0xe5, 0x6b, 0xf1, 0xac, 0xb4, 0x50, 0x5a, 0x16, 0xed, 0x14, 0x55, 0xb5,
	0xd9, 0x1a, 0x38, 0x74, 0x3c, 0x2a, 0x85, 0x45, 0x47, 0x2b, 0xda, 0x61, 0xe4, 0xac, 0x83, 0xab,
	0x43, 0x74, 0xa5, 0xc2, 0x52, 0x6b, 0xb3, 0x2d, 0xb7, 0xca, 0x68, 0xeb, 0xdc, 0xc1, 0x1b, 0xe8,
	0xf3, 0x36, 0x3b, 0x75, 0xd1, 0x49, 0x63, 0x4a, 0xf4, 0xcb, 0x26, 0x58, 0x00, 0x78, 0xfa, 0x07,
	0x5e, 0x80, 0xde, 0x8c, 0xdd, 0xf3, 0x64, 0x4c, 0x53, 0x9a, 0xc4, 0xfe, 0x3f, 0xd8, 0x03, 0x67,
	0x33, 0x76, 0xcb, 0xb2, 0x07, 0xe6, 0x77, 0x20, 0x04, 0xe7, 0x3c, 0xbb, 0xa3, 0xe3, 0xc7, 0x22,
	0xa5, 0x2c, 0xa6, 0x6c, 0xe2, 0x7b, 0x70, 0x00, 0xfa, 0xfb, 0xd9, 0x34, 0xe3, 0x74, 0x5c, 0xb0,
	0x6c, 0x5a, 0xa4, 0xd9, 0x8c, 0xc5, 0x7e, 0x77, 0xf4, 0xdd, 0x01, 0xc1, 0xd2, 0x6c, 0xd0, 0xdf,
	0x2b, 0x8c, 0x2e, 0x4f, 0x8b, 0xf0, 0xa6, 0x3d, 0xef, 0x3c, 0xc5, 0x7b, 0xab, 0x34, 0xeb, 0x52,
	0x4b, 0x64, 0x6a, 0x19, 0xca, 0x95, 0x6e, 0x77, 0x3b, 0x1c, 0xb2, 0x52, 0xf6, 0xb7, 0xbb, 0xde,
	0xb8, 0xcf, 0xbb, 0xd7, 0x9d, 0x44, 0xd1, 0x87, 0x37, 0x9c, 0xb8, 0xb0, 0x48, 0x58, 0xe4, 0x64,
	0xa3, 0xe6, 0x18, 0xb5, 0x4f, 0xda, 0xcf, 0x03, 0x90, 0x47, 0xc2, 0xe6, 0x47, 0x20, 0x9f, 0xe3,
	0xdc, 0x01, 0x5f, 0x5e, 0xe0, 0xa6, 0x84, 0x44, 0xc2, 0x12, 0x72, 0x44, 0x08, 0x99, 0x63, 0x42,
	0x1c, 0xb4, 0xf8, 0xdf, 0xb6, 0xc3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x42, 0x34, 0x20,
	0xf4, 0x01, 0x00, 0x00,
}
