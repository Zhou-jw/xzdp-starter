package redis

import (
	"strconv"
	"time"

	"github.com/go-redis/redis/v7"
)


const (
	loginCodeSuffix = ":login_code"
	loginMeSuffix = ":login_me"
	loginTTL = time.Minute * 2
	tokenTTL = time.Hour * 1
)

func logincodekey(phone string) string {
	return phone + loginCodeSuffix
}

func AddSmsCode(phone, code string) error {
	sadd(rdb, phone+loginCodeSuffix, code)
	return nil
}

func CheckSmsCode(phone, code string) (bool, error) {
	return sexist(rdb, phone, code), nil
}

func loginmekey(phone string) string {
	return phone + loginMeSuffix
}

func AddUser(phone string, token string) error {
	hadd(rdb, token, phone)
	return nil
}

func GetUserPhoneByToken(token string) (string, error) {
	return hget(rdb, token), nil
}

func hadd(c *redis.Client, k string, v string) {
	k = loginmekey(k)
	tx := c.TxPipeline()
	tx.HSet(k, "phone", v)
	tx.Expire(k, tokenTTL)
	tx.Exec()
}

func hget(c *redis.Client, k string) string {
	k = loginmekey(k)
	v, _ := c.HGet(k, "phone").Result()
	c.Expire(k, tokenTTL)
	return v
}

// add k & v to redis
func sadd(c *redis.Client, k string, v string) {
	tx := c.TxPipeline()
	tx.SAdd(k, v)
	tx.Expire(k, loginTTL)
	tx.Exec()
}

// check the set of k if exist
func scheck(c *redis.Client, k string) bool {
	k = logincodekey(k)
	if e, _ := c.Exists(k).Result(); e > 0 {
		return true
	}
	return false
}

// exist check the relation k and v if exist
func sexist(c *redis.Client, k string, v string) bool {
	k = logincodekey(k)
	if e, _ := c.SIsMember(k, v).Result(); e {
		c.Expire(k, loginTTL)
		return true
	}
	return false
}

func sget(c *redis.Client, k string) (vt []int64) {
	k = logincodekey(k)
	v, _ := c.SMembers(k).Result()
	c.Expire(k, loginTTL)
	for _, vs := range v {
		v_i64, _ := strconv.ParseInt(vs, 10, 64)
		vt = append(vt, v_i64)
	}
	return vt
}