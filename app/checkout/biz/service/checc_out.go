package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	thrifter "github.com/thrift-iterator/go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"gomall/app/checkout/infra/mq"
	"gomall/app/checkout/infra/rpc"
	"gomall/rpc_gen/kitex_gen/cart"
	check_out "gomall/rpc_gen/kitex_gen/check_out"
	"gomall/rpc_gen/kitex_gen/email"
	"gomall/rpc_gen/kitex_gen/order"
	"gomall/rpc_gen/kitex_gen/payment"
	"gomall/rpc_gen/kitex_gen/product"
	"strconv"
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

	var (
		total float32
		oi    []*order.OrderItem
	)
	for _, cartItem := range cartResult.Cart.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: cartItem.ProductId})
		if resultErr != nil {
			klog.Error(resultErr)
			err = resultErr
			return
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		cost := float32(p.Price) * float32(cartItem.Quantity)
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{ProductId: cartItem.ProductId, Quantity: cartItem.Quantity},
			Cost: float64(cost),
		})
	}

	orderReq := &order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "USD",
		Items:        oi,
		Email:        req.Email,
	}
	if req.Address != nil {
		addr := req.Address
		zipCodeInt, _ := strconv.Atoi(addr.ZipCode)
		orderReq.Address = &order.Adress{
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			Country:       addr.Country,
			State:         addr.State,
			ZipCope:       int32(zipCodeInt),
		}
	}

	orderResult, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	if err != nil {
		err = fmt.Errorf("PlaceOrder.err:%v", err)
		return
	}
	klog.Info("orderResult", orderResult)
	// empty cart
	emptyResult, err := rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		err = fmt.Errorf("EmptyCart.err:%v", err)
		return
	}
	klog.Info(emptyResult)

	var orderId string
	if orderResult != nil && orderResult.Order != nil {
		orderId = orderResult.Order.OrderId
	}

	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  float64(total),
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
		},
	}
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		err = fmt.Errorf("Charge.err:%v", err)
		return
	}

	data, _ := thrifter.Marshal(&email.EmailReq{
		From:        "from@example.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "Order Confirmation",
		Content:     "Dear " + req.Email + ",\n\nThank you for your order. Your order number is " + orderId + ".\n\n",
	})

	msg := &nats.Msg{
		Subject: "email",
		Data:    data,
		Header:  make(nats.Header),
	}
	otel.GetTextMapPropagator().Inject(s.ctx, propagation.HeaderCarrier(msg.Header))

	_ = mq.Nc.PublishMsg(msg)
	klog.Info(paymentResult)

	return &check_out.CheckOutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}, nil
}
