// Code generated by protoc-gen-go. DO NOT EDIT.
// source: multicast.proto

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

type MulticastProto int32

const (
	MulticastProto_igmp          MulticastProto = 0
	MulticastProto_pim           MulticastProto = 1
	MulticastProto_anycast_rp    MulticastProto = 2
	MulticastProto_mld           MulticastProto = 3
	MulticastProto_igmp_snooping MulticastProto = 4
	MulticastProto_mld_snooping  MulticastProto = 5
)

var MulticastProto_name = map[int32]string{
	0: "igmp",
	1: "pim",
	2: "anycast_rp",
	3: "mld",
	4: "igmp_snooping",
	5: "mld_snooping",
}

var MulticastProto_value = map[string]int32{
	"igmp":          0,
	"pim":           1,
	"anycast_rp":    2,
	"mld":           3,
	"igmp_snooping": 4,
	"mld_snooping":  5,
}

func (x MulticastProto) String() string {
	return proto.EnumName(MulticastProto_name, int32(x))
}

func (MulticastProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eedbde62517e047e, []int{0}
}

type MCCommonReq struct {
	Ifname               string   `protobuf:"bytes,1,opt,name=ifname,proto3" json:"ifname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MCCommonReq) Reset()         { *m = MCCommonReq{} }
func (m *MCCommonReq) String() string { return proto.CompactTextString(m) }
func (*MCCommonReq) ProtoMessage()    {}
func (*MCCommonReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_eedbde62517e047e, []int{0}
}

func (m *MCCommonReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MCCommonReq.Unmarshal(m, b)
}
func (m *MCCommonReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MCCommonReq.Marshal(b, m, deterministic)
}
func (m *MCCommonReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MCCommonReq.Merge(m, src)
}
func (m *MCCommonReq) XXX_Size() int {
	return xxx_messageInfo_MCCommonReq.Size(m)
}
func (m *MCCommonReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MCCommonReq.DiscardUnknown(m)
}

var xxx_messageInfo_MCCommonReq proto.InternalMessageInfo

func (m *MCCommonReq) GetIfname() string {
	if m != nil {
		return m.Ifname
	}
	return ""
}

//组播开关，在每个端口上打开或关闭指定的组播协议
type MulticastSwitch struct {
	Ifname               string         `protobuf:"bytes,1,opt,name=ifname,proto3" json:"ifname,omitempty"`
	Proto                MulticastProto `protobuf:"varint,2,opt,name=proto,proto3,enum=cml.MulticastProto" json:"proto,omitempty"`
	Enabled              bool           `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MulticastSwitch) Reset()         { *m = MulticastSwitch{} }
func (m *MulticastSwitch) String() string { return proto.CompactTextString(m) }
func (*MulticastSwitch) ProtoMessage()    {}
func (*MulticastSwitch) Descriptor() ([]byte, []int) {
	return fileDescriptor_eedbde62517e047e, []int{1}
}

func (m *MulticastSwitch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MulticastSwitch.Unmarshal(m, b)
}
func (m *MulticastSwitch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MulticastSwitch.Marshal(b, m, deterministic)
}
func (m *MulticastSwitch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MulticastSwitch.Merge(m, src)
}
func (m *MulticastSwitch) XXX_Size() int {
	return xxx_messageInfo_MulticastSwitch.Size(m)
}
func (m *MulticastSwitch) XXX_DiscardUnknown() {
	xxx_messageInfo_MulticastSwitch.DiscardUnknown(m)
}

var xxx_messageInfo_MulticastSwitch proto.InternalMessageInfo

func (m *MulticastSwitch) GetIfname() string {
	if m != nil {
		return m.Ifname
	}
	return ""
}

func (m *MulticastSwitch) GetProto() MulticastProto {
	if m != nil {
		return m.Proto
	}
	return MulticastProto_igmp
}

func (m *MulticastSwitch) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type MulticastSwitchList struct {
	McList               []*MulticastSwitch `protobuf:"bytes,1,rep,name=mcList,proto3" json:"mcList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MulticastSwitchList) Reset()         { *m = MulticastSwitchList{} }
func (m *MulticastSwitchList) String() string { return proto.CompactTextString(m) }
func (*MulticastSwitchList) ProtoMessage()    {}
func (*MulticastSwitchList) Descriptor() ([]byte, []int) {
	return fileDescriptor_eedbde62517e047e, []int{2}
}

func (m *MulticastSwitchList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MulticastSwitchList.Unmarshal(m, b)
}
func (m *MulticastSwitchList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MulticastSwitchList.Marshal(b, m, deterministic)
}
func (m *MulticastSwitchList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MulticastSwitchList.Merge(m, src)
}
func (m *MulticastSwitchList) XXX_Size() int {
	return xxx_messageInfo_MulticastSwitchList.Size(m)
}
func (m *MulticastSwitchList) XXX_DiscardUnknown() {
	xxx_messageInfo_MulticastSwitchList.DiscardUnknown(m)
}

var xxx_messageInfo_MulticastSwitchList proto.InternalMessageInfo

func (m *MulticastSwitchList) GetMcList() []*MulticastSwitch {
	if m != nil {
		return m.McList
	}
	return nil
}

type MCGroupMember struct {
	Ifname               string         `protobuf:"bytes,1,opt,name=ifname,proto3" json:"ifname,omitempty"`
	Proto                MulticastProto `protobuf:"varint,2,opt,name=proto,proto3,enum=cml.MulticastProto" json:"proto,omitempty"`
	MulticastAddr        string         `protobuf:"bytes,3,opt,name=multicastAddr,proto3" json:"multicastAddr,omitempty"`
	Members              []string       `protobuf:"bytes,4,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MCGroupMember) Reset()         { *m = MCGroupMember{} }
func (m *MCGroupMember) String() string { return proto.CompactTextString(m) }
func (*MCGroupMember) ProtoMessage()    {}
func (*MCGroupMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_eedbde62517e047e, []int{3}
}

func (m *MCGroupMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MCGroupMember.Unmarshal(m, b)
}
func (m *MCGroupMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MCGroupMember.Marshal(b, m, deterministic)
}
func (m *MCGroupMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MCGroupMember.Merge(m, src)
}
func (m *MCGroupMember) XXX_Size() int {
	return xxx_messageInfo_MCGroupMember.Size(m)
}
func (m *MCGroupMember) XXX_DiscardUnknown() {
	xxx_messageInfo_MCGroupMember.DiscardUnknown(m)
}

var xxx_messageInfo_MCGroupMember proto.InternalMessageInfo

func (m *MCGroupMember) GetIfname() string {
	if m != nil {
		return m.Ifname
	}
	return ""
}

func (m *MCGroupMember) GetProto() MulticastProto {
	if m != nil {
		return m.Proto
	}
	return MulticastProto_igmp
}

func (m *MCGroupMember) GetMulticastAddr() string {
	if m != nil {
		return m.MulticastAddr
	}
	return ""
}

func (m *MCGroupMember) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type MCGroupMemberReq struct {
	Proto                MulticastProto `protobuf:"varint,1,opt,name=proto,proto3,enum=cml.MulticastProto" json:"proto,omitempty"`
	MulticastAddr        string         `protobuf:"bytes,2,opt,name=multicastAddr,proto3" json:"multicastAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MCGroupMemberReq) Reset()         { *m = MCGroupMemberReq{} }
func (m *MCGroupMemberReq) String() string { return proto.CompactTextString(m) }
func (*MCGroupMemberReq) ProtoMessage()    {}
func (*MCGroupMemberReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_eedbde62517e047e, []int{4}
}

func (m *MCGroupMemberReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MCGroupMemberReq.Unmarshal(m, b)
}
func (m *MCGroupMemberReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MCGroupMemberReq.Marshal(b, m, deterministic)
}
func (m *MCGroupMemberReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MCGroupMemberReq.Merge(m, src)
}
func (m *MCGroupMemberReq) XXX_Size() int {
	return xxx_messageInfo_MCGroupMemberReq.Size(m)
}
func (m *MCGroupMemberReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MCGroupMemberReq.DiscardUnknown(m)
}

var xxx_messageInfo_MCGroupMemberReq proto.InternalMessageInfo

func (m *MCGroupMemberReq) GetProto() MulticastProto {
	if m != nil {
		return m.Proto
	}
	return MulticastProto_igmp
}

func (m *MCGroupMemberReq) GetMulticastAddr() string {
	if m != nil {
		return m.MulticastAddr
	}
	return ""
}

func init() {
	proto.RegisterEnum("cml.MulticastProto", MulticastProto_name, MulticastProto_value)
	proto.RegisterType((*MCCommonReq)(nil), "cml.MCCommonReq")
	proto.RegisterType((*MulticastSwitch)(nil), "cml.MulticastSwitch")
	proto.RegisterType((*MulticastSwitchList)(nil), "cml.MulticastSwitchList")
	proto.RegisterType((*MCGroupMember)(nil), "cml.MCGroupMember")
	proto.RegisterType((*MCGroupMemberReq)(nil), "cml.MCGroupMemberReq")
}

func init() {
	proto.RegisterFile("multicast.proto", fileDescriptor_eedbde62517e047e)
}

var fileDescriptor_eedbde62517e047e = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xdd, 0xce, 0x93, 0x40,
	0x10, 0x65, 0x4b, 0x7f, 0xa7, 0xb6, 0x5d, 0x17, 0x35, 0xa4, 0x57, 0x84, 0x68, 0x82, 0xc6, 0xf4,
	0xa2, 0xbe, 0x80, 0xca, 0x85, 0x5e, 0x48, 0x62, 0xe8, 0x03, 0x34, 0x74, 0x59, 0xdb, 0x4d, 0x18,
	0x76, 0x05, 0x6a, 0xd3, 0xd7, 0xf0, 0x09, 0x7d, 0x14, 0xc3, 0x5f, 0x1b, 0x48, 0xbd, 0x30, 0xf9,
	0xee, 0x98, 0xb3, 0x67, 0xce, 0x19, 0x66, 0x0e, 0xac, 0xf0, 0x9c, 0x14, 0x92, 0x47, 0x79, 0xb1,
	0xd1, 0x99, 0x2a, 0x14, 0x33, 0x39, 0x26, 0x6b, 0xca, 0x15, 0xe2, 0xbe, 0xb8, 0x6a, 0x91, 0xd7,
	0xb0, 0xfb, 0x06, 0xe6, 0x81, 0xef, 0x2b, 0x44, 0x95, 0x86, 0xe2, 0x27, 0x7b, 0x05, 0x63, 0xf9,
	0x23, 0x8d, 0x50, 0xd8, 0xc4, 0x21, 0xde, 0x2c, 0x6c, 0x2a, 0x37, 0x85, 0x55, 0xd0, 0x0a, 0xee,
	0x2e, 0xb2, 0xe0, 0xa7, 0x7f, 0x51, 0xd9, 0x5b, 0x18, 0x55, 0xd2, 0xf6, 0xc0, 0x21, 0xde, 0x72,
	0x6b, 0x6d, 0x38, 0x26, 0x9b, 0x5b, 0xf3, 0xf7, 0xf2, 0x29, 0xac, 0x19, 0xcc, 0x86, 0x89, 0x48,
	0xa3, 0x43, 0x22, 0x62, 0xdb, 0x74, 0x88, 0x37, 0x0d, 0xdb, 0xd2, 0xf5, 0xc1, 0xea, 0xf9, 0x7d,
	0x93, 0x79, 0xc1, 0xde, 0xc3, 0x18, 0x79, 0xf9, 0x65, 0x13, 0xc7, 0xf4, 0xe6, 0xdb, 0x17, 0x5d,
	0xf1, 0x9a, 0x19, 0x36, 0x1c, 0xf7, 0x37, 0x81, 0x45, 0xe0, 0x7f, 0xc9, 0xd4, 0x59, 0x07, 0x02,
	0x0f, 0x22, 0x7b, 0x8a, 0x99, 0x5f, 0xc3, 0xe2, 0xb6, 0xda, 0x4f, 0x71, 0x9c, 0x55, 0x93, 0xcf,
	0xc2, 0x2e, 0x58, 0xfe, 0x19, 0x56, 0x96, 0xb9, 0x3d, 0x74, 0x4c, 0x6f, 0x16, 0xb6, 0xa5, 0xcb,
	0x81, 0x76, 0x66, 0x2a, 0xb7, 0x7e, 0xb3, 0x27, 0xff, 0x6f, 0x3f, 0x78, 0x60, 0xff, 0x2e, 0x82,
	0x65, 0xb7, 0x9d, 0x4d, 0x61, 0x28, 0x8f, 0xa8, 0xa9, 0xc1, 0x26, 0x60, 0x6a, 0x89, 0x94, 0xb0,
	0x25, 0x40, 0x94, 0x5e, 0x4b, 0xca, 0x3e, 0xd3, 0x74, 0x50, 0x3e, 0x60, 0x12, 0x53, 0x93, 0x3d,
	0x87, 0x45, 0xc9, 0xdd, 0xe7, 0xa9, 0x52, 0x5a, 0xa6, 0x47, 0x3a, 0x64, 0x14, 0x9e, 0x61, 0x12,
	0xdf, 0x91, 0xd1, 0xf6, 0x0f, 0x01, 0xcb, 0xc7, 0xe4, 0xbe, 0x7b, 0x91, 0xfd, 0x92, 0x5c, 0xb0,
	0x8f, 0xc0, 0x76, 0xa2, 0xe8, 0x87, 0xe5, 0xe1, 0xa1, 0xd6, 0xac, 0x42, 0xdb, 0xf4, 0xe5, 0xfa,
	0x6b, 0x9c, 0xb9, 0x06, 0xf3, 0xc1, 0xda, 0x9d, 0xd4, 0xa5, 0x2f, 0x41, 0x6b, 0x89, 0x7b, 0x58,
	0xd7, 0xf6, 0x23, 0xd1, 0xea, 0xf2, 0x06, 0xfb, 0xdc, 0x13, 0x69, 0x02, 0xf0, 0xb2, 0x11, 0xe9,
	0x1e, 0xa0, 0x19, 0xa4, 0x03, 0xbb, 0xc6, 0x61, 0x5c, 0xad, 0xfc, 0xc3, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd3, 0xd8, 0xc4, 0xb0, 0x4c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CmlMulticastServiceClient is the client API for CmlMulticastService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CmlMulticastServiceClient interface {
	//打开/关闭指定端口的组播协议
	SetMulticastSwitch(ctx context.Context, in *MulticastSwitch, opts ...grpc.CallOption) (*CommonRespHdr, error)
	//Show指定端口的组播协议开关，只认入参的接口名称
	ShowMulticastSwitch(ctx context.Context, in *MCCommonReq, opts ...grpc.CallOption) (*MulticastSwitchList, error)
	//Show组播成员
	ShowMulticastMember(ctx context.Context, in *MCGroupMemberReq, opts ...grpc.CallOption) (*MCGroupMember, error)
}

type cmlMulticastServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCmlMulticastServiceClient(cc grpc.ClientConnInterface) CmlMulticastServiceClient {
	return &cmlMulticastServiceClient{cc}
}

func (c *cmlMulticastServiceClient) SetMulticastSwitch(ctx context.Context, in *MulticastSwitch, opts ...grpc.CallOption) (*CommonRespHdr, error) {
	out := new(CommonRespHdr)
	err := c.cc.Invoke(ctx, "/cml.CmlMulticastService/SetMulticastSwitch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmlMulticastServiceClient) ShowMulticastSwitch(ctx context.Context, in *MCCommonReq, opts ...grpc.CallOption) (*MulticastSwitchList, error) {
	out := new(MulticastSwitchList)
	err := c.cc.Invoke(ctx, "/cml.CmlMulticastService/ShowMulticastSwitch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmlMulticastServiceClient) ShowMulticastMember(ctx context.Context, in *MCGroupMemberReq, opts ...grpc.CallOption) (*MCGroupMember, error) {
	out := new(MCGroupMember)
	err := c.cc.Invoke(ctx, "/cml.CmlMulticastService/ShowMulticastMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CmlMulticastServiceServer is the server API for CmlMulticastService service.
type CmlMulticastServiceServer interface {
	//打开/关闭指定端口的组播协议
	SetMulticastSwitch(context.Context, *MulticastSwitch) (*CommonRespHdr, error)
	//Show指定端口的组播协议开关，只认入参的接口名称
	ShowMulticastSwitch(context.Context, *MCCommonReq) (*MulticastSwitchList, error)
	//Show组播成员
	ShowMulticastMember(context.Context, *MCGroupMemberReq) (*MCGroupMember, error)
}

// UnimplementedCmlMulticastServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCmlMulticastServiceServer struct {
}

func (*UnimplementedCmlMulticastServiceServer) SetMulticastSwitch(ctx context.Context, req *MulticastSwitch) (*CommonRespHdr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMulticastSwitch not implemented")
}
func (*UnimplementedCmlMulticastServiceServer) ShowMulticastSwitch(ctx context.Context, req *MCCommonReq) (*MulticastSwitchList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowMulticastSwitch not implemented")
}
func (*UnimplementedCmlMulticastServiceServer) ShowMulticastMember(ctx context.Context, req *MCGroupMemberReq) (*MCGroupMember, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowMulticastMember not implemented")
}

func RegisterCmlMulticastServiceServer(s *grpc.Server, srv CmlMulticastServiceServer) {
	s.RegisterService(&_CmlMulticastService_serviceDesc, srv)
}

func _CmlMulticastService_SetMulticastSwitch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MulticastSwitch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmlMulticastServiceServer).SetMulticastSwitch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.CmlMulticastService/SetMulticastSwitch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmlMulticastServiceServer).SetMulticastSwitch(ctx, req.(*MulticastSwitch))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmlMulticastService_ShowMulticastSwitch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MCCommonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmlMulticastServiceServer).ShowMulticastSwitch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.CmlMulticastService/ShowMulticastSwitch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmlMulticastServiceServer).ShowMulticastSwitch(ctx, req.(*MCCommonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmlMulticastService_ShowMulticastMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MCGroupMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmlMulticastServiceServer).ShowMulticastMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.CmlMulticastService/ShowMulticastMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmlMulticastServiceServer).ShowMulticastMember(ctx, req.(*MCGroupMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CmlMulticastService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cml.CmlMulticastService",
	HandlerType: (*CmlMulticastServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetMulticastSwitch",
			Handler:    _CmlMulticastService_SetMulticastSwitch_Handler,
		},
		{
			MethodName: "ShowMulticastSwitch",
			Handler:    _CmlMulticastService_ShowMulticastSwitch_Handler,
		},
		{
			MethodName: "ShowMulticastMember",
			Handler:    _CmlMulticastService_ShowMulticastMember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "multicast.proto",
}
