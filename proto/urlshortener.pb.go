// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/urlshortener.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// request for short URL
type ShortenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OriginalURL   string                 `protobuf:"bytes,1,opt,name=originalURL,proto3" json:"originalURL,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortenRequest) Reset() {
	*x = ShortenRequest{}
	mi := &file_proto_urlshortener_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenRequest) ProtoMessage() {}

func (x *ShortenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urlshortener_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenRequest.ProtoReflect.Descriptor instead.
func (*ShortenRequest) Descriptor() ([]byte, []int) {
	return file_proto_urlshortener_proto_rawDescGZIP(), []int{0}
}

func (x *ShortenRequest) GetOriginalURL() string {
	if x != nil {
		return x.OriginalURL
	}
	return ""
}

// response URL
type ShortenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortURL      string                 `protobuf:"bytes,1,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortenResponse) Reset() {
	*x = ShortenResponse{}
	mi := &file_proto_urlshortener_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenResponse) ProtoMessage() {}

func (x *ShortenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urlshortener_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenResponse.ProtoReflect.Descriptor instead.
func (*ShortenResponse) Descriptor() ([]byte, []int) {
	return file_proto_urlshortener_proto_rawDescGZIP(), []int{1}
}

func (x *ShortenResponse) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

// request for original URL
type OriginalRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortURL      string                 `protobuf:"bytes,1,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OriginalRequest) Reset() {
	*x = OriginalRequest{}
	mi := &file_proto_urlshortener_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OriginalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginalRequest) ProtoMessage() {}

func (x *OriginalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urlshortener_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginalRequest.ProtoReflect.Descriptor instead.
func (*OriginalRequest) Descriptor() ([]byte, []int) {
	return file_proto_urlshortener_proto_rawDescGZIP(), []int{2}
}

func (x *OriginalRequest) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

// Response with original URL
type OriginalResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OriginalURL   string                 `protobuf:"bytes,1,opt,name=originalURL,proto3" json:"originalURL,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OriginalResponse) Reset() {
	*x = OriginalResponse{}
	mi := &file_proto_urlshortener_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OriginalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginalResponse) ProtoMessage() {}

func (x *OriginalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urlshortener_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginalResponse.ProtoReflect.Descriptor instead.
func (*OriginalResponse) Descriptor() ([]byte, []int) {
	return file_proto_urlshortener_proto_rawDescGZIP(), []int{3}
}

func (x *OriginalResponse) GetOriginalURL() string {
	if x != nil {
		return x.OriginalURL
	}
	return ""
}

// Req stat
type StatRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortURL      string                 `protobuf:"bytes,1,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatRequest) Reset() {
	*x = StatRequest{}
	mi := &file_proto_urlshortener_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatRequest) ProtoMessage() {}

func (x *StatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urlshortener_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatRequest.ProtoReflect.Descriptor instead.
func (*StatRequest) Descriptor() ([]byte, []int) {
	return file_proto_urlshortener_proto_rawDescGZIP(), []int{4}
}

func (x *StatRequest) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

// Resp stat
type StatResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClickCount    int64                  `protobuf:"varint,1,opt,name=clickCount,proto3" json:"clickCount,omitempty"`
	LastAccess    string                 `protobuf:"bytes,2,opt,name=lastAccess,proto3" json:"lastAccess,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatResponse) Reset() {
	*x = StatResponse{}
	mi := &file_proto_urlshortener_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatResponse) ProtoMessage() {}

func (x *StatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urlshortener_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatResponse.ProtoReflect.Descriptor instead.
func (*StatResponse) Descriptor() ([]byte, []int) {
	return file_proto_urlshortener_proto_rawDescGZIP(), []int{5}
}

func (x *StatResponse) GetClickCount() int64 {
	if x != nil {
		return x.ClickCount
	}
	return 0
}

func (x *StatResponse) GetLastAccess() string {
	if x != nil {
		return x.LastAccess
	}
	return ""
}

var File_proto_urlshortener_proto protoreflect.FileDescriptor

var file_proto_urlshortener_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x52, 0x50, 0x43,
	0x74, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x32, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x22, 0x2d, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x52, 0x4c, 0x22, 0x2d, 0x0a, 0x0f, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x52, 0x4c, 0x22, 0x34, 0x0a, 0x10, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x22, 0x29, 0x0a, 0x0b, 0x53, 0x74,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x52, 0x4c, 0x22, 0x4e, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x63, 0x6b,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x98, 0x02, 0x0a, 0x0c, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x08, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x12, 0x18, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67,
	0x52, 0x50, 0x43, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a,
	0x01, 0x2a, 0x22, 0x08, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x12, 0x5c, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x12, 0x19,
	0x2e, 0x67, 0x52, 0x50, 0x43, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x52, 0x50, 0x43,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f,
	0x7b, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x7d, 0x12, 0x54, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x15, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x52, 0x50, 0x43, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x7b, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x7d,
	0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_urlshortener_proto_rawDescOnce sync.Once
	file_proto_urlshortener_proto_rawDescData []byte
)

func file_proto_urlshortener_proto_rawDescGZIP() []byte {
	file_proto_urlshortener_proto_rawDescOnce.Do(func() {
		file_proto_urlshortener_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_urlshortener_proto_rawDesc), len(file_proto_urlshortener_proto_rawDesc)))
	})
	return file_proto_urlshortener_proto_rawDescData
}

var file_proto_urlshortener_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_urlshortener_proto_goTypes = []any{
	(*ShortenRequest)(nil),   // 0: gRPCtest.ShortenRequest
	(*ShortenResponse)(nil),  // 1: gRPCtest.ShortenResponse
	(*OriginalRequest)(nil),  // 2: gRPCtest.OriginalRequest
	(*OriginalResponse)(nil), // 3: gRPCtest.OriginalResponse
	(*StatRequest)(nil),      // 4: gRPCtest.StatRequest
	(*StatResponse)(nil),     // 5: gRPCtest.StatResponse
}
var file_proto_urlshortener_proto_depIdxs = []int32{
	0, // 0: gRPCtest.URLShortener.URLShort:input_type -> gRPCtest.ShortenRequest
	2, // 1: gRPCtest.URLShortener.GetOriginalURL:input_type -> gRPCtest.OriginalRequest
	4, // 2: gRPCtest.URLShortener.GetStats:input_type -> gRPCtest.StatRequest
	1, // 3: gRPCtest.URLShortener.URLShort:output_type -> gRPCtest.ShortenResponse
	3, // 4: gRPCtest.URLShortener.GetOriginalURL:output_type -> gRPCtest.OriginalResponse
	5, // 5: gRPCtest.URLShortener.GetStats:output_type -> gRPCtest.StatResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_urlshortener_proto_init() }
func file_proto_urlshortener_proto_init() {
	if File_proto_urlshortener_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_urlshortener_proto_rawDesc), len(file_proto_urlshortener_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_urlshortener_proto_goTypes,
		DependencyIndexes: file_proto_urlshortener_proto_depIdxs,
		MessageInfos:      file_proto_urlshortener_proto_msgTypes,
	}.Build()
	File_proto_urlshortener_proto = out.File
	file_proto_urlshortener_proto_goTypes = nil
	file_proto_urlshortener_proto_depIdxs = nil
}
