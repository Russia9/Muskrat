// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: raidService.proto

package protobuf

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
	Raid_GetRaidInfo_FullMethodName = "/raidService.Raid/GetRaidInfo"
)

// RaidClient is the client API for Raid service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaidClient interface {
	GetRaidInfo(ctx context.Context, in *RaidInfo, opts ...grpc.CallOption) (*Null, error)
}

type raidClient struct {
	cc grpc.ClientConnInterface
}

func NewRaidClient(cc grpc.ClientConnInterface) RaidClient {
	return &raidClient{cc}
}

func (c *raidClient) GetRaidInfo(ctx context.Context, in *RaidInfo, opts ...grpc.CallOption) (*Null, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Null)
	err := c.cc.Invoke(ctx, Raid_GetRaidInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaidServer is the server API for Raid service.
// All implementations must embed UnimplementedRaidServer
// for forward compatibility.
type RaidServer interface {
	GetRaidInfo(context.Context, *RaidInfo) (*Null, error)
	mustEmbedUnimplementedRaidServer()
}

// UnimplementedRaidServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRaidServer struct{}

func (UnimplementedRaidServer) GetRaidInfo(context.Context, *RaidInfo) (*Null, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRaidInfo not implemented")
}
func (UnimplementedRaidServer) mustEmbedUnimplementedRaidServer() {}
func (UnimplementedRaidServer) testEmbeddedByValue()              {}

// UnsafeRaidServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaidServer will
// result in compilation errors.
type UnsafeRaidServer interface {
	mustEmbedUnimplementedRaidServer()
}

func RegisterRaidServer(s grpc.ServiceRegistrar, srv RaidServer) {
	// If the following call pancis, it indicates UnimplementedRaidServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Raid_ServiceDesc, srv)
}

func _Raid_GetRaidInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaidInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaidServer).GetRaidInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Raid_GetRaidInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaidServer).GetRaidInfo(ctx, req.(*RaidInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// Raid_ServiceDesc is the grpc.ServiceDesc for Raid service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Raid_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "raidService.Raid",
	HandlerType: (*RaidServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRaidInfo",
			Handler:    _Raid_GetRaidInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raidService.proto",
}