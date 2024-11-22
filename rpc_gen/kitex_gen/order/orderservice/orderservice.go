// Code generated by Kitex v0.9.1. DO NOT EDIT.

package orderservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	order "gomall/rpc_gen/kitex_gen/order"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"PlaceOrder": kitex.NewMethodInfo(
		placeOrderHandler,
		newOrderServicePlaceOrderArgs,
		newOrderServicePlaceOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ListOrders": kitex.NewMethodInfo(
		listOrdersHandler,
		newOrderServiceListOrdersArgs,
		newOrderServiceListOrdersResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	orderServiceServiceInfo                = NewServiceInfo()
	orderServiceServiceInfoForClient       = NewServiceInfoForClient()
	orderServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return orderServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return orderServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return orderServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "OrderService"
	handlerType := (*order.OrderService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "order",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func placeOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderServicePlaceOrderArgs)
	realResult := result.(*order.OrderServicePlaceOrderResult)
	success, err := handler.(order.OrderService).PlaceOrder(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderServicePlaceOrderArgs() interface{} {
	return order.NewOrderServicePlaceOrderArgs()
}

func newOrderServicePlaceOrderResult() interface{} {
	return order.NewOrderServicePlaceOrderResult()
}

func listOrdersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderServiceListOrdersArgs)
	realResult := result.(*order.OrderServiceListOrdersResult)
	success, err := handler.(order.OrderService).ListOrders(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderServiceListOrdersArgs() interface{} {
	return order.NewOrderServiceListOrdersArgs()
}

func newOrderServiceListOrdersResult() interface{} {
	return order.NewOrderServiceListOrdersResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (r *order.PlaceOrderResp, err error) {
	var _args order.OrderServicePlaceOrderArgs
	_args.Req = req
	var _result order.OrderServicePlaceOrderResult
	if err = p.c.Call(ctx, "PlaceOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListOrders(ctx context.Context, req *order.ListOrdersReq) (r *order.ListOrdersResp, err error) {
	var _args order.OrderServiceListOrdersArgs
	_args.Req = req
	var _result order.OrderServiceListOrdersResult
	if err = p.c.Call(ctx, "ListOrders", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}