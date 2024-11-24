// Code generated by Kitex v0.9.1. DO NOT EDIT.

package emailservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	email "gomall/rpc_gen/kitex_gen/email"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Send": kitex.NewMethodInfo(
		sendHandler,
		newEmailServiceSendArgs,
		newEmailServiceSendResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	emailServiceServiceInfo                = NewServiceInfo()
	emailServiceServiceInfoForClient       = NewServiceInfoForClient()
	emailServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return emailServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return emailServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return emailServiceServiceInfoForClient
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
	serviceName := "EmailService"
	handlerType := (*email.EmailService)(nil)
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
		"PackageName": "email",
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

func sendHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*email.EmailServiceSendArgs)
	realResult := result.(*email.EmailServiceSendResult)
	success, err := handler.(email.EmailService).Send(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newEmailServiceSendArgs() interface{} {
	return email.NewEmailServiceSendArgs()
}

func newEmailServiceSendResult() interface{} {
	return email.NewEmailServiceSendResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Send(ctx context.Context, req *email.EmailReq) (r *email.EmailResp, err error) {
	var _args email.EmailServiceSendArgs
	_args.Req = req
	var _result email.EmailServiceSendResult
	if err = p.c.Call(ctx, "Send", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
