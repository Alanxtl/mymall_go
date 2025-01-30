package service

import (
	"context"
	"errors"
	"github.com/Alanxtl/mymall_go/app/user/biz/dal/mysql"
	"github.com/Alanxtl/mymall_go/app/user/biz/model"
	user "github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}

	userGet, err := model.GetByEmail(mysql.DB, req.Email)

	err = bcrypt.CompareHashAndPassword([]byte(userGet.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	resp = &user.LoginResp{
		UserId: uint32(userGet.ID),
	}

	return resp, nil
}
