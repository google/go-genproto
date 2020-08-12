// Copyright 2017 Google Inc.
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
// source: google/bigtable/admin/table/v1/bigtable_table_service.proto

package table

import (
	context "context"
	reflect "reflect"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

var File_google_bigtable_admin_table_v1_bigtable_table_service_proto protoreflect.FileDescriptor

var file_google_bigtable_admin_table_v1_bigtable_table_service_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x67, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x62, 0x69,
	0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbe, 0x0c, 0x0a, 0x14, 0x42, 0x69, 0x67,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xa4, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62,
	0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x3a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x34, 0x22, 0x2f, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73,
	0x2f, 0x2a, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xac, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x12, 0x2f, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x7a, 0x6f, 0x6e,
	0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d,
	0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x9d, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69,
	0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62,
	0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x39, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x33, 0x12, 0x31, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73,
	0x2f, 0x2a, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x94, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x39, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x2a, 0x31, 0x2f, 0x76, 0x31,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x2a, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x9e,
	0x01, 0x0a, 0x0b, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x32,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x43, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x3d, 0x22, 0x38, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x2a,
	0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0xca, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x22,
	0x4b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x45, 0x22, 0x40, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x7a, 0x6f,
	0x6e, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a,
	0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x69, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xbf, 0x01, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x46, 0x61, 0x6d,
	0x69, 0x6c, 0x79, 0x12, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x46, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x1a, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x22,
	0x4d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x47, 0x1a, 0x42, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x7a, 0x6f,
	0x6e, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a,
	0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x46, 0x61, 0x6d, 0x69, 0x6c, 0x69, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0xb3,
	0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x46,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62,
	0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x4a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x44,
	0x2a, 0x42, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x2f, 0x2a, 0x2f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x69, 0x65,
	0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xb2, 0x01, 0x0a, 0x0e, 0x42, 0x75, 0x6c, 0x6b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x51, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4b, 0x22, 0x46,
	0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73,
	0x2f, 0x2a, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x62, 0x75, 0x6c, 0x6b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x6f, 0x77, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x87, 0x01, 0x0a, 0x22, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x1a, 0x42, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_google_bigtable_admin_table_v1_bigtable_table_service_proto_goTypes = []interface{}{
	(*CreateTableRequest)(nil),        // 0: google.bigtable.admin.table.v1.CreateTableRequest
	(*ListTablesRequest)(nil),         // 1: google.bigtable.admin.table.v1.ListTablesRequest
	(*GetTableRequest)(nil),           // 2: google.bigtable.admin.table.v1.GetTableRequest
	(*DeleteTableRequest)(nil),        // 3: google.bigtable.admin.table.v1.DeleteTableRequest
	(*RenameTableRequest)(nil),        // 4: google.bigtable.admin.table.v1.RenameTableRequest
	(*CreateColumnFamilyRequest)(nil), // 5: google.bigtable.admin.table.v1.CreateColumnFamilyRequest
	(*ColumnFamily)(nil),              // 6: google.bigtable.admin.table.v1.ColumnFamily
	(*DeleteColumnFamilyRequest)(nil), // 7: google.bigtable.admin.table.v1.DeleteColumnFamilyRequest
	(*BulkDeleteRowsRequest)(nil),     // 8: google.bigtable.admin.table.v1.BulkDeleteRowsRequest
	(*Table)(nil),                     // 9: google.bigtable.admin.table.v1.Table
	(*ListTablesResponse)(nil),        // 10: google.bigtable.admin.table.v1.ListTablesResponse
	(*empty.Empty)(nil),               // 11: google.protobuf.Empty
}
var file_google_bigtable_admin_table_v1_bigtable_table_service_proto_depIdxs = []int32{
	0,  // 0: google.bigtable.admin.table.v1.BigtableTableService.CreateTable:input_type -> google.bigtable.admin.table.v1.CreateTableRequest
	1,  // 1: google.bigtable.admin.table.v1.BigtableTableService.ListTables:input_type -> google.bigtable.admin.table.v1.ListTablesRequest
	2,  // 2: google.bigtable.admin.table.v1.BigtableTableService.GetTable:input_type -> google.bigtable.admin.table.v1.GetTableRequest
	3,  // 3: google.bigtable.admin.table.v1.BigtableTableService.DeleteTable:input_type -> google.bigtable.admin.table.v1.DeleteTableRequest
	4,  // 4: google.bigtable.admin.table.v1.BigtableTableService.RenameTable:input_type -> google.bigtable.admin.table.v1.RenameTableRequest
	5,  // 5: google.bigtable.admin.table.v1.BigtableTableService.CreateColumnFamily:input_type -> google.bigtable.admin.table.v1.CreateColumnFamilyRequest
	6,  // 6: google.bigtable.admin.table.v1.BigtableTableService.UpdateColumnFamily:input_type -> google.bigtable.admin.table.v1.ColumnFamily
	7,  // 7: google.bigtable.admin.table.v1.BigtableTableService.DeleteColumnFamily:input_type -> google.bigtable.admin.table.v1.DeleteColumnFamilyRequest
	8,  // 8: google.bigtable.admin.table.v1.BigtableTableService.BulkDeleteRows:input_type -> google.bigtable.admin.table.v1.BulkDeleteRowsRequest
	9,  // 9: google.bigtable.admin.table.v1.BigtableTableService.CreateTable:output_type -> google.bigtable.admin.table.v1.Table
	10, // 10: google.bigtable.admin.table.v1.BigtableTableService.ListTables:output_type -> google.bigtable.admin.table.v1.ListTablesResponse
	9,  // 11: google.bigtable.admin.table.v1.BigtableTableService.GetTable:output_type -> google.bigtable.admin.table.v1.Table
	11, // 12: google.bigtable.admin.table.v1.BigtableTableService.DeleteTable:output_type -> google.protobuf.Empty
	11, // 13: google.bigtable.admin.table.v1.BigtableTableService.RenameTable:output_type -> google.protobuf.Empty
	6,  // 14: google.bigtable.admin.table.v1.BigtableTableService.CreateColumnFamily:output_type -> google.bigtable.admin.table.v1.ColumnFamily
	6,  // 15: google.bigtable.admin.table.v1.BigtableTableService.UpdateColumnFamily:output_type -> google.bigtable.admin.table.v1.ColumnFamily
	11, // 16: google.bigtable.admin.table.v1.BigtableTableService.DeleteColumnFamily:output_type -> google.protobuf.Empty
	11, // 17: google.bigtable.admin.table.v1.BigtableTableService.BulkDeleteRows:output_type -> google.protobuf.Empty
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_google_bigtable_admin_table_v1_bigtable_table_service_proto_init() }
func file_google_bigtable_admin_table_v1_bigtable_table_service_proto_init() {
	if File_google_bigtable_admin_table_v1_bigtable_table_service_proto != nil {
		return
	}
	file_google_bigtable_admin_table_v1_bigtable_table_data_proto_init()
	file_google_bigtable_admin_table_v1_bigtable_table_service_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_bigtable_admin_table_v1_bigtable_table_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_bigtable_admin_table_v1_bigtable_table_service_proto_goTypes,
		DependencyIndexes: file_google_bigtable_admin_table_v1_bigtable_table_service_proto_depIdxs,
	}.Build()
	File_google_bigtable_admin_table_v1_bigtable_table_service_proto = out.File
	file_google_bigtable_admin_table_v1_bigtable_table_service_proto_rawDesc = nil
	file_google_bigtable_admin_table_v1_bigtable_table_service_proto_goTypes = nil
	file_google_bigtable_admin_table_v1_bigtable_table_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BigtableTableServiceClient is the client API for BigtableTableService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BigtableTableServiceClient interface {
	// Creates a new table, to be served from a specified cluster.
	// The table can be created with a full set of initial column families,
	// specified in the request.
	CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*Table, error)
	// Lists the names of all tables served from a specified cluster.
	ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error)
	// Gets the schema of the specified table, including its column families.
	GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*Table, error)
	// Permanently deletes a specified table and all of its data.
	DeleteTable(ctx context.Context, in *DeleteTableRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Changes the name of a specified table.
	// Cannot be used to move tables between clusters, zones, or projects.
	RenameTable(ctx context.Context, in *RenameTableRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Creates a new column family within a specified table.
	CreateColumnFamily(ctx context.Context, in *CreateColumnFamilyRequest, opts ...grpc.CallOption) (*ColumnFamily, error)
	// Changes the configuration of a specified column family.
	UpdateColumnFamily(ctx context.Context, in *ColumnFamily, opts ...grpc.CallOption) (*ColumnFamily, error)
	// Permanently deletes a specified column family and all of its data.
	DeleteColumnFamily(ctx context.Context, in *DeleteColumnFamilyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete all rows in a table corresponding to a particular prefix
	BulkDeleteRows(ctx context.Context, in *BulkDeleteRowsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type bigtableTableServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBigtableTableServiceClient(cc grpc.ClientConnInterface) BigtableTableServiceClient {
	return &bigtableTableServiceClient{cc}
}

func (c *bigtableTableServiceClient) CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*Table, error) {
	out := new(Table)
	err := c.cc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/CreateTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error) {
	out := new(ListTablesResponse)
	err := c.cc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/ListTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*Table, error) {
	out := new(Table)
	err := c.cc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/GetTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) DeleteTable(ctx context.Context, in *DeleteTableRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/DeleteTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) RenameTable(ctx context.Context, in *RenameTableRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/RenameTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) CreateColumnFamily(ctx context.Context, in *CreateColumnFamilyRequest, opts ...grpc.CallOption) (*ColumnFamily, error) {
	out := new(ColumnFamily)
	err := c.cc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/CreateColumnFamily", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) UpdateColumnFamily(ctx context.Context, in *ColumnFamily, opts ...grpc.CallOption) (*ColumnFamily, error) {
	out := new(ColumnFamily)
	err := c.cc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/UpdateColumnFamily", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) DeleteColumnFamily(ctx context.Context, in *DeleteColumnFamilyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/DeleteColumnFamily", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) BulkDeleteRows(ctx context.Context, in *BulkDeleteRowsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/BulkDeleteRows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BigtableTableServiceServer is the server API for BigtableTableService service.
type BigtableTableServiceServer interface {
	// Creates a new table, to be served from a specified cluster.
	// The table can be created with a full set of initial column families,
	// specified in the request.
	CreateTable(context.Context, *CreateTableRequest) (*Table, error)
	// Lists the names of all tables served from a specified cluster.
	ListTables(context.Context, *ListTablesRequest) (*ListTablesResponse, error)
	// Gets the schema of the specified table, including its column families.
	GetTable(context.Context, *GetTableRequest) (*Table, error)
	// Permanently deletes a specified table and all of its data.
	DeleteTable(context.Context, *DeleteTableRequest) (*empty.Empty, error)
	// Changes the name of a specified table.
	// Cannot be used to move tables between clusters, zones, or projects.
	RenameTable(context.Context, *RenameTableRequest) (*empty.Empty, error)
	// Creates a new column family within a specified table.
	CreateColumnFamily(context.Context, *CreateColumnFamilyRequest) (*ColumnFamily, error)
	// Changes the configuration of a specified column family.
	UpdateColumnFamily(context.Context, *ColumnFamily) (*ColumnFamily, error)
	// Permanently deletes a specified column family and all of its data.
	DeleteColumnFamily(context.Context, *DeleteColumnFamilyRequest) (*empty.Empty, error)
	// Delete all rows in a table corresponding to a particular prefix
	BulkDeleteRows(context.Context, *BulkDeleteRowsRequest) (*empty.Empty, error)
}

// UnimplementedBigtableTableServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBigtableTableServiceServer struct {
}

func (*UnimplementedBigtableTableServiceServer) CreateTable(context.Context, *CreateTableRequest) (*Table, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTable not implemented")
}
func (*UnimplementedBigtableTableServiceServer) ListTables(context.Context, *ListTablesRequest) (*ListTablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTables not implemented")
}
func (*UnimplementedBigtableTableServiceServer) GetTable(context.Context, *GetTableRequest) (*Table, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTable not implemented")
}
func (*UnimplementedBigtableTableServiceServer) DeleteTable(context.Context, *DeleteTableRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTable not implemented")
}
func (*UnimplementedBigtableTableServiceServer) RenameTable(context.Context, *RenameTableRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameTable not implemented")
}
func (*UnimplementedBigtableTableServiceServer) CreateColumnFamily(context.Context, *CreateColumnFamilyRequest) (*ColumnFamily, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateColumnFamily not implemented")
}
func (*UnimplementedBigtableTableServiceServer) UpdateColumnFamily(context.Context, *ColumnFamily) (*ColumnFamily, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateColumnFamily not implemented")
}
func (*UnimplementedBigtableTableServiceServer) DeleteColumnFamily(context.Context, *DeleteColumnFamilyRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteColumnFamily not implemented")
}
func (*UnimplementedBigtableTableServiceServer) BulkDeleteRows(context.Context, *BulkDeleteRowsRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkDeleteRows not implemented")
}

func RegisterBigtableTableServiceServer(s *grpc.Server, srv BigtableTableServiceServer) {
	s.RegisterService(&_BigtableTableService_serviceDesc, srv)
}

func _BigtableTableService_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/CreateTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).CreateTable(ctx, req.(*CreateTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_ListTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).ListTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/ListTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).ListTables(ctx, req.(*ListTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_GetTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).GetTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/GetTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).GetTable(ctx, req.(*GetTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_DeleteTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).DeleteTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/DeleteTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).DeleteTable(ctx, req.(*DeleteTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_RenameTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).RenameTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/RenameTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).RenameTable(ctx, req.(*RenameTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_CreateColumnFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateColumnFamilyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).CreateColumnFamily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/CreateColumnFamily",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).CreateColumnFamily(ctx, req.(*CreateColumnFamilyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_UpdateColumnFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ColumnFamily)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).UpdateColumnFamily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/UpdateColumnFamily",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).UpdateColumnFamily(ctx, req.(*ColumnFamily))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_DeleteColumnFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteColumnFamilyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).DeleteColumnFamily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/DeleteColumnFamily",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).DeleteColumnFamily(ctx, req.(*DeleteColumnFamilyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_BulkDeleteRows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkDeleteRowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).BulkDeleteRows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/BulkDeleteRows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).BulkDeleteRows(ctx, req.(*BulkDeleteRowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BigtableTableService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.bigtable.admin.table.v1.BigtableTableService",
	HandlerType: (*BigtableTableServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTable",
			Handler:    _BigtableTableService_CreateTable_Handler,
		},
		{
			MethodName: "ListTables",
			Handler:    _BigtableTableService_ListTables_Handler,
		},
		{
			MethodName: "GetTable",
			Handler:    _BigtableTableService_GetTable_Handler,
		},
		{
			MethodName: "DeleteTable",
			Handler:    _BigtableTableService_DeleteTable_Handler,
		},
		{
			MethodName: "RenameTable",
			Handler:    _BigtableTableService_RenameTable_Handler,
		},
		{
			MethodName: "CreateColumnFamily",
			Handler:    _BigtableTableService_CreateColumnFamily_Handler,
		},
		{
			MethodName: "UpdateColumnFamily",
			Handler:    _BigtableTableService_UpdateColumnFamily_Handler,
		},
		{
			MethodName: "DeleteColumnFamily",
			Handler:    _BigtableTableService_DeleteColumnFamily_Handler,
		},
		{
			MethodName: "BulkDeleteRows",
			Handler:    _BigtableTableService_BulkDeleteRows_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/bigtable/admin/table/v1/bigtable_table_service.proto",
}
