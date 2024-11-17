package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	frontendUtils "gomall/app/frontend/utils"
)

func GlobalAuth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// todo global auth
		s := sessions.Default(ctx)
		c = context.WithValue(c, frontendUtils.SessionUserId, s.Get("user_id"))
		ctx.Next(c)
	}
}

func Auth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// todo l auth
		s := sessions.Default(ctx)
		userId := s.Get("user_id")
		if userId == nil {
			ctx.Redirect(302, []byte("/sign-in?next="+ctx.FullPath()))
			ctx.Abort()
			return
		}
		ctx.Next(c)
	}
}
