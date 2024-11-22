namespace go frontend.order_page

include "frontend/common.thrift"

service OrderPageService {
    common.Empty OrderList(1: common.Empty req) (api.get = "/order")
}