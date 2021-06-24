// Copyright 2021 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/aiplatform/v1/schema/trainingjob/definition/automl_video_classification.proto

package definition

import (
	reflect "reflect"
	sync "sync"

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

type AutoMlVideoClassificationInputs_ModelType int32

const (
	// Should not be set.
	AutoMlVideoClassificationInputs_MODEL_TYPE_UNSPECIFIED AutoMlVideoClassificationInputs_ModelType = 0
	// A model best tailored to be used within Google Cloud, and which cannot
	// be exported. Default.
	AutoMlVideoClassificationInputs_CLOUD AutoMlVideoClassificationInputs_ModelType = 1
	// A model that, in addition to being available within Google Cloud, can
	// also be exported (see ModelService.ExportModel) as a TensorFlow or
	// TensorFlow Lite model and used on a mobile or edge device afterwards.
	AutoMlVideoClassificationInputs_MOBILE_VERSATILE_1 AutoMlVideoClassificationInputs_ModelType = 2
	// A model that, in addition to being available within Google Cloud, can
	// also be exported (see ModelService.ExportModel) to a Jetson device
	// afterwards.
	AutoMlVideoClassificationInputs_MOBILE_JETSON_VERSATILE_1 AutoMlVideoClassificationInputs_ModelType = 3
)

// Enum value maps for AutoMlVideoClassificationInputs_ModelType.
var (
	AutoMlVideoClassificationInputs_ModelType_name = map[int32]string{
		0: "MODEL_TYPE_UNSPECIFIED",
		1: "CLOUD",
		2: "MOBILE_VERSATILE_1",
		3: "MOBILE_JETSON_VERSATILE_1",
	}
	AutoMlVideoClassificationInputs_ModelType_value = map[string]int32{
		"MODEL_TYPE_UNSPECIFIED":    0,
		"CLOUD":                     1,
		"MOBILE_VERSATILE_1":        2,
		"MOBILE_JETSON_VERSATILE_1": 3,
	}
)

func (x AutoMlVideoClassificationInputs_ModelType) Enum() *AutoMlVideoClassificationInputs_ModelType {
	p := new(AutoMlVideoClassificationInputs_ModelType)
	*p = x
	return p
}

func (x AutoMlVideoClassificationInputs_ModelType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AutoMlVideoClassificationInputs_ModelType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_enumTypes[0].Descriptor()
}

func (AutoMlVideoClassificationInputs_ModelType) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_enumTypes[0]
}

func (x AutoMlVideoClassificationInputs_ModelType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AutoMlVideoClassificationInputs_ModelType.Descriptor instead.
func (AutoMlVideoClassificationInputs_ModelType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_rawDescGZIP(), []int{1, 0}
}

// A TrainingJob that trains and uploads an AutoML Video Classification Model.
type AutoMlVideoClassification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The input parameters of this TrainingJob.
	Inputs *AutoMlVideoClassificationInputs `protobuf:"bytes,1,opt,name=inputs,proto3" json:"inputs,omitempty"`
}

func (x *AutoMlVideoClassification) Reset() {
	*x = AutoMlVideoClassification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoMlVideoClassification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoMlVideoClassification) ProtoMessage() {}

func (x *AutoMlVideoClassification) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoMlVideoClassification.ProtoReflect.Descriptor instead.
func (*AutoMlVideoClassification) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_rawDescGZIP(), []int{0}
}

func (x *AutoMlVideoClassification) GetInputs() *AutoMlVideoClassificationInputs {
	if x != nil {
		return x.Inputs
	}
	return nil
}

type AutoMlVideoClassificationInputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelType AutoMlVideoClassificationInputs_ModelType `protobuf:"varint,1,opt,name=model_type,json=modelType,proto3,enum=google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlVideoClassificationInputs_ModelType" json:"model_type,omitempty"`
}

func (x *AutoMlVideoClassificationInputs) Reset() {
	*x = AutoMlVideoClassificationInputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoMlVideoClassificationInputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoMlVideoClassificationInputs) ProtoMessage() {}

func (x *AutoMlVideoClassificationInputs) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoMlVideoClassificationInputs.ProtoReflect.Descriptor instead.
func (*AutoMlVideoClassificationInputs) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_rawDescGZIP(), []int{1}
}

func (x *AutoMlVideoClassificationInputs) GetModelType() AutoMlVideoClassificationInputs_ModelType {
	if x != nil {
		return x.ModelType
	}
	return AutoMlVideoClassificationInputs_MODEL_TYPE_UNSPECIFIED
}

var File_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_rawDesc = []byte{
	0x0a, 0x5a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x6a, 0x6f, 0x62, 0x2f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x6c, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x38, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x6a, 0x6f, 0x62, 0x2e, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x19, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x71, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x59, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x6a,
	0x6f, 0x62, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x75,
	0x74, 0x6f, 0x4d, 0x6c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x52, 0x06, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x73, 0x22, 0x91, 0x02, 0x0a, 0x1f, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x82, 0x01, 0x0a, 0x0a, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x63,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x6a, 0x6f, 0x62, 0x2e, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x22, 0x69,
	0x0a, 0x09, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4d,
	0x4f, 0x44, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x4f, 0x55, 0x44,
	0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x56, 0x45, 0x52,
	0x53, 0x41, 0x54, 0x49, 0x4c, 0x45, 0x5f, 0x31, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x4f,
	0x42, 0x49, 0x4c, 0x45, 0x5f, 0x4a, 0x45, 0x54, 0x53, 0x4f, 0x4e, 0x5f, 0x56, 0x45, 0x52, 0x53,
	0x41, 0x54, 0x49, 0x4c, 0x45, 0x5f, 0x31, 0x10, 0x03, 0x42, 0xc4, 0x01, 0x0a, 0x3c, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x6a, 0x6f, 0x62, 0x2e,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x1e, 0x41, 0x75, 0x74, 0x6f,
	0x4d, 0x4c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x62, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x6a, 0x6f, 0x62, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_rawDescData = file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_goTypes = []interface{}{
	(AutoMlVideoClassificationInputs_ModelType)(0), // 0: google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlVideoClassificationInputs.ModelType
	(*AutoMlVideoClassification)(nil),              // 1: google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlVideoClassification
	(*AutoMlVideoClassificationInputs)(nil),        // 2: google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlVideoClassificationInputs
}
var file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_depIdxs = []int32{
	2, // 0: google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlVideoClassification.inputs:type_name -> google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlVideoClassificationInputs
	0, // 1: google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlVideoClassificationInputs.model_type:type_name -> google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlVideoClassificationInputs.ModelType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_init()
}
func file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_init() {
	if File_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoMlVideoClassification); i {
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
		file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoMlVideoClassificationInputs); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_depIdxs,
		EnumInfos:         file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_enumTypes,
		MessageInfos:      file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto = out.File
	file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_video_classification_proto_depIdxs = nil
}
