namespace go check_out

include "payment.thrift"

service CheckOutService {
    CheckOutResp CheccOut (1: CheckOutReq req)
}

struct CheckOutReq {
    1: i64 user_id,
    2: string firstname,
    3: string lastname,
    4: string email,
    5: Address address,
    6: payment.CreditCardInfo credit_card,
}

struct Address{
    1: string street_address,
    2: string city,
    3: string state,
    4: string country,
    5: string zip_code,
}

struct CheckOutResp {
    1: string order_id,
    2: string transaction_id,
}