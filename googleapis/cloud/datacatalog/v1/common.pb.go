// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datacatalog/v1/common.proto

package datacatalog

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// This enum describes all the possible systems that Data Catalog integrates
// with.
type IntegratedSystem int32

const (
	// Default unknown system.
	IntegratedSystem_INTEGRATED_SYSTEM_UNSPECIFIED IntegratedSystem = 0
	// BigQuery.
	IntegratedSystem_BIGQUERY IntegratedSystem = 1
	// Cloud Pub/Sub.
	IntegratedSystem_CLOUD_PUBSUB IntegratedSystem = 2
)

var IntegratedSystem_name = map[int32]string{
	0: "INTEGRATED_SYSTEM_UNSPECIFIED",
	1: "BIGQUERY",
	2: "CLOUD_PUBSUB",
}

var IntegratedSystem_value = map[string]int32{
	"INTEGRATED_SYSTEM_UNSPECIFIED": 0,
	"BIGQUERY":                      1,
	"CLOUD_PUBSUB":                  2,
}

func (x IntegratedSystem) String() string {
	return proto.EnumName(IntegratedSystem_name, int32(x))
}

func (IntegratedSystem) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fec73966e023eb05, []int{0}
}

func init() {
	proto.RegisterEnum("google.cloud.datacatalog.v1.IntegratedSystem", IntegratedSystem_name, IntegratedSystem_value)
}

func init() {
	proto.RegisterFile("google/cloud/datacatalog/v1/common.proto", fileDescriptor_fec73966e023eb05)
}

var fileDescriptor_fec73966e023eb05 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0x4f, 0x49, 0x2c, 0x49, 0x4c, 0x4e, 0x2c,
	0x49, 0xcc, 0xc9, 0x4f, 0xd7, 0x2f, 0x33, 0xd4, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x86, 0xa8, 0xd4, 0x03, 0xab, 0xd4, 0x43, 0x52, 0xa9, 0x57,
	0x66, 0xa8, 0x15, 0xca, 0x25, 0xe0, 0x99, 0x57, 0x92, 0x9a, 0x5e, 0x94, 0x58, 0x92, 0x9a, 0x12,
	0x5c, 0x59, 0x5c, 0x92, 0x9a, 0x2b, 0xa4, 0xc8, 0x25, 0xeb, 0xe9, 0x17, 0xe2, 0xea, 0x1e, 0xe4,
	0x18, 0xe2, 0xea, 0x12, 0x1f, 0x1c, 0x19, 0x1c, 0xe2, 0xea, 0x1b, 0x1f, 0xea, 0x17, 0x1c, 0xe0,
	0xea, 0xec, 0xe9, 0xe6, 0xe9, 0xea, 0x22, 0xc0, 0x20, 0xc4, 0xc3, 0xc5, 0xe1, 0xe4, 0xe9, 0x1e,
	0x18, 0xea, 0x1a, 0x14, 0x29, 0xc0, 0x28, 0x24, 0xc0, 0xc5, 0xe3, 0xec, 0xe3, 0x1f, 0xea, 0x12,
	0x1f, 0x10, 0xea, 0x14, 0x1c, 0xea, 0x24, 0xc0, 0xe4, 0x94, 0xc7, 0x25, 0x9f, 0x9c, 0x9f, 0xab,
	0x87, 0xc7, 0xe6, 0x00, 0xc6, 0x28, 0x37, 0xa8, 0x74, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e,
	0x7e, 0x51, 0xba, 0x7e, 0x7a, 0x6a, 0x1e, 0xd8, 0xd1, 0xfa, 0x10, 0xa9, 0xc4, 0x82, 0xcc, 0x62,
	0xac, 0x3e, 0xb4, 0x46, 0xe2, 0xfe, 0x60, 0x64, 0x4c, 0x62, 0x03, 0xeb, 0x32, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x29, 0x41, 0xc8, 0x6b, 0x16, 0x01, 0x00, 0x00,
}
