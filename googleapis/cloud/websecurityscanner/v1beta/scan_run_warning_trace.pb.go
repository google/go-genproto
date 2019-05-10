// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/websecurityscanner/v1beta/scan_run_warning_trace.proto

package websecurityscanner // import "google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1beta"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Output only.
// Defines a warning message code.
// Next id: 5
type ScanRunWarningTrace_Code int32

const (
	// Default value is never used.
	ScanRunWarningTrace_CODE_UNSPECIFIED ScanRunWarningTrace_Code = 0
	// Indicates that a scan discovered an unexpectedly low number of URLs. This
	// is sometimes caused by complex navigation features or by using a single
	// URL for numerous pages.
	ScanRunWarningTrace_INSUFFICIENT_CRAWL_RESULTS ScanRunWarningTrace_Code = 1
	// Indicates that a scan discovered too many URLs to test, or excessive
	// redundant URLs.
	ScanRunWarningTrace_TOO_MANY_CRAWL_RESULTS ScanRunWarningTrace_Code = 2
	// Indicates that too many tests have been generated for the scan. Customer
	// should try reducing the number of starting URLs, increasing the QPS rate,
	// or narrowing down the scope of the scan using the excluded patterns.
	ScanRunWarningTrace_TOO_MANY_FUZZ_TASKS ScanRunWarningTrace_Code = 3
	// Indicates that a scan is blocked by IAP.
	ScanRunWarningTrace_BLOCKED_BY_IAP ScanRunWarningTrace_Code = 4
)

var ScanRunWarningTrace_Code_name = map[int32]string{
	0: "CODE_UNSPECIFIED",
	1: "INSUFFICIENT_CRAWL_RESULTS",
	2: "TOO_MANY_CRAWL_RESULTS",
	3: "TOO_MANY_FUZZ_TASKS",
	4: "BLOCKED_BY_IAP",
}
var ScanRunWarningTrace_Code_value = map[string]int32{
	"CODE_UNSPECIFIED":           0,
	"INSUFFICIENT_CRAWL_RESULTS": 1,
	"TOO_MANY_CRAWL_RESULTS":     2,
	"TOO_MANY_FUZZ_TASKS":        3,
	"BLOCKED_BY_IAP":             4,
}

func (x ScanRunWarningTrace_Code) String() string {
	return proto.EnumName(ScanRunWarningTrace_Code_name, int32(x))
}
func (ScanRunWarningTrace_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_scan_run_warning_trace_f2bc1463bc284e39, []int{0, 0}
}

// Output only.
// Defines a warning trace message for ScanRun. Warning traces provide customers
// with useful information that helps make the scanning process more effective.
type ScanRunWarningTrace struct {
	// Output only.
	// Indicates the warning code.
	Code                 ScanRunWarningTrace_Code `protobuf:"varint,1,opt,name=code,proto3,enum=google.cloud.websecurityscanner.v1beta.ScanRunWarningTrace_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ScanRunWarningTrace) Reset()         { *m = ScanRunWarningTrace{} }
func (m *ScanRunWarningTrace) String() string { return proto.CompactTextString(m) }
func (*ScanRunWarningTrace) ProtoMessage()    {}
func (*ScanRunWarningTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_scan_run_warning_trace_f2bc1463bc284e39, []int{0}
}
func (m *ScanRunWarningTrace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanRunWarningTrace.Unmarshal(m, b)
}
func (m *ScanRunWarningTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanRunWarningTrace.Marshal(b, m, deterministic)
}
func (dst *ScanRunWarningTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanRunWarningTrace.Merge(dst, src)
}
func (m *ScanRunWarningTrace) XXX_Size() int {
	return xxx_messageInfo_ScanRunWarningTrace.Size(m)
}
func (m *ScanRunWarningTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanRunWarningTrace.DiscardUnknown(m)
}

var xxx_messageInfo_ScanRunWarningTrace proto.InternalMessageInfo

func (m *ScanRunWarningTrace) GetCode() ScanRunWarningTrace_Code {
	if m != nil {
		return m.Code
	}
	return ScanRunWarningTrace_CODE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*ScanRunWarningTrace)(nil), "google.cloud.websecurityscanner.v1beta.ScanRunWarningTrace")
	proto.RegisterEnum("google.cloud.websecurityscanner.v1beta.ScanRunWarningTrace_Code", ScanRunWarningTrace_Code_name, ScanRunWarningTrace_Code_value)
}

func init() {
	proto.RegisterFile("google/cloud/websecurityscanner/v1beta/scan_run_warning_trace.proto", fileDescriptor_scan_run_warning_trace_f2bc1463bc284e39)
}

var fileDescriptor_scan_run_warning_trace_f2bc1463bc284e39 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd1, 0x41, 0x4b, 0xf3, 0x30,
	0x18, 0xc0, 0xf1, 0x37, 0xef, 0x3b, 0xde, 0x43, 0x0e, 0xa3, 0x64, 0xa2, 0x63, 0x07, 0x91, 0x1d,
	0x44, 0x3c, 0xa4, 0xa8, 0x47, 0x2f, 0xb6, 0x59, 0x0b, 0x65, 0xb3, 0x2d, 0x4d, 0xcb, 0xdc, 0x2e,
	0x21, 0xcb, 0x42, 0x18, 0xcc, 0x64, 0x64, 0xad, 0xc3, 0x0f, 0xe0, 0x17, 0xf1, 0x2b, 0xfa, 0x05,
	0xa4, 0xad, 0x78, 0x70, 0x03, 0x77, 0x7c, 0xf8, 0xc3, 0x0f, 0x9e, 0xe7, 0x81, 0x44, 0x19, 0xa3,
	0xd6, 0xd2, 0x15, 0x6b, 0x53, 0x2d, 0xdd, 0x9d, 0x5c, 0x6c, 0xa5, 0xa8, 0xec, 0xaa, 0x7c, 0xdd,
	0x0a, 0xae, 0xb5, 0xb4, 0xee, 0xcb, 0xcd, 0x42, 0x96, 0xdc, 0xad, 0x47, 0x66, 0x2b, 0xcd, 0x76,
	0xdc, 0xea, 0x95, 0x56, 0xac, 0xb4, 0x5c, 0x48, 0xbc, 0xb1, 0xa6, 0x34, 0xe8, 0xb2, 0x45, 0x70,
	0x83, 0xe0, 0x7d, 0x04, 0xb7, 0xc8, 0xf0, 0x03, 0xc0, 0x1e, 0x15, 0x5c, 0x67, 0x95, 0x9e, 0xb6,
	0x4c, 0x5e, 0x2b, 0x28, 0x87, 0x1d, 0x61, 0x96, 0xb2, 0x0f, 0x2e, 0xc0, 0x55, 0xf7, 0xf6, 0x01,
	0x1f, 0xc7, 0xe1, 0x03, 0x14, 0x26, 0x66, 0x29, 0xb3, 0x46, 0x1b, 0xbe, 0x01, 0xd8, 0xa9, 0x47,
	0x74, 0x02, 0x1d, 0x92, 0x8c, 0x02, 0x56, 0xc4, 0x34, 0x0d, 0x48, 0x14, 0x46, 0xc1, 0xc8, 0xf9,
	0x83, 0xce, 0xe1, 0x20, 0x8a, 0x69, 0x11, 0x86, 0x11, 0x89, 0x82, 0x38, 0x67, 0x24, 0xf3, 0xa6,
	0x13, 0x96, 0x05, 0xb4, 0x98, 0xe4, 0xd4, 0x01, 0x68, 0x00, 0x4f, 0xf3, 0x24, 0x61, 0x8f, 0x5e,
	0x3c, 0xfb, 0xd1, 0xfe, 0xa2, 0x33, 0xd8, 0xfb, 0x6e, 0x61, 0x31, 0x9f, 0xb3, 0xdc, 0xa3, 0x63,
	0xea, 0xfc, 0x43, 0x08, 0x76, 0xfd, 0x49, 0x42, 0xc6, 0xc1, 0x88, 0xf9, 0x33, 0x16, 0x79, 0xa9,
	0xd3, 0xf1, 0xdf, 0x01, 0xbc, 0x16, 0xe6, 0xf9, 0xc8, 0xad, 0xfc, 0xfe, 0x81, 0xb5, 0xd2, 0xfa,
	0xcc, 0x29, 0x98, 0x3f, 0x7d, 0x19, 0xca, 0xac, 0xb9, 0x56, 0xd8, 0x58, 0xe5, 0x2a, 0xa9, 0x9b,
	0x27, 0xb8, 0x6d, 0xe2, 0x9b, 0xd5, 0xf6, 0xb7, 0x67, 0xde, 0xef, 0x97, 0xc5, 0xff, 0x06, 0xb9,
	0xfb, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x13, 0x9b, 0x58, 0x02, 0x10, 0x02, 0x00, 0x00,
}
