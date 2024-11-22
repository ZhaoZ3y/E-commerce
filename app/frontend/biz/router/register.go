// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	frontend_auth_page "gomall/app/frontend/biz/router/frontend/auth_page"
	frontend_cart_page "gomall/app/frontend/biz/router/frontend/cart_page"
	frontend_category "gomall/app/frontend/biz/router/frontend/category"
	frontend_checkout_page "gomall/app/frontend/biz/router/frontend/checkout_page"
	frontend_common "gomall/app/frontend/biz/router/frontend/common"
	frontend_home "gomall/app/frontend/biz/router/frontend/home"
	frontend_order_page "gomall/app/frontend/biz/router/frontend/order_page"
	frontend_product_page "gomall/app/frontend/biz/router/frontend/product_page"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	frontend_order_page.Register(r)

	frontend_checkout_page.Register(r)

	frontend_cart_page.Register(r)

	frontend_category.Register(r)

	frontend_product_page.Register(r)

	frontend_common.Register(r)

	frontend_auth_page.Register(r)

	frontend_home.Register(r)
}
