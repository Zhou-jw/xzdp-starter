// idl/user.thrift

namespace go api.user

// ------------------------------
// 2. 接口请求/响应结构体
// ------------------------------
// 2.1 发送验证码：Request + Response
struct SendCodeRequest {
    1: required string Phone (api.query="phone");  // 手机号（Query参数，?phone=xxx）
}

struct SendCodeResponse {
    1: required i32 Code = 200;             // 状态码：200成功，非200失败
    2: required string Msg = "ok";          // 提示信息：失败时返回具体错误
    3: required string Data;                      // 扩展字段：预留（如验证码ID）
    4: optional i64 Timestamp;                    // 响应时间戳（毫秒）
}

// 2.2 短信登录：Request + Response
struct SmsLoginRequest {
    1: required string Phone (api.body="phone");   // 手机号（JSON请求体）
    2: required string SmsCode (api.body="code");  // 短信验证码（JSON请求体）
}

struct SmsLoginResponse {
    1: required i32 Code = 200;             // 状态码
    2: required string Msg = "ok";          // 提示信息
    3: optional string AccessToken;               // JWT访问令牌（核心）
    4: optional i64 ExpireAt;                     // Token过期时间（毫秒时间戳）
    5: optional string Phone;                     // 登录手机号（替代User结构体的核心字段）
    6: optional i64 Timestamp;                    // 响应时间戳
}

// 2.3 获取当前用户信息
struct UserMeRequest {
    1: optional string Token (api.header="Authorization"); // JWT令牌（请求头）
}

struct UserMeResponse {
    1: required i32 Code = 200;             // 状态码
    2: required string Msg = "ok";          // 提示信息
    3: optional string Phone;                     // 当前登录用户手机号（核心标识）
    // 可选扩展：如需其他基础信息，直接添加字符串字段，无需结构体
    // 4: optional string NickName;
    // 5: optional string Avatar;
    4: optional i64 Timestamp;                    // 响应时间戳
}

// ------------------------------
// 4. 用户核心服务定义
// ------------------------------
service UserService {
    /**
     * 发送短信验证码接口
     */
    SendCodeResponse SendCode(1: SendCodeRequest req) (api.post="/api/user/code", api.path="/api/user/code");

    /**
     * 短信验证码登录接口
     */
    SmsLoginResponse SmsLogin(1: SmsLoginRequest req) (api.post="/api/user/login", api.path="/api/user/login");

    /**
     * 获取当前用户信息接口
     */
    UserMeResponse UserMe(1: UserMeRequest req) (api.get="/api/user/me", api.path="/api/user/me");
}