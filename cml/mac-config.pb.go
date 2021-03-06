// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mac-config.proto

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

//黑洞mac配置信息
type Macblackhole struct {
	Mac                  string   `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	Vlanid               int32    `protobuf:"varint,2,opt,name=vlanid,proto3" json:"vlanid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Macblackhole) Reset()         { *m = Macblackhole{} }
func (m *Macblackhole) String() string { return proto.CompactTextString(m) }
func (*Macblackhole) ProtoMessage()    {}
func (*Macblackhole) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ce91e21daa58402, []int{0}
}

func (m *Macblackhole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Macblackhole.Unmarshal(m, b)
}
func (m *Macblackhole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Macblackhole.Marshal(b, m, deterministic)
}
func (m *Macblackhole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Macblackhole.Merge(m, src)
}
func (m *Macblackhole) XXX_Size() int {
	return xxx_messageInfo_Macblackhole.Size(m)
}
func (m *Macblackhole) XXX_DiscardUnknown() {
	xxx_messageInfo_Macblackhole.DiscardUnknown(m)
}

var xxx_messageInfo_Macblackhole proto.InternalMessageInfo

func (m *Macblackhole) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *Macblackhole) GetVlanid() int32 {
	if m != nil {
		return m.Vlanid
	}
	return 0
}

//undo命令若只指定ifname,则删除从该接口出的所有静态mac
type Macstatic struct {
	Mac                  string   `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	Ifname               string   `protobuf:"bytes,2,opt,name=ifname,proto3" json:"ifname,omitempty"`
	Vlanid               int32    `protobuf:"varint,3,opt,name=vlanid,proto3" json:"vlanid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Macstatic) Reset()         { *m = Macstatic{} }
func (m *Macstatic) String() string { return proto.CompactTextString(m) }
func (*Macstatic) ProtoMessage()    {}
func (*Macstatic) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ce91e21daa58402, []int{1}
}

func (m *Macstatic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Macstatic.Unmarshal(m, b)
}
func (m *Macstatic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Macstatic.Marshal(b, m, deterministic)
}
func (m *Macstatic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Macstatic.Merge(m, src)
}
func (m *Macstatic) XXX_Size() int {
	return xxx_messageInfo_Macstatic.Size(m)
}
func (m *Macstatic) XXX_DiscardUnknown() {
	xxx_messageInfo_Macstatic.DiscardUnknown(m)
}

var xxx_messageInfo_Macstatic proto.InternalMessageInfo

func (m *Macstatic) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *Macstatic) GetIfname() string {
	if m != nil {
		return m.Ifname
	}
	return ""
}

func (m *Macstatic) GetVlanid() int32 {
	if m != nil {
		return m.Vlanid
	}
	return 0
}

//mac漂移相关的配置信息
type Macflapcfg struct {
	Flapingagetime       int32    `protobuf:"varint,1,opt,name=flapingagetime,proto3" json:"flapingagetime,omitempty"`
	Flapdetectlevel      string   `protobuf:"bytes,2,opt,name=flapdetectlevel,proto3" json:"flapdetectlevel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Macflapcfg) Reset()         { *m = Macflapcfg{} }
func (m *Macflapcfg) String() string { return proto.CompactTextString(m) }
func (*Macflapcfg) ProtoMessage()    {}
func (*Macflapcfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ce91e21daa58402, []int{2}
}

func (m *Macflapcfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Macflapcfg.Unmarshal(m, b)
}
func (m *Macflapcfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Macflapcfg.Marshal(b, m, deterministic)
}
func (m *Macflapcfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Macflapcfg.Merge(m, src)
}
func (m *Macflapcfg) XXX_Size() int {
	return xxx_messageInfo_Macflapcfg.Size(m)
}
func (m *Macflapcfg) XXX_DiscardUnknown() {
	xxx_messageInfo_Macflapcfg.DiscardUnknown(m)
}

var xxx_messageInfo_Macflapcfg proto.InternalMessageInfo

func (m *Macflapcfg) GetFlapingagetime() int32 {
	if m != nil {
		return m.Flapingagetime
	}
	return 0
}

func (m *Macflapcfg) GetFlapdetectlevel() string {
	if m != nil {
		return m.Flapdetectlevel
	}
	return ""
}

//mac全局配置信息
type Macglobalcfg struct {
	Agingtime            int32         `protobuf:"varint,1,opt,name=agingtime,proto3" json:"agingtime,omitempty"`
	Machole              *Macblackhole `protobuf:"bytes,2,opt,name=machole,proto3" json:"machole,omitempty"`
	Macflap              *Macflapcfg   `protobuf:"bytes,3,opt,name=macflap,proto3" json:"macflap,omitempty"`
	Static               []*Macstatic  `protobuf:"bytes,4,rep,name=static,proto3" json:"static,omitempty"`
	Updatetime           int64         `protobuf:"varint,6,opt,name=updatetime,proto3" json:"updatetime,omitempty"`
	Cmdid                int32         `protobuf:"varint,7,opt,name=cmdid,proto3" json:"cmdid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Macglobalcfg) Reset()         { *m = Macglobalcfg{} }
func (m *Macglobalcfg) String() string { return proto.CompactTextString(m) }
func (*Macglobalcfg) ProtoMessage()    {}
func (*Macglobalcfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ce91e21daa58402, []int{3}
}

func (m *Macglobalcfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Macglobalcfg.Unmarshal(m, b)
}
func (m *Macglobalcfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Macglobalcfg.Marshal(b, m, deterministic)
}
func (m *Macglobalcfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Macglobalcfg.Merge(m, src)
}
func (m *Macglobalcfg) XXX_Size() int {
	return xxx_messageInfo_Macglobalcfg.Size(m)
}
func (m *Macglobalcfg) XXX_DiscardUnknown() {
	xxx_messageInfo_Macglobalcfg.DiscardUnknown(m)
}

var xxx_messageInfo_Macglobalcfg proto.InternalMessageInfo

func (m *Macglobalcfg) GetAgingtime() int32 {
	if m != nil {
		return m.Agingtime
	}
	return 0
}

func (m *Macglobalcfg) GetMachole() *Macblackhole {
	if m != nil {
		return m.Machole
	}
	return nil
}

func (m *Macglobalcfg) GetMacflap() *Macflapcfg {
	if m != nil {
		return m.Macflap
	}
	return nil
}

func (m *Macglobalcfg) GetStatic() []*Macstatic {
	if m != nil {
		return m.Static
	}
	return nil
}

func (m *Macglobalcfg) GetUpdatetime() int64 {
	if m != nil {
		return m.Updatetime
	}
	return 0
}

func (m *Macglobalcfg) GetCmdid() int32 {
	if m != nil {
		return m.Cmdid
	}
	return 0
}

type Maclearndis struct {
	Learndisable         int32    `protobuf:"varint,1,opt,name=learndisable,proto3" json:"learndisable,omitempty"`
	Action               string   `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Maclearndis) Reset()         { *m = Maclearndis{} }
func (m *Maclearndis) String() string { return proto.CompactTextString(m) }
func (*Maclearndis) ProtoMessage()    {}
func (*Maclearndis) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ce91e21daa58402, []int{4}
}

func (m *Maclearndis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Maclearndis.Unmarshal(m, b)
}
func (m *Maclearndis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Maclearndis.Marshal(b, m, deterministic)
}
func (m *Maclearndis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Maclearndis.Merge(m, src)
}
func (m *Maclearndis) XXX_Size() int {
	return xxx_messageInfo_Maclearndis.Size(m)
}
func (m *Maclearndis) XXX_DiscardUnknown() {
	xxx_messageInfo_Maclearndis.DiscardUnknown(m)
}

var xxx_messageInfo_Maclearndis proto.InternalMessageInfo

func (m *Maclearndis) GetLearndisable() int32 {
	if m != nil {
		return m.Learndisable
	}
	return 0
}

func (m *Maclearndis) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

//mac最大学习量,action,alarm可选配
type Maclimit struct {
	Maxnum               int32    `protobuf:"varint,1,opt,name=maxnum,proto3" json:"maxnum,omitempty"`
	Action               string   `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Alarm                int32    `protobuf:"varint,3,opt,name=alarm,proto3" json:"alarm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Maclimit) Reset()         { *m = Maclimit{} }
func (m *Maclimit) String() string { return proto.CompactTextString(m) }
func (*Maclimit) ProtoMessage()    {}
func (*Maclimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ce91e21daa58402, []int{5}
}

func (m *Maclimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Maclimit.Unmarshal(m, b)
}
func (m *Maclimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Maclimit.Marshal(b, m, deterministic)
}
func (m *Maclimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Maclimit.Merge(m, src)
}
func (m *Maclimit) XXX_Size() int {
	return xxx_messageInfo_Maclimit.Size(m)
}
func (m *Maclimit) XXX_DiscardUnknown() {
	xxx_messageInfo_Maclimit.DiscardUnknown(m)
}

var xxx_messageInfo_Maclimit proto.InternalMessageInfo

func (m *Maclimit) GetMaxnum() int32 {
	if m != nil {
		return m.Maxnum
	}
	return 0
}

func (m *Maclimit) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Maclimit) GetAlarm() int32 {
	if m != nil {
		return m.Alarm
	}
	return 0
}

//接口下的mac配置信息
type Macifcfg struct {
	Ifname               string       `protobuf:"bytes,1,opt,name=ifname,proto3" json:"ifname,omitempty"`
	Maclearn             *Maclearndis `protobuf:"bytes,2,opt,name=maclearn,proto3" json:"maclearn,omitempty"`
	Maclimit             *Maclimit    `protobuf:"bytes,3,opt,name=maclimit,proto3" json:"maclimit,omitempty"`
	Portbridge           int32        `protobuf:"varint,4,opt,name=portbridge,proto3" json:"portbridge,omitempty"`
	Updatetime           int64        `protobuf:"varint,5,opt,name=updatetime,proto3" json:"updatetime,omitempty"`
	Cmdid                int32        `protobuf:"varint,6,opt,name=cmdid,proto3" json:"cmdid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Macifcfg) Reset()         { *m = Macifcfg{} }
func (m *Macifcfg) String() string { return proto.CompactTextString(m) }
func (*Macifcfg) ProtoMessage()    {}
func (*Macifcfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ce91e21daa58402, []int{6}
}

func (m *Macifcfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Macifcfg.Unmarshal(m, b)
}
func (m *Macifcfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Macifcfg.Marshal(b, m, deterministic)
}
func (m *Macifcfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Macifcfg.Merge(m, src)
}
func (m *Macifcfg) XXX_Size() int {
	return xxx_messageInfo_Macifcfg.Size(m)
}
func (m *Macifcfg) XXX_DiscardUnknown() {
	xxx_messageInfo_Macifcfg.DiscardUnknown(m)
}

var xxx_messageInfo_Macifcfg proto.InternalMessageInfo

func (m *Macifcfg) GetIfname() string {
	if m != nil {
		return m.Ifname
	}
	return ""
}

func (m *Macifcfg) GetMaclearn() *Maclearndis {
	if m != nil {
		return m.Maclearn
	}
	return nil
}

func (m *Macifcfg) GetMaclimit() *Maclimit {
	if m != nil {
		return m.Maclimit
	}
	return nil
}

func (m *Macifcfg) GetPortbridge() int32 {
	if m != nil {
		return m.Portbridge
	}
	return 0
}

func (m *Macifcfg) GetUpdatetime() int64 {
	if m != nil {
		return m.Updatetime
	}
	return 0
}

func (m *Macifcfg) GetCmdid() int32 {
	if m != nil {
		return m.Cmdid
	}
	return 0
}

type Macinfo struct {
	Mac                  string   `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Macinfo) Reset()         { *m = Macinfo{} }
func (m *Macinfo) String() string { return proto.CompactTextString(m) }
func (*Macinfo) ProtoMessage()    {}
func (*Macinfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ce91e21daa58402, []int{7}
}

func (m *Macinfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Macinfo.Unmarshal(m, b)
}
func (m *Macinfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Macinfo.Marshal(b, m, deterministic)
}
func (m *Macinfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Macinfo.Merge(m, src)
}
func (m *Macinfo) XXX_Size() int {
	return xxx_messageInfo_Macinfo.Size(m)
}
func (m *Macinfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Macinfo.DiscardUnknown(m)
}

var xxx_messageInfo_Macinfo proto.InternalMessageInfo

func (m *Macinfo) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

type Vlanidinfo struct {
	Vlanid               int32    `protobuf:"varint,1,opt,name=vlanid,proto3" json:"vlanid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vlanidinfo) Reset()         { *m = Vlanidinfo{} }
func (m *Vlanidinfo) String() string { return proto.CompactTextString(m) }
func (*Vlanidinfo) ProtoMessage()    {}
func (*Vlanidinfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ce91e21daa58402, []int{8}
}

func (m *Vlanidinfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vlanidinfo.Unmarshal(m, b)
}
func (m *Vlanidinfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vlanidinfo.Marshal(b, m, deterministic)
}
func (m *Vlanidinfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vlanidinfo.Merge(m, src)
}
func (m *Vlanidinfo) XXX_Size() int {
	return xxx_messageInfo_Vlanidinfo.Size(m)
}
func (m *Vlanidinfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Vlanidinfo.DiscardUnknown(m)
}

var xxx_messageInfo_Vlanidinfo proto.InternalMessageInfo

func (m *Vlanidinfo) GetVlanid() int32 {
	if m != nil {
		return m.Vlanid
	}
	return 0
}

type Ifinfocfg struct {
	Ifname               string   `protobuf:"bytes,1,opt,name=ifname,proto3" json:"ifname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ifinfocfg) Reset()         { *m = Ifinfocfg{} }
func (m *Ifinfocfg) String() string { return proto.CompactTextString(m) }
func (*Ifinfocfg) ProtoMessage()    {}
func (*Ifinfocfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ce91e21daa58402, []int{9}
}

func (m *Ifinfocfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ifinfocfg.Unmarshal(m, b)
}
func (m *Ifinfocfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ifinfocfg.Marshal(b, m, deterministic)
}
func (m *Ifinfocfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ifinfocfg.Merge(m, src)
}
func (m *Ifinfocfg) XXX_Size() int {
	return xxx_messageInfo_Ifinfocfg.Size(m)
}
func (m *Ifinfocfg) XXX_DiscardUnknown() {
	xxx_messageInfo_Ifinfocfg.DiscardUnknown(m)
}

var xxx_messageInfo_Ifinfocfg proto.InternalMessageInfo

func (m *Ifinfocfg) GetIfname() string {
	if m != nil {
		return m.Ifname
	}
	return ""
}

//mac动态操作信息
type Macclearact struct {
	Mac                  []*Macinfo    `protobuf:"bytes,1,rep,name=mac,proto3" json:"mac,omitempty"`
	Vlanid               []*Vlanidinfo `protobuf:"bytes,2,rep,name=vlanid,proto3" json:"vlanid,omitempty"`
	Ifinfo               []*Ifinfocfg  `protobuf:"bytes,3,rep,name=ifinfo,proto3" json:"ifinfo,omitempty"`
	Cmdid                int32         `protobuf:"varint,4,opt,name=cmdid,proto3" json:"cmdid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Macclearact) Reset()         { *m = Macclearact{} }
func (m *Macclearact) String() string { return proto.CompactTextString(m) }
func (*Macclearact) ProtoMessage()    {}
func (*Macclearact) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ce91e21daa58402, []int{10}
}

func (m *Macclearact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Macclearact.Unmarshal(m, b)
}
func (m *Macclearact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Macclearact.Marshal(b, m, deterministic)
}
func (m *Macclearact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Macclearact.Merge(m, src)
}
func (m *Macclearact) XXX_Size() int {
	return xxx_messageInfo_Macclearact.Size(m)
}
func (m *Macclearact) XXX_DiscardUnknown() {
	xxx_messageInfo_Macclearact.DiscardUnknown(m)
}

var xxx_messageInfo_Macclearact proto.InternalMessageInfo

func (m *Macclearact) GetMac() []*Macinfo {
	if m != nil {
		return m.Mac
	}
	return nil
}

func (m *Macclearact) GetVlanid() []*Vlanidinfo {
	if m != nil {
		return m.Vlanid
	}
	return nil
}

func (m *Macclearact) GetIfinfo() []*Ifinfocfg {
	if m != nil {
		return m.Ifinfo
	}
	return nil
}

func (m *Macclearact) GetCmdid() int32 {
	if m != nil {
		return m.Cmdid
	}
	return 0
}

func init() {
	proto.RegisterType((*Macblackhole)(nil), "cml.macblackhole")
	proto.RegisterType((*Macstatic)(nil), "cml.macstatic")
	proto.RegisterType((*Macflapcfg)(nil), "cml.macflapcfg")
	proto.RegisterType((*Macglobalcfg)(nil), "cml.macglobalcfg")
	proto.RegisterType((*Maclearndis)(nil), "cml.maclearndis")
	proto.RegisterType((*Maclimit)(nil), "cml.maclimit")
	proto.RegisterType((*Macifcfg)(nil), "cml.macifcfg")
	proto.RegisterType((*Macinfo)(nil), "cml.macinfo")
	proto.RegisterType((*Vlanidinfo)(nil), "cml.vlanidinfo")
	proto.RegisterType((*Ifinfocfg)(nil), "cml.ifinfocfg")
	proto.RegisterType((*Macclearact)(nil), "cml.macclearact")
}

func init() {
	proto.RegisterFile("mac-config.proto", fileDescriptor_1ce91e21daa58402)
}

var fileDescriptor_1ce91e21daa58402 = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0xfc, 0xb9, 0x49, 0xc9, 0x24, 0x6d, 0xc2, 0x0a, 0xa1, 0xa8, 0xa0, 0x2a, 0x32, 0xa8,
	0xa4, 0x02, 0x8a, 0x14, 0x40, 0xed, 0x2d, 0x2a, 0x12, 0xea, 0x45, 0x25, 0xe4, 0xde, 0x23, 0x6d,
	0xc6, 0x6b, 0x77, 0xc5, 0xae, 0x1d, 0xb9, 0xdb, 0x96, 0x67, 0xe0, 0x9a, 0x17, 0xe3, 0x15, 0x78,
	0x12, 0x34, 0xfb, 0x63, 0x3b, 0xa5, 0x45, 0xf4, 0x6e, 0x67, 0x76, 0x7e, 0xce, 0x9c, 0x39, 0xbb,
	0x30, 0xd1, 0x1c, 0x5f, 0x63, 0x55, 0xe6, 0xb2, 0x38, 0x58, 0xd5, 0x95, 0xa9, 0x58, 0x8c, 0x5a,
	0xed, 0x8c, 0x31, 0x2f, 0xb0, 0xd2, 0xba, 0x2a, 0x9d, 0x37, 0x39, 0x82, 0x91, 0xe6, 0xb8, 0x54,
	0x1c, 0xbf, 0x9e, 0x57, 0x4a, 0xb0, 0x09, 0xc4, 0x9a, 0xe3, 0x34, 0x9a, 0x45, 0xf3, 0x41, 0x4a,
	0x47, 0xf6, 0x18, 0xfa, 0x57, 0x8a, 0x97, 0x32, 0x9b, 0xfe, 0x3f, 0x8b, 0xe6, 0xbd, 0xd4, 0x5b,
	0xc9, 0x29, 0x0c, 0x34, 0xc7, 0x0b, 0xc3, 0x8d, 0xc4, 0xdb, 0xd3, 0x64, 0x5e, 0x72, 0x2d, 0x6c,
	0xda, 0x20, 0xf5, 0x56, 0xa7, 0x5c, 0xbc, 0x56, 0xee, 0x0b, 0x80, 0xe6, 0x98, 0x2b, 0xbe, 0xc2,
	0xbc, 0x60, 0x7b, 0xb0, 0x4d, 0x47, 0x59, 0x16, 0xbc, 0x10, 0x46, 0x6a, 0x61, 0x4b, 0xf7, 0xd2,
	0x1b, 0x5e, 0x36, 0x87, 0x31, 0x79, 0x32, 0x61, 0x04, 0x1a, 0x25, 0xae, 0x84, 0xf2, 0xed, 0x6e,
	0xba, 0x93, 0x5f, 0x91, 0x9d, 0xb4, 0x50, 0xd5, 0x92, 0x2b, 0x6a, 0xf1, 0x14, 0x06, 0xbc, 0x90,
	0x65, 0xd1, 0xa9, 0xde, 0x3a, 0xd8, 0x4b, 0xd8, 0xd4, 0x1c, 0x89, 0x12, 0x5b, 0x70, 0xb8, 0x78,
	0x78, 0x80, 0x5a, 0x1d, 0x74, 0xb9, 0x4a, 0x43, 0x04, 0xdb, 0xb7, 0xc1, 0xd4, 0xd1, 0x0e, 0x35,
	0x5c, 0x8c, 0x43, 0xb0, 0x9f, 0x27, 0x0d, 0xf7, 0x6c, 0x0f, 0xfa, 0x8e, 0xb2, 0xe9, 0xc6, 0x2c,
	0x9e, 0x0f, 0x17, 0xdb, 0x21, 0xd2, 0x79, 0x53, 0x7f, 0xcb, 0x76, 0x01, 0x2e, 0x57, 0x19, 0x37,
	0x6e, 0xf8, 0xfe, 0x2c, 0x9a, 0xc7, 0x69, 0xc7, 0xc3, 0x1e, 0x41, 0x0f, 0x75, 0x26, 0xb3, 0xe9,
	0xa6, 0x45, 0xee, 0x8c, 0xe4, 0x04, 0x86, 0x9a, 0xa3, 0x12, 0xbc, 0x2e, 0x33, 0x79, 0xc1, 0x12,
	0x18, 0x85, 0x33, 0x5f, 0xaa, 0x30, 0xe5, 0x9a, 0x8f, 0xf6, 0xc1, 0xd1, 0xc8, 0xaa, 0x0c, 0x7b,
	0x72, 0x56, 0xf2, 0x19, 0x1e, 0x50, 0x29, 0xa9, 0xa5, 0xa1, 0x18, 0xcd, 0xbf, 0x95, 0x97, 0xda,
	0x57, 0xf0, 0xd6, 0x5d, 0xb9, 0x04, 0x8e, 0x2b, 0x5e, 0x6b, 0xbf, 0x62, 0x67, 0x24, 0x3f, 0x23,
	0x5b, 0x52, 0xe6, 0xc4, 0x7e, 0x2b, 0x8f, 0x68, 0x4d, 0x1e, 0xaf, 0x5c, 0x5b, 0x42, 0xe8, 0x89,
	0x9f, 0x04, 0x86, 0x02, 0xec, 0xb4, 0x89, 0x60, 0xfb, 0x2d, 0x48, 0xcf, 0xfc, 0x56, 0x13, 0x4d,
	0xce, 0xb4, 0x9d, 0x61, 0x17, 0x60, 0x55, 0xd5, 0x66, 0x59, 0xcb, 0xac, 0x10, 0xd3, 0x0d, 0x0b,
	0xac, 0xe3, 0xb9, 0x41, 0x78, 0xef, 0x6e, 0xc2, 0xfb, 0x5d, 0xc2, 0x9f, 0xd8, 0xcd, 0xcb, 0x32,
	0xaf, 0xfe, 0x7c, 0x02, 0xc9, 0x73, 0x00, 0x27, 0x6e, 0x7b, 0xdf, 0x0a, 0x3f, 0x5a, 0x13, 0xfe,
	0x33, 0x18, 0xc8, 0x9c, 0x22, 0xfe, 0x42, 0x4b, 0xf2, 0x23, 0xb2, 0x9b, 0xb5, 0x63, 0x73, 0xa4,
	0x69, 0x7c, 0x33, 0xd2, 0xd0, 0x28, 0xcc, 0x4c, 0x55, 0xdc, 0xeb, 0x7b, 0xd1, 0x79, 0xb4, 0x71,
	0x23, 0xc8, 0x16, 0x4d, 0xe8, 0x4e, 0x7a, 0x74, 0xdd, 0xa7, 0x71, 0x47, 0x8f, 0x0d, 0xa0, 0xd4,
	0xdf, 0xb6, 0xe3, 0x6f, 0x74, 0xc6, 0x5f, 0x7c, 0x8f, 0x61, 0xeb, 0x58, 0x2b, 0x92, 0xaf, 0xa8,
	0xaf, 0x24, 0x0a, 0xf6, 0x06, 0x86, 0x67, 0xc2, 0x9c, 0x72, 0x3c, 0xc9, 0x8f, 0xf3, 0x82, 0x35,
	0xeb, 0xb0, 0x5b, 0xdf, 0x71, 0x30, 0x30, 0x2f, 0x6a, 0x53, 0x52, 0xd9, 0xe4, 0x3f, 0x4a, 0xf8,
	0x28, 0xd4, 0x3d, 0x12, 0xde, 0xc1, 0xe8, 0xec, 0xbc, 0xba, 0x6e, 0x32, 0x9c, 0x3e, 0x2e, 0xce,
	0xab, 0x6b, 0xcc, 0x0b, 0x8a, 0xd9, 0x69, 0x3d, 0x6d, 0xd6, 0x21, 0x8c, 0x1d, 0xae, 0x4f, 0xf6,
	0x03, 0xa0, 0xc4, 0xe6, 0x45, 0x37, 0x7f, 0xc2, 0x6d, 0xed, 0x0e, 0x61, 0xec, 0xf0, 0xdd, 0x37,
	0xf1, 0x08, 0x26, 0x1e, 0x67, 0x9b, 0xf9, 0x6f, 0x58, 0xdf, 0xc3, 0xb6, 0xc3, 0x7a, 0x4c, 0xeb,
	0xfe, 0x80, 0x86, 0x35, 0x6f, 0x20, 0x08, 0xe0, 0x96, 0x86, 0xcb, 0xbe, 0xfd, 0xd1, 0xdf, 0xfe,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xc2, 0xc5, 0x14, 0xfb, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CmlmacserviceClient is the client API for Cmlmacservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CmlmacserviceClient interface {
	//mac在接口上的配置
	SetMacIfCfg(ctx context.Context, in *Macifcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error)
	DelMacIfCfg(ctx context.Context, in *Macifcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error)
	ShowMacIfCfg(ctx context.Context, in *Showcfginfo, opts ...grpc.CallOption) (*Showrtninfo, error)
	//mac全局配置
	SetMacGlobalCfg(ctx context.Context, in *Macglobalcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error)
	DelMacGlobalCfg(ctx context.Context, in *Macglobalcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error)
	ShowMacGlobalCfg(ctx context.Context, in *Showcfginfo, opts ...grpc.CallOption) (*Showrtninfo, error)
	//mac clear动作命令
	SetMacClearAct(ctx context.Context, in *Macclearact, opts ...grpc.CallOption) (*Cfgrtninfo, error)
}

type cmlmacserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewCmlmacserviceClient(cc grpc.ClientConnInterface) CmlmacserviceClient {
	return &cmlmacserviceClient{cc}
}

func (c *cmlmacserviceClient) SetMacIfCfg(ctx context.Context, in *Macifcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error) {
	out := new(Cfgrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmlmacservice/SetMacIfCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmlmacserviceClient) DelMacIfCfg(ctx context.Context, in *Macifcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error) {
	out := new(Cfgrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmlmacservice/DelMacIfCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmlmacserviceClient) ShowMacIfCfg(ctx context.Context, in *Showcfginfo, opts ...grpc.CallOption) (*Showrtninfo, error) {
	out := new(Showrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmlmacservice/ShowMacIfCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmlmacserviceClient) SetMacGlobalCfg(ctx context.Context, in *Macglobalcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error) {
	out := new(Cfgrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmlmacservice/SetMacGlobalCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmlmacserviceClient) DelMacGlobalCfg(ctx context.Context, in *Macglobalcfg, opts ...grpc.CallOption) (*Cfgrtninfo, error) {
	out := new(Cfgrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmlmacservice/DelMacGlobalCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmlmacserviceClient) ShowMacGlobalCfg(ctx context.Context, in *Showcfginfo, opts ...grpc.CallOption) (*Showrtninfo, error) {
	out := new(Showrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmlmacservice/ShowMacGlobalCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmlmacserviceClient) SetMacClearAct(ctx context.Context, in *Macclearact, opts ...grpc.CallOption) (*Cfgrtninfo, error) {
	out := new(Cfgrtninfo)
	err := c.cc.Invoke(ctx, "/cml.Cmlmacservice/SetMacClearAct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CmlmacserviceServer is the server API for Cmlmacservice service.
type CmlmacserviceServer interface {
	//mac在接口上的配置
	SetMacIfCfg(context.Context, *Macifcfg) (*Cfgrtninfo, error)
	DelMacIfCfg(context.Context, *Macifcfg) (*Cfgrtninfo, error)
	ShowMacIfCfg(context.Context, *Showcfginfo) (*Showrtninfo, error)
	//mac全局配置
	SetMacGlobalCfg(context.Context, *Macglobalcfg) (*Cfgrtninfo, error)
	DelMacGlobalCfg(context.Context, *Macglobalcfg) (*Cfgrtninfo, error)
	ShowMacGlobalCfg(context.Context, *Showcfginfo) (*Showrtninfo, error)
	//mac clear动作命令
	SetMacClearAct(context.Context, *Macclearact) (*Cfgrtninfo, error)
}

// UnimplementedCmlmacserviceServer can be embedded to have forward compatible implementations.
type UnimplementedCmlmacserviceServer struct {
}

func (*UnimplementedCmlmacserviceServer) SetMacIfCfg(ctx context.Context, req *Macifcfg) (*Cfgrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMacIfCfg not implemented")
}
func (*UnimplementedCmlmacserviceServer) DelMacIfCfg(ctx context.Context, req *Macifcfg) (*Cfgrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelMacIfCfg not implemented")
}
func (*UnimplementedCmlmacserviceServer) ShowMacIfCfg(ctx context.Context, req *Showcfginfo) (*Showrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowMacIfCfg not implemented")
}
func (*UnimplementedCmlmacserviceServer) SetMacGlobalCfg(ctx context.Context, req *Macglobalcfg) (*Cfgrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMacGlobalCfg not implemented")
}
func (*UnimplementedCmlmacserviceServer) DelMacGlobalCfg(ctx context.Context, req *Macglobalcfg) (*Cfgrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelMacGlobalCfg not implemented")
}
func (*UnimplementedCmlmacserviceServer) ShowMacGlobalCfg(ctx context.Context, req *Showcfginfo) (*Showrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowMacGlobalCfg not implemented")
}
func (*UnimplementedCmlmacserviceServer) SetMacClearAct(ctx context.Context, req *Macclearact) (*Cfgrtninfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMacClearAct not implemented")
}

func RegisterCmlmacserviceServer(s *grpc.Server, srv CmlmacserviceServer) {
	s.RegisterService(&_Cmlmacservice_serviceDesc, srv)
}

func _Cmlmacservice_SetMacIfCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Macifcfg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmlmacserviceServer).SetMacIfCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmlmacservice/SetMacIfCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmlmacserviceServer).SetMacIfCfg(ctx, req.(*Macifcfg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmlmacservice_DelMacIfCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Macifcfg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmlmacserviceServer).DelMacIfCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmlmacservice/DelMacIfCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmlmacserviceServer).DelMacIfCfg(ctx, req.(*Macifcfg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmlmacservice_ShowMacIfCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Showcfginfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmlmacserviceServer).ShowMacIfCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmlmacservice/ShowMacIfCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmlmacserviceServer).ShowMacIfCfg(ctx, req.(*Showcfginfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmlmacservice_SetMacGlobalCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Macglobalcfg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmlmacserviceServer).SetMacGlobalCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmlmacservice/SetMacGlobalCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmlmacserviceServer).SetMacGlobalCfg(ctx, req.(*Macglobalcfg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmlmacservice_DelMacGlobalCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Macglobalcfg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmlmacserviceServer).DelMacGlobalCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmlmacservice/DelMacGlobalCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmlmacserviceServer).DelMacGlobalCfg(ctx, req.(*Macglobalcfg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmlmacservice_ShowMacGlobalCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Showcfginfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmlmacserviceServer).ShowMacGlobalCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmlmacservice/ShowMacGlobalCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmlmacserviceServer).ShowMacGlobalCfg(ctx, req.(*Showcfginfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmlmacservice_SetMacClearAct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Macclearact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmlmacserviceServer).SetMacClearAct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cml.Cmlmacservice/SetMacClearAct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmlmacserviceServer).SetMacClearAct(ctx, req.(*Macclearact))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cmlmacservice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cml.Cmlmacservice",
	HandlerType: (*CmlmacserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetMacIfCfg",
			Handler:    _Cmlmacservice_SetMacIfCfg_Handler,
		},
		{
			MethodName: "DelMacIfCfg",
			Handler:    _Cmlmacservice_DelMacIfCfg_Handler,
		},
		{
			MethodName: "ShowMacIfCfg",
			Handler:    _Cmlmacservice_ShowMacIfCfg_Handler,
		},
		{
			MethodName: "SetMacGlobalCfg",
			Handler:    _Cmlmacservice_SetMacGlobalCfg_Handler,
		},
		{
			MethodName: "DelMacGlobalCfg",
			Handler:    _Cmlmacservice_DelMacGlobalCfg_Handler,
		},
		{
			MethodName: "ShowMacGlobalCfg",
			Handler:    _Cmlmacservice_ShowMacGlobalCfg_Handler,
		},
		{
			MethodName: "SetMacClearAct",
			Handler:    _Cmlmacservice_SetMacClearAct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mac-config.proto",
}
