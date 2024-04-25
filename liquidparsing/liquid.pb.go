package liquidparsingpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HtmlPart  string            `protobuf:"bytes,1,opt,name=html_part,json=htmlPart,proto3" json:"html_part,omitempty"`
	Variables map[string]string `protobuf:"bytes,2,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Template) Reset() {
	*x = Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liquid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_liquid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template.ProtoReflect.Descriptor instead.
func (*Template) Descriptor() ([]byte, []int) {
	return file_liquid_proto_rawDescGZIP(), []int{0}
}

func (x *Template) GetHtmlPart() string {
	if x != nil {
		return x.HtmlPart
	}
	return ""
}

func (x *Template) GetVariables() map[string]string {
	if x != nil {
		return x.Variables
	}
	return nil
}

type ParsedResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ParsedResult) Reset() {
	*x = ParsedResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liquid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParsedResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParsedResult) ProtoMessage() {}

func (x *ParsedResult) ProtoReflect() protoreflect.Message {
	mi := &file_liquid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParsedResult.ProtoReflect.Descriptor instead.
func (*ParsedResult) Descriptor() ([]byte, []int) {
	return file_liquid_proto_rawDescGZIP(), []int{1}
}

func (x *ParsedResult) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_liquid_proto protoreflect.FileDescriptor

var file_liquid_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x22, 0xa4, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x74, 0x6d, 0x6c, 0x50, 0x61, 0x72, 0x74,
	0x12, 0x3d, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x2e, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x1a,
	0x3c, 0x0a, 0x0e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x26, 0x0a,
	0x0c, 0x50, 0x61, 0x72, 0x73, 0x65, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x51, 0x0a, 0x0d, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x50,
	0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x40, 0x0a, 0x14, 0x50, 0x61, 0x72, 0x73, 0x65, 0x41,
	0x6e, 0x64, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x10,
	0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x1a, 0x14, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x6c, 0x69, 0x71, 0x75,
	0x69, 0x64, 0x70, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x70, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_liquid_proto_rawDescOnce sync.Once
	file_liquid_proto_rawDescData = file_liquid_proto_rawDesc
)

func file_liquid_proto_rawDescGZIP() []byte {
	file_liquid_proto_rawDescOnce.Do(func() {
		file_liquid_proto_rawDescData = protoimpl.X.CompressGZIP(file_liquid_proto_rawDescData)
	})
	return file_liquid_proto_rawDescData
}

var file_liquid_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_liquid_proto_goTypes = []interface{}{
	(*Template)(nil),     // 0: liquid.Template
	(*ParsedResult)(nil), // 1: liquid.ParsedResult
	nil,                  // 2: liquid.Template.VariablesEntry
}
var file_liquid_proto_depIdxs = []int32{
	2, // 0: liquid.Template.variables:type_name -> liquid.Template.VariablesEntry
	0, // 1: liquid.LiquidParsing.ParseAndRenderString:input_type -> liquid.Template
	1, // 2: liquid.LiquidParsing.ParseAndRenderString:output_type -> liquid.ParsedResult
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_liquid_proto_init() }
func file_liquid_proto_init() {
	if File_liquid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_liquid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template); i {
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
		file_liquid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParsedResult); i {
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
			RawDescriptor: file_liquid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_liquid_proto_goTypes,
		DependencyIndexes: file_liquid_proto_depIdxs,
		MessageInfos:      file_liquid_proto_msgTypes,
	}.Build()
	File_liquid_proto = out.File
	file_liquid_proto_rawDesc = nil
	file_liquid_proto_goTypes = nil
	file_liquid_proto_depIdxs = nil
}
