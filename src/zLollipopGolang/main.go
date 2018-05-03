package main

import (
	//	"cache2go"
	"encoding/json"
	"flag"
	"fmt"
	"glog-master"
	"go-concurrentMap-master" //https://github.com/fanliao/go-concurrentMap
	"log"
	"net/http"
	"os"
	"runtime"
	"zLollipopGolang/protocolfile"

	"code.google.com/p/go.net/websocket"
)

////////////////////////////////////////////////////////
//
//主程序
//
//参数说明：
//	启动服务器端：  goServer [port]				eg: goServer port = 6001
//  var url = "ws://www.strawberry.com:8000/goServer?data="
////////////////////////////////////////////////////////
func main() {

	log.Print(os.Args[1:])
	log.Printf(flag.Arg(1))
	// 第三方日志系统
	glog.Info("Entry Main"+"Server "+"! Port:", os.Args[1])
	glog.Info("System NumCPU: ", runtime.NumCPU())
	// 指定程序使用多核
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	glog.Flush()
	///////////////////////////////////////////////
	// 注册事件给服务器，建立成功。路由分发处理；服务器的路由功能。
	http.Handle("/GolangLTD", websocket.Handler(BuildConnection))
	//	if err := http.ListenAndServe(":"+os.Args[1], nil); err != nil {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		glog.Info("Entry nil", err.Error())
		glog.Flush()
		return
	}
	return

}

// 初始化服务器启动数据
func init() {

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			strerr := fmt.Sprintf("%s", err)
			glog.Info("init:", strerr)
		}
	}()
	// 初始化数据
	runningGoldMinerRoom.OnlineUsers = make(map[string]*OnlineUser)
	// 获取指定系统参数
	var SysCanShu = runtime.NumCPU()

	// 初始化cache
	//	var cache = cache2go.Cache("myCache")
	// 初始化map
	M = concurrent.NewConcurrentMap()
	// 屏蔽日志
	if SysCanShu != 4 {
		//f.exe -log_dir="./" -v=3
		flag.Set("alsologtostderr", "true") // 日志写入文件的同时，输出到stderr
		flag.Set("log_dir", "./log")        // 日志文件保存目录
		flag.Set("v", "3")                  // 配置V输出的等级。
		flag.Parse()
	}

	// 获取配置文件
	//ReadCsv_ConfigFile()
	return
}

// 建立链接
func BuildConnection(ws *websocket.Conn) {

	data := ws.Request().URL.Query().Get("data")
	glog.Info(data)
	//	if data == "" {
	//		// 直接返回数据
	//		glog.Info("data is Nil")
	//		glog.Flush()
	//		return
	//	}
	// 指针和地址都可以看做是数据块在内存存储的地址
	onlineUser := &OnlineUser{
		Connection: ws, // 链接的数据信息== 广播的数据的信息，广播给用户的数据；所有的链接的数据的信息
		MapSafe:    M,  // 并发安全的map
	}
	// 从客户端读取消息并发送到正在运行的 ActiveRoom 实例。
	onlineUser.PullFromClient()
}

// 数据分发处理
func (this *OnlineUser) PullFromClient() {

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			strerr := fmt.Sprintf("%s", err)
			glog.Info("PullFromClient", strerr)
		}
	}()
	glog.Info("Entry PullFromClient")
	for {
		var content string
		if err := websocket.Message.Receive(this.Connection, &content); err != nil {
			//glog.Info("content get error ：", err.Error())
			glog.Flush()
			break
		}
		if len(content) == 0 {
			glog.Info("len(content) == 0 ")
			glog.Flush()
			//this.Connection.Close()
			break
		}
		glog.Info("server get data ：" + content)

		// 高并发处理
		go this.SyncMessageFun(content)
	}
}

// 异步处理函数
func (this *OnlineUser) SyncMessageFun(content string) {

	var r Requestbody
	r.req = content
	if ProtocolData, err := r.Json2map(); err == nil {
		this.HandleCltProtocol(ProtocolData["Protocol"], ProtocolData["Protocol2"], ProtocolData)
	} else {
		glog.Info("SyncMessageFun :", err.Error())
	}

	// 清除并发连接
	runtime.Gosched()
	runtime.Goexit()

	return
}

// 结构体数据类型
type Requestbody struct {
	req string
}

// json转化为map:数据的处理
func (r *Requestbody) Json2map() (s map[string]interface{}, err error) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(r.req), &result); err != nil {
		glog.Info("Json2map:", err.Error())
		return nil, err
	}
	return result, nil
}

//--------------------------------------------------------------------------------
// 【错误提示】 G_Error_All_Proto
type G_Error_All struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	ErrCode   string // 错误码
	ErrMsg    string // 错误说明
}

// 消息处理函数,通过协议去处理数据;主协议处理
func (this *OnlineUser) HandleCltProtocol(protocol interface{}, protocol2 interface{}, ProtocolData map[string]interface{}) {

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			strerr := fmt.Sprintf("%s", err)
			//发消息给客户端
			ErrorST := G_Error_All{
				Protocol: Proto_Data.G_Error_Proto, // 主协议
				//				Protocol2: Proto2_Data.G_Error_All_Proto, // 子协议
				ErrCode: "80006",
				ErrMsg:  "亲，您发的数据的格式不对！" + strerr,
			}
			// 发送给玩家数据
			this.PlayerSendMessage(ErrorST)
		}
	}()
	// 通过过去的协议惊醒解析操作
	switch protocol {
	case float64(Proto_Data.Network_Data_Proto):
		{
			// 网络相关数据，心跳协议
			this.HandleCltProtocol2WD(protocol2, ProtocolData)
			break
		}
	}
	return
}
