package service

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/frontend/infra/rpc"
	frontendUtils "github.com/Alanxtl/mymall_go/app/frontend/utils"
	rpcCart "github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/cart"
	rpcProduct "github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"

	"github.com/Alanxtl/mymall_go/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *common.Empty) (resp map[string]any, err error) {
	userId, err := frontendUtils.GetUserIdFromCtx(h.Context)
	if err != nil {
		return nil, err
	}

	cartResp, err := rpc.CartClient.GetCart(h.Context, &rpcCart.GetCartReq{UserId: userId})
	if err != nil {
		return nil, err
	}

	var cartItems []map[string]string
	var total float32

	for _, item := range cartResp.Items {

		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcProduct.GetProductReq{Id: item.ProductId})
		if err != nil {
			return nil, err
		}
		if productResp == nil || productResp.Product == nil {
			continue
		}

		p := productResp.Product

		total += p.Price * float32(item.Quantity)

		cartItems = append(cartItems, map[string]string{
			"Name":     p.Name,
			"Price":    strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture":  p.Picture,
			"quantity": strconv.Itoa(int(item.Quantity)),
		})

	}

	return utils.H{
		"title": "Checkout",
		"items": cartItems,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
