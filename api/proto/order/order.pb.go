package order

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price         float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	mi := &file_order_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *OrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderItem) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items         []*OrderItem           `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_order_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateOrderRequest) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	mi := &file_order_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreateOrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_order_order_proto protoreflect.FileDescriptor

const file_order_order_proto_rawDesc = "" +
	"\n" +
	"\x11order/order.proto\x12\x05order\"\\\n" +
	"\tOrderItem\x12\x1d\n" +
	"\n" +
	"product_id\x18\x01 \x01(\tR\tproductId\x12\x1a\n" +
	"\bquantity\x18\x02 \x01(\x05R\bquantity\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x01R\x05price\"U\n" +
	"\x12CreateOrderRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12&\n" +
	"\x05items\x18\x02 \x03(\v2\x10.order.OrderItemR\x05items\"H\n" +
	"\x13CreateOrderResponse\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status2T\n" +
	"\fOrderService\x12D\n" +
	"\vCreateOrder\x12\x19.order.CreateOrderRequest\x1a\x1a.order.CreateOrderResponseB7Z5github.com/user/payment-microservices/api/proto/orderb\x06proto3"

var (
	file_order_order_proto_rawDescOnce sync.Once
	file_order_order_proto_rawDescData []byte
)

func file_order_order_proto_rawDescGZIP() []byte {
	file_order_order_proto_rawDescOnce.Do(func() {
		file_order_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_order_order_proto_rawDesc), len(file_order_order_proto_rawDesc)))
	})
	return file_order_order_proto_rawDescData
}

var file_order_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_order_order_proto_goTypes = []any{
	(*OrderItem)(nil),
	(*CreateOrderRequest)(nil),
	(*CreateOrderResponse)(nil),
}
var file_order_order_proto_depIdxs = []int32{
	0,
	1,
	2,
	2,
	1,
	1,
	1,
	0,
}

func init() { file_order_order_proto_init() }
func file_order_order_proto_init() {
	if File_order_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_order_order_proto_rawDesc), len(file_order_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_order_proto_goTypes,
		DependencyIndexes: file_order_order_proto_depIdxs,
		MessageInfos:      file_order_order_proto_msgTypes,
	}.Build()
	File_order_order_proto = out.File
	file_order_order_proto_goTypes = nil
	file_order_order_proto_depIdxs = nil
}
