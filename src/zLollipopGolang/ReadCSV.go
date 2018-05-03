package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"zLollipopGolang/globalData"
)

func ReadCsv_ConfigFile() {
	// 获取配置文件
	ReadCsv_ConfigFile_GameInfoST_Fun()
	ReadCsv_ConfigFile_ActivitiesInfoST_Fun()
	ReadCsv_ConfigFile_GameConfigInfoST_Fun()
	ReadCsv_ConfigFile_StCard2List_Fun()
	return
}

// 游戏的基本的ID的数据信息
func ReadCsv_ConfigFile_GameInfoST_Fun() bool {
	// 获取数据，按照文件
	fileName := "GameInfo.csv"
	fileName = "./csv/" + fileName
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		return false
	}
	// 读取文件数据
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, _ := r2.ReadAll()
	sz := len(ss)
	// 循环取数据
	for i := 1; i < sz; i++ {

		Infotmp := new(Global_Define.StGameListInfo)
		igame, _ := strconv.Atoi(ss[i][0])
		Infotmp.GameId = uint32(igame)
		Infotmp.GameName = ss[i][1]
		Infotmp.Ip = ss[i][2]
		iport, _ := strconv.Atoi(ss[i][3])
		Infotmp.Port = uint32(iport)
		Infotmp.Type = ss[i][4]
		// G_GameInfoST[strconv.Itoa(int(Infotmp.GameId))] = Infotmp
		G_GameInfoST[Infotmp.GameName] = Infotmp
	}

	fmt.Println(G_GameInfoST)
	return true
}

// 读取游戏活动的数据
func ReadCsv_ConfigFile_ActivitiesInfoST_Fun() bool {
	// 获取数据，按照文件
	fileName := "ActivitiesInfo.csv"
	fileName = "./csv/" + fileName
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		return false
	}
	// 读取文件数据
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, _ := r2.ReadAll()
	sz := len(ss)
	// 循环取数据
	for i := 1; i < sz; i++ {

		Infotmp := new(Global_Define.StActivitiesInfocsv)
		Infotmp.ID = ss[i][0]
		Infotmp.LoginName = ss[i][1]
		Infotmp.PicUrl = ss[i][2]
		Infotmp.Type = ss[i][3]
		Infotmp.JiLV = ss[i][4]
		G_ActivitiesInfoST[Infotmp.ID] = Infotmp
	}

	fmt.Println("ActivitiesInfo.csv:", G_ActivitiesInfoST)
	return true
}

// 读取游戏配置文件
func ReadCsv_ConfigFile_GameConfigInfoST_Fun() bool {
	// 获取数据，按照文件
	fileName := "GameConfig.csv"
	fileName = "./csv/" + fileName
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		return false
	}
	// 读取文件数据
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, _ := r2.ReadAll()
	sz := len(ss)
	// 循环取数据
	for i := 1; i < sz; i++ {

		Infotmp := new(Global_Define.StGameConfigInfo)
		Infotmp.GameName = ss[i][0]
		Infotmp.GameID = ss[i][1]
		Infotmp.PlayerMax = ss[i][2]
		Infotmp.PlayerTimeOut = ss[i][3]
		Infotmp.GmaeTime = ss[i][4]
		Infotmp.PCTimeOut = ss[i][5]
		Infotmp.GetGameJiFen = ss[i][6]

		G_GameConfigInfoST[Infotmp.GameName] = Infotmp
	}

	fmt.Println(G_GameConfigInfoST)
	return true
}

// 获取卡牌的列表
func ReadCsv_ConfigFile_StCard2List_Fun() bool {
	// 获取数据，按照文件
	fileName := "puke.csv"
	fileName = "./csv/" + fileName
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		return false
	}
	// 读取文件数据
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, _ := r2.ReadAll()
	sz := len(ss)

	// 循环取数据
	for i := 1; i < sz; i++ {

		Infotmp := new(Global_Define.Card2InfoBase)
		Infotmp.Card2ID = ss[i][0]
		Infotmp.Card2Msg = ss[i][1]
		Infotmp.Card2GameName = ss[i][2]
		Infotmp.Card2GameID = ss[i][3]
		Infotmp.PicPath = ss[i][4]
		Global_Define.G_StCard2InfoBaseST[Infotmp.Card2ID] = Infotmp
	}
	fmt.Println(Global_Define.G_StCard2InfoBaseST)
	return true
}
