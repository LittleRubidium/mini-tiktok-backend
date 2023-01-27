namespace go api

struct CheckUserRequest {
    1: string username (api.query="username", api.vd="len($) > 0")
    2: string password (api.query="password", api.vd="len($) > 0")
}

struct CheckUserResponse {
    1: i64 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

service ApiService {
    CheckUserResponse CheckUser(1: CheckUserRequest req) (api.post="/douyin/user/login/")
}