// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: grpc/coordinate_transporter.proto

package grpc

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

type PostCoordinateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X        int32  `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y        int32  `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	ClientId string `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *PostCoordinateRequest) Reset() {
	*x = PostCoordinateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_coordinate_transporter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostCoordinateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCoordinateRequest) ProtoMessage() {}

func (x *PostCoordinateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_coordinate_transporter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCoordinateRequest.ProtoReflect.Descriptor instead.
func (*PostCoordinateRequest) Descriptor() ([]byte, []int) {
	return file_grpc_coordinate_transporter_proto_rawDescGZIP(), []int{0}
}

func (x *PostCoordinateRequest) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *PostCoordinateRequest) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *PostCoordinateRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type PostCoordinateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *PostCoordinateResponse) Reset() {
	*x = PostCoordinateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_coordinate_transporter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostCoordinateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCoordinateResponse) ProtoMessage() {}

func (x *PostCoordinateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_coordinate_transporter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCoordinateResponse.ProtoReflect.Descriptor instead.
func (*PostCoordinateResponse) Descriptor() ([]byte, []int) {
	return file_grpc_coordinate_transporter_proto_rawDescGZIP(), []int{1}
}

func (x *PostCoordinateResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_grpc_coordinate_transporter_proto protoreflect.FileDescriptor

var file_grpc_coordinate_transporter_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x15, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32,
	0x61, 0x0a, 0x15, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74,
	0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_coordinate_transporter_proto_rawDescOnce sync.Once
	file_grpc_coordinate_transporter_proto_rawDescData = file_grpc_coordinate_transporter_proto_rawDesc
)

func file_grpc_coordinate_transporter_proto_rawDescGZIP() []byte {
	file_grpc_coordinate_transporter_proto_rawDescOnce.Do(func() {
		file_grpc_coordinate_transporter_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_coordinate_transporter_proto_rawDescData)
	})
	return file_grpc_coordinate_transporter_proto_rawDescData
}

var file_grpc_coordinate_transporter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_coordinate_transporter_proto_goTypes = []interface{}{
	(*PostCoordinateRequest)(nil),  // 0: PostCoordinateRequest
	(*PostCoordinateResponse)(nil), // 1: PostCoordinateResponse
}
var file_grpc_coordinate_transporter_proto_depIdxs = []int32{
	0, // 0: CoordinateTransporter.PostCoordinates:input_type -> PostCoordinateRequest
	1, // 1: CoordinateTransporter.PostCoordinates:output_type -> PostCoordinateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_coordinate_transporter_proto_init() }
func file_grpc_coordinate_transporter_proto_init() {
	if File_grpc_coordinate_transporter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_coordinate_transporter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostCoordinateRequest); i {
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
		file_grpc_coordinate_transporter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostCoordinateResponse); i {
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
			RawDescriptor: file_grpc_coordinate_transporter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_coordinate_transporter_proto_goTypes,
		DependencyIndexes: file_grpc_coordinate_transporter_proto_depIdxs,
		MessageInfos:      file_grpc_coordinate_transporter_proto_msgTypes,
	}.Build()
	File_grpc_coordinate_transporter_proto = out.File
	file_grpc_coordinate_transporter_proto_rawDesc = nil
	file_grpc_coordinate_transporter_proto_goTypes = nil
	file_grpc_coordinate_transporter_proto_depIdxs = nil
}
