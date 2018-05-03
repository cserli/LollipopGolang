package Redis_DB

import (
	//"github.com.Go-Redis/github.com/alphazero/Go-Redis"
	"github.com/go-redis/redis" // 内存数据库--用于测试
)

////////////////////////////////////////////////////////////////////////////////
//
//                        内存数据库链接
//
//
////////////////////////////////////////////////////////////////////////////////

// 全局 redis 链接
var GRedis_Client1 redis.Client

// 链接内存数据库函数
func Redis_ConnFun() bool {

	return true
}

// 数据库ping（）程序
func GameDBRedisTimer() {

	//	GGameStartTimer := time.NewTicker(3 * time.Second)
	//	for {
	//		select {
	//		case <-GGameStartTimer.C:
	//			{
	//				// 数据库ping()
	//				if GRedis_Client != nil {

	//					err := GRedis_Client.Ping()
	//					if err != nil {
	//						glog.Info("db redis ping() err!")
	//						// 需要重新链接
	//						Redis_ConnFun()
	//					}
	//				} else {
	//					// 需要重新链接
	//					Redis_ConnFun()
	//				}
	//			}
	//		}
	//	}
}
