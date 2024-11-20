package cart

import (
	"context"
	cart "gomall/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"gomall/rpc_gen/kitex_gen/cart/cartservice"
)

type RPCClient interface {
	KitexClient() cartservice.Client
	Service() string
	AddItem(ctx context.Context, req *cart.AddItemReq, callOptions ...callopt.Option) (r *cart.AddItemResp, err error)
	GetCart(ctx context.Context, req *cart.GetCartReq, callOptions ...callopt.Option) (r *cart.GetCartResp, err error)
	EmptyCart(ctx context.Context, req *cart.EmptyCartReq, callOptions ...callopt.Option) (r *cart.EmptyCartResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := cartservice.NewClient(dstService, opts...)
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
	kitexClient cartservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() cartservice.Client {
	return c.kitexClient
}

func (c *clientImpl) AddItem(ctx context.Context, req *cart.AddItemReq, callOptions ...callopt.Option) (r *cart.AddItemResp, err error) {
	return c.kitexClient.AddItem(ctx, req, callOptions...)
}

func (c *clientImpl) GetCart(ctx context.Context, req *cart.GetCartReq, callOptions ...callopt.Option) (r *cart.GetCartResp, err error) {
	return c.kitexClient.GetCart(ctx, req, callOptions...)
}

func (c *clientImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartReq, callOptions ...callopt.Option) (r *cart.EmptyCartResp, err error) {
	return c.kitexClient.EmptyCart(ctx, req, callOptions...)
}
