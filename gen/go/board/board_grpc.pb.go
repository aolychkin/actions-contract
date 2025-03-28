// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: board.proto

package board_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BoardService_GetBoard_FullMethodName = "/board.BoardService/GetBoard"
)

// BoardServiceClient is the client API for BoardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoardServiceClient interface {
	GetBoard(ctx context.Context, in *GetBoardRequest, opts ...grpc.CallOption) (*GetBoardResponse, error)
}

type boardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBoardServiceClient(cc grpc.ClientConnInterface) BoardServiceClient {
	return &boardServiceClient{cc}
}

func (c *boardServiceClient) GetBoard(ctx context.Context, in *GetBoardRequest, opts ...grpc.CallOption) (*GetBoardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBoardResponse)
	err := c.cc.Invoke(ctx, BoardService_GetBoard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoardServiceServer is the server API for BoardService service.
// All implementations must embed UnimplementedBoardServiceServer
// for forward compatibility.
type BoardServiceServer interface {
	GetBoard(context.Context, *GetBoardRequest) (*GetBoardResponse, error)
	mustEmbedUnimplementedBoardServiceServer()
}

// UnimplementedBoardServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBoardServiceServer struct{}

func (UnimplementedBoardServiceServer) GetBoard(context.Context, *GetBoardRequest) (*GetBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBoard not implemented")
}
func (UnimplementedBoardServiceServer) mustEmbedUnimplementedBoardServiceServer() {}
func (UnimplementedBoardServiceServer) testEmbeddedByValue()                      {}

// UnsafeBoardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoardServiceServer will
// result in compilation errors.
type UnsafeBoardServiceServer interface {
	mustEmbedUnimplementedBoardServiceServer()
}

func RegisterBoardServiceServer(s grpc.ServiceRegistrar, srv BoardServiceServer) {
	// If the following call pancis, it indicates UnimplementedBoardServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BoardService_ServiceDesc, srv)
}

func _BoardService_GetBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).GetBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BoardService_GetBoard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).GetBoard(ctx, req.(*GetBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BoardService_ServiceDesc is the grpc.ServiceDesc for BoardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BoardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "board.BoardService",
	HandlerType: (*BoardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBoard",
			Handler:    _BoardService_GetBoard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "board.proto",
}
