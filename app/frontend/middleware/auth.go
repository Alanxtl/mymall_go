package middleware

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)

		ctx = context.WithValue(ctx, utils.SessionUserId, session.Get("user_id"))
		//ctx = context.WithValue(ctx, utils.SessionUserNAME, session.Get("user_name"))
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)

		userId := session.Get("user_id")

		if userId == nil {
			c.Redirect(302, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}

		c.Next(ctx)
	}
}

func AuthLogin() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		//session := sessions.Default(c)
		//
		//userId := session.Get("user_id")
		//
		//if userId != nil {
		//	c.Redirect(302, []byte("/?from="+c.FullPath()))
		//	c.Abort()
		//	return
		//}

		c.Next(ctx)
	}
}
