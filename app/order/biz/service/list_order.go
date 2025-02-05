package service

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/order/biz/dal/mysql"
	"github.com/Alanxtl/mymall_go/app/order/biz/model"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/common"
	order "github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {

	listOrder, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(500001, err.Error())
	}

	var orders []*order.Order
	for _, v := range listOrder {
		var items []*order.OrderItem
		for _, item := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Item: &common.ProductItem{
					ProductId: item.ProductId,
					Quantity:  item.Quantity,
				},
				Cost: item.Cost,
			})
		}
		orders = append(orders, &order.Order{
			Items:        items,
			UserId:       v.UserId,
			OrderId:      v.OrderId,
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			Address: &common.Address{
				Country: v.Consignee.Country,
				State:   v.Consignee.State,
				City:    v.Consignee.State,
				Street:  v.Consignee.StreetAddress,
				Zip:     string(v.Consignee.ZipCode),
			},
			CreatedAt: int32(v.CreatedAt.Unix()),
		})
	}

	resp = &order.ListOrderResp{
		Orders: orders,
	}

	return resp, nil
}
