package service

import (
	"context"
	"github.com/hertz-contrib/sessions"

	"github.com/cloudwego/hertz/pkg/app"
	auth_page "gomall/app/frontend/hertz_gen/frontend/auth_page"
	common "gomall/app/frontend/hertz_gen/frontend/common"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth_page.LoginReq) (resp *common.Empty, err error) {
	// todo user SVC api
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 1)
	session.Save()
	return
}
