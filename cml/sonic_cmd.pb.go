// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sonic_cmd.proto

package cml

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//命令执行返回码
//对执行成功的配置命令，默认不输出信息
//对交互式命令，第一级交互在klish完成，执行成功的也不输出信息
//执行失败的命令都必须输出信息
//对ping,show等存在多次输出的,中间存在交互方式(用户可终止执行)
type CMD_RTNCODE int32

const (
	CMD_RTNCODE_CMD_EXEC_SUCC       CMD_RTNCODE = 0
	CMD_RTNCODE_CMD_EXEC_FAILED     CMD_RTNCODE = 1
	CMD_RTNCODE_CMD_EXEC_FINISHED   CMD_RTNCODE = 2
	CMD_RTNCODE_CMD_EXEC_NOT_FINISH CMD_RTNCODE = 3
)

var CMD_RTNCODE_name = map[int32]string{
	0: "CMD_EXEC_SUCC",
	1: "CMD_EXEC_FAILED",
	2: "CMD_EXEC_FINISHED",
	3: "CMD_EXEC_NOT_FINISH",
}

var CMD_RTNCODE_value = map[string]int32{
	"CMD_EXEC_SUCC":       0,
	"CMD_EXEC_FAILED":     1,
	"CMD_EXEC_FINISHED":   2,
	"CMD_EXEC_NOT_FINISH": 3,
}

func (x CMD_RTNCODE) String() string {
	return proto.EnumName(CMD_RTNCODE_name, int32(x))
}

func (CMD_RTNCODE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aada0c871e733071, []int{0}
}

//SONIC命令执行消息（配置和show等共用一套）
type SONICCmdInputProfile struct {
	CmdId                uint32   `protobuf:"varint,1,opt,name=cmd_id,json=cmdId,proto3" json:"cmd_id,omitempty"`
	CmdinputLine         string   `protobuf:"bytes,2,opt,name=cmdinput_line,json=cmdinputLine,proto3" json:"cmdinput_line,omitempty"`
	CmdviewList          string   `protobuf:"bytes,3,opt,name=cmdview_list,json=cmdviewList,proto3" json:"cmdview_list,omitempty"`
	UpdatedAt            uint64   `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	KlishVtyNo           uint32   `protobuf:"varint,5,opt,name=klish_vty_no,json=klishVtyNo,proto3" json:"klish_vty_no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SONICCmdInputProfile) Reset()         { *m = SONICCmdInputProfile{} }
func (m *SONICCmdInputProfile) String() string { return proto.CompactTextString(m) }
func (*SONICCmdInputProfile) ProtoMessage()    {}
func (*SONICCmdInputProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_aada0c871e733071, []int{0}
}

func (m *SONICCmdInputProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SONICCmdInputProfile.Unmarshal(m, b)
}
func (m *SONICCmdInputProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SONICCmdInputProfile.Marshal(b, m, deterministic)
}
func (m *SONICCmdInputProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SONICCmdInputProfile.Merge(m, src)
}
func (m *SONICCmdInputProfile) XXX_Size() int {
	return xxx_messageInfo_SONICCmdInputProfile.Size(m)
}
func (m *SONICCmdInputProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_SONICCmdInputProfile.DiscardUnknown(m)
}

var xxx_messageInfo_SONICCmdInputProfile proto.InternalMessageInfo

func (m *SONICCmdInputProfile) GetCmdId() uint32 {
	if m != nil {
		return m.CmdId
	}
	return 0
}

func (m *SONICCmdInputProfile) GetCmdinputLine() string {
	if m != nil {
		return m.CmdinputLine
	}
	return ""
}

func (m *SONICCmdInputProfile) GetCmdviewList() string {
	if m != nil {
		return m.CmdviewList
	}
	return ""
}

func (m *SONICCmdInputProfile) GetUpdatedAt() uint64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *SONICCmdInputProfile) GetKlishVtyNo() uint32 {
	if m != nil {
		return m.KlishVtyNo
	}
	return 0
}

//SONIC 原生非配置命令如show或诊断等 返回定义。部分show存在多个返回信息（rpc采用流方式）
type SONICCmdRtnProfile struct {
	CmdId                uint32      `protobuf:"varint,1,opt,name=cmd_id,json=cmdId,proto3" json:"cmd_id,omitempty"`
	CmdRtncode           CMD_RTNCODE `protobuf:"varint,2,opt,name=cmd_rtncode,json=cmdRtncode,proto3,enum=cml.CMD_RTNCODE" json:"cmd_rtncode,omitempty"`
	CmdrtnStr            string      `protobuf:"bytes,3,opt,name=cmdrtn_str,json=cmdrtnStr,proto3" json:"cmdrtn_str,omitempty"`
	CmdrtnSerial         uint32      `protobuf:"varint,4,opt,name=cmdrtn_serial,json=cmdrtnSerial,proto3" json:"cmdrtn_serial,omitempty"`
	UpdatedAt            uint64      `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	KlishVtyNo           uint32      `protobuf:"varint,6,opt,name=klish_vty_no,json=klishVtyNo,proto3" json:"klish_vty_no,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SONICCmdRtnProfile) Reset()         { *m = SONICCmdRtnProfile{} }
func (m *SONICCmdRtnProfile) String() string { return proto.CompactTextString(m) }
func (*SONICCmdRtnProfile) ProtoMessage()    {}
func (*SONICCmdRtnProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_aada0c871e733071, []int{1}
}

func (m *SONICCmdRtnProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SONICCmdRtnProfile.Unmarshal(m, b)
}
func (m *SONICCmdRtnProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SONICCmdRtnProfile.Marshal(b, m, deterministic)
}
func (m *SONICCmdRtnProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SONICCmdRtnProfile.Merge(m, src)
}
func (m *SONICCmdRtnProfile) XXX_Size() int {
	return xxx_messageInfo_SONICCmdRtnProfile.Size(m)
}
func (m *SONICCmdRtnProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_SONICCmdRtnProfile.DiscardUnknown(m)
}

var xxx_messageInfo_SONICCmdRtnProfile proto.InternalMessageInfo

func (m *SONICCmdRtnProfile) GetCmdId() uint32 {
	if m != nil {
		return m.CmdId
	}
	return 0
}

func (m *SONICCmdRtnProfile) GetCmdRtncode() CMD_RTNCODE {
	if m != nil {
		return m.CmdRtncode
	}
	return CMD_RTNCODE_CMD_EXEC_SUCC
}

func (m *SONICCmdRtnProfile) GetCmdrtnStr() string {
	if m != nil {
		return m.CmdrtnStr
	}
	return ""
}

func (m *SONICCmdRtnProfile) GetCmdrtnSerial() uint32 {
	if m != nil {
		return m.CmdrtnSerial
	}
	return 0
}

func (m *SONICCmdRtnProfile) GetUpdatedAt() uint64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *SONICCmdRtnProfile) GetKlishVtyNo() uint32 {
	if m != nil {
		return m.KlishVtyNo
	}
	return 0
}

func init() {
	proto.RegisterEnum("cml.CMD_RTNCODE", CMD_RTNCODE_name, CMD_RTNCODE_value)
	proto.RegisterType((*SONICCmdInputProfile)(nil), "cml.SONICCmdInputProfile")
	proto.RegisterType((*SONICCmdRtnProfile)(nil), "cml.SONICCmdRtnProfile")
}

func init() {
	proto.RegisterFile("sonic_cmd.proto", fileDescriptor_aada0c871e733071)
}

var fileDescriptor_aada0c871e733071 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xd1, 0x8e, 0xd2, 0x40,
	0x14, 0x86, 0xb7, 0xcb, 0x42, 0xc2, 0xd9, 0xad, 0x74, 0x07, 0x37, 0x5b, 0x4d, 0x4c, 0x2a, 0xde,
	0x10, 0x2f, 0x48, 0xc4, 0x27, 0x20, 0xd3, 0x1a, 0x1b, 0xb1, 0x98, 0x16, 0x8d, 0x57, 0x4e, 0x70,
	0x66, 0xd0, 0x89, 0x9d, 0x29, 0x99, 0x0e, 0x28, 0x2f, 0xe6, 0x0b, 0xf9, 0x22, 0x66, 0x86, 0x02,
	0x4a, 0x36, 0xe1, 0xf2, 0x7c, 0xe7, 0x4f, 0xfb, 0xff, 0xff, 0x1c, 0xe8, 0xd5, 0x95, 0x12, 0x94,
	0x50, 0xc9, 0x46, 0x2b, 0x5d, 0x99, 0x0a, 0xb5, 0xa8, 0x2c, 0x07, 0xbf, 0x3d, 0x78, 0x5c, 0xcc,
	0xb2, 0x14, 0x63, 0xc9, 0x52, 0xb5, 0x5a, 0x9b, 0x0f, 0xba, 0x5a, 0x8a, 0x92, 0xa3, 0x3b, 0xe8,
	0x50, 0xc9, 0x88, 0x60, 0xa1, 0x17, 0x79, 0x43, 0x3f, 0x6f, 0x53, 0xc9, 0x52, 0x86, 0x5e, 0x80,
	0x4f, 0x25, 0x13, 0x56, 0x49, 0x4a, 0xa1, 0x78, 0x78, 0x19, 0x79, 0xc3, 0x6e, 0x7e, 0xb3, 0x87,
	0x53, 0xa1, 0x38, 0x7a, 0x0e, 0x76, 0xde, 0x08, 0xfe, 0x93, 0x94, 0xa2, 0x36, 0x61, 0xcb, 0x69,
	0xae, 0x1b, 0x36, 0x15, 0xb5, 0x41, 0xcf, 0x00, 0xd6, 0x2b, 0xb6, 0x30, 0x9c, 0x91, 0x85, 0x09,
	0xaf, 0x22, 0x6f, 0x78, 0x95, 0x77, 0x1b, 0x32, 0x31, 0x28, 0x82, 0x9b, 0x1f, 0xa5, 0xa8, 0xbf,
	0x93, 0x8d, 0xd9, 0x12, 0x55, 0x85, 0x6d, 0xe7, 0x01, 0x1c, 0xfb, 0x64, 0xb6, 0x59, 0x35, 0xf8,
	0xe3, 0x01, 0xda, 0x1b, 0xcf, 0x8d, 0x3a, 0x63, 0xfb, 0x15, 0xd8, 0xbf, 0x13, 0x6d, 0x14, 0xad,
	0xd8, 0xce, 0xf4, 0xa3, 0x71, 0x30, 0xa2, 0xb2, 0x1c, 0xe1, 0xf7, 0x31, 0xc9, 0xe7, 0x19, 0x9e,
	0xc5, 0x49, 0x0e, 0xd4, 0x7d, 0xcc, 0x6a, 0xac, 0x43, 0x2a, 0x99, 0x36, 0x8a, 0xd4, 0x46, 0x37,
	0x11, 0xba, 0x3b, 0x52, 0x18, 0xdd, 0x14, 0xe1, 0xd6, 0x5c, 0x8b, 0x45, 0xe9, 0x32, 0xf8, 0xae,
	0x08, 0xab, 0x70, 0xec, 0x24, 0x65, 0xfb, 0x5c, 0xca, 0xce, 0x69, 0xca, 0x97, 0x1c, 0xae, 0xff,
	0xf1, 0x87, 0x6e, 0xc1, 0xb7, 0x63, 0xf2, 0x39, 0xc1, 0xa4, 0xf8, 0x88, 0x71, 0x70, 0x81, 0xfa,
	0xd0, 0x3b, 0xa0, 0x37, 0x93, 0x74, 0x9a, 0xc4, 0x81, 0x87, 0xee, 0xe0, 0xf6, 0x08, 0xd3, 0x2c,
	0x2d, 0xde, 0x26, 0x71, 0x70, 0x89, 0xee, 0xa1, 0x7f, 0xc0, 0xd9, 0x6c, 0xde, 0xac, 0x82, 0xd6,
	0xf8, 0x0b, 0xf4, 0x0a, 0x7b, 0x1d, 0x58, 0xb2, 0x82, 0xeb, 0x8d, 0xa0, 0x1c, 0xbd, 0x83, 0x7e,
	0xf2, 0x8b, 0xd3, 0x1d, 0x5e, 0x7e, 0xdb, 0xf7, 0xfb, 0xc4, 0x75, 0xf6, 0xd0, 0xc5, 0x3c, 0xbd,
	0xff, 0x6f, 0x75, 0x7c, 0x93, 0xc1, 0xc5, 0xd7, 0x8e, 0xbb, 0xb8, 0xd7, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x39, 0xe9, 0x2f, 0x8f, 0x84, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SonicCmdServiceClient is the client API for SonicCmdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SonicCmdServiceClient interface {
	//sonic原生命令处理流程,配置和show等共用一套。由klish通过grpc到cml，由cml分发
	ExecSonicCfgProfile(ctx context.Context, in *SONICCmdInputProfile, opts ...grpc.CallOption) (*SONICCmdRtnProfile, error)
}

type sonicCmdServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSonicCmdServiceClient(cc grpc.ClientConnInterface) SonicCmdServiceClient {
	return &sonicCmdServiceClient{cc}
}

func (c *sonicCmdServiceClient) ExecSonicCfgProfile(ctx context.Context, in *SONICCmdInputProfile, opts ...grpc.CallOption) (*SONICCmdRtnProfile, error) {
	out := new(SONICCmdRtnProfile)
	err := c.cc.Invoke(ctx, "/cml.SonicCmdService/ExecSonicCfgProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SonicCmdServiceServer is the server API for SonicCmdService service.
type SonicCmdServiceServer interface {
	//sonic原生命令处理流程,配置和show等共用一套。由klish通过grpc到cml，由cml分发
	ExecSonicCfgProfile(context.Context, *SONICCmdInputProfile) (*SONICCmdRtnProfile, error)
}

// UnimplementedSonicCmdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSonicCmdServiceServer struct {
}

func (*UnimplementedSonicCmdServiceServer) ExecSonicCfgProfile(ctx context.Context, req *SONICCmdInputProfile) (*SONICCmdRtnProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecSonicCfgProfile not implemented")
}

func RegisterSonicCmdServiceServer(s *grpc.Server, srv SonicCmdServiceServer) {
	s.RegisterService(&_SonicCmdService_serviceDesc, srv)
}

func _SonicCmdService_ExecSonicCfgProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SONICCmdInputProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SonicCmdServiceServer).ExecSonicCfgProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.SonicCmdService/ExecSonicCfgProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SonicCmdServiceServer).ExecSonicCfgProfile(ctx, req.(*SONICCmdInputProfile))
	}
	return interceptor(ctx, in, info, handler)
}

var _SonicCmdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cml.SonicCmdService",
	HandlerType: (*SonicCmdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecSonicCfgProfile",
			Handler:    _SonicCmdService_ExecSonicCfgProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sonic_cmd.proto",
}
