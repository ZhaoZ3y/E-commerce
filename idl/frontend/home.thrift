namespace go frontend.home

struct Empty {}

service HomeService {
    Empty Home(1: Empty e) (api.get="/home");
}
