package main

import (
	"encoding/json"
	"glog-master"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"zLollipopGolang/db_mysql"
	"zLollipopGolang/db_redis"
	"zLollipopGolang/globalData"
	"zLollipopGolang/loglogic"
)

// 结构体定义
type STQQtoken struct {
	Access_token  string
	Expires_in    int
	Rufresh_token string
	Openid        string
	Scope         string
}

// qq的人物的结构体
type STQQPlayer struct {
	Ret                int
	Msg                string
	Is_lost            int
	Nickname           string
	Gender             string // 男女
	Province           string
	City               string
	Year               string
	Figureurl          string
	Figureurl_1        string
	Figureurl_2        string
	Figureurl_qq_1     string
	Figureurl_qq_2     string
	Is_yellow_vip      string
	Vip                string
	Yellow_vip_level   string
	Level              string
	Is_yellow_year_vip string
}

// QQ的方式获取数据，
// GET方式获取网页的数据信息
func httpQQGet(code string, strOpenID string, StrNoMd5 string, StrMd5 string) (int, string, string, string) {

	resp1, err1 := http.Get("https://graph.qq.com/user/get_user_info?oauth_consumer_key=&access_token=" + code + "&openid=" + strOpenID + "&format=json")
	if err1 != nil {
		glog.Info("playerdata", err1.Error())
		return 0, "", "", ""
	}
	// 数据的请求
	body2, err2 := ioutil.ReadAll(resp1.Body)
	if err2 != nil {
		glog.Info("playerdata", err2.Error())
		return 0, "", "", ""
	}
	// 测试的消息
	glog.Info("playerdata:", string(body2))
	// 解析数据，保存数据库操作
	//	stbtmp := &STWeiXinPlayer{}
	stbtmp := &STQQPlayer{}
	err3 := json.Unmarshal([]byte(body2), &stbtmp)
	if err3 != nil {
		glog.Info("Unmarshal faild")
	} else {
		glog.Info("Unmarshal success")
		glog.Info("Unmarshal success stbtmp:", stbtmp)
		glog.Info("Unmarshal success stbtmp.Nickname:", stbtmp.Nickname)
	}
	//==========================================================================
	// 保存数据库的操作，主要是记录数据的操作
	//时间戳到具体显示的转化
	iLastLoginTime := time.Now().Unix()
	// 结构体赋值
	var StWeiXinData Global_Define.StWeiXinUserInfo
	StWeiXinData.Openid = strOpenID
	StWeiXinData.Nickname = stbtmp.Nickname
	if stbtmp.Gender == "男" {
		StWeiXinData.Sex = uint32(1)
	} else if stbtmp.Gender == "女" {
		StWeiXinData.Sex = uint32(2)
	}
	StWeiXinData.Province = stbtmp.Province
	StWeiXinData.City = stbtmp.City
	StWeiXinData.Headimgurl = stbtmp.Figureurl_qq_2
	StWeiXinData.IdentifykeySJ = Log_Eio.FilePort                  // 端口数据就可以
	StWeiXinData.IdentifykeyHD = strconv.Itoa(int(iLastLoginTime)) // 存保存的时间,报名的时间

	glog.Info(" 结构体的数据结构：StWeiXinData", StWeiXinData)

	// 保存数据库
	var iiret = 1
	if len(StWeiXinData.Openid) == 0 {
		glog.Info(" Send BaoMing　fail!!! ,Openid  == 0")
		iiret = 0
	} else {
		dbif.RegNewUserForWeiXin(StWeiXinData)

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
		Infotmp.Openid = strOpenID
		Infotmp.Nickname = stbtmp.Nickname
		if stbtmp.Gender == "男" {
			Infotmp.Sex = uint32(1)
		} else if stbtmp.Gender == "女" {
			Infotmp.Sex = uint32(2)
		}
		Infotmp.Province = stbtmp.Province
		Infotmp.City = stbtmp.City
		Infotmp.Headimgurl = stbtmp.Figureurl_qq_2
		Infotmp.IdentifykeySJ = Log_Eio.FilePort
		Infotmp.IdentifykeyHD = strconv.Itoa(int(iLastLoginTime)) // 存保存的时间,报名的时间

		G_StWeiXinDatatmp[Infotmp.Openid+"|"+StrMd5] = Infotmp
	}
	// 释放资源
	resp1.Body.Close()
	return int(iiret), strOpenID, stbtmp.Nickname, stbtmp.Figureurl_qq_2
}

// WeiBo的人物的结构体
type STWeiBoPlayer struct {
	Id                  int64
	Idstr               string
	Class               int
	Screen_name         string
	Name                string
	Province            string
	City                int
	Location            string
	Description         string
	Url                 string
	Profile_image_url   string
	Profile_url         string
	Domain              string
	Gender              string
	Followers_count     string
	Friends_count       int
	Pagefriends_count   int
	Statuses_count      int
	Favourites_count    int
	Created_at          string
	Following           bool
	Allow_all_act_msg   bool
	Geo_enabled         bool
	Verified            bool
	Verified_type       int
	Remark              string
	Status              string
	Ptype               int
	Allow_all_comment   bool
	Avatar_large        string
	Avatar_hd           string
	Verified_reason     string
	Verified_trade      string
	Verified_reason_url string
	Verified_source     string
	Verified_source_url string
	Follow_me           bool
	Online_status       int
	Bi_followers_count  int
	Lang                string
	Star                int
	Mbtype              int
	Mbrank              int
	Block_word          int
	Block_app           int
	Credit_score        int
	User_ability        int
	Urank               int
}

// 微博的方式获取数据，
// GET方式获取网页的数据信息
func httpWeiBoGet(code string, strOpenID string, StrNoMd5 string, StrMd5 string) (int, string, string, string) {

	resp1, err1 := http.Get("https://api.weibo.com/2/users/show.json?access_token=" + code + "&uid=" + strOpenID)
	if err1 != nil {
		glog.Info("playerdata", err1.Error())
		return 0, "", "", ""
	}
	// 数据的请求
	body2, err2 := ioutil.ReadAll(resp1.Body)
	if err2 != nil {
		glog.Info("playerdata", err2.Error())
		return 0, "", "", ""
	}
	// 测试的消息
	glog.Info("playerdata:", string(body2))
	// 解析数据，保存数据库操作
	stbtmp := &STWeiBoPlayer{}
	err3 := json.Unmarshal([]byte(body2), &stbtmp)
	if err3 != nil {
		glog.Info("Unmarshal faild", err3)
	} else {
		glog.Info("Unmarshal success")
		glog.Info("Unmarshal success stbtmp:", stbtmp)
		glog.Info("Unmarshal success stbtmp.Screen_name:", stbtmp.Screen_name)
	}
	//==========================================================================
	// 保存数据库的操作，主要是记录数据的操作
	//时间戳到具体显示的转化
	iLastLoginTime := time.Now().Unix()
	// 结构体赋值
	var StWeiXinData Global_Define.StWeiXinUserInfo
	StWeiXinData.Openid = strOpenID
	StWeiXinData.Nickname = stbtmp.Screen_name
	if stbtmp.Gender == "m" {
		StWeiXinData.Sex = uint32(1)
	} else if stbtmp.Gender == "女" {
		StWeiXinData.Sex = uint32(2)
	}
	StWeiXinData.Province = stbtmp.Province
	StWeiXinData.City = strconv.Itoa(stbtmp.City)
	StWeiXinData.Headimgurl = stbtmp.Profile_image_url
	StWeiXinData.IdentifykeySJ = Log_Eio.FilePort                  // 端口数据就可以
	StWeiXinData.IdentifykeyHD = strconv.Itoa(int(iLastLoginTime)) // 存保存的时间,报名的时间

	glog.Info(" 结构体的数据结构：StWeiXinData", StWeiXinData)

	// 保存数据库
	var iiret = 1
	if len(StWeiXinData.Openid) == 0 {
		glog.Info(" Send BaoMing　fail!!! ,Openid  == 0")
		iiret = 0
	} else {
		dbif.RegNewUserForWeiXin(StWeiXinData)

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
		Infotmp.Openid = strOpenID
		Infotmp.Nickname = stbtmp.Screen_name
		if stbtmp.Gender == "m" {
			Infotmp.Sex = uint32(1)
		} else if stbtmp.Gender == "f" {
			Infotmp.Sex = uint32(2)
		}
		Infotmp.Province = stbtmp.Province
		Infotmp.City = strconv.Itoa(stbtmp.City)
		Infotmp.Headimgurl = stbtmp.Profile_image_url
		Infotmp.IdentifykeySJ = Log_Eio.FilePort
		Infotmp.IdentifykeyHD = strconv.Itoa(int(iLastLoginTime)) // 存保存的时间,报名的时间

		G_StWeiXinDatatmp[Infotmp.Openid+"|"+StrMd5] = Infotmp
	}
	// 释放资源
	resp1.Body.Close()
	return int(iiret), strOpenID, stbtmp.Screen_name, stbtmp.Profile_image_url
}
