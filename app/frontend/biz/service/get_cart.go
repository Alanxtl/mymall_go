package service

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/frontend/infra/rpc"
	frontendUtils "github.com/Alanxtl/mymall_go/app/frontend/utils"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/cart"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"

	common "github.com/Alanxtl/mymall_go/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	userId, err := frontendUtils.GetUserIdFromCtx(h.Context)
	cartResp, err := rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{
		UserId: userId,
	})

	if err != nil {
		return nil, err
	}

	var items []map[string]string
	var total float64
	for _, item := range cartResp.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{
			Id: item.ProductId,
		})

		if err != nil {
			continue
		}

		p := productResp.Product

		items = append(items, map[string]string{
			"Name":        p.Name,
			"Description": p.Description,
			"Picture":     p.Picture,
			"Price":       strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Quantity":    strconv.FormatUint(uint64(item.Quantity), 10),
			"Id":          strconv.Itoa(int(p.Id)),
		})

		total += float64(item.Quantity) * float64(p.Price)
	}

	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(total, 'f', 2, 64),
	}, nil
}
