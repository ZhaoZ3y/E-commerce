namespace go product

service ProductCatalogService{
    ListProductsResp ListProduct(1:ListProductsReq req),
    GetProductResp GetProduct(1:GetProductReq req),
    SearchProductsResp SearchProduct(1:SearchProductsReq req),
}

struct ListProductsReq{
    1: i32 page,
    2: i32 page_size,
    3: string category_name
}

struct Product{
    1: i64 id,
    2: string name,
    3: string description,
    4: string picture,
    5: double price,

    6: list<string> categories,
}

struct ListProductsResp{
    1: list<Product> products,
}

struct GetProductReq{
    1: i64 id,
}

struct GetProductResp{
    1: Product product,
}

struct SearchProductsReq{
    1: string query,
}

struct SearchProductsResp{
    1: list<Product> results,
}