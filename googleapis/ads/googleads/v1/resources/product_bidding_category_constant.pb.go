// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/product_bidding_category_constant.proto

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

// A Product Bidding Category.
type ProductBiddingCategoryConstant struct {
	// The resource name of the product bidding category.
	// Product bidding category resource names have the form:
	//
	// `productBiddingCategoryConstants/{country_code}~{level}~{id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// ID of the product bidding category.
	//
	// This ID is equivalent to the google_product_category ID as described in
	// this article: https://support.google.com/merchants/answer/6324436.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Two-letter upper-case country code of the product bidding category.
	CountryCode *wrappers.StringValue `protobuf:"bytes,3,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	// Resource name of the parent product bidding category.
	ProductBiddingCategoryConstantParent *wrappers.StringValue `protobuf:"bytes,4,opt,name=product_bidding_category_constant_parent,json=productBiddingCategoryConstantParent,proto3" json:"product_bidding_category_constant_parent,omitempty"`
	// Level of the product bidding category.
	Level enums.ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel `protobuf:"varint,5,opt,name=level,proto3,enum=google.ads.googleads.v1.enums.ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel" json:"level,omitempty"`
	// Status of the product bidding category.
	Status enums.ProductBiddingCategoryStatusEnum_ProductBiddingCategoryStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.ProductBiddingCategoryStatusEnum_ProductBiddingCategoryStatus" json:"status,omitempty"`
	// Language code of the product bidding category.
	LanguageCode *wrappers.StringValue `protobuf:"bytes,7,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Display value of the product bidding category localized according to
	// language_code.
	LocalizedName        *wrappers.StringValue `protobuf:"bytes,8,opt,name=localized_name,json=localizedName,proto3" json:"localized_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ProductBiddingCategoryConstant) Reset()         { *m = ProductBiddingCategoryConstant{} }
func (m *ProductBiddingCategoryConstant) String() string { return proto.CompactTextString(m) }
func (*ProductBiddingCategoryConstant) ProtoMessage()    {}
func (*ProductBiddingCategoryConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_bidding_category_constant_d0a899e5743d0566, []int{0}
}
func (m *ProductBiddingCategoryConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductBiddingCategoryConstant.Unmarshal(m, b)
}
func (m *ProductBiddingCategoryConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductBiddingCategoryConstant.Marshal(b, m, deterministic)
}
func (dst *ProductBiddingCategoryConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductBiddingCategoryConstant.Merge(dst, src)
}
func (m *ProductBiddingCategoryConstant) XXX_Size() int {
	return xxx_messageInfo_ProductBiddingCategoryConstant.Size(m)
}
func (m *ProductBiddingCategoryConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductBiddingCategoryConstant.DiscardUnknown(m)
}

var xxx_messageInfo_ProductBiddingCategoryConstant proto.InternalMessageInfo

func (m *ProductBiddingCategoryConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ProductBiddingCategoryConstant) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ProductBiddingCategoryConstant) GetCountryCode() *wrappers.StringValue {
	if m != nil {
		return m.CountryCode
	}
	return nil
}

func (m *ProductBiddingCategoryConstant) GetProductBiddingCategoryConstantParent() *wrappers.StringValue {
	if m != nil {
		return m.ProductBiddingCategoryConstantParent
	}
	return nil
}

func (m *ProductBiddingCategoryConstant) GetLevel() enums.ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel {
	if m != nil {
		return m.Level
	}
	return enums.ProductBiddingCategoryLevelEnum_UNSPECIFIED
}

func (m *ProductBiddingCategoryConstant) GetStatus() enums.ProductBiddingCategoryStatusEnum_ProductBiddingCategoryStatus {
	if m != nil {
		return m.Status
	}
	return enums.ProductBiddingCategoryStatusEnum_UNSPECIFIED
}

func (m *ProductBiddingCategoryConstant) GetLanguageCode() *wrappers.StringValue {
	if m != nil {
		return m.LanguageCode
	}
	return nil
}

func (m *ProductBiddingCategoryConstant) GetLocalizedName() *wrappers.StringValue {
	if m != nil {
		return m.LocalizedName
	}
	return nil
}

func init() {
	proto.RegisterType((*ProductBiddingCategoryConstant)(nil), "google.ads.googleads.v1.resources.ProductBiddingCategoryConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/product_bidding_category_constant.proto", fileDescriptor_product_bidding_category_constant_d0a899e5743d0566)
}

var fileDescriptor_product_bidding_category_constant_d0a899e5743d0566 = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0xb1, 0xbb, 0x64, 0x9b, 0x9a, 0xf4, 0xe0, 0x93, 0xe9, 0x4a, 0x49, 0xd7, 0x15, 0x02,
	0x03, 0x19, 0x77, 0x63, 0x07, 0xef, 0x30, 0x1c, 0x33, 0x4a, 0xc7, 0x18, 0xc1, 0x85, 0x1c, 0x8a,
	0xc1, 0x28, 0x96, 0x26, 0x0c, 0xb6, 0x64, 0x24, 0x39, 0x63, 0xfb, 0x14, 0xbb, 0xee, 0xba, 0xe3,
	0x3e, 0xca, 0x3e, 0xca, 0x3e, 0xc5, 0xb0, 0x64, 0xf9, 0xb2, 0xd5, 0x0d, 0xbd, 0xbd, 0xe4, 0xbd,
	0xff, 0xff, 0xf7, 0xf2, 0xde, 0x53, 0xc0, 0x35, 0xe5, 0x9c, 0x56, 0x24, 0x40, 0x58, 0x06, 0x26,
	0xec, 0xa2, 0x5d, 0x18, 0x08, 0x22, 0x79, 0x2b, 0x0a, 0x22, 0x83, 0x46, 0x70, 0xdc, 0x16, 0x2a,
	0xdf, 0x96, 0x18, 0x97, 0x8c, 0xe6, 0x05, 0x52, 0x84, 0x72, 0xf1, 0x35, 0x2f, 0x38, 0x93, 0x0a,
	0x31, 0x05, 0x1b, 0xc1, 0x15, 0xf7, 0xce, 0x8c, 0x1e, 0x22, 0x2c, 0xe1, 0x60, 0x05, 0x77, 0x21,
	0x1c, 0xac, 0x8e, 0x57, 0x77, 0xd1, 0x08, 0x6b, 0xeb, 0x11, 0x52, 0x45, 0x76, 0xa4, 0x32, 0x98,
	0xe3, 0xe4, 0x81, 0x1e, 0x52, 0x21, 0xd5, 0xca, 0xde, 0xe4, 0xb4, 0x37, 0xd1, 0x9f, 0xb6, 0xed,
	0xe7, 0xe0, 0x8b, 0x40, 0x4d, 0x43, 0x84, 0xcd, 0x9f, 0x58, 0x48, 0x53, 0x06, 0x88, 0x31, 0xae,
	0x90, 0x2a, 0x39, 0xeb, 0xb3, 0xcf, 0x7f, 0x4c, 0xc0, 0xe9, 0xda, 0x70, 0x56, 0x06, 0x93, 0xf4,
	0x94, 0xa4, 0x1f, 0x89, 0x77, 0x0e, 0xe6, 0xf6, 0x67, 0xe7, 0x0c, 0xd5, 0xc4, 0x77, 0x16, 0xce,
	0xf2, 0x69, 0x3a, 0xb3, 0x5f, 0x7e, 0x42, 0x35, 0xf1, 0x5e, 0x02, 0xb7, 0xc4, 0xbe, 0xbb, 0x70,
	0x96, 0x87, 0x97, 0xcf, 0xfa, 0x99, 0x41, 0xdb, 0x12, 0xbc, 0x66, 0xea, 0xcd, 0xeb, 0x0d, 0xaa,
	0x5a, 0x92, 0xba, 0x25, 0xf6, 0xde, 0x81, 0x59, 0xc1, 0x5b, 0xa6, 0xf4, 0xe0, 0x31, 0xf1, 0x0f,
	0xb4, 0xec, 0xe4, 0x1f, 0xd9, 0x8d, 0x12, 0x25, 0xa3, 0x46, 0x77, 0xd8, 0x2b, 0x12, 0x8e, 0x89,
	0xa7, 0xc0, 0xf2, 0xde, 0x55, 0xe6, 0x0d, 0x12, 0x84, 0x29, 0xff, 0xd1, 0x1e, 0xe6, 0x2f, 0x9a,
	0xd1, 0x11, 0xac, 0xb5, 0x93, 0xd7, 0x80, 0x89, 0xde, 0x9e, 0x3f, 0x59, 0x38, 0xcb, 0xa3, 0xcb,
	0x5b, 0x78, 0xd7, 0x95, 0xe8, 0xf5, 0xc1, 0xff, 0x8f, 0xf5, 0x63, 0xe7, 0xf0, 0x9e, 0xb5, 0xf5,
	0x58, 0x3e, 0x35, 0x20, 0x4f, 0x81, 0xa9, 0xd9, 0xb5, 0x3f, 0xd5, 0xc8, 0xec, 0x41, 0xc8, 0x1b,
	0x6d, 0x31, 0xc2, 0x34, 0x05, 0x69, 0xcf, 0xf2, 0x62, 0x30, 0xaf, 0x10, 0xa3, 0x2d, 0xa2, 0xc4,
	0xec, 0xe7, 0xf1, 0x1e, 0x23, 0x9c, 0x59, 0x89, 0x5e, 0x50, 0x02, 0x8e, 0x2a, 0x5e, 0xa0, 0xaa,
	0xfc, 0x46, 0xb0, 0x39, 0x9a, 0x27, 0x7b, 0x78, 0xcc, 0x07, 0x4d, 0x77, 0x53, 0xab, 0xef, 0x2e,
	0xb8, 0x28, 0x78, 0x0d, 0xef, 0x7d, 0x8c, 0xab, 0xf3, 0xf1, 0x13, 0x5e, 0x77, 0xb0, 0xb5, 0x73,
	0xfb, 0xa1, 0x77, 0xa2, 0xbc, 0x6b, 0x16, 0x72, 0x41, 0x03, 0x4a, 0x98, 0x6e, 0xc5, 0xbe, 0xbf,
	0xa6, 0x94, 0x23, 0x7f, 0x20, 0x6f, 0x87, 0xe8, 0xa7, 0x7b, 0x70, 0x15, 0xc7, 0xbf, 0xdc, 0xb3,
	0x2b, 0x63, 0x19, 0x63, 0x09, 0x4d, 0xd8, 0x45, 0x9b, 0x10, 0xa6, 0xb6, 0xf2, 0xb7, 0xad, 0xc9,
	0x62, 0x2c, 0xb3, 0xa1, 0x26, 0xdb, 0x84, 0xd9, 0x50, 0xf3, 0xc7, 0xbd, 0x30, 0x89, 0x28, 0x8a,
	0xb1, 0x8c, 0xa2, 0xa1, 0x2a, 0x8a, 0x36, 0x61, 0x14, 0x0d, 0x75, 0xdb, 0xa9, 0x6e, 0xf6, 0xd5,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x85, 0x4e, 0x23, 0x27, 0xec, 0x04, 0x00, 0x00,
}
