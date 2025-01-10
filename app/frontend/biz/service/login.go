package service

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/frontend/infra/rpc"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/user"
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
	res, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}

	session := sessions.Default(h.RequestContext)

	session.Set("user_id", res.UserId)
	session.Set("user_name", req.Email)
	err = session.Save()
	if err != nil {
		return req.From, err
	}

	return req.From, nil
}
