package service

import (
	"context"
	"gomall/app/frontend/hertz_gen/frontend/common"

	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	// todo edit your code
	var resp = make(map[string]any) // 初始化 map
	//定义商品
	item := []map[string]any{
		{"Name": "RedRock Backend T-shirt", "Price": 25, "Picture": "../static/image/backend.jpg"},
		{"Name": "RedRock Frontend T-shirt", "Price": 25, "Picture": "../static/image/frontend.jpg"},
		{"Name": "RedRock Design T-shirt", "Price": 25, "Picture": "../static/image/design.jpg"},
		{"Name": "RedRock Mobile T-shirt", "Price": 25, "Picture": "../static/image/mobile.jpg"},
		{"Name": "RedRock Product T-shirt", "Price": 25, "Picture": "../static/image/product.jpg"},
		{"Name": "RedRock SRE T-shirt", "Price": 25, "Picture": "../static/image/SRE.jpg"},
		{"Name": "RedRock Work Card White", "Price": 10, "Picture": "../static/image/work_card_white.jpg"},
		{"Name": "RedRock Work Card Black", "Price": 10, "Picture": "../static/image/work_card_black.jpg"},
	}
	//返回数据
	resp["Title"] = "Hot Sales"
	resp["Items"] = item

	return resp, nil
}
