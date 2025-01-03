package utils

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/frontend/middleware"
	"github.com/hertz-contrib/sessions"

	"github.com/cloudwego/hertz/pkg/app"
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

func WarpResponse(ctx context.Context, c *app.RequestContext, context map[string]any) map[string]any {
	session := sessions.Default(c)

	if context == nil {
		context = make(map[string]any)
	}

	context["user_id"] = ctx.Value(middleware.SessionUserId)
	context["user_name"] = session.Get("user_name")

	return context
}
