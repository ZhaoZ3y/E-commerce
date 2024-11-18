package service

import (
	"context"
	product "gomall/app/product/kitex_gen/product"
	"testing"
)

func TestListProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListProductService(ctx)
	// init req and assert value

	req := &product.ListProductsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
