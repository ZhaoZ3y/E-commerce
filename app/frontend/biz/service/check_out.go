package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"gomall/app/frontend/infra/rpc"
	frontendUtils "gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/cart"
	"gomall/rpc_gen/kitex_gen/product"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	common "gomall/app/frontend/hertz_gen/frontend/common"
)

type CheckOutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckOutService(Context context.Context, RequestContext *app.RequestContext) *CheckOutService {
	return &CheckOutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckOutService) Run(req *common.Empty) (resp map[string]any, err error) {
	// todo checkout svc api
	var items []map[string]string
	userId := frontendUtils.GetUserIDFromCtx(h.Context)

	carts, err := rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{UserId: int64(userId)})
	if err != nil {
		return nil, err
	}

	var total float32
	for _, item := range carts.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: item.ProductId})
		if err != nil {
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":     p.Name,
			"Price":    strconv.FormatFloat(p.Price, 'f', 2, 64),
			"Picture":  p.Picture,
			"Quantity": strconv.Itoa(int(item.Quantity)),
		})
		total += float32(p.Price) * float32(item.Quantity)
	}

	return utils.H{
		"Title": "Checkout",
		"Items": items,
		"Total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
