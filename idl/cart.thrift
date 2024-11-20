namespace go cart

service CartService{
    AddItemResp AddItem(1:AddItemReq req),
    GetCartResp GetCart(1:GetCartReq req),
    EmptyCartResp EmptyCart(1:EmptyCartReq req),
}

struct CartItem{
    1: i64 product_id,
    2: i64 quantity,
}

struct AddItemReq{
    1: i64 user_id,
    2: CartItem item,
}

struct AddItemResp{}

struct GetCartReq{
    1: i64 user_id,
}

struct GetCartResp{
    1: Cart cart,
}

struct EmptyCartReq{
    1: i64 user_id,
}

struct Cart {
  1: i64 user_id = 1;
  2: list<CartItem> items = 2;
}

struct EmptyCartResp{}