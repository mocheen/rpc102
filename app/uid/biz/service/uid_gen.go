package service

import (
	"context"
	uid "rpc102/app/uid/kitex_gen/uid"
)

type UidGenService struct {
	ctx context.Context
}

// NewUidGenService new UidGenService
func NewUidGenService(ctx context.Context) *UidGenService {
	return &UidGenService{ctx: ctx}
}

// Run create note info
func (s *UidGenService) Run(req *uid.Empty) (resp *uid.Uid, err error) {
	// Finish your business logic.

	r := "ssss"
	return &uid.Uid{Uid: r}, nil
}
