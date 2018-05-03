package main

import (
	"glog-master"
	"zLollipopGolang/globalData"
	"zLollipopGolang/protocolfile"
	"zLollipopGolang/protocolfile/Proto2"
)

// 处理服务器间的数据的处理
func (this *OnlineServer) HandleCltProtocol2ServerData(protocol2 interface{}, ProtocolData map[string]interface{}) {

	// 通过过去的协议惊醒解析操作
	glog.Info("protocol2:", protocol2)
	switch protocol2 {
	case float64(Proto2_Data.Net_Link_Data_Proto):
		{ // 绑定子游戏服务器的协议
			this.NetHeartBeatingServer(ProtocolData)
			break
		}
	case float64(Proto2_Data.HSS2C_FirmsUser_SendGameInfo_Proto):
		{ // 提示PC更新
			this.Update_Data_Link_To_ZiGame(ProtocolData)
			break
		}
	case float64(Proto2_Data.H52S_CSVFile_RST_Proto):
		{ // 重新加载CSV文件
			this.Update_Data_CSVFile_RST()
			break
		}
	default:
		panic("主协议：8，子协议 不存在！！！")
	}
	return
}

// 重新加载CSV文件
func (this *OnlineServer) Update_Data_CSVFile_RST() {
	glog.Info("重新加载配置文件，CSV!")
	ReadCsv_ConfigFile()
	return
}

// 获取数据，主动推送给服务器，实时推送给游戏
// 查找游戏服务器的数据的链接
func (this *OnlineServer) Update_Data_Link_To_ZiGame(ProtocolData map[string]interface{}) {
	glog.Info("Update_Data_Link_To_ZiGame")
	//  获取登陆服务器传递的数据信息
	strLoginName := ProtocolData["StrLoginName"].(string)
	strGameName := ProtocolData["StrGameName"].(string)
	strXCName := ProtocolData["StrXCName"].(string)

	var StrAwardLev map[string]interface{}
	if ProtocolData["GXCAwardInfo"] != nil {
		StrAwardLev = ProtocolData["GXCAwardInfo"].(map[string]interface{})
	}

	var StrAwardLev1 map[string]interface{}
	if ProtocolData["GXCPublicizeInfo"] != nil {
		StrAwardLev1 = ProtocolData["GXCPublicizeInfo"].(map[string]interface{})
	}

	strname := ""
	if ProtocolData["GFirmsNameListInfo"] != nil {
		strname = ProtocolData["GFirmsNameListInfo"].(string)
	}

	strewm := ""
	if ProtocolData["EWMUrlPath"] != nil {
		if len(ProtocolData["EWMUrlPath"].(string)) != 0 {
			strewm = RES_Path + ProtocolData["EWMUrlPath"].(string)
		}
	}

	strlongo := ""
	if ProtocolData["UseLogoPath"] != nil {
		if len(ProtocolData["UseLogoPath"].(string)) != 0 {
			strlongo = RES_Path + ProtocolData["UseLogoPath"].(string)
		}
	}
	//StrAwardLev := ProtocolData["GXCAwardInfo"].(map[string]interface{})
	//StrAwardLev1 := ProtocolData["GXCPublicizeInfo"].(map[string]interface{})
	//strname := ProtocolData["GFirmsNameListInfo"].(string)
	//strlongo := ProtocolData["UseLogoPath"].(string)
	//strewm := ProtocolData["EWMUrlPath"].(string)
	// 链接数据
	strMD5data := Global_Define.MD5ToStringFromString(strLoginName + "|" + strGameName + "|" + strXCName)
	//--------------------------------------------------------------------------
	// 解析数据
	var PaiHangBInfotmp1 map[string]*Global_Define.StAwardInfo
	PaiHangBInfotmp1 = make(map[string]*Global_Define.StAwardInfo)

	for first, _ := range StrAwardLev {
		ItemInfo := StrAwardLev[first].(map[string]interface{})
		Infotmp := new(Global_Define.StAwardInfo)

		Infotmp.Name = string(ItemInfo["Name"].(string))
		Infotmp.Lev = string(ItemInfo["Lev"].(string))
		Infotmp.StrMsg = string(ItemInfo["StrMsg"].(string))
		Infotmp.AwardNum = uint32(ItemInfo["AwardNum"].(float64))
		Infotmp.BMark = bool(ItemInfo["BMark"].(bool))

		PaiHangBInfotmp1[string(ItemInfo["Name"].(string))] = Infotmp
	}

	// 获取某一个
	var PaiHangBInfotmp11 map[string]*Global_Define.StPublicizeInfo
	PaiHangBInfotmp11 = make(map[string]*Global_Define.StPublicizeInfo)

	for first, _ := range StrAwardLev1 {
		ItemInfo := StrAwardLev1[first].(map[string]interface{})
		Infotmp := new(Global_Define.StPublicizeInfo)

		Infotmp.Name = string(ItemInfo["Name"].(string))
		Infotmp.Lev = string(ItemInfo["Lev"].(string))
		Infotmp.StrMsg = string(ItemInfo["StrMsg"].(string))
		Infotmp.BMark = bool(ItemInfo["BMark"].(bool))

		PaiHangBInfotmp11[string(ItemInfo["Name"].(string))] = Infotmp
	}
	//--------------------------------------------------------------------------
	// 发给商家的数据更新的
	GameDataInfo := Proto2_Data.HS2C_FirmsUser_SendGameInfo{
		Protocol:           Proto_Data.Firms_Data_Proto,                   // 主协议
		Protocol2:          Proto2_Data.HS2C_FirmsUser_SendGameInfo_Proto, // 子协议
		GXCAwardInfo:       PaiHangBInfotmp1,                              // 商品列表
		GXCPublicizeInfo:   PaiHangBInfotmp11,                             // 宣传语列表
		GFirmsNameListInfo: strname,                                       // 商家名字列表数据
		UseLogoPath:        strlongo,                                      // 正在使用的LOGO地址
		EWMUrlPath:         strewm,                                        // 正在使用的二维码的地址
	}

	glog.Info("GameDataInfo:", GameDataInfo)

	// 发送数据
	if runningGoldMinerRoom.OnlineUsers[strMD5data] != nil {
		runningGoldMinerRoom.OnlineUsers[strMD5data].PlayerSendMessage(GameDataInfo)
		return
	}
	glog.Info("runningGoldMinerRoom:is nil!")
	return
}

func (this *OnlineServer) NetHeartBeatingServer(ProtocolData map[string]interface{}) {

	if ProtocolData["OpenID"] == nil {
		panic("NetHeartBeatingServer is nil!!!")
	}
	// 保存的登陆服务器的链接，以后发送到登陆服务器
	StrOpenID := ProtocolData["OpenID"].(string)
	// 需要传输的数据的结构
	heartbeat := Proto2_Data.Net_Link_Data{
		Protocol:  Proto_Data.G_ServerToServer_Proto, // 主协议
		Protocol2: Proto2_Data.Net_Link_Data_Proto,   // 子协议
		OpenID:    StrOpenID,                         // 数据链接的验证码
	}
	//================================推送消息处理===================================
	// 保存在线的玩家的数据信息
	onlineServer := runningServerRoom.OnlineServers[StrOpenID]
	onlineServer = &OnlineServer{
		Connection: this.Connection, // 链接的数据信息
		StrMD5:     StrOpenID,       // 绑定商家信息,必须map
	}
	// 赋值操作数据
	runningServerRoom.OnlineServers[StrOpenID] = onlineServer
	//==============================================================================
	// 发送给玩家数据
	this.SendHallServerDataFun(heartbeat)
	return
}
