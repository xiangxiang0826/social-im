// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: giftRpc.proto

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

type GiftSendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	GiftId   int64  `protobuf:"varint,2,opt,name=giftId,proto3" json:"giftId,omitempty"`
	GiftNum  int64  `protobuf:"varint,3,opt,name=gift_num,json=giftNum,proto3" json:"gift_num,omitempty"`
	SendTo   string `protobuf:"bytes,4,opt,name=send_to,json=sendTo,proto3" json:"send_to,omitempty"`
	RoomMark string `protobuf:"bytes,5,opt,name=room_mark,json=roomMark,proto3" json:"room_mark,omitempty"`
}

func (x *GiftSendReq) Reset() {
	*x = GiftSendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_giftRpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftSendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftSendReq) ProtoMessage() {}

func (x *GiftSendReq) ProtoReflect() protoreflect.Message {
	mi := &file_giftRpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftSendReq.ProtoReflect.Descriptor instead.
func (*GiftSendReq) Descriptor() ([]byte, []int) {
	return file_giftRpc_proto_rawDescGZIP(), []int{0}
}

func (x *GiftSendReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GiftSendReq) GetGiftId() int64 {
	if x != nil {
		return x.GiftId
	}
	return 0
}

func (x *GiftSendReq) GetGiftNum() int64 {
	if x != nil {
		return x.GiftNum
	}
	return 0
}

func (x *GiftSendReq) GetSendTo() string {
	if x != nil {
		return x.SendTo
	}
	return ""
}

func (x *GiftSendReq) GetRoomMark() string {
	if x != nil {
		return x.RoomMark
	}
	return ""
}

type GiftSendResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iret int64  `protobuf:"varint,1,opt,name=Iret,proto3" json:"Iret,omitempty"`
	Smsg string `protobuf:"bytes,2,opt,name=smsg,proto3" json:"smsg,omitempty"`
}

func (x *GiftSendResp) Reset() {
	*x = GiftSendResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_giftRpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftSendResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftSendResp) ProtoMessage() {}

func (x *GiftSendResp) ProtoReflect() protoreflect.Message {
	mi := &file_giftRpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftSendResp.ProtoReflect.Descriptor instead.
func (*GiftSendResp) Descriptor() ([]byte, []int) {
	return file_giftRpc_proto_rawDescGZIP(), []int{1}
}

func (x *GiftSendResp) GetIret() int64 {
	if x != nil {
		return x.Iret
	}
	return 0
}

func (x *GiftSendResp) GetSmsg() string {
	if x != nil {
		return x.Smsg
	}
	return ""
}

type GetItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GiftId int64 `protobuf:"varint,1,opt,name=GiftId,proto3" json:"GiftId,omitempty"`
}

func (x *GetItemReq) Reset() {
	*x = GetItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_giftRpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemReq) ProtoMessage() {}

func (x *GetItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_giftRpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemReq.ProtoReflect.Descriptor instead.
func (*GetItemReq) Descriptor() ([]byte, []int) {
	return file_giftRpc_proto_rawDescGZIP(), []int{2}
}

func (x *GetItemReq) GetGiftId() int64 {
	if x != nil {
		return x.GiftId
	}
	return 0
}

type GetItemResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iret      int64  `protobuf:"varint,1,opt,name=Iret,proto3" json:"Iret,omitempty"`
	Smsg      string `protobuf:"bytes,2,opt,name=Smsg,proto3" json:"Smsg,omitempty"`
	Id        int64  `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`
	GiftName  string `protobuf:"bytes,4,opt,name=GiftName,proto3" json:"GiftName,omitempty"`
	ImgUrl    string `protobuf:"bytes,5,opt,name=ImgUrl,proto3" json:"ImgUrl,omitempty"`
	CreatedAt int64  `protobuf:"varint,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Currency  int64  `protobuf:"varint,7,opt,name=Currency,proto3" json:"Currency,omitempty"`
	ChargeNum int64  `protobuf:"varint,8,opt,name=ChargeNum,proto3" json:"ChargeNum,omitempty"`
}

func (x *GetItemResp) Reset() {
	*x = GetItemResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_giftRpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemResp) ProtoMessage() {}

func (x *GetItemResp) ProtoReflect() protoreflect.Message {
	mi := &file_giftRpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemResp.ProtoReflect.Descriptor instead.
func (*GetItemResp) Descriptor() ([]byte, []int) {
	return file_giftRpc_proto_rawDescGZIP(), []int{3}
}

func (x *GetItemResp) GetIret() int64 {
	if x != nil {
		return x.Iret
	}
	return 0
}

func (x *GetItemResp) GetSmsg() string {
	if x != nil {
		return x.Smsg
	}
	return ""
}

func (x *GetItemResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetItemResp) GetGiftName() string {
	if x != nil {
		return x.GiftName
	}
	return ""
}

func (x *GetItemResp) GetImgUrl() string {
	if x != nil {
		return x.ImgUrl
	}
	return ""
}

func (x *GetItemResp) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *GetItemResp) GetCurrency() int64 {
	if x != nil {
		return x.Currency
	}
	return 0
}

func (x *GetItemResp) GetChargeNum() int64 {
	if x != nil {
		return x.ChargeNum
	}
	return 0
}

type GiftUpdateBagReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	GiftId  int64 `protobuf:"varint,2,opt,name=giftId,proto3" json:"giftId,omitempty"`
	GiftNum int64 `protobuf:"varint,3,opt,name=gift_num,json=giftNum,proto3" json:"gift_num,omitempty"`
	SendTo  int64 `protobuf:"varint,4,opt,name=send_to,json=sendTo,proto3" json:"send_to,omitempty"`
}

func (x *GiftUpdateBagReq) Reset() {
	*x = GiftUpdateBagReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_giftRpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftUpdateBagReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftUpdateBagReq) ProtoMessage() {}

func (x *GiftUpdateBagReq) ProtoReflect() protoreflect.Message {
	mi := &file_giftRpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftUpdateBagReq.ProtoReflect.Descriptor instead.
func (*GiftUpdateBagReq) Descriptor() ([]byte, []int) {
	return file_giftRpc_proto_rawDescGZIP(), []int{4}
}

func (x *GiftUpdateBagReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GiftUpdateBagReq) GetGiftId() int64 {
	if x != nil {
		return x.GiftId
	}
	return 0
}

func (x *GiftUpdateBagReq) GetGiftNum() int64 {
	if x != nil {
		return x.GiftNum
	}
	return 0
}

func (x *GiftUpdateBagReq) GetSendTo() int64 {
	if x != nil {
		return x.SendTo
	}
	return 0
}

type GiftUpdateBagResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iret int64  `protobuf:"varint,1,opt,name=Iret,proto3" json:"Iret,omitempty"`
	Smsg string `protobuf:"bytes,2,opt,name=smsg,proto3" json:"smsg,omitempty"`
}

func (x *GiftUpdateBagResp) Reset() {
	*x = GiftUpdateBagResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_giftRpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftUpdateBagResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftUpdateBagResp) ProtoMessage() {}

func (x *GiftUpdateBagResp) ProtoReflect() protoreflect.Message {
	mi := &file_giftRpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftUpdateBagResp.ProtoReflect.Descriptor instead.
func (*GiftUpdateBagResp) Descriptor() ([]byte, []int) {
	return file_giftRpc_proto_rawDescGZIP(), []int{5}
}

func (x *GiftUpdateBagResp) GetIret() int64 {
	if x != nil {
		return x.Iret
	}
	return 0
}

func (x *GiftUpdateBagResp) GetSmsg() string {
	if x != nil {
		return x.Smsg
	}
	return ""
}

type GiftUpdateAttrReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	GiftId  int64 `protobuf:"varint,2,opt,name=giftId,proto3" json:"giftId,omitempty"`
	GiftNum int64 `protobuf:"varint,3,opt,name=gift_num,json=giftNum,proto3" json:"gift_num,omitempty"`
	SendTo  int64 `protobuf:"varint,4,opt,name=send_to,json=sendTo,proto3" json:"send_to,omitempty"`
}

func (x *GiftUpdateAttrReq) Reset() {
	*x = GiftUpdateAttrReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_giftRpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftUpdateAttrReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftUpdateAttrReq) ProtoMessage() {}

func (x *GiftUpdateAttrReq) ProtoReflect() protoreflect.Message {
	mi := &file_giftRpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftUpdateAttrReq.ProtoReflect.Descriptor instead.
func (*GiftUpdateAttrReq) Descriptor() ([]byte, []int) {
	return file_giftRpc_proto_rawDescGZIP(), []int{6}
}

func (x *GiftUpdateAttrReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GiftUpdateAttrReq) GetGiftId() int64 {
	if x != nil {
		return x.GiftId
	}
	return 0
}

func (x *GiftUpdateAttrReq) GetGiftNum() int64 {
	if x != nil {
		return x.GiftNum
	}
	return 0
}

func (x *GiftUpdateAttrReq) GetSendTo() int64 {
	if x != nil {
		return x.SendTo
	}
	return 0
}

type GiftUpdateAttrResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iret int64  `protobuf:"varint,1,opt,name=Iret,proto3" json:"Iret,omitempty"`
	Smsg string `protobuf:"bytes,2,opt,name=smsg,proto3" json:"smsg,omitempty"`
}

func (x *GiftUpdateAttrResp) Reset() {
	*x = GiftUpdateAttrResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_giftRpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftUpdateAttrResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftUpdateAttrResp) ProtoMessage() {}

func (x *GiftUpdateAttrResp) ProtoReflect() protoreflect.Message {
	mi := &file_giftRpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftUpdateAttrResp.ProtoReflect.Descriptor instead.
func (*GiftUpdateAttrResp) Descriptor() ([]byte, []int) {
	return file_giftRpc_proto_rawDescGZIP(), []int{7}
}

func (x *GiftUpdateAttrResp) GetIret() int64 {
	if x != nil {
		return x.Iret
	}
	return 0
}

func (x *GiftUpdateAttrResp) GetSmsg() string {
	if x != nil {
		return x.Smsg
	}
	return ""
}

var File_giftRpc_proto protoreflect.FileDescriptor

var file_giftRpc_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x69, 0x66, 0x74, 0x52, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x88, 0x01, 0x0a, 0x0b, 0x47, 0x69, 0x66, 0x74, 0x53, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x69, 0x66, 0x74, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x67, 0x69, 0x66, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x67, 0x69, 0x66, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x54,
	0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4d, 0x61, 0x72, 0x6b, 0x22, 0x36,
	0x0a, 0x0c, 0x47, 0x69, 0x66, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x49, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x49, 0x72,
	0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x6d, 0x73, 0x67, 0x22, 0x24, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x69, 0x66, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x47, 0x69, 0x66, 0x74, 0x49, 0x64, 0x22, 0xd1, 0x01, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x49, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x49, 0x72, 0x65, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x53, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x53, 0x6d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x69, 0x66, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x47, 0x69, 0x66, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x49, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x49, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x4e, 0x75, 0x6d,
	0x22, 0x70, 0x0a, 0x10, 0x47, 0x69, 0x66, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61,
	0x67, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x69, 0x66, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x67, 0x69, 0x66, 0x74, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x67, 0x69, 0x66, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x54, 0x6f, 0x22, 0x3b, 0x0a, 0x11, 0x47, 0x69, 0x66, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x72, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x49, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6d, 0x73, 0x67, 0x22,
	0x71, 0x0a, 0x11, 0x47, 0x69, 0x66, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x69, 0x66, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x67, 0x69, 0x66, 0x74, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x67, 0x69, 0x66, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x54, 0x6f, 0x22, 0x3c, 0x0a, 0x12, 0x47, 0x69, 0x66, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x72, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x49, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6d, 0x73, 0x67,
	0x32, 0xab, 0x01, 0x0a, 0x07, 0x67, 0x69, 0x66, 0x74, 0x52, 0x70, 0x63, 0x12, 0x29, 0x0a, 0x04,
	0x73, 0x65, 0x6e, 0x64, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x69, 0x66, 0x74, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x69, 0x66, 0x74, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x67, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x69, 0x66, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x69, 0x66, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x3b, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x12,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x69, 0x66, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x74, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x69, 0x66, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_giftRpc_proto_rawDescOnce sync.Once
	file_giftRpc_proto_rawDescData = file_giftRpc_proto_rawDesc
)

func file_giftRpc_proto_rawDescGZIP() []byte {
	file_giftRpc_proto_rawDescOnce.Do(func() {
		file_giftRpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_giftRpc_proto_rawDescData)
	})
	return file_giftRpc_proto_rawDescData
}

var file_giftRpc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_giftRpc_proto_goTypes = []interface{}{
	(*GiftSendReq)(nil),        // 0: pb.GiftSendReq
	(*GiftSendResp)(nil),       // 1: pb.GiftSendResp
	(*GetItemReq)(nil),         // 2: pb.GetItemReq
	(*GetItemResp)(nil),        // 3: pb.GetItemResp
	(*GiftUpdateBagReq)(nil),   // 4: pb.GiftUpdateBagReq
	(*GiftUpdateBagResp)(nil),  // 5: pb.GiftUpdateBagResp
	(*GiftUpdateAttrReq)(nil),  // 6: pb.GiftUpdateAttrReq
	(*GiftUpdateAttrResp)(nil), // 7: pb.GiftUpdateAttrResp
}
var file_giftRpc_proto_depIdxs = []int32{
	0, // 0: pb.giftRpc.send:input_type -> pb.GiftSendReq
	4, // 1: pb.giftRpc.updateBag:input_type -> pb.GiftUpdateBagReq
	6, // 2: pb.giftRpc.updateAttr:input_type -> pb.GiftUpdateAttrReq
	1, // 3: pb.giftRpc.send:output_type -> pb.GiftSendResp
	5, // 4: pb.giftRpc.updateBag:output_type -> pb.GiftUpdateBagResp
	7, // 5: pb.giftRpc.updateAttr:output_type -> pb.GiftUpdateAttrResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_giftRpc_proto_init() }
func file_giftRpc_proto_init() {
	if File_giftRpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_giftRpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftSendReq); i {
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
		file_giftRpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftSendResp); i {
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
		file_giftRpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemReq); i {
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
		file_giftRpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemResp); i {
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
		file_giftRpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftUpdateBagReq); i {
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
		file_giftRpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftUpdateBagResp); i {
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
		file_giftRpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftUpdateAttrReq); i {
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
		file_giftRpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftUpdateAttrResp); i {
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
			RawDescriptor: file_giftRpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_giftRpc_proto_goTypes,
		DependencyIndexes: file_giftRpc_proto_depIdxs,
		MessageInfos:      file_giftRpc_proto_msgTypes,
	}.Build()
	File_giftRpc_proto = out.File
	file_giftRpc_proto_rawDesc = nil
	file_giftRpc_proto_goTypes = nil
	file_giftRpc_proto_depIdxs = nil
}
