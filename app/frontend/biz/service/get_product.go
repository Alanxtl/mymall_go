package service

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/frontend/infra/rpc"
	"github.com/cloudwego/hertz/pkg/common/utils"

	product "github.com/Alanxtl/mymall_go/app/frontend/hertz_gen/frontend/product"
	rpcProduct "github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	p, err := rpc.ProductClient.GetProduct(h.Context, &rpcProduct.GetProductReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"item": p.Product,
	}, nil
}
