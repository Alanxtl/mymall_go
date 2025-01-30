package product

import (
	"context"

	"github.com/Alanxtl/mymall_go/app/frontend/biz/service"
	"github.com/Alanxtl/mymall_go/app/frontend/biz/utils"
	common "github.com/Alanxtl/mymall_go/app/frontend/hertz_gen/frontend/common"
	product "github.com/Alanxtl/mymall_go/app/frontend/hertz_gen/frontend/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetProduct .
// @router /product [GET]
func GetProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetProductService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err = utils.WarpResponse(ctx, c, resp)
	if err != nil {
		return
	}

	c.HTML(consts.StatusOK, "product", resp)
	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// SearchProducts .
// @router /search [GET]
func SearchProducts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.SearchProductsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewSearchProductsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "search", resp)
	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
