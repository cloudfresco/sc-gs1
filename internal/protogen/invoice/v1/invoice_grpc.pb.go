// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: invoice/v1/invoice.proto

package v1

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
	InvoiceService_CreateInvoice_FullMethodName                              = "/invoice.v1.InvoiceService/CreateInvoice"
	InvoiceService_GetInvoices_FullMethodName                                = "/invoice.v1.InvoiceService/GetInvoices"
	InvoiceService_GetInvoice_FullMethodName                                 = "/invoice.v1.InvoiceService/GetInvoice"
	InvoiceService_GetInvoiceByPk_FullMethodName                             = "/invoice.v1.InvoiceService/GetInvoiceByPk"
	InvoiceService_UpdateInvoice_FullMethodName                              = "/invoice.v1.InvoiceService/UpdateInvoice"
	InvoiceService_CreateInvoiceLineItem_FullMethodName                      = "/invoice.v1.InvoiceService/CreateInvoiceLineItem"
	InvoiceService_GetInvoiceLineItems_FullMethodName                        = "/invoice.v1.InvoiceService/GetInvoiceLineItems"
	InvoiceService_CreateInvoiceLineItemInformationAfterTaxes_FullMethodName = "/invoice.v1.InvoiceService/CreateInvoiceLineItemInformationAfterTaxes"
	InvoiceService_CreateInvoiceTotal_FullMethodName                         = "/invoice.v1.InvoiceService/CreateInvoiceTotal"
	InvoiceService_CreateInvoiceAllowanceCharge_FullMethodName               = "/invoice.v1.InvoiceService/CreateInvoiceAllowanceCharge"
	InvoiceService_CreateLeviedDutyFeeTax_FullMethodName                     = "/invoice.v1.InvoiceService/CreateLeviedDutyFeeTax"
)

// InvoiceServiceClient is the client API for InvoiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The InvoiceService service definition.
type InvoiceServiceClient interface {
	CreateInvoice(ctx context.Context, in *CreateInvoiceRequest, opts ...grpc.CallOption) (*CreateInvoiceResponse, error)
	GetInvoices(ctx context.Context, in *GetInvoicesRequest, opts ...grpc.CallOption) (*GetInvoicesResponse, error)
	GetInvoice(ctx context.Context, in *GetInvoiceRequest, opts ...grpc.CallOption) (*GetInvoiceResponse, error)
	GetInvoiceByPk(ctx context.Context, in *GetInvoiceByPkRequest, opts ...grpc.CallOption) (*GetInvoiceByPkResponse, error)
	UpdateInvoice(ctx context.Context, in *UpdateInvoiceRequest, opts ...grpc.CallOption) (*UpdateInvoiceResponse, error)
	CreateInvoiceLineItem(ctx context.Context, in *CreateInvoiceLineItemRequest, opts ...grpc.CallOption) (*CreateInvoiceLineItemResponse, error)
	GetInvoiceLineItems(ctx context.Context, in *GetInvoiceLineItemsRequest, opts ...grpc.CallOption) (*GetInvoiceLineItemsResponse, error)
	CreateInvoiceLineItemInformationAfterTaxes(ctx context.Context, in *CreateInvoiceLineItemInformationAfterTaxesRequest, opts ...grpc.CallOption) (*CreateInvoiceLineItemInformationAfterTaxesResponse, error)
	CreateInvoiceTotal(ctx context.Context, in *CreateInvoiceTotalRequest, opts ...grpc.CallOption) (*CreateInvoiceTotalResponse, error)
	CreateInvoiceAllowanceCharge(ctx context.Context, in *CreateInvoiceAllowanceChargeRequest, opts ...grpc.CallOption) (*CreateInvoiceAllowanceChargeResponse, error)
	CreateLeviedDutyFeeTax(ctx context.Context, in *CreateLeviedDutyFeeTaxRequest, opts ...grpc.CallOption) (*CreateLeviedDutyFeeTaxResponse, error)
}

type invoiceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoiceServiceClient(cc grpc.ClientConnInterface) InvoiceServiceClient {
	return &invoiceServiceClient{cc}
}

func (c *invoiceServiceClient) CreateInvoice(ctx context.Context, in *CreateInvoiceRequest, opts ...grpc.CallOption) (*CreateInvoiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateInvoiceResponse)
	err := c.cc.Invoke(ctx, InvoiceService_CreateInvoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) GetInvoices(ctx context.Context, in *GetInvoicesRequest, opts ...grpc.CallOption) (*GetInvoicesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInvoicesResponse)
	err := c.cc.Invoke(ctx, InvoiceService_GetInvoices_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) GetInvoice(ctx context.Context, in *GetInvoiceRequest, opts ...grpc.CallOption) (*GetInvoiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInvoiceResponse)
	err := c.cc.Invoke(ctx, InvoiceService_GetInvoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) GetInvoiceByPk(ctx context.Context, in *GetInvoiceByPkRequest, opts ...grpc.CallOption) (*GetInvoiceByPkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInvoiceByPkResponse)
	err := c.cc.Invoke(ctx, InvoiceService_GetInvoiceByPk_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) UpdateInvoice(ctx context.Context, in *UpdateInvoiceRequest, opts ...grpc.CallOption) (*UpdateInvoiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateInvoiceResponse)
	err := c.cc.Invoke(ctx, InvoiceService_UpdateInvoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) CreateInvoiceLineItem(ctx context.Context, in *CreateInvoiceLineItemRequest, opts ...grpc.CallOption) (*CreateInvoiceLineItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateInvoiceLineItemResponse)
	err := c.cc.Invoke(ctx, InvoiceService_CreateInvoiceLineItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) GetInvoiceLineItems(ctx context.Context, in *GetInvoiceLineItemsRequest, opts ...grpc.CallOption) (*GetInvoiceLineItemsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInvoiceLineItemsResponse)
	err := c.cc.Invoke(ctx, InvoiceService_GetInvoiceLineItems_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) CreateInvoiceLineItemInformationAfterTaxes(ctx context.Context, in *CreateInvoiceLineItemInformationAfterTaxesRequest, opts ...grpc.CallOption) (*CreateInvoiceLineItemInformationAfterTaxesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateInvoiceLineItemInformationAfterTaxesResponse)
	err := c.cc.Invoke(ctx, InvoiceService_CreateInvoiceLineItemInformationAfterTaxes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) CreateInvoiceTotal(ctx context.Context, in *CreateInvoiceTotalRequest, opts ...grpc.CallOption) (*CreateInvoiceTotalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateInvoiceTotalResponse)
	err := c.cc.Invoke(ctx, InvoiceService_CreateInvoiceTotal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) CreateInvoiceAllowanceCharge(ctx context.Context, in *CreateInvoiceAllowanceChargeRequest, opts ...grpc.CallOption) (*CreateInvoiceAllowanceChargeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateInvoiceAllowanceChargeResponse)
	err := c.cc.Invoke(ctx, InvoiceService_CreateInvoiceAllowanceCharge_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) CreateLeviedDutyFeeTax(ctx context.Context, in *CreateLeviedDutyFeeTaxRequest, opts ...grpc.CallOption) (*CreateLeviedDutyFeeTaxResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateLeviedDutyFeeTaxResponse)
	err := c.cc.Invoke(ctx, InvoiceService_CreateLeviedDutyFeeTax_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvoiceServiceServer is the server API for InvoiceService service.
// All implementations must embed UnimplementedInvoiceServiceServer
// for forward compatibility.
//
// The InvoiceService service definition.
type InvoiceServiceServer interface {
	CreateInvoice(context.Context, *CreateInvoiceRequest) (*CreateInvoiceResponse, error)
	GetInvoices(context.Context, *GetInvoicesRequest) (*GetInvoicesResponse, error)
	GetInvoice(context.Context, *GetInvoiceRequest) (*GetInvoiceResponse, error)
	GetInvoiceByPk(context.Context, *GetInvoiceByPkRequest) (*GetInvoiceByPkResponse, error)
	UpdateInvoice(context.Context, *UpdateInvoiceRequest) (*UpdateInvoiceResponse, error)
	CreateInvoiceLineItem(context.Context, *CreateInvoiceLineItemRequest) (*CreateInvoiceLineItemResponse, error)
	GetInvoiceLineItems(context.Context, *GetInvoiceLineItemsRequest) (*GetInvoiceLineItemsResponse, error)
	CreateInvoiceLineItemInformationAfterTaxes(context.Context, *CreateInvoiceLineItemInformationAfterTaxesRequest) (*CreateInvoiceLineItemInformationAfterTaxesResponse, error)
	CreateInvoiceTotal(context.Context, *CreateInvoiceTotalRequest) (*CreateInvoiceTotalResponse, error)
	CreateInvoiceAllowanceCharge(context.Context, *CreateInvoiceAllowanceChargeRequest) (*CreateInvoiceAllowanceChargeResponse, error)
	CreateLeviedDutyFeeTax(context.Context, *CreateLeviedDutyFeeTaxRequest) (*CreateLeviedDutyFeeTaxResponse, error)
	mustEmbedUnimplementedInvoiceServiceServer()
}

// UnimplementedInvoiceServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInvoiceServiceServer struct{}

func (UnimplementedInvoiceServiceServer) CreateInvoice(context.Context, *CreateInvoiceRequest) (*CreateInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvoice not implemented")
}
func (UnimplementedInvoiceServiceServer) GetInvoices(context.Context, *GetInvoicesRequest) (*GetInvoicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoices not implemented")
}
func (UnimplementedInvoiceServiceServer) GetInvoice(context.Context, *GetInvoiceRequest) (*GetInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoice not implemented")
}
func (UnimplementedInvoiceServiceServer) GetInvoiceByPk(context.Context, *GetInvoiceByPkRequest) (*GetInvoiceByPkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoiceByPk not implemented")
}
func (UnimplementedInvoiceServiceServer) UpdateInvoice(context.Context, *UpdateInvoiceRequest) (*UpdateInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInvoice not implemented")
}
func (UnimplementedInvoiceServiceServer) CreateInvoiceLineItem(context.Context, *CreateInvoiceLineItemRequest) (*CreateInvoiceLineItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvoiceLineItem not implemented")
}
func (UnimplementedInvoiceServiceServer) GetInvoiceLineItems(context.Context, *GetInvoiceLineItemsRequest) (*GetInvoiceLineItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoiceLineItems not implemented")
}
func (UnimplementedInvoiceServiceServer) CreateInvoiceLineItemInformationAfterTaxes(context.Context, *CreateInvoiceLineItemInformationAfterTaxesRequest) (*CreateInvoiceLineItemInformationAfterTaxesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvoiceLineItemInformationAfterTaxes not implemented")
}
func (UnimplementedInvoiceServiceServer) CreateInvoiceTotal(context.Context, *CreateInvoiceTotalRequest) (*CreateInvoiceTotalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvoiceTotal not implemented")
}
func (UnimplementedInvoiceServiceServer) CreateInvoiceAllowanceCharge(context.Context, *CreateInvoiceAllowanceChargeRequest) (*CreateInvoiceAllowanceChargeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvoiceAllowanceCharge not implemented")
}
func (UnimplementedInvoiceServiceServer) CreateLeviedDutyFeeTax(context.Context, *CreateLeviedDutyFeeTaxRequest) (*CreateLeviedDutyFeeTaxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLeviedDutyFeeTax not implemented")
}
func (UnimplementedInvoiceServiceServer) mustEmbedUnimplementedInvoiceServiceServer() {}
func (UnimplementedInvoiceServiceServer) testEmbeddedByValue()                        {}

// UnsafeInvoiceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvoiceServiceServer will
// result in compilation errors.
type UnsafeInvoiceServiceServer interface {
	mustEmbedUnimplementedInvoiceServiceServer()
}

func RegisterInvoiceServiceServer(s grpc.ServiceRegistrar, srv InvoiceServiceServer) {
	// If the following call pancis, it indicates UnimplementedInvoiceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InvoiceService_ServiceDesc, srv)
}

func _InvoiceService_CreateInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).CreateInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_CreateInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).CreateInvoice(ctx, req.(*CreateInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_GetInvoices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).GetInvoices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_GetInvoices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).GetInvoices(ctx, req.(*GetInvoicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_GetInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).GetInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_GetInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).GetInvoice(ctx, req.(*GetInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_GetInvoiceByPk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoiceByPkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).GetInvoiceByPk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_GetInvoiceByPk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).GetInvoiceByPk(ctx, req.(*GetInvoiceByPkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_UpdateInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).UpdateInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_UpdateInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).UpdateInvoice(ctx, req.(*UpdateInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_CreateInvoiceLineItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvoiceLineItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).CreateInvoiceLineItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_CreateInvoiceLineItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).CreateInvoiceLineItem(ctx, req.(*CreateInvoiceLineItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_GetInvoiceLineItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoiceLineItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).GetInvoiceLineItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_GetInvoiceLineItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).GetInvoiceLineItems(ctx, req.(*GetInvoiceLineItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_CreateInvoiceLineItemInformationAfterTaxes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvoiceLineItemInformationAfterTaxesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).CreateInvoiceLineItemInformationAfterTaxes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_CreateInvoiceLineItemInformationAfterTaxes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).CreateInvoiceLineItemInformationAfterTaxes(ctx, req.(*CreateInvoiceLineItemInformationAfterTaxesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_CreateInvoiceTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvoiceTotalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).CreateInvoiceTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_CreateInvoiceTotal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).CreateInvoiceTotal(ctx, req.(*CreateInvoiceTotalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_CreateInvoiceAllowanceCharge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvoiceAllowanceChargeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).CreateInvoiceAllowanceCharge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_CreateInvoiceAllowanceCharge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).CreateInvoiceAllowanceCharge(ctx, req.(*CreateInvoiceAllowanceChargeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_CreateLeviedDutyFeeTax_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLeviedDutyFeeTaxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).CreateLeviedDutyFeeTax(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_CreateLeviedDutyFeeTax_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).CreateLeviedDutyFeeTax(ctx, req.(*CreateLeviedDutyFeeTaxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InvoiceService_ServiceDesc is the grpc.ServiceDesc for InvoiceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InvoiceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "invoice.v1.InvoiceService",
	HandlerType: (*InvoiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInvoice",
			Handler:    _InvoiceService_CreateInvoice_Handler,
		},
		{
			MethodName: "GetInvoices",
			Handler:    _InvoiceService_GetInvoices_Handler,
		},
		{
			MethodName: "GetInvoice",
			Handler:    _InvoiceService_GetInvoice_Handler,
		},
		{
			MethodName: "GetInvoiceByPk",
			Handler:    _InvoiceService_GetInvoiceByPk_Handler,
		},
		{
			MethodName: "UpdateInvoice",
			Handler:    _InvoiceService_UpdateInvoice_Handler,
		},
		{
			MethodName: "CreateInvoiceLineItem",
			Handler:    _InvoiceService_CreateInvoiceLineItem_Handler,
		},
		{
			MethodName: "GetInvoiceLineItems",
			Handler:    _InvoiceService_GetInvoiceLineItems_Handler,
		},
		{
			MethodName: "CreateInvoiceLineItemInformationAfterTaxes",
			Handler:    _InvoiceService_CreateInvoiceLineItemInformationAfterTaxes_Handler,
		},
		{
			MethodName: "CreateInvoiceTotal",
			Handler:    _InvoiceService_CreateInvoiceTotal_Handler,
		},
		{
			MethodName: "CreateInvoiceAllowanceCharge",
			Handler:    _InvoiceService_CreateInvoiceAllowanceCharge_Handler,
		},
		{
			MethodName: "CreateLeviedDutyFeeTax",
			Handler:    _InvoiceService_CreateLeviedDutyFeeTax_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "invoice/v1/invoice.proto",
}
