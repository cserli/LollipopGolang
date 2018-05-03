package Redis_DB

//import (
//	"encoding/json"
//	"fmt"
//	"glog-master"
//	"io/ioutil"
//	"net/http"
//	"zLollipopGolang/globalData"
//	"zLollipopGolang/playerstruct"

//	"code.google.com/p/go.net/websocket"
//	//"net/url"
//	"strconv"
//	//"sync"
//	"time"
//)

////var msuo sync.Mutex
////var mmlock sync.RWMutex

//// 结构体数据类型
//type Requestbody struct {
//	req string
//}

//// 在线玩家的数据的结构体
//type OnlineUser struct {
//	Connection *websocket.Conn
//}

//// json转化为map:数据的处理
//func (r *Requestbody) Json2map() (s map[string]interface{}, err error) {
//	var result map[string]interface{}
//	if err := json.Unmarshal([]byte(r.req), &result); err != nil {
//		glog.Info("Json2map:", err.Error())
//		return nil, err
//	}
//	return result, nil
//}

//////////////////////////////////////////////////////////////////////////////////
////
////                 Get_Redis_GameInfo // 全部是内存的数据保存的
////
////
//////////////////////////////////////////////////////////////////////////////////

//func Redis_Read_GameInfo() {

//	glog.Info("Entry Redis_Read_GameInfo")

//	return
//}

////玩家登陆数据的保存:直接数据的操作内存数据库
////hkey ：为 玩家的AccountID
////func Redis_Read_PlayerInfo(AccountID uint64) PlayerData.Plyerdata {

////	//	glog.Info("Entry Redis_Write_PlayerInfo")

////	//	tmpplayer := PlayerData.Plyerdata{}
////	//	glog.Info(AccountID)
////	//	// 写入玩家数据到内存数据库
////	//	data_byte, error := GetREDIS().Redis_Client.Hget(PlayerLoginKey, strconv.Itoa(int(AccountID)))
////	//	if error != nil {
////	//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error)
////	//		glog.Info("GetREDIS().Redis_Client.Hset: Set data 66666666", data_byte)
////	//		return tmpplayer
////	//	}
////	//	// 数据转换
////	//	var r Requestbody
////	//	// pPlayerTmp := new(PlayerData.Plyerdata)
////	//	r.req = string(data_byte)
////	//	glog.Info("++++++++++++++++++ ", string(data_byte))
////	//	var result PlayerData.Plyerdata
////	//	json.Unmarshal([]byte(r.req), &result)
////	//	tmpplayer = result
////		return tmpplayer
////}

//// 获取指定的玩家的签到的天数，内存素具保存的数据的结果；数据库需要做容错处理
//// hkey ： 为 玩家的accountid的信息
////func Redis_Read_PlayerSignInDayData(strkey string) int {
////	//	glog.Info("Entry Redis_Read_PlayerSignInDayData")
////	//	var count int
////	//	// 获取内存数据的数据信息
////	//	// 1 获取数据信息 内存的数据库的信息
////	//	data_byte, _ := GetREDIS().Redis_Client.Hget(PlayerSinginDayKey, strkey)
////	//	tmpstring := string(data_byte)
////	//	count, _ = strconv.Atoi(tmpstring)

////	//	return count
////}

//// 获取账单的数据信息：规则是按照时间最新的在最前面的数据的获取的方式
//// strkeytime ： 为是账单产生的时间
//// 返回数值为，一个map数据，这个数据是多个数据的组合的操作,最多数据为也数据
////func Redis_Read_HistoryPayInfoData(strkeytime string, iRoleUID int64) (PayInfo map[int]*Global_Define.StHistoryPayInfo) {
////	//	glog.Info("Entry Redis_Read_HistoryPayInfoData")

////	//	ftmpData, _ := strconv.Atoi(strkeytime)
////	//	glog.Info("strconv.Atoi(strkeytime):", ftmpData)
////	//	// 获取内存数据库的数据
////	//	bytesMap, _ := GetREDIS().Redis_Client.Zrevrange(HistoryPayInfoKey, 0, -1)
////	//	glog.Info("GetREDIS().Redis_Client.Zrevrange return:", bytesMap)
////	//	// 这里就需要对比时间，数据的操作，主要的数据的
////	//	// 1 循环取到数据。
////	//	var vcount = 0
////	//	var PayInfotmp Global_Define.StHistoryPayInfo
////	//	for member, score := range bytesMap {
////	//		// 如果时间戳大于或者等于member
////	//		if ftmpData >= member {

////	//			// 临时的 转换
////	//			json.Unmarshal(score, PayInfotmp) // score转换结构体
////	//			// 判断数据与自己有关不
////	//			if iRoleUID == PayInfotmp.SrcRoleUID || iRoleUID == PayInfotmp.DecRoleUID {
////	//				vcount++
////	//				json.Unmarshal(score, PayInfo[vcount]) // score转换结构体
////	//			}
////	//		}
////	//		// 退出循环,当达到20条数据的时候
////	//		if vcount >= 20 {
////	//			return PayInfo
////	//		}
////	//	}

////	//	return PayInfo
////}

//// 获取问答社区评论的的内存的数据的信息
//// strkeytime ： 为是账单产生的时间
//// 返回数值为，一个map数据，这个数据是多个数据的组合的操作,最多数据为也数据
//func Redis_Read_AnswerCommentInfoData(strkeytime string, iAnswerid int64) (CommentInfo map[string]*Global_Define.StCommentInfo) {

//	glog.Info("Entry Redis_Read_AnswerCommentInfoData")
//	CommentInfo = make(map[string]*Global_Define.StCommentInfo)

//	ftmpData, _ := strconv.Atoi(strkeytime)
//	glog.Info("strconv.Atoi(strkeytime):", ftmpData)
//	// 获取内存数据库的数据
//	bytesMap, _ := GetREDIS().Redis_Client.Zrevrange(AnswerCommentInfoKey, 0, -1)
//	glog.Info("GetREDIS().Redis_Client.Zrevrange return:", bytesMap)
//	// 这里就需要对比时间，数据的操作，主要的数据的
//	// 1 循环取到数据。
//	var vcount = 0
//	var CommentInfotmp *Global_Define.StCommentInfo
//	for member, score := range bytesMap {
//		// 如果时间戳大于或者等于member
//		CommentInfotmp = new(Global_Define.StCommentInfo)
//		var CommentIdtmp = ""
//		// 如果时间戳大于或者等于member
//		if ftmpData >= member {

//			// 临时的 转换
//			json.Unmarshal(score, CommentInfotmp) // score转换结构体
//			// 判断数据与自己有关不
//			// vcount++
//			glog.Info("GetREDIS().Redis_Client.Zrevrange string(score)", string(score))
//			//test
//			//json转map
//			var r Requestbody
//			r.req = string(score)
//			if req2map, err := r.Json2map(); err == nil {

//				//				s := fmt.Sprintf("a %s", "string")
//				//				fmt.Println(s)

//				CommentId := fmt.Sprintf("%s", req2map["CommentId"])
//				CommentAuthor := fmt.Sprintf("%s", req2map["CommentAuthor"])
//				CommentAuthorIcon := fmt.Sprintf("%s", req2map["CommentAuthorIcon"])
//				CommentAuthrID := fmt.Sprintf("%s", req2map["CommentAuthrID"])
//				CommentContent := fmt.Sprintf("%s", req2map["CommentContent"])
//				CommentTime := fmt.Sprintf("%s", req2map["CommentTime"])

//				// 转换的
//				CommentInfotmp.CommentId = CommentId
//				CommentIdtmp = CommentId
//				CommentInfotmp.CommentAuthor = CommentAuthor
//				CommentInfotmp.CommentAuthorIcon = CommentAuthorIcon
//				CommentInfotmp.CommentAuthrID = CommentAuthrID
//				CommentInfotmp.CommentContent = CommentContent
//				CommentInfotmp.CommentTime = CommentTime

//				glog.Info("------------------------------------------------")
//				glog.Info("CommentInfotmp.Answerid", CommentInfotmp.CommentId)
//				glog.Info("------------------------------------------------")

//				glog.Info("GetREDIS().Redis_Client.Zrevrange CommentInfotmp:", CommentInfotmp)
//			} else {
//				glog.Info("GetREDIS().Redis_Client.Zrevrange CommentInfotmp[vcount]:error")
//			}
//			// 评论ID 相等
//			if strconv.Itoa(int(iAnswerid)) == CommentIdtmp {
//				vcount++
//				CommentInfo[strconv.Itoa(vcount)] = CommentInfotmp
//				glog.Info("GetREDIS().Redis_Client.Zrevrange CommentInfo[vcount]:", CommentInfo[strconv.Itoa(vcount)])
//			}

//		}

//		// 退出循环,当达到20条数据的时候
//		if vcount >= 20 {
//			return CommentInfo
//		}
//	}

//	return CommentInfo
//}

//// 获取问答社区帖子列表的内存的数据的信息
//// strkeytime ： 为是账单产生的时间
//// 返回数值为，一个map数据，这个数据是多个数据的组合的操作,最多数据为也数据
//func Redis_Read_AnswerListInfoData(strkeytime string) (AnswerInfo map[string]*Global_Define.StAnswerInfo) {
//	glog.Info("Entry Redis_Read_AnswerListInfoData")
//	AnswerInfo = make(map[string]*Global_Define.StAnswerInfo)

//	ftmpData, _ := strconv.Atoi(strkeytime)
//	glog.Info("strconv.Atoi(strkeytime):", ftmpData)
//	// 获取内存数据库的数据
//	bytesMap, _ := GetREDIS().Redis_Client.Zrevrange(AnswerInfoKey, 0, -1)
//	glog.Info("GetREDIS().Redis_Client.Zrevrange return:", bytesMap)
//	// 这里就需要对比时间，数据的操作，主要的数据的
//	// 1 循环取到数据。
//	var vcount = 0
//	//  这里需要修改数据操作
//	//  时间上判断等于100的操作是不是刷新最新的
//	if ftmpData == 100 {

//		var AnswerInfotmp *Global_Define.StAnswerInfo
//		for member, score := range bytesMap {
//			AnswerInfotmp = new(Global_Define.StAnswerInfo)
//			// 如果时间戳大于或者等于member
//			if ftmpData >= member {

//				// 临时的 转换
//				json.Unmarshal(score, AnswerInfotmp) // score转换结构体
//				// 判断数据与自己有关不
//				vcount++
//				glog.Info("GetREDIS().Redis_Client.Zrevrange string(score)", string(score))
//				//test
//				//json转map
//				var r Requestbody
//				r.req = string(score)
//				if req2map, err := r.Json2map(); err == nil {

//					//				s := fmt.Sprintf("a %s", "string")
//					//				fmt.Println(s)

//					AuthorRoleUID := fmt.Sprintf("%s", req2map["AuthorRoleUID"])
//					AuthorName := fmt.Sprintf("%s", req2map["AuthorName"])
//					AuthorIcon := fmt.Sprintf("%s", req2map["AuthorIcon"])
//					AnswerContent := fmt.Sprintf("%s", req2map["AnswerContent"])
//					Instertime := fmt.Sprintf("%s", req2map["Instertime"])
//					Answerid := fmt.Sprintf("%s", req2map["Answerid"])

//					// 转换的
//					AnswerInfotmp.AuthorRoleUID = AuthorRoleUID
//					AnswerInfotmp.AuthorName = AuthorName
//					AnswerInfotmp.AuthorIcon = AuthorIcon
//					AnswerInfotmp.AnswerContent = AnswerContent
//					AnswerInfotmp.Instertime = Instertime
//					AnswerInfotmp.Answerid = Answerid

//					glog.Info("------------------------------------------------")
//					glog.Info("AnswerInfotmp.Answerid", AnswerInfotmp.Answerid)
//					glog.Info("------------------------------------------------")

//					glog.Info("GetREDIS().Redis_Client.Zrevrange AnswerInfotmp:", AnswerInfotmp)
//				} else {
//					glog.Info("GetREDIS().Redis_Client.Zrevrange AnswerInfo[vcount]:error")
//				}

//				AnswerInfo[strconv.Itoa(vcount)] = AnswerInfotmp
//				glog.Info("GetREDIS().Redis_Client.Zrevrange AnswerInfo[vcount]:", AnswerInfo[strconv.Itoa(vcount)])
//			}
//			// 退出循环,当达到20条数据的时候
//			if vcount >= 20 {
//				return AnswerInfo
//			}
//		}
//		return AnswerInfo
//	}
//	// ------------------------------------------------------------------------------
//	// 真正的时间戳的效果
//	var AnswerInfotmp *Global_Define.StAnswerInfo
//	for member, score := range bytesMap {
//		glog.Info("member：", member)
//		AnswerInfotmp = new(Global_Define.StAnswerInfo)
//		// 临时的 转换
//		json.Unmarshal(score, AnswerInfotmp) // score转换结构体
//		// 判断数据与自己有关不
//		vcount++
//		glog.Info("GetREDIS().Redis_Client.Zrevrange string(score)", string(score))
//		//test
//		//json转map
//		var r Requestbody
//		r.req = string(score)
//		if req2map, err := r.Json2map(); err == nil {
//			// 如果时间戳大于或者等于member
//			Instertimetmp := fmt.Sprintf("%s", req2map["Instertime"])
//			inttemp, _ := strconv.Atoi(Instertimetmp)
//			if ftmpData > inttemp {

//				//			// 临时的 转换
//				//			json.Unmarshal(score, AnswerInfotmp) // score转换结构体
//				//			// 判断数据与自己有关不
//				//			vcount++
//				//			glog.Info("GetREDIS().Redis_Client.Zrevrange string(score)", string(score))
//				//			//test
//				//			//json转map
//				//			var r Requestbody
//				//			r.req = string(score)
//				//			if req2map, err := r.Json2map(); err == nil { //======

//				//				s := fmt.Sprintf("a %s", "string")
//				//				fmt.Println(s)

//				AuthorRoleUID := fmt.Sprintf("%s", req2map["AuthorRoleUID"])
//				AuthorName := fmt.Sprintf("%s", req2map["AuthorName"])
//				AuthorIcon := fmt.Sprintf("%s", req2map["AuthorIcon"])
//				AnswerContent := fmt.Sprintf("%s", req2map["AnswerContent"])
//				Instertime := fmt.Sprintf("%s", req2map["Instertime"])
//				Answerid := fmt.Sprintf("%s", req2map["Answerid"])

//				// 转换的
//				AnswerInfotmp.AuthorRoleUID = AuthorRoleUID
//				AnswerInfotmp.AuthorName = AuthorName
//				AnswerInfotmp.AuthorIcon = AuthorIcon
//				AnswerInfotmp.AnswerContent = AnswerContent
//				AnswerInfotmp.Instertime = Instertime
//				AnswerInfotmp.Answerid = Answerid

//				glog.Info("------------------------------------------------")
//				glog.Info("AnswerInfotmp.Answerid", AnswerInfotmp.Answerid)
//				glog.Info("------------------------------------------------")

//				glog.Info("GetREDIS().Redis_Client.Zrevrange AnswerInfotmp:", AnswerInfotmp)
//			}

//			AnswerInfo[strconv.Itoa(vcount)] = AnswerInfotmp
//			glog.Info("GetREDIS().Redis_Client.Zrevrange AnswerInfo[vcount]:", AnswerInfo[strconv.Itoa(vcount)])
//			// 退出循环,当达到20条数据的时候
//			if vcount >= 20 {
//				return AnswerInfo
//			}

//		} else {
//			glog.Info("GetREDIS().Redis_Client.Zrevrange AnswerInfo[vcount]:error")
//		}
//		return AnswerInfo
//	}

//	return AnswerInfo
//}

//// 获取问答社区帖子列表的内存的数据的信息
//// strkeytime ： 为是账单产生的时间
//// 返回数值为，一个map数据，这个数据是多个数据的组合的操作,最多数据为也数据
//func Redis_Read_ArticleListInfoData(strkeytime string, iiCircleid int) (ArticleInfo map[string]*Global_Define.StArticleInfo) {
//	glog.Info("Entry Redis_Read_ArticleListInfoData")
//	ArticleInfo = make(map[string]*Global_Define.StArticleInfo)

//	ftmpData, _ := strconv.Atoi(strkeytime)
//	glog.Info("strconv.Atoi(strkeytime):", ftmpData)

//	var GoAnguageInfoKey string

//	// 首先判断是那个圈子的数据的保存
//	if iiCircleid == 100 {
//		GoAnguageInfoKey = "GoAnguage_Base" // GO 语言基础知识

//	} else if iiCircleid == 101 {
//		GoAnguageInfoKey = "GoAnguage_NetWork" // GO 语言网路知识

//	} else if iiCircleid == 102 {
//		GoAnguageInfoKey = "GoAnguage_DataBase" // GO 语言数据库知识

//	} else if iiCircleid == 103 {
//		GoAnguageInfoKey = "GoAnguage_MemoryDataBase" // GO 语言内存数据库知识

//	} else if iiCircleid == 104 {
//		GoAnguageInfoKey = "GoAnguage_Server" // GO 语言服务器开发

//	} else if iiCircleid == 105 {
//		GoAnguageInfoKey = "GoAnguage_OfficialRewardMK" // GO 语言官方模块悬赏，不一定是模块，也可以是工具的开发的悬赏

//	} else if iiCircleid == 106 {
//		GoAnguageInfoKey = "GoAnguage_OfficialRewardXM" // GO 语言官方项目悬赏，项目的进度的悬赏。
//	}

//	// 转换以下
//	if ftmpData == 100 {
//		// 获取时间，转化为基础的
//		// 获取内存数据库的数据
//		bytesMap, _ := GetREDIS().Redis_Client.Zrevrange(GoAnguageInfoKey, 0, -1)
//		glog.Info("GetREDIS().Redis_Client.Zrevrange return:", bytesMap)
//		// 这里就需要对比时间，数据的操作，主要的数据的
//		// 1 循环取到数据。
//		var vcount = 0
//		var ArticleInfotmp *Global_Define.StArticleInfo
//		for member, score := range bytesMap {
//			ArticleInfotmp = new(Global_Define.StArticleInfo)
//			// 如果时间戳大于或者等于member
//			if ftmpData >= member {

//				// 临时的 转换
//				json.Unmarshal(score, ArticleInfotmp) // score转换结构体
//				// 判断数据与自己有关不
//				vcount++
//				glog.Info("GetREDIS().Redis_Client.Zrevrange string(score)", string(score))
//				//test
//				//json转map
//				var r Requestbody
//				r.req = string(score)
//				if req2map, err := r.Json2map(); err == nil {

//					//				s := fmt.Sprintf("a %s", "string")
//					//				fmt.Println(s)

//					ArticleId := fmt.Sprintf("%s", req2map["ArticleId"])
//					ArticleName := fmt.Sprintf("%s", req2map["ArticleName"])
//					ArticleAuthor := fmt.Sprintf("%s", req2map["ArticleAuthor"])
//					ArticleAuthorIcon := fmt.Sprintf("%s", req2map["ArticleAuthorIcon"])
//					ArticleContent := fmt.Sprintf("%s", req2map["ArticleContent"])
//					ArticleNewComment := fmt.Sprintf("%s", req2map["ArticleNewComment"])
//					ArticleIssueTime := fmt.Sprintf("%s", req2map["ArticleIssueTime"])
//					ICircleId := fmt.Sprintf("%s", req2map["ICircleId"])
//					IArticleAuthorUID := fmt.Sprintf("%s", req2map["IArticleAuthorUID"])

//					// 转换的
//					ArticleInfotmp.ArticleId = ArticleId
//					ArticleInfotmp.ArticleName = ArticleName
//					ArticleInfotmp.ArticleAuthor = ArticleAuthor
//					ArticleInfotmp.ArticleAuthorIcon = ArticleAuthorIcon
//					ArticleInfotmp.ArticleContent = ArticleContent
//					ArticleInfotmp.ArticleNewComment = ArticleNewComment
//					ArticleInfotmp.ArticleIssueTime = ArticleIssueTime
//					ArticleInfotmp.ICircleId = ICircleId
//					ArticleInfotmp.IArticleAuthorUID = IArticleAuthorUID

//					glog.Info("------------------------------------------------")
//					glog.Info("ArticleInfotmp.ICircleId", ArticleInfotmp.ICircleId)
//					glog.Info("ArticleInfotmp.ArticleId", ArticleInfotmp.ArticleId)
//					glog.Info("------------------------------------------------")

//					glog.Info("GetREDIS().Redis_Client.Zrevrange AnswerInfotmp:", ArticleInfotmp)
//				} else {
//					glog.Info("GetREDIS().Redis_Client.Zrevrange ArticleInfotmp[vcount]:error")
//				}

//				ArticleInfo[strconv.Itoa(vcount)] = ArticleInfotmp
//				glog.Info("GetREDIS().Redis_Client.Zrevrange ArticleInfo[vcount]:", ArticleInfo[strconv.Itoa(vcount)])
//			}
//			// 退出循环,当达到5条数据的时候
//			if vcount >= 5 {
//				return ArticleInfo
//			}
//		}
//		return ArticleInfo
//	}

//	// 获取内存数据库的数据
//	bytesMap, _ := GetREDIS().Redis_Client.Zrevrange(GoAnguageInfoKey, 0, -1)
//	glog.Info("GetREDIS().Redis_Client.Zrevrange return:", bytesMap)
//	// 这里就需要对比时间，数据的操作，主要的数据的
//	// 1 循环取到数据。
//	var vcount = 0
//	var ArticleInfotmp *Global_Define.StArticleInfo
//	for member, score := range bytesMap {
//		glog.Info("member:", member)
//		ArticleInfotmp = new(Global_Define.StArticleInfo)
//		// 临时的 转换
//		json.Unmarshal(score, ArticleInfotmp) // score转换结构体
//		glog.Info("GetREDIS().Redis_Client.Zrevrange string(score)", string(score))
//		//test
//		//json转map
//		var r Requestbody
//		r.req = string(score)
//		if req2map, err := r.Json2map(); err == nil {

//			ArticleIssueTimetemp := fmt.Sprintf("%s", req2map["ArticleIssueTime"])
//			inttemp, _ := strconv.Atoi(ArticleIssueTimetemp)
//			// 如果时间戳大于或者等于member
//			if ftmpData > inttemp {
//				// 判断数据与自己有关不
//				vcount++
//				//				s := fmt.Sprintf("a %s", "string")
//				//				fmt.Println(s)

//				ArticleId := fmt.Sprintf("%s", req2map["ArticleId"])
//				ArticleName := fmt.Sprintf("%s", req2map["ArticleName"])
//				ArticleAuthor := fmt.Sprintf("%s", req2map["ArticleAuthor"])
//				ArticleAuthorIcon := fmt.Sprintf("%s", req2map["ArticleAuthorIcon"])
//				ArticleContent := fmt.Sprintf("%s", req2map["ArticleContent"])
//				ArticleNewComment := fmt.Sprintf("%s", req2map["ArticleNewComment"])
//				ArticleIssueTime := fmt.Sprintf("%s", req2map["ArticleIssueTime"])
//				ICircleId := fmt.Sprintf("%s", req2map["ICircleId"])
//				IArticleAuthorUID := fmt.Sprintf("%s", req2map["IArticleAuthorUID"])

//				// 转换的
//				ArticleInfotmp.ArticleId = ArticleId
//				ArticleInfotmp.ArticleName = ArticleName
//				ArticleInfotmp.ArticleAuthor = ArticleAuthor
//				ArticleInfotmp.ArticleAuthorIcon = ArticleAuthorIcon
//				ArticleInfotmp.ArticleContent = ArticleContent
//				ArticleInfotmp.ArticleNewComment = ArticleNewComment
//				ArticleInfotmp.ArticleIssueTime = ArticleIssueTime
//				ArticleInfotmp.ICircleId = ICircleId
//				ArticleInfotmp.IArticleAuthorUID = IArticleAuthorUID

//				glog.Info("------------------------------------------------")
//				glog.Info("ArticleInfotmp.ICircleId", ArticleInfotmp.ICircleId)
//				glog.Info("ArticleInfotmp.ArticleId", ArticleInfotmp.ArticleId)
//				glog.Info("------------------------------------------------")

//				glog.Info("GetREDIS().Redis_Client.Zrevrange AnswerInfotmp:", ArticleInfotmp)

//				ArticleInfo[strconv.Itoa(vcount)] = ArticleInfotmp
//				glog.Info("GetREDIS().Redis_Client.Zrevrange ArticleInfo[vcount]:", ArticleInfo[strconv.Itoa(vcount)])
//			}

//		} else {
//			glog.Info("GetREDIS().Redis_Client.Zrevrange ArticleInfotmp[vcount]:error")
//		}
//		// 退出循环,当达到5条数据的时候
//		if vcount >= 5 {
//			return ArticleInfo
//		}
//	}

//	return ArticleInfo
//}

//// 获取问答社区评论的的内存的数据的信息
//// strkeytime ： 为是账单产生的时间
//// 返回数值为，一个map数据，这个数据是多个数据的组合的操作,最多数据为也数据
//func Redis_Read_ArticleCommentInfoData(strkeytime string, iAnswerid int64, icircleid int) (CommentInfo map[string]*Global_Define.StCommentInfo) {
//	glog.Info("Entry Redis_Read_ArticleCommentInfoData")
//	CommentInfo = make(map[string]*Global_Define.StCommentInfo)

//	ftmpData, _ := strconv.Atoi(strkeytime)
//	glog.Info("strconv.Atoi(strkeytime):", ftmpData)

//	var GoAnguageInfoKey string

//	// 首先判断是那个圈子的数据的保存
//	if strconv.Itoa(icircleid) == "100" {
//		GoAnguageInfoKey = "GoAnguage_Base" // GO 语言基础知识

//	} else if strconv.Itoa(icircleid) == "101" {
//		GoAnguageInfoKey = "GoAnguage_NetWork" // GO 语言网路知识

//	} else if strconv.Itoa(icircleid) == "102" {
//		GoAnguageInfoKey = "GoAnguage_DataBase" // GO 语言数据库知识

//	} else if strconv.Itoa(icircleid) == "103" {
//		GoAnguageInfoKey = "GoAnguage_MemoryDataBase" // GO 语言内存数据库知识

//	} else if strconv.Itoa(icircleid) == "104" {
//		GoAnguageInfoKey = "GoAnguage_Server" // GO 语言服务器开发

//	} else if strconv.Itoa(icircleid) == "105" {
//		GoAnguageInfoKey = "GoAnguage_OfficialRewardMK" // GO 语言官方模块悬赏，不一定是模块，也可以是工具的开发的悬赏

//	} else if strconv.Itoa(icircleid) == "106" {
//		GoAnguageInfoKey = "GoAnguage_OfficialRewardXM" // GO 语言官方项目悬赏，项目的进度的悬赏。
//	}

//	// 获取内存数据库的数据
//	bytesMap, _ := GetREDIS().Redis_Client.Zrevrange(GoAnguageInfoKey, 0, -1)
//	glog.Info("GetREDIS().Redis_Client.Zrevrange return:", bytesMap)
//	// 这里就需要对比时间，数据的操作，主要的数据的
//	// 1 循环取到数据。
//	var vcount = 0
//	var CommentInfotmp *Global_Define.StCommentInfo
//	for member, score := range bytesMap {
//		// 如果时间戳大于或者等于member
//		CommentInfotmp = new(Global_Define.StCommentInfo)
//		var CommentIdtmp = ""
//		// 如果时间戳大于或者等于member
//		if ftmpData >= member {

//			// 临时的 转换
//			json.Unmarshal(score, CommentInfotmp) // score转换结构体
//			// 判断数据与自己有关不
//			// vcount++
//			glog.Info("GetREDIS().Redis_Client.Zrevrange string(score)", string(score))
//			//test
//			//json转map
//			var r Requestbody
//			r.req = string(score)
//			if req2map, err := r.Json2map(); err == nil {

//				//				s := fmt.Sprintf("a %s", "string")
//				//				fmt.Println(s)

//				CommentId := fmt.Sprintf("%s", req2map["CommentId"])
//				CommentAuthor := fmt.Sprintf("%s", req2map["CommentAuthor"])
//				CommentAuthorIcon := fmt.Sprintf("%s", req2map["CommentAuthorIcon"])
//				CommentAuthrID := fmt.Sprintf("%s", req2map["CommentAuthrID"])
//				CommentContent := fmt.Sprintf("%s", req2map["CommentContent"])
//				CommentTime := fmt.Sprintf("%s", req2map["CommentTime"])

//				// 转换的
//				CommentInfotmp.CommentId = CommentId
//				CommentIdtmp = CommentId
//				CommentInfotmp.CommentAuthor = CommentAuthor
//				CommentInfotmp.CommentAuthorIcon = CommentAuthorIcon
//				CommentInfotmp.CommentAuthrID = CommentAuthrID
//				CommentInfotmp.CommentContent = CommentContent
//				CommentInfotmp.CommentTime = CommentTime

//				glog.Info("------------------------------------------------")
//				glog.Info("CommentInfotmp.Answerid", CommentInfotmp.CommentId)
//				glog.Info("------------------------------------------------")

//				glog.Info("GetREDIS().Redis_Client.Zrevrange CommentInfotmp:", CommentInfotmp)
//			} else {
//				glog.Info("GetREDIS().Redis_Client.Zrevrange CommentInfotmp[vcount]:error")
//			}
//			// 评论ID 相等
//			if strconv.Itoa(int(iAnswerid)) == CommentIdtmp {
//				vcount++
//				CommentInfo[strconv.Itoa(vcount)] = CommentInfotmp
//				glog.Info("GetREDIS().Redis_Client.Zrevrange CommentInfo[vcount]:", CommentInfo[strconv.Itoa(vcount)])
//			}

//		}

//		// 退出循环,当达到20条数据的时候
//		if vcount >= 20 {
//			return CommentInfo
//		}
//	}

//	return CommentInfo
//}

////==============================================================================
//// 玩家点击报名处理，有报名成功的加1
//// hkey ： 为商家的信息： 二2key   为官方的活动的数据
//func Redis_Read_YaoYiYaoBaoMingData(strkey string) int {
//	glog.Info("Entry Redis_Read_YaoYiYaoBaoMingData")
//	var count int = 0
//	// 获取内存数据的数据信息
//	// 1 获取数据信息 内存的数据库的信息
//	data_byte, _ := GetREDIS().Redis_Client.Hget(G_BaoMingKey, strkey)
//	tmpstring := string(data_byte)
//	count, _ = strconv.Atoi(tmpstring)

//	return count
//}

//// 玩家摇一摇处理，每摇一次加1
//// hkey ： 为商家的信息+现场名称： 二2key   OpenID
//func Redis_Read_YaoYiYaoJiShuData(strkey string) int {
//	glog.Info("Entry Redis_Read_YaoYiYaoJiShuData")
//	var count = 0
//	// 获取内存数据的数据信息
//	// 1 获取数据信息 内存的数据库的信息
//	data_byte, _ := GetREDIS().Redis_Client.Hget(G_DataSaveDongKey, strkey)
//	tmpstring := string(data_byte)
//	count, _ = strconv.Atoi(tmpstring)

//	return count
//}

//// 定时器会调用，当有3名用户到达终点的时候发，倒计时的数据的消息
//// 玩家摇一摇到达终点，每到一位加1
//// hkey ： GameOver+商家名字+场次： 二2key   GameOver
//func Redis_Read_YaoYiYaoGameOver3Data(strkey string) int {
//	glog.Info("Entry Redis_Read_YaoYiYaoGameOver3Data")
//	var count int = 0
//	// 获取内存数据的数据信息
//	// 1 获取数据信息 内存的数据库的信息
//	data_byte, _ := GetREDIS().Redis_Client.Hget(G_TimerGameOverKey, strkey)
//	tmpstring := string(data_byte)
//	count, _ = strconv.Atoi(tmpstring)

//	return count
//}

//// 摇一摇游戏排行榜的设计，获取数据，倒序获取数据
//// strkey ： 摇动的步数
//// 返回数值为，openid
//func Redis_Read_YaoYiYaoPaiHangBangInfoData() (MapOpenId map[string]*Global_Define.StPaiHangBang) {
//	glog.Info("Entry Redis_Read_YaoYiYaoPaiHangBangInfoData")
//	MapOpenId = make(map[string]*Global_Define.StPaiHangBang)

//	// 获取内存数据库的数据
//	glog.Info("GetREDIS().Redis_Client 是否为空:", GetREDIS().Redis_Client)
//	// 重新链接存储结构
//	if GetREDIS().Redis_Client == nil {
//		Redis_ConnFun()
//	}
//	// bytesMap, _ := GetREDIS().Redis_Client.Zrevrange(keys, 0, -1)
//	bytesMap, _ := GetREDIS().Redis_Client.Zrevrange(G_PaiHangKey, 0, -1)
//	glog.Info("GetREDIS().Redis_Client.Zrevrange return:", bytesMap)
//	// 1 循环取到数据。
//	var vcount = 0
//	//
//	var CommentInfotmp *Global_Define.StPaiHangBang
//	for _, score := range bytesMap {
//		// 如果时间戳大于或者等于member
//		CommentInfotmp = new(Global_Define.StPaiHangBang)
//		// 如果时间戳大于或者等于member
//		if 1 == 1 {

//			// 临时的 转换
//			json.Unmarshal(score, CommentInfotmp) // score转换结构体
//			// 判断数据与自己有关不
//			vcount++
//			glog.Info("GetREDIS().Redis_Client.Zrevrange string(score)", string(score))
//			//test
//			//json转map
//			var r Requestbody
//			r.req = string(score)
//			if req2map, err := r.Json2map(); err == nil {

//				//				s := fmt.Sprintf("a %s", "string")
//				//				fmt.Println(s)

//				OpenID := fmt.Sprintf("%s", req2map["OpenID"])
//				//				PaiHang := fmt.Sprintf("%s", req2map["PaiHang"])
//				//				YaoCiShu := fmt.Sprintf("%s", req2map["YaoCiShu"])

//				// 转换的
//				CommentInfotmp.OpenID = OpenID
//				//				CommentInfotmp.PaiHang = PaiHang
//				//				CommentInfotmp.YaoCiShu = YaoCiShu

//				glog.Info("------------------------------------------------")
//				glog.Info("OpenID", CommentInfotmp.OpenID)
//				glog.Info("------------------------------------------------")

//				glog.Info("GetREDIS().Redis_Client.Zrevrange MapOpenId:", CommentInfotmp)
//			} else {
//				glog.Info("GetREDIS().Redis_Client.Zrevrange MapOpenId[vcount]:error")
//			}
//			// 保存数据
//			MapOpenId[strconv.Itoa(vcount)] = CommentInfotmp
//			glog.Info("GetREDIS().Redis_Client.Zrevrange MapOpenId[vcount]:", MapOpenId[strconv.Itoa(vcount)])

//		}

//		// return MapOpenId
//	}

//	return MapOpenId
//}

////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////抽奖///////////////////////////////////////
////// 抽奖CD时间的保存
////func Redis_Read_ChouJiangCDTimeInfoData(stropenid string, StrMd5 string) (string, bool) {

////	glog.Info("Entry Redis_Read_ChouJiangCDTimeInfoData!!!")
////	time.Sleep(10)
////	msuo.Lock()
////	defer msuo.Unlock()
////	data_byte, err := GetREDIS().Redis_Client.Hget(CDTime+"yiy"+StrMd5, stropenid)
////	if err != nil {
////		return "", false
////	}
////	tmpstring := string(data_byte)
////	return tmpstring, true

////}

//////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////抽奖///////////////////////////////////////
//// 抽奖CD时间的保存
//func Redis_Read_ChouJiangCDTimeInfoData(stropenid string, StrMd5 string) (string, bool) {

//	glog.Info("Entry Redis_Read_ChouJiangCDTimeInfoData!!!")
//	time.Sleep(10)
//	//	msuo.Lock()
//	//	defer msuo.Unlock()
//	data_byte, err := GetREDIS().Redis_Client.Hget(CDTime+"yiy"+StrMd5, stropenid)
//	if err != nil {
//		glog.Info("Entry Redis_Read_ChouJiangCDTimeInfoData!!!", err)
//		return "", false
//	}
//	tmpstring := string(data_byte)
//	glog.Info("=============================CDtime==========================tmpstring：" + tmpstring)
//	return tmpstring, true

//}

///////////////////////////////////////更新二维码的数据///////////////////////////////////////
//func Redis_Read_EWMResData(StrMd5 string) bool {

//	glog.Info("Entry Redis_Read_EWMResData!!!")
//	var count int = 0
//	// 获取内存数据的数据信息
//	// 1 获取数据信息 内存的数据库的信息
//	data_byte, err := GetREDIS().Redis_Client.Hget("EWN_Update"+StrMd5, StrMd5)
//	if err != nil {
//		return true
//	}
//	tmpstring := string(data_byte)
//	glog.Info("=========================Redis_Read_EWMResData==============================tmpstring" + tmpstring)
//	count, _ = strconv.Atoi(tmpstring)
//	if count == 0 {
//		return false
//	}
//	return true
//}

///////////////////////////////////////更新二维码的数据///////////////////////////////////////
//func Redis_Read_EWMResDataPort(StrMd5 string) bool {

//	glog.Info("Entry Redis_Read_EWMResDataPort!!!")
//	var count int = 0
//	// 获取内存数据的数据信息
//	// 1 获取数据信息 内存的数据库的信息
//	data_byte, err := GetREDIS().Redis_Client.Hget("EWN_Update"+StrMd5, StrMd5)
//	if err != nil {
//		return true
//	}
//	tmpstring := string(data_byte)
//	glog.Info("=========================Redis_Read_EWMResData==============================tmpstring" + tmpstring)
//	count, _ = strconv.Atoi(tmpstring)
//	if count == 0 {
//		return false
//	}
//	return true
//}

//// 获取数据，用户现场信息
//// hkey:MD5(登陆名字|游戏名字|现场名字)
//func Redis_Read_InfoData_About_GameLogoOrEwm(hkey string, itype uint32) string {
//	//glog.Info("Entry Redis_Read_InfoData_About_GameLogoOrEwm!")
//	if itype == 1 {
//		M_Log := "log"
//		logtime, error := GetREDIS().Redis_Client.Hget(M_Log, hkey)
//		if error != nil {
//			glog.Info("GetREDIS().Redis_Client.Hget: Set data ", error.Error())
//			return ""
//		}
//		return string(logtime)

//	} else if itype == 2 {
//		M_Ewm := "ewm"
//		ewmtime, error := GetREDIS().Redis_Client.Hget(M_Ewm, hkey)
//		if error != nil {
//			glog.Info("GetREDIS().Redis_Client.Hget: Set data ", error.Error())
//			return ""
//		}
//		return string(ewmtime)
//	}
//	return ""
//}

///////////////////////////////////////更新水军的数据///////////////////////////////////////
//func Redis_Read_ShuiJunData(strOpenID string, strLoginName string) bool {

//	glog.Info("Entry Redis_Read_ShuiJunData!!!")

//	StrFirmUser := Global_Define.MD5ToStringFromString(strLoginName)
//	data_byte, _ := GetREDIS().Redis_Client.Hget("ShuiJun|"+StrFirmUser, strOpenID)
//	tmpstring := string(data_byte)
//	if strOpenID == tmpstring {
//		return true
//	}
//	return false
//}

////// CD次数
////func Redis_Read_ChouJiangCDCiShuInfoData(stropenid string, StrMd5 string) (int, bool) {

////	glog.Info("Entry Redis_Read_ChouJiangCDCiShuInfoData!!!")
////	time.Sleep(10)
////	msuo.Lock()
////	defer msuo.Unlock()
////	var count int = 0
////	// 1 获取数据信息 内存的数据库的信息
////	data_byte, err := GetREDIS().Redis_Client.Hget(CDTime+"yiycishu"+StrMd5, stropenid)
////	if err != nil {
////		return 0, false
////	}
////	tmpstring := string(data_byte)
////	count, _ = strconv.Atoi(tmpstring)
////	glog.Info("=======================================================count：" + strconv.Itoa(int(count)))
////	glog.Info("=======================================================tmpstring：" + tmpstring)
////	if count > 1000 {
////		count = 0
////	}
////	return count, true
////}

//// CD次数
//func Redis_Read_ChouJiangCDCiShuInfoData(stropenid string) (int, bool) {

//	glog.Info("Entry Redis_Read_ChouJiangCDCiShuInfoData!!!")
//	time.Sleep(10)
//	//	msuo.Lock()
//	//	defer msuo.Unlock()
//	var count int = 0
//	data_byte, err := GetREDIS().Redis_Client.Get(stropenid)
//	if err != nil {
//		glog.Info("Entry Redis_Read_ChouJiangCDCiShuInfoData!!!", err)
//		return 0, false
//	}
//	tmpstring := string(data_byte)
//	count, _ = strconv.Atoi(tmpstring)
//	glog.Info("===========================CDCiShu============================count：" + strconv.Itoa(int(count)))
//	if count > 1000 {
//		count = 0
//	}
//	return count, true

//}

///////////////////////////////////////读取修改配置时间///////////////////////////////////////
//func Redis_Read_ConfigData(StrMD5 string) string {

//	glog.Info("Entry Redis_Read_ConfigData!!!")
//	// 写入玩家数据到内存数据库
//	data_byte, error1 := GetREDIS().Redis_Client.Hget("GameConfig", StrMD5)
//	if error1 != nil {
//		glog.Info("GetREDIS().Redis_Client.Hset: Set data ", error1.Error())
//		return ""
//	}
//	tmpstring := string(data_byte)
//	return tmpstring
//}

///////////////////////////////////////积分///////////////////////////////////////
//// 游戏积分
//func Redis_Read_Player_JiFenDatabak(stropenid string) int {

//	glog.Info("Entry Redis_Read_Player_JiFenData!!!")
//	if len(stropenid) == 0 {
//		return 0
//	}
//	// 读写所
//	//	mmlock.RLock()
//	//	defer mmlock.RUnlock()
//	var strOPenID = ""
//	strsplit := Strings_Split(stropenid, "|")
//	for i := 0; i < len(strsplit); i++ {
//		if i == 0 {
//			strOPenID = strsplit[i]
//		}
//	}
//	//////////////////////
//	var count int = 0
//	// 获取内存数据的数据信息
//	// 1 获取数据信息 内存的数据库的信息
//	data_byte, _ := GetREDIS().Redis_Client.Get(strOPenID)
//	tmpstring := string(data_byte)
//	glog.Info("=========================-------------------------==============================tmpstring" + tmpstring)
//	count, _ = strconv.Atoi(tmpstring)
//	return count
//}

/////////////////////////////////////////积分///////////////////////////////////////
////// 游戏积分
////func Redis_Read_Player_JiFenData(stropenid string) int {

////	glog.Info("Entry Redis_Read_Player_JiFenData!!!")
////	if len(stropenid) == 0 {
////		return 0
////	}
////	// 读写所
////	mmlock.RLock()
////	defer mmlock.RUnlock()
////	var strOPenID = ""
////	strsplit := Strings_Split(stropenid, "|")
////	for i := 0; i < len(strsplit); i++ {
////		if i == 0 {
////			strOPenID = strsplit[i]
////		}
////	}
////	//////////////////////
////	var count int = 0
////	// 获取内存数据的数据信息
////	// 1 获取数据信息 内存的数据库的信息
////	data_byte, _ := GetREDIS().Redis_Client.Hget("JiFenTJ", strOPenID)
////	tmpstring := string(data_byte)
////	glog.Info("=========================-------------------------==============================tmpstring" + tmpstring)
////	count, _ = strconv.Atoi(tmpstring)
////	return count
////}

///////////////////////////////////////积分///////////////////////////////////////
//// 游戏积分
//func Redis_Read_Player_JiFenData(stropenid string) (int, int) {

//	glog.Info("Entry Redis_Read_Player_JiFenData!!!")
//	if len(stropenid) == 0 {
//		return 0, -1
//	}
//	// 读写所
//	//	mmlock.RLock()
//	//	defer mmlock.RUnlock()
//	var strOPenID = ""
//	strsplit := Strings_Split(stropenid, "|")
//	for i := 0; i < len(strsplit); i++ {
//		if i == 0 {
//			strOPenID = strsplit[i]
//		}
//	}
//	strcount, iret := Get(strOPenID)
//	count, _ := strconv.Atoi(strcount)
//	return count, iret
//}

//var client = &http.Client{}

//func Get(openid string) (string, int) {

//	request, _ := http.NewRequest("GET", "http://127.0.0.1:9092/?openid="+openid, nil)
//	response, _ := client.Do(request)
//	defer response.Body.Close()
//	if response.StatusCode == 200 {
//		str, _ := ioutil.ReadAll(response.Body)
//		bodystr := string(str)
//		fmt.Println(bodystr)
//		return bodystr, 1
//	}
//	//	//向服务端发送get请求
//	//	if Log_Eio.BTest == true {
//	//		request, _ := http.NewRequest("GET", "http://112.74.201.179:9092/?openid="+openid, nil)
//	//		response, _ := client.Do(request)
//	//		defer response.Body.Close()
//	//		if response.StatusCode == 200 {
//	//			str, _ := ioutil.ReadAll(response.Body)
//	//			bodystr := string(str)
//	//			fmt.Println(bodystr)
//	//			return bodystr, 1
//	//		}
//	//	} else {
//	//		request, _ := http.NewRequest("GET", "http://120.25.83.66:9092/?openid="+openid, nil)
//	//		response, _ := client.Do(request)
//	//		defer response.Body.Close()
//	//		if response.StatusCode == 200 {
//	//			str, _ := ioutil.ReadAll(response.Body)
//	//			bodystr := string(str)
//	//			fmt.Println(bodystr)
//	//			return bodystr, 1
//	//		}
//	//	}
//	glog.Info("Entry Redis_Read_Player_JiFenData  http 9092 error!!!")
//	return "", 0
//}

//// 查找玩家结构数据
//func Redis_Read_PlayerSt_InfoData(stropenid string) (PlayST *PlayerData.YLGame_Plyerdata) {

//	glog.Info("Entry Redis_Read_PlayerSt_InfoData!!!")
//	if len(stropenid) == 0 {
//		glog.Error("Entry Redis_Read_PlayerSt_InfoData!!!,stropenid", stropenid)
//		return PlayST
//	}
//	data_byte, err := GetREDIS().Redis_Client.Hget("PlayerST", stropenid)
//	if err != nil {
//		glog.Error("Entry Redis_Read_PlayerSt_InfoData!!!", err)
//		return PlayST
//	}
//	if len(string(data_byte)) == 0 {
//		return PlayST
//	}
//	stb := &PlayerData.YLGame_Plyerdata{}
//	json.Unmarshal(data_byte, stb)
//	glog.Info("==========================="+stropenid+"============================stb", stb)

//	return stb

//}

//// 获取玩家的道具的结构
//func Redis_Read_GetCardInfo_Of_Game(stropenid string) string {

//	glog.Info("Entry Redis_Read_GetCardInfo_Of_Game!!!")
//	if len(stropenid) == 0 {
//		glog.Error("Entry Redis_Read_GetCardInfo_Of_Game!!!,stropenid", stropenid)
//		return ""
//	}
//	data_byte, err := GetREDIS().Redis_Client.Hget("Player_GetItem", stropenid)
//	if err != nil {
//		glog.Error("Entry Redis_Read_PlayerSt_InfoData!!!", err)
//		return ""
//	}
//	if len(string(data_byte)) == 0 {
//		return ""
//	}

//	glog.Info("==========================="+stropenid+"============================data_byte", string(data_byte))

//	return string(data_byte)

//}
