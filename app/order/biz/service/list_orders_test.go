package service

import (
	"context"
	order "gomall/rpc_gen/kitex_gen/order"
	"testing"
)

func TestListOrders_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListOrdersService(ctx)
	// init req and assert value

	req := &order.ListOrdersReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}