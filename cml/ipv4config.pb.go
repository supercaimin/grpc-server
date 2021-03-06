// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv4config.proto

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

//接口下IP地址配置; 接口下可支持配多个地址,不再采用主从模式(如果要用接口地址与外部建协议链接,
//需要考虑配置的地址被删除的情况(原来采用主从,建链使用主地址,主地址是接口上最后一个被删除的地址))
type Ipv4Ifcfg struct {
	Ifname                string   `protobuf:"bytes,1,opt,name=ifname,proto3" json:"ifname,omitempty"`
	Ipaddr                string   `protobuf:"bytes,2,opt,name=ipaddr,proto3" json:"ipaddr,omitempty"`
	Addrmask              string   `protobuf:"bytes,3,opt,name=addrmask,proto3" json:"addrmask,omitempty"`
	Masklen               int32    `protobuf:"varint,4,opt,name=masklen,proto3" json:"masklen,omitempty"`
	Unnumberifname        string   `protobuf:"bytes,5,opt,name=unnumberifname,proto3" json:"unnumberifname,omitempty"`
	Icmpttlexceedsource   int32    `protobuf:"varint,6,opt,name=icmpttlexceedsource,proto3" json:"icmpttlexceedsource,omitempty"`
	Icmpportunreachsource int32    `protobuf:"varint,7,opt,name=icmpportunreachsource,proto3" json:"icmpportunreachsource,omitempty"`
	Updatetime            int64    `protobuf:"varint,8,opt,name=updatetime,proto3" json:"updatetime,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Ipv4Ifcfg) Reset()         { *m = Ipv4Ifcfg{} }
func (m *Ipv4Ifcfg) String() string { return proto.CompactTextString(m) }
func (*Ipv4Ifcfg) ProtoMessage()    {}
func (*Ipv4Ifcfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb0be90e7a1434cb, []int{0}
}

func (m *Ipv4Ifcfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4Ifcfg.Unmarshal(m, b)
}
func (m *Ipv4Ifcfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4Ifcfg.Marshal(b, m, deterministic)
}
func (m *Ipv4Ifcfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4Ifcfg.Merge(m, src)
}
func (m *Ipv4Ifcfg) XXX_Size() int {
	return xxx_messageInfo_Ipv4Ifcfg.Size(m)
}
func (m *Ipv4Ifcfg) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4Ifcfg.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4Ifcfg proto.InternalMessageInfo

func (m *Ipv4Ifcfg) GetIfname() string {
	if m != nil {
		return m.Ifname
	}
	return ""
}

func (m *Ipv4Ifcfg) GetIpaddr() string {
	if m != nil {
		return m.Ipaddr
	}
	return ""
}

func (m *Ipv4Ifcfg) GetAddrmask() string {
	if m != nil {
		return m.Addrmask
	}
	return ""
}

func (m *Ipv4Ifcfg) GetMasklen() int32 {
	if m != nil {
		return m.Masklen
	}
	return 0
}

func (m *Ipv4Ifcfg) GetUnnumberifname() string {
	if m != nil {
		return m.Unnumberifname
	}
	return ""
}

func (m *Ipv4Ifcfg) GetIcmpttlexceedsource() int32 {
	if m != nil {
		return m.Icmpttlexceedsource
	}
	return 0
}

func (m *Ipv4Ifcfg) GetIcmpportunreachsource() int32 {
	if m != nil {
		return m.Icmpportunreachsource
	}
	return 0
}

func (m *Ipv4Ifcfg) GetUpdatetime() int64 {
	if m != nil {
		return m.Updatetime
	}
	return 0
}

//arp的相关配置
//接口下的arp的相关配置
type Arpifcfg struct {
	Ifname               string   `protobuf:"bytes,1,opt,name=ifname,proto3" json:"ifname,omitempty"`
	Linkdowndelete       int32    `protobuf:"varint,2,opt,name=linkdowndelete,proto3" json:"linkdowndelete,omitempty"`
	Detectinterval       int32    `protobuf:"varint,3,opt,name=detectinterval,proto3" json:"detectinterval,omitempty"`
	Detecttimes          int32    `protobuf:"varint,4,opt,name=detecttimes,proto3" json:"detecttimes,omitempty"`
	L2Proxyenable        int32    `protobuf:"varint,5,opt,name=l2proxyenable,proto3" json:"l2proxyenable,omitempty"`
	Proxyenable          int32    `protobuf:"varint,6,opt,name=proxyenable,proto3" json:"proxyenable,omitempty"`
	Intervlanproxy       int32    `protobuf:"varint,7,opt,name=intervlanproxy,proto3" json:"intervlanproxy,omitempty"`
	Timeout              int32    `protobuf:"varint,8,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Updatetime           int64    `protobuf:"varint,9,opt,name=updatetime,proto3" json:"updatetime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Arpifcfg) Reset()         { *m = Arpifcfg{} }
func (m *Arpifcfg) String() string { return proto.CompactTextString(m) }
func (*Arpifcfg) ProtoMessage()    {}
func (*Arpifcfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb0be90e7a1434cb, []int{1}
}

func (m *Arpifcfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Arpifcfg.Unmarshal(m, b)
}
func (m *Arpifcfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Arpifcfg.Marshal(b, m, deterministic)
}
func (m *Arpifcfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Arpifcfg.Merge(m, src)
}
func (m *Arpifcfg) XXX_Size() int {
	return xxx_messageInfo_Arpifcfg.Size(m)
}
func (m *Arpifcfg) XXX_DiscardUnknown() {
	xxx_messageInfo_Arpifcfg.DiscardUnknown(m)
}

var xxx_messageInfo_Arpifcfg proto.InternalMessageInfo

func (m *Arpifcfg) GetIfname() string {
	if m != nil {
		return m.Ifname
	}
	return ""
}

func (m *Arpifcfg) GetLinkdowndelete() int32 {
	if m != nil {
		return m.Linkdowndelete
	}
	return 0
}

func (m *Arpifcfg) GetDetectinterval() int32 {
	if m != nil {
		return m.Detectinterval
	}
	return 0
}

func (m *Arpifcfg) GetDetecttimes() int32 {
	if m != nil {
		return m.Detecttimes
	}
	return 0
}

func (m *Arpifcfg) GetL2Proxyenable() int32 {
	if m != nil {
		return m.L2Proxyenable
	}
	return 0
}

func (m *Arpifcfg) GetProxyenable() int32 {
	if m != nil {
		return m.Proxyenable
	}
	return 0
}

func (m *Arpifcfg) GetIntervlanproxy() int32 {
	if m != nil {
		return m.Intervlanproxy
	}
	return 0
}

func (m *Arpifcfg) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Arpifcfg) GetUpdatetime() int64 {
	if m != nil {
		return m.Updatetime
	}
	return 0
}

//静态arp配置
type Arpstaticcfg struct {
	Ipaddr               string   `protobuf:"bytes,1,opt,name=ipaddr,proto3" json:"ipaddr,omitempty"`
	Macaddr              string   `protobuf:"bytes,2,opt,name=macaddr,proto3" json:"macaddr,omitempty"`
	Vpnname              string   `protobuf:"bytes,3,opt,name=vpnname,proto3" json:"vpnname,omitempty"`
	Vlanid               int32    `protobuf:"varint,4,opt,name=vlanid,proto3" json:"vlanid,omitempty"`
	Cvlanid              int32    `protobuf:"varint,5,opt,name=cvlanid,proto3" json:"cvlanid,omitempty"`
	Iftype               string   `protobuf:"bytes,6,opt,name=iftype,proto3" json:"iftype,omitempty"`
	Ifnumber             int32    `protobuf:"varint,7,opt,name=ifnumber,proto3" json:"ifnumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Arpstaticcfg) Reset()         { *m = Arpstaticcfg{} }
func (m *Arpstaticcfg) String() string { return proto.CompactTextString(m) }
func (*Arpstaticcfg) ProtoMessage()    {}
func (*Arpstaticcfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb0be90e7a1434cb, []int{2}
}

func (m *Arpstaticcfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Arpstaticcfg.Unmarshal(m, b)
}
func (m *Arpstaticcfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Arpstaticcfg.Marshal(b, m, deterministic)
}
func (m *Arpstaticcfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Arpstaticcfg.Merge(m, src)
}
func (m *Arpstaticcfg) XXX_Size() int {
	return xxx_messageInfo_Arpstaticcfg.Size(m)
}
func (m *Arpstaticcfg) XXX_DiscardUnknown() {
	xxx_messageInfo_Arpstaticcfg.DiscardUnknown(m)
}

var xxx_messageInfo_Arpstaticcfg proto.InternalMessageInfo

func (m *Arpstaticcfg) GetIpaddr() string {
	if m != nil {
		return m.Ipaddr
	}
	return ""
}

func (m *Arpstaticcfg) GetMacaddr() string {
	if m != nil {
		return m.Macaddr
	}
	return ""
}

func (m *Arpstaticcfg) GetVpnname() string {
	if m != nil {
		return m.Vpnname
	}
	return ""
}

func (m *Arpstaticcfg) GetVlanid() int32 {
	if m != nil {
		return m.Vlanid
	}
	return 0
}

func (m *Arpstaticcfg) GetCvlanid() int32 {
	if m != nil {
		return m.Cvlanid
	}
	return 0
}

func (m *Arpstaticcfg) GetIftype() string {
	if m != nil {
		return m.Iftype
	}
	return ""
}

func (m *Arpstaticcfg) GetIfnumber() int32 {
	if m != nil {
		return m.Ifnumber
	}
	return 0
}

//全局arp的相关配置
type Arpglobalcfg struct {
	Detectinterval       int32           `protobuf:"varint,1,opt,name=detectinterval,proto3" json:"detectinterval,omitempty"`
	Detecttimes          int32           `protobuf:"varint,2,opt,name=detecttimes,proto3" json:"detecttimes,omitempty"`
	Conflictdetect       int32           `protobuf:"varint,3,opt,name=conflictdetect,proto3" json:"conflictdetect,omitempty"`
	Arpstatic            []*Arpstaticcfg `protobuf:"bytes,4,rep,name=arpstatic,proto3" json:"arpstatic,omitempty"`
	Arpstaticmaxnum      int32           `protobuf:"varint,5,opt,name=arpstaticmaxnum,proto3" json:"arpstaticmaxnum,omitempty"`
	Timeout              int32           `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Updatetime           int64           `protobuf:"varint,7,opt,name=updatetime,proto3" json:"updatetime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Arpglobalcfg) Reset()         { *m = Arpglobalcfg{} }
func (m *Arpglobalcfg) String() string { return proto.CompactTextString(m) }
func (*Arpglobalcfg) ProtoMessage()    {}
func (*Arpglobalcfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb0be90e7a1434cb, []int{3}
}

func (m *Arpglobalcfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Arpglobalcfg.Unmarshal(m, b)
}
func (m *Arpglobalcfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Arpglobalcfg.Marshal(b, m, deterministic)
}
func (m *Arpglobalcfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Arpglobalcfg.Merge(m, src)
}
func (m *Arpglobalcfg) XXX_Size() int {
	return xxx_messageInfo_Arpglobalcfg.Size(m)
}
func (m *Arpglobalcfg) XXX_DiscardUnknown() {
	xxx_messageInfo_Arpglobalcfg.DiscardUnknown(m)
}

var xxx_messageInfo_Arpglobalcfg proto.InternalMessageInfo

func (m *Arpglobalcfg) GetDetectinterval() int32 {
	if m != nil {
		return m.Detectinterval
	}
	return 0
}

func (m *Arpglobalcfg) GetDetecttimes() int32 {
	if m != nil {
		return m.Detecttimes
	}
	return 0
}

func (m *Arpglobalcfg) GetConflictdetect() int32 {
	if m != nil {
		return m.Conflictdetect
	}
	return 0
}

func (m *Arpglobalcfg) GetArpstatic() []*Arpstaticcfg {
	if m != nil {
		return m.Arpstatic
	}
	return nil
}

func (m *Arpglobalcfg) GetArpstaticmaxnum() int32 {
	if m != nil {
		return m.Arpstaticmaxnum
	}
	return 0
}

func (m *Arpglobalcfg) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Arpglobalcfg) GetUpdatetime() int64 {
	if m != nil {
		return m.Updatetime
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4Ifcfg)(nil), "cml.ipv4ifcfg")
	proto.RegisterType((*Arpifcfg)(nil), "cml.arpifcfg")
	proto.RegisterType((*Arpstaticcfg)(nil), "cml.arpstaticcfg")
	proto.RegisterType((*Arpglobalcfg)(nil), "cml.arpglobalcfg")
}

func init() {
	proto.RegisterFile("ipv4config.proto", fileDescriptor_eb0be90e7a1434cb)
}

var fileDescriptor_eb0be90e7a1434cb = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xd1, 0x4e, 0xdb, 0x30,
	0x14, 0x5d, 0x5a, 0x5a, 0xe8, 0x2d, 0xb4, 0xcc, 0x68, 0x53, 0xd4, 0x87, 0xa9, 0xaa, 0x26, 0xd4,
	0x27, 0xd8, 0x18, 0x13, 0x93, 0xf6, 0x34, 0x81, 0x34, 0x21, 0x6d, 0x0f, 0x0b, 0x5f, 0xe0, 0x3a,
	0x4e, 0xb0, 0x70, 0x6c, 0xcb, 0x71, 0x0a, 0x7c, 0xc8, 0x3e, 0x60, 0x3f, 0x32, 0xed, 0xb3, 0xf6,
	0x38, 0xd9, 0x4e, 0xd2, 0x34, 0x30, 0x44, 0x9f, 0xaa, 0x73, 0x7c, 0x8f, 0xef, 0xcd, 0xb9, 0x27,
	0x29, 0xec, 0x33, 0xb5, 0x3c, 0x25, 0x52, 0x24, 0x2c, 0x3d, 0x52, 0x5a, 0x1a, 0x89, 0xba, 0x24,
	0xe3, 0x93, 0x31, 0x49, 0x52, 0x22, 0xb3, 0x4c, 0x0a, 0xcf, 0xce, 0x7e, 0x75, 0x60, 0x60, 0x4b,
	0x59, 0x42, 0x92, 0x14, 0xbd, 0x86, 0x3e, 0x4b, 0x04, 0xce, 0x68, 0x18, 0x4c, 0x83, 0xf9, 0x20,
	0x2a, 0x91, 0xe3, 0x15, 0x8e, 0x63, 0x1d, 0x76, 0x4a, 0xde, 0x21, 0x34, 0x81, 0x1d, 0xfb, 0x9b,
	0xe1, 0xfc, 0x26, 0xec, 0xba, 0x93, 0x1a, 0xa3, 0x10, 0xb6, 0xed, 0x2f, 0xa7, 0x22, 0xdc, 0x9a,
	0x06, 0xf3, 0x5e, 0x54, 0x41, 0x74, 0x08, 0xa3, 0x42, 0x88, 0x22, 0x5b, 0x50, 0x5d, 0x76, 0xeb,
	0x39, 0x6d, 0x8b, 0x45, 0xef, 0xe0, 0x80, 0x91, 0x4c, 0x19, 0xc3, 0xe9, 0x1d, 0xa1, 0x34, 0xce,
	0x65, 0xa1, 0x09, 0x0d, 0xfb, 0xee, 0xb6, 0xc7, 0x8e, 0xd0, 0x29, 0xbc, 0xb2, 0xb4, 0x92, 0xda,
	0x14, 0x42, 0x53, 0x4c, 0xae, 0x4b, 0xcd, 0xb6, 0xd3, 0x3c, 0x7e, 0x88, 0xde, 0x00, 0x14, 0x2a,
	0xc6, 0x86, 0x1a, 0x96, 0xd1, 0x70, 0x67, 0x1a, 0xcc, 0xbb, 0x51, 0x83, 0x99, 0xfd, 0xee, 0xc0,
	0x0e, 0xd6, 0xea, 0x69, 0x8b, 0x0e, 0x61, 0xc4, 0x99, 0xb8, 0x89, 0xe5, 0xad, 0x88, 0x29, 0xa7,
	0x86, 0x3a, 0xab, 0x7a, 0x51, 0x8b, 0xb5, 0x75, 0x31, 0x35, 0x94, 0x18, 0x26, 0x0c, 0xd5, 0x4b,
	0xcc, 0x9d, 0x71, 0xbd, 0xa8, 0xc5, 0xa2, 0x29, 0x0c, 0x3d, 0x63, 0x47, 0xc8, 0x4b, 0x0b, 0x9b,
	0x14, 0x7a, 0x0b, 0x7b, 0xfc, 0x44, 0x69, 0x79, 0x77, 0x4f, 0x05, 0x5e, 0x70, 0xef, 0x62, 0x2f,
	0x5a, 0x27, 0xed, 0x3d, 0xcd, 0x1a, 0x6f, 0x5e, 0x93, 0xb2, 0x13, 0xf9, 0xae, 0x1c, 0x0b, 0xc7,
	0x97, 0x6e, 0xb5, 0x58, 0xbb, 0x50, 0xdb, 0x58, 0x16, 0xc6, 0x79, 0xd4, 0x8b, 0x2a, 0xd8, 0x32,
	0x70, 0xf0, 0xc0, 0xc0, 0x3f, 0x01, 0xec, 0x62, 0xad, 0x72, 0x83, 0x0d, 0x23, 0x95, 0x89, 0x3e,
	0x4f, 0xc1, 0x5a, 0x9e, 0x5c, 0x66, 0x48, 0x23, 0x68, 0x15, 0xb4, 0x27, 0x4b, 0x25, 0x9c, 0xef,
	0x3e, 0x68, 0x15, 0xb4, 0x77, 0xd9, 0x19, 0x59, 0x5c, 0x7a, 0x54, 0x22, 0xab, 0x20, 0xe5, 0x81,
	0x37, 0xa6, 0x82, 0x7e, 0x85, 0xe6, 0x5e, 0x79, 0x37, 0xdc, 0x0a, 0x2d, 0xb2, 0x69, 0x66, 0x89,
	0x4f, 0x60, 0x69, 0x41, 0x8d, 0x67, 0x3f, 0x3b, 0xee, 0x11, 0x52, 0x2e, 0x17, 0x98, 0xdb, 0x47,
	0x78, 0xb8, 0xc7, 0xe0, 0x39, 0x7b, 0xec, 0x3c, 0xdc, 0xe3, 0x21, 0x8c, 0xec, 0x8b, 0xca, 0x19,
	0x31, 0x9e, 0xae, 0x12, 0xb1, 0xce, 0xa2, 0x63, 0x18, 0xd4, 0x26, 0x86, 0x5b, 0xd3, 0xee, 0x7c,
	0x78, 0xf2, 0xf2, 0x88, 0x64, 0xfc, 0xa8, 0x69, 0x6d, 0xb4, 0xaa, 0x41, 0x73, 0x18, 0xd7, 0x20,
	0xc3, 0x77, 0xa2, 0xc8, 0x4a, 0x27, 0xda, 0x74, 0x73, 0xb5, 0xfd, 0xa7, 0x56, 0xbb, 0xdd, 0x5e,
	0xed, 0xc9, 0xdf, 0x2e, 0x8c, 0xce, 0x33, 0x6e, 0x3f, 0x21, 0x39, 0xd5, 0x4b, 0x46, 0x28, 0x7a,
	0x0f, 0xbb, 0x57, 0xd4, 0x5c, 0xaa, 0xe5, 0xe9, 0x65, 0x72, 0x9e, 0xa4, 0x68, 0xe4, 0x86, 0xac,
	0x3f, 0x32, 0x93, 0xb1, 0xc3, 0x24, 0x49, 0xb5, 0x11, 0x4c, 0x24, 0x72, 0xf6, 0xc2, 0x4a, 0x2e,
	0x28, 0xdf, 0x48, 0xf2, 0x11, 0xf6, 0xae, 0xae, 0xe5, 0xed, 0x4a, 0xb3, 0xef, 0x6a, 0xf2, 0x6b,
	0x79, 0x4b, 0x92, 0xd4, 0x16, 0x4d, 0x56, 0xcc, 0x4a, 0x76, 0x0c, 0xc3, 0x2b, 0x6a, 0xbe, 0x68,
	0xe5, 0x45, 0x7b, 0x95, 0x81, 0xff, 0xed, 0x73, 0x0c, 0xc3, 0x0b, 0xca, 0x37, 0x10, 0x7c, 0x86,
	0x03, 0x3b, 0xd8, 0x0f, 0x99, 0x7f, 0xb3, 0xaf, 0xd7, 0x77, 0x19, 0xd3, 0xe7, 0x8f, 0x77, 0x06,
	0x63, 0x3f, 0xde, 0x57, 0x17, 0x34, 0x2b, 0xac, 0x77, 0x5c, 0x67, 0xef, 0xb1, 0xae, 0x67, 0x30,
	0xf6, 0x63, 0x6e, 0x2a, 0xfc, 0x04, 0xfb, 0x76, 0xdc, 0x35, 0xe5, 0xb3, 0x66, 0x5d, 0xf4, 0xdd,
	0x3f, 0xc8, 0x87, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x36, 0x38, 0x2d, 0x6b, 0x06, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// Cmlipv4ServiceClient is the client API for Cmlipv4Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Cmlipv4ServiceClient interface {
	//ipv4在接口上的基础配置
	SetIpv4IfCfg(ctx context.Context, in *Ipv4Ifcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error)
	DelIpv4IfCfg(ctx context.Context, in *Ipv4Ifcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error)
	ShowIpv4IfCfg(ctx context.Context, in *Showcfginfo, opts ...grpc.CallOption) (*Showrtninfo, error)
	//arp在接口下的相关配置
	SetArpIfCfg(ctx context.Context, in *Arpifcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error)
	DelArpIfCfg(ctx context.Context, in *Arpifcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error)
	ShowQosLableModeCfg(ctx context.Context, in *Showcfginfo, opts ...grpc.CallOption) (*Showrtninfo, error)
	//arp 全局配置
	SetArpGlobalCfg(ctx context.Context, in *Arpglobalcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error)
	DelArpGlobalCfg(ctx context.Context, in *Arpglobalcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error)
	ShowArpGlobalCfg(ctx context.Context, in *Showcfginfo, opts ...grpc.CallOption) (*Showrtninfo, error)
}

type cmlipv4ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCmlipv4ServiceClient(cc grpc.ClientConnInterface) Cmlipv4ServiceClient {
	return &cmlipv4ServiceClient{cc}
}

func (c *cmlipv4ServiceClient) SetIpv4IfCfg(ctx context.Context, in *Ipv4Ifcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error) {
	out := new(Cfgrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmlipv4service/SetIpv4IfCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmlipv4ServiceClient) DelIpv4IfCfg(ctx context.Context, in *Ipv4Ifcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error) {
	out := new(Cfgrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmlipv4service/DelIpv4IfCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmlipv4ServiceClient) ShowIpv4IfCfg(ctx context.Context, in *Showcfginfo, opts ...grpc.CallOption) (*Showrtninfo, error) {
	out := new(Showrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmlipv4service/ShowIpv4IfCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmlipv4ServiceClient) SetArpIfCfg(ctx context.Context, in *Arpifcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error) {
	out := new(Cfgrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmlipv4service/SetArpIfCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmlipv4ServiceClient) DelArpIfCfg(ctx context.Context, in *Arpifcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error) {
	out := new(Cfgrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmlipv4service/DelArpIfCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmlipv4ServiceClient) ShowQosLableModeCfg(ctx context.Context, in *Showcfginfo, opts ...grpc.CallOption) (*Showrtninfo, error) {
	out := new(Showrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmlipv4service/ShowQosLableModeCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmlipv4ServiceClient) SetArpGlobalCfg(ctx context.Context, in *Arpglobalcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error) {
	out := new(Cfgrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmlipv4service/SetArpGlobalCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmlipv4ServiceClient) DelArpGlobalCfg(ctx context.Context, in *Arpglobalcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error) {
	out := new(Cfgrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmlipv4service/DelArpGlobalCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmlipv4ServiceClient) ShowArpGlobalCfg(ctx context.Context, in *Showcfginfo, opts ...grpc.CallOption) (*Showrtninfo, error) {
	out := new(Showrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmlipv4service/ShowArpGlobalCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Cmlipv4ServiceServer is the server API for Cmlipv4Service service.
type Cmlipv4ServiceServer interface {
	//ipv4在接口上的基础配置
	SetIpv4IfCfg(context.Context, *Ipv4Ifcfg) (*Cfgrtninfo, error)
	DelIpv4IfCfg(context.Context, *Ipv4Ifcfg) (*Cfgrtninfo, error)
	ShowIpv4IfCfg(context.Context, *Showcfginfo) (*Showrtninfo, error)
	//arp在接口下的相关配置
	SetArpIfCfg(context.Context, *Arpifcfg) (*Cfgrtninfo, error)
	DelArpIfCfg(context.Context, *Arpifcfg) (*Cfgrtninfo, error)
	ShowQosLableModeCfg(context.Context, *Showcfginfo) (*Showrtninfo, error)
	//arp 全局配置
	SetArpGlobalCfg(context.Context, *Arpglobalcfg) (*Cfgrtninfo, error)
	DelArpGlobalCfg(context.Context, *Arpglobalcfg) (*Cfgrtninfo, error)
	ShowArpGlobalCfg(context.Context, *Showcfginfo) (*Showrtninfo, error)
}

// UnimplementedCmlipv4ServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCmlipv4ServiceServer struct {
}

func (*UnimplementedCmlipv4ServiceServer) SetIpv4IfCfg(ctx context.Context, req *Ipv4Ifcfg) (*Cfgrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIpv4IfCfg not implemented")
}
func (*UnimplementedCmlipv4ServiceServer) DelIpv4IfCfg(ctx context.Context, req *Ipv4Ifcfg) (*Cfgrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelIpv4IfCfg not implemented")
}
func (*UnimplementedCmlipv4ServiceServer) ShowIpv4IfCfg(ctx context.Context, req *Showcfginfo) (*Showrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowIpv4IfCfg not implemented")
}
func (*UnimplementedCmlipv4ServiceServer) SetArpIfCfg(ctx context.Context, req *Arpifcfg) (*Cfgrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetArpIfCfg not implemented")
}
func (*UnimplementedCmlipv4ServiceServer) DelArpIfCfg(ctx context.Context, req *Arpifcfg) (*Cfgrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelArpIfCfg not implemented")
}
func (*UnimplementedCmlipv4ServiceServer) ShowQosLableModeCfg(ctx context.Context, req *Showcfginfo) (*Showrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowQosLableModeCfg not implemented")
}
func (*UnimplementedCmlipv4ServiceServer) SetArpGlobalCfg(ctx context.Context, req *Arpglobalcfg) (*Cfgrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetArpGlobalCfg not implemented")
}
func (*UnimplementedCmlipv4ServiceServer) DelArpGlobalCfg(ctx context.Context, req *Arpglobalcfg) (*Cfgrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelArpGlobalCfg not implemented")
}
func (*UnimplementedCmlipv4ServiceServer) ShowArpGlobalCfg(ctx context.Context, req *Showcfginfo) (*Showrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowArpGlobalCfg not implemented")
}

func RegisterCmlipv4ServiceServer(s *grpc.Server, srv Cmlipv4ServiceServer) {
	s.RegisterService(&_Cmlipv4Service_serviceDesc, srv)
}

func _Cmlipv4Service_SetIpv4IfCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ipv4Ifcfg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Cmlipv4ServiceServer).SetIpv4IfCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmlipv4service/SetIpv4IfCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Cmlipv4ServiceServer).SetIpv4IfCfg(ctx, req.(*Ipv4Ifcfg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmlipv4Service_DelIpv4IfCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ipv4Ifcfg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Cmlipv4ServiceServer).DelIpv4IfCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmlipv4service/DelIpv4IfCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Cmlipv4ServiceServer).DelIpv4IfCfg(ctx, req.(*Ipv4Ifcfg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmlipv4Service_ShowIpv4IfCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Showcfginfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Cmlipv4ServiceServer).ShowIpv4IfCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmlipv4service/ShowIpv4IfCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Cmlipv4ServiceServer).ShowIpv4IfCfg(ctx, req.(*Showcfginfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmlipv4Service_SetArpIfCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Arpifcfg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Cmlipv4ServiceServer).SetArpIfCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmlipv4service/SetArpIfCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Cmlipv4ServiceServer).SetArpIfCfg(ctx, req.(*Arpifcfg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmlipv4Service_DelArpIfCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Arpifcfg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Cmlipv4ServiceServer).DelArpIfCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmlipv4service/DelArpIfCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Cmlipv4ServiceServer).DelArpIfCfg(ctx, req.(*Arpifcfg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmlipv4Service_ShowQosLableModeCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Showcfginfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Cmlipv4ServiceServer).ShowQosLableModeCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmlipv4service/ShowQosLableModeCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Cmlipv4ServiceServer).ShowQosLableModeCfg(ctx, req.(*Showcfginfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmlipv4Service_SetArpGlobalCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Arpglobalcfg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Cmlipv4ServiceServer).SetArpGlobalCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmlipv4service/SetArpGlobalCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Cmlipv4ServiceServer).SetArpGlobalCfg(ctx, req.(*Arpglobalcfg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmlipv4Service_DelArpGlobalCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Arpglobalcfg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Cmlipv4ServiceServer).DelArpGlobalCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmlipv4service/DelArpGlobalCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Cmlipv4ServiceServer).DelArpGlobalCfg(ctx, req.(*Arpglobalcfg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmlipv4Service_ShowArpGlobalCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Showcfginfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Cmlipv4ServiceServer).ShowArpGlobalCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmlipv4service/ShowArpGlobalCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Cmlipv4ServiceServer).ShowArpGlobalCfg(ctx, req.(*Showcfginfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cmlipv4Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cml.Cmlipv4service",
	HandlerType: (*Cmlipv4ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetIpv4IfCfg",
			Handler:    _Cmlipv4Service_SetIpv4IfCfg_Handler,
		},
		{
			MethodName: "DelIpv4IfCfg",
			Handler:    _Cmlipv4Service_DelIpv4IfCfg_Handler,
		},
		{
			MethodName: "ShowIpv4IfCfg",
			Handler:    _Cmlipv4Service_ShowIpv4IfCfg_Handler,
		},
		{
			MethodName: "SetArpIfCfg",
			Handler:    _Cmlipv4Service_SetArpIfCfg_Handler,
		},
		{
			MethodName: "DelArpIfCfg",
			Handler:    _Cmlipv4Service_DelArpIfCfg_Handler,
		},
		{
			MethodName: "ShowQosLableModeCfg",
			Handler:    _Cmlipv4Service_ShowQosLableModeCfg_Handler,
		},
		{
			MethodName: "SetArpGlobalCfg",
			Handler:    _Cmlipv4Service_SetArpGlobalCfg_Handler,
		},
		{
			MethodName: "DelArpGlobalCfg",
			Handler:    _Cmlipv4Service_DelArpGlobalCfg_Handler,
		},
		{
			MethodName: "ShowArpGlobalCfg",
			Handler:    _Cmlipv4Service_ShowArpGlobalCfg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ipv4config.proto",
}
