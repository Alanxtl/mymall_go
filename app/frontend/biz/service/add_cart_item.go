package service

import (
	"context"
	cart "github.com/Alanxtl/mymall_go/app/frontend/hertz_gen/frontend/cart"
	rpcfrontendcommon "github.com/Alanxtl/mymall_go/app/frontend/hertz_gen/frontend/common"
	"github.com/Alanxtl/mymall_go/app/frontend/infra/rpc"
	frontendUtils "github.com/Alanxtl/mymall_go/app/frontend/utils"
	rpccart "github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/cart"
	rpccommon "github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp *rpcfrontendcommon.Empty, err error) {
	userId, err := frontendUtils.GetUserIdFromCtx(h.Context)

	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: userId,
		Item: &rpccommon.ProductItem{
			ProductId: req.ProductId,
			Quantity:  req.ProductNum,
		},
	})
	if err != nil {
		return nil, err
	}
	return
}
