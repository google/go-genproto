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
// 	protoc        v3.11.2
// source: google/cloud/billing/budgets/v1beta1/budget_model.proto

package budgets

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	money "google.golang.org/genproto/googleapis/type/money"
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

// The type of basis used to determine if spend has passed the threshold.
type ThresholdRule_Basis int32

const (
	// Unspecified threshold basis.
	ThresholdRule_BASIS_UNSPECIFIED ThresholdRule_Basis = 0
	// Use current spend as the basis for comparison against the threshold.
	ThresholdRule_CURRENT_SPEND ThresholdRule_Basis = 1
	// Use forecasted spend for the period as the basis for comparison against
	// the threshold.
	ThresholdRule_FORECASTED_SPEND ThresholdRule_Basis = 2
)

// Enum value maps for ThresholdRule_Basis.
var (
	ThresholdRule_Basis_name = map[int32]string{
		0: "BASIS_UNSPECIFIED",
		1: "CURRENT_SPEND",
		2: "FORECASTED_SPEND",
	}
	ThresholdRule_Basis_value = map[string]int32{
		"BASIS_UNSPECIFIED": 0,
		"CURRENT_SPEND":     1,
		"FORECASTED_SPEND":  2,
	}
)

func (x ThresholdRule_Basis) Enum() *ThresholdRule_Basis {
	p := new(ThresholdRule_Basis)
	*p = x
	return p
}

func (x ThresholdRule_Basis) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ThresholdRule_Basis) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_billing_budgets_v1beta1_budget_model_proto_enumTypes[0].Descriptor()
}

func (ThresholdRule_Basis) Type() protoreflect.EnumType {
	return &file_google_cloud_billing_budgets_v1beta1_budget_model_proto_enumTypes[0]
}

func (x ThresholdRule_Basis) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ThresholdRule_Basis.Descriptor instead.
func (ThresholdRule_Basis) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDescGZIP(), []int{3, 0}
}

// Specifies how credits should be treated when determining spend for
// threshold calculations.
type Filter_CreditTypesTreatment int32

const (
	Filter_CREDIT_TYPES_TREATMENT_UNSPECIFIED Filter_CreditTypesTreatment = 0
	// All types of credit are subtracted from the gross cost to determine the
	// spend for threshold calculations.
	Filter_INCLUDE_ALL_CREDITS Filter_CreditTypesTreatment = 1
	// All types of credit are added to the net cost to determine the spend for
	// threshold calculations.
	Filter_EXCLUDE_ALL_CREDITS Filter_CreditTypesTreatment = 2
)

// Enum value maps for Filter_CreditTypesTreatment.
var (
	Filter_CreditTypesTreatment_name = map[int32]string{
		0: "CREDIT_TYPES_TREATMENT_UNSPECIFIED",
		1: "INCLUDE_ALL_CREDITS",
		2: "EXCLUDE_ALL_CREDITS",
	}
	Filter_CreditTypesTreatment_value = map[string]int32{
		"CREDIT_TYPES_TREATMENT_UNSPECIFIED": 0,
		"INCLUDE_ALL_CREDITS":                1,
		"EXCLUDE_ALL_CREDITS":                2,
	}
)

func (x Filter_CreditTypesTreatment) Enum() *Filter_CreditTypesTreatment {
	p := new(Filter_CreditTypesTreatment)
	*p = x
	return p
}

func (x Filter_CreditTypesTreatment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Filter_CreditTypesTreatment) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_billing_budgets_v1beta1_budget_model_proto_enumTypes[1].Descriptor()
}

func (Filter_CreditTypesTreatment) Type() protoreflect.EnumType {
	return &file_google_cloud_billing_budgets_v1beta1_budget_model_proto_enumTypes[1]
}

func (x Filter_CreditTypesTreatment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Filter_CreditTypesTreatment.Descriptor instead.
func (Filter_CreditTypesTreatment) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDescGZIP(), []int{5, 0}
}

// A budget is a plan that describes what you expect to spend on Cloud
// projects, plus the rules to execute as spend is tracked against that plan,
// (for example, send an alert when 90% of the target spend is met).
// Currently all plans are monthly budgets so the usage period(s) tracked are
// implied (calendar months of usage back-to-back).
type Budget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Resource name of the budget.
	// The resource name implies the scope of a budget. Values are of the form
	// `billingAccounts/{billingAccountId}/budgets/{budgetId}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// User data for display name in UI.
	// Validation: <= 60 chars.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. Filters that define which resources are used to compute
	// the actual spend against the budget.
	BudgetFilter *Filter `protobuf:"bytes,3,opt,name=budget_filter,json=budgetFilter,proto3" json:"budget_filter,omitempty"`
	// Required. Budgeted amount.
	Amount *BudgetAmount `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	// Optional. Rules that trigger alerts (notifications of thresholds
	// being crossed) when spend exceeds the specified percentages of the budget.
	ThresholdRules []*ThresholdRule `protobuf:"bytes,5,rep,name=threshold_rules,json=thresholdRules,proto3" json:"threshold_rules,omitempty"`
	// Optional. Rules to apply to all updates to the actual spend, regardless
	// of the thresholds set in `threshold_rules`.
	AllUpdatesRule *AllUpdatesRule `protobuf:"bytes,6,opt,name=all_updates_rule,json=allUpdatesRule,proto3" json:"all_updates_rule,omitempty"`
	// Optional. Etag to validate that the object is unchanged for a
	// read-modify-write operation.
	// An empty etag will cause an update to overwrite other changes.
	Etag string `protobuf:"bytes,7,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (x *Budget) Reset() {
	*x = Budget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Budget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Budget) ProtoMessage() {}

func (x *Budget) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Budget.ProtoReflect.Descriptor instead.
func (*Budget) Descriptor() ([]byte, []int) {
	return file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDescGZIP(), []int{0}
}

func (x *Budget) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Budget) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Budget) GetBudgetFilter() *Filter {
	if x != nil {
		return x.BudgetFilter
	}
	return nil
}

func (x *Budget) GetAmount() *BudgetAmount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Budget) GetThresholdRules() []*ThresholdRule {
	if x != nil {
		return x.ThresholdRules
	}
	return nil
}

func (x *Budget) GetAllUpdatesRule() *AllUpdatesRule {
	if x != nil {
		return x.AllUpdatesRule
	}
	return nil
}

func (x *Budget) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

// The budgeted amount for each usage period.
type BudgetAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specification for what amount to use as the budget.
	//
	// Types that are assignable to BudgetAmount:
	//	*BudgetAmount_SpecifiedAmount
	//	*BudgetAmount_LastPeriodAmount
	BudgetAmount isBudgetAmount_BudgetAmount `protobuf_oneof:"budget_amount"`
}

func (x *BudgetAmount) Reset() {
	*x = BudgetAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetAmount) ProtoMessage() {}

func (x *BudgetAmount) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetAmount.ProtoReflect.Descriptor instead.
func (*BudgetAmount) Descriptor() ([]byte, []int) {
	return file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDescGZIP(), []int{1}
}

func (m *BudgetAmount) GetBudgetAmount() isBudgetAmount_BudgetAmount {
	if m != nil {
		return m.BudgetAmount
	}
	return nil
}

func (x *BudgetAmount) GetSpecifiedAmount() *money.Money {
	if x, ok := x.GetBudgetAmount().(*BudgetAmount_SpecifiedAmount); ok {
		return x.SpecifiedAmount
	}
	return nil
}

func (x *BudgetAmount) GetLastPeriodAmount() *LastPeriodAmount {
	if x, ok := x.GetBudgetAmount().(*BudgetAmount_LastPeriodAmount); ok {
		return x.LastPeriodAmount
	}
	return nil
}

type isBudgetAmount_BudgetAmount interface {
	isBudgetAmount_BudgetAmount()
}

type BudgetAmount_SpecifiedAmount struct {
	// A specified amount to use as the budget.
	// `currency_code` is optional. If specified, it must match the
	// currency of the billing account. The `currency_code` is provided on
	// output.
	SpecifiedAmount *money.Money `protobuf:"bytes,1,opt,name=specified_amount,json=specifiedAmount,proto3,oneof"`
}

type BudgetAmount_LastPeriodAmount struct {
	// Use the last period's actual spend as the budget for the present period.
	LastPeriodAmount *LastPeriodAmount `protobuf:"bytes,2,opt,name=last_period_amount,json=lastPeriodAmount,proto3,oneof"`
}

func (*BudgetAmount_SpecifiedAmount) isBudgetAmount_BudgetAmount() {}

func (*BudgetAmount_LastPeriodAmount) isBudgetAmount_BudgetAmount() {}

// Describes a budget amount targeted to last period's spend.
// At this time, the amount is automatically 100% of last period's spend;
// that is, there are no other options yet.
// Future configuration will be described here (for example, configuring a
// percentage of last period's spend).
type LastPeriodAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LastPeriodAmount) Reset() {
	*x = LastPeriodAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LastPeriodAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastPeriodAmount) ProtoMessage() {}

func (x *LastPeriodAmount) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LastPeriodAmount.ProtoReflect.Descriptor instead.
func (*LastPeriodAmount) Descriptor() ([]byte, []int) {
	return file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDescGZIP(), []int{2}
}

// ThresholdRule contains a definition of a threshold which triggers
// an alert (a notification of a threshold being crossed) to be sent when
// spend goes above the specified amount.
// Alerts are automatically e-mailed to users with the Billing Account
// Administrator role or the Billing Account User role.
// The thresholds here have no effect on notifications sent to anything
// configured under `Budget.all_updates_rule`.
type ThresholdRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Send an alert when this threshold is exceeded.
	// This is a 1.0-based percentage, so 0.5 = 50%.
	// Validation: non-negative number.
	ThresholdPercent float64 `protobuf:"fixed64,1,opt,name=threshold_percent,json=thresholdPercent,proto3" json:"threshold_percent,omitempty"`
	// Optional. The type of basis used to determine if spend has passed the
	// threshold. Behavior defaults to CURRENT_SPEND if not set.
	SpendBasis ThresholdRule_Basis `protobuf:"varint,2,opt,name=spend_basis,json=spendBasis,proto3,enum=google.cloud.billing.budgets.v1beta1.ThresholdRule_Basis" json:"spend_basis,omitempty"`
}

func (x *ThresholdRule) Reset() {
	*x = ThresholdRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThresholdRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThresholdRule) ProtoMessage() {}

func (x *ThresholdRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThresholdRule.ProtoReflect.Descriptor instead.
func (*ThresholdRule) Descriptor() ([]byte, []int) {
	return file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDescGZIP(), []int{3}
}

func (x *ThresholdRule) GetThresholdPercent() float64 {
	if x != nil {
		return x.ThresholdPercent
	}
	return 0
}

func (x *ThresholdRule) GetSpendBasis() ThresholdRule_Basis {
	if x != nil {
		return x.SpendBasis
	}
	return ThresholdRule_BASIS_UNSPECIFIED
}

// AllUpdatesRule defines notifications that are sent on every update to the
// billing account's spend, regardless of the thresholds defined using
// threshold rules.
type AllUpdatesRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the Cloud Pub/Sub topic where budget related messages will be
	// published, in the form `projects/{project_id}/topics/{topic_id}`. Updates
	// are sent at regular intervals to the topic.
	// The topic needs to be created before the budget is created; see
	// https://cloud.google.com/billing/docs/how-to/budgets#manage-notifications
	// for more details.
	// Caller is expected to have
	// `pubsub.topics.setIamPolicy` permission on the topic when it's set for a
	// budget, otherwise, the API call will fail with PERMISSION_DENIED. See
	// https://cloud.google.com/pubsub/docs/access-control for more details on
	// Pub/Sub roles and permissions.
	PubsubTopic string `protobuf:"bytes,1,opt,name=pubsub_topic,json=pubsubTopic,proto3" json:"pubsub_topic,omitempty"`
	// Required. The schema version of the notification.
	// Only "1.0" is accepted. It represents the JSON schema as defined in
	// https://cloud.google.com/billing/docs/how-to/budgets#notification_format
	SchemaVersion string `protobuf:"bytes,2,opt,name=schema_version,json=schemaVersion,proto3" json:"schema_version,omitempty"`
}

func (x *AllUpdatesRule) Reset() {
	*x = AllUpdatesRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllUpdatesRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllUpdatesRule) ProtoMessage() {}

func (x *AllUpdatesRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllUpdatesRule.ProtoReflect.Descriptor instead.
func (*AllUpdatesRule) Descriptor() ([]byte, []int) {
	return file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDescGZIP(), []int{4}
}

func (x *AllUpdatesRule) GetPubsubTopic() string {
	if x != nil {
		return x.PubsubTopic
	}
	return ""
}

func (x *AllUpdatesRule) GetSchemaVersion() string {
	if x != nil {
		return x.SchemaVersion
	}
	return ""
}

// A filter for a budget, limiting the scope of the cost to calculate.
type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. A set of projects of the form `projects/{project}`,
	// specifying that usage from only this set of projects should be
	// included in the budget. If omitted, the report will include all usage for
	// the billing account, regardless of which project the usage occurred on.
	// Only zero or one project can be specified currently.
	Projects []string `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	// Optional. If not set, default behavior is `INCLUDE_ALL_CREDITS`.
	CreditTypesTreatment Filter_CreditTypesTreatment `protobuf:"varint,4,opt,name=credit_types_treatment,json=creditTypesTreatment,proto3,enum=google.cloud.billing.budgets.v1beta1.Filter_CreditTypesTreatment" json:"credit_types_treatment,omitempty"`
	// Optional. A set of services of the form `services/{service_id}`,
	// specifying that usage from only this set of services should be
	// included in the budget. If omitted, the report will include usage for
	// all the services.
	// The service names are available through the Catalog API:
	// https://cloud.google.com/billing/v1/how-tos/catalog-api.
	Services []string `protobuf:"bytes,3,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDescGZIP(), []int{5}
}

func (x *Filter) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *Filter) GetCreditTypesTreatment() Filter_CreditTypesTreatment {
	if x != nil {
		return x.CreditTypesTreatment
	}
	return Filter_CREDIT_TYPES_TREATMENT_UNSPECIFIED
}

func (x *Filter) GetServices() []string {
	if x != nil {
		return x.Services
	}
	return nil
}

var File_google_cloud_billing_budgets_v1beta1_budget_model_proto protoreflect.FileDescriptor

var file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x04, 0x0a, 0x06, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x56, 0x0a, 0x0d, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0c, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x4f, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x61, 0x0a, 0x0f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x75,
	0x6c, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0e, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x63, 0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x73, 0x52, 0x75, 0x6c, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0e, 0x61, 0x6c,
	0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x04,
	0x65, 0x74, 0x61, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x04, 0x65, 0x74, 0x61, 0x67, 0x3a, 0x5d, 0xea, 0x41, 0x5a, 0x0a, 0x24, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x12, 0x32, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x2f, 0x7b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x7d, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x62, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x7d, 0x22, 0xc8, 0x01, 0x0a, 0x0c, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x10, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d, 0x6f,
	0x6e, 0x65, 0x79, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x66, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x10, 0x6c, 0x61,
	0x73, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0f,
	0x0a, 0x0d, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x12, 0x0a, 0x10, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0xeb, 0x01, 0x0a, 0x0d, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x11, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x10, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x5f, 0x0a, 0x0b, 0x73, 0x70, 0x65, 0x6e, 0x64,
	0x5f, 0x62, 0x61, 0x73, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x75, 0x6c,
	0x65, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x73, 0x70,
	0x65, 0x6e, 0x64, 0x42, 0x61, 0x73, 0x69, 0x73, 0x22, 0x47, 0x0a, 0x05, 0x42, 0x61, 0x73, 0x69,
	0x73, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x41, 0x53, 0x49, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x55, 0x52, 0x52,
	0x45, 0x4e, 0x54, 0x5f, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x46,
	0x4f, 0x52, 0x45, 0x43, 0x41, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x10,
	0x02, 0x22, 0x64, 0x0a, 0x0e, 0x41, 0x6c, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b,
	0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x2a, 0x0a, 0x0e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xba, 0x02, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x12, 0x7c, 0x0a, 0x16, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x5f, 0x74, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x54, 0x72, 0x65,
	0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x14, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x54, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x22, 0x70, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x54, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x22, 0x43, 0x52,
	0x45, 0x44, 0x49, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x53, 0x5f, 0x54, 0x52, 0x45, 0x41, 0x54,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x5f, 0x41, 0x4c,
	0x4c, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x49, 0x54, 0x53, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x45,
	0x58, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x49,
	0x54, 0x53, 0x10, 0x02, 0x42, 0x79, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDescOnce sync.Once
	file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDescData = file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDesc
)

func file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDescGZIP() []byte {
	file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDescOnce.Do(func() {
		file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDescData)
	})
	return file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDescData
}

var file_google_cloud_billing_budgets_v1beta1_budget_model_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_billing_budgets_v1beta1_budget_model_proto_goTypes = []interface{}{
	(ThresholdRule_Basis)(0),         // 0: google.cloud.billing.budgets.v1beta1.ThresholdRule.Basis
	(Filter_CreditTypesTreatment)(0), // 1: google.cloud.billing.budgets.v1beta1.Filter.CreditTypesTreatment
	(*Budget)(nil),                   // 2: google.cloud.billing.budgets.v1beta1.Budget
	(*BudgetAmount)(nil),             // 3: google.cloud.billing.budgets.v1beta1.BudgetAmount
	(*LastPeriodAmount)(nil),         // 4: google.cloud.billing.budgets.v1beta1.LastPeriodAmount
	(*ThresholdRule)(nil),            // 5: google.cloud.billing.budgets.v1beta1.ThresholdRule
	(*AllUpdatesRule)(nil),           // 6: google.cloud.billing.budgets.v1beta1.AllUpdatesRule
	(*Filter)(nil),                   // 7: google.cloud.billing.budgets.v1beta1.Filter
	(*money.Money)(nil),              // 8: google.type.Money
}
var file_google_cloud_billing_budgets_v1beta1_budget_model_proto_depIdxs = []int32{
	7, // 0: google.cloud.billing.budgets.v1beta1.Budget.budget_filter:type_name -> google.cloud.billing.budgets.v1beta1.Filter
	3, // 1: google.cloud.billing.budgets.v1beta1.Budget.amount:type_name -> google.cloud.billing.budgets.v1beta1.BudgetAmount
	5, // 2: google.cloud.billing.budgets.v1beta1.Budget.threshold_rules:type_name -> google.cloud.billing.budgets.v1beta1.ThresholdRule
	6, // 3: google.cloud.billing.budgets.v1beta1.Budget.all_updates_rule:type_name -> google.cloud.billing.budgets.v1beta1.AllUpdatesRule
	8, // 4: google.cloud.billing.budgets.v1beta1.BudgetAmount.specified_amount:type_name -> google.type.Money
	4, // 5: google.cloud.billing.budgets.v1beta1.BudgetAmount.last_period_amount:type_name -> google.cloud.billing.budgets.v1beta1.LastPeriodAmount
	0, // 6: google.cloud.billing.budgets.v1beta1.ThresholdRule.spend_basis:type_name -> google.cloud.billing.budgets.v1beta1.ThresholdRule.Basis
	1, // 7: google.cloud.billing.budgets.v1beta1.Filter.credit_types_treatment:type_name -> google.cloud.billing.budgets.v1beta1.Filter.CreditTypesTreatment
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_cloud_billing_budgets_v1beta1_budget_model_proto_init() }
func file_google_cloud_billing_budgets_v1beta1_budget_model_proto_init() {
	if File_google_cloud_billing_budgets_v1beta1_budget_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Budget); i {
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
		file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BudgetAmount); i {
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
		file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LastPeriodAmount); i {
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
		file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThresholdRule); i {
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
		file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllUpdatesRule); i {
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
		file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
	file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*BudgetAmount_SpecifiedAmount)(nil),
		(*BudgetAmount_LastPeriodAmount)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_billing_budgets_v1beta1_budget_model_proto_goTypes,
		DependencyIndexes: file_google_cloud_billing_budgets_v1beta1_budget_model_proto_depIdxs,
		EnumInfos:         file_google_cloud_billing_budgets_v1beta1_budget_model_proto_enumTypes,
		MessageInfos:      file_google_cloud_billing_budgets_v1beta1_budget_model_proto_msgTypes,
	}.Build()
	File_google_cloud_billing_budgets_v1beta1_budget_model_proto = out.File
	file_google_cloud_billing_budgets_v1beta1_budget_model_proto_rawDesc = nil
	file_google_cloud_billing_budgets_v1beta1_budget_model_proto_goTypes = nil
	file_google_cloud_billing_budgets_v1beta1_budget_model_proto_depIdxs = nil
}
