package checkout_page

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"gomall/app/frontend/biz/service"
	"gomall/app/frontend/biz/utils"
	checkout_page "gomall/app/frontend/hertz_gen/frontend/checkout_page"
	common "gomall/app/frontend/hertz_gen/frontend/common"
)

// CheckOut .
// @router /checkout [GET]
func CheckOut(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCheckOutService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "checkout", utils.WarpResponse(ctx, c, resp))
}

// CheckOutWaiting .
// @router /checkout/waiting [POST]
func CheckOutWaiting(ctx context.Context, c *app.RequestContext) {
	var err error
	var req checkout_page.CheckoutReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCheckOutWaitingService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "waiting", utils.WarpResponse(ctx, c, resp))
}

// CheckOutResult .
// @router /checkout/result [GET]
func CheckOutResult(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCheckOutResultService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "result", utils.WarpResponse(ctx, c, resp))
}
