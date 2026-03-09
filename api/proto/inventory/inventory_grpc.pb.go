package inventory

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	InventoryService_CheckStock_FullMethodName  = "/inventory.InventoryService/CheckStock"
	InventoryService_DeductStock_FullMethodName = "/inventory.InventoryService/DeductStock"
)

type InventoryServiceClient interface {
	CheckStock(ctx context.Context, in *CheckStockRequest, opts ...grpc.CallOption) (*CheckStockResponse, error)
	DeductStock(ctx context.Context, in *DeductStockRequest, opts ...grpc.CallOption) (*DeductStockResponse, error)
}

type inventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryServiceClient(cc grpc.ClientConnInterface) InventoryServiceClient {
	return &inventoryServiceClient{cc}
}

func (c *inventoryServiceClient) CheckStock(ctx context.Context, in *CheckStockRequest, opts ...grpc.CallOption) (*CheckStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckStockResponse)
	err := c.cc.Invoke(ctx, InventoryService_CheckStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) DeductStock(ctx context.Context, in *DeductStockRequest, opts ...grpc.CallOption) (*DeductStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeductStockResponse)
	err := c.cc.Invoke(ctx, InventoryService_DeductStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type InventoryServiceServer interface {
	CheckStock(context.Context, *CheckStockRequest) (*CheckStockResponse, error)
	DeductStock(context.Context, *DeductStockRequest) (*DeductStockResponse, error)
	mustEmbedUnimplementedInventoryServiceServer()
}

type UnimplementedInventoryServiceServer struct{}

func (UnimplementedInventoryServiceServer) CheckStock(context.Context, *CheckStockRequest) (*CheckStockResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method CheckStock not implemented")
}
func (UnimplementedInventoryServiceServer) DeductStock(context.Context, *DeductStockRequest) (*DeductStockResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method DeductStock not implemented")
}
func (UnimplementedInventoryServiceServer) mustEmbedUnimplementedInventoryServiceServer() {}
func (UnimplementedInventoryServiceServer) testEmbeddedByValue()                          {}

type UnsafeInventoryServiceServer interface {
	mustEmbedUnimplementedInventoryServiceServer()
}

func RegisterInventoryServiceServer(s grpc.ServiceRegistrar, srv InventoryServiceServer) {

	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InventoryService_ServiceDesc, srv)
}

func _InventoryService_CheckStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).CheckStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_CheckStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).CheckStock(ctx, req.(*CheckStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_DeductStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeductStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).DeductStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_DeductStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).DeductStock(ctx, req.(*DeductStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var InventoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.InventoryService",
	HandlerType: (*InventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckStock",
			Handler:    _InventoryService_CheckStock_Handler,
		},
		{
			MethodName: "DeductStock",
			Handler:    _InventoryService_DeductStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory/inventory.proto",
}
