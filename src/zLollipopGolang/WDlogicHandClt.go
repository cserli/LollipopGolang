package main

import (
	"glog-master"
	"strconv"
	"zLollipopGolang/db_mysql"
	"zLollipopGolang/protocolfile"
	_ "zLollipopGolang/protocolfile/Proto2"
	"zLollipopGolang/protocolfile/proto3"
)

// 消息处理函数,通过协议去处理数据； 网络相关的子协议处理； 主要是登陆和注册等数据。
func (this *OnlineUser) HandleCltProtocol2GM(protocol2 interface{}, ProtocolData map[string]interface{}) {

	// 通过过去的协议惊醒解析操作
	switch protocol2 {
	case float64(Proto3_Data.GM2S_Add_Apply_Code_Proto):
		{ // 设置邀请码--包括邀请激活码--金牌教师邀请码
			this.Gm_AddApply_Code_Fun(ProtocolData)
			break
		}
	default:
		panic("主协议：3，子协议 不存在！！！")
	}
	return
}

// 生产激活码==也就是系统的邀请码
func (this *OnlineUser) Gm_AddApply_Code_Fun(ProtocolData map[string]interface{}) {
	strOPUID := ProtocolData["OPUID"].(string) // 问答数据
	glog.Info("增加 Add  Apply Code", strOPUID)
	// 数据库生产--激活码
	strcode := dbif.RegNewApplyCodeWD()
	// 数据结构--发送---给网页
	Pc_PlayeNInfo := Proto3_Data.S2GM_Add_Apply_Code{
		Protocol:   Proto_Data.G_Gm_Data_Proto,            // 主协议
		Protocol2:  Proto3_Data.S2GM_Add_Apply_Code_Proto, // 子协议
		Apply_Code: strcode,                               // 邀请码
	}
	this.PlayerSendMessage(Pc_PlayeNInfo)
	return
}

// 消息处理函数,通过协议去处理数据； 网络相关的子协议处理； 主要是登陆和注册等数据。
func (this *OnlineUser) HandleCltProtocol2WD(protocol2 interface{}, ProtocolData map[string]interface{}) {

	// 通过过去的协议惊醒解析操作
	switch protocol2 {
	case float64(Proto3_Data.Net_Heart_Beating_Proto):
		{ // 心跳检测的子协议
			this.NetHeartBeatingWD(ProtocolData)
			break
		}
	case float64(Proto3_Data.C2S_WinXinPlayer_Chat_WT_Proto):
		{ // 聊天
			this.Player_Chat_Fun(ProtocolData)
			break
		}
	case float64(Proto3_Data.C2S_WinXinPlayer_GetList_T_Proto):
		{ // 获取老师的列表信息
			this.Player_GetWDCommentListOFBeatT_Fun(ProtocolData)
			break
		}
	case float64(Proto3_Data.C2S_WinXinPlayer_CommentList_Of_Best_WT_Proto):
		{ // 选择最优的评论
			this.Player_GetWDCommentListOFBeat_Fun(ProtocolData)
			break
		}
	case float64(Proto3_Data.C2S_WinXinPlayer_CommentList_WT_Proto):
		{ // 获取问答的评论列表
			this.Player_GetWDCommentList_Fun(ProtocolData)
			break
		}
	case float64(Proto3_Data.C2S_WinXinPlayer_My_WT_Proto):
		{ // 获取我的问答列表
			this.Player_GetWDListOfMy_Fun(ProtocolData)
			break
		}
	case float64(Proto3_Data.C2S_WinXinPlayer_Apply_Teacher_Proto):
		{ // 申请当老师--需要邀请码
			this.Player_Apply_Teach_Fun(ProtocolData)
			break
		}
	case float64(Proto3_Data.C2S_WinXinPlayer_Comment_WT_Proto):
		{ // 发表评论
			this.Player_YaoYiYao_FaBiaoPingLun_Fun(ProtocolData)
			break
		}
	case float64(Proto3_Data.C2S_WinXinPlayer_GetList_WT_Proto):
		{ // 获取提问列表
			this.Player_YaoYiYao_WenDaList_Fun(ProtocolData)
			break
		}
	case float64(Proto3_Data.C2S_WinXinPlayer_TiWen_Proto):
		{ // 提问的功能--数据保存
			this.Player_YaoYiYao_WenDa_Fun(ProtocolData)
			break
		}
	case float64(Proto3_Data.C2S_Net_WinXin_OpenID_BaoMing_Proto):
		{ // 微信处理协议
			this.Player_YaoYiYao_BaoMing_FunWd(ProtocolData)
			break
		}

	default:
		panic("主协议：3，子协议 不存在！！！")
	}
	return
}

// 心跳检测函数实现
func (this *OnlineUser) NetHeartBeatingWD(ProtocolData map[string]interface{}) {
	if ProtocolData["OpenID"] == nil {
		panic("主协议：3，子协议 1；协议说明：心跳检测函数实现")
	}
	// 解析
	StrOpenID := ProtocolData["OpenID"].(string)
	if len(StrOpenID) == 0 {
		return
	}

	onlineUser := &OnlineUser{
		Connection: this.Connection, // 链接的数据信息
		MapSafe:    this.MapSafe,
	}
	// 测试阶段
	this.MapSafe.Put(StrOpenID+"|connect", onlineUser)

	return
}

// 聊天
func (this *OnlineUser) Player_Chat_Fun(ProtocolData map[string]interface{}) {

	strTOpenID := ProtocolData["TOpenID"].(string)   // 老师的ID
	strMyOpenID := ProtocolData["MyOpenID"].(string) // 问答者的ID
	strData := ProtocolData["Data"].(string)         // 聊天数据

	dataplayer, _ := dbif.GetUserInfoWDBaoMing(strMyOpenID)
	// 回消息
	Pc_PlayeNInfo := Proto3_Data.S2C_WinXinPlayer_Chat_WT{
		Protocol:  Proto_Data.Network_Data_Proto,              // 主协议
		Protocol2: Proto3_Data.S2C_WinXinPlayer_Chat_WT_Proto, // 子协议
		TOpenID:   strMyOpenID,
		WOpenID:   strTOpenID,
		WHeadUrl:  dataplayer["1"].HeadUrl,
		WName:     dataplayer["1"].Name,
		Data:      strData, //
	}
	// 发送数据
	val, _ := this.MapSafe.Get(strTOpenID + "|connect")
	if val == nil {
		glog.Info("老师已经离线")
		Pc_PlayeNInfo.Data = "老师已经离线"
		this.PlayerSendMessage(Pc_PlayeNInfo)
		return
	} else {
		if val.(interface{}).(*OnlineUser).PlayerSendMessage("test") == 2 {
			glog.Info("老师已经离线")
			Pc_PlayeNInfo.Data = "老师已经离线"
			this.PlayerSendMessage(Pc_PlayeNInfo)
			return
		} else {
			val.(interface{}).(*OnlineUser).PlayerSendMessage(Pc_PlayeNInfo)
			return
		}

	}

	return
}

// 获取老师列表
func (this *OnlineUser) Player_GetWDCommentListOFBeatT_Fun(ProtocolData map[string]interface{}) {

	data, _ := dbif.GetUserInfoWDBaoMingOfOnline()
	// 回消息
	Pc_PlayeNInfo := Proto3_Data.S2C_WinXinPlayer_GetList_T{
		Protocol:  Proto_Data.Network_Data_Proto,                // 主协议
		Protocol2: Proto3_Data.S2C_WinXinPlayer_GetList_T_Proto, // 子协议
		Data:      data,                                         //
	}
	this.PlayerSendMessage(Pc_PlayeNInfo)

	return
}

// 获取评列表
func (this *OnlineUser) Player_GetWDCommentListOFBeat_Fun(ProtocolData map[string]interface{}) {
	strWenDaID := ProtocolData["WenDaID"].(string)     // 个人中心的openid
	strCommentID := ProtocolData["CommentID"].(string) // 评论id

	// 数据库操作
	strCommentIDi, _ := strconv.Atoi(strCommentID)
	// 更新到问答列表
	dbif.UpdateWDData(strWenDaID, strCommentIDi)
	// 更新到评论列表
	dbif.UpdateWDPLData(strCommentID, 1)
	// 回消息
	Pc_PlayeNInfo := Proto3_Data.S2C_WinXinPlayer_CommentList_Of_Best_WT{
		Protocol:  Proto_Data.Network_Data_Proto,                             // 主协议
		Protocol2: Proto3_Data.S2C_WinXinPlayer_CommentList_Of_Best_WT_Proto, // 子协议
		IS_Succ:   true,                                                      //
	}
	this.PlayerSendMessage(Pc_PlayeNInfo)

	return
}

// 获取评列表
func (this *OnlineUser) Player_GetWDCommentList_Fun(ProtocolData map[string]interface{}) {

	// strOpenID := ProtocolData["OpenID"].(string)   // 个人中心的openid
	strWenDaID := ProtocolData["WenDaID"].(string) // 个人中心的openid
	strPageNum := ProtocolData["PageNum"].(string) // 分页获取

	// 数据库
	data, _ := dbif.GetWDCommentListOfMy(strPageNum, strWenDaID)
	// 回消息
	Pc_PlayeNInfo := Proto3_Data.S2C_WinXinPlayer_CommentList_WT{
		Protocol:  Proto_Data.Network_Data_Proto,                     // 主协议
		Protocol2: Proto3_Data.S2C_WinXinPlayer_CommentList_WT_Proto, // 子协议
		Data:      data,                                              //
	}
	this.PlayerSendMessage(Pc_PlayeNInfo)

	return
}

// 个人中心
func (this *OnlineUser) Player_GetWDListOfMy_Fun(ProtocolData map[string]interface{}) {

	strOpenID := ProtocolData["OpenID"].(string)   // 个人中心的openid
	strPageNum := ProtocolData["PageNum"].(string) // 分页获取

	data, _ := dbif.GetWDListOfMy(strPageNum, strOpenID)

	Pc_PlayeNInfo := Proto3_Data.S2C_WinXinPlayer_My_WT{
		Protocol:  Proto_Data.Network_Data_Proto,            // 主协议
		Protocol2: Proto3_Data.S2C_WinXinPlayer_My_WT_Proto, // 子协议
		Data:      data,                                     //
	}
	this.PlayerSendMessage(Pc_PlayeNInfo)

	return
}

// 申请为老师
func (this *OnlineUser) Player_Apply_Teach_Fun(ProtocolData map[string]interface{}) {

	strOpenID := ProtocolData["OpenID"].(string)         // 提出问题的openid
	strApply_Code := ProtocolData["Apply_Code"].(string) // 申请码

	// 更新数据
	retstr := dbif.UpdateApllyData(strOpenID, strApply_Code)

	Pc_PlayeNInfo := Proto3_Data.S2C_WinXinPlayer_Apply_Teacher{
		Protocol:  Proto_Data.Network_Data_Proto,                    // 主协议
		Protocol2: Proto3_Data.S2C_WinXinPlayer_Apply_Teacher_Proto, // 子协议
		Code:      retstr,                                           //
	}
	this.PlayerSendMessage(Pc_PlayeNInfo)
	return
}

// 发表评论
func (this *OnlineUser) Player_YaoYiYao_FaBiaoPingLun_Fun(ProtocolData map[string]interface{}) {

	strWenTiID := ProtocolData["WenTiID"].(string)   // 问题的id
	strOpenID := ProtocolData["OpenID"].(string)     // 提出问题的openid
	strPLOpenID := ProtocolData["PLOpenID"].(string) // 评论人的openid
	strData := ProtocolData["Data"].(string)         // 评论的数据
	strType := ProtocolData["Type"].(string)         // 提问者选择一个最优的
	// 保存数据到评论列表的数据中
	dbif.RegNewWDPingLunData(strWenTiID, strOpenID, strPLOpenID, strData, strType)
	Pc_PlayeNInfo := Proto3_Data.S2C_WinXinPlayer_Comment_WT{
		Protocol:  Proto_Data.Network_Data_Proto,                 // 主协议
		Protocol2: Proto3_Data.S2C_WinXinPlayer_Comment_WT_Proto, // 子协议
		IS_Succ:   true,                                          //
	}
	this.PlayerSendMessage(Pc_PlayeNInfo)
	return
}

// 获取提问列表
func (this *OnlineUser) Player_YaoYiYao_WenDaList_Fun(ProtocolData map[string]interface{}) {
	// 解析数据
	// strOpenID := ProtocolData["OpenID"].(string)   // 微信用户的唯一的ID
	strPageNum := ProtocolData["PageNum"].(string) // 问答数据
	// 获取数据库数据
	data, _ := dbif.GetWDList(strPageNum)

	Pc_PlayeNInfo := Proto3_Data.S2C_WinXinPlayer_GetList_WT{
		Protocol:  Proto_Data.Network_Data_Proto,                 // 主协议
		Protocol2: Proto3_Data.S2C_WinXinPlayer_GetList_WT_Proto, // 子协议
		Data:      data,                                          //
	}
	this.PlayerSendMessage(Pc_PlayeNInfo)
	return
}

func (this *OnlineUser) Player_YaoYiYao_WenDa_Fun(ProtocolData map[string]interface{}) {

	// 解析数据
	strOpenID := ProtocolData["OpenID"].(string)   // 微信用户的唯一的ID
	strData := ProtocolData["Data"].(string)       // 问答数据
	strPicData := ProtocolData["PicData"].(string) // 图片数据
	strCoin := ProtocolData["Coin"].(string)       // 付费问答的费用
	// 服务的费用不可以为空--
	if strCoin == "" {
		panic("提问问题金币不可以为0")
		return
	}
	// 数据的保存
	dbif.RegNewUserForWeiXinTiWD(strOpenID, strData, strPicData, strCoin)
	// 数据发送
	Pc_PlayeNInfo := Proto3_Data.S2C_WinXinPlayer_TiWen{
		Protocol:  Proto_Data.Network_Data_Proto,            // 主协议
		Protocol2: Proto3_Data.S2C_WinXinPlayer_TiWen_Proto, // 子协议
		IS_Succ:   true,                                     //
	}
	this.PlayerSendMessage(Pc_PlayeNInfo)
	return
}

// 游戏操作的步骤:报名
func (this *OnlineUser) Player_YaoYiYao_BaoMing_FunWd(ProtocolData map[string]interface{}) {

	if ProtocolData["Code"] == nil {
		panic("主协议：2，子协议：1,；用户请求报名")
		return
	}
	// 解析数据
	strCode := ProtocolData["Code"].(string) // 微信用户的唯一的ID
	glog.Info("Player_YaoYiYao_BaoMing_Fun:", strCode)
	// 获取微信数据
	_, strtmpopenid, _, _ := httpGetWD(strCode)
	//	if iiret == 0 { // 微信授权失败
	//		glog.Info("Player_YaoYiYao_BaoMing_Fun:", strnikename, strheadurl)

	//		// 发数据给用户
	//		YaoYiYaoBaoMingInfo := Proto2_Data.G_Error_All{
	//			Protocol:  Proto_Data.G_Error_Proto,      // 主协议
	//			Protocol2: Proto2_Data.G_Error_All_Proto, // 子协议
	//			ErrCode:   "40001",                       // 报名失败
	//		}
	//		// 发送给玩家数据
	//		this.PlayerSendMessage(YaoYiYaoBaoMingInfo)
	//		return
	//	} else {
	// 处理内测试邀请码
	if dbif.GetNeiCeZiGeData(strtmpopenid) == 0 {
		// 返回数据--无内测资格
		//		YaoYiYaoBaoMingInfo := Proto2_Data.G_Error_All{
		//			Protocol:  Proto_Data.G_Error_Proto,      // 主协议
		//			Protocol2: Proto2_Data.G_Error_All_Proto, // 子协议
		//			ErrCode:   "40004",                       // 报名失败
		//		}
		//		// 发送给玩家数据
		//		this.PlayerSendMessage(YaoYiYaoBaoMingInfo)
		//return
	}
	// 获取用户的基本的信息
	Data, _ := dbif.GetUserInfoWDBaoMing(strtmpopenid)
	// 推送PC 报名的信息
	Pc_PlayeNInfo := Proto3_Data.S2C_Net_WinXin_OpenID_BaoMing{
		Protocol:  Proto_Data.Network_Data_Proto,                   // 主协议
		Protocol2: Proto3_Data.S2C_Net_WinXin_OpenID_BaoMing_Proto, // 子协议
		Data:      Data,                                            //
	}

	// 发广播所有的用户--如果是老师上线的话进行广播
	// this.XC_Data_Send_AllPlayer_State(strMD5data, Pc_PlayeNInfo)
	// 发送数据
	this.PlayerSendMessage(Pc_PlayeNInfo)
	//================================推送消息处理===================================
	// 保存在线的玩家的数据信息
	onlineUser222 := &OnlineUser{
		Connection: this.Connection, // 链接的数据信息
		MapSafe:    this.MapSafe,
	}
	// 测试阶段
	this.MapSafe.Put(strtmpopenid+"|connect", onlineUser222)
	//==============================================================================
	return
	//	}
	//	return
}
