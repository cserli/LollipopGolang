package Redis_DB

//import (
//	"encoding/json"
//	"fmt"
//	"glog-master"
//	"strconv"
//	"strings"
//	"time"
//	"zLollipopGolang/globalData"
//	"zLollipopGolang/playerstruct"
//)

//////////////////////////////////////////////////////////////////////////////////
////
////                      内存数据库的写入的操作
////
////
////
//////////////////////////////////////////////////////////////////////////////////

//// 游戏的GameInfo数据写入到内存
//func Redis_Write_GameInfo() bool {

//	//	glog.Info("Entry Redis_Write_GameInfo")
//	//	glog.Info("Entry Redis_Write_GameInfo")

//	//	// 测试写如数据库 ok
//	//	error := GetREDIS().Redis_Client.Hset("xiaoF_Game", "xiaoF_Game1", []byte("sadsd"))
//	//	glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error)

//	//	beytr, _ := GetREDIS().Redis_Client.Hget("xiaoF_Game", "xiaoF_Game1")
//	//	glog.Info("GetREDIS().Redis_Client.Hset: Set data ", string(beytr))

//	//	// 测试数据：
//	//	playerinfo := PlayerData.Plyerdata{
//	//		PlayerMoney: 1100000,
//	//	}

//	//	Redis_Write_PlayerInfo("1222", playerinfo)
//	// 读取内存的数据
//	//	strert := Redis_Read_PlayerInfo(44)
//	//	glog.Info("redis  read :", strert.PlayerName)
//	return true
//}

////玩家登陆数据的保存:直接数据的操作内存数据库
////hkey ：为 玩家的AccountID
//func Redis_Write_PlayerInfo(key string, playerinfo PlayerData.Plyerdata) bool {

//	glog.Info("Entry Redis_Write_PlayerInfo")
//	glog.Info("Entry Redis_Write_PlayerInfo")
//	p, err := json.Marshal(playerinfo)
//	if err != nil {
//		glog.Info(err.Error())
//		glog.Info(err.Error())
//	}

//	// 写入玩家数据到内存数据库
//	error := GetREDIS().Redis_Client.Hset(PlayerLoginKey, key, p)
//	if error != nil {
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error)
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error.Error())
//		return false
//	}
//	return true
//}

//// 所有的玩家的游戏结束后的数据
//// hkey ： 为 数据量的大小 hkey：数字
//func Redis_Write_SaveGoldMinerGameOverData(strkey string) bool {
//	//   服务器获取数据
//	glog.Info("Entry Redis_Write_SaveGoldMinerGameOverData")
//	var count int
//	// 获取内存数据的数据信息
//	// 1 获取数据信息 内存的数据库的信息
//	data_byte, _ := GetREDIS().Redis_Client.Hget(PlayerGameOverKey, strkey)
//	tmpstring := string(data_byte)
//	count, _ = strconv.Atoi(tmpstring)
//	count = count + 1
//	// 写入玩家数据到内存数据库
//	error1 := GetREDIS().Redis_Client.Hset(PlayerGameOverKey, strkey, []byte(strconv.Itoa(count)))
//	if error1 != nil {
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1)
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1.Error())
//		return false
//	}

//	return true
//}

//// 玩家签到数据保存到内存数据库
//// hkey ： 为 每月的第几天 如12月的第一天 hkey：1
//func Redis_Write_PlayerSignInDay(strkey string) bool {

//	glog.Info("Entry Redis_Write_PopularityRecordData")
//	var count int
//	// 获取内存数据的数据信息
//	// 1 获取数据信息 内存的数据库的信息
//	data_byte, _ := GetREDIS().Redis_Client.Hget(PlayerSinginDayKey, strkey)
//	tmpstring := string(data_byte)
//	count, _ = strconv.Atoi(tmpstring)
//	count = count + 1
//	// 写入玩家数据到内存数据库
//	error1 := GetREDIS().Redis_Client.Hset(PlayerSinginDayKey, strkey, []byte(strconv.Itoa(count)))
//	if error1 != nil {
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1)
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1.Error())
//		return false
//	}

//	//	// 保存数据到数据库，更新数据库的字段数据
//	//	tmpint, _ := strconv.Atoi(strkey)
//	//	dbif.UpdatePlayerSignInDataFun(int64(tmpint), count)

//	return true
//}

//// 玩家点击游戏后数据的处理，记录游戏的人气的数据
//// hkey ： 为 游戏的ID： GameID   为官方的数据
//func Redis_Write_PopularityRecordData(strkey string) bool {
//	glog.Info("Entry Redis_Write_PopularityRecordData")
//	var count int
//	// 获取内存数据的数据信息
//	// 1 获取数据信息 内存的数据库的信息
//	data_byte, _ := GetREDIS().Redis_Client.Hget(PopularityRecordKey, strkey)
//	tmpstring := string(data_byte)
//	count, _ = strconv.Atoi(tmpstring)
//	count = count + 1
//	// 写入玩家数据到内存数据库
//	error1 := GetREDIS().Redis_Client.Hset(PopularityRecordKey, strkey, []byte(strconv.Itoa(count)))
//	if error1 != nil {
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1)
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1.Error())
//		return false
//	}

//	return true
//}

//// 用户赠送豆豆的时候产生的的一条记录数据
//// hkey ： 为 账单记录产生的时间戳： Instertime
//func Redis_Write_HistoryPayInfoData(strkey_Instertime string, stPayInfo Global_Define.StHistoryPayInfo) bool {
//	glog.Info("Entry Redis_Write_HistoryPayInfoData")
//	glog.Info("Entry Redis_Write_HistoryPayInfoData")

//	// 数据的存储
//	itimekey, _ := strconv.Atoi(stPayInfo.Instertime)
//	bytetmp, _ := json.Marshal(stPayInfo)
//	_, error := GetREDIS().Redis_Client.Zadd(HistoryPayInfoKey, float64(itimekey), bytetmp)
//	if error != nil {
//		glog.Info("GetREDIS().Redis_Client.Zadd: Set data ", error)
//		glog.Info("GetREDIS().Redis_Client.Zadd: Set data ", error.Error())
//		return false
//	}
//	return true
//}

//// 用户产生一条问答
//// hkey ： 为 问答产生的时间戳： Instertime
//func Redis_Write_AnswerInfoData(strkey_Instertime string, stAnswerInfo Global_Define.StAnswerInfo) bool {
//	glog.Info("Entry Redis_Write_AnswerInfoData")
//	glog.Info("Entry Redis_Write_AnswerInfoData")

//	// 数据的存储
//	itimekey, _ := strconv.Atoi(stAnswerInfo.Instertime)
//	bytetmp, _ := json.Marshal(stAnswerInfo)
//	_, error := GetREDIS().Redis_Client.Zadd(AnswerInfoKey, float64(itimekey), bytetmp)
//	if error != nil {
//		glog.Info("GetREDIS().Redis_Client.Zadd: Set data ", error)
//		glog.Info("GetREDIS().Redis_Client.Zadd: Set data ", error.Error())
//		return false
//	}
//	return true
//}

//// 用户产生一条问答的评论
//// hkey ： 为 问答评论产生的时间戳： Instertime
//func Redis_Write_AnswerCommentInfoData(strkey_Instertime string, stAnswerCommentInfo Global_Define.StCommentInfo) bool {
//	glog.Info("Entry Redis_Write_AnswerCommentInfoData")
//	glog.Info("Entry Redis_Write_AnswerCommentInfoData")

//	// 数据的存储
//	itimekey, _ := strconv.Atoi(stAnswerCommentInfo.CommentTime)
//	bytetmp, _ := json.Marshal(stAnswerCommentInfo)
//	_, error := GetREDIS().Redis_Client.Zadd(AnswerCommentInfoKey, float64(itimekey), bytetmp)
//	if error != nil {
//		glog.Info("GetREDIS().Redis_Client.Zadd: Set data ", error)
//		glog.Info("GetREDIS().Redis_Client.Zadd: Set data ", error.Error())
//		return false
//	}
//	return true
//}

//// 用户产生一篇文章
//// hkey ： 为 问答产生的时间戳： Instertime
//func Redis_Write_ArticleInfoData(strkey_Instertime string, StArticleInfo Global_Define.StArticleInfo) bool {
//	glog.Info("Entry Redis_Write_ArticleInfoData")
//	glog.Info("Entry Redis_Write_ArticleInfoData")

//	// 数据的存储
//	itimekey, _ := strconv.Atoi(StArticleInfo.ArticleIssueTime)
//	bytetmp, _ := json.Marshal(StArticleInfo)

//	var GoAnguageInfoKey string

//	// 首先判断是那个圈子的数据的保存
//	if StArticleInfo.ICircleId == "100" {
//		GoAnguageInfoKey = "GoAnguage_Base" // GO 语言基础知识

//	} else if StArticleInfo.ICircleId == "101" {
//		GoAnguageInfoKey = "GoAnguage_NetWork" // GO 语言网路知识

//	} else if StArticleInfo.ICircleId == "102" {
//		GoAnguageInfoKey = "GoAnguage_DataBase" // GO 语言数据库知识

//	} else if StArticleInfo.ICircleId == "103" {
//		GoAnguageInfoKey = "GoAnguage_MemoryDataBase" // GO 语言内存数据库知识

//	} else if StArticleInfo.ICircleId == "104" {
//		GoAnguageInfoKey = "GoAnguage_Server" // GO 语言服务器开发

//	} else if StArticleInfo.ICircleId == "105" {
//		GoAnguageInfoKey = "GoAnguage_OfficialRewardMK" // GO 语言官方模块悬赏，不一定是模块，也可以是工具的开发的悬赏

//	} else if StArticleInfo.ICircleId == "106" {
//		GoAnguageInfoKey = "GoAnguage_OfficialRewardXM" // GO 语言官方项目悬赏，项目的进度的悬赏。
//	}

//	_, error := GetREDIS().Redis_Client.Zadd(GoAnguageInfoKey, float64(itimekey), bytetmp)
//	if error != nil {
//		glog.Info("GetREDIS().Redis_Client.Zadd: Set data ", error)
//		glog.Info("GetREDIS().Redis_Client.Zadd: Set data ", error.Error())
//		return false
//	}
//	return true
//}

//// 用户产生一条问答的评论
//// hkey ： 为 问答评论产生的时间戳： Instertime
//func Redis_Write_ArticleCommentInfoData(strkey_Instertime string, icircleid int, stAnswerCommentInfo Global_Define.StCommentInfo) bool {
//	glog.Info("Entry Redis_Write_ArticleCommentInfoData")
//	glog.Info("Entry Redis_Write_ArticleCommentInfoData")

//	// 数据的存储
//	itimekey, _ := strconv.Atoi(stAnswerCommentInfo.CommentTime)
//	bytetmp, _ := json.Marshal(stAnswerCommentInfo)

//	var GoAnguageInfoKey string

//	// 首先判断是那个圈子的数据的保存
//	if strconv.Itoa(icircleid) == "100" {
//		GoAnguageInfoKey = "GoAnguage_Base" // GO 语言基础知识

//	} else if strconv.Itoa(icircleid) == "101" {
//		GoAnguageInfoKey = "GoAnguage_NetWork" // GO 语言网路知识

//	} else if strconv.Itoa(icircleid) == "102" {
//		GoAnguageInfoKey = "GoAnguage_DataBase" // GO 语言数据库知识

//	} else if strconv.Itoa(icircleid) == "103" {
//		GoAnguageInfoKey = "GoAnguage_MemoryDataBase" // GO 语言内存数据库知识

//	} else if strconv.Itoa(icircleid) == "104" {
//		GoAnguageInfoKey = "GoAnguage_Server" // GO 语言服务器开发

//	} else if strconv.Itoa(icircleid) == "105" {
//		GoAnguageInfoKey = "GoAnguage_OfficialRewardMK" // GO 语言官方模块悬赏，不一定是模块，也可以是工具的开发的悬赏

//	} else if strconv.Itoa(icircleid) == "106" {
//		GoAnguageInfoKey = "GoAnguage_OfficialRewardXM" // GO 语言官方项目悬赏，项目的进度的悬赏。
//	}

//	_, error := GetREDIS().Redis_Client.Zadd(GoAnguageInfoKey, float64(itimekey), bytetmp)
//	if error != nil {
//		glog.Info("GetREDIS().Redis_Client.Zadd: Set data ", error)
//		glog.Info("GetREDIS().Redis_Client.Zadd: Set data ", error.Error())
//		return false
//	}
//	return true
//}

//// redis的数据更新操作
//// 数据更新：
//// redis 127.0.0.1:6379> ZADD myset 1 "hello"
//// (integer) 1
//// redis 127.0.0.1:6379> ZADD myset 1 "foo"
//// (integer) 1
//// redis 127.0.0.1:6379> ZADD myset 2 "world" 3 "bar"
//// (integer) 2

//// redis 127.0.0.1:6379> ZRANGE myset 0 -1 WITHSCORES  //正序
//// 1) "hello"
//// 2) "1"
//// 3) "foo"
//// 4) "1"
//// 5) "world"
//// 6) "2"
//// 7) "bar"
//// 8) "3"
//// zrevrange myzset 0 -1 WITHSCORES  // 反序
////======================================================================================
//// 玩家点击报名处理，有报名成功的加1
//// hkey ： 为商家的信息： 二2key   为官方的活动的数据
//func Redis_Write_YaoYiYaoBaoMingData(strkey string) bool {
//	glog.Info("Entry Redis_Write_YaoYiYaoBaoMingData")
//	var count int
//	// 获取内存数据的数据信息
//	// 1 获取数据信息 内存的数据库的信息
//	data_byte, _ := GetREDIS().Redis_Client.Hget(G_BaoMingKey, strkey)
//	tmpstring := string(data_byte)
//	count, _ = strconv.Atoi(tmpstring)
//	count = count + 1
//	// 写入玩家数据到内存数据库
//	error1 := GetREDIS().Redis_Client.Hset(G_BaoMingKey, strkey, []byte(strconv.Itoa(count)))
//	if error1 != nil {
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1)
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1.Error())
//		return false
//	}

//	return true
//}

//// 玩家摇一摇处理，每摇一次加1
//// hkey ： 为商家的信息+现场名称： 二2key   OpenID
//func Redis_Write_YaoYiYaoJiShuData(strkey string, iYaoCiShu uint32) bool {
//	glog.Info("Entry Redis_Write_YaoYiYaoJiShuData")
//	glog.Info("Entry Redis_Write_YaoYiYaoJiShuData")
//	var count = 0
//	// 获取内存数据的数据信息
//	// 1 获取数据信息 内存的数据库的信息
//	data_byte, _ := GetREDIS().Redis_Client.Hget(G_DataSaveDongKey, strkey)
//	tmpstring := string(data_byte)
//	count, _ = strconv.Atoi(tmpstring)
//	glog.Info("---------------------------------------------------------")
//	glog.Info("内存保存的数据：count", count)
//	count = count + 1
//	if count == 0 {
//		return false
//	}
//	// 写入玩家数据到内存数据库
//	error1 := GetREDIS().Redis_Client.Hset(G_DataSaveDongKey, strkey, []byte(strconv.Itoa(int(iYaoCiShu))))
//	if error1 != nil {
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1)
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1.Error())
//		return false
//	}

//	data_byte1, _ := GetREDIS().Redis_Client.Hget(G_DataSaveDongKey, strkey)
//	tmpstring1 := string(data_byte1)
//	count1, _ := strconv.Atoi(tmpstring1)
//	glog.Info("内存保存后的数据：count", count1)
//	glog.Info("---------------------------------------------------------")

//	return true
//}

//// 玩家摇一摇到达终点，每到一位加1
//// hkey ： GameOver+商家名字+场次： 二2key   GameOver
//func Redis_Write_YaoYiYaoGameOver3Data(strkey string) bool {
//	glog.Info("Entry Redis_Write_YaoYiYaoJiShuData")
//	var count int
//	// 获取内存数据的数据信息
//	// 1 获取数据信息 内存的数据库的信息
//	data_byte, _ := GetREDIS().Redis_Client.Hget(G_TimerGameOverKey, strkey)
//	tmpstring := string(data_byte)
//	count, _ = strconv.Atoi(tmpstring)
//	count = count + 1
//	// 写入玩家数据到内存数据库
//	error1 := GetREDIS().Redis_Client.Hset(G_TimerGameOverKey, strkey, []byte(strconv.Itoa(count)))
//	if error1 != nil {
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1)
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1.Error())
//		return false
//	}

//	return true
//}

//// 摇一摇的排行榜的实现
//// hkey ： 为 用户摇一摇的数据： 步数
//func Redis_Write_YaoYiYaoPaiHangBangInfoData(stPaiHangBang Global_Define.StPaiHangBang, iPaiMing uint32) bool {
//	glog.Info("Entry Redis_Write_YaoYiYaoPaiHangBangInfoData")
//	glog.Info("Entry Redis_Write_YaoYiYaoPaiHangBangInfoData")

//	//	iopenid := stPaiHangBang.OpenID
//	//	iiopendi, err := strconv.Atoi(iopenid)
//	//	if err != nil {
//	//		glog.Info("Redis_Write_YaoYiYaoPaiHangBangInfoData", err.Error())
//	//	}
//	// 数据的存储
//	bytetmp, _ := json.Marshal(stPaiHangBang)
//	_, error := GetREDIS().Redis_Client.Zadd(G_PaiHangKey, float64(iPaiMing), bytetmp)
//	if error != nil {
//		glog.Info("GetREDIS().Redis_Client.Zadd: Set data ", error)
//		glog.Info("GetREDIS().Redis_Client.Zadd: Set data ", error.Error())
//		return false
//	}
//	return true
//}

///////////////////////////////////////抽奖///////////////////////////////////////
//// 抽奖CD时间的保存
//func Redis_Write_ChouJiangCDTimeInfoData(strCDtime string, stropenid string, StrMd5 string) {

//	glog.Info("Entry Redis_Write_ChouJiangCDTimeInfoData!!!")

//	// 写入玩家数据到内存数据库
//	error1 := GetREDIS().Redis_Client.Hset(CDTime+"yiy"+StrMd5, stropenid, []byte(strCDtime))
//	if error1 != nil {
//		glog.Info("=========================-----------SaveCDTime--------------==============================tmpstring", error1)
//		return
//	}
//	return
//}

///////////////////////////////////////更新二维码的数据///////////////////////////////////////
//func Redis_Write_EWMResData(StrMd5 string) bool {

//	glog.Info("Entry Redis_Write_EWMResData!!!")
//	var count int = 1
//	// 写入玩家数据到内存数据库
//	error1 := GetREDIS().Redis_Client.Hset("EWN_Update"+StrMd5, StrMd5, []byte(strconv.Itoa(count)))
//	if error1 != nil {
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1.Error())
//		return false
//	}
//	return true
//}

///////////////////////////////////////设置游戏在经行中///////////////////////////////////////
//func Redis_Write_Game_Running_Data(strL_G_X_Name string, strAllName string) bool {

//	glog.Info("Entry Redis_Write_Game_Running_Data!!!")
//	if len(strAllName) == 0 {
//		glog.Info("strAllName is nil!!!")
//		return false
//	}

//	//	var count int
//	//	count = 1
//	StrFirmUser := Global_Define.MD5ToStringFromString(strL_G_X_Name)
//	// 删除数据先
//	// GetREDIS().Redis_Client.Del("Game_Running" + "|" + StrFirmUser)
//	// 写入玩家数据到内存数据库
//	glog.Info("------------------------------------------")
//	glog.Info("------------------------------------------strL_G_X_Name", strL_G_X_Name)
//	glog.Info("------------------------------------------strAllName", strAllName)
//	glog.Info("------------------------------------------StrFirmUser", StrFirmUser)
//	glog.Info("------------------------------------------[]byte(strAllName)", string([]byte(strAllName)))
//	glog.Info("------------------------------------------")
//	// 写入玩家数据到内存数据库
//	error1 := GetREDIS().Redis_Client.Hset("Game_Running"+"|"+StrFirmUser, strL_G_X_Name, []byte(strAllName))
//	if error1 != nil {
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1.Error())
//		return false
//	}
//	return true
//}

/////////////////////////////////////////次数///////////////////////////////////////
////// 抽奖CD次数
////func Redis_Write_ChouJiangCDCishuInfoData(stropenid string, StrMd5 string) bool {

////	glog.Info("Entry Redis_Write_ChouJiangCDCishuInfoData!!!")
////	var count int = 0
////	// 获取内存数据的数据信息
////	// 1 获取数据信息 内存的数据库的信息
////	data_byte, _ := GetREDIS().Redis_Client.Hget(CDTime+"yiycishu"+StrMd5, stropenid)
////	tmpstring := string(data_byte)
////	glog.Info("=========================-------------------------==============================tmpstring" + tmpstring)
////	count, _ = strconv.Atoi(tmpstring)
////	count = count + 1
////	// 写入玩家数据到内存数据库
////	error1 := GetREDIS().Redis_Client.Hset(CDTime+"yiycishu"+StrMd5, stropenid, []byte(strconv.Itoa(count)))
////	if error1 != nil {
////		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1.Error())
////		return false
////	}
////	return true
////}

///////////////////////////////////////次数///////////////////////////////////////
//// 抽奖CD时间的保存
//func Redis_Write_ChouJiangCDCishuInfoData(stropenid string) bool {

//	glog.Info("Entry Redis_Write_ChouJiangCDCishuInfoData!!!")
//	iret, err := GetREDIS().Redis_Client.Incr(stropenid)
//	if err != nil {
//		glog.Info("Entry Redis_Write_ChouJiangCDCishuInfoData err", err)
//		return false
//	}
//	retime := Redis_key_Expire()
//	retime = (24 - retime) * 60 * 60
//	GetREDIS().Redis_Client.Expire(stropenid, int64(retime))
//	glog.Info("Entry Redis_Write_ChouJiangCDCishuInfoData ", iret)
//	return true

//}

//// 返回当前时间
//func Redis_key_Expire() int {

//	strime := time.Now().Format("2006-01-02 15:04:05")
//	rs := []rune(strime)
//	fmt.Println("rs:", string(rs[11:13]))
//	//fmt(string(rs[11:13]))
//	iret, _ := strconv.Atoi(string(rs[11:13]))
//	return iret
//}

///////////////////////////////////////保存用户数据到内存数据库///////////////////////////////////////
//// 商家对应的所有的游戏
//func Redis_Write_WeiXinPlayData_InRedis_Data(StrLoginName string, StrGameName string, StrXCName string, StWeiXinData Global_Define.StWeiXinUserInfo) bool {

//	////	msuo.Lock()
//	//	defer msuo.Unlock()
//	StrLoginNameMD5 := Global_Define.MD5ToStringFromString(StrLoginName)
//	iLastLoginTime := time.Now().UnixNano()
//	// 架构转化
//	byteStWeiXinData, _ := json.Marshal(StWeiXinData)
//	strLastLoginTime := strconv.FormatInt(iLastLoginTime, 10)
//	if GetREDIS().Redis_Client == nil {
//		Redis_ConnFun()
//	}
//	// 写入玩家数据到内存数据库
//	error1 := GetREDIS().Redis_Client.Hset("Game_WeiXinData"+"|"+StrLoginNameMD5, strLastLoginTime, byteStWeiXinData)
//	if error1 != nil {
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1.Error())
//		return false
//	}
//	return true
//}

///////////////////////////////////////保存用户数据到内存数据库///////////////////////////////////////
//// 商家对应的某个的游戏
//func Redis_Write_WeiXinPlayData_InRedis_Data_ByLoginOrGame(StrLoginName string, StrGameName string, StrXCName string, StWeiXinData Global_Define.StWeiXinUserInfo) bool {

//	//	msuo.Lock()
//	//	defer msuo.Unlock()
//	StrLoginNameMD5 := Global_Define.MD5ToStringFromString(StrLoginName + StrGameName)
//	iLastLoginTime := time.Now().UnixNano()
//	// 架构转化
//	byteStWeiXinData, _ := json.Marshal(StWeiXinData)
//	strLastLoginTime := strconv.FormatInt(iLastLoginTime, 10)
//	if GetREDIS().Redis_Client == nil {
//		Redis_ConnFun()
//	}
//	// 写入玩家数据到内存数据库
//	error1 := GetREDIS().Redis_Client.Hset("Game_WeiXinData"+"|"+StrLoginNameMD5, strLastLoginTime, byteStWeiXinData)
//	if error1 != nil {
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1.Error())
//		return false
//	}
//	return true
//}

///////////////////////////////////////保存用户数据到内存数据库///////////////////////////////////////
//// 商家对应的某个的游戏的某个现场
//func Redis_Write_WeiXinPlayData_InRedis_Data_ByLoginOrGameOrXC(StrLoginName string, StrGameName string, StrXCName string, StWeiXinData Global_Define.StWeiXinUserInfo) bool {

//	//	msuo.Lock()
//	//	defer msuo.Unlock()
//	StrLoginNameMD5 := Global_Define.MD5ToStringFromString(StrLoginName + StrGameName + StrXCName)
//	iLastLoginTime := time.Now().UnixNano()
//	// 架构转化
//	byteStWeiXinData, _ := json.Marshal(StWeiXinData)
//	strLastLoginTime := strconv.FormatInt(iLastLoginTime, 10)
//	if GetREDIS().Redis_Client == nil {
//		Redis_ConnFun()
//	}
//	// 写入玩家数据到内存数据库
//	error1 := GetREDIS().Redis_Client.Hset("Game_WeiXinData"+"|"+StrLoginNameMD5, strLastLoginTime, byteStWeiXinData)
//	if error1 != nil {
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1.Error())
//		return false
//	}
//	return true
//}

//// 字符串分割函数
//func Strings_Split(Data string, Split string) []string {
//	return strings.Split(Data, Split)
//}

///////////////////////////////////////积分///////////////////////////////////////
//// 游戏积分
//func Redis_Write_Player_JiFenDatabka(stropenid string, iJiFen int) bool {

//	glog.Info("Entry Redis_Write_Player_JiFenData!!!")
//	if len(stropenid) == 0 || iJiFen <= 0 {
//		return false
//	}
//	//////////////////////
//	// 判断数据库存在不
//	var StrOpenID = ""
//	// 差分key
//	strsplit := Strings_Split(stropenid, "|")
//	for i := 0; i < len(strsplit); i++ {
//		if i == 0 {
//			StrOpenID = strsplit[i]
//		} else if i == 1 {

//		}
//	}
//	//////////////////////
//	var count int = 0
//	// 获取内存数据的数据信息
//	// 1 获取数据信息 内存的数据库的信息
//	data_byte, _ := GetREDIS().Redis_Client.Get(StrOpenID)
//	tmpstring := string(data_byte)
//	glog.Info("=========================-------------------------==============================tmpstring" + tmpstring)
//	count, _ = strconv.Atoi(tmpstring)
//	count = count + iJiFen
//	// 写入玩家数据到内存数据库
//	error1 := GetREDIS().Redis_Client.Set(StrOpenID, []byte(strconv.Itoa(count)))
//	if error1 != nil {
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1.Error())
//		return false
//	}
//	return true
//}

///////////////////////////////////////积分///////////////////////////////////////
//// 游戏积分
//func Redis_Write_Player_JiFenData(stropenid string, iJiFen int) bool {

//	glog.Info("Entry Redis_Write_Player_JiFenData!!!")
//	if len(stropenid) == 0 || iJiFen <= 0 {
//		return false
//	}

//	//	msuo.Lock()
//	//	defer msuo.Unlock()
//	//////////////////////
//	// 判断数据库存在不
//	var StrOpenID = ""
//	// 差分key
//	strsplit := Strings_Split(stropenid, "|")
//	for i := 0; i < len(strsplit); i++ {
//		if i == 0 {
//			StrOpenID = strsplit[i]
//		}
//	}
//	//////////////////////
//	var countq int = 0
//	var counth int = 0
//	var iret int = 0
//	// 获取内存数据的数据信息
//	// 1 获取数据信息 内存的数据库的信息
//	countq, iret = Redis_Read_Player_JiFenData(StrOpenID)
//	//countq = dbif.Get_JiFen_System_Data(StrOpenID)
//	glog.Info("GetREDIS().Redis_Client.Hset: 积分前： ", countq)
//	counth = countq + iJiFen
//	glog.Info("GetREDIS().Redis_Client.Hset: 积分后： ", counth)
//	if iret == 0 || countq == counth {
//		return false
//	}
//	//	if countq == counth {
//	//		return false
//	//	}
//	// 写入玩家数据到内存数据库
//	error1 := GetREDIS().Redis_Client.Hset("JiFenTJ", StrOpenID, []byte(strconv.Itoa(counth)))
//	if error1 != nil {
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1.Error())
//		return false
//	}
//	return true
//}

//// 用户获取积分的数据的保存
//// hkey ： 为 问答产生的时间戳： Instertime
//func Redis_Write_JiFenInfoData(stJiFenLogInfo Global_Define.StJiFenLogInfo) bool {

//	//	msuo.Lock()
//	//	defer msuo.Unlock()
//	glog.Info("111111111111111111111111111111")
//	glog.Info(stJiFenLogInfo)
//	glog.Info("111111111111111111111111111111")
//	// 数据的存储
//	itimekey, _ := strconv.Atoi(stJiFenLogInfo.Instertime)
//	bytetmp, _ := json.Marshal(stJiFenLogInfo)
//	_, error := GetREDIS().Redis_Client.Zadd("JiFenLogInfo", float64(itimekey), bytetmp)
//	if error != nil {
//		glog.Info("GetREDIS().Redis_Client.Zadd: Set data ", error.Error())
//		return false
//	}
//	return true
//}

//// 游戏中获取卡片
//func Redis_Write_GetCardInfo_Of_Game(stropenid string, data string) {

//	glog.Info("Entry Redis_Write_GetCardInfo_Of_Game!!!")
//	// 判断数据库存在不
//	var StrMD5 = ""
//	// 差分key
//	strsplit := Strings_Split(stropenid, "|")
//	for i := 0; i < len(strsplit); i++ {
//		if i == 0 {
//			StrMD5 = strsplit[i]
//		}
//	}
//	strdata := Redis_Read_GetCardInfo_Of_Game(StrMD5)

//	if len(strdata) == 0 {
//		data = data
//	} else {
//		data = strdata + "|" + data
//	}

//	// 写入玩家数据到内存数据库
//	error1 := GetREDIS().Redis_Client.Hset("Player_GetItem", StrMD5, []byte(data))
//	if error1 != nil {
//		glog.Info("=========================-----------Redis_Write_GetCardInfo_Of_Game--------------==============================tmpstring", error1)
//		return
//	}
//	return
//}
