package service

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/frontend/infra/rpc"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/user"
	"github.com/hertz-contrib/sessions"

	auth "github.com/Alanxtl/mymall_go/app/frontend/hertz_gen/frontend/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp string, err error) {
	res, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
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
