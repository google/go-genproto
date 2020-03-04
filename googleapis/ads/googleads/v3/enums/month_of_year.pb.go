// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/month_of_year.proto

package enums

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

// Enumerates months of the year, e.g., "January".
type MonthOfYearEnum_MonthOfYear int32

const (
	// Not specified.
	MonthOfYearEnum_UNSPECIFIED MonthOfYearEnum_MonthOfYear = 0
	// The value is unknown in this version.
	MonthOfYearEnum_UNKNOWN MonthOfYearEnum_MonthOfYear = 1
	// January.
	MonthOfYearEnum_JANUARY MonthOfYearEnum_MonthOfYear = 2
	// February.
	MonthOfYearEnum_FEBRUARY MonthOfYearEnum_MonthOfYear = 3
	// March.
	MonthOfYearEnum_MARCH MonthOfYearEnum_MonthOfYear = 4
	// April.
	MonthOfYearEnum_APRIL MonthOfYearEnum_MonthOfYear = 5
	// May.
	MonthOfYearEnum_MAY MonthOfYearEnum_MonthOfYear = 6
	// June.
	MonthOfYearEnum_JUNE MonthOfYearEnum_MonthOfYear = 7
	// July.
	MonthOfYearEnum_JULY MonthOfYearEnum_MonthOfYear = 8
	// August.
	MonthOfYearEnum_AUGUST MonthOfYearEnum_MonthOfYear = 9
	// September.
	MonthOfYearEnum_SEPTEMBER MonthOfYearEnum_MonthOfYear = 10
	// October.
	MonthOfYearEnum_OCTOBER MonthOfYearEnum_MonthOfYear = 11
	// November.
	MonthOfYearEnum_NOVEMBER MonthOfYearEnum_MonthOfYear = 12
	// December.
	MonthOfYearEnum_DECEMBER MonthOfYearEnum_MonthOfYear = 13
)

var MonthOfYearEnum_MonthOfYear_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "JANUARY",
	3:  "FEBRUARY",
	4:  "MARCH",
	5:  "APRIL",
	6:  "MAY",
	7:  "JUNE",
	8:  "JULY",
	9:  "AUGUST",
	10: "SEPTEMBER",
	11: "OCTOBER",
	12: "NOVEMBER",
	13: "DECEMBER",
}

var MonthOfYearEnum_MonthOfYear_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"JANUARY":     2,
	"FEBRUARY":    3,
	"MARCH":       4,
	"APRIL":       5,
	"MAY":         6,
	"JUNE":        7,
	"JULY":        8,
	"AUGUST":      9,
	"SEPTEMBER":   10,
	"OCTOBER":     11,
	"NOVEMBER":    12,
	"DECEMBER":    13,
}

func (x MonthOfYearEnum_MonthOfYear) String() string {
	return proto.EnumName(MonthOfYearEnum_MonthOfYear_name, int32(x))
}

func (MonthOfYearEnum_MonthOfYear) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9fad203756f7601c, []int{0, 0}
}

// Container for enumeration of months of the year, e.g., "January".
type MonthOfYearEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MonthOfYearEnum) Reset()         { *m = MonthOfYearEnum{} }
func (m *MonthOfYearEnum) String() string { return proto.CompactTextString(m) }
func (*MonthOfYearEnum) ProtoMessage()    {}
func (*MonthOfYearEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fad203756f7601c, []int{0}
}

func (m *MonthOfYearEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonthOfYearEnum.Unmarshal(m, b)
}
func (m *MonthOfYearEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonthOfYearEnum.Marshal(b, m, deterministic)
}
func (m *MonthOfYearEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonthOfYearEnum.Merge(m, src)
}
func (m *MonthOfYearEnum) XXX_Size() int {
	return xxx_messageInfo_MonthOfYearEnum.Size(m)
}
func (m *MonthOfYearEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MonthOfYearEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MonthOfYearEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.MonthOfYearEnum_MonthOfYear", MonthOfYearEnum_MonthOfYear_name, MonthOfYearEnum_MonthOfYear_value)
	proto.RegisterType((*MonthOfYearEnum)(nil), "google.ads.googleads.v3.enums.MonthOfYearEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/month_of_year.proto", fileDescriptor_9fad203756f7601c)
}

var fileDescriptor_9fad203756f7601c = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x25, 0xe9, 0x4c, 0x1f, 0xce, 0x8c, 0xc6, 0xf2, 0x12, 0x31, 0x8b, 0x99, 0x0f, 0x70, 0x84,
	0xb2, 0x33, 0x2b, 0x27, 0xe3, 0x29, 0x2d, 0xcd, 0x43, 0x69, 0x13, 0x14, 0x14, 0xa9, 0x0a, 0x24,
	0x0d, 0x95, 0x1a, 0xbb, 0x8a, 0xd3, 0x4a, 0xfc, 0x0e, 0x4b, 0xfe, 0x81, 0x1f, 0x60, 0xc7, 0x47,
	0xb0, 0xe1, 0x2b, 0x90, 0x63, 0x5a, 0x75, 0x03, 0x9b, 0xe8, 0x9c, 0x7b, 0xee, 0xb9, 0x37, 0x3e,
	0x17, 0xbc, 0xae, 0x85, 0xa8, 0x77, 0x95, 0x5d, 0x94, 0xd2, 0xd6, 0x50, 0xa1, 0xa3, 0x63, 0x57,
	0xfc, 0xd0, 0x48, 0xbb, 0x11, 0xbc, 0xfb, 0xbc, 0x16, 0x9b, 0xf5, 0x97, 0xaa, 0x68, 0xf1, 0xbe,
	0x15, 0x9d, 0x40, 0xf7, 0xba, 0x0f, 0x17, 0xa5, 0xc4, 0x67, 0x0b, 0x3e, 0x3a, 0xb8, 0xb7, 0xbc,
	0x7c, 0x75, 0x9a, 0xb8, 0xdf, 0xda, 0x05, 0xe7, 0xa2, 0x2b, 0xba, 0xad, 0xe0, 0x52, 0x9b, 0x1f,
	0x7f, 0x1a, 0xe0, 0xce, 0x57, 0x43, 0xc3, 0x4d, 0x56, 0x15, 0x2d, 0xe3, 0x87, 0xe6, 0xf1, 0xbb,
	0x01, 0xac, 0x8b, 0x1a, 0xba, 0x03, 0x56, 0x12, 0x2c, 0x23, 0xe6, 0xcd, 0x9e, 0x67, 0xec, 0x09,
	0xbe, 0x40, 0x16, 0x18, 0x25, 0xc1, 0xbb, 0x20, 0x7c, 0x1f, 0x40, 0x43, 0x91, 0x39, 0x0d, 0x12,
	0x1a, 0x67, 0xd0, 0x44, 0x37, 0x60, 0xfc, 0xcc, 0xdc, 0xb8, 0x67, 0x03, 0x34, 0x01, 0xd7, 0x3e,
	0x8d, 0xbd, 0xb7, 0xf0, 0x4a, 0x41, 0x1a, 0xc5, 0xb3, 0x05, 0xbc, 0x46, 0x23, 0x30, 0xf0, 0x69,
	0x06, 0x87, 0x68, 0x0c, 0xae, 0xe6, 0x49, 0xc0, 0xe0, 0x48, 0xa3, 0x45, 0x06, 0xc7, 0x08, 0x80,
	0x21, 0x4d, 0xa6, 0xc9, 0x72, 0x05, 0x27, 0xe8, 0x16, 0x4c, 0x96, 0x2c, 0x5a, 0x31, 0xdf, 0x65,
	0x31, 0x04, 0x6a, 0x51, 0xe8, 0xad, 0x42, 0x45, 0x2c, 0xb5, 0x28, 0x08, 0x53, 0x2d, 0xdd, 0x28,
	0xf6, 0xc4, 0x3c, 0xcd, 0x6e, 0xdd, 0x5f, 0x06, 0x78, 0xf8, 0x24, 0x1a, 0xfc, 0xdf, 0x5c, 0x5c,
	0x78, 0xf1, 0xc4, 0x48, 0x65, 0x11, 0x19, 0x1f, 0xdc, 0xbf, 0x96, 0x5a, 0xec, 0x0a, 0x5e, 0x63,
	0xd1, 0xd6, 0x76, 0x5d, 0xf1, 0x3e, 0xa9, 0xd3, 0x35, 0xf6, 0x5b, 0xf9, 0x8f, 0xe3, 0xbc, 0xe9,
	0xbf, 0x5f, 0xcd, 0xc1, 0x94, 0xd2, 0x6f, 0xe6, 0xfd, 0x54, 0x8f, 0xa2, 0xa5, 0xc4, 0x1a, 0x2a,
	0x94, 0x3a, 0x58, 0x45, 0x2c, 0x7f, 0x9c, 0xf4, 0x9c, 0x96, 0x32, 0x3f, 0xeb, 0x79, 0xea, 0xe4,
	0xbd, 0xfe, 0xdb, 0x7c, 0xd0, 0x45, 0x42, 0x68, 0x29, 0x09, 0x39, 0x77, 0x10, 0x92, 0x3a, 0x84,
	0xf4, 0x3d, 0x1f, 0x87, 0xfd, 0x8f, 0x39, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x68, 0x60, 0x63,
	0xea, 0x34, 0x02, 0x00, 0x00,
}
