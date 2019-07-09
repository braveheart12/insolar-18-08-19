// Code generated by protoc-gen-go. DO NOT EDIT.
// source: instrumentation/introspector/introproto/publisher.proto

package introproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// EmptyArgs is just a stub for grpc methods without arguments.
type EmptyArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyArgs) Reset()         { *m = EmptyArgs{} }
func (m *EmptyArgs) String() string { return proto.CompactTextString(m) }
func (*EmptyArgs) ProtoMessage()    {}
func (*EmptyArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_38901606595998ea, []int{0}
}

func (m *EmptyArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyArgs.Unmarshal(m, b)
}
func (m *EmptyArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyArgs.Marshal(b, m, deterministic)
}
func (m *EmptyArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyArgs.Merge(m, src)
}
func (m *EmptyArgs) XXX_Size() int {
	return xxx_messageInfo_EmptyArgs.Size(m)
}
func (m *EmptyArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyArgs.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyArgs proto.InternalMessageInfo

// AllMessageFilterStats map of MessageFilterWithStat per message type.
type AllMessageFilterStats struct {
	Filters              map[string]*MessageFilterWithStat `protobuf:"bytes,1,rep,name=Filters,proto3" json:"Filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *AllMessageFilterStats) Reset()         { *m = AllMessageFilterStats{} }
func (m *AllMessageFilterStats) String() string { return proto.CompactTextString(m) }
func (*AllMessageFilterStats) ProtoMessage()    {}
func (*AllMessageFilterStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_38901606595998ea, []int{1}
}

func (m *AllMessageFilterStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllMessageFilterStats.Unmarshal(m, b)
}
func (m *AllMessageFilterStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllMessageFilterStats.Marshal(b, m, deterministic)
}
func (m *AllMessageFilterStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllMessageFilterStats.Merge(m, src)
}
func (m *AllMessageFilterStats) XXX_Size() int {
	return xxx_messageInfo_AllMessageFilterStats.Size(m)
}
func (m *AllMessageFilterStats) XXX_DiscardUnknown() {
	xxx_messageInfo_AllMessageFilterStats.DiscardUnknown(m)
}

var xxx_messageInfo_AllMessageFilterStats proto.InternalMessageInfo

func (m *AllMessageFilterStats) GetFilters() map[string]*MessageFilterWithStat {
	if m != nil {
		return m.Filters
	}
	return nil
}

// MessageFilterByType represents filter state for message type.
type MessageFilterByType struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Enable               bool     `protobuf:"varint,2,opt,name=Enable,proto3" json:"Enable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageFilterByType) Reset()         { *m = MessageFilterByType{} }
func (m *MessageFilterByType) String() string { return proto.CompactTextString(m) }
func (*MessageFilterByType) ProtoMessage()    {}
func (*MessageFilterByType) Descriptor() ([]byte, []int) {
	return fileDescriptor_38901606595998ea, []int{2}
}

func (m *MessageFilterByType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageFilterByType.Unmarshal(m, b)
}
func (m *MessageFilterByType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageFilterByType.Marshal(b, m, deterministic)
}
func (m *MessageFilterByType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageFilterByType.Merge(m, src)
}
func (m *MessageFilterByType) XXX_Size() int {
	return xxx_messageInfo_MessageFilterByType.Size(m)
}
func (m *MessageFilterByType) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageFilterByType.DiscardUnknown(m)
}

var xxx_messageInfo_MessageFilterByType proto.InternalMessageInfo

func (m *MessageFilterByType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MessageFilterByType) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

// MessageFilterWithStat represents filter state and count of filtered messages.
type MessageFilterWithStat struct {
	Enable               bool     `protobuf:"varint,2,opt,name=Enable,proto3" json:"Enable,omitempty"`
	Filtered             int64    `protobuf:"varint,3,opt,name=Filtered,proto3" json:"Filtered,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageFilterWithStat) Reset()         { *m = MessageFilterWithStat{} }
func (m *MessageFilterWithStat) String() string { return proto.CompactTextString(m) }
func (*MessageFilterWithStat) ProtoMessage()    {}
func (*MessageFilterWithStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_38901606595998ea, []int{3}
}

func (m *MessageFilterWithStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageFilterWithStat.Unmarshal(m, b)
}
func (m *MessageFilterWithStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageFilterWithStat.Marshal(b, m, deterministic)
}
func (m *MessageFilterWithStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageFilterWithStat.Merge(m, src)
}
func (m *MessageFilterWithStat) XXX_Size() int {
	return xxx_messageInfo_MessageFilterWithStat.Size(m)
}
func (m *MessageFilterWithStat) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageFilterWithStat.DiscardUnknown(m)
}

var xxx_messageInfo_MessageFilterWithStat proto.InternalMessageInfo

func (m *MessageFilterWithStat) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *MessageFilterWithStat) GetFiltered() int64 {
	if m != nil {
		return m.Filtered
	}
	return 0
}

// MessageStatByType is a counter for message type.
type MessageStatByType struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageStatByType) Reset()         { *m = MessageStatByType{} }
func (m *MessageStatByType) String() string { return proto.CompactTextString(m) }
func (*MessageStatByType) ProtoMessage()    {}
func (*MessageStatByType) Descriptor() ([]byte, []int) {
	return fileDescriptor_38901606595998ea, []int{4}
}

func (m *MessageStatByType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageStatByType.Unmarshal(m, b)
}
func (m *MessageStatByType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageStatByType.Marshal(b, m, deterministic)
}
func (m *MessageStatByType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageStatByType.Merge(m, src)
}
func (m *MessageStatByType) XXX_Size() int {
	return xxx_messageInfo_MessageStatByType.Size(m)
}
func (m *MessageStatByType) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageStatByType.DiscardUnknown(m)
}

var xxx_messageInfo_MessageStatByType proto.InternalMessageInfo

func (m *MessageStatByType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MessageStatByType) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

// AllMessageStatByType is a list of counters per message type.
type AllMessageStatByType struct {
	Counters             []*MessageStatByType `protobuf:"bytes,1,rep,name=Counters,proto3" json:"Counters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AllMessageStatByType) Reset()         { *m = AllMessageStatByType{} }
func (m *AllMessageStatByType) String() string { return proto.CompactTextString(m) }
func (*AllMessageStatByType) ProtoMessage()    {}
func (*AllMessageStatByType) Descriptor() ([]byte, []int) {
	return fileDescriptor_38901606595998ea, []int{5}
}

func (m *AllMessageStatByType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllMessageStatByType.Unmarshal(m, b)
}
func (m *AllMessageStatByType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllMessageStatByType.Marshal(b, m, deterministic)
}
func (m *AllMessageStatByType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllMessageStatByType.Merge(m, src)
}
func (m *AllMessageStatByType) XXX_Size() int {
	return xxx_messageInfo_AllMessageStatByType.Size(m)
}
func (m *AllMessageStatByType) XXX_DiscardUnknown() {
	xxx_messageInfo_AllMessageStatByType.DiscardUnknown(m)
}

var xxx_messageInfo_AllMessageStatByType proto.InternalMessageInfo

func (m *AllMessageStatByType) GetCounters() []*MessageStatByType {
	if m != nil {
		return m.Counters
	}
	return nil
}

func init() {
	proto.RegisterType((*EmptyArgs)(nil), "introproto.EmptyArgs")
	proto.RegisterType((*AllMessageFilterStats)(nil), "introproto.AllMessageFilterStats")
	proto.RegisterMapType((map[string]*MessageFilterWithStat)(nil), "introproto.AllMessageFilterStats.FiltersEntry")
	proto.RegisterType((*MessageFilterByType)(nil), "introproto.MessageFilterByType")
	proto.RegisterType((*MessageFilterWithStat)(nil), "introproto.MessageFilterWithStat")
	proto.RegisterType((*MessageStatByType)(nil), "introproto.MessageStatByType")
	proto.RegisterType((*AllMessageStatByType)(nil), "introproto.AllMessageStatByType")
}

func init() {
	proto.RegisterFile("instrumentation/introspector/introproto/publisher.proto", fileDescriptor_38901606595998ea)
}

var fileDescriptor_38901606595998ea = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0xe5, 0x84, 0x8d, 0xf6, 0x5f, 0x24, 0x36, 0x6f, 0x81, 0x28, 0x30, 0xc8, 0x7c, 0x8a,
	0x76, 0x48, 0xa4, 0x72, 0x18, 0x54, 0xe2, 0x50, 0x50, 0x01, 0x09, 0x81, 0x20, 0x43, 0xe2, 0xc4,
	0xc1, 0x05, 0x2b, 0x8b, 0x48, 0xed, 0xc8, 0x76, 0x90, 0x72, 0xe5, 0x15, 0x78, 0x21, 0xde, 0x81,
	0x1b, 0x67, 0x1e, 0x04, 0xc5, 0xce, 0x9a, 0xac, 0xa4, 0xf4, 0x96, 0x2f, 0xfe, 0xbe, 0xef, 0x67,
	0xfd, 0xfd, 0x87, 0xf3, 0x9c, 0x2b, 0x2d, 0xab, 0x15, 0xe3, 0x9a, 0xea, 0x5c, 0xf0, 0x24, 0xe7,
	0x5a, 0x0a, 0x55, 0xb2, 0xcf, 0x5a, 0x48, 0x2b, 0x4a, 0x29, 0xb4, 0x48, 0xca, 0x6a, 0x59, 0xe4,
	0xea, 0x92, 0xc9, 0xd8, 0x68, 0x0c, 0xdd, 0x59, 0x70, 0x3f, 0x13, 0x22, 0x2b, 0x58, 0x42, 0xcb,
	0x3c, 0xa1, 0x9c, 0x0b, 0x5b, 0xa5, 0xac, 0x93, 0x4c, 0x60, 0xbc, 0x58, 0x95, 0xba, 0x9e, 0xcb,
	0x4c, 0x91, 0x9f, 0x08, 0xbc, 0x79, 0x51, 0xbc, 0x61, 0x4a, 0xd1, 0x8c, 0xbd, 0xc8, 0x0b, 0xcd,
	0xe4, 0x85, 0xa6, 0x5a, 0xe1, 0x57, 0x70, 0xd3, 0x4a, 0xe5, 0xa3, 0xd0, 0x8d, 0x26, 0xd3, 0x38,
	0xee, 0x10, 0xf1, 0x60, 0x26, 0x6e, 0x03, 0x0b, 0xae, 0x65, 0x9d, 0x5e, 0xc5, 0x83, 0x4f, 0x70,
	0xab, 0x7f, 0x80, 0x0f, 0xc0, 0xfd, 0xca, 0x6a, 0x1f, 0x85, 0x28, 0x1a, 0xa7, 0xcd, 0x27, 0x3e,
	0x87, 0xbd, 0x6f, 0xb4, 0xa8, 0x98, 0xef, 0x84, 0x28, 0x9a, 0x4c, 0x4f, 0xfb, 0xa4, 0x6b, 0x98,
	0x8f, 0xb9, 0xbe, 0x6c, 0x50, 0xa9, 0xf5, 0xcf, 0x9c, 0xc7, 0x88, 0xcc, 0xe1, 0xe8, 0x9a, 0xe7,
	0x59, 0xfd, 0xa1, 0x2e, 0x19, 0xc6, 0x70, 0xe3, 0x2d, 0x5d, 0xb1, 0x16, 0x63, 0xbe, 0xf1, 0x1d,
	0xd8, 0x5f, 0x70, 0xba, 0x2c, 0x2c, 0x68, 0x94, 0xb6, 0x8a, 0xbc, 0x06, 0x6f, 0x10, 0xb3, 0x2d,
	0x80, 0x03, 0x18, 0x59, 0x27, 0xfb, 0xe2, 0xbb, 0x21, 0x8a, 0xdc, 0x74, 0xad, 0xc9, 0x53, 0x38,
	0x6c, 0xcb, 0x9a, 0x8a, 0xff, 0xdc, 0xe6, 0x18, 0xf6, 0x9e, 0x8b, 0x8a, 0x6b, 0xd3, 0xed, 0xa6,
	0x56, 0x90, 0xf7, 0x70, 0xdc, 0x0d, 0xb7, 0xd7, 0xf0, 0x04, 0x46, 0xc6, 0xd0, 0x3d, 0xc8, 0xc9,
	0xc0, 0x98, 0xba, 0x40, 0xba, 0xb6, 0x4f, 0x7f, 0x3b, 0x30, 0x7e, 0x77, 0xb5, 0x2f, 0x58, 0xc3,
	0xe1, 0x05, 0xd3, 0xad, 0x5f, 0xd9, 0x6b, 0xe3, 0x87, 0x5b, 0x47, 0x6e, 0xdb, 0x82, 0x5d, 0x06,
	0x72, 0xf2, 0xfd, 0xd7, 0x9f, 0x1f, 0xce, 0x5d, 0x82, 0x13, 0xb5, 0xd9, 0x3e, 0x43, 0x67, 0x98,
	0x03, 0x7e, 0xb9, 0xf9, 0x5f, 0x61, 0xaf, 0xdf, 0xba, 0xde, 0xca, 0xe0, 0x74, 0xe7, 0xaa, 0x91,
	0x07, 0x06, 0xe7, 0x93, 0xa3, 0x24, 0xfb, 0xa7, 0xb6, 0xe1, 0x65, 0x70, 0xbb, 0xc7, 0x33, 0x8f,
	0xb9, 0x05, 0x16, 0x0e, 0xc3, 0xba, 0x49, 0x92, 0x7b, 0x86, 0xe5, 0x91, 0x83, 0x3e, 0xab, 0x39,
	0x9f, 0xa1, 0xb3, 0xe5, 0xbe, 0x09, 0x3e, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x97, 0xa4, 0xe9,
	0x22, 0xba, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PublisherClient is the client API for Publisher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PublisherClient interface {
	// SetMessagesFilter enables/disables messages publishing by type.
	SetMessagesFilter(ctx context.Context, in *MessageFilterByType, opts ...grpc.CallOption) (*MessageFilterByType, error)
	// GetMessagesFilters returns map with filter state for every message type.
	GetMessagesFilters(ctx context.Context, in *EmptyArgs, opts ...grpc.CallOption) (*AllMessageFilterStats, error)
	// GetMessagesStat returns statistic for published messages by type.
	GetMessagesStat(ctx context.Context, in *EmptyArgs, opts ...grpc.CallOption) (*AllMessageStatByType, error)
}

type publisherClient struct {
	cc *grpc.ClientConn
}

func NewPublisherClient(cc *grpc.ClientConn) PublisherClient {
	return &publisherClient{cc}
}

func (c *publisherClient) SetMessagesFilter(ctx context.Context, in *MessageFilterByType, opts ...grpc.CallOption) (*MessageFilterByType, error) {
	out := new(MessageFilterByType)
	err := c.cc.Invoke(ctx, "/introproto.Publisher/SetMessagesFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publisherClient) GetMessagesFilters(ctx context.Context, in *EmptyArgs, opts ...grpc.CallOption) (*AllMessageFilterStats, error) {
	out := new(AllMessageFilterStats)
	err := c.cc.Invoke(ctx, "/introproto.Publisher/GetMessagesFilters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publisherClient) GetMessagesStat(ctx context.Context, in *EmptyArgs, opts ...grpc.CallOption) (*AllMessageStatByType, error) {
	out := new(AllMessageStatByType)
	err := c.cc.Invoke(ctx, "/introproto.Publisher/GetMessagesStat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublisherServer is the server API for Publisher service.
type PublisherServer interface {
	// SetMessagesFilter enables/disables messages publishing by type.
	SetMessagesFilter(context.Context, *MessageFilterByType) (*MessageFilterByType, error)
	// GetMessagesFilters returns map with filter state for every message type.
	GetMessagesFilters(context.Context, *EmptyArgs) (*AllMessageFilterStats, error)
	// GetMessagesStat returns statistic for published messages by type.
	GetMessagesStat(context.Context, *EmptyArgs) (*AllMessageStatByType, error)
}

// UnimplementedPublisherServer can be embedded to have forward compatible implementations.
type UnimplementedPublisherServer struct {
}

func (*UnimplementedPublisherServer) SetMessagesFilter(ctx context.Context, req *MessageFilterByType) (*MessageFilterByType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMessagesFilter not implemented")
}
func (*UnimplementedPublisherServer) GetMessagesFilters(ctx context.Context, req *EmptyArgs) (*AllMessageFilterStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessagesFilters not implemented")
}
func (*UnimplementedPublisherServer) GetMessagesStat(ctx context.Context, req *EmptyArgs) (*AllMessageStatByType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessagesStat not implemented")
}

func RegisterPublisherServer(s *grpc.Server, srv PublisherServer) {
	s.RegisterService(&_Publisher_serviceDesc, srv)
}

func _Publisher_SetMessagesFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageFilterByType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherServer).SetMessagesFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/introproto.Publisher/SetMessagesFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherServer).SetMessagesFilter(ctx, req.(*MessageFilterByType))
	}
	return interceptor(ctx, in, info, handler)
}

func _Publisher_GetMessagesFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherServer).GetMessagesFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/introproto.Publisher/GetMessagesFilters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherServer).GetMessagesFilters(ctx, req.(*EmptyArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Publisher_GetMessagesStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherServer).GetMessagesStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/introproto.Publisher/GetMessagesStat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherServer).GetMessagesStat(ctx, req.(*EmptyArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _Publisher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "introproto.Publisher",
	HandlerType: (*PublisherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetMessagesFilter",
			Handler:    _Publisher_SetMessagesFilter_Handler,
		},
		{
			MethodName: "GetMessagesFilters",
			Handler:    _Publisher_GetMessagesFilters_Handler,
		},
		{
			MethodName: "GetMessagesStat",
			Handler:    _Publisher_GetMessagesStat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "instrumentation/introspector/introproto/publisher.proto",
}
