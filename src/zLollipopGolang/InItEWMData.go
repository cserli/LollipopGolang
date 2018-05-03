package main

import (
	"barcode-master"
	"barcode-master/qr"
	"glog-master"
	"image"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	//	"zLollipopGolang/db_mysql"
	"zLollipopGolang/db_redis"
	"zLollipopGolang/globalData"
	"zLollipopGolang/loglogic"
)

// 创建账号，自动刷新二维码：
//http://2/Default.ashx?url=xxxxx

// 1 http 获取 二维码数据并保存到IO磁盘
func GetEWM_Res_By_Httpbak(StrLoginNmae string, StrGameName string, StrXCName string, strLGXMD5 string) bool {
	glog.Info("Entry GetEWM_Res_By_Http!!!")
	// 获取内存的数据
	bret := Redis_DB.Redis_Read_EWMResData(strLGXMD5)
	if bret == true {
		glog.Info("Redis_Read_EWMResData is Updated!!!")
		return false
	}
	// 判断参数为空不
	if len(StrLoginNmae) == 0 || len(StrGameName) == 0 || len(StrXCName) == 0 || len(strLGXMD5) == 0 {
		glog.Info("Entry GetEWM_Res_By_Http can shu is nil!!!")
		return false
	}
	// 读取数据库获取手机的url数据
	var StXianChangInfotmp map[string]*Global_Define.StXianChangInfo
	StXianChangInfotmp, _ = dbif.GetPayGamePublicizeInfoss(StrLoginNmae, StrGameName, StrXCName, strLGXMD5)
	if StXianChangInfotmp[strLGXMD5] == nil {
		glog.Info(" GetEWM_Res_By_Http StXianChangInfotmp is nil!!")
		return false
	}
	// 获取手机的二维码的EWM数据
	url := StXianChangInfotmp[strLGXMD5].ResPath
	if len(url) == 0 {
		glog.Info(" GetEWM_Res_By_Http url is nil!!")
		return false
	}
	glog.Info("url:" + url)
	// 1 获取gameid
	iretgameid := GetGameIdByGameName(StrGameName)
	var IsTest = 0
	if Log_Eio.SysCanShu == 4 {
		// 正式环境
		IsTest = 1
	}
	glog.Info(IsTest)
	//-----------------------------------------------------------------------------------------------------------

	base64 := "http://.cn?s=" + strconv.Itoa(int(StXianChangInfotmp[strLGXMD5].ID)) + "&g=" + strconv.Itoa(iretgameid) + "&u=" + StrGameName
	glog.Info("Original data:", base64)
	code, err := qr.Encode(base64, qr.L, qr.Unicode)
	// code, err := code39.Encode(base64)
	if err != nil {
		glog.Info(err)
		return false
	}
	glog.Info("Encoded data: ", code.Content())

	if base64 != code.Content() {
		glog.Info("data differs")
		return false
	}

	code, err = barcode.Scale(code, 150, 150)
	if err != nil {
		glog.Info(err)
		return false
	}

	//-----------------------------------------------------------------------------------------------------------
	// 测试的消息
	// StrPath = gameid+"/"+EWMPic
	StrPath := strconv.Itoa(iretgameid) + "/" + "EWMPic"
	// 保存磁盘的操作
	data := StrPath + "/" + strLGXMD5 + ".png"
	writePng(data, code)
	//SaveFiles(StrPath, string(code), "png", strLGXMD5)
	// 更新二维码数据库的数据
	dbif.UpdatePayGameEWMInfo(StrLoginNmae, StrGameName, StrXCName, data)
	//  更新内存数据库
	Redis_DB.Redis_Write_EWMResData(strLGXMD5)
	return true
}

// 写入操作
func writePng(filename string, img image.Image) {
	file, err := os.Create(filename)
	if err != nil {
		glog.Info(err)
		return
	}
	err = png.Encode(file, img)
	// err = jpeg.Encode(file, img, &jpeg.Options{100})      //图像质量值为100，是最好的图像显示
	if err != nil {
		glog.Info(err)
		return
	}
	file.Close()
	glog.Info(file.Name())
	return
}

func GetEWM_Res_By_Http(StrLoginNmae string, StrGameName string, StrXCName string, strLGXMD5 string) bool {
	glog.Info("Entry GetEWM_Res_By_Http!!!")
	// 获取内存的数据
	bret := Redis_DB.Redis_Read_EWMResData(strLGXMD5)
	if bret == true {
		glog.Info("Redis_Read_EWMResData is Updated!!!")
		return false
	}
	// 判断参数为空不
	if len(StrLoginNmae) == 0 || len(StrGameName) == 0 || len(StrXCName) == 0 || len(strLGXMD5) == 0 {
		glog.Info("Entry GetEWM_Res_By_Http can shu is nil!!!")
		return false
	}
	// 读取数据库获取手机的url数据
	var StXianChangInfotmp map[string]*Global_Define.StXianChangInfo
	StXianChangInfotmp, _ = dbif.GetPayGamePublicizeInfoss(StrLoginNmae, StrGameName, StrXCName, strLGXMD5)
	if StXianChangInfotmp[strLGXMD5] == nil {
		glog.Info(" GetEWM_Res_By_Http StXianChangInfotmp is nil!!")
		return false
	}
	// 获取手机的二维码的EWM数据
	url := StXianChangInfotmp[strLGXMD5].ResPath
	if len(url) == 0 {
		glog.Info(" GetEWM_Res_By_Http url is nil!!")
		return false
	}
	glog.Info("url:" + url)
	// 1 获取gameid
	iretgameid := GetGameIdByGameName(StrGameName)
	var IsTest = 0
	if Log_Eio.SysCanShu == 4 {
		// 正式环境
		IsTest = 1
	}
	// 解析数据
	resp1, err1 := http.Get("http://.c/Default.ashx?i=" + strconv.Itoa(int(IsTest)) + "&s=" + strconv.Itoa(int(StXianChangInfotmp[strLGXMD5].ID)) + "&g=" + strconv.Itoa(iretgameid))
	if err1 != nil {
		glog.Info("playerdata", err1.Error())
		return false
	}
	// 解析数据
	body2, err2 := ioutil.ReadAll(resp1.Body)
	if err2 != nil {
		glog.Info("playerdata", err2.Error())
		return false
	}
	// 测试的消息
	glog.Info("base64:", string(body2))
	// 保存磁盘IO (StrPath string, StrBase64Data string, StrPicType string, StrPicName string)
	// StrPath = gameid+"/"+EWMPic
	StrPath := strconv.Itoa(iretgameid) + "/" + "EWMPic"
	// 保存磁盘的操作
	SaveFiles(StrPath, string(body2), "png", strLGXMD5)
	// 更新二维码数据库的数据
	data := StrPath + "/" + strLGXMD5 + ".png"
	dbif.UpdatePayGameEWMInfo(StrLoginNmae, StrGameName, StrXCName, data)
	//  更新内存数据库
	Redis_DB.Redis_Write_EWMResData(strLGXMD5)
	return true
}
