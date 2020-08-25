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
// source: google/ads/googleads/v2/services/account_budget_proposal_service.proto

package services

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// Request message for
// [AccountBudgetProposalService.GetAccountBudgetProposal][google.ads.googleads.v2.services.AccountBudgetProposalService.GetAccountBudgetProposal].
type GetAccountBudgetProposalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the account-level budget proposal to fetch.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *GetAccountBudgetProposalRequest) Reset() {
	*x = GetAccountBudgetProposalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountBudgetProposalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountBudgetProposalRequest) ProtoMessage() {}

func (x *GetAccountBudgetProposalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountBudgetProposalRequest.ProtoReflect.Descriptor instead.
func (*GetAccountBudgetProposalRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAccountBudgetProposalRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Request message for
// [AccountBudgetProposalService.MutateAccountBudgetProposal][google.ads.googleads.v2.services.AccountBudgetProposalService.MutateAccountBudgetProposal].
type MutateAccountBudgetProposalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The operation to perform on an individual account-level budget proposal.
	Operation *AccountBudgetProposalOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *MutateAccountBudgetProposalRequest) Reset() {
	*x = MutateAccountBudgetProposalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAccountBudgetProposalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAccountBudgetProposalRequest) ProtoMessage() {}

func (x *MutateAccountBudgetProposalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAccountBudgetProposalRequest.ProtoReflect.Descriptor instead.
func (*MutateAccountBudgetProposalRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_rawDescGZIP(), []int{1}
}

func (x *MutateAccountBudgetProposalRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateAccountBudgetProposalRequest) GetOperation() *AccountBudgetProposalOperation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *MutateAccountBudgetProposalRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// A single operation to propose the creation of a new account-level budget or
// edit/end/remove an existing one.
type AccountBudgetProposalOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask that determines which budget fields are modified.  While budgets
	// may be modified, proposals that propose such modifications are final.
	// Therefore, update operations are not supported for proposals.
	//
	// Proposals that modify budgets have the 'update' proposal type.  Specifying
	// a mask for any other proposal type is considered an error.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*AccountBudgetProposalOperation_Create
	//	*AccountBudgetProposalOperation_Remove
	Operation isAccountBudgetProposalOperation_Operation `protobuf_oneof:"operation"`
}

func (x *AccountBudgetProposalOperation) Reset() {
	*x = AccountBudgetProposalOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountBudgetProposalOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountBudgetProposalOperation) ProtoMessage() {}

func (x *AccountBudgetProposalOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountBudgetProposalOperation.ProtoReflect.Descriptor instead.
func (*AccountBudgetProposalOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_rawDescGZIP(), []int{2}
}

func (x *AccountBudgetProposalOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *AccountBudgetProposalOperation) GetOperation() isAccountBudgetProposalOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *AccountBudgetProposalOperation) GetCreate() *resources.AccountBudgetProposal {
	if x, ok := x.GetOperation().(*AccountBudgetProposalOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *AccountBudgetProposalOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*AccountBudgetProposalOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isAccountBudgetProposalOperation_Operation interface {
	isAccountBudgetProposalOperation_Operation()
}

type AccountBudgetProposalOperation_Create struct {
	// Create operation: A new proposal to create a new budget, edit an
	// existing budget, end an actively running budget, or remove an approved
	// budget scheduled to start in the future.
	// No resource name is expected for the new proposal.
	Create *resources.AccountBudgetProposal `protobuf:"bytes,2,opt,name=create,proto3,oneof"`
}

type AccountBudgetProposalOperation_Remove struct {
	// Remove operation: A resource name for the removed proposal is expected,
	// in this format:
	//
	// `customers/{customer_id}/accountBudgetProposals/{account_budget_proposal_id}`
	// A request may be cancelled iff it is pending.
	Remove string `protobuf:"bytes,1,opt,name=remove,proto3,oneof"`
}

func (*AccountBudgetProposalOperation_Create) isAccountBudgetProposalOperation_Operation() {}

func (*AccountBudgetProposalOperation_Remove) isAccountBudgetProposalOperation_Operation() {}

// Response message for account-level budget mutate operations.
type MutateAccountBudgetProposalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The result of the mutate.
	Result *MutateAccountBudgetProposalResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MutateAccountBudgetProposalResponse) Reset() {
	*x = MutateAccountBudgetProposalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAccountBudgetProposalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAccountBudgetProposalResponse) ProtoMessage() {}

func (x *MutateAccountBudgetProposalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAccountBudgetProposalResponse.ProtoReflect.Descriptor instead.
func (*MutateAccountBudgetProposalResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateAccountBudgetProposalResponse) GetResult() *MutateAccountBudgetProposalResult {
	if x != nil {
		return x.Result
	}
	return nil
}

// The result for the account budget proposal mutate.
type MutateAccountBudgetProposalResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateAccountBudgetProposalResult) Reset() {
	*x = MutateAccountBudgetProposalResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAccountBudgetProposalResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAccountBudgetProposalResult) ProtoMessage() {}

func (x *MutateAccountBudgetProposalResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAccountBudgetProposalResult.ProtoReflect.Descriptor instead.
func (*MutateAccountBudgetProposalResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_rawDescGZIP(), []int{4}
}

func (x *MutateAccountBudgetProposalResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v2_services_account_budget_proposal_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_rawDesc = []byte{
	0x0a, 0x46, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x3f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7e, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x5b, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xe0, 0x41, 0x02, 0xfa,
	0x41, 0x30, 0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xd4, 0x01, 0x0a, 0x22, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x63, 0x0a,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f,
	0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xd8, 0x01, 0x0a, 0x1e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x52, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x82, 0x01, 0x0a, 0x23, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x48, 0x0a, 0x21, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x32, 0xb4, 0x04, 0x0a, 0x1c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xe9, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12,
	0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x22, 0x50, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x3a, 0x12, 0x38, 0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x8a,
	0x02, 0x0a, 0x1b, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x44,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x40, 0x22, 0x3b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d,
	0x2a, 0x7d, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x3a, 0x01, 0x2a, 0xda, 0x41, 0x15, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0xca, 0x41, 0x18,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0x88, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x42, 0x21, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x32,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x5c, 0x56, 0x32, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x24, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_rawDescData = file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_rawDesc
)

func file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_rawDescData
}

var file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_goTypes = []interface{}{
	(*GetAccountBudgetProposalRequest)(nil),     // 0: google.ads.googleads.v2.services.GetAccountBudgetProposalRequest
	(*MutateAccountBudgetProposalRequest)(nil),  // 1: google.ads.googleads.v2.services.MutateAccountBudgetProposalRequest
	(*AccountBudgetProposalOperation)(nil),      // 2: google.ads.googleads.v2.services.AccountBudgetProposalOperation
	(*MutateAccountBudgetProposalResponse)(nil), // 3: google.ads.googleads.v2.services.MutateAccountBudgetProposalResponse
	(*MutateAccountBudgetProposalResult)(nil),   // 4: google.ads.googleads.v2.services.MutateAccountBudgetProposalResult
	(*fieldmaskpb.FieldMask)(nil),               // 5: google.protobuf.FieldMask
	(*resources.AccountBudgetProposal)(nil),     // 6: google.ads.googleads.v2.resources.AccountBudgetProposal
}
var file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v2.services.MutateAccountBudgetProposalRequest.operation:type_name -> google.ads.googleads.v2.services.AccountBudgetProposalOperation
	5, // 1: google.ads.googleads.v2.services.AccountBudgetProposalOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 2: google.ads.googleads.v2.services.AccountBudgetProposalOperation.create:type_name -> google.ads.googleads.v2.resources.AccountBudgetProposal
	4, // 3: google.ads.googleads.v2.services.MutateAccountBudgetProposalResponse.result:type_name -> google.ads.googleads.v2.services.MutateAccountBudgetProposalResult
	0, // 4: google.ads.googleads.v2.services.AccountBudgetProposalService.GetAccountBudgetProposal:input_type -> google.ads.googleads.v2.services.GetAccountBudgetProposalRequest
	1, // 5: google.ads.googleads.v2.services.AccountBudgetProposalService.MutateAccountBudgetProposal:input_type -> google.ads.googleads.v2.services.MutateAccountBudgetProposalRequest
	6, // 6: google.ads.googleads.v2.services.AccountBudgetProposalService.GetAccountBudgetProposal:output_type -> google.ads.googleads.v2.resources.AccountBudgetProposal
	3, // 7: google.ads.googleads.v2.services.AccountBudgetProposalService.MutateAccountBudgetProposal:output_type -> google.ads.googleads.v2.services.MutateAccountBudgetProposalResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_init() }
func file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_init() {
	if File_google_ads_googleads_v2_services_account_budget_proposal_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountBudgetProposalRequest); i {
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
		file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAccountBudgetProposalRequest); i {
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
		file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountBudgetProposalOperation); i {
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
		file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAccountBudgetProposalResponse); i {
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
		file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAccountBudgetProposalResult); i {
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
	file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*AccountBudgetProposalOperation_Create)(nil),
		(*AccountBudgetProposalOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v2_services_account_budget_proposal_service_proto = out.File
	file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_rawDesc = nil
	file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_goTypes = nil
	file_google_ads_googleads_v2_services_account_budget_proposal_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AccountBudgetProposalServiceClient is the client API for AccountBudgetProposalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountBudgetProposalServiceClient interface {
	// Returns an account-level budget proposal in full detail.
	GetAccountBudgetProposal(ctx context.Context, in *GetAccountBudgetProposalRequest, opts ...grpc.CallOption) (*resources.AccountBudgetProposal, error)
	// Creates, updates, or removes account budget proposals.  Operation statuses
	// are returned.
	MutateAccountBudgetProposal(ctx context.Context, in *MutateAccountBudgetProposalRequest, opts ...grpc.CallOption) (*MutateAccountBudgetProposalResponse, error)
}

type accountBudgetProposalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountBudgetProposalServiceClient(cc grpc.ClientConnInterface) AccountBudgetProposalServiceClient {
	return &accountBudgetProposalServiceClient{cc}
}

func (c *accountBudgetProposalServiceClient) GetAccountBudgetProposal(ctx context.Context, in *GetAccountBudgetProposalRequest, opts ...grpc.CallOption) (*resources.AccountBudgetProposal, error) {
	out := new(resources.AccountBudgetProposal)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AccountBudgetProposalService/GetAccountBudgetProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountBudgetProposalServiceClient) MutateAccountBudgetProposal(ctx context.Context, in *MutateAccountBudgetProposalRequest, opts ...grpc.CallOption) (*MutateAccountBudgetProposalResponse, error) {
	out := new(MutateAccountBudgetProposalResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AccountBudgetProposalService/MutateAccountBudgetProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountBudgetProposalServiceServer is the server API for AccountBudgetProposalService service.
type AccountBudgetProposalServiceServer interface {
	// Returns an account-level budget proposal in full detail.
	GetAccountBudgetProposal(context.Context, *GetAccountBudgetProposalRequest) (*resources.AccountBudgetProposal, error)
	// Creates, updates, or removes account budget proposals.  Operation statuses
	// are returned.
	MutateAccountBudgetProposal(context.Context, *MutateAccountBudgetProposalRequest) (*MutateAccountBudgetProposalResponse, error)
}

// UnimplementedAccountBudgetProposalServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccountBudgetProposalServiceServer struct {
}

func (*UnimplementedAccountBudgetProposalServiceServer) GetAccountBudgetProposal(context.Context, *GetAccountBudgetProposalRequest) (*resources.AccountBudgetProposal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountBudgetProposal not implemented")
}
func (*UnimplementedAccountBudgetProposalServiceServer) MutateAccountBudgetProposal(context.Context, *MutateAccountBudgetProposalRequest) (*MutateAccountBudgetProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAccountBudgetProposal not implemented")
}

func RegisterAccountBudgetProposalServiceServer(s *grpc.Server, srv AccountBudgetProposalServiceServer) {
	s.RegisterService(&_AccountBudgetProposalService_serviceDesc, srv)
}

func _AccountBudgetProposalService_GetAccountBudgetProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountBudgetProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountBudgetProposalServiceServer).GetAccountBudgetProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AccountBudgetProposalService/GetAccountBudgetProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountBudgetProposalServiceServer).GetAccountBudgetProposal(ctx, req.(*GetAccountBudgetProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountBudgetProposalService_MutateAccountBudgetProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAccountBudgetProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountBudgetProposalServiceServer).MutateAccountBudgetProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AccountBudgetProposalService/MutateAccountBudgetProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountBudgetProposalServiceServer).MutateAccountBudgetProposal(ctx, req.(*MutateAccountBudgetProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountBudgetProposalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.AccountBudgetProposalService",
	HandlerType: (*AccountBudgetProposalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountBudgetProposal",
			Handler:    _AccountBudgetProposalService_GetAccountBudgetProposal_Handler,
		},
		{
			MethodName: "MutateAccountBudgetProposal",
			Handler:    _AccountBudgetProposalService_MutateAccountBudgetProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/account_budget_proposal_service.proto",
}
