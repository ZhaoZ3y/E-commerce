package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"gomall/app/order/biz/dal/mysql"
	"gomall/app/order/biz/model"
	order "gomall/rpc_gen/kitex_gen/order"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	if len(req.Items) == 0 {
		return nil, kerrors.NewBizStatusError(500001, "items is empty")
	}
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		orderId, _ := uuid.NewUUID()
		o := &model.Order{
			OrderId:      orderId.String(),
			UserId:       uint(req.UserId),
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}
		if req.Address != nil {
			a := req.Address
			o.Consignee.StreetAddress = a.StreetAddress
			o.Consignee.City = a.City
			o.Consignee.State = a.State
			o.Consignee.Country = a.Country
		}

		if err := tx.Create(o).Error; err != nil {
			return err
		}

		var items []*model.OrderItem
		for _, v := range req.Items {
			items = append(items, &model.OrderItem{
				OrderIdRefer: orderId.String(),
				ProductId:    uint32(v.Item.ProductId),
				Quantity:     uint32(v.Item.Quantity),
				Cost:         v.Cost,
			})
		}

		if err := tx.Create(items).Error; err != nil {
			return err
		}

		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult_{
				OrderId: orderId.String(),
			},
		}

		return nil
	})
	return
}
