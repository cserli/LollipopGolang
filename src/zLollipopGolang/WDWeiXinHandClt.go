package main

import (
	"encoding/json"
	"glog-master"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"zLollipopGolang/db_mysql"
	"zLollipopGolang/globalData"
)

// 微信的方式获取数据，
// GET方式获取网页的数据信息
func httpGetWD(code string) (int, string, string, string) {

	//=========================GONG===========================================
	Strappid_secret := "appid=&secret=&code="
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
	StWeiXinData.IdentifykeySJ = "2102"                            // 端口数据就可以
	StWeiXinData.IdentifykeyHD = strconv.Itoa(int(iLastLoginTime)) // 存保存的时间,报名的时间

	glog.Info(" 结构体的数据结构：StWeiXinData", StWeiXinData)

	// 保存数据库
	var iiret = 1
	if len(StWeiXinData.Openid) == 0 {
		glog.Info(" Send BaoMing　fail!!! ,Openid  == 0")
		iiret = 0
	} else {
		// 判断是否发送数据
		_, ok := Global_Define.G_StWeiXinDatatmpPaihang[StWeiXinData.Openid]
		if !ok {
			dbif.RegNewUserForWeiXinbakWD(StWeiXinData)
		}
		glog.Info("stbtmp.Headimgurl", stbtmp.Headimgurl)
		//

		if Global_Define.G_StWeiXinDatatmpPaihang[StWeiXinData.Openid] != nil {
			if Global_Define.G_StWeiXinDatatmpPaihang[StWeiXinData.Openid].Headimgurl != StWeiXinData.Headimgurl || Global_Define.G_StWeiXinDatatmpPaihang[StWeiXinData.Openid].Nickname != StWeiXinData.Nickname {

				// 更新内存
				StWeiXinDataok := new(Global_Define.StWeiXinUserInfo)
				StWeiXinDataok.Openid = stbtmp.Openid
				StWeiXinDataok.Nickname = stbtmp.Nickname
				StWeiXinDataok.Sex = uint32(stbtmp.Sex)
				StWeiXinDataok.Province = stbtmp.Province
				StWeiXinDataok.City = stbtmp.City
				StWeiXinDataok.Country = stbtmp.Country
				StWeiXinDataok.Headimgurl = stbtmp.Headimgurl
				StWeiXinDataok.IdentifykeyHD = strconv.Itoa(int(iLastLoginTime)) // 存保存的时间,报名的时间
				Global_Define.G_StWeiXinDatatmpPaihang[StWeiXinData.Openid] = StWeiXinDataok
				// 更新数据库数据
				dbif.UpdateHeadUrlWeiXin(stbtmp.Openid, stbtmp.Headimgurl, stbtmp.Nickname)
			}
		}
	}
	// 释放资源
	resp.Body.Close()
	resp1.Body.Close()
	return int(iiret), stbtmp.Openid, stbtmp.Nickname, stbtmp.Headimgurl
}
