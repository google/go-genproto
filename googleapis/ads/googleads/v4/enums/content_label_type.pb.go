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
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/ads/googleads/v4/enums/content_label_type.proto

package enums

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Enum listing the content label types supported by ContentLabel criterion.
type ContentLabelTypeEnum_ContentLabelType int32

const (
	// Not specified.
	ContentLabelTypeEnum_UNSPECIFIED ContentLabelTypeEnum_ContentLabelType = 0
	// Used for return value only. Represents value unknown in this version.
	ContentLabelTypeEnum_UNKNOWN ContentLabelTypeEnum_ContentLabelType = 1
	// Sexually suggestive content.
	ContentLabelTypeEnum_SEXUALLY_SUGGESTIVE ContentLabelTypeEnum_ContentLabelType = 2
	// Below the fold placement.
	ContentLabelTypeEnum_BELOW_THE_FOLD ContentLabelTypeEnum_ContentLabelType = 3
	// Parked domain.
	ContentLabelTypeEnum_PARKED_DOMAIN ContentLabelTypeEnum_ContentLabelType = 4
	// Juvenile, gross & bizarre content.
	ContentLabelTypeEnum_JUVENILE ContentLabelTypeEnum_ContentLabelType = 6
	// Profanity & rough language.
	ContentLabelTypeEnum_PROFANITY ContentLabelTypeEnum_ContentLabelType = 7
	// Death & tragedy.
	ContentLabelTypeEnum_TRAGEDY ContentLabelTypeEnum_ContentLabelType = 8
	// Video.
	ContentLabelTypeEnum_VIDEO ContentLabelTypeEnum_ContentLabelType = 9
	// Content rating: G.
	ContentLabelTypeEnum_VIDEO_RATING_DV_G ContentLabelTypeEnum_ContentLabelType = 10
	// Content rating: PG.
	ContentLabelTypeEnum_VIDEO_RATING_DV_PG ContentLabelTypeEnum_ContentLabelType = 11
	// Content rating: T.
	ContentLabelTypeEnum_VIDEO_RATING_DV_T ContentLabelTypeEnum_ContentLabelType = 12
	// Content rating: MA.
	ContentLabelTypeEnum_VIDEO_RATING_DV_MA ContentLabelTypeEnum_ContentLabelType = 13
	// Content rating: not yet rated.
	ContentLabelTypeEnum_VIDEO_NOT_YET_RATED ContentLabelTypeEnum_ContentLabelType = 14
	// Embedded video.
	ContentLabelTypeEnum_EMBEDDED_VIDEO ContentLabelTypeEnum_ContentLabelType = 15
	// Live streaming video.
	ContentLabelTypeEnum_LIVE_STREAMING_VIDEO ContentLabelTypeEnum_ContentLabelType = 16
	// Sensitive social issues.
	ContentLabelTypeEnum_SOCIAL_ISSUES ContentLabelTypeEnum_ContentLabelType = 17
)

// Enum value maps for ContentLabelTypeEnum_ContentLabelType.
var (
	ContentLabelTypeEnum_ContentLabelType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "SEXUALLY_SUGGESTIVE",
		3:  "BELOW_THE_FOLD",
		4:  "PARKED_DOMAIN",
		6:  "JUVENILE",
		7:  "PROFANITY",
		8:  "TRAGEDY",
		9:  "VIDEO",
		10: "VIDEO_RATING_DV_G",
		11: "VIDEO_RATING_DV_PG",
		12: "VIDEO_RATING_DV_T",
		13: "VIDEO_RATING_DV_MA",
		14: "VIDEO_NOT_YET_RATED",
		15: "EMBEDDED_VIDEO",
		16: "LIVE_STREAMING_VIDEO",
		17: "SOCIAL_ISSUES",
	}
	ContentLabelTypeEnum_ContentLabelType_value = map[string]int32{
		"UNSPECIFIED":          0,
		"UNKNOWN":              1,
		"SEXUALLY_SUGGESTIVE":  2,
		"BELOW_THE_FOLD":       3,
		"PARKED_DOMAIN":        4,
		"JUVENILE":             6,
		"PROFANITY":            7,
		"TRAGEDY":              8,
		"VIDEO":                9,
		"VIDEO_RATING_DV_G":    10,
		"VIDEO_RATING_DV_PG":   11,
		"VIDEO_RATING_DV_T":    12,
		"VIDEO_RATING_DV_MA":   13,
		"VIDEO_NOT_YET_RATED":  14,
		"EMBEDDED_VIDEO":       15,
		"LIVE_STREAMING_VIDEO": 16,
		"SOCIAL_ISSUES":        17,
	}
)

func (x ContentLabelTypeEnum_ContentLabelType) Enum() *ContentLabelTypeEnum_ContentLabelType {
	p := new(ContentLabelTypeEnum_ContentLabelType)
	*p = x
	return p
}

func (x ContentLabelTypeEnum_ContentLabelType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContentLabelTypeEnum_ContentLabelType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v4_enums_content_label_type_proto_enumTypes[0].Descriptor()
}

func (ContentLabelTypeEnum_ContentLabelType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v4_enums_content_label_type_proto_enumTypes[0]
}

func (x ContentLabelTypeEnum_ContentLabelType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContentLabelTypeEnum_ContentLabelType.Descriptor instead.
func (ContentLabelTypeEnum_ContentLabelType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v4_enums_content_label_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing content label types in ContentLabel.
type ContentLabelTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContentLabelTypeEnum) Reset() {
	*x = ContentLabelTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v4_enums_content_label_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentLabelTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentLabelTypeEnum) ProtoMessage() {}

func (x *ContentLabelTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v4_enums_content_label_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentLabelTypeEnum.ProtoReflect.Descriptor instead.
func (*ContentLabelTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v4_enums_content_label_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v4_enums_content_label_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v4_enums_content_label_type_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x34, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x02, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xdd,
	0x02, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x01, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x58, 0x55, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x53, 0x55,
	0x47, 0x47, 0x45, 0x53, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x45,
	0x4c, 0x4f, 0x57, 0x5f, 0x54, 0x48, 0x45, 0x5f, 0x46, 0x4f, 0x4c, 0x44, 0x10, 0x03, 0x12, 0x11,
	0x0a, 0x0d, 0x50, 0x41, 0x52, 0x4b, 0x45, 0x44, 0x5f, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x10,
	0x04, 0x12, 0x0c, 0x0a, 0x08, 0x4a, 0x55, 0x56, 0x45, 0x4e, 0x49, 0x4c, 0x45, 0x10, 0x06, 0x12,
	0x0d, 0x0a, 0x09, 0x50, 0x52, 0x4f, 0x46, 0x41, 0x4e, 0x49, 0x54, 0x59, 0x10, 0x07, 0x12, 0x0b,
	0x0a, 0x07, 0x54, 0x52, 0x41, 0x47, 0x45, 0x44, 0x59, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x56,
	0x49, 0x44, 0x45, 0x4f, 0x10, 0x09, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f,
	0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x56, 0x5f, 0x47, 0x10, 0x0a, 0x12, 0x16, 0x0a,
	0x12, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x56,
	0x5f, 0x50, 0x47, 0x10, 0x0b, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x52,
	0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x56, 0x5f, 0x54, 0x10, 0x0c, 0x12, 0x16, 0x0a, 0x12,
	0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x56, 0x5f,
	0x4d, 0x41, 0x10, 0x0d, 0x12, 0x17, 0x0a, 0x13, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x59, 0x45, 0x54, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x44, 0x10, 0x0e, 0x12, 0x12, 0x0a,
	0x0e, 0x45, 0x4d, 0x42, 0x45, 0x44, 0x44, 0x45, 0x44, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10,
	0x0f, 0x12, 0x18, 0x0a, 0x14, 0x4c, 0x49, 0x56, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d,
	0x49, 0x4e, 0x47, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x10, 0x12, 0x11, 0x0a, 0x0d, 0x53,
	0x4f, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x49, 0x53, 0x53, 0x55, 0x45, 0x53, 0x10, 0x11, 0x42, 0xea,
	0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x42, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x34, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56,
	0x34, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x34, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v4_enums_content_label_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v4_enums_content_label_type_proto_rawDescData = file_google_ads_googleads_v4_enums_content_label_type_proto_rawDesc
)

func file_google_ads_googleads_v4_enums_content_label_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v4_enums_content_label_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v4_enums_content_label_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v4_enums_content_label_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v4_enums_content_label_type_proto_rawDescData
}

var file_google_ads_googleads_v4_enums_content_label_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v4_enums_content_label_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v4_enums_content_label_type_proto_goTypes = []interface{}{
	(ContentLabelTypeEnum_ContentLabelType)(0), // 0: google.ads.googleads.v4.enums.ContentLabelTypeEnum.ContentLabelType
	(*ContentLabelTypeEnum)(nil),               // 1: google.ads.googleads.v4.enums.ContentLabelTypeEnum
}
var file_google_ads_googleads_v4_enums_content_label_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v4_enums_content_label_type_proto_init() }
func file_google_ads_googleads_v4_enums_content_label_type_proto_init() {
	if File_google_ads_googleads_v4_enums_content_label_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v4_enums_content_label_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentLabelTypeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v4_enums_content_label_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v4_enums_content_label_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v4_enums_content_label_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v4_enums_content_label_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v4_enums_content_label_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v4_enums_content_label_type_proto = out.File
	file_google_ads_googleads_v4_enums_content_label_type_proto_rawDesc = nil
	file_google_ads_googleads_v4_enums_content_label_type_proto_goTypes = nil
	file_google_ads_googleads_v4_enums_content_label_type_proto_depIdxs = nil
}
