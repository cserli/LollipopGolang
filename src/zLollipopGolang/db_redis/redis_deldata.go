package Redis_DB

//	"encoding/json"
//	"glog-master"
//	"zLollipopGolang/playerstruct"

// 删除内存数据的数据 ; 会存在问题
func Redis_Del_PlayerInfo(key string) bool {

	//	glog.Info("Entry Redis_Del_PlayerInfo")
	//	// 测试数据：
	//	playerinfo := PlayerData.Plyerdata{}
	//	p, err := json.Marshal(playerinfo)
	//	if err != nil {
	//		glog.Info(err.Error())
	//	}

	//	// 写入玩家数据到内存数据库
	//	error := GetREDIS().Redis_Client.Hset(PlayerLoginKey, key, p)
	//	if error != nil {
	//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error.Error())
	//		return false
	//	}
	return true
}

// 删除签到的数据
func Redis_Del_SigninData() {
	// 清除内存数据库库的数据
	GetREDIS().Redis_Client.Del(PlayerSinginKey)
	GetREDIS().Redis_Client.Del("shangjia1xianchangyi1")
	return
}

//==============================================================================
// 清楚记录数据
func Redis_Del_GameOver3PData() {
	// 清除内存数据库库的数据
	GetREDIS().Redis_Client.Del("GameOvershangjia1xianchangyi1")
	GetREDIS().Redis_Client.Del("shangjia1")
	// GetREDIS().Redis_Client.Del("YaoYiYaoPaiHang")
	GetREDIS().Redis_Client.Del("shangjia1xianchangyi1")
	GetREDIS().Redis_Client.Del("huodong1")
	return
}

// 删除内存数据库CD次数
func Redis_Del_Game_CDCIShu_Data(StrMD5 string) bool {
	//	glog.Info("Entry Redis_Del_Game_CDCIShu_Data StrMD5", StrMD5)
	//	if len(StrMD5) == 0 {
	//		glog.Info("Entry Redis_Del_Game_CDCIShu_Data len(StrMD5) == 0 !!!")
	//		return false
	//	}

	//	if GetREDIS().Redis_Client == nil {
	//		glog.Info("Entry Redis_Del_Game_CDCIShu_Data GetREDIS().Redis_Client is nil!!!")
	//		Redis_ConnFun()
	//	}
	//	bret, err := GetREDIS().Redis_Client.Del(CDTime + "yiycishu" + StrMD5)
	//	if err != nil {
	//		glog.Info(" GetREDIS().Redis_Client.Del,StrMD5:", StrMD5, "error:", err)
	//	}
	//	if bret == true {
	//		glog.Info(" GetREDIS().Redis_Client.Del,SUCC")
	//	}
	return true
}
