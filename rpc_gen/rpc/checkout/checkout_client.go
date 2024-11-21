package checkout

import (
	"context"
	check_out "gomall/rpc_gen/kitex_gen/check_out"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"gomall/rpc_gen/kitex_gen/check_out/checkoutservice"
)

type RPCClient interface {
	KitexClient() checkoutservice.Client
	Service() string
	CheccOut(ctx context.Context, req *check_out.CheckOutReq, callOptions ...callopt.Option) (r *check_out.CheckOutResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := checkoutservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient checkoutservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() checkoutservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CheccOut(ctx context.Context, req *check_out.CheckOutReq, callOptions ...callopt.Option) (r *check_out.CheckOutResp, err error) {
	return c.kitexClient.CheccOut(ctx, req, callOptions...)
}
