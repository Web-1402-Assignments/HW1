// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: Auth_Server.proto

package Auth_Server

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

type PQ_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce     string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	MessageId uint32 `protobuf:"varint,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *PQ_Request) Reset() {
	*x = PQ_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auth_Server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PQ_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PQ_Request) ProtoMessage() {}

func (x *PQ_Request) ProtoReflect() protoreflect.Message {
	mi := &file_Auth_Server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PQ_Request.ProtoReflect.Descriptor instead.
func (*PQ_Request) Descriptor() ([]byte, []int) {
	return file_Auth_Server_proto_rawDescGZIP(), []int{0}
}

func (x *PQ_Request) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *PQ_Request) GetMessageId() uint32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

type PQ_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce       string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce string `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	MessageId   uint32 `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	P           uint32 `protobuf:"varint,4,opt,name=p,proto3" json:"p,omitempty"`
	G           int32  `protobuf:"varint,5,opt,name=g,proto3" json:"g,omitempty"`
}

func (x *PQ_Response) Reset() {
	*x = PQ_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auth_Server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PQ_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PQ_Response) ProtoMessage() {}

func (x *PQ_Response) ProtoReflect() protoreflect.Message {
	mi := &file_Auth_Server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PQ_Response.ProtoReflect.Descriptor instead.
func (*PQ_Response) Descriptor() ([]byte, []int) {
	return file_Auth_Server_proto_rawDescGZIP(), []int{1}
}

func (x *PQ_Response) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *PQ_Response) GetServerNonce() string {
	if x != nil {
		return x.ServerNonce
	}
	return ""
}

func (x *PQ_Response) GetMessageId() uint32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *PQ_Response) GetP() uint32 {
	if x != nil {
		return x.P
	}
	return 0
}

func (x *PQ_Response) GetG() int32 {
	if x != nil {
		return x.G
	}
	return 0
}

type DH_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce       string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce string `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	MessageId   uint32 `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	A           int32  `protobuf:"varint,4,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *DH_Request) Reset() {
	*x = DH_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auth_Server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DH_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DH_Request) ProtoMessage() {}

func (x *DH_Request) ProtoReflect() protoreflect.Message {
	mi := &file_Auth_Server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DH_Request.ProtoReflect.Descriptor instead.
func (*DH_Request) Descriptor() ([]byte, []int) {
	return file_Auth_Server_proto_rawDescGZIP(), []int{2}
}

func (x *DH_Request) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *DH_Request) GetServerNonce() string {
	if x != nil {
		return x.ServerNonce
	}
	return ""
}

func (x *DH_Request) GetMessageId() uint32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *DH_Request) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

type DH_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce       string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce string `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	MessageId   uint32 `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	B           int32  `protobuf:"varint,4,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *DH_Response) Reset() {
	*x = DH_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auth_Server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DH_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DH_Response) ProtoMessage() {}

func (x *DH_Response) ProtoReflect() protoreflect.Message {
	mi := &file_Auth_Server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DH_Response.ProtoReflect.Descriptor instead.
func (*DH_Response) Descriptor() ([]byte, []int) {
	return file_Auth_Server_proto_rawDescGZIP(), []int{3}
}

func (x *DH_Response) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *DH_Response) GetServerNonce() string {
	if x != nil {
		return x.ServerNonce
	}
	return ""
}

func (x *DH_Response) GetMessageId() uint32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *DH_Response) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

var File_Auth_Server_proto protoreflect.FileDescriptor

var file_Auth_Server_proto_rawDesc = []byte{
	0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0a, 0x50, 0x51, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x0b, 0x50, 0x51, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x0c,
	0x0a, 0x01, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x70, 0x12, 0x0c, 0x0a, 0x01,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x67, 0x22, 0x72, 0x0a, 0x0a, 0x44, 0x48,
	0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x6e, 0x63,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x22, 0x73,
	0x0a, 0x0b, 0x44, 0x48, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x62, 0x32, 0x64, 0x0a, 0x06, 0x52, 0x65, 0x71, 0x5f, 0x44, 0x48, 0x12, 0x2c, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x50, 0x51, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0b,
	0x2e, 0x50, 0x51, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x50, 0x51,
	0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x44, 0x48, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0b, 0x2e, 0x44,
	0x48, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x44, 0x48, 0x5f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x57, 0x65, 0x62, 0x2d, 0x31, 0x34, 0x30, 0x32,
	0x2d, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x48, 0x57, 0x31,
	0x2f, 0x41, 0x75, 0x74, 0x68, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Auth_Server_proto_rawDescOnce sync.Once
	file_Auth_Server_proto_rawDescData = file_Auth_Server_proto_rawDesc
)

func file_Auth_Server_proto_rawDescGZIP() []byte {
	file_Auth_Server_proto_rawDescOnce.Do(func() {
		file_Auth_Server_proto_rawDescData = protoimpl.X.CompressGZIP(file_Auth_Server_proto_rawDescData)
	})
	return file_Auth_Server_proto_rawDescData
}

var file_Auth_Server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Auth_Server_proto_goTypes = []interface{}{
	(*PQ_Request)(nil),  // 0: PQ_Request
	(*PQ_Response)(nil), // 1: PQ_Response
	(*DH_Request)(nil),  // 2: DH_Request
	(*DH_Response)(nil), // 3: DH_Response
}
var file_Auth_Server_proto_depIdxs = []int32{
	0, // 0: Req_DH.GetPQResponse:input_type -> PQ_Request
	2, // 1: Req_DH.GetDHResponse:input_type -> DH_Request
	1, // 2: Req_DH.GetPQResponse:output_type -> PQ_Response
	3, // 3: Req_DH.GetDHResponse:output_type -> DH_Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Auth_Server_proto_init() }
func file_Auth_Server_proto_init() {
	if File_Auth_Server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Auth_Server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PQ_Request); i {
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
		file_Auth_Server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PQ_Response); i {
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
		file_Auth_Server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DH_Request); i {
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
		file_Auth_Server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DH_Response); i {
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
			RawDescriptor: file_Auth_Server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Auth_Server_proto_goTypes,
		DependencyIndexes: file_Auth_Server_proto_depIdxs,
		MessageInfos:      file_Auth_Server_proto_msgTypes,
	}.Build()
	File_Auth_Server_proto = out.File
	file_Auth_Server_proto_rawDesc = nil
	file_Auth_Server_proto_goTypes = nil
	file_Auth_Server_proto_depIdxs = nil
}