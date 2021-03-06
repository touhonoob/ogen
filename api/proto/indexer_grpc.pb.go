// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// IndexerClient is the client API for Indexer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IndexerClient interface {
	GetIndexerStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IndexerStatus, error)
	GetBlockHash(ctx context.Context, in *Number, opts ...grpc.CallOption) (*Hash, error)
	GetTransaction(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*Tx, error)
	GetDetailedTransaction(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*TxDetail, error)
	GetAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*AccountDetail, error)
	GetBlock(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*Block, error)
	GetBlockTxs(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*BlockTxs, error)
}

type indexerClient struct {
	cc grpc.ClientConnInterface
}

func NewIndexerClient(cc grpc.ClientConnInterface) IndexerClient {
	return &indexerClient{cc}
}

func (c *indexerClient) GetIndexerStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IndexerStatus, error) {
	out := new(IndexerStatus)
	err := c.cc.Invoke(ctx, "/Indexer/GetIndexerStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexerClient) GetBlockHash(ctx context.Context, in *Number, opts ...grpc.CallOption) (*Hash, error) {
	out := new(Hash)
	err := c.cc.Invoke(ctx, "/Indexer/GetBlockHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexerClient) GetTransaction(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*Tx, error) {
	out := new(Tx)
	err := c.cc.Invoke(ctx, "/Indexer/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexerClient) GetDetailedTransaction(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*TxDetail, error) {
	out := new(TxDetail)
	err := c.cc.Invoke(ctx, "/Indexer/GetDetailedTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexerClient) GetAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*AccountDetail, error) {
	out := new(AccountDetail)
	err := c.cc.Invoke(ctx, "/Indexer/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexerClient) GetBlock(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/Indexer/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexerClient) GetBlockTxs(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*BlockTxs, error) {
	out := new(BlockTxs)
	err := c.cc.Invoke(ctx, "/Indexer/GetBlockTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexerServer is the server API for Indexer service.
// All implementations must embed UnimplementedIndexerServer
// for forward compatibility
type IndexerServer interface {
	GetIndexerStatus(context.Context, *Empty) (*IndexerStatus, error)
	GetBlockHash(context.Context, *Number) (*Hash, error)
	GetTransaction(context.Context, *Hash) (*Tx, error)
	GetDetailedTransaction(context.Context, *Hash) (*TxDetail, error)
	GetAccount(context.Context, *Account) (*AccountDetail, error)
	GetBlock(context.Context, *Hash) (*Block, error)
	GetBlockTxs(context.Context, *Hash) (*BlockTxs, error)
	mustEmbedUnimplementedIndexerServer()
}

// UnimplementedIndexerServer must be embedded to have forward compatible implementations.
type UnimplementedIndexerServer struct {
}

func (UnimplementedIndexerServer) GetIndexerStatus(context.Context, *Empty) (*IndexerStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndexerStatus not implemented")
}
func (UnimplementedIndexerServer) GetBlockHash(context.Context, *Number) (*Hash, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHash not implemented")
}
func (UnimplementedIndexerServer) GetTransaction(context.Context, *Hash) (*Tx, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedIndexerServer) GetDetailedTransaction(context.Context, *Hash) (*TxDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetailedTransaction not implemented")
}
func (UnimplementedIndexerServer) GetAccount(context.Context, *Account) (*AccountDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedIndexerServer) GetBlock(context.Context, *Hash) (*Block, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (UnimplementedIndexerServer) GetBlockTxs(context.Context, *Hash) (*BlockTxs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockTxs not implemented")
}
func (UnimplementedIndexerServer) mustEmbedUnimplementedIndexerServer() {}

// UnsafeIndexerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndexerServer will
// result in compilation errors.
type UnsafeIndexerServer interface {
	mustEmbedUnimplementedIndexerServer()
}

func RegisterIndexerServer(s *grpc.Server, srv IndexerServer) {
	s.RegisterService(&_Indexer_serviceDesc, srv)
}

func _Indexer_GetIndexerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).GetIndexerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Indexer/GetIndexerStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).GetIndexerStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexer_GetBlockHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Number)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).GetBlockHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Indexer/GetBlockHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).GetBlockHash(ctx, req.(*Number))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexer_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Indexer/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).GetTransaction(ctx, req.(*Hash))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexer_GetDetailedTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).GetDetailedTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Indexer/GetDetailedTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).GetDetailedTransaction(ctx, req.(*Hash))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexer_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Indexer/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).GetAccount(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexer_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Indexer/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).GetBlock(ctx, req.(*Hash))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexer_GetBlockTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).GetBlockTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Indexer/GetBlockTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).GetBlockTxs(ctx, req.(*Hash))
	}
	return interceptor(ctx, in, info, handler)
}

var _Indexer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Indexer",
	HandlerType: (*IndexerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIndexerStatus",
			Handler:    _Indexer_GetIndexerStatus_Handler,
		},
		{
			MethodName: "GetBlockHash",
			Handler:    _Indexer_GetBlockHash_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _Indexer_GetTransaction_Handler,
		},
		{
			MethodName: "GetDetailedTransaction",
			Handler:    _Indexer_GetDetailedTransaction_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _Indexer_GetAccount_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _Indexer_GetBlock_Handler,
		},
		{
			MethodName: "GetBlockTxs",
			Handler:    _Indexer_GetBlockTxs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "indexer.proto",
}
