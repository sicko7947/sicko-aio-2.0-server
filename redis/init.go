package redis

import (
	"github.com/gogf/gf/database/gredis"
)

var (
	redisClient *gredis.Redis
)

func init() {
	config := &gredis.Config{
		Host: "127.0.0.1",
		Port: 6379,
		Db:   0,
	}
	groupName := "nike"
	gredis.SetConfig(config, groupName)
	redisClient = gredis.Instance(groupName)
}
