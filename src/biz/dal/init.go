package dal

import (
	"github.com/Zhou-jw/xzdp-starter/src/biz/dal/db"
	"github.com/Zhou-jw/xzdp-starter/src/biz/middleware/redis"
)

// Init init dal
func Init() {
	db.Init() // mysql init
	redis.InitRedis()
}