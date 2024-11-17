package service

import (
	"context"
	"github.com/hertz-contrib/sessions"
	"gomall/app/frontend/infra/rpc"
	"gomall/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	auth_page "gomall/app/frontend/hertz_gen/frontend/auth_page"
	common "gomall/app/frontend/hertz_gen/frontend/common"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth_page.RegisterReq) (resp *common.Empty, err error) {
	// todo user svc api
	UserResp, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	})
	if err != nil {
		return nil, err
	}

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", UserResp.UserId)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return
}
