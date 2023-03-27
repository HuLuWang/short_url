// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: transform.proto

package transform

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExpandReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shorten string `protobuf:"bytes,1,opt,name=shorten,proto3" json:"shorten,omitempty"`
}

func (x *ExpandReq) Reset() {
	*x = ExpandReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandReq) ProtoMessage() {}

func (x *ExpandReq) ProtoReflect() protoreflect.Message {
	mi := &file_transform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandReq.ProtoReflect.Descriptor instead.
func (*ExpandReq) Descriptor() ([]byte, []int) {
	return file_transform_proto_rawDescGZIP(), []int{0}
}

func (x *ExpandReq) GetShorten() string {
	if x != nil {
		return x.Shorten
	}
	return ""
}

type ExpandResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ExpandResp) Reset() {
	*x = ExpandResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandResp) ProtoMessage() {}

func (x *ExpandResp) ProtoReflect() protoreflect.Message {
	mi := &file_transform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandResp.ProtoReflect.Descriptor instead.
func (*ExpandResp) Descriptor() ([]byte, []int) {
	return file_transform_proto_rawDescGZIP(), []int{1}
}

func (x *ExpandResp) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ShortenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ShortenReq) Reset() {
	*x = ShortenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transform_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenReq) ProtoMessage() {}

func (x *ShortenReq) ProtoReflect() protoreflect.Message {
	mi := &file_transform_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenReq.ProtoReflect.Descriptor instead.
func (*ShortenReq) Descriptor() ([]byte, []int) {
	return file_transform_proto_rawDescGZIP(), []int{2}
}

func (x *ShortenReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ShortenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shorten string `protobuf:"bytes,1,opt,name=shorten,proto3" json:"shorten,omitempty"`
}

func (x *ShortenResp) Reset() {
	*x = ShortenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transform_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenResp) ProtoMessage() {}

func (x *ShortenResp) ProtoReflect() protoreflect.Message {
	mi := &file_transform_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenResp.ProtoReflect.Descriptor instead.
func (*ShortenResp) Descriptor() ([]byte, []int) {
	return file_transform_proto_rawDescGZIP(), []int{3}
}

func (x *ShortenResp) GetShorten() string {
	if x != nil {
		return x.Shorten
	}
	return ""
}

var File_transform_proto protoreflect.FileDescriptor

var file_transform_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x25, 0x0a, 0x09,
	0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x22, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x22, 0x1e, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x22, 0x27, 0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x32, 0x7e, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x06, 0x65,
	0x78, 0x70, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x38, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x12, 0x15, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0d, 0x5a, 0x0b,
	0x2e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_transform_proto_rawDescOnce sync.Once
	file_transform_proto_rawDescData = file_transform_proto_rawDesc
)

func file_transform_proto_rawDescGZIP() []byte {
	file_transform_proto_rawDescOnce.Do(func() {
		file_transform_proto_rawDescData = protoimpl.X.CompressGZIP(file_transform_proto_rawDescData)
	})
	return file_transform_proto_rawDescData
}

var file_transform_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_transform_proto_goTypes = []interface{}{
	(*ExpandReq)(nil),   // 0: transform.expandReq
	(*ExpandResp)(nil),  // 1: transform.expandResp
	(*ShortenReq)(nil),  // 2: transform.shortenReq
	(*ShortenResp)(nil), // 3: transform.shortenResp
}
var file_transform_proto_depIdxs = []int32{
	0, // 0: transform.transformer.expand:input_type -> transform.expandReq
	2, // 1: transform.transformer.shorten:input_type -> transform.shortenReq
	1, // 2: transform.transformer.expand:output_type -> transform.expandResp
	3, // 3: transform.transformer.shorten:output_type -> transform.shortenResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transform_proto_init() }
func file_transform_proto_init() {
	if File_transform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandReq); i {
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
		file_transform_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandResp); i {
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
		file_transform_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenReq); i {
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
		file_transform_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenResp); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transform_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transform_proto_goTypes,
		DependencyIndexes: file_transform_proto_depIdxs,
		MessageInfos:      file_transform_proto_msgTypes,
	}.Build()
	File_transform_proto = out.File
	file_transform_proto_rawDesc = nil
	file_transform_proto_goTypes = nil
	file_transform_proto_depIdxs = nil
}
