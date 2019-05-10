// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/field_error.proto

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

// Enum describing possible field errors.
type FieldErrorEnum_FieldError int32

const (
	// Enum unspecified.
	FieldErrorEnum_UNSPECIFIED FieldErrorEnum_FieldError = 0
	// The received error code is not known in this version.
	FieldErrorEnum_UNKNOWN FieldErrorEnum_FieldError = 1
	// The required field was not present.
	FieldErrorEnum_REQUIRED FieldErrorEnum_FieldError = 2
	// The field attempted to be mutated is immutable.
	FieldErrorEnum_IMMUTABLE_FIELD FieldErrorEnum_FieldError = 3
	// The field's value is invalid.
	FieldErrorEnum_INVALID_VALUE FieldErrorEnum_FieldError = 4
	// The field cannot be set.
	FieldErrorEnum_VALUE_MUST_BE_UNSET FieldErrorEnum_FieldError = 5
	// The required repeated field was empty.
	FieldErrorEnum_REQUIRED_NONEMPTY_LIST FieldErrorEnum_FieldError = 6
	// The field cannot be cleared.
	FieldErrorEnum_FIELD_CANNOT_BE_CLEARED FieldErrorEnum_FieldError = 7
	// The field's value is on a blacklist for this field.
	FieldErrorEnum_BLACKLISTED_VALUE FieldErrorEnum_FieldError = 8
)

var FieldErrorEnum_FieldError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "REQUIRED",
	3: "IMMUTABLE_FIELD",
	4: "INVALID_VALUE",
	5: "VALUE_MUST_BE_UNSET",
	6: "REQUIRED_NONEMPTY_LIST",
	7: "FIELD_CANNOT_BE_CLEARED",
	8: "BLACKLISTED_VALUE",
}
var FieldErrorEnum_FieldError_value = map[string]int32{
	"UNSPECIFIED":             0,
	"UNKNOWN":                 1,
	"REQUIRED":                2,
	"IMMUTABLE_FIELD":         3,
	"INVALID_VALUE":           4,
	"VALUE_MUST_BE_UNSET":     5,
	"REQUIRED_NONEMPTY_LIST":  6,
	"FIELD_CANNOT_BE_CLEARED": 7,
	"BLACKLISTED_VALUE":       8,
}

func (x FieldErrorEnum_FieldError) String() string {
	return proto.EnumName(FieldErrorEnum_FieldError_name, int32(x))
}
func (FieldErrorEnum_FieldError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_field_error_2de3d5e282abeaca, []int{0, 0}
}

// Container for enum describing possible field errors.
type FieldErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldErrorEnum) Reset()         { *m = FieldErrorEnum{} }
func (m *FieldErrorEnum) String() string { return proto.CompactTextString(m) }
func (*FieldErrorEnum) ProtoMessage()    {}
func (*FieldErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_field_error_2de3d5e282abeaca, []int{0}
}
func (m *FieldErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldErrorEnum.Unmarshal(m, b)
}
func (m *FieldErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldErrorEnum.Marshal(b, m, deterministic)
}
func (dst *FieldErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldErrorEnum.Merge(dst, src)
}
func (m *FieldErrorEnum) XXX_Size() int {
	return xxx_messageInfo_FieldErrorEnum.Size(m)
}
func (m *FieldErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FieldErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FieldErrorEnum)(nil), "google.ads.googleads.v1.errors.FieldErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.FieldErrorEnum_FieldError", FieldErrorEnum_FieldError_name, FieldErrorEnum_FieldError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/field_error.proto", fileDescriptor_field_error_2de3d5e282abeaca)
}

var fileDescriptor_field_error_2de3d5e282abeaca = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x8e, 0xd3, 0x30,
	0x14, 0xc6, 0x49, 0x06, 0x66, 0x46, 0x1e, 0xa0, 0x19, 0x8f, 0x60, 0xa4, 0x01, 0xcd, 0x22, 0x07,
	0x70, 0x88, 0xd8, 0x99, 0x95, 0x93, 0xb8, 0x23, 0xab, 0x89, 0x1b, 0x9a, 0x3f, 0x08, 0x14, 0xc9,
	0x0a, 0x24, 0x44, 0x91, 0xda, 0xb8, 0x8a, 0x4b, 0x0f, 0xc4, 0x92, 0x9b, 0xc0, 0x19, 0x38, 0x41,
	0x4f, 0x81, 0x12, 0x93, 0x76, 0x05, 0x2b, 0x7f, 0x7e, 0xfa, 0x7d, 0xdf, 0xb3, 0xdf, 0x03, 0x6f,
	0x1a, 0x29, 0x9b, 0x75, 0xed, 0x94, 0x95, 0x72, 0xb4, 0x1c, 0xd4, 0xde, 0x75, 0xea, 0xbe, 0x97,
	0xbd, 0x72, 0xbe, 0xb6, 0xf5, 0xba, 0x12, 0xe3, 0x05, 0x6d, 0x7b, 0xb9, 0x93, 0xf0, 0x5e, 0x63,
	0xa8, 0xac, 0x14, 0x3a, 0x3a, 0xd0, 0xde, 0x45, 0xda, 0x71, 0xf7, 0x7a, 0x4a, 0xdc, 0xb6, 0x4e,
	0xd9, 0x75, 0x72, 0x57, 0xee, 0x5a, 0xd9, 0x29, 0xed, 0xb6, 0x7f, 0x1b, 0xe0, 0xf9, 0x7c, 0xc8,
	0xa4, 0x03, 0x4d, 0xbb, 0x6f, 0x1b, 0xfb, 0xa7, 0x01, 0xc0, 0xa9, 0x04, 0x67, 0xe0, 0x2a, 0xe3,
	0x49, 0x4c, 0x7d, 0x36, 0x67, 0x34, 0xb0, 0x1e, 0xc1, 0x2b, 0x70, 0x91, 0xf1, 0x05, 0x5f, 0x7e,
	0xe0, 0x96, 0x01, 0x9f, 0x82, 0xcb, 0x15, 0x7d, 0x9f, 0xb1, 0x15, 0x0d, 0x2c, 0x13, 0xde, 0x80,
	0x19, 0x8b, 0xa2, 0x2c, 0x25, 0x5e, 0x48, 0xc5, 0x9c, 0xd1, 0x30, 0xb0, 0xce, 0xe0, 0x35, 0x78,
	0xc6, 0x78, 0x4e, 0x42, 0x16, 0x88, 0x9c, 0x84, 0x19, 0xb5, 0x1e, 0xc3, 0x5b, 0x70, 0x33, 0x4a,
	0x11, 0x65, 0x49, 0x2a, 0x3c, 0x2a, 0x32, 0x9e, 0xd0, 0xd4, 0x7a, 0x02, 0xef, 0xc0, 0xcb, 0x29,
	0x4e, 0xf0, 0x25, 0xa7, 0x51, 0x9c, 0x7e, 0x14, 0x21, 0x4b, 0x52, 0xeb, 0x1c, 0xbe, 0x02, 0xb7,
	0x63, 0xa4, 0xf0, 0x09, 0xe7, 0xcb, 0xd1, 0xe6, 0x87, 0x94, 0x0c, 0x9d, 0x2f, 0xe0, 0x0b, 0x70,
	0xed, 0x85, 0xc4, 0x5f, 0x0c, 0x2c, 0x9d, 0x1a, 0x5d, 0x7a, 0x07, 0x03, 0xd8, 0x5f, 0xe4, 0x06,
	0xfd, 0x7f, 0x46, 0xde, 0xec, 0xf4, 0xdf, 0x78, 0x18, 0x4b, 0x6c, 0x7c, 0x0a, 0xfe, 0x5a, 0x1a,
	0xb9, 0x2e, 0xbb, 0x06, 0xc9, 0xbe, 0x71, 0x9a, 0xba, 0x1b, 0x87, 0x36, 0x2d, 0x66, 0xdb, 0xaa,
	0x7f, 0xed, 0xe9, 0x9d, 0x3e, 0xbe, 0x9b, 0x67, 0x0f, 0x84, 0xfc, 0x30, 0xef, 0x1f, 0x74, 0x18,
	0xa9, 0x14, 0xd2, 0x72, 0x50, 0xb9, 0x8b, 0xc6, 0x96, 0xea, 0xd7, 0x04, 0x14, 0xa4, 0x52, 0xc5,
	0x11, 0x28, 0x72, 0xb7, 0xd0, 0xc0, 0xc1, 0xb4, 0x75, 0x15, 0x63, 0x52, 0x29, 0x8c, 0x8f, 0x08,
	0xc6, 0xb9, 0x8b, 0xb1, 0x86, 0x3e, 0x9f, 0x8f, 0xaf, 0x7b, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x9c, 0xd0, 0xb3, 0x96, 0x44, 0x02, 0x00, 0x00,
}
