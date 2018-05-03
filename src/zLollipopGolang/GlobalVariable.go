package main

import (
	"cache2go"
	"concurrent-map-master"
	"go-concurrentMap-master"
	"sync"
	"zLollipopGolang/globalData"

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

// csv数据结构定义
var G_GameInfoST map[string]*Global_Define.StGameListInfo            // StGameListInfo数据
var G_ActivitiesInfoST map[string]*Global_Define.StActivitiesInfocsv // StActivitiesInfo数据
var G_GameConfigInfoST map[string]*Global_Define.StGameConfigInfo    // StGameConfigInfo数据

//-----------------------------------------------------------------------
// serve一个管理指针的结构
var runningServerRoom *ServerRoom

//-----------------------------------------------------------------------
var PlayerNum int = 0

// map的key 和user没有任何关系
type ServerRoom struct {
	OnlineServers map[string]*OnlineServer
}

// 在线玩家的数据的结构体
type OnlineServer struct {
	Connection        *websocket.Conn           // 链接的信息
	OnlineServersLock sync.RWMutex              // 读写锁
	StrMD5            string                    // 游戏的标识，游戏的数据
	M                 *concurrent.ConcurrentMap // 并发安全的map
}

// 定义全局的变量的数据
var cache *cache2go.CacheTable  // 硬件存储
var Gmap cmap.ConcurrentMap     // 并发安全的map
var M *concurrent.ConcurrentMap // 并发安全的map
var ServerURl string = ""

// 游戏中需要的结构
var MapG_LogoTime map[string]string
var MapG_EWMTime map[string]string

/////////////////////

var GameType = "1"

var G_StXianChangInfotmp map[string]*Global_Define.StXianChangInfo // 主动推送数据
var G_StWeiXinDatatmp map[string]*Global_Define.StWeiXinUserInfo   // 微信的用户的结构信息
