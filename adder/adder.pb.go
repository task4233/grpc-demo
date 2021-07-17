// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.9.0
// source: adder.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type NumMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *NumMessage) Reset() {
	*x = NumMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumMessage) ProtoMessage() {}

func (x *NumMessage) ProtoReflect() protoreflect.Message {
	mi := &file_adder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumMessage.ProtoReflect.Descriptor instead.
func (*NumMessage) Descriptor() ([]byte, []int) {
	return file_adder_proto_rawDescGZIP(), []int{0}
}

func (x *NumMessage) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_adder_proto protoreflect.FileDescriptor

var file_adder_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x64, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x64, 0x64, 0x65, 0x72, 0x22, 0x22, 0x0a, 0x0a, 0x4e, 0x75, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x3d, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12,
	0x11, 0x2e, 0x61, 0x64, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x75, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x11, 0x2e, 0x61, 0x64, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x75, 0x6d, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_adder_proto_rawDescOnce sync.Once
	file_adder_proto_rawDescData = file_adder_proto_rawDesc
)

func file_adder_proto_rawDescGZIP() []byte {
	file_adder_proto_rawDescOnce.Do(func() {
		file_adder_proto_rawDescData = protoimpl.X.CompressGZIP(file_adder_proto_rawDescData)
	})
	return file_adder_proto_rawDescData
}

var file_adder_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_adder_proto_goTypes = []interface{}{
	(*NumMessage)(nil), // 0: adder.NumMessage
}
var file_adder_proto_depIdxs = []int32{
	0, // 0: adder.AdderService.Add:input_type -> adder.NumMessage
	0, // 1: adder.AdderService.Add:output_type -> adder.NumMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_adder_proto_init() }
func file_adder_proto_init() {
	if File_adder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_adder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumMessage); i {
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
			RawDescriptor: file_adder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_adder_proto_goTypes,
		DependencyIndexes: file_adder_proto_depIdxs,
		MessageInfos:      file_adder_proto_msgTypes,
	}.Build()
	File_adder_proto = out.File
	file_adder_proto_rawDesc = nil
	file_adder_proto_goTypes = nil
	file_adder_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdderServiceClient is the client API for AdderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdderServiceClient interface {
	Add(ctx context.Context, in *NumMessage, opts ...grpc.CallOption) (*NumMessage, error)
}

type adderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdderServiceClient(cc grpc.ClientConnInterface) AdderServiceClient {
	return &adderServiceClient{cc}
}

func (c *adderServiceClient) Add(ctx context.Context, in *NumMessage, opts ...grpc.CallOption) (*NumMessage, error) {
	out := new(NumMessage)
	err := c.cc.Invoke(ctx, "/adder.AdderService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdderServiceServer is the server API for AdderService service.
type AdderServiceServer interface {
	Add(context.Context, *NumMessage) (*NumMessage, error)
}

// UnimplementedAdderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdderServiceServer struct {
}

func (*UnimplementedAdderServiceServer) Add(context.Context, *NumMessage) (*NumMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterAdderServiceServer(s *grpc.Server, srv AdderServiceServer) {
	s.RegisterService(&_AdderService_serviceDesc, srv)
}

func _AdderService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdderServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adder.AdderService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdderServiceServer).Add(ctx, req.(*NumMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "adder.AdderService",
	HandlerType: (*AdderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _AdderService_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "adder.proto",
}
