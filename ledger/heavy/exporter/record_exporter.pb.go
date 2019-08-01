// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ledger/heavy/exporter/record_exporter.proto

package exporter

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_insolar_insolar_insolar "github.com/insolar/insolar/insolar"
	record "github.com/insolar/insolar/insolar/record"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type GetRecords struct {
	Polymorph    uint32                                         `protobuf:"varint,16,opt,name=Polymorph,proto3" json:"Polymorph,omitempty"`
	PulseNumber  github_com_insolar_insolar_insolar.PulseNumber `protobuf:"bytes,20,opt,name=PulseNumber,proto3,customtype=github.com/insolar/insolar/insolar.PulseNumber" json:"PulseNumber"`
	RecordNumber uint32                                         `protobuf:"varint,21,opt,name=RecordNumber,proto3" json:"RecordNumber,omitempty"`
	Count        uint32                                         `protobuf:"varint,22,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (m *GetRecords) Reset()      { *m = GetRecords{} }
func (*GetRecords) ProtoMessage() {}
func (*GetRecords) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfb4fbd68f50939d, []int{0}
}
func (m *GetRecords) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetRecords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetRecords.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetRecords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecords.Merge(m, src)
}
func (m *GetRecords) XXX_Size() int {
	return m.Size()
}
func (m *GetRecords) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecords.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecords proto.InternalMessageInfo

func (m *GetRecords) GetPolymorph() uint32 {
	if m != nil {
		return m.Polymorph
	}
	return 0
}

func (m *GetRecords) GetRecordNumber() uint32 {
	if m != nil {
		return m.RecordNumber
	}
	return 0
}

func (m *GetRecords) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Record struct {
	Polymorph    uint32          `protobuf:"varint,16,opt,name=Polymorph,proto3" json:"Polymorph,omitempty"`
	RecordNumber uint32          `protobuf:"varint,20,opt,name=RecordNumber,proto3" json:"RecordNumber,omitempty"`
	Record       record.Material `protobuf:"bytes,21,opt,name=Record,proto3" json:"Record"`
}

func (m *Record) Reset()      { *m = Record{} }
func (*Record) ProtoMessage() {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfb4fbd68f50939d, []int{1}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Record.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return m.Size()
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetPolymorph() uint32 {
	if m != nil {
		return m.Polymorph
	}
	return 0
}

func (m *Record) GetRecordNumber() uint32 {
	if m != nil {
		return m.RecordNumber
	}
	return 0
}

func (m *Record) GetRecord() record.Material {
	if m != nil {
		return m.Record
	}
	return record.Material{}
}

func init() {
	proto.RegisterType((*GetRecords)(nil), "exporter.GetRecords")
	proto.RegisterType((*Record)(nil), "exporter.Record")
}

func init() {
	proto.RegisterFile("ledger/heavy/exporter/record_exporter.proto", fileDescriptor_dfb4fbd68f50939d)
}

var fileDescriptor_dfb4fbd68f50939d = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xce, 0x49, 0x4d, 0x49,
	0x4f, 0x2d, 0xd2, 0xcf, 0x48, 0x4d, 0x2c, 0xab, 0xd4, 0x4f, 0xad, 0x28, 0xc8, 0x2f, 0x2a, 0x49,
	0x2d, 0xd2, 0x2f, 0x4a, 0x4d, 0xce, 0x2f, 0x4a, 0x89, 0x87, 0xf1, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x38, 0x60, 0x7c, 0x29, 0xdd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc,
	0x5c, 0xfd, 0xf4, 0xfc, 0xf4, 0x7c, 0x7d, 0xb0, 0x82, 0xa4, 0xd2, 0x34, 0x30, 0x0f, 0xcc, 0x01,
	0xb3, 0x20, 0x1a, 0xa5, 0xcc, 0x90, 0x94, 0x67, 0xe6, 0x15, 0xe7, 0xe7, 0x24, 0x16, 0x61, 0xd0,
	0x10, 0x2b, 0xa1, 0x14, 0x44, 0x9f, 0xd2, 0x3e, 0x46, 0x2e, 0x2e, 0xf7, 0xd4, 0x92, 0x20, 0xb0,
	0x58, 0xb1, 0x90, 0x0c, 0x17, 0x67, 0x40, 0x7e, 0x4e, 0x65, 0x6e, 0x7e, 0x51, 0x41, 0x86, 0x84,
	0x80, 0x02, 0xa3, 0x06, 0x6f, 0x10, 0x42, 0x40, 0x28, 0x82, 0x8b, 0x3b, 0xa0, 0x34, 0xa7, 0x38,
	0xd5, 0xaf, 0x34, 0x37, 0x29, 0xb5, 0x48, 0x42, 0x44, 0x81, 0x51, 0x83, 0xc7, 0xc9, 0xec, 0xc4,
	0x3d, 0x79, 0x86, 0x5b, 0xf7, 0xe4, 0xf5, 0x08, 0xbb, 0x40, 0x0f, 0x49, 0x77, 0x10, 0xb2, 0x51,
	0x42, 0x4a, 0x5c, 0x3c, 0x10, 0x27, 0x40, 0x8d, 0x16, 0x05, 0x5b, 0x8d, 0x22, 0x26, 0x24, 0xc2,
	0xc5, 0xea, 0x9c, 0x5f, 0x9a, 0x57, 0x22, 0x21, 0x06, 0x96, 0x84, 0x70, 0x94, 0xaa, 0xb8, 0xd8,
	0x20, 0xaa, 0x08, 0xb8, 0x1d, 0xdd, 0x06, 0x11, 0x2c, 0x36, 0xe8, 0xc1, 0xcc, 0x02, 0xdb, 0xcf,
	0x6d, 0x24, 0xa0, 0x07, 0x0d, 0x2b, 0xdf, 0xc4, 0x92, 0xd4, 0xa2, 0xcc, 0xc4, 0x1c, 0x27, 0x16,
	0x90, 0x67, 0x83, 0xa0, 0xaa, 0x8c, 0xdc, 0xb8, 0xf8, 0x20, 0x2c, 0x57, 0x68, 0xac, 0x09, 0x99,
	0x70, 0xb1, 0x41, 0xd8, 0x42, 0x22, 0x7a, 0xf0, 0xa8, 0x45, 0x84, 0xaf, 0x94, 0x00, 0x42, 0x14,
	0x22, 0xa4, 0xc4, 0x60, 0xc0, 0xe8, 0x64, 0x72, 0xe1, 0xa1, 0x1c, 0xc3, 0x8d, 0x87, 0x72, 0x0c,
	0x1f, 0x1e, 0xca, 0x31, 0x36, 0x3c, 0x92, 0x63, 0x5c, 0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39,
	0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x7c, 0xf1, 0x48, 0x8e, 0xe1, 0xc3, 0x23,
	0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x89,
	0x0d, 0x1c, 0x83, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x91, 0xa4, 0xb9, 0x61, 0x02,
	0x00, 0x00,
}

func (this *GetRecords) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetRecords)
	if !ok {
		that2, ok := that.(GetRecords)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Polymorph != that1.Polymorph {
		return false
	}
	if !this.PulseNumber.Equal(that1.PulseNumber) {
		return false
	}
	if this.RecordNumber != that1.RecordNumber {
		return false
	}
	if this.Count != that1.Count {
		return false
	}
	return true
}
func (this *Record) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Record)
	if !ok {
		that2, ok := that.(Record)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Polymorph != that1.Polymorph {
		return false
	}
	if this.RecordNumber != that1.RecordNumber {
		return false
	}
	if !this.Record.Equal(&that1.Record) {
		return false
	}
	return true
}
func (this *GetRecords) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&exporter.GetRecords{")
	s = append(s, "Polymorph: "+fmt.Sprintf("%#v", this.Polymorph)+",\n")
	s = append(s, "PulseNumber: "+fmt.Sprintf("%#v", this.PulseNumber)+",\n")
	s = append(s, "RecordNumber: "+fmt.Sprintf("%#v", this.RecordNumber)+",\n")
	s = append(s, "Count: "+fmt.Sprintf("%#v", this.Count)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Record) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&exporter.Record{")
	s = append(s, "Polymorph: "+fmt.Sprintf("%#v", this.Polymorph)+",\n")
	s = append(s, "RecordNumber: "+fmt.Sprintf("%#v", this.RecordNumber)+",\n")
	s = append(s, "Record: "+strings.Replace(this.Record.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringRecordExporter(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RecordExporterClient is the client API for RecordExporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecordExporterClient interface {
	Export(ctx context.Context, in *GetRecords, opts ...grpc.CallOption) (RecordExporter_ExportClient, error)
}

type recordExporterClient struct {
	cc *grpc.ClientConn
}

func NewRecordExporterClient(cc *grpc.ClientConn) RecordExporterClient {
	return &recordExporterClient{cc}
}

func (c *recordExporterClient) Export(ctx context.Context, in *GetRecords, opts ...grpc.CallOption) (RecordExporter_ExportClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RecordExporter_serviceDesc.Streams[0], "/exporter.RecordExporter/Export", opts...)
	if err != nil {
		return nil, err
	}
	x := &recordExporterExportClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RecordExporter_ExportClient interface {
	Recv() (*Record, error)
	grpc.ClientStream
}

type recordExporterExportClient struct {
	grpc.ClientStream
}

func (x *recordExporterExportClient) Recv() (*Record, error) {
	m := new(Record)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RecordExporterServer is the server API for RecordExporter service.
type RecordExporterServer interface {
	Export(*GetRecords, RecordExporter_ExportServer) error
}

func RegisterRecordExporterServer(s *grpc.Server, srv RecordExporterServer) {
	s.RegisterService(&_RecordExporter_serviceDesc, srv)
}

func _RecordExporter_Export_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRecords)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RecordExporterServer).Export(m, &recordExporterExportServer{stream})
}

type RecordExporter_ExportServer interface {
	Send(*Record) error
	grpc.ServerStream
}

type recordExporterExportServer struct {
	grpc.ServerStream
}

func (x *recordExporterExportServer) Send(m *Record) error {
	return x.ServerStream.SendMsg(m)
}

var _RecordExporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "exporter.RecordExporter",
	HandlerType: (*RecordExporterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Export",
			Handler:       _RecordExporter_Export_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ledger/heavy/exporter/record_exporter.proto",
}

func (m *GetRecords) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetRecords) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Polymorph != 0 {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintRecordExporter(dAtA, i, uint64(m.Polymorph))
	}
	dAtA[i] = 0xa2
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintRecordExporter(dAtA, i, uint64(m.PulseNumber.Size()))
	n1, err := m.PulseNumber.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.RecordNumber != 0 {
		dAtA[i] = 0xa8
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintRecordExporter(dAtA, i, uint64(m.RecordNumber))
	}
	if m.Count != 0 {
		dAtA[i] = 0xb0
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintRecordExporter(dAtA, i, uint64(m.Count))
	}
	return i, nil
}

func (m *Record) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Polymorph != 0 {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintRecordExporter(dAtA, i, uint64(m.Polymorph))
	}
	if m.RecordNumber != 0 {
		dAtA[i] = 0xa0
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintRecordExporter(dAtA, i, uint64(m.RecordNumber))
	}
	dAtA[i] = 0xaa
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintRecordExporter(dAtA, i, uint64(m.Record.Size()))
	n2, err := m.Record.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func encodeVarintRecordExporter(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetRecords) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Polymorph != 0 {
		n += 2 + sovRecordExporter(uint64(m.Polymorph))
	}
	l = m.PulseNumber.Size()
	n += 2 + l + sovRecordExporter(uint64(l))
	if m.RecordNumber != 0 {
		n += 2 + sovRecordExporter(uint64(m.RecordNumber))
	}
	if m.Count != 0 {
		n += 2 + sovRecordExporter(uint64(m.Count))
	}
	return n
}

func (m *Record) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Polymorph != 0 {
		n += 2 + sovRecordExporter(uint64(m.Polymorph))
	}
	if m.RecordNumber != 0 {
		n += 2 + sovRecordExporter(uint64(m.RecordNumber))
	}
	l = m.Record.Size()
	n += 2 + l + sovRecordExporter(uint64(l))
	return n
}

func sovRecordExporter(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRecordExporter(x uint64) (n int) {
	return sovRecordExporter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GetRecords) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetRecords{`,
		`Polymorph:` + fmt.Sprintf("%v", this.Polymorph) + `,`,
		`PulseNumber:` + fmt.Sprintf("%v", this.PulseNumber) + `,`,
		`RecordNumber:` + fmt.Sprintf("%v", this.RecordNumber) + `,`,
		`Count:` + fmt.Sprintf("%v", this.Count) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Record) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Record{`,
		`Polymorph:` + fmt.Sprintf("%v", this.Polymorph) + `,`,
		`RecordNumber:` + fmt.Sprintf("%v", this.RecordNumber) + `,`,
		`Record:` + strings.Replace(strings.Replace(this.Record.String(), "Material", "record.Material", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRecordExporter(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GetRecords) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecordExporter
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetRecords: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetRecords: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Polymorph", wireType)
			}
			m.Polymorph = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecordExporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Polymorph |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PulseNumber", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecordExporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRecordExporter
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRecordExporter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PulseNumber.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 21:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordNumber", wireType)
			}
			m.RecordNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecordExporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RecordNumber |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 22:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecordExporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRecordExporter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRecordExporter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRecordExporter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Record) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecordExporter
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Polymorph", wireType)
			}
			m.Polymorph = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecordExporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Polymorph |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordNumber", wireType)
			}
			m.RecordNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecordExporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RecordNumber |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Record", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecordExporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRecordExporter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecordExporter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Record.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecordExporter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRecordExporter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRecordExporter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRecordExporter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRecordExporter
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRecordExporter
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRecordExporter
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRecordExporter
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRecordExporter
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRecordExporter
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRecordExporter(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRecordExporter
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRecordExporter = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRecordExporter   = fmt.Errorf("proto: integer overflow")
)