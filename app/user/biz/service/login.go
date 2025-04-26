package service

import (
	"context"
	"fmt"
	"rpc102/app/user/infra/rpc"
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
func (s *LoginService) Run(req *user.LoginReq) (err error) {
	// Finish your business logic.

	fmt.Println("login is success")
	rpc.Client

}
