// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/asset.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
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

// Asset is a part of an ad which can be shared across multiple ads.
// It can be an image (ImageAsset), a video (YoutubeVideoAsset), etc.
type Asset struct {
	// The resource name of the asset.
	// Asset resource names have the form:
	//
	// `customers/{customer_id}/assets/{asset_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the asset.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Optional name of the asset.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Type of the asset.
	Type enums.AssetTypeEnum_AssetType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.AssetTypeEnum_AssetType" json:"type,omitempty"`
	// The specific type of the asset.
	//
	// Types that are valid to be assigned to AssetData:
	//	*Asset_YoutubeVideoAsset
	//	*Asset_MediaBundleAsset
	//	*Asset_ImageAsset
	//	*Asset_TextAsset
	AssetData            isAsset_AssetData `protobuf_oneof:"asset_data"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_asset_7039888b9b05b05a, []int{0}
}
func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (dst *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(dst, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Asset) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Asset) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Asset) GetType() enums.AssetTypeEnum_AssetType {
	if m != nil {
		return m.Type
	}
	return enums.AssetTypeEnum_UNSPECIFIED
}

type isAsset_AssetData interface {
	isAsset_AssetData()
}

type Asset_YoutubeVideoAsset struct {
	YoutubeVideoAsset *common.YoutubeVideoAsset `protobuf:"bytes,5,opt,name=youtube_video_asset,json=youtubeVideoAsset,proto3,oneof"`
}

type Asset_MediaBundleAsset struct {
	MediaBundleAsset *common.MediaBundleAsset `protobuf:"bytes,6,opt,name=media_bundle_asset,json=mediaBundleAsset,proto3,oneof"`
}

type Asset_ImageAsset struct {
	ImageAsset *common.ImageAsset `protobuf:"bytes,7,opt,name=image_asset,json=imageAsset,proto3,oneof"`
}

type Asset_TextAsset struct {
	TextAsset *common.TextAsset `protobuf:"bytes,8,opt,name=text_asset,json=textAsset,proto3,oneof"`
}

func (*Asset_YoutubeVideoAsset) isAsset_AssetData() {}

func (*Asset_MediaBundleAsset) isAsset_AssetData() {}

func (*Asset_ImageAsset) isAsset_AssetData() {}

func (*Asset_TextAsset) isAsset_AssetData() {}

func (m *Asset) GetAssetData() isAsset_AssetData {
	if m != nil {
		return m.AssetData
	}
	return nil
}

func (m *Asset) GetYoutubeVideoAsset() *common.YoutubeVideoAsset {
	if x, ok := m.GetAssetData().(*Asset_YoutubeVideoAsset); ok {
		return x.YoutubeVideoAsset
	}
	return nil
}

func (m *Asset) GetMediaBundleAsset() *common.MediaBundleAsset {
	if x, ok := m.GetAssetData().(*Asset_MediaBundleAsset); ok {
		return x.MediaBundleAsset
	}
	return nil
}

func (m *Asset) GetImageAsset() *common.ImageAsset {
	if x, ok := m.GetAssetData().(*Asset_ImageAsset); ok {
		return x.ImageAsset
	}
	return nil
}

func (m *Asset) GetTextAsset() *common.TextAsset {
	if x, ok := m.GetAssetData().(*Asset_TextAsset); ok {
		return x.TextAsset
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Asset) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Asset_OneofMarshaler, _Asset_OneofUnmarshaler, _Asset_OneofSizer, []interface{}{
		(*Asset_YoutubeVideoAsset)(nil),
		(*Asset_MediaBundleAsset)(nil),
		(*Asset_ImageAsset)(nil),
		(*Asset_TextAsset)(nil),
	}
}

func _Asset_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Asset)
	// asset_data
	switch x := m.AssetData.(type) {
	case *Asset_YoutubeVideoAsset:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.YoutubeVideoAsset); err != nil {
			return err
		}
	case *Asset_MediaBundleAsset:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MediaBundleAsset); err != nil {
			return err
		}
	case *Asset_ImageAsset:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ImageAsset); err != nil {
			return err
		}
	case *Asset_TextAsset:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TextAsset); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Asset.AssetData has unexpected type %T", x)
	}
	return nil
}

func _Asset_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Asset)
	switch tag {
	case 5: // asset_data.youtube_video_asset
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.YoutubeVideoAsset)
		err := b.DecodeMessage(msg)
		m.AssetData = &Asset_YoutubeVideoAsset{msg}
		return true, err
	case 6: // asset_data.media_bundle_asset
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.MediaBundleAsset)
		err := b.DecodeMessage(msg)
		m.AssetData = &Asset_MediaBundleAsset{msg}
		return true, err
	case 7: // asset_data.image_asset
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ImageAsset)
		err := b.DecodeMessage(msg)
		m.AssetData = &Asset_ImageAsset{msg}
		return true, err
	case 8: // asset_data.text_asset
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.TextAsset)
		err := b.DecodeMessage(msg)
		m.AssetData = &Asset_TextAsset{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Asset_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Asset)
	// asset_data
	switch x := m.AssetData.(type) {
	case *Asset_YoutubeVideoAsset:
		s := proto.Size(x.YoutubeVideoAsset)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Asset_MediaBundleAsset:
		s := proto.Size(x.MediaBundleAsset)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Asset_ImageAsset:
		s := proto.Size(x.ImageAsset)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Asset_TextAsset:
		s := proto.Size(x.TextAsset)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Asset)(nil), "google.ads.googleads.v1.resources.Asset")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/asset.proto", fileDescriptor_asset_7039888b9b05b05a)
}

var fileDescriptor_asset_7039888b9b05b05a = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x49, 0xd6, 0x0d, 0xe6, 0x0d, 0x04, 0xe6, 0x52, 0x8d, 0x09, 0x75, 0xa0, 0x49, 0x05,
	0x84, 0xd3, 0x0c, 0xb4, 0x43, 0x38, 0xa5, 0x12, 0x1a, 0xab, 0x34, 0x34, 0x85, 0x29, 0x12, 0xa8,
	0x52, 0x70, 0x6b, 0x13, 0x59, 0xaa, 0xed, 0x28, 0x76, 0xca, 0xfa, 0x3a, 0x1c, 0x79, 0x11, 0x24,
	0x1e, 0x85, 0x17, 0xe0, 0x8a, 0x62, 0xc7, 0x01, 0x15, 0x75, 0xb9, 0x7d, 0x76, 0x7e, 0xff, 0xdf,
	0xf7, 0xd5, 0xb5, 0xc1, 0xcb, 0x5c, 0xca, 0x7c, 0x41, 0x03, 0x4c, 0x54, 0x60, 0xcb, 0xba, 0x5a,
	0x86, 0x41, 0x49, 0x95, 0xac, 0xca, 0x39, 0x55, 0x01, 0x56, 0x8a, 0x6a, 0x54, 0x94, 0x52, 0x4b,
	0x78, 0x64, 0x19, 0x84, 0x89, 0x42, 0x2d, 0x8e, 0x96, 0x21, 0x6a, 0xf1, 0x83, 0xd1, 0x26, 0xe3,
	0x5c, 0x72, 0x2e, 0x85, 0xd5, 0x65, 0x7a, 0x55, 0x50, 0x65, 0xa5, 0x07, 0x68, 0x53, 0x82, 0x8a,
	0x8a, 0xab, 0x7f, 0x02, 0x0d, 0xff, 0xb8, 0xe1, 0xcd, 0x6a, 0x56, 0x7d, 0x09, 0xbe, 0x96, 0xb8,
	0x28, 0x68, 0xe9, 0x7c, 0x87, 0xce, 0x57, 0xb0, 0x00, 0x0b, 0x21, 0x35, 0xd6, 0x4c, 0x8a, 0xe6,
	0xeb, 0x93, 0x1f, 0x3d, 0xb0, 0x1d, 0xd7, 0x4a, 0xf8, 0x14, 0xdc, 0x75, 0x63, 0x67, 0x02, 0x73,
	0xda, 0xf7, 0x06, 0xde, 0x70, 0x37, 0xd9, 0x77, 0x9b, 0xef, 0x31, 0xa7, 0xf0, 0x05, 0xf0, 0x19,
	0xe9, 0xfb, 0x03, 0x6f, 0xb8, 0x77, 0xf2, 0xa8, 0x99, 0x14, 0xb9, 0xce, 0xe8, 0x5c, 0xe8, 0xd3,
	0xd7, 0x29, 0x5e, 0x54, 0x34, 0xf1, 0x19, 0x81, 0x23, 0xd0, 0x33, 0xa2, 0x2d, 0x83, 0x1f, 0xfe,
	0x87, 0x7f, 0xd0, 0x25, 0x13, 0xb9, 0xe5, 0x0d, 0x09, 0x27, 0xa0, 0x57, 0xff, 0xb2, 0x7e, 0x6f,
	0xe0, 0x0d, 0xef, 0x9d, 0x9c, 0xa2, 0x4d, 0xe7, 0x6b, 0x8e, 0x02, 0x99, 0xb9, 0xaf, 0x56, 0x05,
	0x7d, 0x2b, 0x2a, 0xfe, 0x77, 0x95, 0x18, 0x07, 0x9c, 0x83, 0x87, 0x2b, 0x59, 0xe9, 0x6a, 0x46,
	0xb3, 0x25, 0x23, 0x54, 0x66, 0xe6, 0xe4, 0xfa, 0xdb, 0x66, 0x98, 0x70, 0xa3, 0xda, 0xfe, 0x2f,
	0xe8, 0xa3, 0x8d, 0xa6, 0x75, 0xd2, 0x98, 0xdf, 0xdd, 0x4a, 0x1e, 0xac, 0xd6, 0x37, 0xe1, 0x67,
	0x00, 0x39, 0x25, 0x0c, 0x67, 0xb3, 0x4a, 0x90, 0x05, 0x6d, 0x7a, 0xec, 0x98, 0x1e, 0xa3, 0xae,
	0x1e, 0x17, 0x75, 0x72, 0x6c, 0x82, 0xae, 0xc5, 0x7d, 0xbe, 0xb6, 0x07, 0x2f, 0xc0, 0x1e, 0xe3,
	0x38, 0x77, 0xea, 0xdb, 0x46, 0xfd, 0xbc, 0x4b, 0x7d, 0x5e, 0x47, 0x9c, 0x14, 0xb0, 0x76, 0x05,
	0x27, 0x00, 0x68, 0x7a, 0xad, 0x1b, 0xdb, 0x1d, 0x63, 0x7b, 0xd6, 0x65, 0xbb, 0xa2, 0xd7, 0xda,
	0xc9, 0x76, 0xb5, 0x5b, 0x8c, 0xf7, 0x01, 0xb0, 0xb7, 0x91, 0x60, 0x8d, 0xc7, 0xbf, 0x3d, 0x70,
	0x3c, 0x97, 0x1c, 0x75, 0xbe, 0x89, 0x31, 0x30, 0xf1, 0xcb, 0xfa, 0x1a, 0x5c, 0x7a, 0x9f, 0x26,
	0x4d, 0x20, 0x97, 0x0b, 0x2c, 0x72, 0x24, 0xcb, 0x3c, 0xc8, 0xa9, 0x30, 0x97, 0xc4, 0xdd, 0xff,
	0x82, 0xa9, 0x1b, 0x9e, 0xe4, 0x9b, 0xb6, 0xfa, 0xe6, 0x6f, 0x9d, 0xc5, 0xf1, 0x77, 0xff, 0xe8,
	0xcc, 0x2a, 0x63, 0xa2, 0x90, 0x2d, 0xeb, 0x2a, 0x0d, 0x51, 0xe2, 0xc8, 0x9f, 0x8e, 0x99, 0xc6,
	0x44, 0x4d, 0x5b, 0x66, 0x9a, 0x86, 0xd3, 0x96, 0xf9, 0xe5, 0x1f, 0xdb, 0x0f, 0x51, 0x14, 0x13,
	0x15, 0x45, 0x2d, 0x15, 0x45, 0x69, 0x18, 0x45, 0x2d, 0x37, 0xdb, 0x31, 0xc3, 0xbe, 0xfa, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x7c, 0x61, 0x45, 0xad, 0x3e, 0x04, 0x00, 0x00,
}
