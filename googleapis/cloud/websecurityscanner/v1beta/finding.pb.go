// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/websecurityscanner/v1beta/finding.proto

package websecurityscanner // import "google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1beta"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

// A Finding resource represents a vulnerability instance identified during a
// ScanRun.
type Finding struct {
	// Output only.
	// The resource name of the Finding. The name follows the format of
	// 'projects/{projectId}/scanConfigs/{scanConfigId}/scanruns/{scanRunId}/findings/{findingId}'.
	// The finding IDs are generated by the system.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only.
	// The type of the Finding.
	// Detailed and up-to-date information on findings can be found here:
	// https://cloud.google.com/security-scanner/docs/scan-result-details
	FindingType string `protobuf:"bytes,2,opt,name=finding_type,json=findingType,proto3" json:"finding_type,omitempty"`
	// Output only.
	// The http method of the request that triggered the vulnerability, in
	// uppercase.
	HttpMethod string `protobuf:"bytes,3,opt,name=http_method,json=httpMethod,proto3" json:"http_method,omitempty"`
	// Output only.
	// The URL produced by the server-side fuzzer and used in the request that
	// triggered the vulnerability.
	FuzzedUrl string `protobuf:"bytes,4,opt,name=fuzzed_url,json=fuzzedUrl,proto3" json:"fuzzed_url,omitempty"`
	// Output only.
	// The body of the request that triggered the vulnerability.
	Body string `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	// Output only.
	// The description of the vulnerability.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// Output only.
	// The URL containing human-readable payload that user can leverage to
	// reproduce the vulnerability.
	ReproductionUrl string `protobuf:"bytes,7,opt,name=reproduction_url,json=reproductionUrl,proto3" json:"reproduction_url,omitempty"`
	// Output only.
	// If the vulnerability was originated from nested IFrame, the immediate
	// parent IFrame is reported.
	FrameUrl string `protobuf:"bytes,8,opt,name=frame_url,json=frameUrl,proto3" json:"frame_url,omitempty"`
	// Output only.
	// The URL where the browser lands when the vulnerability is detected.
	FinalUrl string `protobuf:"bytes,9,opt,name=final_url,json=finalUrl,proto3" json:"final_url,omitempty"`
	// Output only.
	// The tracking ID uniquely identifies a vulnerability instance across
	// multiple ScanRuns.
	TrackingId string `protobuf:"bytes,10,opt,name=tracking_id,json=trackingId,proto3" json:"tracking_id,omitempty"`
	// Output only.
	// An addon containing information reported for a vulnerability with an HTML
	// form, if any.
	Form *Form `protobuf:"bytes,16,opt,name=form,proto3" json:"form,omitempty"`
	// Output only.
	// An addon containing information about outdated libraries.
	OutdatedLibrary *OutdatedLibrary `protobuf:"bytes,11,opt,name=outdated_library,json=outdatedLibrary,proto3" json:"outdated_library,omitempty"`
	// Output only.
	// An addon containing detailed information regarding any resource causing the
	// vulnerability such as JavaScript sources, image, audio files, etc.
	ViolatingResource *ViolatingResource `protobuf:"bytes,12,opt,name=violating_resource,json=violatingResource,proto3" json:"violating_resource,omitempty"`
	// Output only.
	// An addon containing information about vulnerable or missing HTTP headers.
	VulnerableHeaders *VulnerableHeaders `protobuf:"bytes,15,opt,name=vulnerable_headers,json=vulnerableHeaders,proto3" json:"vulnerable_headers,omitempty"`
	// Output only.
	// An addon containing information about request parameters which were found
	// to be vulnerable.
	VulnerableParameters *VulnerableParameters `protobuf:"bytes,13,opt,name=vulnerable_parameters,json=vulnerableParameters,proto3" json:"vulnerable_parameters,omitempty"`
	// Output only.
	// An addon containing information reported for an XSS, if any.
	Xss                  *Xss     `protobuf:"bytes,14,opt,name=xss,proto3" json:"xss,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Finding) Reset()         { *m = Finding{} }
func (m *Finding) String() string { return proto.CompactTextString(m) }
func (*Finding) ProtoMessage()    {}
func (*Finding) Descriptor() ([]byte, []int) {
	return fileDescriptor_finding_b0f9e4c54638c23f, []int{0}
}
func (m *Finding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Finding.Unmarshal(m, b)
}
func (m *Finding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Finding.Marshal(b, m, deterministic)
}
func (dst *Finding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Finding.Merge(dst, src)
}
func (m *Finding) XXX_Size() int {
	return xxx_messageInfo_Finding.Size(m)
}
func (m *Finding) XXX_DiscardUnknown() {
	xxx_messageInfo_Finding.DiscardUnknown(m)
}

var xxx_messageInfo_Finding proto.InternalMessageInfo

func (m *Finding) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Finding) GetFindingType() string {
	if m != nil {
		return m.FindingType
	}
	return ""
}

func (m *Finding) GetHttpMethod() string {
	if m != nil {
		return m.HttpMethod
	}
	return ""
}

func (m *Finding) GetFuzzedUrl() string {
	if m != nil {
		return m.FuzzedUrl
	}
	return ""
}

func (m *Finding) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Finding) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Finding) GetReproductionUrl() string {
	if m != nil {
		return m.ReproductionUrl
	}
	return ""
}

func (m *Finding) GetFrameUrl() string {
	if m != nil {
		return m.FrameUrl
	}
	return ""
}

func (m *Finding) GetFinalUrl() string {
	if m != nil {
		return m.FinalUrl
	}
	return ""
}

func (m *Finding) GetTrackingId() string {
	if m != nil {
		return m.TrackingId
	}
	return ""
}

func (m *Finding) GetForm() *Form {
	if m != nil {
		return m.Form
	}
	return nil
}

func (m *Finding) GetOutdatedLibrary() *OutdatedLibrary {
	if m != nil {
		return m.OutdatedLibrary
	}
	return nil
}

func (m *Finding) GetViolatingResource() *ViolatingResource {
	if m != nil {
		return m.ViolatingResource
	}
	return nil
}

func (m *Finding) GetVulnerableHeaders() *VulnerableHeaders {
	if m != nil {
		return m.VulnerableHeaders
	}
	return nil
}

func (m *Finding) GetVulnerableParameters() *VulnerableParameters {
	if m != nil {
		return m.VulnerableParameters
	}
	return nil
}

func (m *Finding) GetXss() *Xss {
	if m != nil {
		return m.Xss
	}
	return nil
}

func init() {
	proto.RegisterType((*Finding)(nil), "google.cloud.websecurityscanner.v1beta.Finding")
}

func init() {
	proto.RegisterFile("google/cloud/websecurityscanner/v1beta/finding.proto", fileDescriptor_finding_b0f9e4c54638c23f)
}

var fileDescriptor_finding_b0f9e4c54638c23f = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x5f, 0x8b, 0x13, 0x31,
	0x14, 0xc5, 0xa9, 0xbb, 0x76, 0xb7, 0x99, 0x6a, 0xbb, 0x41, 0x21, 0xac, 0xca, 0x56, 0x1f, 0x64,
	0xfd, 0xc3, 0x0c, 0xfe, 0x01, 0xf1, 0x1f, 0xc8, 0x3e, 0x2c, 0x0a, 0x8a, 0x65, 0x50, 0x59, 0x7c,
	0x29, 0x99, 0xc9, 0xed, 0x34, 0x38, 0x93, 0x8c, 0x49, 0xa6, 0xda, 0xfd, 0x20, 0x7e, 0x46, 0x3f,
	0x86, 0xe4, 0x66, 0xba, 0xd4, 0x5d, 0xc1, 0xf6, 0x2d, 0x73, 0xce, 0x3d, 0xe7, 0xc7, 0x25, 0x64,
	0xc8, 0xd3, 0x42, 0xeb, 0xa2, 0x84, 0x24, 0x2f, 0x75, 0x23, 0x92, 0x1f, 0x90, 0x59, 0xc8, 0x1b,
	0x23, 0xdd, 0xc2, 0xe6, 0x5c, 0x29, 0x30, 0xc9, 0xfc, 0x51, 0x06, 0x8e, 0x27, 0x53, 0xa9, 0x84,
	0x54, 0x45, 0x5c, 0x1b, 0xed, 0x34, 0xbd, 0x1b, 0x52, 0x31, 0xa6, 0xe2, 0x8b, 0xa9, 0x38, 0xa4,
	0xf6, 0x6f, 0xb6, 0xed, 0xbc, 0x96, 0x09, 0x57, 0x4a, 0x3b, 0xee, 0xa4, 0x56, 0x36, 0xb4, 0xec,
	0xbf, 0xd8, 0x8c, 0x3d, 0xe1, 0x42, 0x68, 0x15, 0xb2, 0x77, 0x7e, 0x77, 0xc9, 0xce, 0x71, 0xd0,
	0x29, 0x25, 0xdb, 0x8a, 0x57, 0xc0, 0x3a, 0xa3, 0xce, 0x61, 0x2f, 0xc5, 0x33, 0xbd, 0x4d, 0xfa,
	0xcb, 0x98, 0x5b, 0xd4, 0xc0, 0x2e, 0xa1, 0x17, 0xb5, 0xda, 0xa7, 0x45, 0x0d, 0xf4, 0x80, 0x44,
	0x33, 0xe7, 0xea, 0x49, 0x05, 0x6e, 0xa6, 0x05, 0xdb, 0xc2, 0x09, 0xe2, 0xa5, 0x0f, 0xa8, 0xd0,
	0x5b, 0x84, 0x4c, 0x9b, 0xd3, 0x53, 0x10, 0x93, 0xc6, 0x94, 0x6c, 0x1b, 0xfd, 0x5e, 0x50, 0x3e,
	0x9b, 0xd2, 0x63, 0x33, 0x2d, 0x16, 0xec, 0x72, 0xc0, 0xfa, 0x33, 0x1d, 0x91, 0x48, 0x80, 0xcd,
	0x8d, 0xac, 0xfd, 0xa2, 0xac, 0x1b, 0xa8, 0x2b, 0x12, 0xbd, 0x47, 0x86, 0x06, 0x6a, 0xa3, 0x45,
	0x93, 0xfb, 0x6f, 0xac, 0xde, 0xc1, 0xb1, 0xc1, 0xaa, 0xee, 0x01, 0x37, 0x48, 0x6f, 0x6a, 0x78,
	0x05, 0x38, 0xb3, 0x8b, 0x33, 0xbb, 0x28, 0x2c, 0x4d, 0xa9, 0x78, 0x89, 0x66, 0xaf, 0x35, 0xbd,
	0xe0, 0xcd, 0x03, 0x12, 0x39, 0xc3, 0xf3, 0x6f, 0x7e, 0x7d, 0x29, 0x18, 0x09, 0xab, 0x2d, 0xa5,
	0x77, 0x82, 0xbe, 0x21, 0xdb, 0x53, 0x6d, 0x2a, 0x36, 0x1c, 0x75, 0x0e, 0xa3, 0xc7, 0x0f, 0xe3,
	0xf5, 0xee, 0x33, 0x3e, 0xd6, 0xa6, 0x4a, 0x31, 0x49, 0x33, 0x32, 0xd4, 0x8d, 0x13, 0xdc, 0x81,
	0x98, 0x94, 0x32, 0x33, 0xdc, 0x2c, 0x58, 0x84, 0x6d, 0xcf, 0xd6, 0x6d, 0xfb, 0xd8, 0xe6, 0xdf,
	0x87, 0x78, 0x3a, 0xd0, 0x7f, 0x0b, 0x74, 0x46, 0xe8, 0x5c, 0xea, 0x92, 0x3b, 0xbf, 0x87, 0x01,
	0xab, 0x1b, 0x93, 0x03, 0xeb, 0x23, 0xe5, 0xf9, 0xba, 0x94, 0x2f, 0xcb, 0x86, 0xb4, 0x2d, 0x48,
	0xf7, 0xe6, 0xe7, 0x25, 0x24, 0x35, 0xa5, 0x02, 0xc3, 0xb3, 0x12, 0x26, 0x33, 0xe0, 0x02, 0x8c,
	0x65, 0x83, 0x0d, 0x49, 0x67, 0x0d, 0x6f, 0x43, 0x41, 0xba, 0x37, 0x3f, 0x2f, 0xd1, 0xef, 0xe4,
	0xfa, 0x0a, 0xa9, 0xe6, 0xfe, 0x3e, 0x9d, 0x87, 0x5d, 0x41, 0xd8, 0xab, 0xcd, 0x61, 0xe3, 0xb3,
	0x8e, 0xf4, 0xda, 0xfc, 0x1f, 0x2a, 0x7d, 0x4d, 0xb6, 0x7e, 0x5a, 0xcb, 0xae, 0x22, 0xe0, 0xc1,
	0xba, 0x80, 0x13, 0x6b, 0x53, 0x9f, 0x3b, 0xfa, 0xd5, 0x21, 0xf7, 0x73, 0x5d, 0xad, 0x99, 0x3b,
	0xea, 0xb7, 0xcf, 0x72, 0xec, 0xdf, 0xe9, 0xb8, 0xf3, 0xf5, 0xa4, 0xcd, 0x15, 0xba, 0xe4, 0xaa,
	0x88, 0xb5, 0x29, 0x92, 0x02, 0x14, 0xbe, 0xe2, 0x24, 0x58, 0xbc, 0x96, 0xf6, 0x7f, 0x3f, 0x81,
	0x97, 0x17, 0x9d, 0xac, 0x8b, 0x25, 0x4f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xa4, 0xbf,
	0x31, 0xc4, 0x04, 0x00, 0x00,
}
