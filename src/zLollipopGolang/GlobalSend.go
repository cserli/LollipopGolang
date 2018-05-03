package main

import (
	"encoding/base64"
	"encoding/json"
	"glog-master"
	"image"
	"image/png"
	"os"
	"strings"

	"code.google.com/p/go.net/websocket"
)

//发送给客户端的数据的函数
func (this *OnlineUser) PlayerSendMessage(senddata interface{}) int {

	// 转化为JSON 这个最好封装下数据
	b, err1 := json.Marshal(senddata)
	if err1 != nil {
		glog.Error("PlayerSendMessage json.Marshal data fail ! err:", err1.Error())
		glog.Flush()
		return 1
	}
	data := ""
	data = "data" + "=" + string(b[0:len(b)])
	//	打印数据
	glog.Info("json.Marshal(data) :", data)
	glog.Flush()
	// 发送数据给玩家 立马发送数据给客户端
	err := websocket.JSON.Send(this.Connection, b)
	if err != nil {
		glog.Error("PlayerSendMessage send data fail ! err:", err.Error())
		glog.Flush()
		//this.Connection.Close()
		return 2
	}
	return 0
}

// 字符串分割函数
func Strings_Split(Data string, Split string) []string {
	return strings.Split(Data, Split)
}

// 保存磁盘的数据的图片处理函数
func SaveFiles(StrPath string, StrBase64Data string, StrPicType string, StrPicName string) bool {
	//glog.Info("SaveFiles  path:" + StrPath)
	// 转换下
	StrBase64Data = strings.Replace(StrBase64Data, "\"", "", -1)
	// 解析
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(StrBase64Data))
	// 转换成png格式的图像，需要导入：_“image/png”
	m, _, ermmr := image.Decode(reader)
	if ermmr != nil {
		glog.Error("image.Decode(reader):", ermmr.Error())
	}
	// 输出到磁盘里:包括路径
	// 文件夹操作
	//	dir, _ := os.Getwd() // 获取当前的程序路径
	//	StrPath = dir + StrPath
	StrPath = "/var/www/html/res/" + StrPath
	err := os.MkdirAll(StrPath, os.ModePerm) //生成多级目录
	if err != nil {
		glog.Error(err.Error())
		glog.Flush()
		return false
	}
	// 保存数据
	StrPicType = StrPath + "/" + StrPicName + "." + StrPicType
	wt, err := os.Create(StrPicType)
	if err != nil {
		glog.Error("Save Image Error!" + err.Error())
		glog.Flush()
		return false
	}
	// defer wt.Close()
	if wt == nil {
		glog.Error("Save Image Error!  wt is nil!!!")
		glog.Flush()
		return false
	}
	// 转换为jpeg格式的图像，这里质量为30（质量取值是1-100）
	// jpeg.Encode(wt, m, &jpeg.Options{30})
	png.Encode(wt, m)
	wt.Close()
	return true
}

// 获取商家的OpenID信息
func Get_MD5_Data(stropenid string) string {
	// 判断数据库存在不
	var StrMD5 = ""
	// 差分key
	strsplit := Strings_Split(stropenid, "|")
	for i := 0; i < len(strsplit); i++ {
		if i == 0 {
		} else if i == 1 {
			StrMD5 = strsplit[i]
		}
	}
	return StrMD5
}

// 获取商家的OpenID信息
func Get_OpenID_Data(stropenid string) string {
	// 判断数据库存在不
	var StrMD5 = ""
	// 差分key
	strsplit := Strings_Split(stropenid, "|")
	for i := 0; i < len(strsplit); i++ {
		if i == 0 {
			StrMD5 = strsplit[i]
		} else if i == 1 {
		}
	}
	return StrMD5
}
