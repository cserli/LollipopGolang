package main

import (
	"zLollipopGolang/db_mysql"
	//	"zLollipopGolang/db_redis"
	"zLollipopGolang/globalData"
	//"zLollipopGolang/loglogic"
	"encoding/json"
	"glog-master"
	"zLollipopGolang/protocolfile"
	"zLollipopGolang/protocolfile/Proto2"
	//	"time"
)

// 摇一摇游戏的逻辑处理，主要是针对游戏
func (this *OnlineUser) HandleCltFirmsProtocol2(protocol2 interface{}, ProtocolData map[string]interface{}) {

	switch protocol2 {
	//	case float64(Proto2_Data.C2HS_FirmsUser_Login_Proto):
	//		{ // 商家登陆的协议
	//			Log_Eio.Log("protocol2 C2HS_FirmsUser_Login_Proto")
	//			// this.FirmsUser_Login_Fun(ProtocolData)
	//			break
	//		}
	default:
		glog.Info("protocol2 default")
	}
	return
}

//  主动推送数据给游戏服务器
func (this *OnlineUser) Server_Send_GameData_To_GameServer(stropenid string, StrMD5 string) {
	glog.Info("Entry Server_Send_GameData_To_GameServer")

	// 判断数据库存在不
	var loginname = ""
	var gamename = ""
	var xcname = ""
	// 差分key
	strsplit := Strings_Split(stropenid, "|")
	for i := 0; i < len(strsplit); i++ {
		if i == 0 {
			loginname = strsplit[i]
		} else if i == 1 {
			gamename = strsplit[i]
		} else if i == 2 {
			xcname = strsplit[i]
		}
	}
	var StXianChangInfotmp map[string]*Global_Define.StXianChangInfo
	var StXianChangInfotmp11 map[string]*Global_Define.StXianChangInfo
	var PaiHangBInfotmp map[string]*Global_Define.StAwardInfo
	PaiHangBInfotmp = make(map[string]*Global_Define.StAwardInfo)

	StXianChangInfotmp, _ = dbif.GetPayGamePublicizeInfoss(loginname, gamename, xcname, StrMD5)
	StXianChangInfotmp11, _ = dbif.GetPayGameAwardInfo(loginname, gamename, xcname, StrMD5)
	if StXianChangInfotmp11[StrMD5] == nil {
		glog.Info("主动推送数据给游戏服务器 Server_Send_GameData_To_GameServer StXianChangInfotmp11 is nil!!")
		return
	}
	if StXianChangInfotmp[StrMD5] == nil {
		glog.Info("主动推送数据给游戏服务器 Server_Send_GameData_To_GameServer StXianChangInfotmp is nil!!")
		return
	}
	json.Unmarshal([]byte(StXianChangInfotmp11[StrMD5].FirmsAward), &PaiHangBInfotmp)
	// 获取数据转义
	var PaiHangBInfotmpxc map[string]*Global_Define.StPublicizeInfo
	PaiHangBInfotmpxc = make(map[string]*Global_Define.StPublicizeInfo)
	json.Unmarshal([]byte(StXianChangInfotmp11[StrMD5].FirmsPublicize), &PaiHangBInfotmpxc)

	// 发给用户
	GameDataInfo := Proto2_Data.HS2C_FirmsUser_SendGameInfo{
		Protocol:           Proto_Data.Firms_Data_Proto,                       // 主协议
		Protocol2:          Proto2_Data.HS2C_FirmsUser_SendGameInfo_Proto,     // 子协议
		GXCAwardInfo:       PaiHangBInfotmp,                                   // 商品列表
		GXCPublicizeInfo:   PaiHangBInfotmpxc,                                 // 宣传语列表
		GFirmsNameListInfo: StXianChangInfotmp[StrMD5].FirmsName,              // 商家名字列表数据
		UseLogoPath:        RES_Path + StXianChangInfotmp[StrMD5].FirmsLogo,   // 正在使用的LOGO地址
		EWMUrlPath:         RES_Path + StXianChangInfotmp[StrMD5].FirmsErCode, // 正在使用的二维码的地址
	}

	this.PlayerSendMessage(GameDataInfo)
	return
}
