package main

import (
	"encoding/json"
	"fmt"
	"log"
	"protocolfile"
	"protocolfile/proto3"

	"golang.org/x/net/websocket"
)

var origin = "http://127.0.0.1:8080/"
var url = "ws://127.0.0.1:8080/GolangLTD"

func main() {

	//for {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	Pc_PlayeNInfo := Proto3_Data.Net_Heart_Beating{
		Protocol: Proto_Data.Network_Data_Proto, // 主协议
	}

	b, err1 := json.Marshal(Pc_PlayeNInfo)
	if err1 != nil {

	}
	data := ""
	data = "data" + "=" + string(b[0:len(b)])
	fmt.Println(data)
	// 发送数据给玩家 立马发送数据给客户端
	datamap := make(map[string]interface{})
	datamap["data"] = data
	websocket.JSON.Send(ws, datamap)

	// 发送数据给服务器
	//	PlayerSendMessage(Pc_PlayeNInfo, ws)
}
