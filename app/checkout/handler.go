package main

import (
	"context"
	"gomall/app/checkout/biz/service"
	check_out "gomall/rpc_gen/kitex_gen/check_out"
)

// CheckOutServiceImpl implements the last service interface defined in the IDL.
type CheckOutServiceImpl struct{}

// CheccOut implements the CheckOutServiceImpl interface.
func (s *CheckOutServiceImpl) CheccOut(ctx context.Context, req *check_out.CheckOutReq) (resp *check_out.CheckOutResp, err error) {
	resp, err = service.NewCheccOutService(ctx).Run(req)

	return resp, err
}
