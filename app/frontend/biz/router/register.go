// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	frontend_auth_page "gomall/app/frontend/biz/router/frontend/auth_page"
	frontend_common "gomall/app/frontend/biz/router/frontend/common"
	frontend_home "gomall/app/frontend/biz/router/frontend/home"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	frontend_common.Register(r)

	frontend_auth_page.Register(r)

	frontend_home.Register(r)
}
