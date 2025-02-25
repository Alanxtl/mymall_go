package service

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/cart/biz/dal/mysql"
	"github.com/Alanxtl/mymall_go/app/cart/biz/model"
	cart "github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/cart"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/common"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	list, err := model.GetCartByUserId(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50002, err.Error())
	}

	var items []*common.ProductItem

	for _, item := range list {
		items = append(items, &common.ProductItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}

	return &cart.GetCartResp{Items: items}, nil
}
