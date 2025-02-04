package service

import (
	"context"
	checkout "github.com/Alanxtl/mymall_go/app/frontend/hertz_gen/frontend/checkout"
	"github.com/Alanxtl/mymall_go/app/frontend/infra/rpc"
	frontendUtils "github.com/Alanxtl/mymall_go/app/frontend/utils"
	rpcCheckout "github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/checkout"
	rpcPayment "github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	userId, err := frontendUtils.GetUserIdFromCtx(h.Context)
	if err != nil {
		return nil, err
	}

	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpcCheckout.CheckoutReq{
		UserId:    userId,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Address: &rpcCheckout.Address{
			Street:  req.Street,
			City:    req.City,
			Country: req.Country,
			Zip:     req.Zipcode,
			State:   req.Province,
		},
		CreditCard: &rpcPayment.CreditCardInfo{
			CreditCardNum:             req.CardNum,
			CreditCardCvv:             req.Cvv,
			CreditCardExpirationMonth: req.ExpirationMonth,
			CreditCardExpirationYear:  req.ExpirationYear,
		},
	})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title":    "Waiting",
		"redirect": "/checkout/result",
	}, nil
}
