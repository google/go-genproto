// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/type/timeofday.proto
// DO NOT EDIT!

package google_type

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Represents a time of day. The date and time zone are either not significant
// or are specified elsewhere. An API may chose to allow leap seconds. Related
// types are [google.type.Date][google.type.Date] and `google.protobuf.Timestamp`.
type TimeOfDay struct {
	// Hours of day in 24 hour format. Should be from 0 to 23. An API may choose
	// to allow the value "24:00:00" for scenarios like business closing time.
	Hours int32 `protobuf:"varint,1,opt,name=hours" json:"hours,omitempty"`
	// Minutes of hour of day. Must be from 0 to 59.
	Minutes int32 `protobuf:"varint,2,opt,name=minutes" json:"minutes,omitempty"`
	// Seconds of minutes of the time. Must normally be from 0 to 59. An API may
	// allow the value 60 if it allows leap-seconds.
	Seconds int32 `protobuf:"varint,3,opt,name=seconds" json:"seconds,omitempty"`
	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	Nanos int32 `protobuf:"varint,4,opt,name=nanos" json:"nanos,omitempty"`
}

func (m *TimeOfDay) Reset()                    { *m = TimeOfDay{} }
func (m *TimeOfDay) String() string            { return proto.CompactTextString(m) }
func (*TimeOfDay) ProtoMessage()               {}
func (*TimeOfDay) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func init() {
	proto.RegisterType((*TimeOfDay)(nil), "google.type.TimeOfDay")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/type/timeofday.proto", fileDescriptor5)
}

var fileDescriptor5 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xb2, 0x4a, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4b, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0xcb, 0x2f, 0x4a, 0xd7, 0x4f, 0x4f,
	0xcd, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x87, 0x48, 0x25, 0x16, 0x64, 0x16, 0xeb, 0x97, 0x54,
	0x16, 0xa4, 0xea, 0x97, 0x64, 0xe6, 0xa6, 0xe6, 0xa7, 0xa5, 0x24, 0x56, 0xea, 0x81, 0xe5, 0x85,
	0xb8, 0xa1, 0x7a, 0x41, 0x92, 0x4a, 0xd9, 0x5c, 0x9c, 0x21, 0x40, 0x79, 0xff, 0x34, 0x97, 0xc4,
	0x4a, 0x21, 0x11, 0x2e, 0xd6, 0x8c, 0xfc, 0xd2, 0xa2, 0x62, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6,
	0x20, 0x08, 0x47, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0x33, 0xaf, 0xb4, 0x24, 0xb5, 0x58, 0x82, 0x09,
	0x2c, 0x0e, 0xe3, 0x82, 0x64, 0x8a, 0x53, 0x93, 0xf3, 0xf3, 0x52, 0x8a, 0x25, 0x98, 0x21, 0x32,
	0x50, 0x2e, 0xc8, 0xa4, 0xbc, 0xc4, 0xbc, 0xfc, 0x62, 0x09, 0x16, 0x88, 0x49, 0x60, 0x8e, 0x93,
	0x26, 0x17, 0x7f, 0x72, 0x7e, 0xae, 0x1e, 0x92, 0xfd, 0x4e, 0x7c, 0x70, 0xdb, 0x03, 0x40, 0x8e,
	0x0b, 0x60, 0x5c, 0xc4, 0xc4, 0xec, 0x1e, 0x12, 0x90, 0xc4, 0x06, 0x76, 0xab, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xe1, 0x0e, 0x56, 0xbd, 0xe9, 0x00, 0x00, 0x00,
}
