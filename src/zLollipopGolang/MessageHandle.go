package main

import (
	"glog-master"
	"os"
	//	"strconv"
	"time"
	//	"zLollipopGolang/db_mysql"
	//	"zLollipopGolang/db_redis"
	//	"zLollipopGolang/globalData"
	//	"zLollipopGolang/protocolfile"
	//	"zLollipopGolang/protocolfile/Proto2"
)

// 处理服务器间的数据的处理
func (this *OnlineUser) HandleCltProtocol2ServerDatabak(protocol2 interface{}, ProtocolData map[string]interface{}) {

	// 通过过去的协议惊醒解析操作
	glog.Info("protocol2:", protocol2)
	switch protocol2 {
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
func (this *OnlineUser) Update_Data_CSVFile_RST() {
	glog.Info("重新加载配置文件，CSV!")
	ReadCsv_ConfigFile()
	return
}

// 消息处理函数,通过协议去处理数据； 网络相关的子协议处理； 主要是登陆和注册等数据。
func (this *OnlineUser) HandleCltProtocol2(protocol2 interface{}, ProtocolData map[string]interface{}) {

	// 通过过去的协议惊醒解析操作
	switch protocol2 {
	case float64(Proto2_Data.G_Test):
		{ // 心跳检测的子协议
			this.NetHeartBeatinssg(ProtocolData)
			break
		}

	default:
		panic("主协议：3，子协议 不存在！！！")
	}
	return
}

func (this *OnlineUser) Player_Re_EntryGame(ProtocolData map[string]interface{}) {

	glog.Info("Entry Player_Re_EntryGame!!!")
	os.Exit(0)
	return
}

// 游戏操作的步骤:报名
func (this *OnlineUser) Player_Re_Entry_TR_Fun(ProtocolData map[string]interface{}) {

	if ProtocolData["OpenID"] == nil {
		panic("主协议：3，子协议：10,；用户重新游戏")
		return
	}

	// 解析state这个字段数据，获取商家的名字的数据的信息等，读取内存数据库
	strIDtmp := ProtocolData["OpenID"].(string)
	if len(strIDtmp) == 0 {
		panic("主协议：3，子协议：11,；  OpenID 字段不对！！！")
		return
	}
	iLastLoginTimes := time.Now().Unix()
	this.MapSafe.Put(strIDtmp+"|GGameConfigOfPlayer", iLastLoginTimes)
	return
}
