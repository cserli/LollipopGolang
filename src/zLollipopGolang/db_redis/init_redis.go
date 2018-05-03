package redis

import (
	_ "glog-master"
	"time"

	"github.com/go-redis/redis" // 内存数据库--用于测试
)

var client *redis.Client

/*初始化*/
func init() {
	client = redis.NewClient(&redis.Options{
		Addr:         "---:6379",
		Password:     "",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     1000,
		PoolTimeout:  30 * time.Second,
	})
	client.FlushDB()
}
