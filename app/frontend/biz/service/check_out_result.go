package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"

	"github.com/cloudwego/hertz/pkg/app"
	common "gomall/app/frontend/hertz_gen/frontend/common"
)

type CheckOutResultService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckOutResultService(Context context.Context, RequestContext *app.RequestContext) *CheckOutResultService {
	return &CheckOutResultService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckOutResultService) Run(req *common.Empty) (resp map[string]any, err error) {
	// todo edit your code
	return utils.H{}, nil
}
