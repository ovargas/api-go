// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/commons/v1/message.proto

package commons

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

type SortType int32

const (
	SortType_SORT_TYPE_UNSPECIFIED SortType = 0
	SortType_ASC                   SortType = 1
	SortType_DESC                  SortType = 2
)

// Enum value maps for SortType.
var (
	SortType_name = map[int32]string{
		0: "SORT_TYPE_UNSPECIFIED",
		1: "ASC",
		2: "DESC",
	}
	SortType_value = map[string]int32{
		"SORT_TYPE_UNSPECIFIED": 0,
		"ASC":                   1,
		"DESC":                  2,
	}
)

func (x SortType) Enum() *SortType {
	p := new(SortType)
	*p = x
	return p
}

func (x SortType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_v1_message_proto_enumTypes[0].Descriptor()
}

func (SortType) Type() protoreflect.EnumType {
	return &file_api_commons_v1_message_proto_enumTypes[0]
}

func (x SortType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortType.Descriptor instead.
func (SortType) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_v1_message_proto_rawDescGZIP(), []int{0}
}

type Sort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Sort  SortType `protobuf:"varint,2,opt,name=sort,proto3,enum=api.commons.v1.SortType" json:"sort,omitempty"`
}

func (x *Sort) Reset() {
	*x = Sort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sort) ProtoMessage() {}

func (x *Sort) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sort.ProtoReflect.Descriptor instead.
func (*Sort) Descriptor() ([]byte, []int) {
	return file_api_commons_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *Sort) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Sort) GetSort() SortType {
	if x != nil {
		return x.Sort
	}
	return SortType_SORT_TYPE_UNSPECIFIED
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Size   int32   `protobuf:"zigzag32,2,opt,name=size,proto3" json:"size,omitempty"`
	Sorts  []*Sort `protobuf:"bytes,3,rep,name=sorts,proto3" json:"sorts,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_api_commons_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *Page) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Page) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Page) GetSorts() []*Sort {
	if x != nil {
		return x.Sorts
	}
	return nil
}

var File_api_commons_v1_message_proto protoreflect.FileDescriptor

var file_api_commons_v1_message_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x4a,
	0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2c, 0x0a, 0x04,
	0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x5e, 0x0a, 0x04, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x2a,
	0x0a, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x6f, 0x72, 0x74, 0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x2a, 0x38, 0x0a, 0x08, 0x53, 0x6f,
	0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45,
	0x53, 0x43, 0x10, 0x02, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6f, 0x76, 0x61, 0x72, 0x67, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_v1_message_proto_rawDescOnce sync.Once
	file_api_commons_v1_message_proto_rawDescData = file_api_commons_v1_message_proto_rawDesc
)

func file_api_commons_v1_message_proto_rawDescGZIP() []byte {
	file_api_commons_v1_message_proto_rawDescOnce.Do(func() {
		file_api_commons_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_v1_message_proto_rawDescData)
	})
	return file_api_commons_v1_message_proto_rawDescData
}

var file_api_commons_v1_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_commons_v1_message_proto_goTypes = []interface{}{
	(SortType)(0), // 0: api.commons.v1.SortType
	(*Sort)(nil),  // 1: api.commons.v1.Sort
	(*Page)(nil),  // 2: api.commons.v1.Page
}
var file_api_commons_v1_message_proto_depIdxs = []int32{
	0, // 0: api.commons.v1.Sort.sort:type_name -> api.commons.v1.SortType
	1, // 1: api.commons.v1.Page.sorts:type_name -> api.commons.v1.Sort
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_commons_v1_message_proto_init() }
func file_api_commons_v1_message_proto_init() {
	if File_api_commons_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_commons_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sort); i {
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
		file_api_commons_v1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
			RawDescriptor: file_api_commons_v1_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_v1_message_proto_goTypes,
		DependencyIndexes: file_api_commons_v1_message_proto_depIdxs,
		EnumInfos:         file_api_commons_v1_message_proto_enumTypes,
		MessageInfos:      file_api_commons_v1_message_proto_msgTypes,
	}.Build()
	File_api_commons_v1_message_proto = out.File
	file_api_commons_v1_message_proto_rawDesc = nil
	file_api_commons_v1_message_proto_goTypes = nil
	file_api_commons_v1_message_proto_depIdxs = nil
}
