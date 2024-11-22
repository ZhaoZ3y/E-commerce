namespace go order

include "cart.thrift"

service OrderService {
    PlaceOrderResp PlaceOrder(1: PlaceOrderReq req)
    ListOrdersResp ListOrders(1: ListOrdersReq req)
}

struct Adress {
    1: string street_address
    2: string city
    3: string state
    4: string country
    5: i32 zip_cope
}

struct OrderItem{
    1: cart.CartItem item
    2: double cost
}

struct PlaceOrderReq {
    1: i64 user_id
    2: string user_currency
    3: Adress address
    4: string email
    5: list<OrderItem> items
}

struct OrderResult{
    1: string order_id
}

struct PlaceOrderResp {
    1: OrderResult order
}

struct ListOrdersReq {
    1: i64 user_id
}

struct Order{
    1: list<OrderItem> items
    2: string order_id
    3: i64 user_id
    4: string user_currency
    5: Adress address
    6: string email
    7: i32 create_at
}

struct ListOrdersResp {
    1: list<Order> orders
}