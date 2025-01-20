package service

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/frontend/infra/rpc"
	rpcProduct "github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"

	product "github.com/Alanxtl/mymall_go/app/frontend/hertz_gen/frontend/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	p, err := rpc.ProductClient.SearchProducts(h.Context, &rpcProduct.SearchProductsReq{
		Query: req.Q,
	})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"items": p.Results,
		"q":     req.Q,
	}, nil
}
