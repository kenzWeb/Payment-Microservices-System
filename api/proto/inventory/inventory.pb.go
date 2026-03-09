package inventory

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

type CheckStockRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckStockRequest) Reset() {
	*x = CheckStockRequest{}
	mi := &file_inventory_inventory_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckStockRequest) ProtoMessage() {}

func (x *CheckStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_inventory_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CheckStockRequest) Descriptor() ([]byte, []int) {
	return file_inventory_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *CheckStockRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CheckStockRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CheckStockResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Available     bool                   `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckStockResponse) Reset() {
	*x = CheckStockResponse{}
	mi := &file_inventory_inventory_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckStockResponse) ProtoMessage() {}

func (x *CheckStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_inventory_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CheckStockResponse) Descriptor() ([]byte, []int) {
	return file_inventory_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *CheckStockResponse) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

type DeductStockRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeductStockRequest) Reset() {
	*x = DeductStockRequest{}
	mi := &file_inventory_inventory_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeductStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeductStockRequest) ProtoMessage() {}

func (x *DeductStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_inventory_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*DeductStockRequest) Descriptor() ([]byte, []int) {
	return file_inventory_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *DeductStockRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *DeductStockRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type DeductStockResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeductStockResponse) Reset() {
	*x = DeductStockResponse{}
	mi := &file_inventory_inventory_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeductStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeductStockResponse) ProtoMessage() {}

func (x *DeductStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_inventory_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*DeductStockResponse) Descriptor() ([]byte, []int) {
	return file_inventory_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *DeductStockResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_inventory_inventory_proto protoreflect.FileDescriptor

const file_inventory_inventory_proto_rawDesc = "" +
	"\n" +
	"\x19inventory/inventory.proto\x12\tinventory\"N\n" +
	"\x11CheckStockRequest\x12\x1d\n" +
	"\n" +
	"product_id\x18\x01 \x01(\tR\tproductId\x12\x1a\n" +
	"\bquantity\x18\x02 \x01(\x05R\bquantity\"2\n" +
	"\x12CheckStockResponse\x12\x1c\n" +
	"\tavailable\x18\x01 \x01(\bR\tavailable\"O\n" +
	"\x12DeductStockRequest\x12\x1d\n" +
	"\n" +
	"product_id\x18\x01 \x01(\tR\tproductId\x12\x1a\n" +
	"\bquantity\x18\x02 \x01(\x05R\bquantity\"/\n" +
	"\x13DeductStockResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\xab\x01\n" +
	"\x10InventoryService\x12I\n" +
	"\n" +
	"CheckStock\x12\x1c.inventory.CheckStockRequest\x1a\x1d.inventory.CheckStockResponse\x12L\n" +
	"\vDeductStock\x12\x1d.inventory.DeductStockRequest\x1a\x1e.inventory.DeductStockResponseB;Z9github.com/user/payment-microservices/api/proto/inventoryb\x06proto3"

var (
	file_inventory_inventory_proto_rawDescOnce sync.Once
	file_inventory_inventory_proto_rawDescData []byte
)

func file_inventory_inventory_proto_rawDescGZIP() []byte {
	file_inventory_inventory_proto_rawDescOnce.Do(func() {
		file_inventory_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_inventory_inventory_proto_rawDesc), len(file_inventory_inventory_proto_rawDesc)))
	})
	return file_inventory_inventory_proto_rawDescData
}

var file_inventory_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_inventory_inventory_proto_goTypes = []any{
	(*CheckStockRequest)(nil),
	(*CheckStockResponse)(nil),
	(*DeductStockRequest)(nil),
	(*DeductStockResponse)(nil),
}
var file_inventory_inventory_proto_depIdxs = []int32{
	0,
	2,
	1,
	3,
	2,
	0,
	0,
	0,
	0,
}

func init() { file_inventory_inventory_proto_init() }
func file_inventory_inventory_proto_init() {
	if File_inventory_inventory_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_inventory_inventory_proto_rawDesc), len(file_inventory_inventory_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventory_inventory_proto_goTypes,
		DependencyIndexes: file_inventory_inventory_proto_depIdxs,
		MessageInfos:      file_inventory_inventory_proto_msgTypes,
	}.Build()
	File_inventory_inventory_proto = out.File
	file_inventory_inventory_proto_goTypes = nil
	file_inventory_inventory_proto_depIdxs = nil
}
