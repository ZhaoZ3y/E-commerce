package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"gomall/app/frontend/infra/rpc"
	frontendUtils "gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/check_out"
	"gomall/rpc_gen/kitex_gen/payment"

	"github.com/cloudwego/hertz/pkg/app"
	checkout_page "gomall/app/frontend/hertz_gen/frontend/checkout_page"
)

type CheckOutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckOutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckOutWaitingService {
	return &CheckOutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckOutWaitingService) Run(req *checkout_page.CheckoutReq) (resp map[string]any, err error) {
	// todo use checkout svc api
	userId := frontendUtils.GetUserIDFromCtx(h.Context)
	_, err = rpc.CheckoutClient.CheccOut(h.Context, &check_out.CheckOutReq{
		UserId:    int64(userId),
		Email:     req.Email,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Address: &check_out.Address{
			Country:       req.Country,
			City:          req.City,
			ZipCode:       req.Zipcode,
			State:         req.Province,
			StreetAddress: req.Street,
		},
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CartNum,
			CreditCardCvv:             req.Cvv,
			CreditCardExpirationMonth: req.ExpirationMonth,
			CreditCardExpirationYear:  req.ExpirationYear,
		},
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"Title":    "Waiting",
		"redirect": "/checkout/result",
	}, nil
}
