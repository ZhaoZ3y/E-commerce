namespace go frontend.home

include "common.thrift"

service HomeService {
    common.Empty Home(1: common.Empty e) (api.get="/");
}
