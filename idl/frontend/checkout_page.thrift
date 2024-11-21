namespace go frontend.checkout_page

include "frontend/common.thrift"

struct CheckoutReq {
  1: string email (api.form = "email")
  2: string firstname (api.form = "firstname")
  3: string lastname (api.form = "lastname")
  4: string street (api.form = "street")
  5: string zipcode (api.form = "zipcode")
  6: string province (api.form = "province")
  7: string country (api.form = "country")
  8: string city (api.form = "city")
  9: string cart_num (api.form = "cartNum")
  10: i32 expiration_month (api.form = "expirationMonth")
  11: i32 expiration_year (api.form = "expirationYear")
  12: i32 cvv (api.form = "cvv")
  13: string payment (api.form = "payment")
}

service CheckOutService {
    common.Empty CheckOut (common.Empty req) (api.get="/checkout")
    common.Empty CheckOutWaiting (CheckoutReq req) (api.post="/checkout/waiting")
    common.Empty CheckOutResult (common.Empty req) (api.get="/checkout/result")
}
