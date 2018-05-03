package main

import (
	"fmt"
	"glog-master"
	"os"
	"time"
	"zLollipopGolang/loglogic"
	"zLollipopGolang/protocolfile"
	"zLollipopGolang/protocolfile/Proto2"
)

// 游戏结束发送给所有的玩家
func PC_Send_AllPlayer_Data() {
	return
}

// 超时踢人
func Time_Out_SeverKill(MD5 string) {

	for key, second := range runningGoldMinerRoom.OnlineUsers {
		glog.Info("玩家", key)
		if len(key) > 32 && second.StrMD5 == MD5 {
			continue
		}
		glog.Info("清除超时玩家", key)
		// 关闭链接
		second.Connection.Close()
		// 清除数据
		delete(runningGoldMinerRoom.OnlineUsers, key)
		//second.MapSafe.Remove(key + "|GGameConfigOfPlayer")
	}
}

// 清楚链接的计数器
func GamePanDuanTimer0428() {
	GGameStartTimer := time.NewTicker(20 * time.Second)
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
			GGameStartTimer.Stop()
			glog.Info("GamePanDuanTimer", ErrorST)
			// 发送给玩家数据
			glog.Flush()
		}
	}()

	for {
		select {
		case <-GGameStartTimer.C:
			{

				// 循环取数据
				// 后面优化，用一个结构专门存储商家，比较好处理
				for key, second := range runningGoldMinerRoom.OnlineUsers {
					if len(key) == 32 {
						// 获取数据
						val, _ := second.MapSafe.Get(key + "|MapG_HartJiShu") // 心跳计数
						if val == nil {
							continue
						}
						valG_HartJiShutmp, _ := second.MapSafe.Get(key + "|G_HartJiShutmp")
						if valG_HartJiShutmp == nil {
							continue
						}
						if valG_HartJiShutmp.(int) < val.(int) {
							second.MapSafe.Put(key+"|G_HartJiShutmp", val.(int))
							continue
						}
						valG_HartJiShutmp1, _ := second.MapSafe.Get(key + "|G_HartJiShutmp")
						if valG_HartJiShutmp1.(int) >= val.(int) {
							second.MapSafe.Put(key+"|G_HartJiShutmp", 0)
							second.MapSafe.Put(key+"|MapG_HartJiShu", 1)
							delete(runningGoldMinerRoom.OnlineUsers, key)
							//second.MapSafe.Remove(key + "|MapConnnetName") // 争议 链接
							second.MapSafe.Remove(key + "|connect")
							second.MapSafe.Remove("server")
							glog.Info("清除 超时链接：", key)
							//Time_Out_SeverKill(key)
						}
					}
				}

			}
		}
	}
}

// 清楚链接的计数器
func GamePanDuanTimer() {
	GGameStartTimerPC := time.NewTicker(20 * time.Second)
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			strerr := fmt.Sprintf("%s", err)
			glog.Info("清除超时PC：", strerr)
			GGameStartTimerPC.Stop() // 暂停计时器
		}
	}()

	for {
		select {
		case <-GGameStartTimerPC.C:
			{
				for itr := M.Iterator(); itr.HasNext(); {
					k, v, _ := itr.Next()
					var key = ""
					var keyName = ""
					// 差分key
					strsplit := Strings_Split(k.(string), "|")
					for i := 0; i < len(strsplit); i++ {
						if len(strsplit) > 2 {
							continue
						}
						if i == 0 {
							key = strsplit[i]
						}
						// 获取链接的名字
						if i == len(strsplit)-1 {
							keyName = strsplit[i]
						}
						// 发消息
						if len(k.(string)) < 41 && keyName == "connect" {

							switch v.(interface{}).(type) {
							case *OnlineUser:
								{
									// 获取数据
									val, _ := v.(interface{}).(*OnlineUser).MapSafe.Get(key + "|MapG_HartJiShu") // 心跳计数
									if val == nil {
										continue
									}
									valG_HartJiShutmp, _ := v.(interface{}).(*OnlineUser).MapSafe.Get(key + "|G_HartJiShutmp")
									if valG_HartJiShutmp == nil {
										continue
									}
									if valG_HartJiShutmp.(int) < val.(int) {
										v.(interface{}).(*OnlineUser).MapSafe.Put(key+"|G_HartJiShutmp", val.(int))
										continue
									}
									valG_HartJiShutmp1, _ := v.(interface{}).(*OnlineUser).MapSafe.Get(key + "|G_HartJiShutmp")
									//									glog.Info("非PC", valG_HartJiShutmp1.(int), "2222wwwwww", val.(int))
									if valG_HartJiShutmp1.(int) >= val.(int) {
										v.(interface{}).(*OnlineUser).MapSafe.Put(key+"|G_HartJiShutmp", 0)
										v.(interface{}).(*OnlineUser).MapSafe.Put(key+"|MapG_HartJiShu", 1)
										v.(interface{}).(*OnlineUser).MapSafe.Remove(key + "|connect")
										v.(interface{}).(*OnlineUser).MapSafe.Remove("server")
										v.(interface{}).(*OnlineUser).Set_GameStart1(false, key)
										glog.Info("清除 超时链接：", key)
									}
								}
							default:
								//glog.Info("非PC")
							}
						}
					}
				}
			}
		}
	}
}

// 清除超时玩家
func Timer_OutPlayer() {

	GGameStartTimer := time.NewTicker(1 * time.Second)
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			strerr := fmt.Sprintf("%s", err)
			glog.Info("清除超时玩家：", strerr)
			GGameStartTimer.Stop() // 暂停计时器
		}
	}()

	for {
		select {
		case <-GGameStartTimer.C:
			{
				iLastLoginTime := time.Now().Unix()
				for itr := M.Iterator(); itr.HasNext(); {
					k, v, _ := itr.Next()
					if len(k.(string)) < 35 {
						continue
					}
					var key = ""
					var keyName = ""
					var keyMD5 = ""
					// 差分key
					strsplit := Strings_Split(k.(string), "|") // key = openid|XC|name
					for i := 0; i < len(strsplit); i++ {
						if len(strsplit) < 2 {
							continue
						}
						switch v.(interface{}).(type) {
						case *OnlineUser:
							{
								if i == 0 {
									key = strsplit[i]
								} else if i == 1 {
									key = key + "|" + strsplit[i]
									keyMD5 = strsplit[i]
								}
								// 获取链接的名字
								if i == len(strsplit)-1 {
									keyName = strsplit[i]
								}
								if keyMD5 == v.(interface{}).(*OnlineUser).StrMD5 && keyName == "connect" {
									// 发消息
									// glog.Info("key:!", k.(string))
									// 获取数据
									val, err1 := v.(interface{}).(*OnlineUser).MapSafe.Get(key + "|GGameConfigOfPlayer") // 心跳计数
									if err1 != nil {
										glog.Info("清除超时玩家定时器!,err:", err1.Error())
										continue
									}
									if val == nil {
										glog.Info("清除超时玩家定时器!,val == nil:")
										continue
									}
									// glog.Info("iLastLoginTime!!!", iLastLoginTime)
									// glog.Info("Save  Time：", val.(int64))
									// glog.Info("  Time：", (iLastLoginTime - val.(int64)))
									if iLastLoginTime-val.(int64) > 30 {

										glog.Info("清除超时玩家定时器!", k.(string))
										// 发送超时
										Pc_PlayeNInfo := Proto2_Data.HS2C_Send_TimeOut_Player_By_Server{
											Protocol:  Proto_Data.FE_User_Data_Proto,                        // 主协议
											Protocol2: Proto2_Data.HS2C_Send_TimeOut_Player_By_Server_Proto, // 子协议
										}
										v.(interface{}).(*OnlineUser).PlayerSendMessage(Pc_PlayeNInfo)
										v.(interface{}).(*OnlineUser).Connection.Close()
										v.(interface{}).(*OnlineUser).MapSafe.Remove(key + "|GGameConfigOfPlayer")
										v.(interface{}).(*OnlineUser).MapSafe.Remove(key + "|connect")
										valGGameMarktmp, _ := v.(interface{}).(*OnlineUser).MapSafe.Get(key + "|GGameMarktmp")
										if valGGameMarktmp != nil {
											v.(interface{}).(*OnlineUser).MapSafe.Remove(key + "|GGameMarktmp")
										}
									}
								}
							}
						default:
							//glog.Info("非玩家")
						}
					}
				}
				//
			}
		}
	}
}

// 清除超时玩家
func Timer_OutPlayer0428() {

	GGameStartTimer := time.NewTicker(1 * time.Second)
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			strerr := fmt.Sprintf("%s", err)
			glog.Info("清除超时玩家：", strerr)
			GGameStartTimer.Stop() // 暂停计时器
		}
	}()

	for {
		select {
		case <-GGameStartTimer.C:
			{
				iLastLoginTime := time.Now().Unix()
				for key, second := range runningGoldMinerRoom.OnlineUsers {

					if len(key) <= 32 {
						//glog.Info("second.MapSafe!")
						continue
					}
					glog.Info("key:!", key)
					// 判断数据
					if second == nil || second.MapSafe == nil {
						glog.Info("second.MapSafe!")
						continue
					}
					// 获取数据
					val, err1 := second.MapSafe.Get(key + "|GGameConfigOfPlayer") // 心跳计数
					if err1 != nil {
						glog.Info("清除超时玩家定时器!,err:", err1.Error())
						continue
					}

					if val == nil {
						glog.Info("清除超时玩家定时器!,val == nil:")
						continue
					}
					glog.Info("iLastLoginTime!!!", iLastLoginTime)
					glog.Info("Save  Time：", val.(int64))
					glog.Info("  Time：", (iLastLoginTime - val.(int64)))
					if iLastLoginTime-val.(int64) > 30 {

						glog.Info("清除超时玩家定时器!")
						// 发送超时
						Pc_PlayeNInfo := Proto2_Data.HS2C_Send_TimeOut_Player_By_Server{
							Protocol:  Proto_Data.FE_User_Data_Proto,                        // 主协议
							Protocol2: Proto2_Data.HS2C_Send_TimeOut_Player_By_Server_Proto, // 子协议
						}
						second.PlayerSendMessage(Pc_PlayeNInfo)
						// 关闭链接
						second.Connection.Close()
						// 清除数据
						delete(runningGoldMinerRoom.OnlineUsers, key)
						second.MapSafe.Remove(key + "|GGameConfigOfPlayer")
						second.MapSafe.Remove(key + "|connect")
						valGGameMarktmp, _ := second.MapSafe.Get(key + "|GGameMarktmp")
						if valGGameMarktmp != nil {
							second.MapSafe.Remove(key + "|GGameMarktmp")
						}
					}
				}
			}
		}
	}
}

// 获取链接数
func GetPlayerSendMax() int {
	var icount = 0
	// 获取链接书的多少
	for _, second := range runningGoldMinerRoom.OnlineUsers {
		if second != nil {
			icount++
		}
	}
	return icount
}

// 清除CD次数
func GameCDClearTimer() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			strerr := fmt.Sprintf("%s", err)
			glog.Info("清除奖券定时器：", strerr)
		}
	}()
	// 50分钟检测一次
	GGameStartTimer := time.NewTicker(3000 * time.Second)
	for {
		select {
		case <-GGameStartTimer.C:
			{
				strime := time.Now().Format("2006-01-02 15:04:05")
				rs := []rune(strime)
				fmt.Println("rs:", string(rs[11:13]))
				if string(rs[11:13]) == "00" {
					// 重启
					if Log_Eio.SysCanShu != 1 { // 去除测试环境
						os.Exit(0)
					}
				}
			}
		}
	}
}
