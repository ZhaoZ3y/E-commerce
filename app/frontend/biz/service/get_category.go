package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"gomall/app/frontend/infra/rpc"
	"gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	category "gomall/app/frontend/hertz_gen/frontend/category"
)

type GetCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCategoryService(Context context.Context, RequestContext *app.RequestContext) *GetCategoryService {
	return &GetCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	// todo use product svc api
	p, err := rpc.ProductClient.ListProduct(h.Context, &product.ListProductsReq{CategoryName: req.Category})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title": req.Category,
		"items": p.Products,
	}, nil
}
