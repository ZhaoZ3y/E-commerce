package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"gomall/app/order/biz/dal/mysql"
	"gomall/app/order/biz/model"
	"gomall/rpc_gen/kitex_gen/cart"
	order "gomall/rpc_gen/kitex_gen/order"
)

type ListOrdersService struct {
	ctx context.Context
} // NewListOrdersService new ListOrdersService
func NewListOrdersService(ctx context.Context) *ListOrdersService {
	return &ListOrdersService{ctx: ctx}
}

// Run create note info
func (s *ListOrdersService) Run(req *order.ListOrdersReq) (resp *order.ListOrdersResp, err error) {
	// Finish your business logic.
	list, err := model.ListOrder(s.ctx, mysql.DB, uint32(req.UserId))
	if err != nil {
		return nil, kerrors.NewBizStatusError(500001, err.Error())
	}

	var orders []*order.Order
	for _, v := range list {
		var items []*order.OrderItem
		for _, oi := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: int64(oi.ProductId),
					Quantity:  int64(oi.Quantity),
				},
				Cost: oi.Cost,
			})
		}

		orders = append(orders, &order.Order{
			OrderId:      v.OrderId,
			UserId:       int64(v.UserId),
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			Address: &order.Adress{
				StreetAddress: v.Consignee.StreetAddress,
				City:          v.Consignee.City,
				State:         v.Consignee.State,
				Country:       v.Consignee.Country,
				ZipCope:       v.Consignee.ZipCode,
			},
		})
	}
	resp = &order.ListOrdersResp{
		Orders: orders,
	}
	return
}
