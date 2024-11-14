package service

import (
	"context"
	pbapi "demo_proto/kitex_gen/pbapi"
	"fmt"
	"github.com/bytedance/gopkg/cloud/metainfo"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(req *pbapi.Request) (resp *pbapi.Response, err error) {
	// Finish your business logic.
	clientName, ok := metainfo.GetPersistentValue(s.ctx, "CLIENT_NAME")
	fmt.Println("clientName:", clientName, ok)
	return &pbapi.Response{Message: req.Message}, nil
}
