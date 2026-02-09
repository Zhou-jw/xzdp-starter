package redis

import (
	"time"

	"github.com/go-redis/redis/v7"

	"github.com/Zhou-jw/xzdp-starter/src/pkg/constants"
)

var (
	expireTime  = time.Minute * 2
	rdb *redis.Client
)



func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       0,
	})
	// rdbFavorite = redis.NewClient(&redis.Options{
	// 	Addr:     constants.RedisAddr,
	// 	Password: constants.RedisPassword,
	// 	DB:       1,
	// })
}