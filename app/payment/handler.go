package main

import (
	"context"
	"gomall/app/payment/biz/service"
	payment "gomall/rpc_gen/kitex_gen/payment"
)

// PaymentImpl implements the last service interface defined in the IDL.
type PaymentImpl struct{}

// Charge implements the PaymentImpl interface.
func (s *PaymentImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	resp, err = service.NewChargeService(ctx).Run(req)

	return resp, err
}
