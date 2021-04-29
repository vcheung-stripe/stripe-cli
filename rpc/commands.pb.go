// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: commands.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_commands_proto protoreflect.FileDescriptor

var file_commands_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x12, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x83, 0x01, 0x0a, 0x09, 0x53, 0x74, 0x72,
	0x69, 0x70, 0x65, 0x43, 0x4c, 0x49, 0x12, 0x40, 0x0a, 0x0b, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x22,
	0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72,
	0x69, 0x70, 0x65, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x2d, 0x63, 0x6c, 0x69, 0x2f, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_commands_proto_goTypes = []interface{}{
	(*SamplesListRequest)(nil),  // 0: rpc.SamplesListRequest
	(*VersionRequest)(nil),      // 1: rpc.VersionRequest
	(*SamplesListResponse)(nil), // 2: rpc.SamplesListResponse
	(*VersionResponse)(nil),     // 3: rpc.VersionResponse
}
var file_commands_proto_depIdxs = []int32{
	0, // 0: rpc.StripeCLI.SamplesList:input_type -> rpc.SamplesListRequest
	1, // 1: rpc.StripeCLI.Version:input_type -> rpc.VersionRequest
	2, // 2: rpc.StripeCLI.SamplesList:output_type -> rpc.SamplesListResponse
	3, // 3: rpc.StripeCLI.Version:output_type -> rpc.VersionResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commands_proto_init() }
func file_commands_proto_init() {
	if File_commands_proto != nil {
		return
	}
	file_samples_list_proto_init()
	file_version_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commands_proto_goTypes,
		DependencyIndexes: file_commands_proto_depIdxs,
	}.Build()
	File_commands_proto = out.File
	file_commands_proto_rawDesc = nil
	file_commands_proto_goTypes = nil
	file_commands_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StripeCLIClient is the client API for StripeCLI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StripeCLIClient interface {
	// Get a list of available Stripe samples. Like `stripe samples list`.
	SamplesList(ctx context.Context, in *SamplesListRequest, opts ...grpc.CallOption) (*SamplesListResponse, error)
	// Get the version of the Stripe CLI. Like `stripe version`.
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type stripeCLIClient struct {
	cc grpc.ClientConnInterface
}

func NewStripeCLIClient(cc grpc.ClientConnInterface) StripeCLIClient {
	return &stripeCLIClient{cc}
}

func (c *stripeCLIClient) SamplesList(ctx context.Context, in *SamplesListRequest, opts ...grpc.CallOption) (*SamplesListResponse, error) {
	out := new(SamplesListResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/SamplesList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StripeCLIServer is the server API for StripeCLI service.
type StripeCLIServer interface {
	// Get a list of available Stripe samples. Like `stripe samples list`.
	SamplesList(context.Context, *SamplesListRequest) (*SamplesListResponse, error)
	// Get the version of the Stripe CLI. Like `stripe version`.
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}

// UnimplementedStripeCLIServer can be embedded to have forward compatible implementations.
type UnimplementedStripeCLIServer struct {
}

func (*UnimplementedStripeCLIServer) SamplesList(context.Context, *SamplesListRequest) (*SamplesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SamplesList not implemented")
}
func (*UnimplementedStripeCLIServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}

func RegisterStripeCLIServer(s *grpc.Server, srv StripeCLIServer) {
	s.RegisterService(&_StripeCLI_serviceDesc, srv)
}

func _StripeCLI_SamplesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SamplesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).SamplesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/SamplesList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).SamplesList(ctx, req.(*SamplesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StripeCLI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.StripeCLI",
	HandlerType: (*StripeCLIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SamplesList",
			Handler:    _StripeCLI_SamplesList_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _StripeCLI_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commands.proto",
}
