package service

import (
	"context"
	"fmt"
	"rpc102/app/uid/infra/rpc"
	uid "rpc102/app/uid/kitex_gen/uid"
	"rpc102/app/user/kitex_gen/user"
)

type UidGenService struct {
	ctx context.Context
}

// NewUidGenService new UidGenService
func NewUidGenService(ctx context.Context) *UidGenService {
	return &UidGenService{ctx: ctx}
}

// Run create note info
func (s *UidGenService) Run(req *uid.Empty) (err error) {
	// Finish your business logic.

	r := "uid is success"
	fmt.Println(r)
	rpc.Client.Login(s.ctx, &user.LoginReq{})
	return nil
}
