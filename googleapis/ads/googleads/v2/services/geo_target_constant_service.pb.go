// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/geo_target_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [GeoTargetConstantService.GetGeoTargetConstant][google.ads.googleads.v2.services.GeoTargetConstantService.GetGeoTargetConstant].
type GetGeoTargetConstantRequest struct {
	// The resource name of the geo target constant to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGeoTargetConstantRequest) Reset()         { *m = GetGeoTargetConstantRequest{} }
func (m *GetGeoTargetConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetGeoTargetConstantRequest) ProtoMessage()    {}
func (*GetGeoTargetConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0b34171a7a67e99, []int{0}
}

func (m *GetGeoTargetConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGeoTargetConstantRequest.Unmarshal(m, b)
}
func (m *GetGeoTargetConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGeoTargetConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetGeoTargetConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGeoTargetConstantRequest.Merge(m, src)
}
func (m *GetGeoTargetConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetGeoTargetConstantRequest.Size(m)
}
func (m *GetGeoTargetConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGeoTargetConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGeoTargetConstantRequest proto.InternalMessageInfo

func (m *GetGeoTargetConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [GeoTargetConstantService.SuggestGeoTargetConstantsRequest][].
type SuggestGeoTargetConstantsRequest struct {
	// If possible, returned geo targets are translated using this locale. If not,
	// en is used by default. This is also used as a hint for returned geo
	// targets.
	Locale *wrappers.StringValue `protobuf:"bytes,3,opt,name=locale,proto3" json:"locale,omitempty"`
	// Returned geo targets are restricted to this country code.
	CountryCode *wrappers.StringValue `protobuf:"bytes,5,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	// Required. A selector of geo target constants.
	//
	// Types that are valid to be assigned to Query:
	//	*SuggestGeoTargetConstantsRequest_LocationNames_
	//	*SuggestGeoTargetConstantsRequest_GeoTargets_
	Query                isSuggestGeoTargetConstantsRequest_Query `protobuf_oneof:"query"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *SuggestGeoTargetConstantsRequest) Reset()         { *m = SuggestGeoTargetConstantsRequest{} }
func (m *SuggestGeoTargetConstantsRequest) String() string { return proto.CompactTextString(m) }
func (*SuggestGeoTargetConstantsRequest) ProtoMessage()    {}
func (*SuggestGeoTargetConstantsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0b34171a7a67e99, []int{1}
}

func (m *SuggestGeoTargetConstantsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuggestGeoTargetConstantsRequest.Unmarshal(m, b)
}
func (m *SuggestGeoTargetConstantsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuggestGeoTargetConstantsRequest.Marshal(b, m, deterministic)
}
func (m *SuggestGeoTargetConstantsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuggestGeoTargetConstantsRequest.Merge(m, src)
}
func (m *SuggestGeoTargetConstantsRequest) XXX_Size() int {
	return xxx_messageInfo_SuggestGeoTargetConstantsRequest.Size(m)
}
func (m *SuggestGeoTargetConstantsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SuggestGeoTargetConstantsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SuggestGeoTargetConstantsRequest proto.InternalMessageInfo

func (m *SuggestGeoTargetConstantsRequest) GetLocale() *wrappers.StringValue {
	if m != nil {
		return m.Locale
	}
	return nil
}

func (m *SuggestGeoTargetConstantsRequest) GetCountryCode() *wrappers.StringValue {
	if m != nil {
		return m.CountryCode
	}
	return nil
}

type isSuggestGeoTargetConstantsRequest_Query interface {
	isSuggestGeoTargetConstantsRequest_Query()
}

type SuggestGeoTargetConstantsRequest_LocationNames_ struct {
	LocationNames *SuggestGeoTargetConstantsRequest_LocationNames `protobuf:"bytes,1,opt,name=location_names,json=locationNames,proto3,oneof"`
}

type SuggestGeoTargetConstantsRequest_GeoTargets_ struct {
	GeoTargets *SuggestGeoTargetConstantsRequest_GeoTargets `protobuf:"bytes,2,opt,name=geo_targets,json=geoTargets,proto3,oneof"`
}

func (*SuggestGeoTargetConstantsRequest_LocationNames_) isSuggestGeoTargetConstantsRequest_Query() {}

func (*SuggestGeoTargetConstantsRequest_GeoTargets_) isSuggestGeoTargetConstantsRequest_Query() {}

func (m *SuggestGeoTargetConstantsRequest) GetQuery() isSuggestGeoTargetConstantsRequest_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *SuggestGeoTargetConstantsRequest) GetLocationNames() *SuggestGeoTargetConstantsRequest_LocationNames {
	if x, ok := m.GetQuery().(*SuggestGeoTargetConstantsRequest_LocationNames_); ok {
		return x.LocationNames
	}
	return nil
}

func (m *SuggestGeoTargetConstantsRequest) GetGeoTargets() *SuggestGeoTargetConstantsRequest_GeoTargets {
	if x, ok := m.GetQuery().(*SuggestGeoTargetConstantsRequest_GeoTargets_); ok {
		return x.GeoTargets
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SuggestGeoTargetConstantsRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SuggestGeoTargetConstantsRequest_LocationNames_)(nil),
		(*SuggestGeoTargetConstantsRequest_GeoTargets_)(nil),
	}
}

// A list of location names.
type SuggestGeoTargetConstantsRequest_LocationNames struct {
	// A list of location names.
	Names                []*wrappers.StringValue `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SuggestGeoTargetConstantsRequest_LocationNames) Reset() {
	*m = SuggestGeoTargetConstantsRequest_LocationNames{}
}
func (m *SuggestGeoTargetConstantsRequest_LocationNames) String() string {
	return proto.CompactTextString(m)
}
func (*SuggestGeoTargetConstantsRequest_LocationNames) ProtoMessage() {}
func (*SuggestGeoTargetConstantsRequest_LocationNames) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0b34171a7a67e99, []int{1, 0}
}

func (m *SuggestGeoTargetConstantsRequest_LocationNames) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuggestGeoTargetConstantsRequest_LocationNames.Unmarshal(m, b)
}
func (m *SuggestGeoTargetConstantsRequest_LocationNames) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuggestGeoTargetConstantsRequest_LocationNames.Marshal(b, m, deterministic)
}
func (m *SuggestGeoTargetConstantsRequest_LocationNames) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuggestGeoTargetConstantsRequest_LocationNames.Merge(m, src)
}
func (m *SuggestGeoTargetConstantsRequest_LocationNames) XXX_Size() int {
	return xxx_messageInfo_SuggestGeoTargetConstantsRequest_LocationNames.Size(m)
}
func (m *SuggestGeoTargetConstantsRequest_LocationNames) XXX_DiscardUnknown() {
	xxx_messageInfo_SuggestGeoTargetConstantsRequest_LocationNames.DiscardUnknown(m)
}

var xxx_messageInfo_SuggestGeoTargetConstantsRequest_LocationNames proto.InternalMessageInfo

func (m *SuggestGeoTargetConstantsRequest_LocationNames) GetNames() []*wrappers.StringValue {
	if m != nil {
		return m.Names
	}
	return nil
}

// A list of geo target constant resource names.
type SuggestGeoTargetConstantsRequest_GeoTargets struct {
	// A list of geo target constant resource names.
	GeoTargetConstants   []*wrappers.StringValue `protobuf:"bytes,1,rep,name=geo_target_constants,json=geoTargetConstants,proto3" json:"geo_target_constants,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SuggestGeoTargetConstantsRequest_GeoTargets) Reset() {
	*m = SuggestGeoTargetConstantsRequest_GeoTargets{}
}
func (m *SuggestGeoTargetConstantsRequest_GeoTargets) String() string {
	return proto.CompactTextString(m)
}
func (*SuggestGeoTargetConstantsRequest_GeoTargets) ProtoMessage() {}
func (*SuggestGeoTargetConstantsRequest_GeoTargets) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0b34171a7a67e99, []int{1, 1}
}

func (m *SuggestGeoTargetConstantsRequest_GeoTargets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuggestGeoTargetConstantsRequest_GeoTargets.Unmarshal(m, b)
}
func (m *SuggestGeoTargetConstantsRequest_GeoTargets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuggestGeoTargetConstantsRequest_GeoTargets.Marshal(b, m, deterministic)
}
func (m *SuggestGeoTargetConstantsRequest_GeoTargets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuggestGeoTargetConstantsRequest_GeoTargets.Merge(m, src)
}
func (m *SuggestGeoTargetConstantsRequest_GeoTargets) XXX_Size() int {
	return xxx_messageInfo_SuggestGeoTargetConstantsRequest_GeoTargets.Size(m)
}
func (m *SuggestGeoTargetConstantsRequest_GeoTargets) XXX_DiscardUnknown() {
	xxx_messageInfo_SuggestGeoTargetConstantsRequest_GeoTargets.DiscardUnknown(m)
}

var xxx_messageInfo_SuggestGeoTargetConstantsRequest_GeoTargets proto.InternalMessageInfo

func (m *SuggestGeoTargetConstantsRequest_GeoTargets) GetGeoTargetConstants() []*wrappers.StringValue {
	if m != nil {
		return m.GeoTargetConstants
	}
	return nil
}

// Response message for [GeoTargetConstantService.SuggestGeoTargetConstants][google.ads.googleads.v2.services.GeoTargetConstantService.SuggestGeoTargetConstants]
type SuggestGeoTargetConstantsResponse struct {
	// Geo target constant suggestions.
	GeoTargetConstantSuggestions []*GeoTargetConstantSuggestion `protobuf:"bytes,1,rep,name=geo_target_constant_suggestions,json=geoTargetConstantSuggestions,proto3" json:"geo_target_constant_suggestions,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                       `json:"-"`
	XXX_unrecognized             []byte                         `json:"-"`
	XXX_sizecache                int32                          `json:"-"`
}

func (m *SuggestGeoTargetConstantsResponse) Reset()         { *m = SuggestGeoTargetConstantsResponse{} }
func (m *SuggestGeoTargetConstantsResponse) String() string { return proto.CompactTextString(m) }
func (*SuggestGeoTargetConstantsResponse) ProtoMessage()    {}
func (*SuggestGeoTargetConstantsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0b34171a7a67e99, []int{2}
}

func (m *SuggestGeoTargetConstantsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuggestGeoTargetConstantsResponse.Unmarshal(m, b)
}
func (m *SuggestGeoTargetConstantsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuggestGeoTargetConstantsResponse.Marshal(b, m, deterministic)
}
func (m *SuggestGeoTargetConstantsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuggestGeoTargetConstantsResponse.Merge(m, src)
}
func (m *SuggestGeoTargetConstantsResponse) XXX_Size() int {
	return xxx_messageInfo_SuggestGeoTargetConstantsResponse.Size(m)
}
func (m *SuggestGeoTargetConstantsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SuggestGeoTargetConstantsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SuggestGeoTargetConstantsResponse proto.InternalMessageInfo

func (m *SuggestGeoTargetConstantsResponse) GetGeoTargetConstantSuggestions() []*GeoTargetConstantSuggestion {
	if m != nil {
		return m.GeoTargetConstantSuggestions
	}
	return nil
}

// A geo target constant suggestion.
type GeoTargetConstantSuggestion struct {
	// The language this GeoTargetConstantSuggestion is currently translated to.
	// It affects the name of geo target fields. For example, if locale=en, then
	// name=Spain. If locale=es, then name=España. The default locale will be
	// returned if no translation exists for the locale in the request.
	Locale *wrappers.StringValue `protobuf:"bytes,1,opt,name=locale,proto3" json:"locale,omitempty"`
	// Approximate user population that will be targeted, rounded to the
	// nearest 100.
	Reach *wrappers.Int64Value `protobuf:"bytes,2,opt,name=reach,proto3" json:"reach,omitempty"`
	// If the request searched by location name, this is the location name that
	// matched the geo target.
	SearchTerm *wrappers.StringValue `protobuf:"bytes,3,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	// The GeoTargetConstant result.
	GeoTargetConstant *resources.GeoTargetConstant `protobuf:"bytes,4,opt,name=geo_target_constant,json=geoTargetConstant,proto3" json:"geo_target_constant,omitempty"`
	// The list of parents of the geo target constant.
	GeoTargetConstantParents []*resources.GeoTargetConstant `protobuf:"bytes,5,rep,name=geo_target_constant_parents,json=geoTargetConstantParents,proto3" json:"geo_target_constant_parents,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                       `json:"-"`
	XXX_unrecognized         []byte                         `json:"-"`
	XXX_sizecache            int32                          `json:"-"`
}

func (m *GeoTargetConstantSuggestion) Reset()         { *m = GeoTargetConstantSuggestion{} }
func (m *GeoTargetConstantSuggestion) String() string { return proto.CompactTextString(m) }
func (*GeoTargetConstantSuggestion) ProtoMessage()    {}
func (*GeoTargetConstantSuggestion) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0b34171a7a67e99, []int{3}
}

func (m *GeoTargetConstantSuggestion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeoTargetConstantSuggestion.Unmarshal(m, b)
}
func (m *GeoTargetConstantSuggestion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeoTargetConstantSuggestion.Marshal(b, m, deterministic)
}
func (m *GeoTargetConstantSuggestion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoTargetConstantSuggestion.Merge(m, src)
}
func (m *GeoTargetConstantSuggestion) XXX_Size() int {
	return xxx_messageInfo_GeoTargetConstantSuggestion.Size(m)
}
func (m *GeoTargetConstantSuggestion) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoTargetConstantSuggestion.DiscardUnknown(m)
}

var xxx_messageInfo_GeoTargetConstantSuggestion proto.InternalMessageInfo

func (m *GeoTargetConstantSuggestion) GetLocale() *wrappers.StringValue {
	if m != nil {
		return m.Locale
	}
	return nil
}

func (m *GeoTargetConstantSuggestion) GetReach() *wrappers.Int64Value {
	if m != nil {
		return m.Reach
	}
	return nil
}

func (m *GeoTargetConstantSuggestion) GetSearchTerm() *wrappers.StringValue {
	if m != nil {
		return m.SearchTerm
	}
	return nil
}

func (m *GeoTargetConstantSuggestion) GetGeoTargetConstant() *resources.GeoTargetConstant {
	if m != nil {
		return m.GeoTargetConstant
	}
	return nil
}

func (m *GeoTargetConstantSuggestion) GetGeoTargetConstantParents() []*resources.GeoTargetConstant {
	if m != nil {
		return m.GeoTargetConstantParents
	}
	return nil
}

func init() {
	proto.RegisterType((*GetGeoTargetConstantRequest)(nil), "google.ads.googleads.v2.services.GetGeoTargetConstantRequest")
	proto.RegisterType((*SuggestGeoTargetConstantsRequest)(nil), "google.ads.googleads.v2.services.SuggestGeoTargetConstantsRequest")
	proto.RegisterType((*SuggestGeoTargetConstantsRequest_LocationNames)(nil), "google.ads.googleads.v2.services.SuggestGeoTargetConstantsRequest.LocationNames")
	proto.RegisterType((*SuggestGeoTargetConstantsRequest_GeoTargets)(nil), "google.ads.googleads.v2.services.SuggestGeoTargetConstantsRequest.GeoTargets")
	proto.RegisterType((*SuggestGeoTargetConstantsResponse)(nil), "google.ads.googleads.v2.services.SuggestGeoTargetConstantsResponse")
	proto.RegisterType((*GeoTargetConstantSuggestion)(nil), "google.ads.googleads.v2.services.GeoTargetConstantSuggestion")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/geo_target_constant_service.proto", fileDescriptor_b0b34171a7a67e99)
}

var fileDescriptor_b0b34171a7a67e99 = []byte{
	// 730 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xbf, 0x6e, 0xd3, 0x40,
	0x18, 0xc7, 0x49, 0x53, 0xc4, 0xa5, 0x45, 0xe2, 0xa8, 0x84, 0x71, 0x4a, 0x09, 0x6e, 0x87, 0xd2,
	0xe1, 0x0c, 0xa6, 0x62, 0x70, 0x55, 0xa1, 0x24, 0x43, 0x8a, 0x04, 0x55, 0x94, 0x56, 0x19, 0x50,
	0xa4, 0xe8, 0x6a, 0x1f, 0xd7, 0x48, 0x8e, 0xcf, 0xbd, 0xb3, 0x8b, 0x2a, 0xc4, 0x82, 0xfa, 0x06,
	0xbc, 0x00, 0x62, 0xec, 0xc8, 0xce, 0x03, 0xd0, 0x85, 0x81, 0x57, 0x60, 0xe2, 0x29, 0x90, 0xed,
	0xb3, 0xd3, 0xca, 0x31, 0x09, 0x2d, 0xdb, 0xe7, 0xfb, 0xbe, 0xef, 0xf7, 0xfb, 0xfe, 0x1b, 0x34,
	0x29, 0x63, 0xd4, 0x25, 0x06, 0x76, 0x84, 0x91, 0x88, 0x91, 0x74, 0x6c, 0x1a, 0x82, 0xf0, 0xe3,
	0xa1, 0x4d, 0x84, 0x41, 0x09, 0x1b, 0x04, 0x98, 0x53, 0x12, 0x0c, 0x6c, 0xe6, 0x89, 0x00, 0x7b,
	0xc1, 0x40, 0x2a, 0x91, 0xcf, 0x59, 0xc0, 0x60, 0x3d, 0x71, 0x44, 0xd8, 0x11, 0x28, 0xc3, 0x40,
	0xc7, 0x26, 0x4a, 0x31, 0xb4, 0xad, 0x22, 0x16, 0x4e, 0x04, 0x0b, 0x79, 0x01, 0x4d, 0x02, 0xaf,
	0x2d, 0xa7, 0xce, 0xfe, 0xd0, 0xc0, 0x9e, 0xc7, 0x02, 0x1c, 0x0c, 0x99, 0x27, 0xa4, 0x76, 0x45,
	0x6a, 0xe3, 0xaf, 0x83, 0xf0, 0xad, 0xf1, 0x8e, 0x63, 0xdf, 0x27, 0x3c, 0xd5, 0xdf, 0xbb, 0xe0,
	0x6d, 0xbb, 0x43, 0x92, 0xc2, 0xea, 0x4d, 0x50, 0x6b, 0x93, 0xa0, 0x4d, 0xd8, 0x7e, 0xcc, 0xda,
	0x92, 0xa4, 0x5d, 0x72, 0x14, 0x12, 0x11, 0xc0, 0x55, 0xb0, 0x98, 0x06, 0x37, 0xf0, 0xf0, 0x88,
	0xa8, 0x4a, 0x5d, 0x59, 0xbf, 0xd5, 0x5d, 0x48, 0x1f, 0x77, 0xf1, 0x88, 0xe8, 0x5f, 0xe7, 0x40,
	0x7d, 0x2f, 0xa4, 0x94, 0x88, 0x3c, 0x90, 0x48, 0x91, 0x36, 0xc1, 0xbc, 0xcb, 0x6c, 0xec, 0x12,
	0xb5, 0x5c, 0x57, 0xd6, 0xab, 0xe6, 0xb2, 0x2c, 0x12, 0x4a, 0x43, 0x46, 0x7b, 0x01, 0x1f, 0x7a,
	0xb4, 0x87, 0xdd, 0x90, 0x74, 0xa5, 0x2d, 0x7c, 0x01, 0x16, 0x6c, 0x16, 0x7a, 0x01, 0x3f, 0x19,
	0xd8, 0xcc, 0x21, 0x6a, 0x65, 0x06, 0xdf, 0xaa, 0xf4, 0x68, 0x31, 0x87, 0xc0, 0x13, 0x70, 0x3b,
	0x82, 0x8a, 0x6a, 0x15, 0x27, 0x20, 0xe2, 0x0c, 0xaa, 0x66, 0x07, 0x4d, 0x6b, 0x17, 0x9a, 0x96,
	0x12, 0x7a, 0x25, 0x81, 0xa3, 0x22, 0x88, 0x9d, 0x1b, 0xdd, 0x45, 0xf7, 0xe2, 0x03, 0xf4, 0x41,
	0x75, 0xdc, 0x4e, 0xa1, 0x96, 0x62, 0xde, 0xd7, 0xff, 0x81, 0x37, 0xd3, 0x44, 0xa4, 0x80, 0x66,
	0x5f, 0x5a, 0x0b, 0x2c, 0x5e, 0x8a, 0x09, 0x9a, 0xa0, 0x92, 0x26, 0x5d, 0x9e, 0x5a, 0xb7, 0xc4,
	0x54, 0xeb, 0x03, 0x30, 0x26, 0x80, 0xbb, 0x60, 0x69, 0xc2, 0x4c, 0xce, 0x06, 0x08, 0x69, 0x2e,
	0x85, 0xe6, 0x4d, 0x50, 0x39, 0x0a, 0x09, 0x3f, 0xd1, 0xcf, 0x14, 0xf0, 0xe8, 0x2f, 0x99, 0x0a,
	0x9f, 0x79, 0x82, 0xc0, 0x53, 0x05, 0x3c, 0x9c, 0xb8, 0x7a, 0x89, 0x67, 0xb4, 0x01, 0x32, 0x94,
	0xed, 0xe9, 0x85, 0xcd, 0xf1, 0xec, 0x65, 0x28, 0xdd, 0x65, 0x5a, 0xac, 0x14, 0xfa, 0xe7, 0x72,
	0xb4, 0x26, 0x85, 0x06, 0x17, 0x86, 0x5b, 0xf9, 0x87, 0xe1, 0x7e, 0x0a, 0x2a, 0x9c, 0x60, 0xfb,
	0x50, 0x8e, 0x46, 0x2d, 0xe7, 0xf4, 0xd2, 0x0b, 0x9e, 0x6f, 0xca, 0xe6, 0xc4, 0x96, 0x70, 0x1b,
	0x54, 0x05, 0xc1, 0xdc, 0x3e, 0x1c, 0x04, 0x84, 0x8f, 0x66, 0x5a, 0x25, 0x90, 0x38, 0xec, 0x13,
	0x3e, 0x82, 0x0e, 0xb8, 0x3b, 0xa1, 0x9a, 0xea, 0x5c, 0x0c, 0xb3, 0x59, 0x58, 0xc1, 0xec, 0x3e,
	0xe5, 0x4b, 0xd8, 0xbd, 0x93, 0x2b, 0x1c, 0x14, 0xa0, 0x36, 0xa9, 0x67, 0x3e, 0xe6, 0x24, 0x1a,
	0x9d, 0x4a, 0xdc, 0xaf, 0xab, 0xb1, 0xa9, 0x39, 0xb6, 0x4e, 0x82, 0x6a, 0x7e, 0x2f, 0x03, 0x35,
	0xdf, 0xa2, 0xa4, 0xf5, 0xf0, 0x9b, 0x02, 0x96, 0x26, 0x9d, 0x39, 0x38, 0xd3, 0xd4, 0x14, 0x9e,
	0x47, 0xed, 0x4a, 0x49, 0xe8, 0x4f, 0x3e, 0xfe, 0xfc, 0xf5, 0xa9, 0xb4, 0x01, 0xd7, 0xa3, 0xdb,
	0xff, 0xfe, 0xd2, 0x7d, 0xdd, 0xce, 0xaf, 0x8c, 0xb1, 0xf1, 0x01, 0xfe, 0x50, 0xc0, 0xfd, 0xc2,
	0x65, 0x81, 0xcd, 0xeb, 0xdf, 0x14, 0xad, 0x75, 0x2d, 0x8c, 0x64, 0x5b, 0xf5, 0xc7, 0x71, 0x62,
	0xab, 0xfa, 0x4a, 0x94, 0x58, 0x3e, 0x13, 0x4b, 0x6e, 0xad, 0xa5, 0x6c, 0x68, 0xb5, 0xf3, 0x86,
	0x3a, 0xa6, 0x91, 0x92, 0x3f, 0x14, 0xc8, 0x66, 0xa3, 0xe6, 0x69, 0x09, 0xac, 0xd9, 0x6c, 0x34,
	0x35, 0xa4, 0xe6, 0x83, 0xa2, 0x8e, 0x77, 0xa2, 0x4d, 0xe8, 0x28, 0x6f, 0x76, 0x24, 0x04, 0x65,
	0x2e, 0xf6, 0x28, 0x62, 0x9c, 0x1a, 0x94, 0x78, 0xf1, 0x9e, 0x18, 0x63, 0xd2, 0xe2, 0xff, 0xfe,
	0x56, 0x2a, 0x7c, 0x29, 0x95, 0xdb, 0x8d, 0xc6, 0x59, 0xa9, 0xde, 0x4e, 0x00, 0x1b, 0x8e, 0x40,
	0x89, 0x18, 0x49, 0x3d, 0x13, 0x49, 0x62, 0x71, 0x9e, 0x9a, 0xf4, 0x1b, 0x8e, 0xe8, 0x67, 0x26,
	0xfd, 0x9e, 0xd9, 0x4f, 0x4d, 0x7e, 0x97, 0xd6, 0x92, 0x77, 0xcb, 0x6a, 0x38, 0xc2, 0xb2, 0x32,
	0x23, 0xcb, 0xea, 0x99, 0x96, 0x95, 0x9a, 0x1d, 0xcc, 0xc7, 0x71, 0x3e, 0xfb, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x6c, 0xd8, 0xc1, 0x65, 0x9e, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GeoTargetConstantServiceClient is the client API for GeoTargetConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeoTargetConstantServiceClient interface {
	// Returns the requested geo target constant in full detail.
	GetGeoTargetConstant(ctx context.Context, in *GetGeoTargetConstantRequest, opts ...grpc.CallOption) (*resources.GeoTargetConstant, error)
	// Returns GeoTargetConstant suggestions by location name or by resource name.
	SuggestGeoTargetConstants(ctx context.Context, in *SuggestGeoTargetConstantsRequest, opts ...grpc.CallOption) (*SuggestGeoTargetConstantsResponse, error)
}

type geoTargetConstantServiceClient struct {
	cc *grpc.ClientConn
}

func NewGeoTargetConstantServiceClient(cc *grpc.ClientConn) GeoTargetConstantServiceClient {
	return &geoTargetConstantServiceClient{cc}
}

func (c *geoTargetConstantServiceClient) GetGeoTargetConstant(ctx context.Context, in *GetGeoTargetConstantRequest, opts ...grpc.CallOption) (*resources.GeoTargetConstant, error) {
	out := new(resources.GeoTargetConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.GeoTargetConstantService/GetGeoTargetConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoTargetConstantServiceClient) SuggestGeoTargetConstants(ctx context.Context, in *SuggestGeoTargetConstantsRequest, opts ...grpc.CallOption) (*SuggestGeoTargetConstantsResponse, error) {
	out := new(SuggestGeoTargetConstantsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.GeoTargetConstantService/SuggestGeoTargetConstants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeoTargetConstantServiceServer is the server API for GeoTargetConstantService service.
type GeoTargetConstantServiceServer interface {
	// Returns the requested geo target constant in full detail.
	GetGeoTargetConstant(context.Context, *GetGeoTargetConstantRequest) (*resources.GeoTargetConstant, error)
	// Returns GeoTargetConstant suggestions by location name or by resource name.
	SuggestGeoTargetConstants(context.Context, *SuggestGeoTargetConstantsRequest) (*SuggestGeoTargetConstantsResponse, error)
}

// UnimplementedGeoTargetConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGeoTargetConstantServiceServer struct {
}

func (*UnimplementedGeoTargetConstantServiceServer) GetGeoTargetConstant(ctx context.Context, req *GetGeoTargetConstantRequest) (*resources.GeoTargetConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGeoTargetConstant not implemented")
}
func (*UnimplementedGeoTargetConstantServiceServer) SuggestGeoTargetConstants(ctx context.Context, req *SuggestGeoTargetConstantsRequest) (*SuggestGeoTargetConstantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuggestGeoTargetConstants not implemented")
}

func RegisterGeoTargetConstantServiceServer(s *grpc.Server, srv GeoTargetConstantServiceServer) {
	s.RegisterService(&_GeoTargetConstantService_serviceDesc, srv)
}

func _GeoTargetConstantService_GetGeoTargetConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGeoTargetConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoTargetConstantServiceServer).GetGeoTargetConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.GeoTargetConstantService/GetGeoTargetConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoTargetConstantServiceServer).GetGeoTargetConstant(ctx, req.(*GetGeoTargetConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoTargetConstantService_SuggestGeoTargetConstants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestGeoTargetConstantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoTargetConstantServiceServer).SuggestGeoTargetConstants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.GeoTargetConstantService/SuggestGeoTargetConstants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoTargetConstantServiceServer).SuggestGeoTargetConstants(ctx, req.(*SuggestGeoTargetConstantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GeoTargetConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.GeoTargetConstantService",
	HandlerType: (*GeoTargetConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGeoTargetConstant",
			Handler:    _GeoTargetConstantService_GetGeoTargetConstant_Handler,
		},
		{
			MethodName: "SuggestGeoTargetConstants",
			Handler:    _GeoTargetConstantService_SuggestGeoTargetConstants_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/geo_target_constant_service.proto",
}
