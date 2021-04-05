// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: src/simple/simple.proto

package _12___Protobuf_with_go

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

type SimpleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsSimple   bool    `protobuf:"varint,2,opt,name=is_simple,json=isSimple,proto3" json:"is_simple,omitempty"`
	Name       string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	SampleList []int32 `protobuf:"varint,4,rep,packed,name=sample_list,json=sampleList,proto3" json:"sample_list,omitempty"`
}

func (x *SimpleMessage) Reset() {
	*x = SimpleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_simple_simple_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleMessage) ProtoMessage() {}

func (x *SimpleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_src_simple_simple_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleMessage.ProtoReflect.Descriptor instead.
func (*SimpleMessage) Descriptor() ([]byte, []int) {
	return file_src_simple_simple_proto_rawDescGZIP(), []int{0}
}

func (x *SimpleMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SimpleMessage) GetIsSimple() bool {
	if x != nil {
		return x.IsSimple
	}
	return false
}

func (x *SimpleMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SimpleMessage) GetSampleList() []int32 {
	if x != nil {
		return x.SampleList
	}
	return nil
}

var File_src_simple_simple_proto protoreflect.FileDescriptor

var file_src_simple_simple_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x72, 0x63, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x71, 0x0a, 0x0d, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x48, 0x5a, 0x46,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x75, 0x73, 0x74, 0x61,
	0x76, 0x6f, 0x68, 0x6d, 0x73, 0x69, 0x6c, 0x76, 0x61, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2d, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2d, 0x73, 0x74, 0x75, 0x64, 0x79,
	0x2f, 0x31, 0x32, 0x5f, 0x2d, 0x5f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x77,
	0x69, 0x74, 0x68, 0x5f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_simple_simple_proto_rawDescOnce sync.Once
	file_src_simple_simple_proto_rawDescData = file_src_simple_simple_proto_rawDesc
)

func file_src_simple_simple_proto_rawDescGZIP() []byte {
	file_src_simple_simple_proto_rawDescOnce.Do(func() {
		file_src_simple_simple_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_simple_simple_proto_rawDescData)
	})
	return file_src_simple_simple_proto_rawDescData
}

var file_src_simple_simple_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_src_simple_simple_proto_goTypes = []interface{}{
	(*SimpleMessage)(nil), // 0: example.simple.SimpleMessage
}
var file_src_simple_simple_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_simple_simple_proto_init() }
func file_src_simple_simple_proto_init() {
	if File_src_simple_simple_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_simple_simple_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleMessage); i {
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
			RawDescriptor: file_src_simple_simple_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_simple_simple_proto_goTypes,
		DependencyIndexes: file_src_simple_simple_proto_depIdxs,
		MessageInfos:      file_src_simple_simple_proto_msgTypes,
	}.Build()
	File_src_simple_simple_proto = out.File
	file_src_simple_simple_proto_rawDesc = nil
	file_src_simple_simple_proto_goTypes = nil
	file_src_simple_simple_proto_depIdxs = nil
}
