package main

import (
	"cache2go"
	"go-concurrentMap-master"

	"code.google.com/p/go.net/websocket"
)

//====================== 游戏全局变量=============================================
// 一个管理指针的结构
var runningGoldMinerRoom *GoldMinerRoom

// map的key 和user没有任何关系
type GoldMinerRoom struct {
	OnlineUsers map[string]*OnlineUser
}

// 在线玩家的数据的结构体;后面优化成每个商家对应一个结构
type OnlineUser struct {
	Connection *websocket.Conn           // 链接的信息
	StrMD5     string                    // PC的标识，主要是针对PC的绑定数据；；区分所有的玩家
	MapSafe    *concurrent.ConcurrentMap // 并发安全的map
}

// 定义全局的变量的数据
var cache *cache2go.CacheTable  // 硬件存储
var M *concurrent.ConcurrentMap // 并发安全的map
