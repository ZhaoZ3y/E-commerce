package service

import (
	"context"
	"gomall/app/product/biz/dal/mysql"
	"gomall/app/product/biz/model"
	product "gomall/rpc_gen/kitex_gen/product"
)

type ListProductService struct {
	ctx context.Context
} // NewListProductService new ListProductService
func NewListProductService(ctx context.Context) *ListProductService {
	return &ListProductService{ctx: ctx}
}

// Run create note info
func (s *ListProductService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	resp = &product.ListProductsResp{}
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
	c, err := categoryQuery.GetProductsByCategoryName(req.CategoryName)
	if err != nil {
		return nil, err
	}

	for _, v1 := range c {
		for _, v2 := range v1.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          int64(v2.ID),
				Name:        v2.Name,
				Description: v2.Description,
				Picture:     v2.Picture,
				Price:       float64(v2.Price),
			})
		}
	}
	return resp, nil
}
