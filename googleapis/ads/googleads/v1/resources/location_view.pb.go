// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/location_view.proto

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

// A location view summarizes the performance of campaigns by
// Location criteria.
type LocationView struct {
	// Immutable. The resource name of the location view.
	// Location view resource names have the form:
	//
	// `customers/{customer_id}/locationViews/{campaign_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocationView) Reset()         { *m = LocationView{} }
func (m *LocationView) String() string { return proto.CompactTextString(m) }
func (*LocationView) ProtoMessage()    {}
func (*LocationView) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aa3d275dc2bd5c6, []int{0}
}

func (m *LocationView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationView.Unmarshal(m, b)
}
func (m *LocationView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationView.Marshal(b, m, deterministic)
}
func (m *LocationView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationView.Merge(m, src)
}
func (m *LocationView) XXX_Size() int {
	return xxx_messageInfo_LocationView.Size(m)
}
func (m *LocationView) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationView.DiscardUnknown(m)
}

var xxx_messageInfo_LocationView proto.InternalMessageInfo

func (m *LocationView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*LocationView)(nil), "google.ads.googleads.v1.resources.LocationView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/location_view.proto", fileDescriptor_3aa3d275dc2bd5c6)
}

var fileDescriptor_3aa3d275dc2bd5c6 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xf3, 0x40,
	0x1c, 0xc6, 0x49, 0x5f, 0x5e, 0xc1, 0x50, 0x07, 0x3b, 0x69, 0x11, 0xb4, 0x42, 0xd1, 0xc5, 0x3b,
	0xa2, 0xb8, 0x9c, 0xd3, 0x75, 0x29, 0x88, 0x48, 0xe9, 0x90, 0x41, 0x83, 0xe5, 0x9a, 0x9c, 0xf1,
	0x20, 0xc9, 0xbf, 0xdc, 0xa5, 0xe9, 0x50, 0xfa, 0x65, 0x1c, 0xfd, 0x18, 0x8e, 0x7e, 0x0a, 0xe7,
	0x7e, 0x04, 0x07, 0x91, 0xe4, 0x72, 0xd7, 0xb8, 0xa8, 0xdb, 0x03, 0xff, 0xdf, 0xf3, 0xdc, 0xc3,
	0x73, 0xee, 0x65, 0x0c, 0x10, 0x27, 0x1c, 0xb3, 0x48, 0x61, 0x2d, 0x4b, 0x55, 0x78, 0x58, 0x72,
	0x05, 0x73, 0x19, 0x72, 0x85, 0x13, 0x08, 0x59, 0x2e, 0x20, 0x9b, 0x14, 0x82, 0x2f, 0xd0, 0x4c,
	0x42, 0x0e, 0x9d, 0x9e, 0x66, 0x11, 0x8b, 0x14, 0xb2, 0x36, 0x54, 0x78, 0xc8, 0xda, 0xba, 0x87,
	0x26, 0x79, 0x26, 0xf0, 0xa3, 0xe0, 0x49, 0x34, 0x99, 0xf2, 0x27, 0x56, 0x08, 0x90, 0x3a, 0xa3,
	0xbb, 0xdf, 0x00, 0x8c, 0xad, 0x3e, 0x1d, 0x34, 0x4e, 0x2c, 0xcb, 0x20, 0xaf, 0x0a, 0x28, 0x7d,
	0x3d, 0x7e, 0x75, 0xdc, 0xf6, 0x4d, 0x5d, 0xca, 0x17, 0x7c, 0xd1, 0x19, 0xbb, 0x3b, 0x26, 0x60,
	0x92, 0xb1, 0x94, 0xef, 0x39, 0x47, 0xce, 0xe9, 0xf6, 0xe0, 0xec, 0x9d, 0xfe, 0xff, 0xa0, 0x27,
	0x6e, 0x7f, 0xd3, 0xb0, 0x56, 0x33, 0xa1, 0x50, 0x08, 0x29, 0x6e, 0xa6, 0x8c, 0xdb, 0x26, 0xe3,
	0x96, 0xa5, 0x9c, 0x3c, 0xac, 0xe9, 0xfd, 0x1f, 0x9d, 0x9d, 0xf3, 0x70, 0xae, 0x72, 0x48, 0xb9,
	0x54, 0x78, 0x69, 0xe4, 0xca, 0xee, 0x56, 0x22, 0x0a, 0x2f, 0xbf, 0xcd, 0xb8, 0x1a, 0x7c, 0x3a,
	0x6e, 0x3f, 0x84, 0x14, 0xfd, 0x3a, 0xe4, 0x60, 0xb7, 0xf9, 0xd6, 0xa8, 0x5c, 0x60, 0xe4, 0xdc,
	0x5d, 0xd7, 0xbe, 0x18, 0x12, 0x96, 0xc5, 0x08, 0x64, 0x8c, 0x63, 0x9e, 0x55, 0xfb, 0xe0, 0x4d,
	0xd5, 0x1f, 0xbe, 0xf5, 0xca, 0xaa, 0xe7, 0xd6, 0xbf, 0x21, 0xa5, 0x2f, 0xad, 0xde, 0x50, 0x47,
	0xd2, 0x48, 0x21, 0x2d, 0x4b, 0xe5, 0x7b, 0x68, 0x6c, 0xc8, 0x37, 0xc3, 0x04, 0x34, 0x52, 0x81,
	0x65, 0x02, 0xdf, 0x0b, 0x2c, 0xb3, 0x6e, 0xf5, 0xf5, 0x81, 0x10, 0x1a, 0x29, 0x42, 0x2c, 0x45,
	0x88, 0xef, 0x11, 0x62, 0xb9, 0xe9, 0x56, 0x55, 0xf6, 0xe2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x69,
	0x65, 0xf7, 0x33, 0x82, 0x02, 0x00, 0x00,
}
