// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: internal/proto/api.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlaceOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	IsBid  bool   `protobuf:"varint,2,opt,name=isBid,proto3" json:"isBid,omitempty"`
	Symbol string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Qty    string `protobuf:"bytes,4,opt,name=qty,proto3" json:"qty,omitempty"`
	Price  string `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	Type   string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *PlaceOrderReq) Reset() {
	*x = PlaceOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderReq) ProtoMessage() {}

func (x *PlaceOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderReq.ProtoReflect.Descriptor instead.
func (*PlaceOrderReq) Descriptor() ([]byte, []int) {
	return file_internal_proto_api_proto_rawDescGZIP(), []int{0}
}

func (x *PlaceOrderReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *PlaceOrderReq) GetIsBid() bool {
	if x != nil {
		return x.IsBid
	}
	return false
}

func (x *PlaceOrderReq) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *PlaceOrderReq) GetQty() string {
	if x != nil {
		return x.Qty
	}
	return ""
}

func (x *PlaceOrderReq) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *PlaceOrderReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Orders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=Orders,proto3" json:"Orders,omitempty"`
}

func (x *Orders) Reset() {
	*x = Orders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orders) ProtoMessage() {}

func (x *Orders) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orders.ProtoReflect.Descriptor instead.
func (*Orders) Descriptor() ([]byte, []int) {
	return file_internal_proto_api_proto_rawDescGZIP(), []int{1}
}

func (x *Orders) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type OrderID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID int64 `protobuf:"varint,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
}

func (x *OrderID) Reset() {
	*x = OrderID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderID) ProtoMessage() {}

func (x *OrderID) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderID.ProtoReflect.Descriptor instead.
func (*OrderID) Descriptor() ([]byte, []int) {
	return file_internal_proto_api_proto_rawDescGZIP(), []int{2}
}

func (x *OrderID) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_internal_proto_api_proto_rawDescGZIP(), []int{3}
}

func (x *UserID) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type OrderBookSymbol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *OrderBookSymbol) Reset() {
	*x = OrderBookSymbol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderBookSymbol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderBookSymbol) ProtoMessage() {}

func (x *OrderBookSymbol) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderBookSymbol.ProtoReflect.Descriptor instead.
func (*OrderBookSymbol) Descriptor() ([]byte, []int) {
	return file_internal_proto_api_proto_rawDescGZIP(), []int{4}
}

func (x *OrderBookSymbol) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         int64                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID     int64                  `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	IsBid      bool                   `protobuf:"varint,3,opt,name=isBid,proto3" json:"isBid,omitempty"`
	Symbol     string                 `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Price      string                 `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	Qty        string                 `protobuf:"bytes,6,opt,name=qty,proto3" json:"qty,omitempty"`
	SizeFilled string                 `protobuf:"bytes,7,opt,name=sizeFilled,proto3" json:"sizeFilled,omitempty"`
	Status     string                 `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Type       string                 `protobuf:"bytes,9,opt,name=type,proto3" json:"type,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ClosedAt   *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=closed_at,json=closedAt,proto3" json:"closed_at,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_internal_proto_api_proto_rawDescGZIP(), []int{5}
}

func (x *Order) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Order) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Order) GetIsBid() bool {
	if x != nil {
		return x.IsBid
	}
	return false
}

func (x *Order) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Order) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Order) GetQty() string {
	if x != nil {
		return x.Qty
	}
	return ""
}

func (x *Order) GetSizeFilled() string {
	if x != nil {
		return x.SizeFilled
	}
	return ""
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Order) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Order) GetClosedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ClosedAt
	}
	return nil
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_internal_proto_api_proto_rawDescGZIP(), []int{6}
}

func (x *Ping) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type OrderBookSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol  string   `protobuf:"bytes,1,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	Bids    []*Limit `protobuf:"bytes,2,rep,name=Bids,proto3" json:"Bids,omitempty"`
	Asks    []*Limit `protobuf:"bytes,3,rep,name=Asks,proto3" json:"Asks,omitempty"`
	BidsQty string   `protobuf:"bytes,4,opt,name=BidsQty,proto3" json:"BidsQty,omitempty"`
	AsksQty string   `protobuf:"bytes,5,opt,name=AsksQty,proto3" json:"AsksQty,omitempty"`
}

func (x *OrderBookSnapshot) Reset() {
	*x = OrderBookSnapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderBookSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderBookSnapshot) ProtoMessage() {}

func (x *OrderBookSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderBookSnapshot.ProtoReflect.Descriptor instead.
func (*OrderBookSnapshot) Descriptor() ([]byte, []int) {
	return file_internal_proto_api_proto_rawDescGZIP(), []int{7}
}

func (x *OrderBookSnapshot) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *OrderBookSnapshot) GetBids() []*Limit {
	if x != nil {
		return x.Bids
	}
	return nil
}

func (x *OrderBookSnapshot) GetAsks() []*Limit {
	if x != nil {
		return x.Asks
	}
	return nil
}

func (x *OrderBookSnapshot) GetBidsQty() string {
	if x != nil {
		return x.BidsQty
	}
	return ""
}

func (x *OrderBookSnapshot) GetAsksQty() string {
	if x != nil {
		return x.AsksQty
	}
	return ""
}

type Limit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price string `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	Qty   string `protobuf:"bytes,2,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *Limit) Reset() {
	*x = Limit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Limit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Limit) ProtoMessage() {}

func (x *Limit) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Limit.ProtoReflect.Descriptor instead.
func (*Limit) Descriptor() ([]byte, []int) {
	return file_internal_proto_api_proto_rawDescGZIP(), []int{8}
}

func (x *Limit) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Limit) GetQty() string {
	if x != nil {
		return x.Qty
	}
	return ""
}

type Trades struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Trade  []*Trade `protobuf:"bytes,2,rep,name=trade,proto3" json:"trade,omitempty"`
}

func (x *Trades) Reset() {
	*x = Trades{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trades) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trades) ProtoMessage() {}

func (x *Trades) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trades.ProtoReflect.Descriptor instead.
func (*Trades) Descriptor() ([]byte, []int) {
	return file_internal_proto_api_proto_rawDescGZIP(), []int{9}
}

func (x *Trades) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Trades) GetTrade() []*Trade {
	if x != nil {
		return x.Trade
	}
	return nil
}

type Trade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsBid bool                   `protobuf:"varint,1,opt,name=isBid,proto3" json:"isBid,omitempty"`
	Price string                 `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
	Qty   string                 `protobuf:"bytes,3,opt,name=qty,proto3" json:"qty,omitempty"`
	Time  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Trade) Reset() {
	*x = Trade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trade) ProtoMessage() {}

func (x *Trade) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trade.ProtoReflect.Descriptor instead.
func (*Trade) Descriptor() ([]byte, []int) {
	return file_internal_proto_api_proto_rawDescGZIP(), []int{10}
}

func (x *Trade) GetIsBid() bool {
	if x != nil {
		return x.IsBid
	}
	return false
}

func (x *Trade) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Trade) GetQty() string {
	if x != nil {
		return x.Qty
	}
	return ""
}

func (x *Trade) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

var File_internal_proto_api_proto protoreflect.FileDescriptor

var file_internal_proto_api_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a,
	0x0d, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x42, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x42, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x71, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x2b, 0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x06, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x23, 0x0a,
	0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x20, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x22, 0x29, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f,
	0x6b, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22,
	0xc5, 0x02, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x42, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x69, 0x73, 0x42, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x71, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x46,
	0x69, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x69, 0x7a,
	0x65, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x37,
	0x0a, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x64, 0x41, 0x74, 0x22, 0x1a, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x69, 0x6e, 0x67, 0x22, 0x9d, 0x01, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f,
	0x6b, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x12, 0x1d, 0x0a, 0x04, 0x42, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x70, 0x62, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x04, 0x42, 0x69, 0x64, 0x73,
	0x12, 0x1d, 0x0a, 0x04, 0x41, 0x73, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x70, 0x62, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x04, 0x41, 0x73, 0x6b, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x42, 0x69, 0x64, 0x73, 0x51, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x42, 0x69, 0x64, 0x73, 0x51, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x73, 0x6b,
	0x73, 0x51, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x73, 0x6b, 0x73,
	0x51, 0x74, 0x79, 0x22, 0x2f, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x71, 0x74, 0x79, 0x22, 0x41, 0x0a, 0x06, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x52, 0x05, 0x74, 0x72, 0x61, 0x64, 0x65, 0x22, 0x75, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x42, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x69, 0x73, 0x42, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x71, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x71, 0x74, 0x79, 0x12, 0x2e,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x32, 0xea,
	0x03, 0x0a, 0x0e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x11, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x00,
	0x12, 0x27, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x70,
	0x62, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f,
	0x6b, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b,
	0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x6f, 0x6f, 0x6b, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f,
	0x6f, 0x6b, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x25, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12,
	0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3b, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x15, 0x2e, 0x70, 0x62,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x43, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65,
	0x63, 0x69, 0x73, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x05, 0x5a, 0x03, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_api_proto_rawDescOnce sync.Once
	file_internal_proto_api_proto_rawDescData = file_internal_proto_api_proto_rawDesc
)

func file_internal_proto_api_proto_rawDescGZIP() []byte {
	file_internal_proto_api_proto_rawDescOnce.Do(func() {
		file_internal_proto_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_api_proto_rawDescData)
	})
	return file_internal_proto_api_proto_rawDescData
}

var file_internal_proto_api_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_internal_proto_api_proto_goTypes = []any{
	(*PlaceOrderReq)(nil),         // 0: pb.PlaceOrderReq
	(*Orders)(nil),                // 1: pb.Orders
	(*OrderID)(nil),               // 2: pb.OrderID
	(*UserID)(nil),                // 3: pb.UserID
	(*OrderBookSymbol)(nil),       // 4: pb.OrderBookSymbol
	(*Order)(nil),                 // 5: pb.order
	(*Ping)(nil),                  // 6: pb.Ping
	(*OrderBookSnapshot)(nil),     // 7: pb.OrderBookSnapshot
	(*Limit)(nil),                 // 8: pb.limit
	(*Trades)(nil),                // 9: pb.Trades
	(*Trade)(nil),                 // 10: pb.trade
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 12: google.protobuf.Empty
}
var file_internal_proto_api_proto_depIdxs = []int32{
	5,  // 0: pb.Orders.Orders:type_name -> pb.order
	11, // 1: pb.order.created_at:type_name -> google.protobuf.Timestamp
	11, // 2: pb.order.closed_at:type_name -> google.protobuf.Timestamp
	8,  // 3: pb.OrderBookSnapshot.Bids:type_name -> pb.limit
	8,  // 4: pb.OrderBookSnapshot.Asks:type_name -> pb.limit
	10, // 5: pb.Trades.trade:type_name -> pb.trade
	11, // 6: pb.trade.time:type_name -> google.protobuf.Timestamp
	0,  // 7: pb.matchingEngine.PlaceOrder:input_type -> pb.PlaceOrderReq
	2,  // 8: pb.matchingEngine.CancelOrder:input_type -> pb.OrderID
	3,  // 9: pb.matchingEngine.GetCurrentOrders:input_type -> pb.UserID
	3,  // 10: pb.matchingEngine.GetOrders:input_type -> pb.UserID
	4,  // 11: pb.matchingEngine.CreateOrderBook:input_type -> pb.OrderBookSymbol
	4,  // 12: pb.matchingEngine.DeleteOrderBook:input_type -> pb.OrderBookSymbol
	6,  // 13: pb.matchingEngine.GetTrades:input_type -> pb.Ping
	6,  // 14: pb.matchingEngine.GetOrderBookSnapshot:input_type -> pb.Ping
	6,  // 15: pb.matchingEngine.GetPrecisedOrderBookSnapshot:input_type -> pb.Ping
	1,  // 16: pb.matchingEngine.PlaceOrder:output_type -> pb.Orders
	5,  // 17: pb.matchingEngine.CancelOrder:output_type -> pb.order
	1,  // 18: pb.matchingEngine.GetCurrentOrders:output_type -> pb.Orders
	1,  // 19: pb.matchingEngine.GetOrders:output_type -> pb.Orders
	12, // 20: pb.matchingEngine.CreateOrderBook:output_type -> google.protobuf.Empty
	12, // 21: pb.matchingEngine.DeleteOrderBook:output_type -> google.protobuf.Empty
	9,  // 22: pb.matchingEngine.GetTrades:output_type -> pb.Trades
	7,  // 23: pb.matchingEngine.GetOrderBookSnapshot:output_type -> pb.OrderBookSnapshot
	7,  // 24: pb.matchingEngine.GetPrecisedOrderBookSnapshot:output_type -> pb.OrderBookSnapshot
	16, // [16:25] is the sub-list for method output_type
	7,  // [7:16] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_internal_proto_api_proto_init() }
func file_internal_proto_api_proto_init() {
	if File_internal_proto_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_api_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PlaceOrderReq); i {
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
		file_internal_proto_api_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Orders); i {
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
		file_internal_proto_api_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*OrderID); i {
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
		file_internal_proto_api_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UserID); i {
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
		file_internal_proto_api_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*OrderBookSymbol); i {
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
		file_internal_proto_api_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Order); i {
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
		file_internal_proto_api_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Ping); i {
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
		file_internal_proto_api_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*OrderBookSnapshot); i {
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
		file_internal_proto_api_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Limit); i {
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
		file_internal_proto_api_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*Trades); i {
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
		file_internal_proto_api_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*Trade); i {
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
			RawDescriptor: file_internal_proto_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_api_proto_goTypes,
		DependencyIndexes: file_internal_proto_api_proto_depIdxs,
		MessageInfos:      file_internal_proto_api_proto_msgTypes,
	}.Build()
	File_internal_proto_api_proto = out.File
	file_internal_proto_api_proto_rawDesc = nil
	file_internal_proto_api_proto_goTypes = nil
	file_internal_proto_api_proto_depIdxs = nil
}
