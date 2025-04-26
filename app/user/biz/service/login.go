package service

import (
	"context"
	user "rpc102/app/user/kitex_gen/user"
)

type LoginService struct {
	ctx context.Context
}

// NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.

	r := &user.LoginResp{
		Token: "11111",
		Roles: []string{"ddd"},
	}
	return r, nil
}
