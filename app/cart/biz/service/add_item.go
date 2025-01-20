package service

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/cart/biz/dal/mysql"
	"github.com/Alanxtl/mymall_go/app/cart/biz/model"
	"github.com/Alanxtl/mymall_go/app/cart/rpc"
	cart "github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/cart"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}

	if productResp == nil || productResp.Products.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product not found")
	}

	cartItem := &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Quantity:  req.Item.Quantity,
	}

	err = model.AddItem(s.ctx, mysql.DB, cartItem)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}

	return &cart.AddItemResp{}, nil
}
