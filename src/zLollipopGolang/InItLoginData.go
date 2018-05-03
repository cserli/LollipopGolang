package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"glog-master"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
	"zLollipopGolang/globalData"
	"zLollipopGolang/protocolfile"
	//	"zLollipopGolang/protocolfile/Proto2"

	"code.google.com/p/go.net/websocket"
)

//------------------------------------------------------------------------------
// 子游戏服务器链接大厅服务器函数封装
var origin = ""

// 链接地址
//var url = "ws://:01/ff?"
var urlre = ""

// 子游戏作为客户端链接登陆服务器
func ZiGameConnetHallServer() bool {

	glog.Info("-------------ZiGameConnetHallServer---------------")
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			strerr := fmt.Sprintf("%s", err)
			//发消息给客户端
			ErrorST := Proto2_Data.G_Error_All{
				Protocol:  Proto_Data.G_Error_Proto,      // 主协议
				Protocol2: Proto2_Data.G_Error_All_Proto, // 子协议
				ErrCode:   "80006",
				ErrMsg:    "server 7878 亲，您发的数据的格式不对！" + strerr,
			}
			// 发送给玩家数据
			//this.GlobalSendHallServerDataFun(ErrorST)
			glog.Info("ErrorST", ErrorST)
		}
	}()
	// md5的数据，
	strOpenid := Global_Define.MD5ToStringFromString("DuiZhan")
	// 心跳包的协议
	test := Proto2_Data.Net_Heart_Beating{
		// 数据赋值
		Protocol:  Proto_Data.Network_Data_Proto,       // 主协议
		Protocol2: Proto2_Data.Net_Heart_Beating_Proto, // 子协议
		OpenID:    strOpenid,
	}

	b, _ := json.Marshal(test)
	data := ""
	data = "data" + "=" + string(b[0:len(b)])
	glog.Info(urlre + data)
	ws, err := websocket.Dial(urlre+data, "", origin)
	if err != nil {
		glog.Info("ZiGameConnetHallServer falied, err:", err.Error())
		return false
	}
	// 初始化在线玩家数据
	serverdata := &OnlineServer{
		Connection: ws, // 链接的数据信息
	}
	// 心跳包的协议
	testCon := Proto2_Data.Net_Link_Data{
		// 数据赋值
		Protocol:  Proto_Data.G_ServerToServer_Proto, // 主协议
		Protocol2: Proto2_Data.Net_Link_Data_Proto,   // 子协议
		OpenID:    strOpenid,
	}
	// 发送给登陆服务器，握手协议
	serverdata.SendHallServerDataFun(testCon)
	// 接受数据
	serverdata.ReceiveHallServerDataFun()
	// WS.Close() //关闭连接
	return true
}

// 子游戏发送给登陆服务器的数据的请求的函数
func (this *OnlineServer) SendHallServerDataFun(senddata interface{}) {

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			strerr := fmt.Sprintf("%s", err)
			glog.Info("清除奖券定时器：", strerr)
		}
	}()
	// 发送消息处理
	glog.Info("SendHallServerDataFun:", senddata)
	b, _ := json.Marshal(senddata)
	data := ""
	//data = "data" + "=" + string(b[0:len(b)])
	data = string(b[0:len(b)])
	err := websocket.Message.Send(this.Connection, data)
	if err != nil {
		glog.Info("SendHallServerDataFun err:", err.Error())
		return
	}
	return
}

// 子游戏发送给登陆服务器的数据的请求的函数
func (this *OnlineServer) GlobalSendHallServerDataFun(senddata interface{}) {

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			strerr := fmt.Sprintf("%s", err)
			glog.Info("清除奖券定时器：", strerr)
		}
	}()
	// 发送消息处理
	glog.Info("SendHallServerDataFun:", senddata)
	senddata, _ = json.Marshal(senddata)
	err := websocket.Message.Send(this.Connection, senddata)
	if err != nil {
		glog.Info("SendHallServerDataFun err:", err.Error())
		return
	}
	return
}

// 接受登陆服务器返回的数据
func (this *OnlineServer) ReceiveHallServerDataFun() {
	glog.Info("Enrty ReceiveHallServerDataFun")
	// 获取消息处理
	var content string
	for {
		err := websocket.Message.Receive(this.Connection, &content)
		if err != nil {
			glog.Info(err.Error())
			//  游戏服务器  强制退出；可以不要求游戏服务器退出
			if err.Error() == "EOF" {
				glog.Info("Server is dead ...ByeBye")
				os.Exit(0)
				return
			}
			continue
		}
		//glog.Info("ReceiveHallServerDataFun content:", content)
		// base64转换
		sTmp := strings.Trim(content, "\"")
		strJSONtmp := Base64String(sTmp)
		// 多线程处理
		go this.SyncServerMessageDataFun(strJSONtmp)
	}
	return
}

// 异步处理服务器处理的函数
func (this *OnlineServer) SyncServerMessageDataFun(content string) {

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			strerr := fmt.Sprintf("%s", err)
			//发消息给客户端
			ErrorST := Proto2_Data.G_Error_All{
				Protocol:  Proto_Data.G_Error_Proto,      // 主协议
				Protocol2: Proto2_Data.G_Error_All_Proto, // 子协议
				ErrCode:   "80006",
				ErrMsg:    "亲，您发的数据的格式不对！" + strerr,
			}
			// 发送给玩家数据
			//this.GlobalSendHallServerDataFun(ErrorST)
			glog.Info("ErrorST", ErrorST)
		}
	}()

	var r Requestbody
	r.req = content
	if ProtocolData, err := r.Json2map(); err == nil {

		// md5 加密
		h := md5.New()
		h.Write([]byte(strconv.Itoa(Global_Define.RandData(1000)))) // 需要加密的字符串为 sharejs.com
		// switch 处理消息数据，主要是处理服务器之间的请求的数据
		this.HandleServerCltProtocol(ProtocolData["Protocol"], ProtocolData["Protocol2"], ProtocolData)
	}
	return
}

// 消息处理函数,通过协议去处理数据;来自于大厅服务器处理数据
func (this *OnlineServer) HandleServerCltProtocol(protocol interface{}, protocol2 interface{}, ProtocolData map[string]interface{}) {

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			strerr := fmt.Sprintf("%s", err)
			//发消息给客户端
			ErrorST := Proto2_Data.G_Error_All{
				Protocol:  Proto_Data.G_Error_Proto,      // 主协议
				Protocol2: Proto2_Data.G_Error_All_Proto, // 子协议
				ErrCode:   "80006",
				ErrMsg:    "亲，您发的数据的格式不对！" + strerr,
			}
			// 发送给玩家数据
			//this.GlobalSendHallServerDataFun(ErrorST)
			glog.Info("ErrorST", ErrorST)
		}
	}()
	glog.Info("protocol", protocol)
	// 通过过去的协议惊醒解析操作
	switch protocol {
	case float64(Proto_Data.G_ServerToServer_Proto):
		{
			// 链接数据相关
			this.HandleCltProtocol2ServerData(protocol2, ProtocolData)
			break
		}
	default:
		panic(" HandleServerCltProtocol default!")
	}
	return
}

func B2S(buf []byte) string {
	return *(*string)(unsafe.Pointer(&buf))
}

func S2B(s *string) []byte {
	return *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(s))))
}

// base64 转换
func Base64String(data string) string {

	//glog.Info(data)
	B64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	bytes, err := B64.DecodeString(data)
	if err != nil {
		glog.Info("Base64String ：", err.Error())
		return ""
	}
	return string(bytes)
}
