package utils

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

// GetCookie2FromRedis : Get Single bm_sz Cookie From Redis
func GetCookie2FromRedis() string {
	con := pool.Get()
	defer con.Close()

	cookieList, err := redis.Strings(con.Do("ZRANGE", "cookie2", "0", "1"))
	if len(cookieList) > 0 {
		res, _ := redis.Int64(con.Do("ZREM", "cookie2", cookieList[0]))
		if res > 0 {
			return cookieList[0]
		}
		return ""
	}
	if err != nil {
		fmt.Println(err)
	}
	return ""
}

// SaveCookie2ToRedis : Save _abck To Redis
func SaveCookie2ToRedis(expiry float64, cookie string) error {
	con := pool.Get()
	defer con.Close()

	expiryStr := fmt.Sprintf("%f", expiry)
	_, err := redis.Int64(con.Do("ZADD", "cookie2", expiryStr, cookie))
	if err != nil {
		return err
	}
	return nil
}

// CheckExpireCookieInRedis : Check Expire Cookie In Redis
func CheckExpireCookieInRedis() {
	con := pool.Get()
	defer con.Close()

	timestamp := fmt.Sprintf(`%d`, time.Now().Unix())
	redis.Int64(con.Do("ZREMRANGEBYSCORE", "cookie2", "0", timestamp))
}
