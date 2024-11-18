namespace go frontend.category

include "frontend/common.thrift"

struct CategoryReq {
  1: string category (api.path = "category")
}

service CategoryService {
  common.Empty GetCategory(1: CategoryReq req) (api.get = "/category/:category")
}
