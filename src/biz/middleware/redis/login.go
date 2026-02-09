package redis

import (
	"strconv"

	"github.com/go-redis/redis/v7"
)


const (
	loginSuffix = ":login_code"
)

func AddSmsCode(phone, code string) error {
	sadd(rdb, phone+loginSuffix, code)
	return nil
}

func CheckSmsCode(phone, code string) (bool, error) {
	return sexist(rdb, phone, code), nil
}

func loginkey(phone string) string {
	return phone + loginSuffix
}

// add k & v to redis
func sadd(c *redis.Client, k string, v string) {
	tx := c.TxPipeline()
	tx.SAdd(k, v)
	tx.Expire(k, expireTime)
	tx.Exec()
}

// check the set of k if exist
func scheck(c *redis.Client, k string) bool {
	k = loginkey(k)
	if e, _ := c.Exists(k).Result(); e > 0 {
		return true
	}
	return false
}

// exist check the relation k and v if exist
func sexist(c *redis.Client, k string, v string) bool {
	k = loginkey(k)
	if e, _ := c.SIsMember(k, v).Result(); e {
		c.Expire(k, expireTime)
		return true
	}
	return false
}

func sget(c *redis.Client, k string) (vt []int64) {
	k = loginkey(k)
	v, _ := c.SMembers(k).Result()
	c.Expire(k, expireTime)
	for _, vs := range v {
		v_i64, _ := strconv.ParseInt(vs, 10, 64)
		vt = append(vt, v_i64)
	}
	return vt
}