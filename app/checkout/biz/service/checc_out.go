package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"gomall/app/checkout/infra/rpc"
	"gomall/rpc_gen/kitex_gen/cart"
	check_out "gomall/rpc_gen/kitex_gen/check_out"
	"gomall/rpc_gen/kitex_gen/payment"
	"gomall/rpc_gen/kitex_gen/product"
)

type CheccOutService struct {
	ctx context.Context
} // NewCheccOutService new CheccOutService
func NewCheccOutService(ctx context.Context) *CheccOutService {
	return &CheccOutService{ctx: ctx}
}

// Run create note info
func (s *CheccOutService) Run(req *check_out.CheckOutReq) (resp *check_out.CheckOutResp, err error) {
	// Finish your business logic.
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Cart.Items == nil {
		return nil, kerrors.NewBizStatusError(5004001, "cart is empty")
	}

	var total float32
	for _, cartItem := range cartResult.Cart.Items {
		productResp, ResultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: cartItem.ProductId})
		if ResultErr != nil {
			return nil, ResultErr
		}
		if productResp.Product == nil {
			continue
		}

		p := productResp.Product.Price
		cost := float32(p) * float32(cartItem.Quantity)
		total += cost
	}

	var OrderId string
	u, _ := uuid.NewRandom()
	OrderId = u.String()

	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: OrderId,
		Amount:  float64(total),
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		},
	}

	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err.Error())
	}

	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}

	klog.Info(paymentResult)

	return &check_out.CheckOutResp{
		OrderId:       OrderId,
		TransactionId: paymentResult.TransactionId,
	}, nil
}
