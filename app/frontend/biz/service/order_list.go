package service

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/frontend/infra/rpc"
	"github.com/Alanxtl/mymall_go/app/frontend/types"
	frontendUtils "github.com/Alanxtl/mymall_go/app/frontend/utils"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/order"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"sort"
	"time"

	common "github.com/Alanxtl/mymall_go/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {

	userId, err := frontendUtils.GetUserIdFromCtx(h.Context)
	if err != nil {
		return nil, err
	}

	listOrderResp, err := rpc.OrderClient.ListOrder(h.Context, &order.ListOrderReq{UserId: userId})
	if err != nil {
		return nil, err
	}

	var orders []*types.Order
	for _, v := range listOrderResp.Orders {
		var (
			orderItems []types.OrderItem
			totalCost  float32
		)

		for _, v := range v.Items {
			item, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: v.Item.ProductId})
			if err != nil {
				return nil, err
			}
			if item == nil || item.Product == nil {
				continue
			}

			p := item.Product
			orderItems = append(orderItems, types.OrderItem{
				ProductName: p.Name,
				Picture:     p.Picture,
				Quantity:    v.Item.Quantity,
				Cost:        p.Price,
			})
			totalCost += v.Cost * float32(v.Item.Quantity)
		}

		orders = append(orders, &types.Order{
			OrderId:     v.OrderId,
			Items:       orderItems,
			Cost:        totalCost,
			CreatedDate: time.Unix(int64(v.CreatedAt), 0).Format("2006-01-02 15:04:05"),
		})
	}


	sort.Slice(orders, func(i, j int) bool {
		return i > j
	})

	return utils.H{
		"title":  "Order",
		"orders": orders,
	}, nil
}
