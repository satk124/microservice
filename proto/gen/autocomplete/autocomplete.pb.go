// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.0
// source: autocomplete.proto

package autocomplete

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Size int32  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autocomplete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_autocomplete_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_autocomplete_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Request) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []string `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autocomplete_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_autocomplete_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_autocomplete_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetResult() []string {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_autocomplete_proto protoreflect.FileDescriptor

var file_autocomplete_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x22, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x2a, 0x0a, 0x0c, 0x41, 0x75, 0x74,
	0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x3b, 0x61, 0x75, 0x74, 0x6f, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_autocomplete_proto_rawDescOnce sync.Once
	file_autocomplete_proto_rawDescData = file_autocomplete_proto_rawDesc
)

func file_autocomplete_proto_rawDescGZIP() []byte {
	file_autocomplete_proto_rawDescOnce.Do(func() {
		file_autocomplete_proto_rawDescData = protoimpl.X.CompressGZIP(file_autocomplete_proto_rawDescData)
	})
	return file_autocomplete_proto_rawDescData
}

var file_autocomplete_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_autocomplete_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: Request
	(*Response)(nil), // 1: Response
}
var file_autocomplete_proto_depIdxs = []int32{
	0, // 0: Autocomplete.Get:input_type -> Request
	1, // 1: Autocomplete.Get:output_type -> Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_autocomplete_proto_init() }
func file_autocomplete_proto_init() {
	if File_autocomplete_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_autocomplete_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_autocomplete_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_autocomplete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_autocomplete_proto_goTypes,
		DependencyIndexes: file_autocomplete_proto_depIdxs,
		MessageInfos:      file_autocomplete_proto_msgTypes,
	}.Build()
	File_autocomplete_proto = out.File
	file_autocomplete_proto_rawDesc = nil
	file_autocomplete_proto_goTypes = nil
	file_autocomplete_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AutocompleteClient is the client API for Autocomplete service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AutocompleteClient interface {
	Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type autocompleteClient struct {
	cc grpc.ClientConnInterface
}

func NewAutocompleteClient(cc grpc.ClientConnInterface) AutocompleteClient {
	return &autocompleteClient{cc}
}

func (c *autocompleteClient) Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Autocomplete/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutocompleteServer is the server API for Autocomplete service.
type AutocompleteServer interface {
	Get(context.Context, *Request) (*Response, error)
}

// UnimplementedAutocompleteServer can be embedded to have forward compatible implementations.
type UnimplementedAutocompleteServer struct {
}

func (*UnimplementedAutocompleteServer) Get(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterAutocompleteServer(s *grpc.Server, srv AutocompleteServer) {
	s.RegisterService(&_Autocomplete_serviceDesc, srv)
}

func _Autocomplete_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutocompleteServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Autocomplete/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutocompleteServer).Get(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Autocomplete_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Autocomplete",
	HandlerType: (*AutocompleteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Autocomplete_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "autocomplete.proto",
}
