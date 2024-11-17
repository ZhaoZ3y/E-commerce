namespace go user

struct RegisterReq{
    1: string email,
    2: string password,
    3: string password_confirm
}

struct RegisterResp{
    1: i32 user_id
}

struct LoginReq{
    1: string email,
    2: string password
}

struct LoginResp{
    1: i32 user_id
}

service UserService{
    RegisterResp Register(RegisterReq req),
    LoginResp Login(LoginReq req)
}