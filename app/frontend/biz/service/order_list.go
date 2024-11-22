package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"gomall/app/frontend/infra/rpc"
	"gomall/app/frontend/types"
	frontendUtils "gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/order"
	"gomall/rpc_gen/kitex_gen/product"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	common "gomall/app/frontend/hertz_gen/frontend/common"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	// todo edit your code
	userId := frontendUtils.GetUserIDFromCtx(h.Context)
	orderResp, err := rpc.OrderClient.ListOrders(h.Context, &order.ListOrdersReq{UserId: int64(userId)})
	if err != nil {
		return nil, err
	}

	var list []types.Order
	for _, v := range orderResp.Orders {

		var (
			total float32
			items []types.OrderItem
		)

		for _, v := range v.Items {
			total += float32(v.Cost)
			i := v.Item
			productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: v.Item.ProductId})
			if err != nil {
				return nil, err
			}

			if productResp == nil || productResp.Product == nil {
				continue
			}

			p := productResp.Product

			items = append(items, types.OrderItem{
				ProductName: p.Name,
				Picture:     p.Picture,
				Qty:         uint32(i.Quantity),
				Cost:        float32(v.Cost),
			})
		}

		created := time.Unix(int64(v.CreateAt), 0)
		list = append(list, types.Order{
			OrderId:     v.OrderId,
			CreatedDate: created.Format("2006-01-02 15:04:05"),
			Cost:        total,
			Items:       items,
		})
	}

	return utils.H{
		"Title":  "Order List",
		"orders": list,
	}, nil
}
