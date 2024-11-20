package service

import (
	"context"
	"gomall/app/frontend/infra/rpc"
	frontendUtils "gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/hertz/pkg/app"
	cart_page "gomall/app/frontend/hertz_gen/frontend/cart_page"
	common "gomall/app/frontend/hertz_gen/frontend/common"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart_page.AddCartItemReq) (resp *common.Empty, err error) {
	// todo use cart SVC api
	_, err = rpc.CartClient.AddItem(h.Context, &cart.AddItemReq{
		UserId: int64(frontendUtils.GetUserIDFromCtx(h.Context)),
		Item: &cart.CartItem{
			ProductId: req.ProductID,
			Quantity:  int64(req.ProductNum),
		},
	})
	if err != nil {
		return nil, err
	}
	return
}
