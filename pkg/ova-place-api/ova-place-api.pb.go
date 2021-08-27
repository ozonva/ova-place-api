// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ova-place-api.proto

package ova_place_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreatePlaceRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Seat   string `protobuf:"bytes,2,opt,name=seat,proto3" json:"seat,omitempty"`
	Memo   string `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *CreatePlaceRequestV1) Reset() {
	*x = CreatePlaceRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_place_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlaceRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlaceRequestV1) ProtoMessage() {}

func (x *CreatePlaceRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_ova_place_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlaceRequestV1.ProtoReflect.Descriptor instead.
func (*CreatePlaceRequestV1) Descriptor() ([]byte, []int) {
	return file_ova_place_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePlaceRequestV1) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreatePlaceRequestV1) GetSeat() string {
	if x != nil {
		return x.Seat
	}
	return ""
}

func (x *CreatePlaceRequestV1) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

type DescribePlaceRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaceId uint64 `protobuf:"varint,1,opt,name=place_id,json=placeId,proto3" json:"place_id,omitempty"`
}

func (x *DescribePlaceRequestV1) Reset() {
	*x = DescribePlaceRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_place_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePlaceRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePlaceRequestV1) ProtoMessage() {}

func (x *DescribePlaceRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_ova_place_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePlaceRequestV1.ProtoReflect.Descriptor instead.
func (*DescribePlaceRequestV1) Descriptor() ([]byte, []int) {
	return file_ova_place_api_proto_rawDescGZIP(), []int{1}
}

func (x *DescribePlaceRequestV1) GetPlaceId() uint64 {
	if x != nil {
		return x.PlaceId
	}
	return 0
}

type ListPlacesRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage uint64 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
}

func (x *ListPlacesRequestV1) Reset() {
	*x = ListPlacesRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_place_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlacesRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlacesRequestV1) ProtoMessage() {}

func (x *ListPlacesRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_ova_place_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlacesRequestV1.ProtoReflect.Descriptor instead.
func (*ListPlacesRequestV1) Descriptor() ([]byte, []int) {
	return file_ova_place_api_proto_rawDescGZIP(), []int{2}
}

func (x *ListPlacesRequestV1) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListPlacesRequestV1) GetPerPage() uint64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type ListPlacesResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Places     []*PlaceV1    `protobuf:"bytes,1,rep,name=places,proto3" json:"places,omitempty"`
	Pagination *PaginationV1 `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListPlacesResponseV1) Reset() {
	*x = ListPlacesResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_place_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlacesResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlacesResponseV1) ProtoMessage() {}

func (x *ListPlacesResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_ova_place_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlacesResponseV1.ProtoReflect.Descriptor instead.
func (*ListPlacesResponseV1) Descriptor() ([]byte, []int) {
	return file_ova_place_api_proto_rawDescGZIP(), []int{3}
}

func (x *ListPlacesResponseV1) GetPlaces() []*PlaceV1 {
	if x != nil {
		return x.Places
	}
	return nil
}

func (x *ListPlacesResponseV1) GetPagination() *PaginationV1 {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type UpdatePlaceRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaceId uint64 `protobuf:"varint,1,opt,name=place_id,json=placeId,proto3" json:"place_id,omitempty"`
	UserId  uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Seat    string `protobuf:"bytes,3,opt,name=seat,proto3" json:"seat,omitempty"`
	Memo    string `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *UpdatePlaceRequestV1) Reset() {
	*x = UpdatePlaceRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_place_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlaceRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlaceRequestV1) ProtoMessage() {}

func (x *UpdatePlaceRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_ova_place_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlaceRequestV1.ProtoReflect.Descriptor instead.
func (*UpdatePlaceRequestV1) Descriptor() ([]byte, []int) {
	return file_ova_place_api_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePlaceRequestV1) GetPlaceId() uint64 {
	if x != nil {
		return x.PlaceId
	}
	return 0
}

func (x *UpdatePlaceRequestV1) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdatePlaceRequestV1) GetSeat() string {
	if x != nil {
		return x.Seat
	}
	return ""
}

func (x *UpdatePlaceRequestV1) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

type RemovePlaceRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaceId uint64 `protobuf:"varint,1,opt,name=place_id,json=placeId,proto3" json:"place_id,omitempty"`
}

func (x *RemovePlaceRequestV1) Reset() {
	*x = RemovePlaceRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_place_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePlaceRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePlaceRequestV1) ProtoMessage() {}

func (x *RemovePlaceRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_ova_place_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePlaceRequestV1.ProtoReflect.Descriptor instead.
func (*RemovePlaceRequestV1) Descriptor() ([]byte, []int) {
	return file_ova_place_api_proto_rawDescGZIP(), []int{5}
}

func (x *RemovePlaceRequestV1) GetPlaceId() uint64 {
	if x != nil {
		return x.PlaceId
	}
	return 0
}

type PlaceV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaceId uint64 `protobuf:"varint,1,opt,name=place_id,json=placeId,proto3" json:"place_id,omitempty"`
	UserId  uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Seat    string `protobuf:"bytes,3,opt,name=seat,proto3" json:"seat,omitempty"`
	Memo    string `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *PlaceV1) Reset() {
	*x = PlaceV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_place_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceV1) ProtoMessage() {}

func (x *PlaceV1) ProtoReflect() protoreflect.Message {
	mi := &file_ova_place_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceV1.ProtoReflect.Descriptor instead.
func (*PlaceV1) Descriptor() ([]byte, []int) {
	return file_ova_place_api_proto_rawDescGZIP(), []int{6}
}

func (x *PlaceV1) GetPlaceId() uint64 {
	if x != nil {
		return x.PlaceId
	}
	return 0
}

func (x *PlaceV1) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PlaceV1) GetSeat() string {
	if x != nil {
		return x.Seat
	}
	return ""
}

func (x *PlaceV1) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

type PaginationV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage uint64 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Total   uint64 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PaginationV1) Reset() {
	*x = PaginationV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_place_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationV1) ProtoMessage() {}

func (x *PaginationV1) ProtoReflect() protoreflect.Message {
	mi := &file_ova_place_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationV1.ProtoReflect.Descriptor instead.
func (*PaginationV1) Descriptor() ([]byte, []int) {
	return file_ova_place_api_proto_rawDescGZIP(), []int{7}
}

func (x *PaginationV1) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationV1) GetPerPage() uint64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *PaginationV1) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type EmptyV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyV1) Reset() {
	*x = EmptyV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_place_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyV1) ProtoMessage() {}

func (x *EmptyV1) ProtoReflect() protoreflect.Message {
	mi := &file_ova_place_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyV1.ProtoReflect.Descriptor instead.
func (*EmptyV1) Descriptor() ([]byte, []int) {
	return file_ova_place_api_proto_rawDescGZIP(), []int{8}
}

var File_ova_place_api_proto protoreflect.FileDescriptor

var file_ova_place_api_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6f, 0x76, 0x61, 0x2d, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6f, 0x76, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x56, 0x31, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x73, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52,
	0x04, 0x73, 0x65, 0x61, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x22, 0x3c, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12,
	0x22, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20,
	0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02,
	0x20, 0x00, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x56, 0x31, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x56, 0x31, 0x52, 0x06, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x9c, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x22, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x04, 0x73, 0x65, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x04, 0x73, 0x65, 0x61, 0x74,
	0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f,
	0x22, 0x3a, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x22, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32,
	0x02, 0x20, 0x00, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x07,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x56, 0x31, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x65, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x65, 0x61, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x22, 0x53, 0x0a, 0x0c, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x56, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x09, 0x0a, 0x07, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x56, 0x31, 0x32, 0xa9, 0x04, 0x0a, 0x0d, 0x4f, 0x76, 0x61, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x41, 0x70, 0x69, 0x56, 0x31, 0x12, 0x60, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x56, 0x31, 0x12, 0x23, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x16, 0x2e, 0x6f, 0x76,
	0x61, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x56, 0x31, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x0a, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x12, 0x6f, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x56, 0x31, 0x12, 0x25, 0x2e, 0x6f, 0x76, 0x61,
	0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56,
	0x31, 0x1a, 0x16, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x56, 0x31, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6b, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x56, 0x31, 0x12, 0x22, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x23, 0x2e, 0x6f,
	0x76, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56,
	0x31, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x73, 0x12, 0x6b, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x56, 0x31, 0x12, 0x23, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x16, 0x2e, 0x6f, 0x76,
	0x61, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x56, 0x31, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x1a, 0x15, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x6b, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x56, 0x31, 0x12, 0x23, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x16, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x56, 0x31,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x2a, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x42,
	0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a,
	0x6f, 0x6e, 0x76, 0x61, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x76, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ova_place_api_proto_rawDescOnce sync.Once
	file_ova_place_api_proto_rawDescData = file_ova_place_api_proto_rawDesc
)

func file_ova_place_api_proto_rawDescGZIP() []byte {
	file_ova_place_api_proto_rawDescOnce.Do(func() {
		file_ova_place_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_ova_place_api_proto_rawDescData)
	})
	return file_ova_place_api_proto_rawDescData
}

var file_ova_place_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ova_place_api_proto_goTypes = []interface{}{
	(*CreatePlaceRequestV1)(nil),   // 0: ova.place.api.CreatePlaceRequestV1
	(*DescribePlaceRequestV1)(nil), // 1: ova.place.api.DescribePlaceRequestV1
	(*ListPlacesRequestV1)(nil),    // 2: ova.place.api.ListPlacesRequestV1
	(*ListPlacesResponseV1)(nil),   // 3: ova.place.api.ListPlacesResponseV1
	(*UpdatePlaceRequestV1)(nil),   // 4: ova.place.api.UpdatePlaceRequestV1
	(*RemovePlaceRequestV1)(nil),   // 5: ova.place.api.RemovePlaceRequestV1
	(*PlaceV1)(nil),                // 6: ova.place.api.PlaceV1
	(*PaginationV1)(nil),           // 7: ova.place.api.PaginationV1
	(*EmptyV1)(nil),                // 8: ova.place.api.EmptyV1
}
var file_ova_place_api_proto_depIdxs = []int32{
	6, // 0: ova.place.api.ListPlacesResponseV1.places:type_name -> ova.place.api.PlaceV1
	7, // 1: ova.place.api.ListPlacesResponseV1.pagination:type_name -> ova.place.api.PaginationV1
	0, // 2: ova.place.api.OvaPlaceApiV1.CreatePlaceV1:input_type -> ova.place.api.CreatePlaceRequestV1
	1, // 3: ova.place.api.OvaPlaceApiV1.DescribePlaceV1:input_type -> ova.place.api.DescribePlaceRequestV1
	2, // 4: ova.place.api.OvaPlaceApiV1.ListPlacesV1:input_type -> ova.place.api.ListPlacesRequestV1
	4, // 5: ova.place.api.OvaPlaceApiV1.UpdatePlaceV1:input_type -> ova.place.api.UpdatePlaceRequestV1
	5, // 6: ova.place.api.OvaPlaceApiV1.RemovePlaceV1:input_type -> ova.place.api.RemovePlaceRequestV1
	6, // 7: ova.place.api.OvaPlaceApiV1.CreatePlaceV1:output_type -> ova.place.api.PlaceV1
	6, // 8: ova.place.api.OvaPlaceApiV1.DescribePlaceV1:output_type -> ova.place.api.PlaceV1
	3, // 9: ova.place.api.OvaPlaceApiV1.ListPlacesV1:output_type -> ova.place.api.ListPlacesResponseV1
	6, // 10: ova.place.api.OvaPlaceApiV1.UpdatePlaceV1:output_type -> ova.place.api.PlaceV1
	8, // 11: ova.place.api.OvaPlaceApiV1.RemovePlaceV1:output_type -> ova.place.api.EmptyV1
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ova_place_api_proto_init() }
func file_ova_place_api_proto_init() {
	if File_ova_place_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ova_place_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlaceRequestV1); i {
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
		file_ova_place_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePlaceRequestV1); i {
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
		file_ova_place_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPlacesRequestV1); i {
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
		file_ova_place_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPlacesResponseV1); i {
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
		file_ova_place_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlaceRequestV1); i {
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
		file_ova_place_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePlaceRequestV1); i {
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
		file_ova_place_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceV1); i {
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
		file_ova_place_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationV1); i {
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
		file_ova_place_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyV1); i {
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
			RawDescriptor: file_ova_place_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ova_place_api_proto_goTypes,
		DependencyIndexes: file_ova_place_api_proto_depIdxs,
		MessageInfos:      file_ova_place_api_proto_msgTypes,
	}.Build()
	File_ova_place_api_proto = out.File
	file_ova_place_api_proto_rawDesc = nil
	file_ova_place_api_proto_goTypes = nil
	file_ova_place_api_proto_depIdxs = nil
}
