// Code generated by hertz generator. DO NOT EDIT.

package product_page

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	product_page "gomall/app/frontend/biz/handler/frontend/product_page"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.GET("/product", append(_getproductMw(), product_page.GetProduct)...)
	root.GET("/search", append(_searchproductMw(), product_page.SearchProduct)...)
}
