namespace go frontend.product_page

include "frontend/common.thrift"

service ProductService{
    common.Empty GetProduct(ProductReq req) (api.get="/product");
    common.Empty SearchProduct(SearchReq req) (api.get="/search");
}

struct ProductReq{
    1: i64 id (api.query = "id")
}

struct SearchReq{
    1:string q (api.query = "q")
}