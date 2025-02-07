package service

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/checkout/infra/rpc"
	"github.com/Alanxtl/mymall_go/app/email/infra/mq"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/cart"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/checkout"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/common"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/email"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/order"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/payment"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	cartResp, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}

	if cartResp == nil || cartResp.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}

	var (
		total      float32
		orderItems []*order.OrderItem
		orderId    string
	)

	for _, item := range cartResp.Items {
		getProductResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: item.ProductId})
		if err != nil {
			return nil, err
		}
		if getProductResp == nil || getProductResp.Product == nil {
			continue
		}

		p := getProductResp.Product
		cost := p.Price * float32(item.Quantity)
		total += cost

		orderItems = append(orderItems, &order.OrderItem{
			Item: &common.ProductItem{
				ProductId: p.Id,
				Quantity:  item.Quantity,
			},
			Cost: cost,
		})
	}

	placeOrderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "USD",
		Address:      req.Address,
		Email:        req.Email,
		Items:        orderItems,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5004002, err.Error())
	}
	if placeOrderResp != nil || placeOrderResp.Order != nil {
		orderId = placeOrderResp.Order.OrderId
	}

	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNum:             req.CreditCard.CreditCardNum,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		},
	}

	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err.Error())
	}

	paymentResp, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}

	data, _ := proto.Marshal(&email.EmailReq{
		From:        "from@example.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "You have just created order in mymall_go",
		Content:     "You have just created order in mymall_go",
	})

	msg := &nats.Msg{Subject: "email", Data: data}

	_ = mq.Nc.PublishMsg(msg)

	klog.Info(paymentResp)

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResp.TransactionId,
	}

	return resp, nil
}
