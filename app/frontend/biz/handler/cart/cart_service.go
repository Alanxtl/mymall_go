package cart

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/frontend/biz/service"
	"github.com/Alanxtl/mymall_go/app/frontend/biz/utils"
	cart "github.com/Alanxtl/mymall_go/app/frontend/hertz_gen/frontend/cart"
	common "github.com/Alanxtl/mymall_go/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetCart .
// @router /cart [POST]
func GetCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetCartService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err = utils.WarpResponse(ctx, c, resp)
	if err != nil {
		return
	}

	c.HTML(consts.StatusOK, "cart", resp)

	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// AddCartItem .
// @router /cart [POST]
func AddCartItem(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.AddCartItemReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewAddCartItemService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.Redirect(consts.StatusFound, []byte("/cart"))
	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
