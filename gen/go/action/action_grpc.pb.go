// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: action.proto

package action_v1

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
	Action_GetAction_FullMethodName = "/action.Action/GetAction"
)

// ActionClient is the client API for Action service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActionClient interface {
	GetAction(ctx context.Context, in *GetActionRequest, opts ...grpc.CallOption) (*GetActionResponse, error)
}

type actionClient struct {
	cc grpc.ClientConnInterface
}

func NewActionClient(cc grpc.ClientConnInterface) ActionClient {
	return &actionClient{cc}
}

func (c *actionClient) GetAction(ctx context.Context, in *GetActionRequest, opts ...grpc.CallOption) (*GetActionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetActionResponse)
	err := c.cc.Invoke(ctx, Action_GetAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionServer is the server API for Action service.
// All implementations must embed UnimplementedActionServer
// for forward compatibility.
type ActionServer interface {
	GetAction(context.Context, *GetActionRequest) (*GetActionResponse, error)
	mustEmbedUnimplementedActionServer()
}

// UnimplementedActionServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedActionServer struct{}

func (UnimplementedActionServer) GetAction(context.Context, *GetActionRequest) (*GetActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAction not implemented")
}
func (UnimplementedActionServer) mustEmbedUnimplementedActionServer() {}
func (UnimplementedActionServer) testEmbeddedByValue()                {}

// UnsafeActionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActionServer will
// result in compilation errors.
type UnsafeActionServer interface {
	mustEmbedUnimplementedActionServer()
}

func RegisterActionServer(s grpc.ServiceRegistrar, srv ActionServer) {
	// If the following call pancis, it indicates UnimplementedActionServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Action_ServiceDesc, srv)
}

func _Action_GetAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServer).GetAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Action_GetAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServer).GetAction(ctx, req.(*GetActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Action_ServiceDesc is the grpc.ServiceDesc for Action service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Action_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "action.Action",
	HandlerType: (*ActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAction",
			Handler:    _Action_GetAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "action.proto",
}
