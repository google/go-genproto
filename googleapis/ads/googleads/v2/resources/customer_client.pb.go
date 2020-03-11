// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/customer_client.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A link between the given customer and a client customer. CustomerClients only
// exist for manager customers. All direct and indirect client customers are
// included, as well as the manager itself.
type CustomerClient struct {
	// Immutable. The resource name of the customer client.
	// CustomerClient resource names have the form:
	// `customers/{customer_id}/customerClients/{client_customer_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The resource name of the client-customer which is linked to
	// the given customer. Read only.
	ClientCustomer *wrappers.StringValue `protobuf:"bytes,3,opt,name=client_customer,json=clientCustomer,proto3" json:"client_customer,omitempty"`
	// Output only. Specifies whether this is a
	// [hidden account](https://support.google.com/google-ads/answer/7519830).
	// Read only.
	Hidden *wrappers.BoolValue `protobuf:"bytes,4,opt,name=hidden,proto3" json:"hidden,omitempty"`
	// Output only. Distance between given customer and client. For self link, the level value
	// will be 0. Read only.
	Level *wrappers.Int64Value `protobuf:"bytes,5,opt,name=level,proto3" json:"level,omitempty"`
	// Output only. Common Locale Data Repository (CLDR) string representation of the
	// time zone of the client, e.g. America/Los_Angeles. Read only.
	TimeZone *wrappers.StringValue `protobuf:"bytes,6,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	// Output only. Identifies if the client is a test account. Read only.
	TestAccount *wrappers.BoolValue `protobuf:"bytes,7,opt,name=test_account,json=testAccount,proto3" json:"test_account,omitempty"`
	// Output only. Identifies if the client is a manager. Read only.
	Manager *wrappers.BoolValue `protobuf:"bytes,8,opt,name=manager,proto3" json:"manager,omitempty"`
	// Output only. Descriptive name for the client. Read only.
	DescriptiveName *wrappers.StringValue `protobuf:"bytes,9,opt,name=descriptive_name,json=descriptiveName,proto3" json:"descriptive_name,omitempty"`
	// Output only. Currency code (e.g. 'USD', 'EUR') for the client. Read only.
	CurrencyCode *wrappers.StringValue `protobuf:"bytes,10,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// Output only. The ID of the client customer. Read only.
	Id                   *wrappers.Int64Value `protobuf:"bytes,11,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CustomerClient) Reset()         { *m = CustomerClient{} }
func (m *CustomerClient) String() string { return proto.CompactTextString(m) }
func (*CustomerClient) ProtoMessage()    {}
func (*CustomerClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_463452e1d1f19c18, []int{0}
}

func (m *CustomerClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerClient.Unmarshal(m, b)
}
func (m *CustomerClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerClient.Marshal(b, m, deterministic)
}
func (m *CustomerClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerClient.Merge(m, src)
}
func (m *CustomerClient) XXX_Size() int {
	return xxx_messageInfo_CustomerClient.Size(m)
}
func (m *CustomerClient) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerClient.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerClient proto.InternalMessageInfo

func (m *CustomerClient) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomerClient) GetClientCustomer() *wrappers.StringValue {
	if m != nil {
		return m.ClientCustomer
	}
	return nil
}

func (m *CustomerClient) GetHidden() *wrappers.BoolValue {
	if m != nil {
		return m.Hidden
	}
	return nil
}

func (m *CustomerClient) GetLevel() *wrappers.Int64Value {
	if m != nil {
		return m.Level
	}
	return nil
}

func (m *CustomerClient) GetTimeZone() *wrappers.StringValue {
	if m != nil {
		return m.TimeZone
	}
	return nil
}

func (m *CustomerClient) GetTestAccount() *wrappers.BoolValue {
	if m != nil {
		return m.TestAccount
	}
	return nil
}

func (m *CustomerClient) GetManager() *wrappers.BoolValue {
	if m != nil {
		return m.Manager
	}
	return nil
}

func (m *CustomerClient) GetDescriptiveName() *wrappers.StringValue {
	if m != nil {
		return m.DescriptiveName
	}
	return nil
}

func (m *CustomerClient) GetCurrencyCode() *wrappers.StringValue {
	if m != nil {
		return m.CurrencyCode
	}
	return nil
}

func (m *CustomerClient) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*CustomerClient)(nil), "google.ads.googleads.v2.resources.CustomerClient")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/customer_client.proto", fileDescriptor_463452e1d1f19c18)
}

var fileDescriptor_463452e1d1f19c18 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdd, 0x6a, 0xd4, 0x4e,
	0x18, 0xc6, 0xd9, 0xec, 0xbf, 0xdb, 0x76, 0xfa, 0xf5, 0x27, 0x9e, 0xc4, 0xb5, 0x68, 0x2b, 0x14,
	0xeb, 0xc9, 0x44, 0xa2, 0x6c, 0x31, 0x82, 0x90, 0x2c, 0x58, 0x2c, 0x22, 0x65, 0x95, 0x3d, 0x28,
	0x0b, 0x61, 0x36, 0xf3, 0x36, 0x1d, 0x48, 0x66, 0xc2, 0xcc, 0x64, 0x45, 0xa5, 0xe0, 0x25, 0x78,
	0x0d, 0x1e, 0x7a, 0x29, 0x5e, 0x45, 0x8f, 0x7b, 0x09, 0x1e, 0xc9, 0x26, 0x99, 0x34, 0xa5, 0xa0,
	0x39, 0x7b, 0xc3, 0xfb, 0xfc, 0x9e, 0x79, 0xde, 0xf9, 0x08, 0x3a, 0x4a, 0x84, 0x48, 0x52, 0x70,
	0x09, 0x55, 0x6e, 0x55, 0x2e, 0xab, 0x85, 0xe7, 0x4a, 0x50, 0xa2, 0x90, 0x31, 0x28, 0x37, 0x2e,
	0x94, 0x16, 0x19, 0xc8, 0x28, 0x4e, 0x19, 0x70, 0x8d, 0x73, 0x29, 0xb4, 0xb0, 0xf7, 0x2b, 0x35,
	0x26, 0x54, 0xe1, 0x06, 0xc4, 0x0b, 0x0f, 0x37, 0xe0, 0xf0, 0x91, 0xf1, 0xce, 0x99, 0x7b, 0xce,
	0x20, 0xa5, 0xd1, 0x1c, 0x2e, 0xc8, 0x82, 0x09, 0x59, 0x79, 0x0c, 0xef, 0xb7, 0x04, 0x06, 0xab,
	0x5b, 0x0f, 0xeb, 0x56, 0xf9, 0x35, 0x2f, 0xce, 0xdd, 0x4f, 0x92, 0xe4, 0x39, 0x48, 0x55, 0xf7,
	0x77, 0x5b, 0x28, 0xe1, 0x5c, 0x68, 0xa2, 0x99, 0xe0, 0x75, 0xf7, 0xf1, 0xf7, 0x01, 0xda, 0x1e,
	0xd7, 0xb1, 0xc7, 0x65, 0x6a, 0xfb, 0x23, 0xda, 0x32, 0x4b, 0x44, 0x9c, 0x64, 0xe0, 0xf4, 0xf6,
	0x7a, 0x87, 0xeb, 0xa1, 0x7b, 0x15, 0xac, 0xfc, 0x0e, 0x9e, 0xa2, 0x27, 0x37, 0x33, 0xd4, 0x55,
	0xce, 0x14, 0x8e, 0x45, 0xe6, 0xde, 0xf6, 0x99, 0x6c, 0x1a, 0x97, 0xf7, 0x24, 0x03, 0xfb, 0x04,
	0xed, 0x54, 0xbb, 0x12, 0x99, 0x5d, 0x72, 0xfa, 0x7b, 0xbd, 0xc3, 0x0d, 0x6f, 0xb7, 0xb6, 0xc1,
	0x66, 0x00, 0xfc, 0x41, 0x4b, 0xc6, 0x93, 0x29, 0x49, 0x0b, 0x08, 0xfb, 0x57, 0x41, 0x7f, 0xb2,
	0x5d, 0x91, 0xc6, 0xdf, 0x3e, 0x42, 0x83, 0x0b, 0x46, 0x29, 0x70, 0xe7, 0xbf, 0xd2, 0x62, 0x78,
	0xc7, 0x22, 0x14, 0x22, 0x6d, 0x19, 0xd4, 0x72, 0x7b, 0x84, 0x56, 0x52, 0x58, 0x40, 0xea, 0xac,
	0x94, 0xdc, 0x83, 0x3b, 0xdc, 0x5b, 0xae, 0x47, 0x2f, 0x5a, 0x60, 0x25, 0xb7, 0x5f, 0xa3, 0x75,
	0xcd, 0x32, 0x88, 0xbe, 0x08, 0x0e, 0xce, 0xa0, 0x6b, 0xec, 0xb5, 0x25, 0x73, 0x26, 0x38, 0xd8,
	0x21, 0xda, 0xd4, 0xa0, 0x74, 0x44, 0xe2, 0x58, 0x14, 0x5c, 0x3b, 0xab, 0xdd, 0x62, 0x6f, 0x2c,
	0xa1, 0xa0, 0x62, 0xec, 0x97, 0x68, 0x35, 0x23, 0x9c, 0x24, 0x20, 0x9d, 0xb5, 0x6e, 0xb8, 0xd1,
	0xdb, 0xef, 0xd0, 0xff, 0x14, 0x54, 0x2c, 0x59, 0xae, 0xd9, 0xa2, 0x3e, 0xd4, 0xf5, 0xae, 0x53,
	0xec, 0xb4, 0xd0, 0xf2, 0x24, 0xdf, 0xa0, 0xad, 0xb8, 0x90, 0x12, 0x78, 0xfc, 0x39, 0x8a, 0x05,
	0x05, 0x07, 0x75, 0xb5, 0xda, 0x34, 0xdc, 0x58, 0x50, 0xb0, 0x9f, 0x21, 0x8b, 0x51, 0x67, 0xa3,
	0xe3, 0x49, 0x58, 0x8c, 0xfa, 0xf4, 0x3a, 0x20, 0x9d, 0xef, 0x9f, 0x3d, 0x32, 0x17, 0x4d, 0xb9,
	0x5f, 0x4d, 0x79, 0xd9, 0xbc, 0xd1, 0x4a, 0xd4, 0x6a, 0xd5, 0x8f, 0xf6, 0x32, 0xfc, 0x66, 0xa1,
	0x83, 0x58, 0x64, 0xf8, 0x9f, 0xcf, 0x36, 0xbc, 0x77, 0x7b, 0xc5, 0xd3, 0x65, 0xf6, 0xd3, 0xde,
	0xd9, 0x49, 0x4d, 0x26, 0x22, 0x25, 0x3c, 0xc1, 0x42, 0x26, 0x6e, 0x02, 0xbc, 0x9c, 0xcc, 0xbd,
	0x89, 0xfc, 0x97, 0x1f, 0xc9, 0xab, 0xa6, 0xfa, 0x61, 0xf5, 0x8f, 0x83, 0xe0, 0xa7, 0xb5, 0x7f,
	0x5c, 0x59, 0x06, 0x54, 0xe1, 0xaa, 0x5c, 0x56, 0x53, 0x0f, 0x4f, 0x8c, 0xf2, 0x97, 0xd1, 0xcc,
	0x02, 0xaa, 0x66, 0x8d, 0x66, 0x36, 0xf5, 0x66, 0x8d, 0xe6, 0xda, 0x3a, 0xa8, 0x1a, 0xbe, 0x1f,
	0x50, 0xe5, 0xfb, 0x8d, 0xca, 0xf7, 0xa7, 0x9e, 0xef, 0x37, 0xba, 0xf9, 0xa0, 0x0c, 0xfb, 0xfc,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xa7, 0xcd, 0xcb, 0xf4, 0x04, 0x00, 0x00,
}
