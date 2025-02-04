package service

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/order/biz/dal/mysql"
	"github.com/Alanxtl/mymall_go/app/order/biz/model"
	order "github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
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
	if len(req.Items) == 0 {
		return nil, kerrors.NewBizStatusError(500001, "item is empty")
	}

	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		orderId, _ := uuid.NewUUID()

		o := &model.Order{
			OrderId:      orderId.String(),
			UserId:       req.UserId,
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}

		if req.Address != nil {
			a := req.Address
			o.Consignee.StreetAddress = a.Street
			o.Consignee.City = a.City
			o.Consignee.State = a.State
			o.Consignee.Country = a.Country
		}

		if err := tx.Create(o).Error; err != nil {
			return err
		}

		var items []model.OrderItem
		for _, item := range req.Items {
			items = append(items, model.OrderItem{
				ProductId:    item.Item.ProductId,
				OrderIdRefer: orderId.String(),
				Quantity:     item.Item.Quantity,
				Cost:         item.Cost,
			})
		}

		if err := tx.Create(items).Error; err != nil {
			return err
		}

		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: orderId.String(),
			},
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return
}
