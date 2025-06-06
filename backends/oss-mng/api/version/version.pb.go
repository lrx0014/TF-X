// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: version/version.proto

package version

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

type VersionReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VersionReq) Reset() {
	*x = VersionReq{}
	mi := &file_version_version_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VersionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionReq) ProtoMessage() {}

func (x *VersionReq) ProtoReflect() protoreflect.Message {
	mi := &file_version_version_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionReq.ProtoReflect.Descriptor instead.
func (*VersionReq) Descriptor() ([]byte, []int) {
	return file_version_version_proto_rawDescGZIP(), []int{0}
}

type VersionReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version       string                 `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VersionReply) Reset() {
	*x = VersionReply{}
	mi := &file_version_version_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VersionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionReply) ProtoMessage() {}

func (x *VersionReply) ProtoReflect() protoreflect.Message {
	mi := &file_version_version_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionReply.ProtoReflect.Descriptor instead.
func (*VersionReply) Descriptor() ([]byte, []int) {
	return file_version_version_proto_rawDescGZIP(), []int{1}
}

func (x *VersionReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VersionReply) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_version_version_proto protoreflect.FileDescriptor

const file_version_version_proto_rawDesc = "" +
	"\n" +
	"\x15version/version.proto\x12\aversion\x1a\x1cgoogle/api/annotations.proto\"\f\n" +
	"\n" +
	"VersionReq\"<\n" +
	"\fVersionReply\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x18\n" +
	"\aversion\x18\x02 \x01(\tR\aversion2Z\n" +
	"\aVersion\x12O\n" +
	"\aversion\x12\x13.version.VersionReq\x1a\x15.version.VersionReply\"\x18\x82\xd3\xe4\x93\x02\x12\x12\x10/oss-mng/versionBE\n" +
	"\x16dev.kratos.api.versionB\fVersionProtoP\x01Z\x1boss-mng/api/version;versionb\x06proto3"

var (
	file_version_version_proto_rawDescOnce sync.Once
	file_version_version_proto_rawDescData []byte
)

func file_version_version_proto_rawDescGZIP() []byte {
	file_version_version_proto_rawDescOnce.Do(func() {
		file_version_version_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_version_version_proto_rawDesc), len(file_version_version_proto_rawDesc)))
	})
	return file_version_version_proto_rawDescData
}

var file_version_version_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_version_version_proto_goTypes = []any{
	(*VersionReq)(nil),   // 0: version.VersionReq
	(*VersionReply)(nil), // 1: version.VersionReply
}
var file_version_version_proto_depIdxs = []int32{
	0, // 0: version.Version.version:input_type -> version.VersionReq
	1, // 1: version.Version.version:output_type -> version.VersionReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_version_version_proto_init() }
func file_version_version_proto_init() {
	if File_version_version_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_version_version_proto_rawDesc), len(file_version_version_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_version_version_proto_goTypes,
		DependencyIndexes: file_version_version_proto_depIdxs,
		MessageInfos:      file_version_version_proto_msgTypes,
	}.Build()
	File_version_version_proto = out.File
	file_version_version_proto_goTypes = nil
	file_version_version_proto_depIdxs = nil
}
