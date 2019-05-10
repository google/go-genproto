// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/extension_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

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

// Possible data types for an extension in an extension setting.
type ExtensionTypeEnum_ExtensionType int32

const (
	// Not specified.
	ExtensionTypeEnum_UNSPECIFIED ExtensionTypeEnum_ExtensionType = 0
	// Used for return value only. Represents value unknown in this version.
	ExtensionTypeEnum_UNKNOWN ExtensionTypeEnum_ExtensionType = 1
	// None.
	ExtensionTypeEnum_NONE ExtensionTypeEnum_ExtensionType = 2
	// App.
	ExtensionTypeEnum_APP ExtensionTypeEnum_ExtensionType = 3
	// Call.
	ExtensionTypeEnum_CALL ExtensionTypeEnum_ExtensionType = 4
	// Callout.
	ExtensionTypeEnum_CALLOUT ExtensionTypeEnum_ExtensionType = 5
	// Message.
	ExtensionTypeEnum_MESSAGE ExtensionTypeEnum_ExtensionType = 6
	// Price.
	ExtensionTypeEnum_PRICE ExtensionTypeEnum_ExtensionType = 7
	// Promotion.
	ExtensionTypeEnum_PROMOTION ExtensionTypeEnum_ExtensionType = 8
	// Review.
	ExtensionTypeEnum_REVIEW ExtensionTypeEnum_ExtensionType = 9
	// Sitelink.
	ExtensionTypeEnum_SITELINK ExtensionTypeEnum_ExtensionType = 10
	// Structured snippet.
	ExtensionTypeEnum_STRUCTURED_SNIPPET ExtensionTypeEnum_ExtensionType = 11
	// Location.
	ExtensionTypeEnum_LOCATION ExtensionTypeEnum_ExtensionType = 12
	// Affiliate location.
	ExtensionTypeEnum_AFFILIATE_LOCATION ExtensionTypeEnum_ExtensionType = 13
)

var ExtensionTypeEnum_ExtensionType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "NONE",
	3:  "APP",
	4:  "CALL",
	5:  "CALLOUT",
	6:  "MESSAGE",
	7:  "PRICE",
	8:  "PROMOTION",
	9:  "REVIEW",
	10: "SITELINK",
	11: "STRUCTURED_SNIPPET",
	12: "LOCATION",
	13: "AFFILIATE_LOCATION",
}
var ExtensionTypeEnum_ExtensionType_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"NONE":               2,
	"APP":                3,
	"CALL":               4,
	"CALLOUT":            5,
	"MESSAGE":            6,
	"PRICE":              7,
	"PROMOTION":          8,
	"REVIEW":             9,
	"SITELINK":           10,
	"STRUCTURED_SNIPPET": 11,
	"LOCATION":           12,
	"AFFILIATE_LOCATION": 13,
}

func (x ExtensionTypeEnum_ExtensionType) String() string {
	return proto.EnumName(ExtensionTypeEnum_ExtensionType_name, int32(x))
}
func (ExtensionTypeEnum_ExtensionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_extension_type_040fb0a485e702bf, []int{0, 0}
}

// Container for enum describing possible data types for an extension in an
// extension setting.
type ExtensionTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtensionTypeEnum) Reset()         { *m = ExtensionTypeEnum{} }
func (m *ExtensionTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ExtensionTypeEnum) ProtoMessage()    {}
func (*ExtensionTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_extension_type_040fb0a485e702bf, []int{0}
}
func (m *ExtensionTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionTypeEnum.Unmarshal(m, b)
}
func (m *ExtensionTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionTypeEnum.Marshal(b, m, deterministic)
}
func (dst *ExtensionTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionTypeEnum.Merge(dst, src)
}
func (m *ExtensionTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ExtensionTypeEnum.Size(m)
}
func (m *ExtensionTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExtensionTypeEnum)(nil), "google.ads.googleads.v1.enums.ExtensionTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.ExtensionTypeEnum_ExtensionType", ExtensionTypeEnum_ExtensionType_name, ExtensionTypeEnum_ExtensionType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/extension_type.proto", fileDescriptor_extension_type_040fb0a485e702bf)
}

var fileDescriptor_extension_type_040fb0a485e702bf = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4d, 0x6a, 0xdc, 0x30,
	0x18, 0xad, 0x67, 0x92, 0xf9, 0xd1, 0x64, 0xa8, 0xaa, 0x45, 0x17, 0xa5, 0x59, 0x24, 0x07, 0x90,
	0x71, 0xbb, 0x53, 0x57, 0x1a, 0x47, 0x33, 0x88, 0x38, 0xb2, 0xb0, 0xe5, 0x09, 0x94, 0x81, 0xc1,
	0xad, 0x8d, 0x19, 0xc8, 0x48, 0x26, 0x72, 0x42, 0x73, 0x9d, 0x2e, 0x7b, 0x94, 0x1e, 0xa0, 0x57,
	0x28, 0x94, 0x1e, 0xa2, 0xc8, 0xae, 0x0d, 0x59, 0xb4, 0x1b, 0xf1, 0xf4, 0xbe, 0xf7, 0x1e, 0xd2,
	0xfb, 0xc0, 0xbb, 0xca, 0x98, 0xea, 0xae, 0xf4, 0xf3, 0xc2, 0xfa, 0x1d, 0x74, 0xe8, 0x31, 0xf0,
	0x4b, 0xfd, 0x70, 0xb4, 0x7e, 0xf9, 0xa5, 0x29, 0xb5, 0x3d, 0x18, 0xbd, 0x6f, 0x9e, 0xea, 0x12,
	0xd7, 0xf7, 0xa6, 0x31, 0xe8, 0xbc, 0x13, 0xe2, 0xbc, 0xb0, 0x78, 0xf0, 0xe0, 0xc7, 0x00, 0xb7,
	0x9e, 0x37, 0x6f, 0xfb, 0xc8, 0xfa, 0xe0, 0xe7, 0x5a, 0x9b, 0x26, 0x6f, 0x0e, 0x46, 0xdb, 0xce,
	0x7c, 0xf9, 0xdb, 0x03, 0xaf, 0x58, 0x9f, 0xaa, 0x9e, 0xea, 0x92, 0xe9, 0x87, 0xe3, 0xe5, 0x0f,
	0x0f, 0x2c, 0x9f, 0xb1, 0xe8, 0x25, 0x58, 0x64, 0x22, 0x95, 0x2c, 0xe4, 0x6b, 0xce, 0xae, 0xe0,
	0x0b, 0xb4, 0x00, 0xd3, 0x4c, 0x5c, 0x8b, 0xf8, 0x56, 0x40, 0x0f, 0xcd, 0xc0, 0x89, 0x88, 0x05,
	0x83, 0x23, 0x34, 0x05, 0x63, 0x2a, 0x25, 0x1c, 0x3b, 0x2a, 0xa4, 0x51, 0x04, 0x4f, 0x9c, 0xd2,
	0xa1, 0x38, 0x53, 0xf0, 0xd4, 0x5d, 0x6e, 0x58, 0x9a, 0xd2, 0x0d, 0x83, 0x13, 0x34, 0x07, 0xa7,
	0x32, 0xe1, 0x21, 0x83, 0x53, 0xb4, 0x04, 0x73, 0x99, 0xc4, 0x37, 0xb1, 0xe2, 0xb1, 0x80, 0x33,
	0x04, 0xc0, 0x24, 0x61, 0x5b, 0xce, 0x6e, 0xe1, 0x1c, 0x9d, 0x81, 0x59, 0xca, 0x15, 0x8b, 0xb8,
	0xb8, 0x86, 0x00, 0xbd, 0x06, 0x28, 0x55, 0x49, 0x16, 0xaa, 0x2c, 0x61, 0x57, 0xfb, 0x54, 0x70,
	0x29, 0x99, 0x82, 0x0b, 0xa7, 0x8a, 0xe2, 0x90, 0xb6, 0xfe, 0x33, 0xa7, 0xa2, 0xeb, 0x35, 0x8f,
	0x38, 0x55, 0x6c, 0x3f, 0xf0, 0xcb, 0xd5, 0x4f, 0x0f, 0x5c, 0x7c, 0x36, 0x47, 0xfc, 0xdf, 0xca,
	0x56, 0xe8, 0xd9, 0xdf, 0xa5, 0x2b, 0x4a, 0x7a, 0x1f, 0x57, 0x7f, 0x4d, 0x95, 0xb9, 0xcb, 0x75,
	0x85, 0xcd, 0x7d, 0xe5, 0x57, 0xa5, 0x6e, 0x6b, 0xec, 0x77, 0x55, 0x1f, 0xec, 0x3f, 0x56, 0xf7,
	0xa1, 0x3d, 0xbf, 0x8e, 0xc6, 0x1b, 0x4a, 0xbf, 0x8d, 0xce, 0x37, 0x5d, 0x14, 0x2d, 0x2c, 0xee,
	0xa0, 0x43, 0xdb, 0x00, 0xbb, 0xf6, 0xed, 0xf7, 0x7e, 0xbe, 0xa3, 0x85, 0xdd, 0x0d, 0xf3, 0xdd,
	0x36, 0xd8, 0xb5, 0xf3, 0x5f, 0xa3, 0x8b, 0x8e, 0x24, 0x84, 0x16, 0x96, 0x90, 0x41, 0x41, 0xc8,
	0x36, 0x20, 0xa4, 0xd5, 0x7c, 0x9a, 0xb4, 0x0f, 0x7b, 0xff, 0x27, 0x00, 0x00, 0xff, 0xff, 0xb6,
	0x23, 0x53, 0xc3, 0x52, 0x02, 0x00, 0x00,
}
