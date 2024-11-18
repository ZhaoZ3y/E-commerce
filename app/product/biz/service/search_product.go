package service

import (
	"context"
	"gomall/app/product/biz/dal/mysql"
	"gomall/app/product/biz/model"
	product "gomall/rpc_gen/kitex_gen/product"
)

type SearchProductService struct {
	ctx context.Context
} // NewSearchProductService new SearchProductService
func NewSearchProductService(ctx context.Context) *SearchProductService {
	return &SearchProductService{ctx: ctx}
}

// Run create note info
func (s *SearchProductService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	products, err := productQuery.SearchProducts(req.Query)
	var results []*product.Product
	for _, v := range products {
		results = append(results, &product.Product{
			Id:          int64(v.ID),
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       float64(v.Price),
		})
	}
	return &product.SearchProductsResp{Results: results}, nil
}
