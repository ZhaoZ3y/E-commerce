namespace go frontend.cart_page

include "frontend/common.thrift"

service CartService {
    common.Empty GetCart(1: common.Empty req) (api.get="/cart")
    common.Empty AddCartItem(1: AddCartItemReq req) (api.post="/cart")

}

struct AddCartItemReq {
    1: i64 product_id (api.form = "productId")
    2: i32 product_num (api.form = "productNum")
}