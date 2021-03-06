// Code generated by protoc-gen-go. DO NOT EDIT.
// source: loopdetect.proto

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

//loopback detection接口下配置
type Lpdetecifcfg struct {
	Ifname               string   `protobuf:"bytes,1,opt,name=ifname,proto3" json:"ifname,omitempty"`
	Action               string   `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Lpdetecenable        int32    `protobuf:"varint,3,opt,name=lpdetecenable,proto3" json:"lpdetecenable,omitempty"`
	Vlanlist             string   `protobuf:"bytes,4,opt,name=vlanlist,proto3" json:"vlanlist,omitempty"`
	Updatetime           int64    `protobuf:"varint,5,opt,name=updatetime,proto3" json:"updatetime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lpdetecifcfg) Reset()         { *m = Lpdetecifcfg{} }
func (m *Lpdetecifcfg) String() string { return proto.CompactTextString(m) }
func (*Lpdetecifcfg) ProtoMessage()    {}
func (*Lpdetecifcfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a8f3bd64231ab91, []int{0}
}

func (m *Lpdetecifcfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lpdetecifcfg.Unmarshal(m, b)
}
func (m *Lpdetecifcfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lpdetecifcfg.Marshal(b, m, deterministic)
}
func (m *Lpdetecifcfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lpdetecifcfg.Merge(m, src)
}
func (m *Lpdetecifcfg) XXX_Size() int {
	return xxx_messageInfo_Lpdetecifcfg.Size(m)
}
func (m *Lpdetecifcfg) XXX_DiscardUnknown() {
	xxx_messageInfo_Lpdetecifcfg.DiscardUnknown(m)
}

var xxx_messageInfo_Lpdetecifcfg proto.InternalMessageInfo

func (m *Lpdetecifcfg) GetIfname() string {
	if m != nil {
		return m.Ifname
	}
	return ""
}

func (m *Lpdetecifcfg) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Lpdetecifcfg) GetLpdetecenable() int32 {
	if m != nil {
		return m.Lpdetecenable
	}
	return 0
}

func (m *Lpdetecifcfg) GetVlanlist() string {
	if m != nil {
		return m.Vlanlist
	}
	return ""
}

func (m *Lpdetecifcfg) GetUpdatetime() int64 {
	if m != nil {
		return m.Updatetime
	}
	return 0
}

//loopbakc global配置
type Lpdetecglobalcfg struct {
	Errordown            int32    `protobuf:"varint,1,opt,name=errordown,proto3" json:"errordown,omitempty"`
	Transinterval        int32    `protobuf:"varint,2,opt,name=transinterval,proto3" json:"transinterval,omitempty"`
	Updatetime           int64    `protobuf:"varint,3,opt,name=updatetime,proto3" json:"updatetime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lpdetecglobalcfg) Reset()         { *m = Lpdetecglobalcfg{} }
func (m *Lpdetecglobalcfg) String() string { return proto.CompactTextString(m) }
func (*Lpdetecglobalcfg) ProtoMessage()    {}
func (*Lpdetecglobalcfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a8f3bd64231ab91, []int{1}
}

func (m *Lpdetecglobalcfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lpdetecglobalcfg.Unmarshal(m, b)
}
func (m *Lpdetecglobalcfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lpdetecglobalcfg.Marshal(b, m, deterministic)
}
func (m *Lpdetecglobalcfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lpdetecglobalcfg.Merge(m, src)
}
func (m *Lpdetecglobalcfg) XXX_Size() int {
	return xxx_messageInfo_Lpdetecglobalcfg.Size(m)
}
func (m *Lpdetecglobalcfg) XXX_DiscardUnknown() {
	xxx_messageInfo_Lpdetecglobalcfg.DiscardUnknown(m)
}

var xxx_messageInfo_Lpdetecglobalcfg proto.InternalMessageInfo

func (m *Lpdetecglobalcfg) GetErrordown() int32 {
	if m != nil {
		return m.Errordown
	}
	return 0
}

func (m *Lpdetecglobalcfg) GetTransinterval() int32 {
	if m != nil {
		return m.Transinterval
	}
	return 0
}

func (m *Lpdetecglobalcfg) GetUpdatetime() int64 {
	if m != nil {
		return m.Updatetime
	}
	return 0
}

func init() {
	proto.RegisterType((*Lpdetecifcfg)(nil), "cml.lpdetecifcfg")
	proto.RegisterType((*Lpdetecglobalcfg)(nil), "cml.lpdetecglobalcfg")
}

func init() {
	proto.RegisterFile("loopdetect.proto", fileDescriptor_0a8f3bd64231ab91)
}

var fileDescriptor_0a8f3bd64231ab91 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0xff, 0xfc, 0x31, 0xc5, 0x5e, 0x94, 0xa6, 0x43, 0x2b, 0x21, 0x88, 0x94, 0xe2, 0xa2,
	0xab, 0x2e, 0x74, 0xa3, 0x0b, 0xa1, 0x50, 0x41, 0x04, 0x57, 0xe9, 0x13, 0x4c, 0xa7, 0x33, 0xe9,
	0xc0, 0xcc, 0xdc, 0x32, 0x19, 0xd3, 0x27, 0xf0, 0x3d, 0x7c, 0x54, 0xc9, 0x24, 0x34, 0x69, 0x71,
	0x21, 0x5d, 0xde, 0x8f, 0x7b, 0xee, 0x39, 0x73, 0x18, 0x88, 0x15, 0xe2, 0x6e, 0xc3, 0x1d, 0x67,
	0x6e, 0xbe, 0xb3, 0xe8, 0x90, 0x84, 0x4c, 0xab, 0x74, 0xc0, 0x44, 0xce, 0x50, 0x6b, 0x34, 0x35,
	0x9d, 0x7e, 0x07, 0x70, 0xa5, 0xea, 0x45, 0x29, 0x98, 0xc8, 0xc9, 0x0d, 0xf4, 0xa4, 0x30, 0x54,
	0xf3, 0x24, 0x98, 0x04, 0xb3, 0x7e, 0xd6, 0x4c, 0x15, 0xa7, 0xcc, 0x49, 0x34, 0xc9, 0xff, 0x9a,
	0xd7, 0x13, 0xb9, 0x87, 0xeb, 0x46, 0xcf, 0x0d, 0x5d, 0x2b, 0x9e, 0x84, 0x93, 0x60, 0x16, 0x65,
	0xc7, 0x90, 0xa4, 0x70, 0x59, 0x2a, 0x6a, 0x94, 0x2c, 0x5c, 0x72, 0xe1, 0xf5, 0x87, 0x99, 0xdc,
	0x01, 0x7c, 0xee, 0x36, 0xd4, 0x71, 0x27, 0x35, 0x4f, 0xa2, 0x49, 0x30, 0x0b, 0xb3, 0x0e, 0x99,
	0x96, 0x10, 0x37, 0xc7, 0x72, 0x85, 0x6b, 0xaa, 0xaa, 0x94, 0xb7, 0xd0, 0xe7, 0xd6, 0xa2, 0xdd,
	0xe0, 0xde, 0xf8, 0xa0, 0x51, 0xd6, 0x82, 0x2a, 0x93, 0xb3, 0xd4, 0x14, 0xd2, 0x38, 0x6e, 0x4b,
	0xaa, 0x7c, 0xe4, 0x28, 0x3b, 0x86, 0x27, 0xbe, 0xe1, 0xa9, 0xef, 0xc3, 0x57, 0x08, 0xc3, 0xa5,
	0x56, 0x8d, 0x77, 0xc1, 0x6d, 0x29, 0x19, 0x27, 0x4f, 0x10, 0xaf, 0xb8, 0xfb, 0x68, 0xba, 0x7d,
	0x17, 0x4b, 0x91, 0x93, 0xe1, 0x9c, 0x69, 0x35, 0xef, 0xd6, 0x98, 0x0e, 0x3c, 0x62, 0x22, 0xb7,
	0xce, 0x48, 0x23, 0x70, 0xfa, 0xaf, 0x52, 0xbe, 0x72, 0x75, 0x8e, 0xf2, 0x19, 0x86, 0xab, 0x2d,
	0xee, 0x8f, 0xa5, 0xb1, 0xdf, 0x2b, 0xb6, 0xb8, 0x67, 0x22, 0xaf, 0x16, 0xd3, 0x96, 0xb4, 0xd2,
	0x05, 0x8c, 0x3a, 0x71, 0xdf, 0x7c, 0x81, 0x95, 0x7a, 0xdc, 0x35, 0x3e, 0xf4, 0xfa, 0x9b, 0xf9,
	0x02, 0x46, 0x9d, 0xd8, 0xe7, 0x5c, 0x78, 0x81, 0x71, 0x37, 0x7e, 0x7b, 0xe2, 0x4f, 0x4f, 0x58,
	0xf7, 0xfc, 0x4f, 0x7d, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x03, 0xb4, 0xa4, 0x2c, 0xd3, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CmllpdetecserviceClient is the client API for Cmllpdetecservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CmllpdetecserviceClient interface {
	//lpdetect在接口上的配置
	SetLpdetectIfCfg(ctx context.Context, in *Lpdetecifcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error)
	DelLpdetectIfCfg(ctx context.Context, in *Lpdetecifcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error)
	ShowLpdetectIfCfg(ctx context.Context, in *Showcfginfo, opts ...grpc.CallOption) (*Showrtninfo, error)
	//lpdetect全局配置
	SetLpdetectGlobalCfg(ctx context.Context, in *Lpdetecglobalcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error)
	DelLpdetectGlobalCfg(ctx context.Context, in *Lpdetecglobalcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error)
	ShowLpdetectGlobalCfg(ctx context.Context, in *Showcfginfo, opts ...grpc.CallOption) (*Showrtninfo, error)
}

type cmllpdetecserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewCmllpdetecserviceClient(cc grpc.ClientConnInterface) CmllpdetecserviceClient {
	return &cmllpdetecserviceClient{cc}
}

func (c *cmllpdetecserviceClient) SetLpdetectIfCfg(ctx context.Context, in *Lpdetecifcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error) {
	out := new(Cfgrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmllpdetecservice/SetLpdetectIfCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmllpdetecserviceClient) DelLpdetectIfCfg(ctx context.Context, in *Lpdetecifcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error) {
	out := new(Cfgrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmllpdetecservice/DelLpdetectIfCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmllpdetecserviceClient) ShowLpdetectIfCfg(ctx context.Context, in *Showcfginfo, opts ...grpc.CallOption) (*Showrtninfo, error) {
	out := new(Showrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmllpdetecservice/ShowLpdetectIfCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmllpdetecserviceClient) SetLpdetectGlobalCfg(ctx context.Context, in *Lpdetecglobalcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error) {
	out := new(Cfgrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmllpdetecservice/SetLpdetectGlobalCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmllpdetecserviceClient) DelLpdetectGlobalCfg(ctx context.Context, in *Lpdetecglobalcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error) {
	out := new(Cfgrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmllpdetecservice/DelLpdetectGlobalCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmllpdetecserviceClient) ShowLpdetectGlobalCfg(ctx context.Context, in *Showcfginfo, opts ...grpc.CallOption) (*Showrtninfo, error) {
	out := new(Showrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmllpdetecservice/ShowLpdetectGlobalCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CmllpdetecserviceServer is the server API for Cmllpdetecservice service.
type CmllpdetecserviceServer interface {
	//lpdetect在接口上的配置
	SetLpdetectIfCfg(context.Context, *Lpdetecifcfg) (*Cfgrtninfo, error)
	DelLpdetectIfCfg(context.Context, *Lpdetecifcfg) (*Cfgrtninfo, error)
	ShowLpdetectIfCfg(context.Context, *Showcfginfo) (*Showrtninfo, error)
	//lpdetect全局配置
	SetLpdetectGlobalCfg(context.Context, *Lpdetecglobalcfg) (*Cfgrtninfo, error)
	DelLpdetectGlobalCfg(context.Context, *Lpdetecglobalcfg) (*Cfgrtninfo, error)
	ShowLpdetectGlobalCfg(context.Context, *Showcfginfo) (*Showrtninfo, error)
}

// UnimplementedCmllpdetecserviceServer can be embedded to have forward compatible implementations.
type UnimplementedCmllpdetecserviceServer struct {
}

func (*UnimplementedCmllpdetecserviceServer) SetLpdetectIfCfg(ctx context.Context, req *Lpdetecifcfg) (*Cfgrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLpdetectIfCfg not implemented")
}
func (*UnimplementedCmllpdetecserviceServer) DelLpdetectIfCfg(ctx context.Context, req *Lpdetecifcfg) (*Cfgrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelLpdetectIfCfg not implemented")
}
func (*UnimplementedCmllpdetecserviceServer) ShowLpdetectIfCfg(ctx context.Context, req *Showcfginfo) (*Showrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowLpdetectIfCfg not implemented")
}
func (*UnimplementedCmllpdetecserviceServer) SetLpdetectGlobalCfg(ctx context.Context, req *Lpdetecglobalcfg) (*Cfgrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLpdetectGlobalCfg not implemented")
}
func (*UnimplementedCmllpdetecserviceServer) DelLpdetectGlobalCfg(ctx context.Context, req *Lpdetecglobalcfg) (*Cfgrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelLpdetectGlobalCfg not implemented")
}
func (*UnimplementedCmllpdetecserviceServer) ShowLpdetectGlobalCfg(ctx context.Context, req *Showcfginfo) (*Showrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowLpdetectGlobalCfg not implemented")
}

func RegisterCmllpdetecserviceServer(s *grpc.Server, srv CmllpdetecserviceServer) {
	s.RegisterService(&_Cmllpdetecservice_serviceDesc, srv)
}

func _Cmllpdetecservice_SetLpdetectIfCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Lpdetecifcfg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmllpdetecserviceServer).SetLpdetectIfCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmllpdetecservice/SetLpdetectIfCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmllpdetecserviceServer).SetLpdetectIfCfg(ctx, req.(*Lpdetecifcfg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmllpdetecservice_DelLpdetectIfCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Lpdetecifcfg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmllpdetecserviceServer).DelLpdetectIfCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmllpdetecservice/DelLpdetectIfCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmllpdetecserviceServer).DelLpdetectIfCfg(ctx, req.(*Lpdetecifcfg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmllpdetecservice_ShowLpdetectIfCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Showcfginfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmllpdetecserviceServer).ShowLpdetectIfCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmllpdetecservice/ShowLpdetectIfCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmllpdetecserviceServer).ShowLpdetectIfCfg(ctx, req.(*Showcfginfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmllpdetecservice_SetLpdetectGlobalCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Lpdetecglobalcfg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmllpdetecserviceServer).SetLpdetectGlobalCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmllpdetecservice/SetLpdetectGlobalCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmllpdetecserviceServer).SetLpdetectGlobalCfg(ctx, req.(*Lpdetecglobalcfg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmllpdetecservice_DelLpdetectGlobalCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Lpdetecglobalcfg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmllpdetecserviceServer).DelLpdetectGlobalCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmllpdetecservice/DelLpdetectGlobalCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmllpdetecserviceServer).DelLpdetectGlobalCfg(ctx, req.(*Lpdetecglobalcfg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmllpdetecservice_ShowLpdetectGlobalCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Showcfginfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmllpdetecserviceServer).ShowLpdetectGlobalCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmllpdetecservice/ShowLpdetectGlobalCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmllpdetecserviceServer).ShowLpdetectGlobalCfg(ctx, req.(*Showcfginfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cmllpdetecservice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cml.Cmllpdetecservice",
	HandlerType: (*CmllpdetecserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetLpdetectIfCfg",
			Handler:    _Cmllpdetecservice_SetLpdetectIfCfg_Handler,
		},
		{
			MethodName: "DelLpdetectIfCfg",
			Handler:    _Cmllpdetecservice_DelLpdetectIfCfg_Handler,
		},
		{
			MethodName: "ShowLpdetectIfCfg",
			Handler:    _Cmllpdetecservice_ShowLpdetectIfCfg_Handler,
		},
		{
			MethodName: "SetLpdetectGlobalCfg",
			Handler:    _Cmllpdetecservice_SetLpdetectGlobalCfg_Handler,
		},
		{
			MethodName: "DelLpdetectGlobalCfg",
			Handler:    _Cmllpdetecservice_DelLpdetectGlobalCfg_Handler,
		},
		{
			MethodName: "ShowLpdetectGlobalCfg",
			Handler:    _Cmllpdetecservice_ShowLpdetectGlobalCfg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loopdetect.proto",
}
