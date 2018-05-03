package main

import (
	"glog-master"
	"zLollipopGolang/db_mysql"
)

////////////////////////////////////////////////////////////////////////////////
//
//
// 获取游戏的配置信息，主要是游戏的配置的信息，一般从数据库或者配置读取数据
//
// 每次进入游戏后获取数据
////////////////////////////////////////////////////////////////////////////////

func GetGameConfigFromDB(strLoginName string, strGameName string, strXCName string) (string, string, string) {
	glog.Info("Entry GetGameConfigFromDB!!! ")
	if len(strLoginName) == 0 || len(strGameName) == 0 || len(strXCName) == 0 {
		glog.Info("Entry GetGameConfigFromDB!!! 数据长度为空！")
		return "", "", ""
	}
	// 获取数据库数据信息
	var datatimelong = ""
	var datacdtime = ""
	var datacdcishu = ""
	datatimelong, datacdtime, datacdcishu = dbif.GetPayGameConfigInfo(strLoginName, strGameName, strXCName)
	glog.Info("GetGameConfigFromDB datatimelong:" + datatimelong + "datacdtime:" + datacdtime + "datacdcishu:" + datacdcishu)

	return datatimelong, datacdtime, datacdcishu
}
