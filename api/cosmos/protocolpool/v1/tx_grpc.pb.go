// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cosmos/protocolpool/v1/tx.proto

package protocolpoolv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_FundCommunityPool_FullMethodName            = "/cosmos.protocolpool.v1.Msg/FundCommunityPool"
	Msg_CommunityPoolSpend_FullMethodName           = "/cosmos.protocolpool.v1.Msg/CommunityPoolSpend"
	Msg_SubmitBudgetProposal_FullMethodName         = "/cosmos.protocolpool.v1.Msg/SubmitBudgetProposal"
	Msg_ClaimBudget_FullMethodName                  = "/cosmos.protocolpool.v1.Msg/ClaimBudget"
	Msg_AddContinuousFunds_FullMethodName           = "/cosmos.protocolpool.v1.Msg/AddContinuousFunds"
	Msg_FundDispensationProposal_FullMethodName     = "/cosmos.protocolpool.v1.Msg/FundDispensationProposal"
	Msg_CancelContinuousFundProposal_FullMethodName = "/cosmos.protocolpool.v1.Msg/CancelContinuousFundProposal"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// FundCommunityPool defines a method to allow an account to directly
	// fund the community pool.
	FundCommunityPool(ctx context.Context, in *MsgFundCommunityPool, opts ...grpc.CallOption) (*MsgFundCommunityPoolResponse, error)
	// CommunityPoolSpend defines a governance operation for sending tokens from
	// the community pool in the x/protocolpool module to another account, which
	// could be the governance module itself. The authority is defined in the
	// keeper.
	CommunityPoolSpend(ctx context.Context, in *MsgCommunityPoolSpend, opts ...grpc.CallOption) (*MsgCommunityPoolSpendResponse, error)
	// SubmitBudgetProposal defines a method to set a budget proposal.
	SubmitBudgetProposal(ctx context.Context, in *MsgSubmitBudgetProposal, opts ...grpc.CallOption) (*MsgSubmitBudgetProposalResponse, error)
	// ClaimBudget defines a method to claim the distributed budget.
	ClaimBudget(ctx context.Context, in *MsgClaimBudget, opts ...grpc.CallOption) (*MsgClaimBudgetResponse, error)
	// AddContinuousFunds defines a method to add funds continuously.
	AddContinuousFunds(ctx context.Context, in *MsgAddContinuousFunds, opts ...grpc.CallOption) (*MsgAddContinuousFundsResponse, error)
	// FundDispensationProposal defines a method for funds dispensation.
	FundDispensationProposal(ctx context.Context, in *MsgFundDispensationProposal, opts ...grpc.CallOption) (*MsgFundDispensationProposalResponse, error)
	// CancelContinuousFundProposal defines a method for cancelling continuous fund proposal.
	CancelContinuousFundProposal(ctx context.Context, in *MsgCancelContinuousFundProposal, opts ...grpc.CallOption) (*MsgCancelContinuousFundProposalResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) FundCommunityPool(ctx context.Context, in *MsgFundCommunityPool, opts ...grpc.CallOption) (*MsgFundCommunityPoolResponse, error) {
	out := new(MsgFundCommunityPoolResponse)
	err := c.cc.Invoke(ctx, Msg_FundCommunityPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CommunityPoolSpend(ctx context.Context, in *MsgCommunityPoolSpend, opts ...grpc.CallOption) (*MsgCommunityPoolSpendResponse, error) {
	out := new(MsgCommunityPoolSpendResponse)
	err := c.cc.Invoke(ctx, Msg_CommunityPoolSpend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitBudgetProposal(ctx context.Context, in *MsgSubmitBudgetProposal, opts ...grpc.CallOption) (*MsgSubmitBudgetProposalResponse, error) {
	out := new(MsgSubmitBudgetProposalResponse)
	err := c.cc.Invoke(ctx, Msg_SubmitBudgetProposal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClaimBudget(ctx context.Context, in *MsgClaimBudget, opts ...grpc.CallOption) (*MsgClaimBudgetResponse, error) {
	out := new(MsgClaimBudgetResponse)
	err := c.cc.Invoke(ctx, Msg_ClaimBudget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddContinuousFunds(ctx context.Context, in *MsgAddContinuousFunds, opts ...grpc.CallOption) (*MsgAddContinuousFundsResponse, error) {
	out := new(MsgAddContinuousFundsResponse)
	err := c.cc.Invoke(ctx, Msg_AddContinuousFunds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) FundDispensationProposal(ctx context.Context, in *MsgFundDispensationProposal, opts ...grpc.CallOption) (*MsgFundDispensationProposalResponse, error) {
	out := new(MsgFundDispensationProposalResponse)
	err := c.cc.Invoke(ctx, Msg_FundDispensationProposal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelContinuousFundProposal(ctx context.Context, in *MsgCancelContinuousFundProposal, opts ...grpc.CallOption) (*MsgCancelContinuousFundProposalResponse, error) {
	out := new(MsgCancelContinuousFundProposalResponse)
	err := c.cc.Invoke(ctx, Msg_CancelContinuousFundProposal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// FundCommunityPool defines a method to allow an account to directly
	// fund the community pool.
	FundCommunityPool(context.Context, *MsgFundCommunityPool) (*MsgFundCommunityPoolResponse, error)
	// CommunityPoolSpend defines a governance operation for sending tokens from
	// the community pool in the x/protocolpool module to another account, which
	// could be the governance module itself. The authority is defined in the
	// keeper.
	CommunityPoolSpend(context.Context, *MsgCommunityPoolSpend) (*MsgCommunityPoolSpendResponse, error)
	// SubmitBudgetProposal defines a method to set a budget proposal.
	SubmitBudgetProposal(context.Context, *MsgSubmitBudgetProposal) (*MsgSubmitBudgetProposalResponse, error)
	// ClaimBudget defines a method to claim the distributed budget.
	ClaimBudget(context.Context, *MsgClaimBudget) (*MsgClaimBudgetResponse, error)
	// AddContinuousFunds defines a method to add funds continuously.
	AddContinuousFunds(context.Context, *MsgAddContinuousFunds) (*MsgAddContinuousFundsResponse, error)
	// FundDispensationProposal defines a method for funds dispensation.
	FundDispensationProposal(context.Context, *MsgFundDispensationProposal) (*MsgFundDispensationProposalResponse, error)
	// CancelContinuousFundProposal defines a method for cancelling continuous fund proposal.
	CancelContinuousFundProposal(context.Context, *MsgCancelContinuousFundProposal) (*MsgCancelContinuousFundProposalResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) FundCommunityPool(context.Context, *MsgFundCommunityPool) (*MsgFundCommunityPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FundCommunityPool not implemented")
}
func (UnimplementedMsgServer) CommunityPoolSpend(context.Context, *MsgCommunityPoolSpend) (*MsgCommunityPoolSpendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommunityPoolSpend not implemented")
}
func (UnimplementedMsgServer) SubmitBudgetProposal(context.Context, *MsgSubmitBudgetProposal) (*MsgSubmitBudgetProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitBudgetProposal not implemented")
}
func (UnimplementedMsgServer) ClaimBudget(context.Context, *MsgClaimBudget) (*MsgClaimBudgetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimBudget not implemented")
}
func (UnimplementedMsgServer) AddContinuousFunds(context.Context, *MsgAddContinuousFunds) (*MsgAddContinuousFundsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddContinuousFunds not implemented")
}
func (UnimplementedMsgServer) FundDispensationProposal(context.Context, *MsgFundDispensationProposal) (*MsgFundDispensationProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FundDispensationProposal not implemented")
}
func (UnimplementedMsgServer) CancelContinuousFundProposal(context.Context, *MsgCancelContinuousFundProposal) (*MsgCancelContinuousFundProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelContinuousFundProposal not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_FundCommunityPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFundCommunityPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FundCommunityPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_FundCommunityPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FundCommunityPool(ctx, req.(*MsgFundCommunityPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CommunityPoolSpend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCommunityPoolSpend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CommunityPoolSpend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CommunityPoolSpend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CommunityPoolSpend(ctx, req.(*MsgCommunityPoolSpend))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitBudgetProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitBudgetProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitBudgetProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SubmitBudgetProposal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitBudgetProposal(ctx, req.(*MsgSubmitBudgetProposal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClaimBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimBudget)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ClaimBudget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimBudget(ctx, req.(*MsgClaimBudget))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddContinuousFunds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddContinuousFunds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddContinuousFunds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddContinuousFunds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddContinuousFunds(ctx, req.(*MsgAddContinuousFunds))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_FundDispensationProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFundDispensationProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FundDispensationProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_FundDispensationProposal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FundDispensationProposal(ctx, req.(*MsgFundDispensationProposal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelContinuousFundProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelContinuousFundProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelContinuousFundProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CancelContinuousFundProposal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelContinuousFundProposal(ctx, req.(*MsgCancelContinuousFundProposal))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.protocolpool.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FundCommunityPool",
			Handler:    _Msg_FundCommunityPool_Handler,
		},
		{
			MethodName: "CommunityPoolSpend",
			Handler:    _Msg_CommunityPoolSpend_Handler,
		},
		{
			MethodName: "SubmitBudgetProposal",
			Handler:    _Msg_SubmitBudgetProposal_Handler,
		},
		{
			MethodName: "ClaimBudget",
			Handler:    _Msg_ClaimBudget_Handler,
		},
		{
			MethodName: "AddContinuousFunds",
			Handler:    _Msg_AddContinuousFunds_Handler,
		},
		{
			MethodName: "FundDispensationProposal",
			Handler:    _Msg_FundDispensationProposal_Handler,
		},
		{
			MethodName: "CancelContinuousFundProposal",
			Handler:    _Msg_CancelContinuousFundProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/protocolpool/v1/tx.proto",
}
