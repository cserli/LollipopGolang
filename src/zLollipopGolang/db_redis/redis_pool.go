package Redis_DB

import (
	//	"glog-master"

	"github.com/go-redis/redis" // 内存数据库--用于测试
	//"github.com.Go-Redis/github.com/alphazero/Go-Redis"
)

// 链接池结构体
type STRedis_Pool struct {
	Redis_Client redis.Client
}

// 链接池的最大链接数量
const MAX_REDIS_POOL_SIZE int = 100

// 全局内存数据库变量
var REDISPool chan *STRedis_Pool

// 获取数据链接
func GetREDIS() *STRedis_Pool {
	// 获取链接
	conn := getREDIS()
	// 压入队列
	putREDIS(conn)
	return conn
}

// 获取链接指针函数
func getREDIS() *STRedis_Pool {
	if REDISPool == nil {
		REDISPool = make(chan *STRedis_Pool, MAX_REDIS_POOL_SIZE)
	}

	return <-REDISPool
}

//存储指针函数
func putREDIS(conn *STRedis_Pool) {
	if REDISPool == nil {
		REDISPool = make(chan *STRedis_Pool, MAX_REDIS_POOL_SIZE)
	}
	if len(REDISPool) == MAX_REDIS_POOL_SIZE {
		//conn.Redis_Client.Close()
		return
	}
	REDISPool <- conn
}
