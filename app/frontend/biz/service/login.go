package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	auth_page "gomall/app/frontend/hertz_gen/frontend/auth_page"
	"gomall/app/frontend/infra/rpc"
	"gomall/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth_page.LoginReq) (redirect string, err error) {
	// todo user SVC api
	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp.UserId)
	err = session.Save()
	if err != nil {
		return "", err
	}
	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}

	return
}
