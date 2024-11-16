namespace go frontend.auth_page

include "frontend/common.thrift"

struct LoginReq{
    1: string email (api.form = "email")
    2: string password (api.form = "password")
}

struct RegisterReq{
    1: string email (api.form = "email")
    2: string password (api.form = "password")
    3: string password_confirm (api.form = "password_confirm")
}

service AuthService{
    common.Empty Login(LoginReq req) (api.post="/auth/login");
    common.Empty Register(RegisterReq req) (api.post="/auth/register");
}