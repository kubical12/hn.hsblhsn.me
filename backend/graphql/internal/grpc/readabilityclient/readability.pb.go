// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: readability.proto

package readabilityclient

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

type GetReadableDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Html       string `protobuf:"bytes,2,opt,name=html,proto3" json:"html,omitempty"`
}

func (x *GetReadableDocumentRequest) Reset() {
	*x = GetReadableDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_readability_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReadableDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReadableDocumentRequest) ProtoMessage() {}

func (x *GetReadableDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_readability_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReadableDocumentRequest.ProtoReflect.Descriptor instead.
func (*GetReadableDocumentRequest) Descriptor() ([]byte, []int) {
	return file_readability_proto_rawDescGZIP(), []int{0}
}

func (x *GetReadableDocumentRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *GetReadableDocumentRequest) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

type GetReadableDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body  string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *GetReadableDocumentResponse) Reset() {
	*x = GetReadableDocumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_readability_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReadableDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReadableDocumentResponse) ProtoMessage() {}

func (x *GetReadableDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_readability_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReadableDocumentResponse.ProtoReflect.Descriptor instead.
func (*GetReadableDocumentResponse) Descriptor() ([]byte, []int) {
	return file_readability_proto_rawDescGZIP(), []int{1}
}

func (x *GetReadableDocumentResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetReadableDocumentResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type GetReadinessInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *GetReadinessInfoRequest) Reset() {
	*x = GetReadinessInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_readability_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReadinessInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReadinessInfoRequest) ProtoMessage() {}

func (x *GetReadinessInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_readability_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReadinessInfoRequest.ProtoReflect.Descriptor instead.
func (*GetReadinessInfoRequest) Descriptor() ([]byte, []int) {
	return file_readability_proto_rawDescGZIP(), []int{2}
}

func (x *GetReadinessInfoRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

type GetReadinessInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ready bool `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
}

func (x *GetReadinessInfoResponse) Reset() {
	*x = GetReadinessInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_readability_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReadinessInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReadinessInfoResponse) ProtoMessage() {}

func (x *GetReadinessInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_readability_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReadinessInfoResponse.ProtoReflect.Descriptor instead.
func (*GetReadinessInfoResponse) Descriptor() ([]byte, []int) {
	return file_readability_proto_rawDescGZIP(), []int{3}
}

func (x *GetReadinessInfoResponse) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

var File_readability_proto protoreflect.FileDescriptor

var file_readability_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x61, 0x64, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x72, 0x65, 0x61, 0x64, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x70, 0x62, 0x22, 0x50, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c,
	0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x74, 0x6d, 0x6c, 0x22, 0x47, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x61,
	0x62, 0x6c, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x39, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x30, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x32, 0xe4, 0x01, 0x0a, 0x0b, 0x52,
	0x65, 0x61, 0x64, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x6e, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x29, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x72,
	0x65, 0x61, 0x64, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26,
	0x2e, 0x72, 0x65, 0x61, 0x64, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_readability_proto_rawDescOnce sync.Once
	file_readability_proto_rawDescData = file_readability_proto_rawDesc
)

func file_readability_proto_rawDescGZIP() []byte {
	file_readability_proto_rawDescOnce.Do(func() {
		file_readability_proto_rawDescData = protoimpl.X.CompressGZIP(file_readability_proto_rawDescData)
	})
	return file_readability_proto_rawDescData
}

var file_readability_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_readability_proto_goTypes = []interface{}{
	(*GetReadableDocumentRequest)(nil),  // 0: readabilitypb.GetReadableDocumentRequest
	(*GetReadableDocumentResponse)(nil), // 1: readabilitypb.GetReadableDocumentResponse
	(*GetReadinessInfoRequest)(nil),     // 2: readabilitypb.GetReadinessInfoRequest
	(*GetReadinessInfoResponse)(nil),    // 3: readabilitypb.GetReadinessInfoResponse
}
var file_readability_proto_depIdxs = []int32{
	0, // 0: readabilitypb.Readability.GetReadableDocument:input_type -> readabilitypb.GetReadableDocumentRequest
	2, // 1: readabilitypb.Readability.GetReadinessInfo:input_type -> readabilitypb.GetReadinessInfoRequest
	1, // 2: readabilitypb.Readability.GetReadableDocument:output_type -> readabilitypb.GetReadableDocumentResponse
	3, // 3: readabilitypb.Readability.GetReadinessInfo:output_type -> readabilitypb.GetReadinessInfoResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_readability_proto_init() }
func file_readability_proto_init() {
	if File_readability_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_readability_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReadableDocumentRequest); i {
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
		file_readability_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReadableDocumentResponse); i {
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
		file_readability_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReadinessInfoRequest); i {
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
		file_readability_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReadinessInfoResponse); i {
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
			RawDescriptor: file_readability_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_readability_proto_goTypes,
		DependencyIndexes: file_readability_proto_depIdxs,
		MessageInfos:      file_readability_proto_msgTypes,
	}.Build()
	File_readability_proto = out.File
	file_readability_proto_rawDesc = nil
	file_readability_proto_goTypes = nil
	file_readability_proto_depIdxs = nil
}
