// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: google/ads/googleads/v3/common/asset_types.proto

package common

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A YouTube asset.
type YoutubeVideoAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// YouTube video id. This is the 11 character string value used in the
	// YouTube video URL.
	YoutubeVideoId *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=youtube_video_id,json=youtubeVideoId,proto3" json:"youtube_video_id,omitempty"`
}

func (x *YoutubeVideoAsset) Reset() {
	*x = YoutubeVideoAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v3_common_asset_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YoutubeVideoAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YoutubeVideoAsset) ProtoMessage() {}

func (x *YoutubeVideoAsset) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v3_common_asset_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YoutubeVideoAsset.ProtoReflect.Descriptor instead.
func (*YoutubeVideoAsset) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_common_asset_types_proto_rawDescGZIP(), []int{0}
}

func (x *YoutubeVideoAsset) GetYoutubeVideoId() *wrapperspb.StringValue {
	if x != nil {
		return x.YoutubeVideoId
	}
	return nil
}

// A MediaBundle asset.
type MediaBundleAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Media bundle (ZIP file) asset data. The format of the uploaded ZIP file
	// depends on the ad field where it will be used. For more information on the
	// format, see the documentation of the ad field where you plan on using the
	// MediaBundleAsset. This field is mutate only.
	Data *wrapperspb.BytesValue `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MediaBundleAsset) Reset() {
	*x = MediaBundleAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v3_common_asset_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaBundleAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaBundleAsset) ProtoMessage() {}

func (x *MediaBundleAsset) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v3_common_asset_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaBundleAsset.ProtoReflect.Descriptor instead.
func (*MediaBundleAsset) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_common_asset_types_proto_rawDescGZIP(), []int{1}
}

func (x *MediaBundleAsset) GetData() *wrapperspb.BytesValue {
	if x != nil {
		return x.Data
	}
	return nil
}

// An Image asset.
type ImageAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The raw bytes data of an image. This field is mutate only.
	Data *wrapperspb.BytesValue `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// File size of the image asset in bytes.
	FileSize *wrapperspb.Int64Value `protobuf:"bytes,2,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	// MIME type of the image asset.
	MimeType enums.MimeTypeEnum_MimeType `protobuf:"varint,3,opt,name=mime_type,json=mimeType,proto3,enum=google.ads.googleads.v3.enums.MimeTypeEnum_MimeType" json:"mime_type,omitempty"`
	// Metadata for this image at its original size.
	FullSize *ImageDimension `protobuf:"bytes,4,opt,name=full_size,json=fullSize,proto3" json:"full_size,omitempty"`
}

func (x *ImageAsset) Reset() {
	*x = ImageAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v3_common_asset_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageAsset) ProtoMessage() {}

func (x *ImageAsset) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v3_common_asset_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageAsset.ProtoReflect.Descriptor instead.
func (*ImageAsset) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_common_asset_types_proto_rawDescGZIP(), []int{2}
}

func (x *ImageAsset) GetData() *wrapperspb.BytesValue {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ImageAsset) GetFileSize() *wrapperspb.Int64Value {
	if x != nil {
		return x.FileSize
	}
	return nil
}

func (x *ImageAsset) GetMimeType() enums.MimeTypeEnum_MimeType {
	if x != nil {
		return x.MimeType
	}
	return enums.MimeTypeEnum_UNSPECIFIED
}

func (x *ImageAsset) GetFullSize() *ImageDimension {
	if x != nil {
		return x.FullSize
	}
	return nil
}

// Metadata for an image at a certain size, either original or resized.
type ImageDimension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Height of the image.
	HeightPixels *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=height_pixels,json=heightPixels,proto3" json:"height_pixels,omitempty"`
	// Width of the image.
	WidthPixels *wrapperspb.Int64Value `protobuf:"bytes,2,opt,name=width_pixels,json=widthPixels,proto3" json:"width_pixels,omitempty"`
	// A URL that returns the image with this height and width.
	Url *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ImageDimension) Reset() {
	*x = ImageDimension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v3_common_asset_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageDimension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageDimension) ProtoMessage() {}

func (x *ImageDimension) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v3_common_asset_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageDimension.ProtoReflect.Descriptor instead.
func (*ImageDimension) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_common_asset_types_proto_rawDescGZIP(), []int{3}
}

func (x *ImageDimension) GetHeightPixels() *wrapperspb.Int64Value {
	if x != nil {
		return x.HeightPixels
	}
	return nil
}

func (x *ImageDimension) GetWidthPixels() *wrapperspb.Int64Value {
	if x != nil {
		return x.WidthPixels
	}
	return nil
}

func (x *ImageDimension) GetUrl() *wrapperspb.StringValue {
	if x != nil {
		return x.Url
	}
	return nil
}

// A Text asset.
type TextAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Text content of the text asset.
	Text *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TextAsset) Reset() {
	*x = TextAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v3_common_asset_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextAsset) ProtoMessage() {}

func (x *TextAsset) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v3_common_asset_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextAsset.ProtoReflect.Descriptor instead.
func (*TextAsset) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_common_asset_types_proto_rawDescGZIP(), []int{4}
}

func (x *TextAsset) GetText() *wrapperspb.StringValue {
	if x != nil {
		return x.Text
	}
	return nil
}

var File_google_ads_googleads_v3_common_asset_types_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v3_common_asset_types_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x1a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5b, 0x0a, 0x11, 0x59, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x12, 0x46, 0x0a, 0x10, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x5f,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x79, 0x6f,
	0x75, 0x74, 0x75, 0x62, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x10,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x97, 0x02, 0x0a, 0x0a, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x38, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x51, 0x0a, 0x09, 0x6d,
	0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4d,
	0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4d, 0x69, 0x6d, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4b,
	0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xc2, 0x01, 0x0a, 0x0e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x40,
	0x0a, 0x0d, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0c, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x50, 0x69, 0x78, 0x65, 0x6c, 0x73,
	0x12, 0x3e, 0x0a, 0x0c, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0b, 0x77, 0x69, 0x64, 0x74, 0x68, 0x50, 0x69, 0x78, 0x65, 0x6c, 0x73,
	0x12, 0x2e, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x22, 0x3d, 0x0a, 0x09, 0x54, 0x65, 0x78, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x30, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42,
	0xea, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa2,
	0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x33, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x33,
	0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x33, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v3_common_asset_types_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v3_common_asset_types_proto_rawDescData = file_google_ads_googleads_v3_common_asset_types_proto_rawDesc
)

func file_google_ads_googleads_v3_common_asset_types_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v3_common_asset_types_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v3_common_asset_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v3_common_asset_types_proto_rawDescData)
	})
	return file_google_ads_googleads_v3_common_asset_types_proto_rawDescData
}

var file_google_ads_googleads_v3_common_asset_types_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_ads_googleads_v3_common_asset_types_proto_goTypes = []interface{}{
	(*YoutubeVideoAsset)(nil),        // 0: google.ads.googleads.v3.common.YoutubeVideoAsset
	(*MediaBundleAsset)(nil),         // 1: google.ads.googleads.v3.common.MediaBundleAsset
	(*ImageAsset)(nil),               // 2: google.ads.googleads.v3.common.ImageAsset
	(*ImageDimension)(nil),           // 3: google.ads.googleads.v3.common.ImageDimension
	(*TextAsset)(nil),                // 4: google.ads.googleads.v3.common.TextAsset
	(*wrapperspb.StringValue)(nil),   // 5: google.protobuf.StringValue
	(*wrapperspb.BytesValue)(nil),    // 6: google.protobuf.BytesValue
	(*wrapperspb.Int64Value)(nil),    // 7: google.protobuf.Int64Value
	(enums.MimeTypeEnum_MimeType)(0), // 8: google.ads.googleads.v3.enums.MimeTypeEnum.MimeType
}
var file_google_ads_googleads_v3_common_asset_types_proto_depIdxs = []int32{
	5,  // 0: google.ads.googleads.v3.common.YoutubeVideoAsset.youtube_video_id:type_name -> google.protobuf.StringValue
	6,  // 1: google.ads.googleads.v3.common.MediaBundleAsset.data:type_name -> google.protobuf.BytesValue
	6,  // 2: google.ads.googleads.v3.common.ImageAsset.data:type_name -> google.protobuf.BytesValue
	7,  // 3: google.ads.googleads.v3.common.ImageAsset.file_size:type_name -> google.protobuf.Int64Value
	8,  // 4: google.ads.googleads.v3.common.ImageAsset.mime_type:type_name -> google.ads.googleads.v3.enums.MimeTypeEnum.MimeType
	3,  // 5: google.ads.googleads.v3.common.ImageAsset.full_size:type_name -> google.ads.googleads.v3.common.ImageDimension
	7,  // 6: google.ads.googleads.v3.common.ImageDimension.height_pixels:type_name -> google.protobuf.Int64Value
	7,  // 7: google.ads.googleads.v3.common.ImageDimension.width_pixels:type_name -> google.protobuf.Int64Value
	5,  // 8: google.ads.googleads.v3.common.ImageDimension.url:type_name -> google.protobuf.StringValue
	5,  // 9: google.ads.googleads.v3.common.TextAsset.text:type_name -> google.protobuf.StringValue
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v3_common_asset_types_proto_init() }
func file_google_ads_googleads_v3_common_asset_types_proto_init() {
	if File_google_ads_googleads_v3_common_asset_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v3_common_asset_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YoutubeVideoAsset); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_ads_googleads_v3_common_asset_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaBundleAsset); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_ads_googleads_v3_common_asset_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageAsset); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_ads_googleads_v3_common_asset_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageDimension); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_ads_googleads_v3_common_asset_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextAsset); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v3_common_asset_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v3_common_asset_types_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v3_common_asset_types_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v3_common_asset_types_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v3_common_asset_types_proto = out.File
	file_google_ads_googleads_v3_common_asset_types_proto_rawDesc = nil
	file_google_ads_googleads_v3_common_asset_types_proto_goTypes = nil
	file_google_ads_googleads_v3_common_asset_types_proto_depIdxs = nil
}
