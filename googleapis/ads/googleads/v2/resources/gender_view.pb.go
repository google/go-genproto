// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/gender_view.proto

package resources

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

// A gender view.
type GenderView struct {
	// Immutable. The resource name of the gender view.
	// Gender view resource names have the form:
	//
	// `customers/{customer_id}/genderViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenderView) Reset()         { *m = GenderView{} }
func (m *GenderView) String() string { return proto.CompactTextString(m) }
func (*GenderView) ProtoMessage()    {}
func (*GenderView) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c741fcdff34173c, []int{0}
}

func (m *GenderView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenderView.Unmarshal(m, b)
}
func (m *GenderView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenderView.Marshal(b, m, deterministic)
}
func (m *GenderView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenderView.Merge(m, src)
}
func (m *GenderView) XXX_Size() int {
	return xxx_messageInfo_GenderView.Size(m)
}
func (m *GenderView) XXX_DiscardUnknown() {
	xxx_messageInfo_GenderView.DiscardUnknown(m)
}

var xxx_messageInfo_GenderView proto.InternalMessageInfo

func (m *GenderView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GenderView)(nil), "google.ads.googleads.v2.resources.GenderView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/gender_view.proto", fileDescriptor_1c741fcdff34173c)
}

var fileDescriptor_1c741fcdff34173c = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xf3, 0x30,
	0x1c, 0xc6, 0xe9, 0x5e, 0x5e, 0xc1, 0xa2, 0x08, 0x3b, 0xe9, 0x10, 0x74, 0xca, 0x40, 0x10, 0x12,
	0xe8, 0x6e, 0xf1, 0x94, 0x5d, 0x06, 0x1e, 0x64, 0x0c, 0x2c, 0x22, 0x85, 0x91, 0x35, 0xb1, 0x06,
	0xd6, 0xfc, 0x47, 0xd2, 0x75, 0x87, 0xb1, 0x2f, 0xe3, 0xd1, 0x4f, 0xe0, 0x67, 0xf0, 0x53, 0x78,
	0xde, 0x47, 0x10, 0x0f, 0xd2, 0xa5, 0x49, 0x77, 0x52, 0x6f, 0x0f, 0xfc, 0x7f, 0xcf, 0x93, 0xa7,
	0x4f, 0xc3, 0x7e, 0x06, 0x90, 0xcd, 0x04, 0x66, 0xdc, 0x60, 0x2b, 0x2b, 0x55, 0x46, 0x58, 0x0b,
	0x03, 0x0b, 0x9d, 0x0a, 0x83, 0x33, 0xa1, 0xb8, 0xd0, 0x93, 0x52, 0x8a, 0x25, 0x9a, 0x6b, 0x28,
	0xa0, 0xdd, 0xb5, 0x24, 0x62, 0xdc, 0x20, 0x6f, 0x42, 0x65, 0x84, 0xbc, 0xa9, 0x73, 0xe6, 0x72,
	0xe7, 0x12, 0x3f, 0x49, 0x31, 0xe3, 0x93, 0xa9, 0x78, 0x66, 0xa5, 0x04, 0x6d, 0x33, 0x3a, 0x27,
	0x3b, 0x80, 0xb3, 0xd5, 0xa7, 0xd3, 0x9d, 0x13, 0x53, 0x0a, 0x0a, 0x56, 0x48, 0x50, 0xc6, 0x5e,
	0x2f, 0xde, 0x82, 0x30, 0x1c, 0x6e, 0x2b, 0xc5, 0x52, 0x2c, 0xdb, 0xa3, 0xf0, 0xd0, 0xd9, 0x27,
	0x8a, 0xe5, 0xe2, 0x38, 0x38, 0x0f, 0xae, 0xf6, 0x07, 0xd7, 0x1f, 0xf4, 0xff, 0x27, 0xed, 0x85,
	0x97, 0x4d, 0xbf, 0x5a, 0xcd, 0xa5, 0x41, 0x29, 0xe4, 0xb8, 0xc9, 0x18, 0x1f, 0xb8, 0x84, 0x3b,
	0x96, 0x0b, 0xf2, 0xb0, 0xa1, 0xf7, 0x7f, 0xf2, 0xb5, 0x51, 0xba, 0x30, 0x05, 0xe4, 0x42, 0x1b,
	0xbc, 0x72, 0x72, 0x5d, 0xef, 0x55, 0x01, 0x06, 0xaf, 0x76, 0xc6, 0x5b, 0x0f, 0xbe, 0x82, 0xb0,
	0x97, 0x42, 0x8e, 0x7e, 0x9d, 0x6f, 0x70, 0xd4, 0xbc, 0x32, 0xaa, 0xbe, 0x7a, 0x14, 0x3c, 0xde,
	0xd6, 0xae, 0x0c, 0x66, 0x4c, 0x65, 0x08, 0x74, 0x56, 0xbd, 0xb4, 0xdd, 0x04, 0x37, 0x15, 0x7f,
	0xf8, 0x91, 0x37, 0x5e, 0xbd, 0xb4, 0xfe, 0x0d, 0x29, 0x7d, 0x6d, 0x75, 0x87, 0x36, 0x92, 0x72,
	0x83, 0xac, 0xac, 0x54, 0x1c, 0xa1, 0xb1, 0x23, 0xdf, 0x1d, 0x93, 0x50, 0x6e, 0x12, 0xcf, 0x24,
	0x71, 0x94, 0x78, 0x66, 0xd3, 0xea, 0xd9, 0x03, 0x21, 0x94, 0x1b, 0x42, 0x3c, 0x45, 0x48, 0x1c,
	0x11, 0xe2, 0xb9, 0xe9, 0xde, 0xb6, 0x6c, 0xff, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x27, 0xe8, 0x9c,
	0x41, 0x74, 0x02, 0x00, 0x00,
}
