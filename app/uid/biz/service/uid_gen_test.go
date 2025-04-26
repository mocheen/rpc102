package service

import (
	"context"
	uid "rpc102/app/uid/kitex_gen/uid"
	"testing"
)

func TestUidGen_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUidGenService(ctx)
	// init req and assert value

	req := &uid.Empty{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
