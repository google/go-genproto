// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/conversion_upload_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v1/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import status "google.golang.org/genproto/googleapis/rpc/status"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [ConversionUploadService.UploadClickConversions][google.ads.googleads.v1.services.ConversionUploadService.UploadClickConversions].
type UploadClickConversionsRequest struct {
	// The ID of the customer performing the upload.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The conversions that are being uploaded.
	Conversions []*ClickConversion `protobuf:"bytes,2,rep,name=conversions,proto3" json:"conversions,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// This should always be set to true.
	PartialFailure       bool     `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadClickConversionsRequest) Reset()         { *m = UploadClickConversionsRequest{} }
func (m *UploadClickConversionsRequest) String() string { return proto.CompactTextString(m) }
func (*UploadClickConversionsRequest) ProtoMessage()    {}
func (*UploadClickConversionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_upload_service_96e37aa981399bfd, []int{0}
}
func (m *UploadClickConversionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadClickConversionsRequest.Unmarshal(m, b)
}
func (m *UploadClickConversionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadClickConversionsRequest.Marshal(b, m, deterministic)
}
func (dst *UploadClickConversionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadClickConversionsRequest.Merge(dst, src)
}
func (m *UploadClickConversionsRequest) XXX_Size() int {
	return xxx_messageInfo_UploadClickConversionsRequest.Size(m)
}
func (m *UploadClickConversionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadClickConversionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadClickConversionsRequest proto.InternalMessageInfo

func (m *UploadClickConversionsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *UploadClickConversionsRequest) GetConversions() []*ClickConversion {
	if m != nil {
		return m.Conversions
	}
	return nil
}

func (m *UploadClickConversionsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

// Response message for [ConversionUploadService.UploadClickConversions][google.ads.googleads.v1.services.ConversionUploadService.UploadClickConversions].
type UploadClickConversionsResponse struct {
	// Errors that pertain to conversion failures in the partial failure mode.
	// Returned when all errors occur inside the conversions. If any errors occur
	// outside the conversions (e.g. auth errors), we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,1,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// Returned for successfully processed conversions. Proto will be empty for
	// rows that received an error. Results are not returned when validate_only is
	// true.
	Results              []*ClickConversionResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *UploadClickConversionsResponse) Reset()         { *m = UploadClickConversionsResponse{} }
func (m *UploadClickConversionsResponse) String() string { return proto.CompactTextString(m) }
func (*UploadClickConversionsResponse) ProtoMessage()    {}
func (*UploadClickConversionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_upload_service_96e37aa981399bfd, []int{1}
}
func (m *UploadClickConversionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadClickConversionsResponse.Unmarshal(m, b)
}
func (m *UploadClickConversionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadClickConversionsResponse.Marshal(b, m, deterministic)
}
func (dst *UploadClickConversionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadClickConversionsResponse.Merge(dst, src)
}
func (m *UploadClickConversionsResponse) XXX_Size() int {
	return xxx_messageInfo_UploadClickConversionsResponse.Size(m)
}
func (m *UploadClickConversionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadClickConversionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadClickConversionsResponse proto.InternalMessageInfo

func (m *UploadClickConversionsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *UploadClickConversionsResponse) GetResults() []*ClickConversionResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// Request message for [ConversionUploadService.UploadCallConversions][google.ads.googleads.v1.services.ConversionUploadService.UploadCallConversions].
type UploadCallConversionsRequest struct {
	// The ID of the customer performing the upload.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The conversions that are being uploaded.
	Conversions []*CallConversion `protobuf:"bytes,2,rep,name=conversions,proto3" json:"conversions,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// This should always be set to true.
	PartialFailure       bool     `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadCallConversionsRequest) Reset()         { *m = UploadCallConversionsRequest{} }
func (m *UploadCallConversionsRequest) String() string { return proto.CompactTextString(m) }
func (*UploadCallConversionsRequest) ProtoMessage()    {}
func (*UploadCallConversionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_upload_service_96e37aa981399bfd, []int{2}
}
func (m *UploadCallConversionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadCallConversionsRequest.Unmarshal(m, b)
}
func (m *UploadCallConversionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadCallConversionsRequest.Marshal(b, m, deterministic)
}
func (dst *UploadCallConversionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadCallConversionsRequest.Merge(dst, src)
}
func (m *UploadCallConversionsRequest) XXX_Size() int {
	return xxx_messageInfo_UploadCallConversionsRequest.Size(m)
}
func (m *UploadCallConversionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadCallConversionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadCallConversionsRequest proto.InternalMessageInfo

func (m *UploadCallConversionsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *UploadCallConversionsRequest) GetConversions() []*CallConversion {
	if m != nil {
		return m.Conversions
	}
	return nil
}

func (m *UploadCallConversionsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

// Response message for [ConversionUploadService.UploadCallConversions][google.ads.googleads.v1.services.ConversionUploadService.UploadCallConversions].
type UploadCallConversionsResponse struct {
	// Errors that pertain to conversion failures in the partial failure mode.
	// Returned when all errors occur inside the conversions. If any errors occur
	// outside the conversions (e.g. auth errors), we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,1,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// Returned for successfully processed conversions. Proto will be empty for
	// rows that received an error. Results are not returned when validate_only is
	// true.
	Results              []*CallConversionResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *UploadCallConversionsResponse) Reset()         { *m = UploadCallConversionsResponse{} }
func (m *UploadCallConversionsResponse) String() string { return proto.CompactTextString(m) }
func (*UploadCallConversionsResponse) ProtoMessage()    {}
func (*UploadCallConversionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_upload_service_96e37aa981399bfd, []int{3}
}
func (m *UploadCallConversionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadCallConversionsResponse.Unmarshal(m, b)
}
func (m *UploadCallConversionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadCallConversionsResponse.Marshal(b, m, deterministic)
}
func (dst *UploadCallConversionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadCallConversionsResponse.Merge(dst, src)
}
func (m *UploadCallConversionsResponse) XXX_Size() int {
	return xxx_messageInfo_UploadCallConversionsResponse.Size(m)
}
func (m *UploadCallConversionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadCallConversionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadCallConversionsResponse proto.InternalMessageInfo

func (m *UploadCallConversionsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *UploadCallConversionsResponse) GetResults() []*CallConversionResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// A click conversion.
type ClickConversion struct {
	// The Google click ID (gclid) associated with this conversion.
	Gclid *wrappers.StringValue `protobuf:"bytes,1,opt,name=gclid,proto3" json:"gclid,omitempty"`
	// Resource name of the conversion action associated with this conversion.
	// Note: Although this resource name consists of a customer id and a
	// conversion action id, validation will ignore the customer id and use the
	// conversion action id as the sole identifier of the conversion action.
	ConversionAction *wrappers.StringValue `protobuf:"bytes,2,opt,name=conversion_action,json=conversionAction,proto3" json:"conversion_action,omitempty"`
	// The date time at which the conversion occurred. Must be after
	// the click time. The timezone must be specified. The format is
	// "yyyy-mm-dd hh:mm:ss+|-hh:mm", e.g. “2019-01-01 12:32:45-08:00”.
	ConversionDateTime *wrappers.StringValue `protobuf:"bytes,3,opt,name=conversion_date_time,json=conversionDateTime,proto3" json:"conversion_date_time,omitempty"`
	// The value of the conversion for the advertiser.
	ConversionValue *wrappers.DoubleValue `protobuf:"bytes,4,opt,name=conversion_value,json=conversionValue,proto3" json:"conversion_value,omitempty"`
	// Currency associated with the conversion value. This is the ISO 4217
	// 3-character currency code. For example: USD, EUR.
	CurrencyCode *wrappers.StringValue `protobuf:"bytes,5,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// The order ID associated with the conversion. An order id can only be used
	// for one conversion per conversion action.
	OrderId *wrappers.StringValue `protobuf:"bytes,6,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	// Additional data about externally attributed conversions. This field
	// is required for conversions with an externally attributed conversion
	// action, but should not be set otherwise.
	ExternalAttributionData *ExternalAttributionData `protobuf:"bytes,7,opt,name=external_attribution_data,json=externalAttributionData,proto3" json:"external_attribution_data,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                 `json:"-"`
	XXX_unrecognized        []byte                   `json:"-"`
	XXX_sizecache           int32                    `json:"-"`
}

func (m *ClickConversion) Reset()         { *m = ClickConversion{} }
func (m *ClickConversion) String() string { return proto.CompactTextString(m) }
func (*ClickConversion) ProtoMessage()    {}
func (*ClickConversion) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_upload_service_96e37aa981399bfd, []int{4}
}
func (m *ClickConversion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClickConversion.Unmarshal(m, b)
}
func (m *ClickConversion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClickConversion.Marshal(b, m, deterministic)
}
func (dst *ClickConversion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClickConversion.Merge(dst, src)
}
func (m *ClickConversion) XXX_Size() int {
	return xxx_messageInfo_ClickConversion.Size(m)
}
func (m *ClickConversion) XXX_DiscardUnknown() {
	xxx_messageInfo_ClickConversion.DiscardUnknown(m)
}

var xxx_messageInfo_ClickConversion proto.InternalMessageInfo

func (m *ClickConversion) GetGclid() *wrappers.StringValue {
	if m != nil {
		return m.Gclid
	}
	return nil
}

func (m *ClickConversion) GetConversionAction() *wrappers.StringValue {
	if m != nil {
		return m.ConversionAction
	}
	return nil
}

func (m *ClickConversion) GetConversionDateTime() *wrappers.StringValue {
	if m != nil {
		return m.ConversionDateTime
	}
	return nil
}

func (m *ClickConversion) GetConversionValue() *wrappers.DoubleValue {
	if m != nil {
		return m.ConversionValue
	}
	return nil
}

func (m *ClickConversion) GetCurrencyCode() *wrappers.StringValue {
	if m != nil {
		return m.CurrencyCode
	}
	return nil
}

func (m *ClickConversion) GetOrderId() *wrappers.StringValue {
	if m != nil {
		return m.OrderId
	}
	return nil
}

func (m *ClickConversion) GetExternalAttributionData() *ExternalAttributionData {
	if m != nil {
		return m.ExternalAttributionData
	}
	return nil
}

// A call conversion.
type CallConversion struct {
	// The caller id from which this call was placed. Caller id is expected to be
	// in E.164 format with preceding '+' sign. e.g. "+16502531234".
	CallerId *wrappers.StringValue `protobuf:"bytes,1,opt,name=caller_id,json=callerId,proto3" json:"caller_id,omitempty"`
	// The date time at which the call occurred. The timezone must be specified.
	// The format is "yyyy-mm-dd hh:mm:ss+|-hh:mm",
	// e.g. "2019-01-01 12:32:45-08:00".
	CallStartDateTime *wrappers.StringValue `protobuf:"bytes,2,opt,name=call_start_date_time,json=callStartDateTime,proto3" json:"call_start_date_time,omitempty"`
	// Resource name of the conversion action associated with this conversion.
	// Note: Although this resource name consists of a customer id and a
	// conversion action id, validation will ignore the customer id and use the
	// conversion action id as the sole identifier of the conversion action.
	ConversionAction *wrappers.StringValue `protobuf:"bytes,3,opt,name=conversion_action,json=conversionAction,proto3" json:"conversion_action,omitempty"`
	// The date time at which the conversion occurred. Must be after the call
	// time. The timezone must be specified. The format is
	// "yyyy-mm-dd hh:mm:ss+|-hh:mm", e.g. "2019-01-01 12:32:45-08:00".
	ConversionDateTime *wrappers.StringValue `protobuf:"bytes,4,opt,name=conversion_date_time,json=conversionDateTime,proto3" json:"conversion_date_time,omitempty"`
	// The value of the conversion for the advertiser.
	ConversionValue *wrappers.DoubleValue `protobuf:"bytes,5,opt,name=conversion_value,json=conversionValue,proto3" json:"conversion_value,omitempty"`
	// Currency associated with the conversion value. This is the ISO 4217
	// 3-character currency code. For example: USD, EUR.
	CurrencyCode         *wrappers.StringValue `protobuf:"bytes,6,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CallConversion) Reset()         { *m = CallConversion{} }
func (m *CallConversion) String() string { return proto.CompactTextString(m) }
func (*CallConversion) ProtoMessage()    {}
func (*CallConversion) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_upload_service_96e37aa981399bfd, []int{5}
}
func (m *CallConversion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallConversion.Unmarshal(m, b)
}
func (m *CallConversion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallConversion.Marshal(b, m, deterministic)
}
func (dst *CallConversion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallConversion.Merge(dst, src)
}
func (m *CallConversion) XXX_Size() int {
	return xxx_messageInfo_CallConversion.Size(m)
}
func (m *CallConversion) XXX_DiscardUnknown() {
	xxx_messageInfo_CallConversion.DiscardUnknown(m)
}

var xxx_messageInfo_CallConversion proto.InternalMessageInfo

func (m *CallConversion) GetCallerId() *wrappers.StringValue {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *CallConversion) GetCallStartDateTime() *wrappers.StringValue {
	if m != nil {
		return m.CallStartDateTime
	}
	return nil
}

func (m *CallConversion) GetConversionAction() *wrappers.StringValue {
	if m != nil {
		return m.ConversionAction
	}
	return nil
}

func (m *CallConversion) GetConversionDateTime() *wrappers.StringValue {
	if m != nil {
		return m.ConversionDateTime
	}
	return nil
}

func (m *CallConversion) GetConversionValue() *wrappers.DoubleValue {
	if m != nil {
		return m.ConversionValue
	}
	return nil
}

func (m *CallConversion) GetCurrencyCode() *wrappers.StringValue {
	if m != nil {
		return m.CurrencyCode
	}
	return nil
}

// Contains additional information about externally attributed conversions.
type ExternalAttributionData struct {
	// Represents the fraction of the conversion that is attributed to the
	// Google Ads click.
	ExternalAttributionCredit *wrappers.DoubleValue `protobuf:"bytes,1,opt,name=external_attribution_credit,json=externalAttributionCredit,proto3" json:"external_attribution_credit,omitempty"`
	// Specifies the attribution model name.
	ExternalAttributionModel *wrappers.StringValue `protobuf:"bytes,2,opt,name=external_attribution_model,json=externalAttributionModel,proto3" json:"external_attribution_model,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}              `json:"-"`
	XXX_unrecognized         []byte                `json:"-"`
	XXX_sizecache            int32                 `json:"-"`
}

func (m *ExternalAttributionData) Reset()         { *m = ExternalAttributionData{} }
func (m *ExternalAttributionData) String() string { return proto.CompactTextString(m) }
func (*ExternalAttributionData) ProtoMessage()    {}
func (*ExternalAttributionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_upload_service_96e37aa981399bfd, []int{6}
}
func (m *ExternalAttributionData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalAttributionData.Unmarshal(m, b)
}
func (m *ExternalAttributionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalAttributionData.Marshal(b, m, deterministic)
}
func (dst *ExternalAttributionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalAttributionData.Merge(dst, src)
}
func (m *ExternalAttributionData) XXX_Size() int {
	return xxx_messageInfo_ExternalAttributionData.Size(m)
}
func (m *ExternalAttributionData) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalAttributionData.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalAttributionData proto.InternalMessageInfo

func (m *ExternalAttributionData) GetExternalAttributionCredit() *wrappers.DoubleValue {
	if m != nil {
		return m.ExternalAttributionCredit
	}
	return nil
}

func (m *ExternalAttributionData) GetExternalAttributionModel() *wrappers.StringValue {
	if m != nil {
		return m.ExternalAttributionModel
	}
	return nil
}

// Identifying information for a successfully processed ClickConversion.
type ClickConversionResult struct {
	// The Google Click ID (gclid) associated with this conversion.
	Gclid *wrappers.StringValue `protobuf:"bytes,1,opt,name=gclid,proto3" json:"gclid,omitempty"`
	// Resource name of the conversion action associated with this conversion.
	ConversionAction *wrappers.StringValue `protobuf:"bytes,2,opt,name=conversion_action,json=conversionAction,proto3" json:"conversion_action,omitempty"`
	// The date time at which the conversion occurred. The format is
	// "yyyy-mm-dd hh:mm:ss+|-hh:mm", e.g. “2019-01-01 12:32:45-08:00”.
	ConversionDateTime   *wrappers.StringValue `protobuf:"bytes,3,opt,name=conversion_date_time,json=conversionDateTime,proto3" json:"conversion_date_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClickConversionResult) Reset()         { *m = ClickConversionResult{} }
func (m *ClickConversionResult) String() string { return proto.CompactTextString(m) }
func (*ClickConversionResult) ProtoMessage()    {}
func (*ClickConversionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_upload_service_96e37aa981399bfd, []int{7}
}
func (m *ClickConversionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClickConversionResult.Unmarshal(m, b)
}
func (m *ClickConversionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClickConversionResult.Marshal(b, m, deterministic)
}
func (dst *ClickConversionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClickConversionResult.Merge(dst, src)
}
func (m *ClickConversionResult) XXX_Size() int {
	return xxx_messageInfo_ClickConversionResult.Size(m)
}
func (m *ClickConversionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ClickConversionResult.DiscardUnknown(m)
}

var xxx_messageInfo_ClickConversionResult proto.InternalMessageInfo

func (m *ClickConversionResult) GetGclid() *wrappers.StringValue {
	if m != nil {
		return m.Gclid
	}
	return nil
}

func (m *ClickConversionResult) GetConversionAction() *wrappers.StringValue {
	if m != nil {
		return m.ConversionAction
	}
	return nil
}

func (m *ClickConversionResult) GetConversionDateTime() *wrappers.StringValue {
	if m != nil {
		return m.ConversionDateTime
	}
	return nil
}

// Identifying information for a successfully processed CallConversionUpload.
type CallConversionResult struct {
	// The caller id from which this call was placed. Caller id is expected to be
	// in E.164 format with preceding '+' sign.
	CallerId *wrappers.StringValue `protobuf:"bytes,1,opt,name=caller_id,json=callerId,proto3" json:"caller_id,omitempty"`
	// The date time at which the call occurred. The format is
	// "yyyy-mm-dd hh:mm:ss+|-hh:mm", e.g. "2019-01-01 12:32:45-08:00".
	CallStartDateTime *wrappers.StringValue `protobuf:"bytes,2,opt,name=call_start_date_time,json=callStartDateTime,proto3" json:"call_start_date_time,omitempty"`
	// Resource name of the conversion action associated with this conversion.
	ConversionAction *wrappers.StringValue `protobuf:"bytes,3,opt,name=conversion_action,json=conversionAction,proto3" json:"conversion_action,omitempty"`
	// The date time at which the conversion occurred. The format is
	// "yyyy-mm-dd hh:mm:ss+|-hh:mm", e.g. "2019-01-01 12:32:45-08:00".
	ConversionDateTime   *wrappers.StringValue `protobuf:"bytes,4,opt,name=conversion_date_time,json=conversionDateTime,proto3" json:"conversion_date_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CallConversionResult) Reset()         { *m = CallConversionResult{} }
func (m *CallConversionResult) String() string { return proto.CompactTextString(m) }
func (*CallConversionResult) ProtoMessage()    {}
func (*CallConversionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_upload_service_96e37aa981399bfd, []int{8}
}
func (m *CallConversionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallConversionResult.Unmarshal(m, b)
}
func (m *CallConversionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallConversionResult.Marshal(b, m, deterministic)
}
func (dst *CallConversionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallConversionResult.Merge(dst, src)
}
func (m *CallConversionResult) XXX_Size() int {
	return xxx_messageInfo_CallConversionResult.Size(m)
}
func (m *CallConversionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CallConversionResult.DiscardUnknown(m)
}

var xxx_messageInfo_CallConversionResult proto.InternalMessageInfo

func (m *CallConversionResult) GetCallerId() *wrappers.StringValue {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *CallConversionResult) GetCallStartDateTime() *wrappers.StringValue {
	if m != nil {
		return m.CallStartDateTime
	}
	return nil
}

func (m *CallConversionResult) GetConversionAction() *wrappers.StringValue {
	if m != nil {
		return m.ConversionAction
	}
	return nil
}

func (m *CallConversionResult) GetConversionDateTime() *wrappers.StringValue {
	if m != nil {
		return m.ConversionDateTime
	}
	return nil
}

func init() {
	proto.RegisterType((*UploadClickConversionsRequest)(nil), "google.ads.googleads.v1.services.UploadClickConversionsRequest")
	proto.RegisterType((*UploadClickConversionsResponse)(nil), "google.ads.googleads.v1.services.UploadClickConversionsResponse")
	proto.RegisterType((*UploadCallConversionsRequest)(nil), "google.ads.googleads.v1.services.UploadCallConversionsRequest")
	proto.RegisterType((*UploadCallConversionsResponse)(nil), "google.ads.googleads.v1.services.UploadCallConversionsResponse")
	proto.RegisterType((*ClickConversion)(nil), "google.ads.googleads.v1.services.ClickConversion")
	proto.RegisterType((*CallConversion)(nil), "google.ads.googleads.v1.services.CallConversion")
	proto.RegisterType((*ExternalAttributionData)(nil), "google.ads.googleads.v1.services.ExternalAttributionData")
	proto.RegisterType((*ClickConversionResult)(nil), "google.ads.googleads.v1.services.ClickConversionResult")
	proto.RegisterType((*CallConversionResult)(nil), "google.ads.googleads.v1.services.CallConversionResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConversionUploadServiceClient is the client API for ConversionUploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConversionUploadServiceClient interface {
	// Processes the given click conversions.
	UploadClickConversions(ctx context.Context, in *UploadClickConversionsRequest, opts ...grpc.CallOption) (*UploadClickConversionsResponse, error)
	// Processes the given call conversions.
	UploadCallConversions(ctx context.Context, in *UploadCallConversionsRequest, opts ...grpc.CallOption) (*UploadCallConversionsResponse, error)
}

type conversionUploadServiceClient struct {
	cc *grpc.ClientConn
}

func NewConversionUploadServiceClient(cc *grpc.ClientConn) ConversionUploadServiceClient {
	return &conversionUploadServiceClient{cc}
}

func (c *conversionUploadServiceClient) UploadClickConversions(ctx context.Context, in *UploadClickConversionsRequest, opts ...grpc.CallOption) (*UploadClickConversionsResponse, error) {
	out := new(UploadClickConversionsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ConversionUploadService/UploadClickConversions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversionUploadServiceClient) UploadCallConversions(ctx context.Context, in *UploadCallConversionsRequest, opts ...grpc.CallOption) (*UploadCallConversionsResponse, error) {
	out := new(UploadCallConversionsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ConversionUploadService/UploadCallConversions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversionUploadServiceServer is the server API for ConversionUploadService service.
type ConversionUploadServiceServer interface {
	// Processes the given click conversions.
	UploadClickConversions(context.Context, *UploadClickConversionsRequest) (*UploadClickConversionsResponse, error)
	// Processes the given call conversions.
	UploadCallConversions(context.Context, *UploadCallConversionsRequest) (*UploadCallConversionsResponse, error)
}

func RegisterConversionUploadServiceServer(s *grpc.Server, srv ConversionUploadServiceServer) {
	s.RegisterService(&_ConversionUploadService_serviceDesc, srv)
}

func _ConversionUploadService_UploadClickConversions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadClickConversionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionUploadServiceServer).UploadClickConversions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ConversionUploadService/UploadClickConversions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionUploadServiceServer).UploadClickConversions(ctx, req.(*UploadClickConversionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConversionUploadService_UploadCallConversions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadCallConversionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionUploadServiceServer).UploadCallConversions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ConversionUploadService/UploadCallConversions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionUploadServiceServer).UploadCallConversions(ctx, req.(*UploadCallConversionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConversionUploadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.ConversionUploadService",
	HandlerType: (*ConversionUploadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadClickConversions",
			Handler:    _ConversionUploadService_UploadClickConversions_Handler,
		},
		{
			MethodName: "UploadCallConversions",
			Handler:    _ConversionUploadService_UploadCallConversions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/conversion_upload_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/conversion_upload_service.proto", fileDescriptor_conversion_upload_service_96e37aa981399bfd)
}

var fileDescriptor_conversion_upload_service_96e37aa981399bfd = []byte{
	// 900 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x57, 0x4d, 0x6f, 0xe3, 0x44,
	0x18, 0x96, 0x93, 0x7e, 0xed, 0x04, 0xb6, 0xec, 0xd0, 0x52, 0x6f, 0x28, 0x4b, 0x64, 0xad, 0x44,
	0xd5, 0x83, 0x4d, 0xb2, 0x88, 0xa5, 0x46, 0x6c, 0xd7, 0xdb, 0xee, 0x96, 0x1e, 0x16, 0x15, 0x07,
	0x7a, 0x58, 0x45, 0xb2, 0xa6, 0x9e, 0x59, 0xcb, 0x62, 0xe2, 0x31, 0x33, 0xe3, 0x00, 0x42, 0x5c,
	0xe0, 0x27, 0x70, 0xe5, 0xc4, 0x91, 0x7f, 0x80, 0x8a, 0x90, 0xb8, 0x72, 0xed, 0x85, 0x03, 0x47,
	0xc4, 0x81, 0x5f, 0x81, 0xfc, 0x31, 0x71, 0x1a, 0x39, 0xc4, 0xad, 0x72, 0x42, 0x7b, 0x9b, 0xcc,
	0xbc, 0xcf, 0xf3, 0x7e, 0x3d, 0x33, 0x79, 0x0d, 0x1e, 0x06, 0x8c, 0x05, 0x94, 0x58, 0x08, 0x0b,
	0x2b, 0x5f, 0xa6, 0xab, 0x51, 0xd7, 0x12, 0x84, 0x8f, 0x42, 0x9f, 0x08, 0xcb, 0x67, 0xd1, 0x88,
	0x70, 0x11, 0xb2, 0xc8, 0x4b, 0x62, 0xca, 0x10, 0xf6, 0x8a, 0x23, 0x33, 0xe6, 0x4c, 0x32, 0xd8,
	0xc9, 0x61, 0x26, 0xc2, 0xc2, 0x1c, 0x33, 0x98, 0xa3, 0xae, 0xa9, 0x18, 0xda, 0xdb, 0xca, 0x47,
	0x1c, 0x5a, 0x28, 0x8a, 0x98, 0x44, 0x32, 0x64, 0x91, 0xc8, 0xf1, 0xed, 0x3b, 0xc5, 0x69, 0xf6,
	0xeb, 0x2c, 0x79, 0x6e, 0x7d, 0xc1, 0x51, 0x1c, 0x13, 0xae, 0xce, 0xb7, 0x8a, 0x73, 0x1e, 0xfb,
	0x96, 0x90, 0x48, 0x26, 0xc5, 0x81, 0xf1, 0xab, 0x06, 0xde, 0xf8, 0x34, 0x8b, 0xe8, 0x80, 0x86,
	0xfe, 0x67, 0x07, 0xe3, 0x38, 0x85, 0x4b, 0x3e, 0x4f, 0x88, 0x90, 0xf0, 0x4d, 0xd0, 0xf2, 0x13,
	0x21, 0xd9, 0x90, 0x70, 0x2f, 0xc4, 0xba, 0xd6, 0xd1, 0x76, 0x6e, 0xb8, 0x40, 0x6d, 0x1d, 0x63,
	0xd8, 0x07, 0xad, 0x32, 0x3d, 0xa1, 0x37, 0x3a, 0xcd, 0x9d, 0x56, 0xaf, 0x6b, 0xce, 0xcb, 0xc8,
	0x9c, 0x72, 0xe8, 0x4e, 0xb2, 0xc0, 0xb7, 0xc0, 0x7a, 0x8c, 0xb8, 0x0c, 0x11, 0xf5, 0x9e, 0xa3,
	0x90, 0x26, 0x9c, 0xe8, 0xcd, 0x8e, 0xb6, 0xb3, 0xe6, 0xde, 0x2c, 0xb6, 0x9f, 0xe4, 0xbb, 0xc6,
	0xb9, 0x06, 0xee, 0xcc, 0x4a, 0x40, 0xc4, 0x2c, 0x12, 0x04, 0x3e, 0x01, 0x9b, 0x53, 0x5c, 0x1e,
	0xe1, 0x9c, 0xf1, 0x2c, 0x97, 0x56, 0x0f, 0xaa, 0x50, 0x79, 0xec, 0x9b, 0xfd, 0xac, 0x38, 0xee,
	0xab, 0x97, 0xbd, 0x3c, 0x4e, 0xcd, 0xe1, 0xc7, 0x60, 0x95, 0x13, 0x91, 0x50, 0xa9, 0x92, 0xbc,
	0x7f, 0xf5, 0x24, 0x33, 0xbc, 0xab, 0x78, 0x8c, 0x5f, 0x34, 0xb0, 0x5d, 0x44, 0x8f, 0x28, 0xbd,
	0x4e, 0xf5, 0xdd, 0xaa, 0xea, 0xbf, 0x5d, 0x23, 0xb0, 0x4b, 0xfe, 0xae, 0x59, 0xfc, 0x9f, 0x4b,
	0xf5, 0x4c, 0x87, 0xbf, 0xe0, 0xda, 0x9f, 0x4c, 0xd7, 0xfe, 0xdd, 0x2b, 0xa7, 0x38, 0x55, 0xfa,
	0x1f, 0x96, 0xc0, 0xfa, 0x54, 0x77, 0x60, 0x0f, 0x2c, 0x07, 0x3e, 0x2d, 0xea, 0xdc, 0xea, 0x6d,
	0x2b, 0x1f, 0xea, 0x5a, 0x99, 0x7d, 0xc9, 0xc3, 0x28, 0x38, 0x45, 0x34, 0x21, 0x6e, 0x6e, 0x0a,
	0x8f, 0xc1, 0xad, 0x89, 0xdb, 0x8d, 0xfc, 0xf4, 0x5a, 0xea, 0x8d, 0x1a, 0xf8, 0x57, 0x4a, 0x98,
	0x93, 0xa1, 0xe0, 0x47, 0x60, 0x63, 0x82, 0x0a, 0x23, 0x49, 0x3c, 0x19, 0x0e, 0xf3, 0xe2, 0xcf,
	0x63, 0x83, 0x25, 0xf2, 0x10, 0x49, 0xf2, 0x49, 0x38, 0x24, 0xf0, 0x08, 0x4c, 0xf8, 0xf0, 0x46,
	0xa9, 0x9d, 0xbe, 0x34, 0x83, 0xeb, 0x90, 0x25, 0x67, 0x94, 0xe4, 0x5c, 0xeb, 0x25, 0x2a, 0xdb,
	0x80, 0x0e, 0x78, 0xd9, 0x4f, 0x38, 0x27, 0x91, 0xff, 0x95, 0xe7, 0x33, 0x4c, 0xf4, 0xe5, 0x1a,
	0x11, 0xbd, 0xa4, 0x20, 0x07, 0x0c, 0x13, 0x78, 0x1f, 0xac, 0x31, 0x8e, 0x73, 0x15, 0xaf, 0xd4,
	0x40, 0xaf, 0x66, 0xd6, 0xc7, 0x18, 0x26, 0xe0, 0x36, 0xf9, 0x52, 0x12, 0x1e, 0x21, 0xea, 0x21,
	0x29, 0x79, 0x78, 0x96, 0xc8, 0xa2, 0x3c, 0x48, 0x5f, 0xcd, 0x98, 0xf6, 0xe6, 0x6b, 0xe1, 0x71,
	0x41, 0xe1, 0x94, 0x0c, 0x87, 0x48, 0x22, 0x77, 0x8b, 0x54, 0x1f, 0x18, 0xbf, 0x35, 0xc1, 0xcd,
	0xcb, 0x02, 0x82, 0x7b, 0xe0, 0x86, 0x8f, 0x28, 0x2d, 0x6f, 0xe2, 0xbc, 0x1c, 0xd6, 0x72, 0xf3,
	0x63, 0x0c, 0x9f, 0x82, 0x8d, 0x74, 0xed, 0x09, 0x89, 0xb8, 0x9c, 0xe8, 0x6c, 0x1d, 0x9d, 0xdc,
	0x4a, 0x91, 0xfd, 0x14, 0x38, 0x6e, 0x6c, 0xa5, 0xe6, 0x9a, 0x0b, 0xd5, 0xdc, 0xd2, 0x02, 0x35,
	0xb7, 0xbc, 0x10, 0xcd, 0xad, 0x5c, 0x55, 0x73, 0xc6, 0x85, 0x06, 0xb6, 0x66, 0x34, 0x1e, 0x0e,
	0xc0, 0xeb, 0x95, 0xb2, 0xf2, 0x39, 0xc1, 0xa1, 0x9c, 0xd9, 0xde, 0xc9, 0x90, 0x6f, 0x57, 0x68,
	0xe7, 0x20, 0x83, 0xc3, 0x67, 0xa0, 0x5d, 0xc9, 0x3e, 0x64, 0x98, 0xd0, 0x5a, 0x5d, 0xd7, 0x2b,
	0xc8, 0x9f, 0xa6, 0x68, 0xe3, 0x6f, 0x0d, 0x6c, 0x56, 0xfe, 0xad, 0xfc, 0xcf, 0x9e, 0x2f, 0xe3,
	0xbc, 0x01, 0x36, 0xaa, 0xde, 0xf0, 0x17, 0x17, 0xb1, 0x46, 0xf5, 0x7a, 0x17, 0x4d, 0xb0, 0x55,
	0x56, 0x2e, 0xff, 0x97, 0xee, 0xe7, 0xcf, 0x21, 0xfc, 0x53, 0x03, 0xaf, 0x55, 0x0f, 0x4d, 0x70,
	0x7f, 0xfe, 0x5b, 0xfa, 0x9f, 0xf3, 0x62, 0xfb, 0xe1, 0xf5, 0x09, 0xf2, 0x99, 0xc1, 0xd8, 0xff,
	0xf6, 0xe2, 0xaf, 0xef, 0x1b, 0x7b, 0xc6, 0x3b, 0xe9, 0x08, 0xad, 0x46, 0x1d, 0x61, 0x7d, 0x3d,
	0x31, 0x08, 0x7d, 0xb0, 0xfb, 0x8d, 0x9d, 0x54, 0xb2, 0xd8, 0xda, 0x2e, 0xfc, 0x43, 0x03, 0x9b,
	0x95, 0x63, 0x09, 0x7c, 0x50, 0x3b, 0xb8, 0xca, 0x71, 0xac, 0xbd, 0x7f, 0x6d, 0x7c, 0x91, 0xdb,
	0x83, 0x2c, 0xb7, 0xf7, 0x8c, 0x7b, 0xb5, 0x72, 0xbb, 0x4c, 0x62, 0x6b, 0xbb, 0x8f, 0xbe, 0x6b,
	0x80, 0xbb, 0x3e, 0x1b, 0xce, 0x0d, 0xe3, 0xd1, 0xf6, 0x8c, 0xde, 0x9f, 0xa4, 0xfa, 0x39, 0xd1,
	0x9e, 0x7d, 0x58, 0x30, 0x04, 0x8c, 0xa2, 0x28, 0x30, 0x19, 0x0f, 0xac, 0x80, 0x44, 0x99, 0xba,
	0xd4, 0x37, 0x4c, 0x1c, 0x8a, 0xd9, 0x9f, 0x34, 0xef, 0xab, 0xc5, 0x8f, 0x8d, 0xe6, 0x91, 0xe3,
	0xfc, 0xd4, 0xe8, 0x1c, 0xe5, 0x84, 0x0e, 0x16, 0x66, 0xbe, 0x4c, 0x57, 0xa7, 0x5d, 0xb3, 0x70,
	0x2c, 0x7e, 0x57, 0x26, 0x03, 0x07, 0x8b, 0xc1, 0xd8, 0x64, 0x70, 0xda, 0x1d, 0x28, 0x93, 0x7f,
	0x1a, 0x77, 0xf3, 0x7d, 0xdb, 0x76, 0xb0, 0xb0, 0xed, 0xb1, 0x91, 0x6d, 0x9f, 0x76, 0x6d, 0x5b,
	0x99, 0x9d, 0xad, 0x64, 0x71, 0xde, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xa3, 0xe2, 0x9c,
	0x79, 0x0d, 0x00, 0x00,
}
