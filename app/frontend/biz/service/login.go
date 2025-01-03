package service

import (
	"context"
	"github.com/hertz-contrib/sessions"

	auth "github.com/Alanxtl/mymall_go/app/frontend/hertz_gen/frontend/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	session := sessions.Default(h.RequestContext)

	session.Set("user_id", 1)
	session.Set("user_name", req.Email)
	err = session.Save()
	if err != nil {
		return req.From, err
	}

	return req.From, nil
}
