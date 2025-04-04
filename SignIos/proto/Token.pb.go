// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: Token.proto

package pb

import (
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

type TokenReq_4 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P_1  *TokenReq_4_1 `protobuf:"bytes,1,opt,name=p_1,json=p1,proto3" json:"p_1,omitempty"`
	P_2  string        `protobuf:"bytes,2,opt,name=p_2,json=p2,proto3" json:"p_2,omitempty"`
	P_3  string        `protobuf:"bytes,3,opt,name=p_3,json=p3,proto3" json:"p_3,omitempty"`
	P_4  string        `protobuf:"bytes,4,opt,name=p_4,json=p4,proto3" json:"p_4,omitempty"`
	P_5  uint64        `protobuf:"varint,5,opt,name=p_5,json=p5,proto3" json:"p_5,omitempty"`
	P_6  string        `protobuf:"bytes,6,opt,name=p_6,json=p6,proto3" json:"p_6,omitempty"`
	P_7  string        `protobuf:"bytes,7,opt,name=p_7,json=p7,proto3" json:"p_7,omitempty"`
	P_8  string        `protobuf:"bytes,8,opt,name=p_8,json=p8,proto3" json:"p_8,omitempty"`
	P_9  string        `protobuf:"bytes,9,opt,name=p_9,json=p9,proto3" json:"p_9,omitempty"`
	P_10 uint64        `protobuf:"varint,10,opt,name=p_10,json=p10,proto3" json:"p_10,omitempty"`
	P_11 uint64        `protobuf:"varint,11,opt,name=p_11,json=p11,proto3" json:"p_11,omitempty"`
	P_12 string        `protobuf:"bytes,12,opt,name=p_12,json=p12,proto3" json:"p_12,omitempty"`
	P_16 uint64        `protobuf:"varint,16,opt,name=p_16,json=p16,proto3" json:"p_16,omitempty"`
}

func (x *TokenReq_4) Reset() {
	*x = TokenReq_4{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenReq_4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenReq_4) ProtoMessage() {}

func (x *TokenReq_4) ProtoReflect() protoreflect.Message {
	mi := &file_Token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenReq_4.ProtoReflect.Descriptor instead.
func (*TokenReq_4) Descriptor() ([]byte, []int) {
	return file_Token_proto_rawDescGZIP(), []int{0}
}

func (x *TokenReq_4) GetP_1() *TokenReq_4_1 {
	if x != nil {
		return x.P_1
	}
	return nil
}

func (x *TokenReq_4) GetP_2() string {
	if x != nil {
		return x.P_2
	}
	return ""
}

func (x *TokenReq_4) GetP_3() string {
	if x != nil {
		return x.P_3
	}
	return ""
}

func (x *TokenReq_4) GetP_4() string {
	if x != nil {
		return x.P_4
	}
	return ""
}

func (x *TokenReq_4) GetP_5() uint64 {
	if x != nil {
		return x.P_5
	}
	return 0
}

func (x *TokenReq_4) GetP_6() string {
	if x != nil {
		return x.P_6
	}
	return ""
}

func (x *TokenReq_4) GetP_7() string {
	if x != nil {
		return x.P_7
	}
	return ""
}

func (x *TokenReq_4) GetP_8() string {
	if x != nil {
		return x.P_8
	}
	return ""
}

func (x *TokenReq_4) GetP_9() string {
	if x != nil {
		return x.P_9
	}
	return ""
}

func (x *TokenReq_4) GetP_10() uint64 {
	if x != nil {
		return x.P_10
	}
	return 0
}

func (x *TokenReq_4) GetP_11() uint64 {
	if x != nil {
		return x.P_11
	}
	return 0
}

func (x *TokenReq_4) GetP_12() string {
	if x != nil {
		return x.P_12
	}
	return ""
}

func (x *TokenReq_4) GetP_16() uint64 {
	if x != nil {
		return x.P_16
	}
	return 0
}

type TokenReq_4_1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P_1  string `protobuf:"bytes,1,opt,name=p_1,json=p1,proto3" json:"p_1,omitempty"`
	P_2  string `protobuf:"bytes,2,opt,name=p_2,json=p2,proto3" json:"p_2,omitempty"`
	P_3  string `protobuf:"bytes,3,opt,name=p_3,json=p3,proto3" json:"p_3,omitempty"`
	P_4  string `protobuf:"bytes,4,opt,name=p_4,json=p4,proto3" json:"p_4,omitempty"`
	P_5  string `protobuf:"bytes,5,opt,name=p_5,json=p5,proto3" json:"p_5,omitempty"`
	P_6  string `protobuf:"bytes,6,opt,name=p_6,json=p6,proto3" json:"p_6,omitempty"`
	P_7  string `protobuf:"bytes,7,opt,name=p_7,json=p7,proto3" json:"p_7,omitempty"`
	P_8  uint64 `protobuf:"varint,8,opt,name=p_8,json=p8,proto3" json:"p_8,omitempty"`
	P_9  string `protobuf:"bytes,9,opt,name=p_9,json=p9,proto3" json:"p_9,omitempty"`
	P_10 uint64 `protobuf:"varint,10,opt,name=p_10,json=p10,proto3" json:"p_10,omitempty"`
	P_11 uint64 `protobuf:"varint,11,opt,name=p_11,json=p11,proto3" json:"p_11,omitempty"`
	P_12 string `protobuf:"bytes,12,opt,name=p_12,json=p12,proto3" json:"p_12,omitempty"`
	P_13 uint64 `protobuf:"varint,13,opt,name=p_13,json=p13,proto3" json:"p_13,omitempty"`
	P_14 uint64 `protobuf:"varint,14,opt,name=p_14,json=p14,proto3" json:"p_14,omitempty"`
	P_15 uint64 `protobuf:"varint,15,opt,name=p_15,json=p15,proto3" json:"p_15,omitempty"`
	P_16 uint64 `protobuf:"varint,16,opt,name=p_16,json=p16,proto3" json:"p_16,omitempty"`
	P_17 uint64 `protobuf:"varint,17,opt,name=p_17,json=p17,proto3" json:"p_17,omitempty"`
	P_18 uint64 `protobuf:"varint,18,opt,name=p_18,json=p18,proto3" json:"p_18,omitempty"`
	P_20 string `protobuf:"bytes,20,opt,name=p_20,json=p20,proto3" json:"p_20,omitempty"`
	P_21 string `protobuf:"bytes,21,opt,name=p_21,json=p21,proto3" json:"p_21,omitempty"`
	P_22 string `protobuf:"bytes,22,opt,name=p_22,json=p22,proto3" json:"p_22,omitempty"`
	P_23 string `protobuf:"bytes,23,opt,name=p_23,json=p23,proto3" json:"p_23,omitempty"`
	P_24 string `protobuf:"bytes,24,opt,name=p_24,json=p24,proto3" json:"p_24,omitempty"`
	P_25 uint64 `protobuf:"varint,25,opt,name=p_25,json=p25,proto3" json:"p_25,omitempty"`
	P_26 string `protobuf:"bytes,26,opt,name=p_26,json=p26,proto3" json:"p_26,omitempty"`
	P_27 uint64 `protobuf:"varint,27,opt,name=p_27,json=p27,proto3" json:"p_27,omitempty"`
	P_28 string `protobuf:"bytes,28,opt,name=p_28,json=p28,proto3" json:"p_28,omitempty"`
	P_29 uint64 `protobuf:"varint,29,opt,name=p_29,json=p29,proto3" json:"p_29,omitempty"`
	P_30 uint64 `protobuf:"varint,30,opt,name=p_30,json=p30,proto3" json:"p_30,omitempty"`
	P_31 string `protobuf:"bytes,31,opt,name=p_31,json=p31,proto3" json:"p_31,omitempty"`
	P_32 string `protobuf:"bytes,32,opt,name=p_32,json=p32,proto3" json:"p_32,omitempty"`
	P_33 string `protobuf:"bytes,33,opt,name=p_33,json=p33,proto3" json:"p_33,omitempty"`
	P_34 string `protobuf:"bytes,34,opt,name=p_34,json=p34,proto3" json:"p_34,omitempty"`
	P_35 string `protobuf:"bytes,35,opt,name=p_35,json=p35,proto3" json:"p_35,omitempty"`
	P_38 string `protobuf:"bytes,38,opt,name=p_38,json=p38,proto3" json:"p_38,omitempty"`
	P_39 string `protobuf:"bytes,39,opt,name=p_39,json=p39,proto3" json:"p_39,omitempty"`
	P_40 uint64 `protobuf:"varint,40,opt,name=p_40,json=p40,proto3" json:"p_40,omitempty"`
	P_42 uint64 `protobuf:"varint,42,opt,name=p_42,json=p42,proto3" json:"p_42,omitempty"`
	P_43 uint64 `protobuf:"varint,43,opt,name=p_43,json=p43,proto3" json:"p_43,omitempty"`
	P_44 string `protobuf:"bytes,44,opt,name=p_44,json=p44,proto3" json:"p_44,omitempty"`
	P_45 uint64 `protobuf:"varint,45,opt,name=p_45,json=p45,proto3" json:"p_45,omitempty"`
	P_46 uint64 `protobuf:"varint,46,opt,name=p_46,json=p46,proto3" json:"p_46,omitempty"`
	P_47 string `protobuf:"bytes,47,opt,name=p_47,json=p47,proto3" json:"p_47,omitempty"`
	P_48 uint64 `protobuf:"varint,48,opt,name=p_48,json=p48,proto3" json:"p_48,omitempty"`
	P_49 uint64 `protobuf:"varint,49,opt,name=p_49,json=p49,proto3" json:"p_49,omitempty"`
	P_50 string `protobuf:"bytes,50,opt,name=p_50,json=p50,proto3" json:"p_50,omitempty"`
	P_51 uint64 `protobuf:"varint,51,opt,name=p_51,json=p51,proto3" json:"p_51,omitempty"`
	P_52 uint64 `protobuf:"varint,52,opt,name=p_52,json=p52,proto3" json:"p_52,omitempty"`
	P_53 string `protobuf:"bytes,53,opt,name=p_53,json=p53,proto3" json:"p_53,omitempty"`
	P_54 string `protobuf:"bytes,54,opt,name=p_54,json=p54,proto3" json:"p_54,omitempty"`
	P_55 string `protobuf:"bytes,55,opt,name=p_55,json=p55,proto3" json:"p_55,omitempty"`
	P_56 uint64 `protobuf:"varint,56,opt,name=p_56,json=p56,proto3" json:"p_56,omitempty"`
}

func (x *TokenReq_4_1) Reset() {
	*x = TokenReq_4_1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenReq_4_1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenReq_4_1) ProtoMessage() {}

func (x *TokenReq_4_1) ProtoReflect() protoreflect.Message {
	mi := &file_Token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenReq_4_1.ProtoReflect.Descriptor instead.
func (*TokenReq_4_1) Descriptor() ([]byte, []int) {
	return file_Token_proto_rawDescGZIP(), []int{1}
}

func (x *TokenReq_4_1) GetP_1() string {
	if x != nil {
		return x.P_1
	}
	return ""
}

func (x *TokenReq_4_1) GetP_2() string {
	if x != nil {
		return x.P_2
	}
	return ""
}

func (x *TokenReq_4_1) GetP_3() string {
	if x != nil {
		return x.P_3
	}
	return ""
}

func (x *TokenReq_4_1) GetP_4() string {
	if x != nil {
		return x.P_4
	}
	return ""
}

func (x *TokenReq_4_1) GetP_5() string {
	if x != nil {
		return x.P_5
	}
	return ""
}

func (x *TokenReq_4_1) GetP_6() string {
	if x != nil {
		return x.P_6
	}
	return ""
}

func (x *TokenReq_4_1) GetP_7() string {
	if x != nil {
		return x.P_7
	}
	return ""
}

func (x *TokenReq_4_1) GetP_8() uint64 {
	if x != nil {
		return x.P_8
	}
	return 0
}

func (x *TokenReq_4_1) GetP_9() string {
	if x != nil {
		return x.P_9
	}
	return ""
}

func (x *TokenReq_4_1) GetP_10() uint64 {
	if x != nil {
		return x.P_10
	}
	return 0
}

func (x *TokenReq_4_1) GetP_11() uint64 {
	if x != nil {
		return x.P_11
	}
	return 0
}

func (x *TokenReq_4_1) GetP_12() string {
	if x != nil {
		return x.P_12
	}
	return ""
}

func (x *TokenReq_4_1) GetP_13() uint64 {
	if x != nil {
		return x.P_13
	}
	return 0
}

func (x *TokenReq_4_1) GetP_14() uint64 {
	if x != nil {
		return x.P_14
	}
	return 0
}

func (x *TokenReq_4_1) GetP_15() uint64 {
	if x != nil {
		return x.P_15
	}
	return 0
}

func (x *TokenReq_4_1) GetP_16() uint64 {
	if x != nil {
		return x.P_16
	}
	return 0
}

func (x *TokenReq_4_1) GetP_17() uint64 {
	if x != nil {
		return x.P_17
	}
	return 0
}

func (x *TokenReq_4_1) GetP_18() uint64 {
	if x != nil {
		return x.P_18
	}
	return 0
}

func (x *TokenReq_4_1) GetP_20() string {
	if x != nil {
		return x.P_20
	}
	return ""
}

func (x *TokenReq_4_1) GetP_21() string {
	if x != nil {
		return x.P_21
	}
	return ""
}

func (x *TokenReq_4_1) GetP_22() string {
	if x != nil {
		return x.P_22
	}
	return ""
}

func (x *TokenReq_4_1) GetP_23() string {
	if x != nil {
		return x.P_23
	}
	return ""
}

func (x *TokenReq_4_1) GetP_24() string {
	if x != nil {
		return x.P_24
	}
	return ""
}

func (x *TokenReq_4_1) GetP_25() uint64 {
	if x != nil {
		return x.P_25
	}
	return 0
}

func (x *TokenReq_4_1) GetP_26() string {
	if x != nil {
		return x.P_26
	}
	return ""
}

func (x *TokenReq_4_1) GetP_27() uint64 {
	if x != nil {
		return x.P_27
	}
	return 0
}

func (x *TokenReq_4_1) GetP_28() string {
	if x != nil {
		return x.P_28
	}
	return ""
}

func (x *TokenReq_4_1) GetP_29() uint64 {
	if x != nil {
		return x.P_29
	}
	return 0
}

func (x *TokenReq_4_1) GetP_30() uint64 {
	if x != nil {
		return x.P_30
	}
	return 0
}

func (x *TokenReq_4_1) GetP_31() string {
	if x != nil {
		return x.P_31
	}
	return ""
}

func (x *TokenReq_4_1) GetP_32() string {
	if x != nil {
		return x.P_32
	}
	return ""
}

func (x *TokenReq_4_1) GetP_33() string {
	if x != nil {
		return x.P_33
	}
	return ""
}

func (x *TokenReq_4_1) GetP_34() string {
	if x != nil {
		return x.P_34
	}
	return ""
}

func (x *TokenReq_4_1) GetP_35() string {
	if x != nil {
		return x.P_35
	}
	return ""
}

func (x *TokenReq_4_1) GetP_38() string {
	if x != nil {
		return x.P_38
	}
	return ""
}

func (x *TokenReq_4_1) GetP_39() string {
	if x != nil {
		return x.P_39
	}
	return ""
}

func (x *TokenReq_4_1) GetP_40() uint64 {
	if x != nil {
		return x.P_40
	}
	return 0
}

func (x *TokenReq_4_1) GetP_42() uint64 {
	if x != nil {
		return x.P_42
	}
	return 0
}

func (x *TokenReq_4_1) GetP_43() uint64 {
	if x != nil {
		return x.P_43
	}
	return 0
}

func (x *TokenReq_4_1) GetP_44() string {
	if x != nil {
		return x.P_44
	}
	return ""
}

func (x *TokenReq_4_1) GetP_45() uint64 {
	if x != nil {
		return x.P_45
	}
	return 0
}

func (x *TokenReq_4_1) GetP_46() uint64 {
	if x != nil {
		return x.P_46
	}
	return 0
}

func (x *TokenReq_4_1) GetP_47() string {
	if x != nil {
		return x.P_47
	}
	return ""
}

func (x *TokenReq_4_1) GetP_48() uint64 {
	if x != nil {
		return x.P_48
	}
	return 0
}

func (x *TokenReq_4_1) GetP_49() uint64 {
	if x != nil {
		return x.P_49
	}
	return 0
}

func (x *TokenReq_4_1) GetP_50() string {
	if x != nil {
		return x.P_50
	}
	return ""
}

func (x *TokenReq_4_1) GetP_51() uint64 {
	if x != nil {
		return x.P_51
	}
	return 0
}

func (x *TokenReq_4_1) GetP_52() uint64 {
	if x != nil {
		return x.P_52
	}
	return 0
}

func (x *TokenReq_4_1) GetP_53() string {
	if x != nil {
		return x.P_53
	}
	return ""
}

func (x *TokenReq_4_1) GetP_54() string {
	if x != nil {
		return x.P_54
	}
	return ""
}

func (x *TokenReq_4_1) GetP_55() string {
	if x != nil {
		return x.P_55
	}
	return ""
}

func (x *TokenReq_4_1) GetP_56() uint64 {
	if x != nil {
		return x.P_56
	}
	return 0
}

type TokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P_1 uint64 `protobuf:"varint,1,opt,name=p_1,json=p1,proto3" json:"p_1,omitempty"`
	P_2 uint64 `protobuf:"varint,2,opt,name=p_2,json=p2,proto3" json:"p_2,omitempty"`
	P_3 uint64 `protobuf:"varint,3,opt,name=p_3,json=p3,proto3" json:"p_3,omitempty"`
	P_4 []byte `protobuf:"bytes,4,opt,name=p_4,json=p4,proto3" json:"p_4,omitempty"`
	P_5 uint64 `protobuf:"varint,5,opt,name=p_5,json=p5,proto3" json:"p_5,omitempty"`
}

func (x *TokenReq) Reset() {
	*x = TokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenReq) ProtoMessage() {}

func (x *TokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_Token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenReq.ProtoReflect.Descriptor instead.
func (*TokenReq) Descriptor() ([]byte, []int) {
	return file_Token_proto_rawDescGZIP(), []int{2}
}

func (x *TokenReq) GetP_1() uint64 {
	if x != nil {
		return x.P_1
	}
	return 0
}

func (x *TokenReq) GetP_2() uint64 {
	if x != nil {
		return x.P_2
	}
	return 0
}

func (x *TokenReq) GetP_3() uint64 {
	if x != nil {
		return x.P_3
	}
	return 0
}

func (x *TokenReq) GetP_4() []byte {
	if x != nil {
		return x.P_4
	}
	return nil
}

func (x *TokenReq) GetP_5() uint64 {
	if x != nil {
		return x.P_5
	}
	return 0
}

type TokenReq_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P_1 uint64      `protobuf:"varint,1,opt,name=p_1,json=p1,proto3" json:"p_1,omitempty"`
	P_2 uint64      `protobuf:"varint,2,opt,name=p_2,json=p2,proto3" json:"p_2,omitempty"`
	P_3 uint64      `protobuf:"varint,3,opt,name=p_3,json=p3,proto3" json:"p_3,omitempty"`
	P_4 *TokenReq_4 `protobuf:"bytes,4,opt,name=p_4,json=p4,proto3" json:"p_4,omitempty"`
	P_5 uint64      `protobuf:"varint,5,opt,name=p_5,json=p5,proto3" json:"p_5,omitempty"`
}

func (x *TokenReq_Res) Reset() {
	*x = TokenReq_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenReq_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenReq_Res) ProtoMessage() {}

func (x *TokenReq_Res) ProtoReflect() protoreflect.Message {
	mi := &file_Token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenReq_Res.ProtoReflect.Descriptor instead.
func (*TokenReq_Res) Descriptor() ([]byte, []int) {
	return file_Token_proto_rawDescGZIP(), []int{3}
}

func (x *TokenReq_Res) GetP_1() uint64 {
	if x != nil {
		return x.P_1
	}
	return 0
}

func (x *TokenReq_Res) GetP_2() uint64 {
	if x != nil {
		return x.P_2
	}
	return 0
}

func (x *TokenReq_Res) GetP_3() uint64 {
	if x != nil {
		return x.P_3
	}
	return 0
}

func (x *TokenReq_Res) GetP_4() *TokenReq_4 {
	if x != nil {
		return x.P_4
	}
	return nil
}

func (x *TokenReq_Res) GetP_5() uint64 {
	if x != nil {
		return x.P_5
	}
	return 0
}

type TokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P_1 uint64 `protobuf:"varint,1,opt,name=p_1,json=p1,proto3" json:"p_1,omitempty"`
	P_2 uint64 `protobuf:"varint,2,opt,name=p_2,json=p2,proto3" json:"p_2,omitempty"`
	P_5 uint64 `protobuf:"varint,5,opt,name=p_5,json=p5,proto3" json:"p_5,omitempty"`
	P_6 []byte `protobuf:"bytes,6,opt,name=p_6,json=p6,proto3" json:"p_6,omitempty"`
}

func (x *TokenRes) Reset() {
	*x = TokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRes) ProtoMessage() {}

func (x *TokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_Token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRes.ProtoReflect.Descriptor instead.
func (*TokenRes) Descriptor() ([]byte, []int) {
	return file_Token_proto_rawDescGZIP(), []int{4}
}

func (x *TokenRes) GetP_1() uint64 {
	if x != nil {
		return x.P_1
	}
	return 0
}

func (x *TokenRes) GetP_2() uint64 {
	if x != nil {
		return x.P_2
	}
	return 0
}

func (x *TokenRes) GetP_5() uint64 {
	if x != nil {
		return x.P_5
	}
	return 0
}

func (x *TokenRes) GetP_6() []byte {
	if x != nil {
		return x.P_6
	}
	return nil
}

type TokenRes_6 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P_1 string `protobuf:"bytes,1,opt,name=p_1,json=p1,proto3" json:"p_1,omitempty"`
	P_2 uint64 `protobuf:"varint,2,opt,name=p_2,json=p2,proto3" json:"p_2,omitempty"`
}

func (x *TokenRes_6) Reset() {
	*x = TokenRes_6{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Token_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRes_6) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRes_6) ProtoMessage() {}

func (x *TokenRes_6) ProtoReflect() protoreflect.Message {
	mi := &file_Token_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRes_6.ProtoReflect.Descriptor instead.
func (*TokenRes_6) Descriptor() ([]byte, []int) {
	return file_Token_proto_rawDescGZIP(), []int{5}
}

func (x *TokenRes_6) GetP_1() string {
	if x != nil {
		return x.P_1
	}
	return ""
}

func (x *TokenRes_6) GetP_2() uint64 {
	if x != nil {
		return x.P_2
	}
	return 0
}

type TokenRes_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P_1 uint64      `protobuf:"varint,1,opt,name=p_1,json=p1,proto3" json:"p_1,omitempty"`
	P_2 uint64      `protobuf:"varint,2,opt,name=p_2,json=p2,proto3" json:"p_2,omitempty"`
	P_5 uint64      `protobuf:"varint,5,opt,name=p_5,json=p5,proto3" json:"p_5,omitempty"`
	P_6 *TokenRes_6 `protobuf:"bytes,6,opt,name=p_6,json=p6,proto3" json:"p_6,omitempty"`
}

func (x *TokenRes_Res) Reset() {
	*x = TokenRes_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Token_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRes_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRes_Res) ProtoMessage() {}

func (x *TokenRes_Res) ProtoReflect() protoreflect.Message {
	mi := &file_Token_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRes_Res.ProtoReflect.Descriptor instead.
func (*TokenRes_Res) Descriptor() ([]byte, []int) {
	return file_Token_proto_rawDescGZIP(), []int{6}
}

func (x *TokenRes_Res) GetP_1() uint64 {
	if x != nil {
		return x.P_1
	}
	return 0
}

func (x *TokenRes_Res) GetP_2() uint64 {
	if x != nil {
		return x.P_2
	}
	return 0
}

func (x *TokenRes_Res) GetP_5() uint64 {
	if x != nil {
		return x.P_5
	}
	return 0
}

func (x *TokenRes_Res) GetP_6() *TokenRes_6 {
	if x != nil {
		return x.P_6
	}
	return nil
}

var File_Token_proto protoreflect.FileDescriptor

var file_Token_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x83, 0x02, 0x0a, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x5f, 0x34,
	0x12, 0x21, 0x0a, 0x03, 0x70, 0x5f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x62, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x5f, 0x34, 0x5f, 0x31, 0x52,
	0x02, 0x70, 0x31, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x70, 0x32, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x70, 0x33, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x34, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x70, 0x34, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x35, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x70, 0x35, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x36, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x36, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x37, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x37, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x38,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x38, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f,
	0x39, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x39, 0x12, 0x11, 0x0a, 0x04, 0x70,
	0x5f, 0x31, 0x30, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x31, 0x30, 0x12, 0x11,
	0x0a, 0x04, 0x70, 0x5f, 0x31, 0x31, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x31,
	0x31, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x31, 0x32, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x70, 0x31, 0x32, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x31, 0x36, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x70, 0x31, 0x36, 0x22, 0xd8, 0x07, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x5f, 0x34, 0x5f, 0x31, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x31, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x31, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x32,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x32, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f,
	0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x33, 0x12, 0x0f, 0x0a, 0x03, 0x70,
	0x5f, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x34, 0x12, 0x0f, 0x0a, 0x03,
	0x70, 0x5f, 0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x35, 0x12, 0x0f, 0x0a,
	0x03, 0x70, 0x5f, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x36, 0x12, 0x0f,
	0x0a, 0x03, 0x70, 0x5f, 0x37, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x37, 0x12,
	0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x38, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x70, 0x38,
	0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x39, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70,
	0x39, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x31, 0x30, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x70, 0x31, 0x30, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x31, 0x31, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x70, 0x31, 0x31, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x31, 0x32, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x31, 0x32, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f,
	0x31, 0x33, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x31, 0x33, 0x12, 0x11, 0x0a,
	0x04, 0x70, 0x5f, 0x31, 0x34, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x31, 0x34,
	0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x31, 0x35, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x70, 0x31, 0x35, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x31, 0x36, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x70, 0x31, 0x36, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x31, 0x37, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x31, 0x37, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x31,
	0x38, 0x18, 0x12, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x31, 0x38, 0x12, 0x11, 0x0a, 0x04,
	0x70, 0x5f, 0x32, 0x30, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x32, 0x30, 0x12,
	0x11, 0x0a, 0x04, 0x70, 0x5f, 0x32, 0x31, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70,
	0x32, 0x31, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x32, 0x32, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x70, 0x32, 0x32, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x32, 0x33, 0x18, 0x17, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x32, 0x33, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x32, 0x34,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x32, 0x34, 0x12, 0x11, 0x0a, 0x04, 0x70,
	0x5f, 0x32, 0x35, 0x18, 0x19, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x32, 0x35, 0x12, 0x11,
	0x0a, 0x04, 0x70, 0x5f, 0x32, 0x36, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x32,
	0x36, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x32, 0x37, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x70, 0x32, 0x37, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x32, 0x38, 0x18, 0x1c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x70, 0x32, 0x38, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x32, 0x39, 0x18,
	0x1d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x32, 0x39, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f,
	0x33, 0x30, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x33, 0x30, 0x12, 0x11, 0x0a,
	0x04, 0x70, 0x5f, 0x33, 0x31, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x33, 0x31,
	0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x33, 0x32, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x70, 0x33, 0x32, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x33, 0x33, 0x18, 0x21, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x70, 0x33, 0x33, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x33, 0x34, 0x18, 0x22,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x33, 0x34, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x33,
	0x35, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x33, 0x35, 0x12, 0x11, 0x0a, 0x04,
	0x70, 0x5f, 0x33, 0x38, 0x18, 0x26, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x33, 0x38, 0x12,
	0x11, 0x0a, 0x04, 0x70, 0x5f, 0x33, 0x39, 0x18, 0x27, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70,
	0x33, 0x39, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x34, 0x30, 0x18, 0x28, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x70, 0x34, 0x30, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x34, 0x32, 0x18, 0x2a, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x34, 0x32, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x34, 0x33,
	0x18, 0x2b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x34, 0x33, 0x12, 0x11, 0x0a, 0x04, 0x70,
	0x5f, 0x34, 0x34, 0x18, 0x2c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x34, 0x34, 0x12, 0x11,
	0x0a, 0x04, 0x70, 0x5f, 0x34, 0x35, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x34,
	0x35, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x34, 0x36, 0x18, 0x2e, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x70, 0x34, 0x36, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x34, 0x37, 0x18, 0x2f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x70, 0x34, 0x37, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x34, 0x38, 0x18,
	0x30, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x34, 0x38, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f,
	0x34, 0x39, 0x18, 0x31, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x34, 0x39, 0x12, 0x11, 0x0a,
	0x04, 0x70, 0x5f, 0x35, 0x30, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x35, 0x30,
	0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x35, 0x31, 0x18, 0x33, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x70, 0x35, 0x31, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x35, 0x32, 0x18, 0x34, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x70, 0x35, 0x32, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x35, 0x33, 0x18, 0x35,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x35, 0x33, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x35,
	0x34, 0x18, 0x36, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x35, 0x34, 0x12, 0x11, 0x0a, 0x04,
	0x70, 0x5f, 0x35, 0x35, 0x18, 0x37, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x35, 0x35, 0x12,
	0x11, 0x0a, 0x04, 0x70, 0x5f, 0x35, 0x36, 0x18, 0x38, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70,
	0x35, 0x36, 0x22, 0x5f, 0x0a, 0x08, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x0f,
	0x0a, 0x03, 0x70, 0x5f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x70, 0x31, 0x12,
	0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x70, 0x32,
	0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x70,
	0x33, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02,
	0x70, 0x34, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x70, 0x35, 0x22, 0x73, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x5f,
	0x52, 0x65, 0x73, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x70, 0x31, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x70, 0x32, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x33, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x70, 0x33, 0x12, 0x1f, 0x0a, 0x03, 0x70, 0x5f, 0x34, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x5f, 0x34, 0x52, 0x02, 0x70, 0x34, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x35, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x70, 0x35, 0x22, 0x4e, 0x0a, 0x08, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x70, 0x31, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x32, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x70, 0x32, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x35, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x70, 0x35, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x36, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x70, 0x36, 0x22, 0x2e, 0x0a, 0x0a, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x5f, 0x36, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x31, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x32, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x70, 0x32, 0x22, 0x62, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x31, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x70, 0x31, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f, 0x32,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x70, 0x32, 0x12, 0x0f, 0x0a, 0x03, 0x70, 0x5f,
	0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x70, 0x35, 0x12, 0x1f, 0x0a, 0x03, 0x70,
	0x5f, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x5f, 0x36, 0x52, 0x02, 0x70, 0x36, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Token_proto_rawDescOnce sync.Once
	file_Token_proto_rawDescData = file_Token_proto_rawDesc
)

func file_Token_proto_rawDescGZIP() []byte {
	file_Token_proto_rawDescOnce.Do(func() {
		file_Token_proto_rawDescData = protoimpl.X.CompressGZIP(file_Token_proto_rawDescData)
	})
	return file_Token_proto_rawDescData
}

var file_Token_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_Token_proto_goTypes = []interface{}{
	(*TokenReq_4)(nil),   // 0: pb.TokenReq_4
	(*TokenReq_4_1)(nil), // 1: pb.TokenReq_4_1
	(*TokenReq)(nil),     // 2: pb.TokenReq
	(*TokenReq_Res)(nil), // 3: pb.TokenReq_Res
	(*TokenRes)(nil),     // 4: pb.TokenRes
	(*TokenRes_6)(nil),   // 5: pb.TokenRes_6
	(*TokenRes_Res)(nil), // 6: pb.TokenRes_Res
}
var file_Token_proto_depIdxs = []int32{
	1, // 0: pb.TokenReq_4.p_1:type_name -> pb.TokenReq_4_1
	0, // 1: pb.TokenReq_Res.p_4:type_name -> pb.TokenReq_4
	5, // 2: pb.TokenRes_Res.p_6:type_name -> pb.TokenRes_6
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Token_proto_init() }
func file_Token_proto_init() {
	if File_Token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenReq_4); i {
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
		file_Token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenReq_4_1); i {
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
		file_Token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenReq); i {
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
		file_Token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenReq_Res); i {
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
		file_Token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRes); i {
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
		file_Token_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRes_6); i {
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
		file_Token_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRes_Res); i {
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
			RawDescriptor: file_Token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Token_proto_goTypes,
		DependencyIndexes: file_Token_proto_depIdxs,
		MessageInfos:      file_Token_proto_msgTypes,
	}.Build()
	File_Token_proto = out.File
	file_Token_proto_rawDesc = nil
	file_Token_proto_goTypes = nil
	file_Token_proto_depIdxs = nil
}
