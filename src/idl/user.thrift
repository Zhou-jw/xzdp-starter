// idl/user.thrift

// 1. 命名空间：生成Go代码时的包名，对应项目模块路径
namespace go api.user

// 2. 空结构体：用于无入参接口（如刷新token、获取当前用户）
struct Empty {}

// 3. 发送验证码请求结构体：匹配POST /api/user/code?phone=xxx的请求格式
// 核心：phone通过Query传参，与HTTP请求URL中的query参数对应
struct SendCodeRequest {
    1: required string Phone (api.query="phone");  // 手机号（必传，Query参数，对应?phone=xxx）
}

struct SmsLoginRequest {
    1: required string Phone (api.query="phone");   // 手机号（Query参数，与发验证码一致）
    2: required string SmsCode (api.query="code");  // 短信验证码（Query参数，如?code=123456）
}

// 4. 通用返回结果结构体：统一接口返回格式，适配所有业务接口
// 替代零散的返回类型，提升项目一致性
struct BaseResponse {
    1: required i32 Code = 200;        // 状态码：200成功，非200失败
    2: required string Message = "ok"; // 提示信息：失败时返回具体错误描述
    3: optional string Data;           // 数据体：成功时返回业务数据（JSON字符串，泛型适配）
    4: optional i64 Timestamp;         // 时间戳：服务端响应时间（毫秒）
}

// 5. 登录成功返回结构体：包含JWT Token和基础用户信息（验证码登录成功后返回）
// 核心：给客户端返回可直接使用的JWT令牌，客户端后续请求携带该token鉴权
struct LoginResponse {
    1: required string AccessToken;    // Hertz-JWT生成的访问令牌（核心，客户端需存在header中）
    2: required i64 ExpireAt;          // token过期时间（毫秒时间戳，客户端可做本地过期判断）
    3: optional string Phone;          // 登录手机号，方便客户端展示
    // 可选：追加轻量用户信息（如昵称、头像），避免客户端二次请求
    // 4: optional string NickName;
    // 5: optional string Avatar;
}

// 6. 用户核心服务定义：聚合用户模块所有接口，IDL即接口契约
service UserService {
    /**
     * 发送短信验证码接口：匹配需求中的POST /api/user/code?phone=xxx
     * 注解说明：api.post指定POST方法，api.path指定完整请求路径
     * 返回值：BaseResponse，成功时Data为""，失败时Message返回错误信息
     */
    BaseResponse SendCode(1: SendCodeRequest req) (api.post="/api/user/code", api.path="/api/user/code");

    /**
     * 短信验证码登录接口：基于发送的验证码完成登录，返回JWT Token
     * 补充：与SendCode配合，完成「发验证码→验证验证码登录」闭环
     */
    BaseResponse SmsLogin(1: SmsLoginRequest req) (api.post="/api/user/login", api.path="/api/user/login");
}

// ：短信验证码登录请求结构体（与SendCode配套，可选追加）
