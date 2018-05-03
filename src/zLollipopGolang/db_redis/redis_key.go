package Redis_DB

////////////////////////////////////////////////////////////////////////////////
//
//
//                    服务器所有的Redis Key
//
//
////////////////////////////////////////////////////////////////////////////////
//服务器的部署方案
var G_ZS_Or_Test = "Server"

var G_GameNameType = "TTAPB" // 天天爱跑步  6001
//var G_GameNameType = "TTXS" // 天天向上    6002
//var G_GameNameType = "AYYS" // 奥运泳士    6003
//  摇一摇游戏
var (
	//G_GameNameType     = ""                               // 游戏的类型名字
	G_BaoMingKey       = "BaoMing" + G_GameNameType       // 报名== 游戏开始前就需要赋值
	G_PaiHangKey       = "PaiHang" + G_GameNameType       // 排行
	G_DataSaveDongKey  = "CaoZuo" + G_GameNameType        // 现场用户操作手机记录
	G_HuoDongKey       = "HuoDong" + G_GameNameType       // 活动
	G_TimerGameOverKey = "GameOverTimer" + G_GameNameType // 游戏结束
)

const CDTime = "CDTime"

////////////////////////////////////////////////////////////////////////////////

// 玩家登陆保存在内存中的key
const PlayerLoginKey = "GamePlayerData"

// 游戏记录玩家登陆人气保存在内存中的key
const PopularityRecordKey = "GamePopularityRecord"

// 签到的初始化数据保存在内存中的key
const PlayerSinginKey = "GamePlayerSingin"

// 玩家的签到的天数保存到内存中
const PlayerSinginDayKey = "GamePlayerSinginDay"

// 玩家游戏结束后保存到内存中
const PlayerGameOverKey = "GamePlayerGameOver"

// 玩家道具信息
const PlayerItemKey = "GamePlayerItem"

// 账单数据保存到内存中
const HistoryPayInfoKey = "HistoryPayInfo"

// 问答社区保存在内存数据库中
const AnswerCommentInfoKey = "AnswerCommentInfo"

// 问答社区的评论保存在内存数据库中
const AnswerInfoKey = "AnswerInfo"

//---------------------------------------------------
// 圈子的ID的内存数据库保存的Key,每个圈子对应一个唯一的Key
const (
	GoAnguage_BaseKey             = "GoAnguage_Base"             // GO 语言基础知识
	GoAnguage_NetWorkKey          = "GoAnguage_NetWork"          // GO 语言网路知识
	GoAnguage_DataBaseKey         = "GoAnguage_DataBase"         // GO 语言数据库知识
	GoAnguage_MemoryDataBaseKey   = "GoAnguage_MemoryDataBase"   // GO 语言内存数据库知识
	GoAnguage_ServerKey           = "GoAnguage_Server"           // GO 语言服务器开发
	GoAnguage_OfficialRewardMKKey = "GoAnguage_OfficialRewardMK" // GO 语言官方模块悬赏，不一定是模块，也可以是工具的开发的悬赏
	GoAnguage_OfficialRewardXMKey = "GoAnguage_OfficialRewardXM" // GO 语言官方项目悬赏，项目的进度的悬赏。
)

//---------------------------------------------------
