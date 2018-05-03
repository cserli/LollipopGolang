package main

import (
	"encoding/json"
	"glog-master"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"zLollipopGolang/db_mysql"
	"zLollipopGolang/db_redis"
	"zLollipopGolang/globalData"
	"zLollipopGolang/loglogic"
	//	"strings"
	"time"
)

//var G_StWeiXinDatatmp map[string]*Global_Define.StWeiXinUserInfo // 微信的用户的结构信息

// 结构体定义
type STtoken struct {
	Access_token  string
	Expires_in    int
	Rufresh_token string
	Openid        string
	Scope         string
}

// 微信的人物的结构体
type STWeiXinPlayer struct {
	Openid     string
	Nickname   string
	Sex        uint32
	Language   string
	City       string
	Province   string
	Country    string
	Headimgurl string
	Privilege  string
}

// 微信的方式获取数据，
// GET方式获取网页的数据信息
func httpGet(code string, StrNoMd5 string, StrMd5 string) (int, string, string, string) {

	//=========================GONG===========================================
	// 判断数据库存在不
	var xcnameTMP = ""
	// 差分key
	strsplit := Strings_Split(StrNoMd5, "|")
	for i := 0; i < len(strsplit); i++ {
		if i == 0 {
		} else if i == 1 {
		} else if i == 2 {
			xcnameTMP = strsplit[i]
		}
	}
	//=========================GONG===========================================
	Strappid_secret := "appid=&secret=&code="
	if xcnameTMP == "beta" {
		Strappid_secret = "appid=&secret=&code="
	}

	resp, err := http.Get("https://api.weixin.qq.com/sns/oauth2/access_token?" + Strappid_secret + code + "&grant_type=authorization_code")
	if err != nil {
		glog.Info("resp err:", err.Error())
		return 0, "", "", ""
	}
	// 数据的请求
	//defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Info("body err:", err.Error())
		return 0, "", "", ""
	}
	// 测试的消息
	glog.Info("body:", string(body))
	// 通过josn 转化为结构体格式
	stb := &STtoken{}
	err = json.Unmarshal([]byte(body), &stb)
	if err != nil {
		glog.Info("Unmarshal faild")
	} else {
		glog.Info("Unmarshal success")
		glog.Info("Unmarshal success stb.Access_token:", stb.Access_token)
	}
	//======================================================================================================
	glog.Info("======================================================================================================")
	glog.Info("stb.Access_token：", stb.Access_token)
	glog.Info("stb.Openid：", stb.Openid)
	glog.Info("======================================================================================================")
	// 通过token 和openid 去获取用户信息  "https://api.weixin.qq.com/sns/userinfo?access_token="+access_token,
	resp1, err1 := http.Get("https://api.weixin.qq.com/sns/userinfo?access_token=" + stb.Access_token + "&openid=" + stb.Openid + "&lang=zh_CN")
	// https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
	if err1 != nil {
		glog.Info("playerdata", err1.Error())
		return 0, "", "", ""
	}
	// 数据的请求
	//defer resp1.Body.Close()
	body2, err2 := ioutil.ReadAll(resp1.Body)
	if err2 != nil {
		glog.Info("playerdata", err2.Error())
		return 0, "", "", ""
	}
	// 测试的消息
	glog.Info("playerdata:", string(body2))
	// 解析数据，保存数据库操作
	stbtmp := &STWeiXinPlayer{}
	err3 := json.Unmarshal([]byte(body2), &stbtmp)
	if err3 != nil {
		glog.Info("Unmarshal faild")
	} else {
		glog.Info("Unmarshal success")
		glog.Info("Unmarshal success stbtmp:", stbtmp)
		glog.Info("Unmarshal success stbtmp.headimgurl:", stbtmp.Headimgurl)
	}
	//==========================================================================
	// 保存数据库的操作，主要是记录数据的操作
	//时间戳到具体显示的转化
	iLastLoginTime := time.Now().Unix()
	// 结构体赋值
	var StWeiXinData Global_Define.StWeiXinUserInfo
	StWeiXinData.Openid = stbtmp.Openid
	StWeiXinData.Nickname = stbtmp.Nickname
	//StWeiXinData.Nickname = strings.Replace(StWeiXinData.Nickname, " ", "%20", -1)
	StWeiXinData.Sex = uint32(stbtmp.Sex)
	StWeiXinData.Province = stbtmp.Province
	//StWeiXinData.Province = strings.Replace(StWeiXinData.Province, " ", "", -1)
	StWeiXinData.City = stbtmp.City
	//StWeiXinData.City = strings.Replace(StWeiXinData.City, " ", "", -1)
	StWeiXinData.Country = stbtmp.Country
	//StWeiXinData.Country = strings.Replace(StWeiXinData.Country, " ", "", -1)
	StWeiXinData.Headimgurl = stbtmp.Headimgurl
	StWeiXinData.IdentifykeySJ = Log_Eio.FilePort                  // 端口数据就可以
	StWeiXinData.IdentifykeyHD = strconv.Itoa(int(iLastLoginTime)) // 存保存的时间,报名的时间

	glog.Info(" 结构体的数据结构：StWeiXinData", StWeiXinData)

	// 保存数据库
	var iiret = 1
	if len(StWeiXinData.Openid) == 0 {
		glog.Info(" Send BaoMing　fail!!! ,Openid  == 0")
		iiret = 0
	} else {
		//
		// go dbif.RegNewUserForWeiXin(StWeiXinData)
		// dbif.RegNewUserForWeiXinbak(StWeiXinData)
		// 判断是否发送数据
		_, ok := Global_Define.G_StWeiXinDatatmpPaihang[StWeiXinData.Openid]
		if !ok {
			dbif.RegNewUserForWeiXinbak(StWeiXinData)
		}
		glog.Info("stbtmp.Headimgurl", stbtmp.Headimgurl)
		//

		if Global_Define.G_StWeiXinDatatmpPaihang[StWeiXinData.Openid] != nil {
			if Global_Define.G_StWeiXinDatatmpPaihang[StWeiXinData.Openid].Headimgurl != StWeiXinData.Headimgurl || Global_Define.G_StWeiXinDatatmpPaihang[StWeiXinData.Openid].Nickname != StWeiXinData.Nickname {
				// 发送更新程序，同步http服务器
				glog.Info(" 发送微信，同步http服务器")
				//StWeiXinData.Nickname = strings.Replace(StWeiXinData.Nickname, " ", "%20", -1)

				req := "http://" + Log_Eio.ServerURl + "9094/TJData?Protocol=6&Protocol2=14&Openid=" + StWeiXinData.Openid + "&Nickname=" + StWeiXinData.Nickname + "&Sex=" + strconv.Itoa(int(StWeiXinData.Sex)) + "&Language=" +
					StWeiXinData.Language + "&City=" + StWeiXinData.City + "&Province=" + StWeiXinData.Privilege + "&Country=" + StWeiXinData.Country + "&Headimgurl=" + StWeiXinData.Headimgurl + "&Privilege=" +
					StWeiXinData.Privilege + "&IdentifykeySJ=" + StWeiXinData.IdentifykeySJ + "&IdentifykeyHD=" + StWeiXinData.IdentifykeyHD + "&Createtime=" + StWeiXinData.IdentifykeyHD

				u, _ := url.Parse(req)
				q := u.Query()
				u.RawQuery = q.Encode() //urlencode
				go http.Get(u.String())

				//go http.Get(url)
				// 更新内存
				StWeiXinDataok := new(Global_Define.StWeiXinUserInfo)
				StWeiXinDataok.Openid = stbtmp.Openid
				StWeiXinDataok.Nickname = stbtmp.Nickname
				StWeiXinDataok.Sex = uint32(stbtmp.Sex)
				StWeiXinDataok.Province = stbtmp.Province
				StWeiXinDataok.City = stbtmp.City
				StWeiXinDataok.Country = stbtmp.Country
				StWeiXinDataok.Headimgurl = stbtmp.Headimgurl
				StWeiXinDataok.IdentifykeySJ = Log_Eio.FilePort                  // 端口数据就可以
				StWeiXinDataok.IdentifykeyHD = strconv.Itoa(int(iLastLoginTime)) // 存保存的时间,报名的时间
				Global_Define.G_StWeiXinDatatmpPaihang[StWeiXinData.Openid] = StWeiXinDataok
				// 更新数据库数据
				dbif.UpdateHeadUrlWeiXin(stbtmp.Openid, stbtmp.Headimgurl, stbtmp.Nickname)
			}
		}

		//=========================TJ===========================================
		// 判断数据库存在不
		var loginname = ""
		var gamename = ""
		var xcname = ""
		// 差分key
		strsplit := Strings_Split(StrNoMd5, "|")
		for i := 0; i < len(strsplit); i++ {
			if i == 0 {
				loginname = strsplit[i]
			} else if i == 1 {
				gamename = strsplit[i]
			} else if i == 2 {
				xcname = strsplit[i]
			}
		}

		// 商家的所有的游戏的对应的玩家
		// key = "TJDataByL"
		// hkey = MD5(LoginName)
		// WXPlayerStruct{}
		Redis_DB.Redis_Write_WeiXinPlayData_InRedis_Data(loginname, gamename, xcname, StWeiXinData)

		// key = "TJDataByLAndG"
		// hkey = MD5(LoginName+GameName)
		// WXPlayerStruct{}
		Redis_DB.Redis_Write_WeiXinPlayData_InRedis_Data_ByLoginOrGame(loginname, gamename, xcname, StWeiXinData)

		// key = "TJDataByLAndGAndXC"
		// hkey = MD5(LoginName+GameName+XCName)
		// WXPlayerStruct{}
		Redis_DB.Redis_Write_WeiXinPlayData_InRedis_Data_ByLoginOrGameOrXC(loginname, gamename, xcname, StWeiXinData)

		//=========================TJ===========================================
		// 保存微信用户的数据
		Infotmp := new(Global_Define.StWeiXinUserInfo)
		Infotmp.Openid = stbtmp.Openid
		Infotmp.Nickname = stbtmp.Nickname
		Infotmp.Sex = uint32(stbtmp.Sex)
		Infotmp.Province = stbtmp.Province
		Infotmp.City = stbtmp.City
		Infotmp.Country = stbtmp.Country
		Infotmp.Headimgurl = stbtmp.Headimgurl
		Infotmp.IdentifykeySJ = Log_Eio.FilePort
		Infotmp.IdentifykeyHD = strconv.Itoa(int(iLastLoginTime)) // 存保存的时间,报名的时间

		G_StWeiXinDatatmp[Infotmp.Openid+"|"+StrMd5] = Infotmp

		//  同步数据
		//StWeiXinData.Nickname = strings.Replace(StWeiXinData.Nickname, " ", "%20", -1)
		go Sync_WeiXin_Data(StWeiXinData, gamename, loginname)

		go Sync_TJWeiXin_Data(StWeiXinData, StrNoMd5+"|"+StrMd5)
	}
	// 释放资源
	resp.Body.Close()
	resp1.Body.Close()
	return int(iiret), stbtmp.Openid, stbtmp.Nickname, stbtmp.Headimgurl
}

// 重玩统计
// StWeiXinData Global_Define.StWeiXinUserInfo
func Sync_WeiXin_Data_ReStart(Data *Global_Define.StWeiXinUserInfo, gamename string, loginname string, xcname string) {

	var StWeiXinData Global_Define.StWeiXinUserInfo
	iLastLoginTime := time.Now().Unix()

	StWeiXinData.Openid = Data.Openid
	StWeiXinData.Nickname = Data.Nickname
	StWeiXinData.Sex = uint32(Data.Sex)
	StWeiXinData.Province = Data.Province
	StWeiXinData.City = Data.City
	StWeiXinData.Country = Data.Country
	StWeiXinData.Headimgurl = Data.Headimgurl
	// StWeiXinData.IdentifykeySJ = loginname
	StWeiXinData.IdentifykeyHD = strconv.Itoa(int(iLastLoginTime))

	//
	dbif.RegNewUserForWeiXinbak(StWeiXinData)
	// 商家的所有的游戏的对应的玩家
	// key = "TJDataByL"
	// hkey = MD5(LoginName)
	// WXPlayerStruct{}
	Redis_DB.Redis_Write_WeiXinPlayData_InRedis_Data(loginname, gamename, xcname, StWeiXinData)

	// key = "TJDataByLAndG"
	// hkey = MD5(LoginName+GameName)
	// WXPlayerStruct{}
	Redis_DB.Redis_Write_WeiXinPlayData_InRedis_Data_ByLoginOrGame(loginname, gamename, xcname, StWeiXinData)

	// key = "TJDataByLAndGAndXC"
	// hkey = MD5(LoginName+GameName+XCName)
	// WXPlayerStruct{}
	Redis_DB.Redis_Write_WeiXinPlayData_InRedis_Data_ByLoginOrGameOrXC(loginname, gamename, xcname, StWeiXinData)

	//  同步数据
	go Sync_WeiXin_Data(StWeiXinData, gamename, loginname)

	StrNoMd5 := loginname + "|" + gamename + "|" + xcname
	Sync_TJWeiXin_Data(StWeiXinData, StrNoMd5+"|"+Global_Define.MD5ToStringFromString(loginname))

	return
}

//  同步数据
// http://api1..cn:9000/API/SetGameData.ashx?method=InsertData&OpenID=xx&NickName=xx&Sex=0&Language=xx&City=xx&Province=xx&Country=xx&Headimgurl=xx&IdentifykeyHD=xx&GameName=xx
// GameType 表示  1 =测试  2 = beta   3=  正式
func Sync_WeiXin_Data(data Global_Define.StWeiXinUserInfo, strGamename string, LoginName string) {

	// 解析数据
	req := "http://api1..cn:9000/API/SetGameData.ashx?method=InsertData&OpenID=" + data.Openid + "&NickName=" + data.Nickname + "&Sex=" + strconv.Itoa(int(data.Sex)) + "&Language=" + data.Language +
		"&City=" + data.City + "&Province=" + data.Province + "&Country=" + data.Country + "&Headimgurl=" + data.Headimgurl + "&IdentifykeyHD=" + data.IdentifykeyHD + "&GameName=" + strGamename + "&GameType=" + GameType +
		"&StoreName=" + LoginName

	u, _ := url.Parse(req)
	q := u.Query()
	u.RawQuery = q.Encode() //urlencode
	http.Get(u.String())
	//http.Get(url)
	//	resp, err1 := http.Get(url)
	//	if err1 != nil {
	//		glog.Error("Send_FirmGetAward_JiLu_By_Hsttp get error", err1.Error())
	//		return
	//	}
	//	_, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		glog.Error("body err:", err.Error())
	//		return
	//	}
	//glog.Info("body :", body)
	return
}

func Sync_TJWeiXin_Data(StWeiXinData Global_Define.StWeiXinUserInfo, L_G_X_Name string) {
	// 解析数据
	req := "http://" + Log_Eio.ServerURl + "9094/TJData?Protocol=6&Protocol2=16&Openid=" + StWeiXinData.Openid + "&Nickname=" + StWeiXinData.Nickname + "&Sex=" + strconv.Itoa(int(StWeiXinData.Sex)) + "&Language=" +
		StWeiXinData.Language + "&City=" + StWeiXinData.City + "&Province=" + StWeiXinData.Privilege + "&Country=" + StWeiXinData.Country + "&Headimgurl=" + StWeiXinData.Headimgurl + "&Privilege=" +
		StWeiXinData.Privilege + "&IdentifykeySJ=" + StWeiXinData.IdentifykeySJ + "&IdentifykeyHD=" + StWeiXinData.IdentifykeyHD + "&L_G_X_Name=" + L_G_X_Name + "&Createtime=" + StWeiXinData.IdentifykeyHD
	glog.Info("req:", req)
	u, _ := url.Parse(req)
	q := u.Query()
	u.RawQuery = q.Encode() //urlencode
	http.Get(u.String())
	//http.Get(url)
	return
}

// 同步数据
func Sync_TJJF_Data(strOpenID string, strGameName string, iJiFen uint32, strOpenIDAndX string) {
	// 解析数据
	// http://api1..cn:9000/api/SetGameData.ashx?method=InsertJiFen&OpenID=xxx&NickName=XXX&JiFen=xxx&Game=xxx
	req := "http://api1..cn:9000/api/SetGameData.ashx?method=InsertJiFen&OpenID=" + strOpenID + "&NickName=" + G_StWeiXinDatatmp[strOpenIDAndX].Nickname + "&JiFen=" + strconv.Itoa(int(iJiFen)) + "&Game=" + strGameName
	u, _ := url.Parse(req)
	q := u.Query()
	u.RawQuery = q.Encode() //urlencode
	http.Get(u.String())
	//http.Get(url)
	return
}
