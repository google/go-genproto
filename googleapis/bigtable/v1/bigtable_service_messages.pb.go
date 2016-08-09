// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/bigtable/v1/bigtable_service_messages.proto
// DO NOT EDIT!

package google_bigtable_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_rpc "google.golang.org/genproto/googleapis/rpc/status"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Request message for BigtableServer.ReadRows.
type ReadRowsRequest struct {
	// The unique name of the table from which to read.
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	// If neither row_key nor row_range is set, reads from all rows.
	//
	// Types that are valid to be assigned to Target:
	//	*ReadRowsRequest_RowKey
	//	*ReadRowsRequest_RowRange
	//	*ReadRowsRequest_RowSet
	Target isReadRowsRequest_Target `protobuf_oneof:"target"`
	// The filter to apply to the contents of the specified row(s). If unset,
	// reads the entire table.
	Filter *RowFilter `protobuf:"bytes,5,opt,name=filter" json:"filter,omitempty"`
	// By default, rows are read sequentially, producing results which are
	// guaranteed to arrive in increasing row order. Setting
	// "allow_row_interleaving" to true allows multiple rows to be interleaved in
	// the response stream, which increases throughput but breaks this guarantee,
	// and may force the client to use more memory to buffer partially-received
	// rows. Cannot be set to true when specifying "num_rows_limit".
	AllowRowInterleaving bool `protobuf:"varint,6,opt,name=allow_row_interleaving,json=allowRowInterleaving" json:"allow_row_interleaving,omitempty"`
	// The read will terminate after committing to N rows' worth of results. The
	// default (zero) is to return all results.
	// Note that "allow_row_interleaving" cannot be set to true when this is set.
	NumRowsLimit int64 `protobuf:"varint,7,opt,name=num_rows_limit,json=numRowsLimit" json:"num_rows_limit,omitempty"`
}

func (m *ReadRowsRequest) Reset()                    { *m = ReadRowsRequest{} }
func (m *ReadRowsRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRowsRequest) ProtoMessage()               {}
func (*ReadRowsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type isReadRowsRequest_Target interface {
	isReadRowsRequest_Target()
}

type ReadRowsRequest_RowKey struct {
	RowKey []byte `protobuf:"bytes,2,opt,name=row_key,json=rowKey,proto3,oneof"`
}
type ReadRowsRequest_RowRange struct {
	RowRange *RowRange `protobuf:"bytes,3,opt,name=row_range,json=rowRange,oneof"`
}
type ReadRowsRequest_RowSet struct {
	RowSet *RowSet `protobuf:"bytes,8,opt,name=row_set,json=rowSet,oneof"`
}

func (*ReadRowsRequest_RowKey) isReadRowsRequest_Target()   {}
func (*ReadRowsRequest_RowRange) isReadRowsRequest_Target() {}
func (*ReadRowsRequest_RowSet) isReadRowsRequest_Target()   {}

func (m *ReadRowsRequest) GetTarget() isReadRowsRequest_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *ReadRowsRequest) GetRowKey() []byte {
	if x, ok := m.GetTarget().(*ReadRowsRequest_RowKey); ok {
		return x.RowKey
	}
	return nil
}

func (m *ReadRowsRequest) GetRowRange() *RowRange {
	if x, ok := m.GetTarget().(*ReadRowsRequest_RowRange); ok {
		return x.RowRange
	}
	return nil
}

func (m *ReadRowsRequest) GetRowSet() *RowSet {
	if x, ok := m.GetTarget().(*ReadRowsRequest_RowSet); ok {
		return x.RowSet
	}
	return nil
}

func (m *ReadRowsRequest) GetFilter() *RowFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ReadRowsRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ReadRowsRequest_OneofMarshaler, _ReadRowsRequest_OneofUnmarshaler, _ReadRowsRequest_OneofSizer, []interface{}{
		(*ReadRowsRequest_RowKey)(nil),
		(*ReadRowsRequest_RowRange)(nil),
		(*ReadRowsRequest_RowSet)(nil),
	}
}

func _ReadRowsRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ReadRowsRequest)
	// target
	switch x := m.Target.(type) {
	case *ReadRowsRequest_RowKey:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.RowKey)
	case *ReadRowsRequest_RowRange:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RowRange); err != nil {
			return err
		}
	case *ReadRowsRequest_RowSet:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RowSet); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ReadRowsRequest.Target has unexpected type %T", x)
	}
	return nil
}

func _ReadRowsRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ReadRowsRequest)
	switch tag {
	case 2: // target.row_key
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Target = &ReadRowsRequest_RowKey{x}
		return true, err
	case 3: // target.row_range
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RowRange)
		err := b.DecodeMessage(msg)
		m.Target = &ReadRowsRequest_RowRange{msg}
		return true, err
	case 8: // target.row_set
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RowSet)
		err := b.DecodeMessage(msg)
		m.Target = &ReadRowsRequest_RowSet{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ReadRowsRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ReadRowsRequest)
	// target
	switch x := m.Target.(type) {
	case *ReadRowsRequest_RowKey:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RowKey)))
		n += len(x.RowKey)
	case *ReadRowsRequest_RowRange:
		s := proto.Size(x.RowRange)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ReadRowsRequest_RowSet:
		s := proto.Size(x.RowSet)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for BigtableService.ReadRows.
type ReadRowsResponse struct {
	// The key of the row for which we're receiving data.
	// Results will be received in increasing row key order, unless
	// "allow_row_interleaving" was specified in the request.
	RowKey []byte `protobuf:"bytes,1,opt,name=row_key,json=rowKey,proto3" json:"row_key,omitempty"`
	// One or more chunks of the row specified by "row_key".
	Chunks []*ReadRowsResponse_Chunk `protobuf:"bytes,2,rep,name=chunks" json:"chunks,omitempty"`
}

func (m *ReadRowsResponse) Reset()                    { *m = ReadRowsResponse{} }
func (m *ReadRowsResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadRowsResponse) ProtoMessage()               {}
func (*ReadRowsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ReadRowsResponse) GetChunks() []*ReadRowsResponse_Chunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

// Specifies a piece of a row's contents returned as part of the read
// response stream.
type ReadRowsResponse_Chunk struct {
	// Types that are valid to be assigned to Chunk:
	//	*ReadRowsResponse_Chunk_RowContents
	//	*ReadRowsResponse_Chunk_ResetRow
	//	*ReadRowsResponse_Chunk_CommitRow
	Chunk isReadRowsResponse_Chunk_Chunk `protobuf_oneof:"chunk"`
}

func (m *ReadRowsResponse_Chunk) Reset()                    { *m = ReadRowsResponse_Chunk{} }
func (m *ReadRowsResponse_Chunk) String() string            { return proto.CompactTextString(m) }
func (*ReadRowsResponse_Chunk) ProtoMessage()               {}
func (*ReadRowsResponse_Chunk) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

type isReadRowsResponse_Chunk_Chunk interface {
	isReadRowsResponse_Chunk_Chunk()
}

type ReadRowsResponse_Chunk_RowContents struct {
	RowContents *Family `protobuf:"bytes,1,opt,name=row_contents,json=rowContents,oneof"`
}
type ReadRowsResponse_Chunk_ResetRow struct {
	ResetRow bool `protobuf:"varint,2,opt,name=reset_row,json=resetRow,oneof"`
}
type ReadRowsResponse_Chunk_CommitRow struct {
	CommitRow bool `protobuf:"varint,3,opt,name=commit_row,json=commitRow,oneof"`
}

func (*ReadRowsResponse_Chunk_RowContents) isReadRowsResponse_Chunk_Chunk() {}
func (*ReadRowsResponse_Chunk_ResetRow) isReadRowsResponse_Chunk_Chunk()    {}
func (*ReadRowsResponse_Chunk_CommitRow) isReadRowsResponse_Chunk_Chunk()   {}

func (m *ReadRowsResponse_Chunk) GetChunk() isReadRowsResponse_Chunk_Chunk {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func (m *ReadRowsResponse_Chunk) GetRowContents() *Family {
	if x, ok := m.GetChunk().(*ReadRowsResponse_Chunk_RowContents); ok {
		return x.RowContents
	}
	return nil
}

func (m *ReadRowsResponse_Chunk) GetResetRow() bool {
	if x, ok := m.GetChunk().(*ReadRowsResponse_Chunk_ResetRow); ok {
		return x.ResetRow
	}
	return false
}

func (m *ReadRowsResponse_Chunk) GetCommitRow() bool {
	if x, ok := m.GetChunk().(*ReadRowsResponse_Chunk_CommitRow); ok {
		return x.CommitRow
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ReadRowsResponse_Chunk) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ReadRowsResponse_Chunk_OneofMarshaler, _ReadRowsResponse_Chunk_OneofUnmarshaler, _ReadRowsResponse_Chunk_OneofSizer, []interface{}{
		(*ReadRowsResponse_Chunk_RowContents)(nil),
		(*ReadRowsResponse_Chunk_ResetRow)(nil),
		(*ReadRowsResponse_Chunk_CommitRow)(nil),
	}
}

func _ReadRowsResponse_Chunk_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ReadRowsResponse_Chunk)
	// chunk
	switch x := m.Chunk.(type) {
	case *ReadRowsResponse_Chunk_RowContents:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RowContents); err != nil {
			return err
		}
	case *ReadRowsResponse_Chunk_ResetRow:
		t := uint64(0)
		if x.ResetRow {
			t = 1
		}
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *ReadRowsResponse_Chunk_CommitRow:
		t := uint64(0)
		if x.CommitRow {
			t = 1
		}
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("ReadRowsResponse_Chunk.Chunk has unexpected type %T", x)
	}
	return nil
}

func _ReadRowsResponse_Chunk_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ReadRowsResponse_Chunk)
	switch tag {
	case 1: // chunk.row_contents
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Family)
		err := b.DecodeMessage(msg)
		m.Chunk = &ReadRowsResponse_Chunk_RowContents{msg}
		return true, err
	case 2: // chunk.reset_row
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Chunk = &ReadRowsResponse_Chunk_ResetRow{x != 0}
		return true, err
	case 3: // chunk.commit_row
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Chunk = &ReadRowsResponse_Chunk_CommitRow{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _ReadRowsResponse_Chunk_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ReadRowsResponse_Chunk)
	// chunk
	switch x := m.Chunk.(type) {
	case *ReadRowsResponse_Chunk_RowContents:
		s := proto.Size(x.RowContents)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ReadRowsResponse_Chunk_ResetRow:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += 1
	case *ReadRowsResponse_Chunk_CommitRow:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Request message for BigtableService.SampleRowKeys.
type SampleRowKeysRequest struct {
	// The unique name of the table from which to sample row keys.
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
}

func (m *SampleRowKeysRequest) Reset()                    { *m = SampleRowKeysRequest{} }
func (m *SampleRowKeysRequest) String() string            { return proto.CompactTextString(m) }
func (*SampleRowKeysRequest) ProtoMessage()               {}
func (*SampleRowKeysRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

// Response message for BigtableService.SampleRowKeys.
type SampleRowKeysResponse struct {
	// Sorted streamed sequence of sample row keys in the table. The table might
	// have contents before the first row key in the list and after the last one,
	// but a key containing the empty string indicates "end of table" and will be
	// the last response given, if present.
	// Note that row keys in this list may not have ever been written to or read
	// from, and users should therefore not make any assumptions about the row key
	// structure that are specific to their use case.
	RowKey []byte `protobuf:"bytes,1,opt,name=row_key,json=rowKey,proto3" json:"row_key,omitempty"`
	// Approximate total storage space used by all rows in the table which precede
	// "row_key". Buffering the contents of all rows between two subsequent
	// samples would require space roughly equal to the difference in their
	// "offset_bytes" fields.
	OffsetBytes int64 `protobuf:"varint,2,opt,name=offset_bytes,json=offsetBytes" json:"offset_bytes,omitempty"`
}

func (m *SampleRowKeysResponse) Reset()                    { *m = SampleRowKeysResponse{} }
func (m *SampleRowKeysResponse) String() string            { return proto.CompactTextString(m) }
func (*SampleRowKeysResponse) ProtoMessage()               {}
func (*SampleRowKeysResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

// Request message for BigtableService.MutateRow.
type MutateRowRequest struct {
	// The unique name of the table to which the mutation should be applied.
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	// The key of the row to which the mutation should be applied.
	RowKey []byte `protobuf:"bytes,2,opt,name=row_key,json=rowKey,proto3" json:"row_key,omitempty"`
	// Changes to be atomically applied to the specified row. Entries are applied
	// in order, meaning that earlier mutations can be masked by later ones.
	// Must contain at least one entry and at most 100000.
	Mutations []*Mutation `protobuf:"bytes,3,rep,name=mutations" json:"mutations,omitempty"`
}

func (m *MutateRowRequest) Reset()                    { *m = MutateRowRequest{} }
func (m *MutateRowRequest) String() string            { return proto.CompactTextString(m) }
func (*MutateRowRequest) ProtoMessage()               {}
func (*MutateRowRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *MutateRowRequest) GetMutations() []*Mutation {
	if m != nil {
		return m.Mutations
	}
	return nil
}

// Request message for BigtableService.MutateRows.
type MutateRowsRequest struct {
	// The unique name of the table to which the mutations should be applied.
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	// The row keys/mutations to be applied in bulk.
	// Each entry is applied as an atomic mutation, but the entries may be
	// applied in arbitrary order (even between entries for the same row).
	// At least one entry must be specified, and in total the entries may
	// contain at most 100000 mutations.
	Entries []*MutateRowsRequest_Entry `protobuf:"bytes,2,rep,name=entries" json:"entries,omitempty"`
}

func (m *MutateRowsRequest) Reset()                    { *m = MutateRowsRequest{} }
func (m *MutateRowsRequest) String() string            { return proto.CompactTextString(m) }
func (*MutateRowsRequest) ProtoMessage()               {}
func (*MutateRowsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *MutateRowsRequest) GetEntries() []*MutateRowsRequest_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type MutateRowsRequest_Entry struct {
	// The key of the row to which the `mutations` should be applied.
	RowKey []byte `protobuf:"bytes,1,opt,name=row_key,json=rowKey,proto3" json:"row_key,omitempty"`
	// Changes to be atomically applied to the specified row. Mutations are
	// applied in order, meaning that earlier mutations can be masked by
	// later ones.
	// At least one mutation must be specified.
	Mutations []*Mutation `protobuf:"bytes,2,rep,name=mutations" json:"mutations,omitempty"`
}

func (m *MutateRowsRequest_Entry) Reset()                    { *m = MutateRowsRequest_Entry{} }
func (m *MutateRowsRequest_Entry) String() string            { return proto.CompactTextString(m) }
func (*MutateRowsRequest_Entry) ProtoMessage()               {}
func (*MutateRowsRequest_Entry) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5, 0} }

func (m *MutateRowsRequest_Entry) GetMutations() []*Mutation {
	if m != nil {
		return m.Mutations
	}
	return nil
}

// Response message for BigtableService.MutateRows.
type MutateRowsResponse struct {
	// The results for each Entry from the request, presented in the order
	// in which the entries were originally given.
	// Depending on how requests are batched during execution, it is possible
	// for one Entry to fail due to an error with another Entry. In the event
	// that this occurs, the same error will be reported for both entries.
	Statuses []*google_rpc.Status `protobuf:"bytes,1,rep,name=statuses" json:"statuses,omitempty"`
}

func (m *MutateRowsResponse) Reset()                    { *m = MutateRowsResponse{} }
func (m *MutateRowsResponse) String() string            { return proto.CompactTextString(m) }
func (*MutateRowsResponse) ProtoMessage()               {}
func (*MutateRowsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *MutateRowsResponse) GetStatuses() []*google_rpc.Status {
	if m != nil {
		return m.Statuses
	}
	return nil
}

// Request message for BigtableService.CheckAndMutateRowRequest
type CheckAndMutateRowRequest struct {
	// The unique name of the table to which the conditional mutation should be
	// applied.
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	// The key of the row to which the conditional mutation should be applied.
	RowKey []byte `protobuf:"bytes,2,opt,name=row_key,json=rowKey,proto3" json:"row_key,omitempty"`
	// The filter to be applied to the contents of the specified row. Depending
	// on whether or not any results are yielded, either "true_mutations" or
	// "false_mutations" will be executed. If unset, checks that the row contains
	// any values at all.
	PredicateFilter *RowFilter `protobuf:"bytes,6,opt,name=predicate_filter,json=predicateFilter" json:"predicate_filter,omitempty"`
	// Changes to be atomically applied to the specified row if "predicate_filter"
	// yields at least one cell when applied to "row_key". Entries are applied in
	// order, meaning that earlier mutations can be masked by later ones.
	// Must contain at least one entry if "false_mutations" is empty, and at most
	// 100000.
	TrueMutations []*Mutation `protobuf:"bytes,4,rep,name=true_mutations,json=trueMutations" json:"true_mutations,omitempty"`
	// Changes to be atomically applied to the specified row if "predicate_filter"
	// does not yield any cells when applied to "row_key". Entries are applied in
	// order, meaning that earlier mutations can be masked by later ones.
	// Must contain at least one entry if "true_mutations" is empty, and at most
	// 100000.
	FalseMutations []*Mutation `protobuf:"bytes,5,rep,name=false_mutations,json=falseMutations" json:"false_mutations,omitempty"`
}

func (m *CheckAndMutateRowRequest) Reset()                    { *m = CheckAndMutateRowRequest{} }
func (m *CheckAndMutateRowRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckAndMutateRowRequest) ProtoMessage()               {}
func (*CheckAndMutateRowRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *CheckAndMutateRowRequest) GetPredicateFilter() *RowFilter {
	if m != nil {
		return m.PredicateFilter
	}
	return nil
}

func (m *CheckAndMutateRowRequest) GetTrueMutations() []*Mutation {
	if m != nil {
		return m.TrueMutations
	}
	return nil
}

func (m *CheckAndMutateRowRequest) GetFalseMutations() []*Mutation {
	if m != nil {
		return m.FalseMutations
	}
	return nil
}

// Response message for BigtableService.CheckAndMutateRowRequest.
type CheckAndMutateRowResponse struct {
	// Whether or not the request's "predicate_filter" yielded any results for
	// the specified row.
	PredicateMatched bool `protobuf:"varint,1,opt,name=predicate_matched,json=predicateMatched" json:"predicate_matched,omitempty"`
}

func (m *CheckAndMutateRowResponse) Reset()                    { *m = CheckAndMutateRowResponse{} }
func (m *CheckAndMutateRowResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckAndMutateRowResponse) ProtoMessage()               {}
func (*CheckAndMutateRowResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

// Request message for BigtableService.ReadModifyWriteRowRequest.
type ReadModifyWriteRowRequest struct {
	// The unique name of the table to which the read/modify/write rules should be
	// applied.
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	// The key of the row to which the read/modify/write rules should be applied.
	RowKey []byte `protobuf:"bytes,2,opt,name=row_key,json=rowKey,proto3" json:"row_key,omitempty"`
	// Rules specifying how the specified row's contents are to be transformed
	// into writes. Entries are applied in order, meaning that earlier rules will
	// affect the results of later ones.
	Rules []*ReadModifyWriteRule `protobuf:"bytes,3,rep,name=rules" json:"rules,omitempty"`
}

func (m *ReadModifyWriteRowRequest) Reset()                    { *m = ReadModifyWriteRowRequest{} }
func (m *ReadModifyWriteRowRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadModifyWriteRowRequest) ProtoMessage()               {}
func (*ReadModifyWriteRowRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *ReadModifyWriteRowRequest) GetRules() []*ReadModifyWriteRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterType((*ReadRowsRequest)(nil), "google.bigtable.v1.ReadRowsRequest")
	proto.RegisterType((*ReadRowsResponse)(nil), "google.bigtable.v1.ReadRowsResponse")
	proto.RegisterType((*ReadRowsResponse_Chunk)(nil), "google.bigtable.v1.ReadRowsResponse.Chunk")
	proto.RegisterType((*SampleRowKeysRequest)(nil), "google.bigtable.v1.SampleRowKeysRequest")
	proto.RegisterType((*SampleRowKeysResponse)(nil), "google.bigtable.v1.SampleRowKeysResponse")
	proto.RegisterType((*MutateRowRequest)(nil), "google.bigtable.v1.MutateRowRequest")
	proto.RegisterType((*MutateRowsRequest)(nil), "google.bigtable.v1.MutateRowsRequest")
	proto.RegisterType((*MutateRowsRequest_Entry)(nil), "google.bigtable.v1.MutateRowsRequest.Entry")
	proto.RegisterType((*MutateRowsResponse)(nil), "google.bigtable.v1.MutateRowsResponse")
	proto.RegisterType((*CheckAndMutateRowRequest)(nil), "google.bigtable.v1.CheckAndMutateRowRequest")
	proto.RegisterType((*CheckAndMutateRowResponse)(nil), "google.bigtable.v1.CheckAndMutateRowResponse")
	proto.RegisterType((*ReadModifyWriteRowRequest)(nil), "google.bigtable.v1.ReadModifyWriteRowRequest")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/bigtable/v1/bigtable_service_messages.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 770 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x3e, 0x49, 0x4e, 0xfe, 0xa6, 0x39, 0xfd, 0xb1, 0x7a, 0x5a, 0x37, 0x6a, 0x75, 0x0e, 0x16,
	0x12, 0x88, 0x4a, 0x8e, 0x0a, 0x54, 0x42, 0xa0, 0x0a, 0x91, 0xd0, 0xaa, 0x08, 0x82, 0xca, 0xe6,
	0x82, 0x1b, 0xa4, 0x68, 0xe3, 0x6c, 0x5c, 0xab, 0xfe, 0x09, 0xbb, 0x9b, 0x54, 0xb9, 0x46, 0xe2,
	0x1e, 0x9e, 0x82, 0x37, 0xe2, 0x86, 0x87, 0x61, 0x76, 0xbd, 0x71, 0x42, 0x69, 0x44, 0x10, 0xbd,
	0x88, 0x62, 0xcf, 0xcc, 0xf7, 0xcd, 0xcc, 0xb7, 0xe3, 0x59, 0x78, 0xe3, 0x27, 0x89, 0x1f, 0x32,
	0xd7, 0x4f, 0x42, 0x1a, 0xfb, 0x6e, 0xc2, 0xfd, 0x86, 0xcf, 0xe2, 0x21, 0x4f, 0x64, 0xd2, 0x48,
	0x5d, 0x74, 0x18, 0x88, 0x46, 0x2f, 0xf0, 0x25, 0xed, 0x85, 0xac, 0x31, 0x3e, 0xc8, 0x9e, 0xbb,
	0x82, 0xf1, 0x71, 0xe0, 0xb1, 0x6e, 0xc4, 0x84, 0xa0, 0x3e, 0x13, 0xae, 0x86, 0x59, 0x96, 0xa1,
	0x9c, 0xc6, 0xb9, 0xe3, 0x83, 0xfa, 0xf1, 0x1f, 0xa4, 0xe9, 0x53, 0x49, 0x53, 0xea, 0xfa, 0xd1,
	0x72, 0x34, 0x7c, 0xe8, 0x35, 0x84, 0xa4, 0x72, 0x24, 0xcc, 0x5f, 0x0a, 0x77, 0xbe, 0xe5, 0x61,
	0x8d, 0x30, 0xda, 0x27, 0xc9, 0xa5, 0x20, 0xec, 0xfd, 0x88, 0x09, 0x69, 0xed, 0x01, 0xa4, 0x69,
	0x62, 0x1a, 0x31, 0x3b, 0xf7, 0x7f, 0xee, 0x6e, 0x95, 0x54, 0xb5, 0xe5, 0x35, 0x1a, 0xac, 0x1d,
	0x28, 0xf3, 0xe4, 0xb2, 0x7b, 0xc1, 0x26, 0x76, 0x1e, 0x7d, 0xb5, 0xd3, 0xbf, 0x48, 0x09, 0x0d,
	0x2f, 0xd9, 0xc4, 0x7a, 0x02, 0x55, 0xe5, 0xe2, 0x58, 0x09, 0xb3, 0x0b, 0xe8, 0x5c, 0xb9, 0xbf,
	0xeb, 0xfe, 0xdc, 0xbb, 0x8b, 0xd9, 0x88, 0x8a, 0x41, 0x68, 0x85, 0x9b, 0x67, 0xeb, 0x30, 0xe5,
	0x15, 0x4c, 0xda, 0x15, 0x0d, 0xad, 0x2f, 0x80, 0x76, 0x98, 0x34, 0x39, 0xf1, 0x09, 0x61, 0xa5,
	0x41, 0x10, 0x4a, 0xc6, 0xed, 0xa2, 0x46, 0xed, 0x2d, 0x40, 0x9d, 0xe8, 0x20, 0x62, 0x82, 0xad,
	0x87, 0xb0, 0x45, 0xc3, 0x50, 0x15, 0x8b, 0xbf, 0x20, 0x46, 0x13, 0x2a, 0x35, 0x0e, 0x62, 0xdf,
	0x2e, 0x21, 0x4d, 0x85, 0x6c, 0x6a, 0x2f, 0xe2, 0x5e, 0xcc, 0xf9, 0xac, 0xdb, 0xb0, 0x1a, 0x8f,
	0x22, 0x85, 0x11, 0xdd, 0x30, 0x88, 0x02, 0x69, 0x97, 0x31, 0xba, 0x40, 0x6a, 0x68, 0x55, 0x12,
	0xbe, 0x52, 0xb6, 0x66, 0x05, 0x4a, 0x92, 0x72, 0x9f, 0x49, 0xe7, 0x43, 0x1e, 0xd6, 0x67, 0xf2,
	0x8a, 0x61, 0x12, 0x0b, 0x66, 0x6d, 0xcf, 0x04, 0x54, 0xe2, 0xd6, 0x32, 0xf9, 0x9a, 0x50, 0xf2,
	0xce, 0x47, 0xf1, 0x85, 0x40, 0x61, 0x0b, 0xd8, 0xca, 0xbd, 0x6b, 0x5b, 0xb9, 0x42, 0xe7, 0xb6,
	0x14, 0x84, 0x18, 0x64, 0xfd, 0x53, 0x0e, 0x8a, 0xda, 0x62, 0x3d, 0x85, 0x9a, 0x4a, 0xe3, 0x25,
	0xd8, 0x40, 0x2c, 0x85, 0xce, 0xb5, 0x40, 0xd4, 0x13, 0x1a, 0x05, 0xe1, 0x04, 0x45, 0x5d, 0x41,
	0x44, 0xcb, 0x00, 0x70, 0x0e, 0xaa, 0x9c, 0xe1, 0x71, 0xa8, 0x76, 0xf5, 0x51, 0x57, 0xf4, 0x79,
	0x29, 0x13, 0x16, 0x60, 0xfd, 0x07, 0xe0, 0x25, 0x11, 0xf6, 0xab, 0xfd, 0x05, 0xe3, 0xaf, 0xa6,
	0x36, 0x0c, 0x68, 0x96, 0xa1, 0xa8, 0x8b, 0x72, 0x0e, 0x61, 0xb3, 0x43, 0xa3, 0x61, 0xc8, 0x88,
	0xee, 0x73, 0xc9, 0x41, 0x73, 0x3a, 0xf0, 0xef, 0x15, 0xd8, 0xaf, 0x04, 0xbc, 0x05, 0xb5, 0x64,
	0x30, 0x50, 0x25, 0xf7, 0x26, 0x92, 0x09, 0x5d, 0x74, 0x81, 0xac, 0xa4, 0xb6, 0xa6, 0x32, 0x39,
	0x1f, 0x73, 0xb0, 0xde, 0x1e, 0xe1, 0x27, 0xa0, 0x58, 0x97, 0x9c, 0xf8, 0xed, 0x2b, 0x13, 0x9f,
	0xe5, 0x7b, 0x0c, 0xd5, 0x48, 0x71, 0x05, 0x58, 0x16, 0x2a, 0x50, 0x58, 0x34, 0xef, 0x6d, 0x13,
	0x44, 0x66, 0xe1, 0xce, 0xd7, 0x1c, 0x6c, 0x64, 0x85, 0x2c, 0xfb, 0xed, 0x1d, 0x43, 0x19, 0x8f,
	0x86, 0x07, 0x6c, 0x3a, 0x22, 0xfb, 0x0b, 0xd3, 0xcd, 0xd3, 0xba, 0xc7, 0x08, 0x9a, 0x90, 0x29,
	0xb6, 0xfe, 0x0e, 0x8a, 0xda, 0xb2, 0x58, 0xc9, 0x1f, 0x3a, 0xcb, 0xff, 0x5e, 0x67, 0xcf, 0xc1,
	0x9a, 0xaf, 0xc0, 0x1c, 0x9a, 0x0b, 0x95, 0x74, 0xf3, 0x30, 0x35, 0x8a, 0x8a, 0xd0, 0x9a, 0x12,
	0xe2, 0x72, 0x72, 0x3b, 0xda, 0x47, 0xb2, 0x18, 0xe7, 0x4b, 0x1e, 0xec, 0xd6, 0x39, 0xf3, 0x2e,
	0x9e, 0xc5, 0xfd, 0x1b, 0x3b, 0xb0, 0x53, 0x58, 0x1f, 0x72, 0xd6, 0x0f, 0x3c, 0xa4, 0xeb, 0x9a,
	0xb5, 0x51, 0x5a, 0x66, 0x6d, 0xac, 0x65, 0xb0, 0xd4, 0x60, 0xb5, 0x60, 0x55, 0xf2, 0x11, 0x6e,
	0xfa, 0x4c, 0xa5, 0xbf, 0x97, 0x50, 0xe9, 0x1f, 0x85, 0x99, 0xbe, 0x09, 0x3c, 0xce, 0xb5, 0x01,
	0x0d, 0xc5, 0x3c, 0x4b, 0x71, 0x09, 0x96, 0x55, 0x0d, 0xca, 0x68, 0x9c, 0x53, 0xd8, 0xb9, 0x46,
	0x29, 0xa3, 0xfb, 0x3e, 0x6c, 0xcc, 0x5a, 0x8e, 0xa8, 0xf4, 0xce, 0x59, 0x5f, 0x2b, 0x56, 0x21,
	0x33, 0x2d, 0xda, 0xa9, 0xdd, 0xf9, 0x9c, 0x83, 0x1d, 0xb5, 0x60, 0xda, 0x49, 0x3f, 0x18, 0x4c,
	0xde, 0xf2, 0xe0, 0x46, 0x54, 0x3f, 0x82, 0x22, 0x1f, 0x85, 0x6c, 0xfa, 0x89, 0xdc, 0x59, 0xb4,
	0xd6, 0xe6, 0xb3, 0x62, 0x3c, 0x49, 0x51, 0xcd, 0x47, 0xb0, 0x85, 0x4b, 0xe5, 0x1a, 0x50, 0x73,
	0xb7, 0x69, 0x5e, 0x3a, 0xe9, 0xbd, 0xdb, 0x36, 0xd7, 0xee, 0x99, 0xba, 0xdb, 0xce, 0x72, 0xbd,
	0x92, 0xbe, 0xe4, 0x1e, 0x7c, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x88, 0x02, 0xc9, 0xd3, 0x07,
	0x00, 0x00,
}
