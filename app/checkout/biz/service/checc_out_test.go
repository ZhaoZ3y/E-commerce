package service

import (
	"context"
	check_out "gomall/rpc_gen/kitex_gen/check_out"
	"testing"
)

func TestCheccOut_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCheccOutService(ctx)
	// init req and assert value

	req := &check_out.CheckOutReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
