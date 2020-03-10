// Code generated by protoc-gen-go. DO NOT EDIT.
// source: carts.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type Cart struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Items                []*Cart_Item         `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Cart) Reset()         { *m = Cart{} }
func (m *Cart) String() string { return proto.CompactTextString(m) }
func (*Cart) ProtoMessage()    {}
func (*Cart) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb9f7b8935c2f010, []int{0}
}

func (m *Cart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cart.Unmarshal(m, b)
}
func (m *Cart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cart.Marshal(b, m, deterministic)
}
func (m *Cart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cart.Merge(m, src)
}
func (m *Cart) XXX_Size() int {
	return xxx_messageInfo_Cart.Size(m)
}
func (m *Cart) XXX_DiscardUnknown() {
	xxx_messageInfo_Cart.DiscardUnknown(m)
}

var xxx_messageInfo_Cart proto.InternalMessageInfo

func (m *Cart) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Cart) GetItems() []*Cart_Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Cart) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Cart) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

type Cart_Item struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Qty                  string   `protobuf:"bytes,2,opt,name=qty,proto3" json:"qty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cart_Item) Reset()         { *m = Cart_Item{} }
func (m *Cart_Item) String() string { return proto.CompactTextString(m) }
func (*Cart_Item) ProtoMessage()    {}
func (*Cart_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb9f7b8935c2f010, []int{0, 0}
}

func (m *Cart_Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cart_Item.Unmarshal(m, b)
}
func (m *Cart_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cart_Item.Marshal(b, m, deterministic)
}
func (m *Cart_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cart_Item.Merge(m, src)
}
func (m *Cart_Item) XXX_Size() int {
	return xxx_messageInfo_Cart_Item.Size(m)
}
func (m *Cart_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Cart_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Cart_Item proto.InternalMessageInfo

func (m *Cart_Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Cart_Item) GetQty() string {
	if m != nil {
		return m.Qty
	}
	return ""
}

type CreateRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Cart                 *Cart    `protobuf:"bytes,2,opt,name=cart,proto3" json:"cart,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb9f7b8935c2f010, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateRequest) GetCart() *Cart {
	if m != nil {
		return m.Cart
	}
	return nil
}

type CreateResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb9f7b8935c2f010, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb9f7b8935c2f010, []int{3}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Cart                 *Cart    `protobuf:"bytes,2,opt,name=cart,proto3" json:"cart,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb9f7b8935c2f010, []int{4}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadResponse) GetCart() *Cart {
	if m != nil {
		return m.Cart
	}
	return nil
}

type UpdateRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Cart                 *Cart    `protobuf:"bytes,2,opt,name=cart,proto3" json:"cart,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb9f7b8935c2f010, []int{5}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateRequest) GetCart() *Cart {
	if m != nil {
		return m.Cart
	}
	return nil
}

type UpdateResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Updated              int64    `protobuf:"varint,2,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb9f7b8935c2f010, []int{6}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateResponse) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type DeleteRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb9f7b8935c2f010, []int{7}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Deleted              int64    `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb9f7b8935c2f010, []int{8}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteResponse) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

type ListRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb9f7b8935c2f010, []int{9}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

type ListResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Data                 []*Cart  `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb9f7b8935c2f010, []int{10}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ListResponse) GetData() []*Cart {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Cart)(nil), "v1.Cart")
	proto.RegisterType((*Cart_Item)(nil), "v1.Cart.Item")
	proto.RegisterType((*CreateRequest)(nil), "v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "v1.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "v1.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "v1.ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "v1.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "v1.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "v1.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "v1.DeleteResponse")
	proto.RegisterType((*ListRequest)(nil), "v1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "v1.ListResponse")
}

func init() {
	proto.RegisterFile("carts.proto", fileDescriptor_eb9f7b8935c2f010)
}

var fileDescriptor_eb9f7b8935c2f010 = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcf, 0x6f, 0xd3, 0x30,
	0x18, 0x55, 0xd2, 0xd2, 0x31, 0x97, 0x76, 0xc5, 0xfc, 0x50, 0x89, 0x26, 0x2d, 0x0a, 0x97, 0xaa,
	0xa2, 0xc9, 0x1a, 0x7a, 0x40, 0x65, 0x62, 0x1b, 0x03, 0x21, 0xa4, 0x71, 0x09, 0x70, 0x80, 0x9b,
	0xd7, 0x7c, 0x6b, 0xcd, 0x9a, 0xd8, 0x8b, 0xdd, 0x0e, 0x84, 0x76, 0xe1, 0xce, 0x05, 0x6e, 0xfc,
	0x57, 0x08, 0x71, 0xe0, 0xce, 0x1f, 0x82, 0x6c, 0x27, 0x5b, 0xbb, 0xd1, 0x22, 0x71, 0x4a, 0xfc,
	0xf9, 0x7d, 0xef, 0x3d, 0xfb, 0x7b, 0x09, 0xaa, 0x0e, 0x48, 0x26, 0x85, 0xcf, 0x33, 0x26, 0x19,
	0xb6, 0xa7, 0x5d, 0x67, 0x63, 0xc8, 0xd8, 0x70, 0x0c, 0x81, 0xae, 0x1c, 0x4c, 0x0e, 0x03, 0x49,
	0x13, 0x10, 0x92, 0x24, 0xdc, 0x80, 0x9c, 0xf5, 0x1c, 0x40, 0x38, 0x0d, 0x48, 0x9a, 0x32, 0x49,
	0x24, 0x65, 0x69, 0x4e, 0xe1, 0xdc, 0xd3, 0x8f, 0x41, 0x67, 0x08, 0x69, 0x47, 0x9c, 0x90, 0xe1,
	0x10, 0xb2, 0x80, 0x71, 0x8d, 0xb8, 0x8c, 0xf6, 0xbe, 0x5b, 0xa8, 0xbc, 0x47, 0x32, 0x89, 0xeb,
	0xc8, 0xa6, 0x71, 0xd3, 0x72, 0xad, 0xd6, 0x6a, 0x64, 0xd3, 0x18, 0xdf, 0x45, 0x57, 0xa8, 0x84,
	0x44, 0x34, 0x6d, 0xb7, 0xd4, 0xaa, 0x86, 0x35, 0x7f, 0xda, 0xf5, 0x15, 0xd0, 0x7f, 0x2e, 0x21,
	0x89, 0xcc, 0x1e, 0xee, 0xa1, 0x95, 0x41, 0x06, 0x44, 0x42, 0xdc, 0x2c, 0xb9, 0x56, 0xab, 0x1a,
	0x3a, 0xbe, 0xf1, 0xe6, 0x17, 0xe6, 0xfd, 0x57, 0x85, 0xf9, 0xa8, 0x80, 0xaa, 0xae, 0x09, 0x8f,
	0x75, 0x57, 0xf9, 0xdf, 0x5d, 0x39, 0xd4, 0x69, 0xa1, 0xb2, 0x92, 0xbe, 0x64, 0xb4, 0x81, 0x4a,
	0xc7, 0xf2, 0x43, 0xd3, 0xd6, 0x05, 0xf5, 0xea, 0x6d, 0xa3, 0xda, 0x9e, 0x96, 0x8a, 0xe0, 0x78,
	0x02, 0x42, 0x2a, 0x08, 0xe1, 0x34, 0xef, 0x51, 0xaf, 0x78, 0x1d, 0x95, 0xd5, 0xb5, 0xeb, 0xae,
	0x6a, 0x78, 0xb5, 0x38, 0x5c, 0xa4, 0xab, 0x5e, 0x88, 0xea, 0x05, 0x81, 0xe0, 0x2c, 0x15, 0xf0,
	0x17, 0x06, 0x63, 0xc3, 0x2e, 0x6c, 0x78, 0x01, 0xaa, 0x46, 0x40, 0xe2, 0xc5, 0x92, 0x17, 0x1b,
	0x1e, 0xa1, 0x6b, 0xa6, 0x61, 0xa1, 0xc4, 0x72, 0x93, 0xdb, 0xa8, 0xf6, 0x5a, 0x5f, 0xcd, 0xff,
	0x9e, 0x72, 0x0b, 0xd5, 0x0b, 0x82, 0x85, 0x16, 0x9a, 0xe7, 0xa3, 0x52, 0x24, 0xa5, 0xb3, 0x71,
	0x78, 0x5d, 0x54, 0x7b, 0x02, 0x63, 0x58, 0x26, 0x7f, 0xf1, 0xc4, 0x5b, 0xa8, 0x5e, 0xb4, 0x2c,
	0x13, 0x8c, 0x35, 0xe6, 0x4c, 0x30, 0x5f, 0x7a, 0x1b, 0xa8, 0xba, 0x4f, 0x85, 0x5c, 0x28, 0xa7,
	0x2e, 0xd4, 0x00, 0x96, 0x5d, 0x68, 0x4c, 0x24, 0xc9, 0x23, 0x3d, 0x73, 0x1f, 0xaa, 0x1a, 0x7e,
	0x2e, 0xa1, 0xaa, 0x5a, 0xbe, 0x84, 0x6c, 0x4a, 0x07, 0x80, 0x77, 0x50, 0x59, 0xf1, 0xe1, 0x35,
	0x85, 0x9b, 0x91, 0x76, 0x1a, 0xe7, 0x05, 0x23, 0xe5, 0xdd, 0xfa, 0xf4, 0xe3, 0xf7, 0x57, 0x7b,
	0x0d, 0xd7, 0x82, 0x69, 0x37, 0xd0, 0xdf, 0x73, 0x40, 0xc6, 0x63, 0xfc, 0x14, 0x55, 0x4c, 0x8e,
	0xf0, 0x75, 0xad, 0x35, 0x1b, 0x4a, 0x07, 0xcf, 0x96, 0x72, 0x9e, 0x9b, 0x9a, 0xa7, 0xee, 0xad,
	0x9e, 0xf1, 0xf4, 0xad, 0x36, 0xde, 0x45, 0x65, 0x95, 0x14, 0x63, 0x64, 0x26, 0x64, 0xc6, 0xc8,
	0x6c, 0x88, 0xbc, 0xdb, 0x9a, 0xa0, 0x81, 0xeb, 0xe7, 0x46, 0x3e, 0xd2, 0xf8, 0x14, 0x53, 0x54,
	0x31, 0xb3, 0x36, 0x4e, 0xe6, 0x82, 0x63, 0x9c, 0xcc, 0x47, 0xc1, 0x7b, 0xa0, 0x89, 0x42, 0xe7,
	0xc6, 0x0c, 0x91, 0x7a, 0xf8, 0x34, 0x3e, 0xed, 0x5b, 0xed, 0xb7, 0xcd, 0x70, 0xc1, 0x0e, 0x7e,
	0x86, 0x2a, 0x66, 0xca, 0x46, 0x6a, 0x2e, 0x24, 0x46, 0x6a, 0x3e, 0x04, 0x85, 0xe7, 0xf6, 0x05,
	0xcf, 0x8f, 0x7f, 0x59, 0x5f, 0x76, 0x7f, 0x5a, 0x98, 0xa1, 0x9a, 0x9a, 0x8a, 0x70, 0xf3, 0xb9,
	0x78, 0x6f, 0x10, 0x7a, 0x41, 0x8f, 0xc0, 0xdd, 0x27, 0x20, 0x09, 0xde, 0x1c, 0x49, 0xc9, 0x45,
	0x3f, 0x08, 0xc6, 0x6a, 0xd9, 0x11, 0x23, 0xc6, 0x39, 0x4d, 0x87, 0x1d, 0x45, 0xd4, 0x89, 0x21,
	0x61, 0xfe, 0x21, 0xcd, 0xe0, 0x80, 0x08, 0x20, 0x9c, 0xfb, 0x03, 0x96, 0x38, 0x77, 0x12, 0x7a,
	0x04, 0x3b, 0x1a, 0x1e, 0xc3, 0x14, 0xc6, 0x8c, 0x27, 0x90, 0x4a, 0xb5, 0x15, 0x96, 0xba, 0xfe,
	0x66, 0xdb, 0xb2, 0xc2, 0x06, 0xe1, 0x7c, 0x4c, 0x07, 0xfa, 0x4f, 0x19, 0xbc, 0x13, 0x2c, 0xed,
	0x5f, 0xaa, 0x44, 0x0f, 0x51, 0xa9, 0xb7, 0xd9, 0xc3, 0x3d, 0xd4, 0x8e, 0x40, 0x4e, 0xb2, 0x14,
	0x62, 0xf7, 0x64, 0x04, 0xa9, 0x2b, 0x47, 0xe0, 0x66, 0x20, 0xd8, 0x24, 0x1b, 0x80, 0x1b, 0x33,
	0x10, 0x6e, 0xca, 0xa4, 0x0b, 0xef, 0xa9, 0x90, 0x3e, 0xae, 0xa0, 0xf2, 0x37, 0xdb, 0x5a, 0x39,
	0xa8, 0xe8, 0x3f, 0xdd, 0xfd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x7a, 0x58, 0x41, 0xfd,
	0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CartServiceClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/v1.CartService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/v1.CartService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/v1.CartService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/v1.CartService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.CartService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
type CartServiceServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

// UnimplementedCartServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (*UnimplementedCartServiceServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedCartServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCartServiceServer) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedCartServiceServer) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedCartServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterCartServiceServer(s *grpc.Server, srv CartServiceServer) {
	s.RegisterService(&_CartService_serviceDesc, srv)
}

func _CartService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CartService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CartService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CartService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CartService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CartService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CartService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _CartService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CartService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _CartService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CartService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CartService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "carts.proto",
}
