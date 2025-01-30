package utils

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/frontend/infra/rpc"
	frontendUtils "github.com/Alanxtl/mymall_go/app/frontend/utils"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpResponse(ctx context.Context, c *app.RequestContext, context map[string]any) (map[string]any, error) {

	if context == nil {
		context = make(map[string]any)
	}

	session := sessions.Default(c)

	fromCtx, err := frontendUtils.GetUserIdFromCtx(ctx)
	if err != nil {
		return context, err
	}
	context["user_id"] = fromCtx
	context["user_name"] = session.Get("user_name")

	if fromCtx > 0 {
		cartResp, err := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{
			UserId: fromCtx,
		})
		if err != nil {
			return nil, err
		}
		context["cart_num"] = len(cartResp.Items)
	}


	return context, nil
}

func QueryFrom(ctx context.Context, c *app.RequestContext, context map[string]any) (map[string]any, error) {
	// 确保 context 被初始化
	if context == nil {
		context = make(map[string]any)
	}

	// 取 Referer 并赋值
	if referer := c.Request.Header.Get("Referer"); referer != "" {
		context["from"] = referer
	}

	if c.Query("from") != "" {
		context["from"] = c.Query("from")
	}

	return context, nil
}
