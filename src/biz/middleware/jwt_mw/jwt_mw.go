package jwtmw

import (
	"context"
	"log"
	"time"

	"github.com/Zhou-jw/xzdp-starter/src/biz/model/api/user" // 引入你的用户模型
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
)

// 全局身份Key，用于从请求上下文获取用户信息
var (
	identity   = "phone"
	JwtMiddleware *jwt.HertzJWTMiddleware
)

// 初始化JWT中间件，返回全局的JWT中间件实例
func Init() {
	var err error
	JwtMiddleware ,err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "xzdp-starter",                 // 认证域，自定义即可
		Key:         []byte("xzdp_2026_jwt_secret"), // 加密密钥，生产环境用环境变量！
		Timeout:     time.Hour,                      // Token有效期
		MaxRefresh:  time.Hour,                      // Token最大刷新时间
		IdentityKey: identity,
		
		// 认证函数
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginReq user.SmsLoginRequest
			err := c.BindAndValidate(&loginReq)
			if err != nil {
				log.Printf("Authenticator error in jwt")
				return nil, err
			}
			
			c.Set("user_phone", loginReq.Phone)

			return loginReq.Phone, nil
		},

		// Payload构造：将用户信息存入JWT载荷
		// data is the return value of Authenticator, which is the phone number in this case
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// data是短信登录成功后传入的用户ID/手机号等信息
			if phone, ok := data.(string); ok {
				return jwt.MapClaims{
					identity: phone, // 载荷中存入用户唯一标识（手机号/用户ID）
				}
			}
			return jwt.MapClaims{}
		},

		// token 生成成功后的响应：返回标准化JSON，包含token和过期时间
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			// add user token to redis with expiration
			c.Set("token", token)
		},

		// 解析JWT后，从载荷构造用户身份，存入请求上下文
		// IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
		// 	claims := ExtractClaims(ctx, c)
		// 	return claims[identity]
		// },

		// 权限校验：按需扩展（如管理员/普通用户），暂时返回true，后续可自定义
		// data is the return value of IdentityHandler, default is identity value from PayloadFunc
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			if _, ok := data.(string); ok {
				return true
			}
			return false
		},

		// 未授权统一响应：401/403时返回标准化JSON
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, utils.H{
				"code":    code,
				"message": message,
			})
		},
	})

	if err != nil {
		log.Fatalf("JWT中间件初始化失败: %v", err)
	}
}
