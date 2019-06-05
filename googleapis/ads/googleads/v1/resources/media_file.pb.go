// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/media_file.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// A media file.
type MediaFile struct {
	// The resource name of the media file.
	// Media file resource names have the form:
	//
	// `customers/{customer_id}/mediaFiles/{media_file_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the media file.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Type of the media file.
	Type enums.MediaTypeEnum_MediaType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.MediaTypeEnum_MediaType" json:"type,omitempty"`
	// The mime type of the media file.
	MimeType enums.MimeTypeEnum_MimeType `protobuf:"varint,6,opt,name=mime_type,json=mimeType,proto3,enum=google.ads.googleads.v1.enums.MimeTypeEnum_MimeType" json:"mime_type,omitempty"`
	// The URL of where the original media file was downloaded from (or a file
	// name).
	SourceUrl *wrappers.StringValue `protobuf:"bytes,7,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	// The name of the media file. The name can be used by clients to help
	// identify previously uploaded media.
	Name *wrappers.StringValue `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	// The size of the media file in bytes.
	FileSize *wrappers.Int64Value `protobuf:"bytes,9,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	// The specific type of the media file.
	//
	// Types that are valid to be assigned to Mediatype:
	//	*MediaFile_Image
	//	*MediaFile_MediaBundle
	//	*MediaFile_Audio
	//	*MediaFile_Video
	Mediatype            isMediaFile_Mediatype `protobuf_oneof:"mediatype"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MediaFile) Reset()         { *m = MediaFile{} }
func (m *MediaFile) String() string { return proto.CompactTextString(m) }
func (*MediaFile) ProtoMessage()    {}
func (*MediaFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ffee613242502eb, []int{0}
}

func (m *MediaFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaFile.Unmarshal(m, b)
}
func (m *MediaFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaFile.Marshal(b, m, deterministic)
}
func (m *MediaFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaFile.Merge(m, src)
}
func (m *MediaFile) XXX_Size() int {
	return xxx_messageInfo_MediaFile.Size(m)
}
func (m *MediaFile) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaFile.DiscardUnknown(m)
}

var xxx_messageInfo_MediaFile proto.InternalMessageInfo

func (m *MediaFile) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *MediaFile) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MediaFile) GetType() enums.MediaTypeEnum_MediaType {
	if m != nil {
		return m.Type
	}
	return enums.MediaTypeEnum_UNSPECIFIED
}

func (m *MediaFile) GetMimeType() enums.MimeTypeEnum_MimeType {
	if m != nil {
		return m.MimeType
	}
	return enums.MimeTypeEnum_UNSPECIFIED
}

func (m *MediaFile) GetSourceUrl() *wrappers.StringValue {
	if m != nil {
		return m.SourceUrl
	}
	return nil
}

func (m *MediaFile) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MediaFile) GetFileSize() *wrappers.Int64Value {
	if m != nil {
		return m.FileSize
	}
	return nil
}

type isMediaFile_Mediatype interface {
	isMediaFile_Mediatype()
}

type MediaFile_Image struct {
	Image *MediaImage `protobuf:"bytes,3,opt,name=image,proto3,oneof"`
}

type MediaFile_MediaBundle struct {
	MediaBundle *MediaBundle `protobuf:"bytes,4,opt,name=media_bundle,json=mediaBundle,proto3,oneof"`
}

type MediaFile_Audio struct {
	Audio *MediaAudio `protobuf:"bytes,10,opt,name=audio,proto3,oneof"`
}

type MediaFile_Video struct {
	Video *MediaVideo `protobuf:"bytes,11,opt,name=video,proto3,oneof"`
}

func (*MediaFile_Image) isMediaFile_Mediatype() {}

func (*MediaFile_MediaBundle) isMediaFile_Mediatype() {}

func (*MediaFile_Audio) isMediaFile_Mediatype() {}

func (*MediaFile_Video) isMediaFile_Mediatype() {}

func (m *MediaFile) GetMediatype() isMediaFile_Mediatype {
	if m != nil {
		return m.Mediatype
	}
	return nil
}

func (m *MediaFile) GetImage() *MediaImage {
	if x, ok := m.GetMediatype().(*MediaFile_Image); ok {
		return x.Image
	}
	return nil
}

func (m *MediaFile) GetMediaBundle() *MediaBundle {
	if x, ok := m.GetMediatype().(*MediaFile_MediaBundle); ok {
		return x.MediaBundle
	}
	return nil
}

func (m *MediaFile) GetAudio() *MediaAudio {
	if x, ok := m.GetMediatype().(*MediaFile_Audio); ok {
		return x.Audio
	}
	return nil
}

func (m *MediaFile) GetVideo() *MediaVideo {
	if x, ok := m.GetMediatype().(*MediaFile_Video); ok {
		return x.Video
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MediaFile) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MediaFile_Image)(nil),
		(*MediaFile_MediaBundle)(nil),
		(*MediaFile_Audio)(nil),
		(*MediaFile_Video)(nil),
	}
}

// Encapsulates an Image.
type MediaImage struct {
	// Raw image data.
	Data                 *wrappers.BytesValue `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MediaImage) Reset()         { *m = MediaImage{} }
func (m *MediaImage) String() string { return proto.CompactTextString(m) }
func (*MediaImage) ProtoMessage()    {}
func (*MediaImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ffee613242502eb, []int{1}
}

func (m *MediaImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaImage.Unmarshal(m, b)
}
func (m *MediaImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaImage.Marshal(b, m, deterministic)
}
func (m *MediaImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaImage.Merge(m, src)
}
func (m *MediaImage) XXX_Size() int {
	return xxx_messageInfo_MediaImage.Size(m)
}
func (m *MediaImage) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaImage.DiscardUnknown(m)
}

var xxx_messageInfo_MediaImage proto.InternalMessageInfo

func (m *MediaImage) GetData() *wrappers.BytesValue {
	if m != nil {
		return m.Data
	}
	return nil
}

// Represents a ZIP archive media the content of which contains HTML5 assets.
type MediaBundle struct {
	// Raw zipped data.
	Data                 *wrappers.BytesValue `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MediaBundle) Reset()         { *m = MediaBundle{} }
func (m *MediaBundle) String() string { return proto.CompactTextString(m) }
func (*MediaBundle) ProtoMessage()    {}
func (*MediaBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ffee613242502eb, []int{2}
}

func (m *MediaBundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaBundle.Unmarshal(m, b)
}
func (m *MediaBundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaBundle.Marshal(b, m, deterministic)
}
func (m *MediaBundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaBundle.Merge(m, src)
}
func (m *MediaBundle) XXX_Size() int {
	return xxx_messageInfo_MediaBundle.Size(m)
}
func (m *MediaBundle) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaBundle.DiscardUnknown(m)
}

var xxx_messageInfo_MediaBundle proto.InternalMessageInfo

func (m *MediaBundle) GetData() *wrappers.BytesValue {
	if m != nil {
		return m.Data
	}
	return nil
}

// Encapsulates an Audio.
type MediaAudio struct {
	// The duration of the Audio in milliseconds.
	AdDurationMillis     *wrappers.Int64Value `protobuf:"bytes,1,opt,name=ad_duration_millis,json=adDurationMillis,proto3" json:"ad_duration_millis,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MediaAudio) Reset()         { *m = MediaAudio{} }
func (m *MediaAudio) String() string { return proto.CompactTextString(m) }
func (*MediaAudio) ProtoMessage()    {}
func (*MediaAudio) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ffee613242502eb, []int{3}
}

func (m *MediaAudio) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaAudio.Unmarshal(m, b)
}
func (m *MediaAudio) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaAudio.Marshal(b, m, deterministic)
}
func (m *MediaAudio) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaAudio.Merge(m, src)
}
func (m *MediaAudio) XXX_Size() int {
	return xxx_messageInfo_MediaAudio.Size(m)
}
func (m *MediaAudio) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaAudio.DiscardUnknown(m)
}

var xxx_messageInfo_MediaAudio proto.InternalMessageInfo

func (m *MediaAudio) GetAdDurationMillis() *wrappers.Int64Value {
	if m != nil {
		return m.AdDurationMillis
	}
	return nil
}

// Encapsulates a Video.
type MediaVideo struct {
	// The duration of the Video in milliseconds.
	AdDurationMillis *wrappers.Int64Value `protobuf:"bytes,1,opt,name=ad_duration_millis,json=adDurationMillis,proto3" json:"ad_duration_millis,omitempty"`
	// The YouTube video ID (as seen in YouTube URLs).
	YoutubeVideoId *wrappers.StringValue `protobuf:"bytes,2,opt,name=youtube_video_id,json=youtubeVideoId,proto3" json:"youtube_video_id,omitempty"`
	// The Advertising Digital Identification code for this video, as defined by
	// the American Association of Advertising Agencies, used mainly for
	// television commercials.
	AdvertisingIdCode *wrappers.StringValue `protobuf:"bytes,3,opt,name=advertising_id_code,json=advertisingIdCode,proto3" json:"advertising_id_code,omitempty"`
	// The Industry Standard Commercial Identifier code for this video, used
	// mainly for television commercials.
	IsciCode             *wrappers.StringValue `protobuf:"bytes,4,opt,name=isci_code,json=isciCode,proto3" json:"isci_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MediaVideo) Reset()         { *m = MediaVideo{} }
func (m *MediaVideo) String() string { return proto.CompactTextString(m) }
func (*MediaVideo) ProtoMessage()    {}
func (*MediaVideo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ffee613242502eb, []int{4}
}

func (m *MediaVideo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaVideo.Unmarshal(m, b)
}
func (m *MediaVideo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaVideo.Marshal(b, m, deterministic)
}
func (m *MediaVideo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaVideo.Merge(m, src)
}
func (m *MediaVideo) XXX_Size() int {
	return xxx_messageInfo_MediaVideo.Size(m)
}
func (m *MediaVideo) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaVideo.DiscardUnknown(m)
}

var xxx_messageInfo_MediaVideo proto.InternalMessageInfo

func (m *MediaVideo) GetAdDurationMillis() *wrappers.Int64Value {
	if m != nil {
		return m.AdDurationMillis
	}
	return nil
}

func (m *MediaVideo) GetYoutubeVideoId() *wrappers.StringValue {
	if m != nil {
		return m.YoutubeVideoId
	}
	return nil
}

func (m *MediaVideo) GetAdvertisingIdCode() *wrappers.StringValue {
	if m != nil {
		return m.AdvertisingIdCode
	}
	return nil
}

func (m *MediaVideo) GetIsciCode() *wrappers.StringValue {
	if m != nil {
		return m.IsciCode
	}
	return nil
}

func init() {
	proto.RegisterType((*MediaFile)(nil), "google.ads.googleads.v1.resources.MediaFile")
	proto.RegisterType((*MediaImage)(nil), "google.ads.googleads.v1.resources.MediaImage")
	proto.RegisterType((*MediaBundle)(nil), "google.ads.googleads.v1.resources.MediaBundle")
	proto.RegisterType((*MediaAudio)(nil), "google.ads.googleads.v1.resources.MediaAudio")
	proto.RegisterType((*MediaVideo)(nil), "google.ads.googleads.v1.resources.MediaVideo")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/media_file.proto", fileDescriptor_7ffee613242502eb)
}

var fileDescriptor_7ffee613242502eb = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xdf, 0x6a, 0xd4, 0x4c,
	0x18, 0xc6, 0xbb, 0xe9, 0xb6, 0xdf, 0x66, 0xb6, 0x5f, 0xa9, 0xe3, 0x49, 0xa8, 0x45, 0xda, 0x95,
	0x42, 0x41, 0x3a, 0x71, 0x6b, 0x29, 0x9a, 0xa2, 0x90, 0xd5, 0xfe, 0xd9, 0x62, 0xa5, 0xa6, 0xba,
	0x82, 0x2c, 0x84, 0xd9, 0x9d, 0x69, 0x18, 0x48, 0x32, 0x21, 0x93, 0xac, 0x6c, 0x2f, 0xc7, 0x23,
	0xf1, 0x52, 0xbc, 0x06, 0xaf, 0xc0, 0x5b, 0xf0, 0x44, 0x66, 0x32, 0xc9, 0x2e, 0x68, 0xbb, 0x5d,
	0xf0, 0xec, 0x9d, 0xce, 0xf3, 0xfc, 0xf2, 0x3e, 0x6f, 0xdf, 0x64, 0xc1, 0x5e, 0xc0, 0x79, 0x10,
	0x52, 0x1b, 0x13, 0x61, 0x17, 0xa5, 0xac, 0x46, 0x6d, 0x3b, 0xa5, 0x82, 0xe7, 0xe9, 0x90, 0x0a,
	0x3b, 0xa2, 0x84, 0x61, 0xff, 0x8a, 0x85, 0x14, 0x25, 0x29, 0xcf, 0x38, 0xdc, 0x2a, 0x84, 0x08,
	0x13, 0x81, 0x2a, 0x0f, 0x1a, 0xb5, 0x51, 0xe5, 0x59, 0x47, 0x37, 0x61, 0x69, 0x9c, 0x47, 0x25,
	0x32, 0x1b, 0x27, 0x1a, 0xb9, 0xbe, 0x3b, 0x43, 0xcf, 0x22, 0x3a, 0x2d, 0xdf, 0x28, 0xe5, 0x09,
	0xb3, 0x71, 0x1c, 0xf3, 0x0c, 0x67, 0x8c, 0xc7, 0x42, 0xdf, 0x3e, 0xd4, 0xb7, 0xea, 0x34, 0xc8,
	0xaf, 0xec, 0xcf, 0x29, 0x4e, 0x12, 0x9a, 0xea, 0xfb, 0xd6, 0x8f, 0x25, 0x60, 0x9e, 0xcb, 0x0e,
	0x8e, 0x59, 0x48, 0xe1, 0x23, 0xf0, 0x7f, 0xd9, 0xb7, 0x1f, 0xe3, 0x88, 0x5a, 0xb5, 0xcd, 0xda,
	0x8e, 0xe9, 0xad, 0x94, 0x7f, 0x7c, 0x8b, 0x23, 0x0a, 0x1f, 0x03, 0x83, 0x11, 0xcb, 0xd8, 0xac,
	0xed, 0x34, 0xf7, 0x1e, 0xe8, 0x70, 0xa8, 0xe4, 0xa3, 0x6e, 0x9c, 0x1d, 0xec, 0xf7, 0x70, 0x98,
	0x53, 0xcf, 0x60, 0x04, 0x9e, 0x81, 0xba, 0xec, 0xd5, 0x5a, 0xda, 0xac, 0xed, 0xac, 0xee, 0x1d,
	0xa0, 0x9b, 0xc6, 0xa5, 0xb2, 0x21, 0xd5, 0xc9, 0xfb, 0x71, 0x42, 0x8f, 0xe2, 0x3c, 0x9a, 0x9c,
	0x3c, 0xc5, 0x80, 0xef, 0x80, 0x59, 0x85, 0xb7, 0x96, 0x15, 0x70, 0x7f, 0x16, 0x90, 0x45, 0x74,
	0xc2, 0xd3, 0x07, 0xaf, 0x11, 0xe9, 0x0a, 0x1e, 0x02, 0xa0, 0xe3, 0xe6, 0x69, 0x68, 0xfd, 0xa7,
	0x32, 0x6d, 0xfc, 0x91, 0xe9, 0x32, 0x4b, 0x59, 0x1c, 0x14, 0xa1, 0xcc, 0x42, 0xff, 0x21, 0x0d,
	0xe1, 0x13, 0x50, 0x57, 0x43, 0x6a, 0xdc, 0xc1, 0xa6, 0x94, 0xf0, 0x19, 0x30, 0xe5, 0xee, 0xf8,
	0x82, 0x5d, 0x53, 0xcb, 0x9c, 0x3d, 0xc1, 0x86, 0x54, 0x5f, 0xb2, 0x6b, 0x0a, 0x8f, 0xc0, 0x12,
	0x8b, 0x70, 0x40, 0xad, 0x45, 0xe5, 0xda, 0x45, 0x33, 0xf7, 0xae, 0x18, 0x5f, 0x57, 0x9a, 0x4e,
	0x17, 0xbc, 0xc2, 0x0d, 0x2f, 0xc1, 0x4a, 0xb1, 0x6f, 0x83, 0x3c, 0x26, 0x21, 0xb5, 0xea, 0x8a,
	0x86, 0xee, 0x4a, 0xeb, 0x28, 0xd7, 0xe9, 0x82, 0xd7, 0x8c, 0x26, 0x47, 0xd9, 0x1b, 0xce, 0x09,
	0xe3, 0x16, 0x98, 0xaf, 0x37, 0x57, 0x9a, 0x64, 0x6f, 0xca, 0x2d, 0x31, 0x23, 0x46, 0x28, 0xb7,
	0x9a, 0xf3, 0x61, 0x7a, 0xd2, 0x24, 0x31, 0xca, 0xdd, 0x69, 0x02, 0x53, 0x35, 0x27, 0xb7, 0xa4,
	0xf5, 0x02, 0x80, 0xc9, 0x18, 0xa0, 0x0d, 0xea, 0x04, 0x67, 0x58, 0x6d, 0xf5, 0xdf, 0x26, 0xdf,
	0x19, 0x67, 0x54, 0xe8, 0xff, 0x97, 0x14, 0xb6, 0x5e, 0x82, 0xe6, 0x54, 0xee, 0xf9, 0xfd, 0x1f,
	0xf5, 0xe3, 0x55, 0x52, 0xd8, 0x05, 0x10, 0x13, 0x9f, 0xe4, 0xa9, 0x7a, 0x43, 0xfd, 0x88, 0x85,
	0x21, 0x13, 0x37, 0xc2, 0xa6, 0xd6, 0x60, 0x0d, 0x93, 0xd7, 0xda, 0x75, 0xae, 0x4c, 0xad, 0xaf,
	0x86, 0x26, 0xab, 0xf0, 0xff, 0x90, 0x0c, 0x8f, 0xc1, 0xda, 0x98, 0xe7, 0x59, 0x3e, 0xa0, 0xbe,
	0x9a, 0xa7, 0x5f, 0xbd, 0xeb, 0xb7, 0x2f, 0xf8, 0xaa, 0x76, 0xa9, 0x86, 0xba, 0x04, 0xbe, 0x01,
	0xf7, 0x31, 0x19, 0xd1, 0x34, 0x63, 0x82, 0xc5, 0x81, 0xcf, 0x88, 0x3f, 0xe4, 0xa4, 0x5c, 0xdf,
	0xdb, 0x51, 0xf7, 0xa6, 0x8c, 0x5d, 0xf2, 0x8a, 0x13, 0x0a, 0x9f, 0x03, 0x93, 0x89, 0x21, 0x2b,
	0x18, 0xf5, 0x3b, 0x30, 0x1a, 0x52, 0x2e, 0xad, 0x9d, 0x5f, 0x35, 0xb0, 0x3d, 0xe4, 0xd1, 0xec,
	0x6d, 0xea, 0xac, 0x56, 0x1f, 0xc2, 0x0b, 0x89, 0xbc, 0xa8, 0x7d, 0x3a, 0xd3, 0xa6, 0x80, 0x87,
	0x38, 0x0e, 0x10, 0x4f, 0x03, 0x3b, 0xa0, 0xb1, 0x7a, 0x60, 0xf9, 0x69, 0x4e, 0x98, 0xb8, 0xe5,
	0x07, 0xe3, 0xb0, 0xaa, 0xbe, 0x18, 0x8b, 0x27, 0xae, 0xfb, 0xcd, 0xd8, 0x3a, 0x29, 0x90, 0x2e,
	0x11, 0xa8, 0x28, 0x65, 0xd5, 0x6b, 0x23, 0xaf, 0x54, 0x7e, 0x2f, 0x35, 0x7d, 0x97, 0x88, 0x7e,
	0xa5, 0xe9, 0xf7, 0xda, 0xfd, 0x4a, 0xf3, 0xd3, 0xd8, 0x2e, 0x2e, 0x1c, 0xc7, 0x25, 0xc2, 0x71,
	0x2a, 0x95, 0xe3, 0xf4, 0xda, 0x8e, 0x53, 0xe9, 0x06, 0xcb, 0xaa, 0xd9, 0xa7, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xa4, 0x7e, 0xab, 0xbf, 0xdc, 0x06, 0x00, 0x00,
}
