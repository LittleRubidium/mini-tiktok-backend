namespace go user

struct BaseResp {
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}

struct CheckUserRequest {
    1: string username (vt.min_size = "1")
    2: string password (vt.min_size = "1")
}

struct CheckUserResponse {
    1: i64 user_id
    2: BaseResp base_resp
}

service UserService {
    CheckUserResponse CheckUser(1: CheckUserRequest req)
}