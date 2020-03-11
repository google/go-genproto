// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/campaign_shared_set.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// CampaignSharedSets are used for managing the shared sets associated with a
// campaign.
type CampaignSharedSet struct {
	// Immutable. The resource name of the campaign shared set.
	// Campaign shared set resource names have the form:
	//
	// `customers/{customer_id}/campaignSharedSets/{campaign_id}~{shared_set_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The campaign to which the campaign shared set belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,3,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Immutable. The shared set associated with the campaign. This may be a negative keyword
	// shared set of another customer. This customer should be a manager of the
	// other customer, otherwise the campaign shared set will exist but have no
	// serving effect. Only negative keyword shared sets can be associated with
	// Shopping campaigns. Only negative placement shared sets can be associated
	// with Display mobile app campaigns.
	SharedSet *wrappers.StringValue `protobuf:"bytes,4,opt,name=shared_set,json=sharedSet,proto3" json:"shared_set,omitempty"`
	// Output only. The status of this campaign shared set. Read only.
	Status               enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus `protobuf:"varint,2,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                  `json:"-"`
	XXX_unrecognized     []byte                                                    `json:"-"`
	XXX_sizecache        int32                                                     `json:"-"`
}

func (m *CampaignSharedSet) Reset()         { *m = CampaignSharedSet{} }
func (m *CampaignSharedSet) String() string { return proto.CompactTextString(m) }
func (*CampaignSharedSet) ProtoMessage()    {}
func (*CampaignSharedSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c77b63cdc30c4b2, []int{0}
}

func (m *CampaignSharedSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignSharedSet.Unmarshal(m, b)
}
func (m *CampaignSharedSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignSharedSet.Marshal(b, m, deterministic)
}
func (m *CampaignSharedSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignSharedSet.Merge(m, src)
}
func (m *CampaignSharedSet) XXX_Size() int {
	return xxx_messageInfo_CampaignSharedSet.Size(m)
}
func (m *CampaignSharedSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignSharedSet.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignSharedSet proto.InternalMessageInfo

func (m *CampaignSharedSet) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignSharedSet) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignSharedSet) GetSharedSet() *wrappers.StringValue {
	if m != nil {
		return m.SharedSet
	}
	return nil
}

func (m *CampaignSharedSet) GetStatus() enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus {
	if m != nil {
		return m.Status
	}
	return enums.CampaignSharedSetStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CampaignSharedSet)(nil), "google.ads.googleads.v1.resources.CampaignSharedSet")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/campaign_shared_set.proto", fileDescriptor_5c77b63cdc30c4b2)
}

var fileDescriptor_5c77b63cdc30c4b2 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6b, 0x13, 0x41,
	0x18, 0x66, 0x37, 0xb5, 0xd8, 0xf1, 0x03, 0xdc, 0x83, 0xc4, 0x52, 0xb4, 0xa9, 0x16, 0x62, 0x91,
	0x19, 0x36, 0xde, 0xa6, 0x28, 0xcc, 0x8a, 0x14, 0x3c, 0x48, 0xd9, 0x40, 0x04, 0x09, 0x2c, 0x93,
	0xdd, 0xe9, 0x74, 0x31, 0x3b, 0xb3, 0xcc, 0xcc, 0xc6, 0x83, 0xf4, 0xe4, 0xcf, 0xf0, 0xe6, 0xd1,
	0x9f, 0xe2, 0xaf, 0xe8, 0xb9, 0x3f, 0xc1, 0x93, 0x64, 0x77, 0x66, 0x52, 0x88, 0xb1, 0x7a, 0x7b,
	0xc2, 0xfb, 0x7c, 0xbc, 0x79, 0xf6, 0x1d, 0x70, 0xcc, 0xa5, 0xe4, 0x73, 0x86, 0x68, 0xa1, 0x51,
	0x07, 0x97, 0x68, 0x11, 0x23, 0xc5, 0xb4, 0x6c, 0x54, 0xce, 0x34, 0xca, 0x69, 0x55, 0xd3, 0x92,
	0x8b, 0x4c, 0x9f, 0x53, 0xc5, 0x8a, 0x4c, 0x33, 0x03, 0x6b, 0x25, 0x8d, 0x8c, 0x06, 0x9d, 0x02,
	0xd2, 0x42, 0x43, 0x2f, 0x86, 0x8b, 0x18, 0x7a, 0xf1, 0xee, 0xeb, 0x4d, 0xfe, 0x4c, 0x34, 0xd5,
	0x1f, 0xbd, 0x33, 0x6d, 0xa8, 0x69, 0x74, 0x17, 0xb1, 0xfb, 0xc4, 0xe9, 0xeb, 0x12, 0x9d, 0x95,
	0x6c, 0x5e, 0x64, 0x33, 0x76, 0x4e, 0x17, 0xa5, 0x54, 0x96, 0xf0, 0xe8, 0x1a, 0xc1, 0xc5, 0xda,
	0xd1, 0x63, 0x3b, 0x6a, 0x7f, 0xcd, 0x9a, 0x33, 0xf4, 0x59, 0xd1, 0xba, 0x66, 0xca, 0x79, 0xef,
	0x5d, 0x93, 0x52, 0x21, 0xa4, 0xa1, 0xa6, 0x94, 0xc2, 0x4e, 0x0f, 0xbe, 0x6d, 0x81, 0x07, 0x6f,
	0xec, 0x7a, 0xe3, 0x76, 0xbb, 0x31, 0x33, 0xd1, 0x07, 0x70, 0xcf, 0xa5, 0x64, 0x82, 0x56, 0xac,
	0x1f, 0xec, 0x07, 0xc3, 0x9d, 0x64, 0x74, 0x49, 0x6e, 0xfd, 0x22, 0x2f, 0xc0, 0xd1, 0xaa, 0x06,
	0x8b, 0xea, 0x52, 0xc3, 0x5c, 0x56, 0x68, 0xcd, 0x2a, 0xbd, 0xeb, 0x8c, 0xde, 0xd3, 0x8a, 0x45,
	0x39, 0xb8, 0xed, 0xca, 0xe8, 0xf7, 0xf6, 0x83, 0xe1, 0x9d, 0xd1, 0x9e, 0xb5, 0x80, 0x6e, 0x7f,
	0x38, 0x36, 0xaa, 0x14, 0x7c, 0x42, 0xe7, 0x0d, 0x4b, 0x9e, 0xb7, 0x89, 0x4f, 0xc1, 0xe0, 0xc6,
	0xc4, 0xd4, 0x1b, 0x47, 0x1c, 0x80, 0x55, 0xd1, 0xfd, 0xad, 0x7f, 0x88, 0x39, 0x6a, 0x63, 0x9e,
	0x81, 0x83, 0x8d, 0x31, 0xab, 0x3f, 0xb4, 0xa3, 0x7d, 0x4d, 0x12, 0x6c, 0x77, 0x9f, 0xb1, 0x1f,
	0xee, 0x07, 0xc3, 0xfb, 0xa3, 0x14, 0x6e, 0x3a, 0x95, 0xf6, 0x0e, 0xe0, 0x5a, 0x3b, 0xe3, 0x56,
	0xfd, 0x56, 0x34, 0xd5, 0xa6, 0x59, 0xd2, 0xbb, 0x24, 0xbd, 0xd4, 0xc6, 0x60, 0x71, 0x45, 0x3e,
	0xfd, 0x4f, 0xfb, 0xd1, 0xab, 0xbc, 0xd1, 0x46, 0x56, 0x4c, 0x69, 0xf4, 0xc5, 0xc1, 0x0b, 0x7f,
	0x8f, 0x9e, 0xb7, 0x9c, 0xae, 0xdf, 0xe8, 0x45, 0xf2, 0x35, 0x04, 0x87, 0xb9, 0xac, 0xe0, 0x8d,
	0x2f, 0x20, 0x79, 0xb8, 0x96, 0x7d, 0xba, 0x2c, 0xfa, 0x34, 0xf8, 0xf8, 0xce, 0x8a, 0xb9, 0x9c,
	0x53, 0xc1, 0xa1, 0x54, 0x1c, 0x71, 0x26, 0xda, 0xcf, 0x80, 0x56, 0xfb, 0xff, 0xe5, 0x69, 0x1e,
	0x7b, 0xf4, 0x3d, 0xec, 0x9d, 0x10, 0xf2, 0x23, 0x1c, 0x9c, 0x74, 0x96, 0xa4, 0xd0, 0xb0, 0x83,
	0x4b, 0x34, 0x89, 0x61, 0xea, 0x98, 0x3f, 0x1d, 0x67, 0x4a, 0x0a, 0x3d, 0xf5, 0x9c, 0xe9, 0x24,
	0x9e, 0x7a, 0xce, 0x55, 0x78, 0xd8, 0x0d, 0x30, 0x26, 0x85, 0xc6, 0xd8, 0xb3, 0x30, 0x9e, 0xc4,
	0x18, 0x7b, 0xde, 0x6c, 0xbb, 0x5d, 0xf6, 0xe5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0xec,
	0x9c, 0x67, 0x46, 0x04, 0x00, 0x00,
}
