// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/region_code_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible region code errors.
type RegionCodeErrorEnum_RegionCodeError int32

const (
	// Enum unspecified.
	RegionCodeErrorEnum_UNSPECIFIED RegionCodeErrorEnum_RegionCodeError = 0
	// The received error code is not known in this version.
	RegionCodeErrorEnum_UNKNOWN RegionCodeErrorEnum_RegionCodeError = 1
	// Invalid region code.
	RegionCodeErrorEnum_INVALID_REGION_CODE RegionCodeErrorEnum_RegionCodeError = 2
)

var RegionCodeErrorEnum_RegionCodeError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "INVALID_REGION_CODE",
}
var RegionCodeErrorEnum_RegionCodeError_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"INVALID_REGION_CODE": 2,
}

func (x RegionCodeErrorEnum_RegionCodeError) String() string {
	return proto.EnumName(RegionCodeErrorEnum_RegionCodeError_name, int32(x))
}
func (RegionCodeErrorEnum_RegionCodeError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_region_code_error_3ef96732f9283c65, []int{0, 0}
}

// Container for enum describing possible region code errors.
type RegionCodeErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegionCodeErrorEnum) Reset()         { *m = RegionCodeErrorEnum{} }
func (m *RegionCodeErrorEnum) String() string { return proto.CompactTextString(m) }
func (*RegionCodeErrorEnum) ProtoMessage()    {}
func (*RegionCodeErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_region_code_error_3ef96732f9283c65, []int{0}
}
func (m *RegionCodeErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegionCodeErrorEnum.Unmarshal(m, b)
}
func (m *RegionCodeErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegionCodeErrorEnum.Marshal(b, m, deterministic)
}
func (dst *RegionCodeErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegionCodeErrorEnum.Merge(dst, src)
}
func (m *RegionCodeErrorEnum) XXX_Size() int {
	return xxx_messageInfo_RegionCodeErrorEnum.Size(m)
}
func (m *RegionCodeErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_RegionCodeErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_RegionCodeErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegionCodeErrorEnum)(nil), "google.ads.googleads.v1.errors.RegionCodeErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.RegionCodeErrorEnum_RegionCodeError", RegionCodeErrorEnum_RegionCodeError_name, RegionCodeErrorEnum_RegionCodeError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/region_code_error.proto", fileDescriptor_region_code_error_3ef96732f9283c65)
}

var fileDescriptor_region_code_error_3ef96732f9283c65 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xf3, 0x30,
	0x1c, 0xc5, 0xbf, 0xf5, 0x03, 0x85, 0xec, 0x62, 0xa3, 0x13, 0x04, 0x91, 0x5d, 0xf4, 0x01, 0x12,
	0x8a, 0xe0, 0x45, 0xbc, 0xca, 0xd6, 0x38, 0x8b, 0x92, 0x8d, 0xc9, 0x2a, 0x48, 0xa1, 0xd4, 0x25,
	0x84, 0xc2, 0x96, 0xff, 0x48, 0xe6, 0x1e, 0xc8, 0x4b, 0x1f, 0xc5, 0x47, 0xf1, 0xc6, 0x57, 0x90,
	0x36, 0xb6, 0x17, 0x03, 0xbd, 0xca, 0xe1, 0xf0, 0x3b, 0x27, 0x87, 0x3f, 0xba, 0xd6, 0x00, 0x7a,
	0xa3, 0x48, 0x29, 0x1d, 0xf1, 0xb2, 0x56, 0x87, 0x98, 0x28, 0x6b, 0xc1, 0x3a, 0x62, 0x95, 0xae,
	0xc0, 0x14, 0x6b, 0x90, 0xaa, 0x68, 0x2c, 0xbc, 0xb3, 0xb0, 0x87, 0x70, 0xec, 0x61, 0x5c, 0x4a,
	0x87, 0xbb, 0x1c, 0x3e, 0xc4, 0xd8, 0xe7, 0x2e, 0x2e, 0xdb, 0xde, 0x5d, 0x45, 0x4a, 0x63, 0x60,
	0x5f, 0xee, 0x2b, 0x30, 0xce, 0xa7, 0xa3, 0x02, 0x8d, 0x96, 0x4d, 0xf1, 0x14, 0xa4, 0xe2, 0x75,
	0x82, 0x9b, 0xd7, 0x6d, 0x74, 0x87, 0x06, 0x47, 0x76, 0x38, 0x40, 0xfd, 0x95, 0x78, 0x5c, 0xf0,
	0x69, 0x7a, 0x9b, 0xf2, 0x64, 0xf8, 0x2f, 0xec, 0xa3, 0xd3, 0x95, 0xb8, 0x17, 0xf3, 0x27, 0x31,
	0xec, 0x85, 0xe7, 0x68, 0x94, 0x8a, 0x8c, 0x3d, 0xa4, 0x49, 0xb1, 0xe4, 0xb3, 0x74, 0x2e, 0x8a,
	0xe9, 0x3c, 0xe1, 0xc3, 0x60, 0xf2, 0xd5, 0x43, 0xd1, 0x1a, 0xb6, 0xf8, 0xef, 0x95, 0x93, 0xb3,
	0xa3, 0xef, 0x16, 0xf5, 0xba, 0x45, 0xef, 0x39, 0xf9, 0xc9, 0x69, 0xd8, 0x94, 0x46, 0x63, 0xb0,
	0x9a, 0x68, 0x65, 0x9a, 0xed, 0xed, 0x95, 0x76, 0x95, 0xfb, 0xed, 0x68, 0x37, 0xfe, 0x79, 0x0b,
	0xfe, 0xcf, 0x18, 0x7b, 0x0f, 0xc6, 0x33, 0x5f, 0xc6, 0xa4, 0xc3, 0x5e, 0xd6, 0x2a, 0x8b, 0x71,
	0xf3, 0xa5, 0xfb, 0x68, 0x81, 0x9c, 0x49, 0x97, 0x77, 0x40, 0x9e, 0xc5, 0xb9, 0x07, 0x3e, 0x83,
	0xc8, 0xbb, 0x94, 0x32, 0xe9, 0x28, 0xed, 0x10, 0x4a, 0xb3, 0x98, 0x52, 0x0f, 0xbd, 0x9c, 0x34,
	0xeb, 0xae, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x2b, 0x1e, 0xbc, 0xd1, 0x01, 0x00, 0x00,
}
