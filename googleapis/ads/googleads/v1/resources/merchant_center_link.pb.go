// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/merchant_center_link.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// A data sharing connection, proposed or in use,
// between a Google Ads Customer and a Merchant Center account.
type MerchantCenterLink struct {
	// The resource name of the merchant center link.
	// Merchant center link resource names have the form:
	//
	// `customers/{customer_id}/merchantCenterLinks/{merchant_center_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the Merchant Center account.
	// This field is readonly.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the Merchant Center account.
	// This field is readonly.
	MerchantCenterAccountName *wrappers.StringValue `protobuf:"bytes,4,opt,name=merchant_center_account_name,json=merchantCenterAccountName,proto3" json:"merchant_center_account_name,omitempty"`
	// The status of the link.
	Status               enums.MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                    `json:"-"`
	XXX_unrecognized     []byte                                                      `json:"-"`
	XXX_sizecache        int32                                                       `json:"-"`
}

func (m *MerchantCenterLink) Reset()         { *m = MerchantCenterLink{} }
func (m *MerchantCenterLink) String() string { return proto.CompactTextString(m) }
func (*MerchantCenterLink) ProtoMessage()    {}
func (*MerchantCenterLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_merchant_center_link_a6c67adc47332ce3, []int{0}
}
func (m *MerchantCenterLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantCenterLink.Unmarshal(m, b)
}
func (m *MerchantCenterLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantCenterLink.Marshal(b, m, deterministic)
}
func (dst *MerchantCenterLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantCenterLink.Merge(dst, src)
}
func (m *MerchantCenterLink) XXX_Size() int {
	return xxx_messageInfo_MerchantCenterLink.Size(m)
}
func (m *MerchantCenterLink) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantCenterLink.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantCenterLink proto.InternalMessageInfo

func (m *MerchantCenterLink) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *MerchantCenterLink) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MerchantCenterLink) GetMerchantCenterAccountName() *wrappers.StringValue {
	if m != nil {
		return m.MerchantCenterAccountName
	}
	return nil
}

func (m *MerchantCenterLink) GetStatus() enums.MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.MerchantCenterLinkStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*MerchantCenterLink)(nil), "google.ads.googleads.v1.resources.MerchantCenterLink")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/merchant_center_link.proto", fileDescriptor_merchant_center_link_a6c67adc47332ce3)
}

var fileDescriptor_merchant_center_link_a6c67adc47332ce3 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0x6a, 0xdb, 0x30,
	0x18, 0xc7, 0xce, 0x16, 0x98, 0xf7, 0xe7, 0xe0, 0xcb, 0xb2, 0x2c, 0x8c, 0x64, 0x23, 0x10, 0x18,
	0xc8, 0x38, 0x1b, 0x3b, 0x68, 0x83, 0xe1, 0x8c, 0x11, 0x36, 0xb6, 0x12, 0x9c, 0xe2, 0x43, 0x71,
	0x31, 0x8a, 0xad, 0xba, 0x22, 0xb6, 0x64, 0x24, 0x39, 0x7d, 0x81, 0x3e, 0x49, 0x8f, 0xbd, 0xf6,
	0x2d, 0xfa, 0x28, 0x7d, 0x8a, 0x12, 0xc9, 0x32, 0xb4, 0x69, 0xda, 0xdb, 0x97, 0x7c, 0xbf, 0xbf,
	0x9f, 0xe5, 0xfc, 0xc8, 0x19, 0xcb, 0x0b, 0xec, 0xa1, 0x4c, 0x78, 0x7a, 0xdc, 0x4e, 0x1b, 0xdf,
	0xe3, 0x58, 0xb0, 0x9a, 0xa7, 0x58, 0x78, 0x25, 0xe6, 0xe9, 0x29, 0xa2, 0x32, 0x49, 0x31, 0x95,
	0x98, 0x27, 0x05, 0xa1, 0x6b, 0x50, 0x71, 0x26, 0x99, 0x3b, 0xd2, 0x14, 0x80, 0x32, 0x01, 0x5a,
	0x36, 0xd8, 0xf8, 0xa0, 0x65, 0xf7, 0x7f, 0xee, 0x33, 0xc0, 0xb4, 0x2e, 0x1f, 0x16, 0x4f, 0x84,
	0x44, 0xb2, 0x16, 0xda, 0xa3, 0xff, 0xa1, 0x11, 0x50, 0xbf, 0x56, 0xf5, 0x89, 0x77, 0xc6, 0x51,
	0x55, 0x61, 0x6e, 0xf6, 0x03, 0x63, 0x50, 0x11, 0x0f, 0x51, 0xca, 0x24, 0x92, 0x84, 0xd1, 0x66,
	0xfb, 0xf1, 0xca, 0x76, 0xdc, 0xff, 0x8d, 0xc7, 0x2f, 0x65, 0xf1, 0x8f, 0xd0, 0xb5, 0xfb, 0xc9,
	0x79, 0x6d, 0x22, 0x26, 0x14, 0x95, 0xb8, 0x67, 0x0d, 0xad, 0xc9, 0x8b, 0xf0, 0x95, 0xf9, 0xf3,
	0x00, 0x95, 0xd8, 0xfd, 0xec, 0xd8, 0x24, 0xeb, 0x75, 0x86, 0xd6, 0xe4, 0xe5, 0xf4, 0x7d, 0xd3,
	0x0f, 0x98, 0x18, 0xe0, 0x0f, 0x95, 0xdf, 0xbe, 0x46, 0xa8, 0xa8, 0x71, 0x68, 0x93, 0xcc, 0x3d,
	0x76, 0x06, 0xf7, 0xbb, 0xa0, 0x34, 0x65, 0x35, 0x95, 0xda, 0xe0, 0x99, 0x92, 0x19, 0xec, 0xc8,
	0x2c, 0x25, 0x27, 0x34, 0xd7, 0x3a, 0xef, 0xca, 0x3b, 0x49, 0x03, 0xcd, 0x57, 0x59, 0x0a, 0xa7,
	0xab, 0xaf, 0xd2, 0x7b, 0x3e, 0xb4, 0x26, 0x6f, 0xa6, 0x87, 0x60, 0xdf, 0xe9, 0xd5, 0x5d, 0xc1,
	0x6e, 0xe7, 0xa5, 0xa2, 0xff, 0xa6, 0x75, 0xb9, 0x77, 0x19, 0x36, 0x1e, 0xb3, 0x73, 0xdb, 0x19,
	0xa7, 0xac, 0x04, 0x4f, 0x7e, 0xde, 0xd9, 0xdb, 0x5d, 0xad, 0xc5, 0xb6, 0xda, 0xc2, 0x3a, 0xfa,
	0xdb, 0xb0, 0x73, 0x56, 0x20, 0x9a, 0x03, 0xc6, 0x73, 0x2f, 0xc7, 0x54, 0x15, 0x37, 0x2f, 0xa1,
	0x22, 0xe2, 0x91, 0x97, 0xf7, 0xbd, 0x9d, 0x2e, 0xec, 0xce, 0x3c, 0x08, 0x2e, 0xed, 0xd1, 0x5c,
	0x4b, 0x06, 0x99, 0x00, 0x7a, 0xdc, 0x4e, 0x91, 0x0f, 0x42, 0x83, 0xbc, 0x36, 0x98, 0x38, 0xc8,
	0x44, 0xdc, 0x62, 0xe2, 0xc8, 0x8f, 0x5b, 0xcc, 0x8d, 0x3d, 0xd6, 0x0b, 0x08, 0x83, 0x4c, 0x40,
	0xd8, 0xa2, 0x20, 0x8c, 0x7c, 0x08, 0x5b, 0xdc, 0xaa, 0xab, 0xc2, 0x7e, 0xb9, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x9f, 0x1d, 0x01, 0xf6, 0x25, 0x03, 0x00, 0x00,
}
