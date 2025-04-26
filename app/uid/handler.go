package main

import (
	"context"
	"rpc102/app/uid/biz/service"
	uid "rpc102/app/uid/kitex_gen/uid"
)

// UidServiceImpl implements the last service interface defined in the IDL.
type UidServiceImpl struct{}

// UidGen implements the UidServiceImpl interface.
func (s *UidServiceImpl) UidGen(ctx context.Context, req *uid.Empty) (resp *uid.Uid, err error) {
	resp, err = service.NewUidGenService(ctx).Run(req)

	return resp, err
}
