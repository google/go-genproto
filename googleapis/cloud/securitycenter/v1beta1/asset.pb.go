// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1beta1/asset.proto

package securitycenter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Cloud Security Command Center's (Cloud SCC) representation of a Google Cloud
// Platform (GCP) resource.
//
// The Asset is a Cloud SCC resource that captures information about a single GCP
// resource. All modifications to an Asset are only within the context of Cloud
// SCC and don't affect the referenced GCP resource.
type Asset struct {
	// The relative resource name of this asset. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/123/assets/456".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Cloud SCC managed properties. These properties are managed by
	// Cloud SCC and cannot be modified by the user.
	SecurityCenterProperties *Asset_SecurityCenterProperties `protobuf:"bytes,2,opt,name=security_center_properties,json=securityCenterProperties,proto3" json:"security_center_properties,omitempty"`
	// Resource managed properties. These properties are managed and defined by
	// the GCP resource and cannot be modified by the user.
	ResourceProperties map[string]*_struct.Value `protobuf:"bytes,7,rep,name=resource_properties,json=resourceProperties,proto3" json:"resource_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// User specified security marks. These marks are entirely managed by the user
	// and come from the SecurityMarks resource that belongs to the asset.
	SecurityMarks *SecurityMarks `protobuf:"bytes,8,opt,name=security_marks,json=securityMarks,proto3" json:"security_marks,omitempty"`
	// The time at which the asset was created in Cloud SCC.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,9,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The time at which the asset was last updated, added, or deleted in Cloud
	// SCC.
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,10,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_51f3b31ca795196d, []int{0}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Asset) GetSecurityCenterProperties() *Asset_SecurityCenterProperties {
	if m != nil {
		return m.SecurityCenterProperties
	}
	return nil
}

func (m *Asset) GetResourceProperties() map[string]*_struct.Value {
	if m != nil {
		return m.ResourceProperties
	}
	return nil
}

func (m *Asset) GetSecurityMarks() *SecurityMarks {
	if m != nil {
		return m.SecurityMarks
	}
	return nil
}

func (m *Asset) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Asset) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

// Cloud SCC managed properties. These properties are managed by Cloud SCC and
// cannot be modified by the user.
type Asset_SecurityCenterProperties struct {
	// The full resource name of the GCP resource this asset
	// represents. This field is immutable after create time. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The type of the GCP resource. Examples include: APPLICATION,
	// PROJECT, and ORGANIZATION. This is a case insensitive field defined by
	// Cloud SCC and/or the producer of the resource and is immutable
	// after create time.
	ResourceType string `protobuf:"bytes,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// The full resource name of the immediate parent of the resource. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	ResourceParent string `protobuf:"bytes,3,opt,name=resource_parent,json=resourceParent,proto3" json:"resource_parent,omitempty"`
	// The full resource name of the project the resource belongs to. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	ResourceProject string `protobuf:"bytes,4,opt,name=resource_project,json=resourceProject,proto3" json:"resource_project,omitempty"`
	// Owners of the Google Cloud resource.
	ResourceOwners       []string `protobuf:"bytes,5,rep,name=resource_owners,json=resourceOwners,proto3" json:"resource_owners,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Asset_SecurityCenterProperties) Reset()         { *m = Asset_SecurityCenterProperties{} }
func (m *Asset_SecurityCenterProperties) String() string { return proto.CompactTextString(m) }
func (*Asset_SecurityCenterProperties) ProtoMessage()    {}
func (*Asset_SecurityCenterProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_51f3b31ca795196d, []int{0, 0}
}

func (m *Asset_SecurityCenterProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset_SecurityCenterProperties.Unmarshal(m, b)
}
func (m *Asset_SecurityCenterProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset_SecurityCenterProperties.Marshal(b, m, deterministic)
}
func (m *Asset_SecurityCenterProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset_SecurityCenterProperties.Merge(m, src)
}
func (m *Asset_SecurityCenterProperties) XXX_Size() int {
	return xxx_messageInfo_Asset_SecurityCenterProperties.Size(m)
}
func (m *Asset_SecurityCenterProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset_SecurityCenterProperties.DiscardUnknown(m)
}

var xxx_messageInfo_Asset_SecurityCenterProperties proto.InternalMessageInfo

func (m *Asset_SecurityCenterProperties) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceParent() string {
	if m != nil {
		return m.ResourceParent
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceProject() string {
	if m != nil {
		return m.ResourceProject
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceOwners() []string {
	if m != nil {
		return m.ResourceOwners
	}
	return nil
}

func init() {
	proto.RegisterType((*Asset)(nil), "google.cloud.securitycenter.v1beta1.Asset")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "google.cloud.securitycenter.v1beta1.Asset.ResourcePropertiesEntry")
	proto.RegisterType((*Asset_SecurityCenterProperties)(nil), "google.cloud.securitycenter.v1beta1.Asset.SecurityCenterProperties")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1beta1/asset.proto", fileDescriptor_51f3b31ca795196d)
}

var fileDescriptor_51f3b31ca795196d = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0x25, 0xdd, 0x6e, 0xb5, 0x77, 0xed, 0x07, 0x23, 0xe8, 0x10, 0x04, 0x17, 0xfb, 0xd0, 0x15,
	0x64, 0x42, 0xd7, 0x97, 0x62, 0x9f, 0x6c, 0xf1, 0x51, 0xad, 0xb1, 0x08, 0x0a, 0xb2, 0xcc, 0xa6,
	0xd7, 0x10, 0xbb, 0x99, 0x19, 0x66, 0x26, 0x95, 0xbc, 0x08, 0xfe, 0x1f, 0xff, 0x94, 0xff, 0x44,
	0x66, 0x32, 0x49, 0xb3, 0x96, 0xd5, 0xf5, 0x2d, 0x39, 0xf7, 0x9c, 0x73, 0x73, 0xef, 0xb9, 0xbb,
	0x90, 0xe4, 0x52, 0xe6, 0x0b, 0x4c, 0xb2, 0x85, 0xac, 0x2e, 0x13, 0x83, 0x59, 0xa5, 0x0b, 0x5b,
	0x67, 0x28, 0x2c, 0xea, 0xe4, 0xfa, 0x68, 0x8e, 0x96, 0x1f, 0x25, 0xdc, 0x18, 0xb4, 0x4c, 0x69,
	0x69, 0x25, 0x39, 0x68, 0x04, 0xcc, 0x0b, 0xd8, 0xb2, 0x80, 0x05, 0x41, 0xfc, 0x28, 0xb8, 0x72,
	0x55, 0x24, 0x5c, 0x08, 0x69, 0xb9, 0x2d, 0xa4, 0x30, 0x8d, 0x45, 0x7c, 0xbc, 0x4e, 0xcf, 0x16,
	0x9e, 0x95, 0x5c, 0x5f, 0xb5, 0xca, 0xd6, 0xd7, 0xbf, 0xcd, 0xab, 0x2f, 0x89, 0xb1, 0xba, 0xca,
	0xc2, 0xa7, 0xc5, 0x8f, 0xff, 0xac, 0xda, 0xa2, 0x44, 0x63, 0x79, 0xa9, 0x1a, 0xc2, 0x93, 0x9f,
	0x5b, 0x30, 0x7c, 0xe9, 0x66, 0x21, 0x04, 0x36, 0x05, 0x2f, 0x91, 0x46, 0xe3, 0x68, 0xb2, 0x9d,
	0xfa, 0x67, 0xf2, 0x23, 0x82, 0xb8, 0xeb, 0xda, 0x7c, 0xcd, 0x4c, 0x69, 0xa9, 0x50, 0xdb, 0x02,
	0x0d, 0xdd, 0x18, 0x47, 0x93, 0xd1, 0xf4, 0x8c, 0xad, 0x31, 0x3f, 0xf3, 0x4d, 0xd8, 0xfb, 0x50,
	0x3c, 0xf3, 0xc5, 0xf3, 0xce, 0x2a, 0xa5, 0x66, 0x45, 0x85, 0x18, 0xb8, 0xaf, 0xd1, 0xc8, 0x4a,
	0x67, 0xd8, 0xef, 0x7d, 0x67, 0x3c, 0x98, 0x8c, 0xa6, 0xa7, 0xff, 0xd1, 0x3b, 0x0d, 0x2e, 0x37,
	0xde, 0xaf, 0x84, 0xd5, 0x75, 0x4a, 0xf4, 0xad, 0x02, 0xf9, 0x08, 0xbb, 0xcb, 0xdb, 0xa6, 0x77,
	0xfd, 0xac, 0xd3, 0xb5, 0xfa, 0xb5, 0x53, 0xbe, 0x76, 0xca, 0x74, 0xc7, 0xf4, 0x5f, 0xc9, 0x09,
	0x8c, 0x32, 0x8d, 0xdc, 0xe2, 0xcc, 0x65, 0x41, 0xb7, 0xbd, 0x6f, 0xdc, 0xfa, 0xb6, 0x41, 0xb1,
	0x8b, 0x36, 0xa8, 0x14, 0x1a, 0xba, 0x03, 0x9c, 0xb8, 0x52, 0x97, 0x9d, 0x18, 0xfe, 0x2d, 0x6e,
	0xe8, 0x0e, 0x88, 0x7f, 0x45, 0x40, 0x57, 0x05, 0x40, 0x0e, 0x60, 0xa7, 0x5b, 0x73, 0xef, 0x0e,
	0xee, 0xb5, 0xe0, 0x1b, 0x77, 0x0f, 0x7d, 0x92, 0xad, 0x15, 0xfa, 0x0b, 0xe8, 0x91, 0x2e, 0x6a,
	0x85, 0xe4, 0x10, 0xf6, 0x6e, 0x02, 0xe3, 0x1a, 0x85, 0xa5, 0x03, 0x4f, 0xdb, 0xed, 0x16, 0xed,
	0x51, 0xf2, 0x14, 0xf6, 0xfb, 0xc9, 0x7e, 0xc5, 0xcc, 0xd2, 0x4d, 0xcf, 0xdc, 0xeb, 0x45, 0xe2,
	0xe0, 0x25, 0x4f, 0xf9, 0x4d, 0xa0, 0x36, 0x74, 0x38, 0x1e, 0xf4, 0x3d, 0xdf, 0x7a, 0x34, 0xfe,
	0x0c, 0x0f, 0x57, 0xe4, 0x4c, 0xf6, 0x61, 0x70, 0x85, 0x75, 0x98, 0xcb, 0x3d, 0x92, 0x67, 0x30,
	0xbc, 0xe6, 0x8b, 0x0a, 0xc3, 0x21, 0x3f, 0xb8, 0xb5, 0xc7, 0x0f, 0xae, 0x9a, 0x36, 0xa4, 0x17,
	0x1b, 0xc7, 0xd1, 0xe9, 0x77, 0x38, 0xcc, 0x64, 0xb9, 0xce, 0x11, 0x9c, 0x47, 0x9f, 0xde, 0x05,
	0x5a, 0x2e, 0x17, 0x5c, 0xe4, 0x4c, 0xea, 0x3c, 0xc9, 0x51, 0x78, 0xf3, 0xf0, 0x1f, 0xc3, 0x55,
	0x61, 0xfe, 0xfa, 0x9b, 0x3f, 0x59, 0x86, 0xe7, 0x5b, 0x5e, 0xfd, 0xfc, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd3, 0x05, 0x13, 0x94, 0xa4, 0x04, 0x00, 0x00,
}
