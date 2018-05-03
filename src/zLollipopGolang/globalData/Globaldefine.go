package Global_Define

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"

	"code.google.com/p/mahonia"
)

//清除卡卷结构
var G_IsUpdateDB map[string]bool                          // 是否定时更新 key 是MD5（登录名+游戏名字+现场的名字） +1 是更新奖品  +2 更新产品库
var G_StWeiXinDatatmpbak map[string]*StWeiXinUserInfobak  // 微信的用户的结构信息
var G_StWeiXinDatatmp map[string]*StWeiXinUserInfo        // 微信的用户的结构信息
var G_StWeiXinDatatmpPaihang map[string]*StWeiXinUserInfo // 微信的用户的结构信息

// 活动数据
type StActivitiesInfoBase struct {
	Id          string // ID
	AccountName string // 商家的名字
	Data        string
	Time        string
}

// 微信用户的结构拓展
type StWeiXinUserInfobak struct {
	ID            string
	Openid        string
	Nickname      string
	Sex           uint32
	Language      string
	City          string
	Province      string
	Country       string
	Headimgurl    string
	Privilege     string
	IdentifykeySJ string
	IdentifykeyHD string
	Createtime    string
}

// 新版活动界面
//// 活动数据
//type StActivitiesInfo struct {
//	Id            string // ID
//	AccountName   string // 商家的名字
//	Pic           string // 图片地址
//	Type          string // 发奖类型
//	JiLv          string // 等奖几率
//	Name          string // 图片名字
//	DataPicType   string // 图片类型jpg  ，png
//	DataPicBase64 string // 图片数据
//}

// 游戏活动
type StActivitiesInfo struct {
	ID        string // 序号
	LoginName string // 商家的登陆名字
	PicUrl    string // 图片地址
	Type      string // 1：全部获奖；2：非quanbu
	JiLV      string
}

// 积分的结构
type StJiFenLogInfo struct {
	FirmsName  string // 商家的名字
	XCName     string // 现场的名字
	GameName   string // 游戏的名字
	OpenID     string // 谁获取了积分
	IJiFen     uint32 // 获得的积分
	Instertime string // 获取积分的时间
}

// 商家游戏现场的数据的结构
type StXianChangInfo struct {
	ID             uint32
	SJName         string
	XianChangName  string // 现场名字
	GameID         uint32 // 游戏的ID
	GameName       string // 游戏的名字
	GameTime       string // 游戏的结束时间
	InsertTime     string // 插入的时间
	FirmsLogo      string // 商家的LOGO
	FirmsErCode    string // 商家的二维码
	FirmsName      string // 商家的名字
	FirmsAward     string // 商品列表
	FirmsPublicize string // 宣传语列表
	//FirmsAward     map[string]*StAwardInfo     // 商品列表
	//FirmsPublicize map[string]*StPublicizeInfo // 宣传语列表
	ResPath   string // 生成二维码的URL
	IpAndPort string //ip 端口
}

// 宣传语的结构
type StPublicizeInfo struct {
	Name   string
	Lev    string
	StrMsg string
	BMark  bool // false :没有被使用   true： 正在使用
}

// 奖品的结构
type StAwardInfo struct {
	Name     string
	Lev      string
	StrMsg   string
	AwardNum uint32
	BMark    bool // false :没有被使用   true： 正在使用
}

// 寻宝的结构的信息
type StXunBao struct {
	ItemId   string //奖品的ID
	ItemIcon string //奖品的Icon
	ItemName string //奖品的Name
	ItemMsg  string //奖品的描述
}

// 抽奖的结构的信息
type StChouJiang struct {
	OpenID string // 用户的唯一ID
	ItemID string // 奖品的ID
	StrMsg string // 描述信息
}

/////////////////////////////////////////////////////////////////////////////////
// 排行榜结构体
type StPaiHangBang struct {
	OpenID string // 用户的唯一ID
	//	PaiHang  string // 实时排行
	//	YaoCiShu string // 摇的次数
}

// 排行榜结构体
type StPaiHangBangP1C struct {
	OpenID   string // 用户的唯一ID
	PaiHang  string // 实时排行
	YaoCiShu string // 摇的次数
}

// 微信用户的结构,数据库的保存的数据
type StWeiXinUserInfo struct {
	Openid        string
	Nickname      string
	Sex           uint32
	Language      string
	City          string
	Province      string
	Country       string
	Headimgurl    string
	Privilege     string
	IdentifykeySJ string
	IdentifykeyHD string
}

// 玩家报名的结构
type StPlayerBaoMingInfo struct {
	PlayerUID              uint32 // 用户的唯一的UID
	PlayerWeiXinName       string // 现场用户报名的 微信的名字
	PlayerWeiXinHeadPicUrl string // 现场用户报名的 微信的头像地址
	PlayerPaiMingPos       uint8  // 报名的位置，只有前8名才显示在屏幕上，Pos :1 2 3 4 5 6 7 8
}

// 玩家中奖的结构
type StPlayerAwardInfo struct {
	PlayerUID              uint32 // 用户的唯一的UID
	PlayerWeiXinName       string // 现场用户报名的 微信的名字
	PlayerWeiXinHeadPicUrl string // 现场用户报名的 微信的头像地址
	PlayerPaiMingPos       uint8  // 报名的位置，只有前8名才显示在屏幕上，Pos :1 2 3 4
	AwardID                uint32 // 奖品的ID
}

//=============================================================================================

// 用户注册的时候，赠送给随机的数量。
const (
	PLAYE_MONEY     int64 = 500 // 赠送 GO豆
	PLAYE_PAY_MONEY int64 = 100 // 赠送 GO币
)

// 圈子列表信息
type StCircleInfo struct {
	CircleId             int64  // 圈子ID
	CircleName           string // 圈子名子
	CircleIcon           string // 圈子图标 客户端获取路径：
	CircleState          int8   // 圈子的状态 1 ：表示最新  2 ：表示最热  3 ：表示最火
	ArticleNum           int64  // 圈子总共文章的数量
	NewArticleTitle      string // 最新的文章的标题
	NewArticleAuthorUID  int64  // 最新的文章的人物的UID
	NewArticleAuthorIcon string // 最新的文章的人物的头像
	NewArticleAuthorName string // 发表最新文章的用户的名字
	CircleSummary        string // 圈子的简介
}

// 文章的结构信息
type StArticleInfo struct {
	ArticleId         string // 文章的ID
	ArticleName       string // 文章的名字
	ArticleAuthor     string // 文章的作者
	ArticleAuthorIcon string // 文章作者的ICON
	ArticleContent    string // 文章的内容
	ArticleTitle      string // 文章的标题
	ArticleNewComment string // 文章的最新的评论；XX：评论 xxxxx
	ArticleIssueTime  string // 文章的发表的时间;时间戳
	ArticleState      string // 文章的类型；1 表示 精华 2 表示 最新  3 表示 最热
	ICircleId         string // 圈子的ID
	IArticleAuthorUID string // 作者的UID
}

// 文章的评论结构信息
type StCommentInfo struct {
	CommentId         string // 评论的ID
	CommentAuthor     string // 评论的作者
	CommentAuthorIcon string // 评论作者的ICON
	CommentAuthrID    string // 评论作者的UID
	CommentContent    string // 评论的内容
	CommentTime       string // 评论的时间；时间戳
}

type StUserRegInfo struct {
	Name          string  // 用户名
	Pwd           string  // 密码
	IMei          string  // 设备类型
	HeadPic       string  // 头像
	Sex           float64 // 性别
	DescText      string  // 用户签名
	Type          float64 // 用户类型
	MobileType    string  // 手机类型
	OsType        string  // 系统类型
	SdkVer        string  // sdk版本
	RegTime       float64 // 注册时间
	LastLoginTime float64 // 最后一次登录时间
	CltVer        float64 // 客户端版本
	LoginType     float64 // 登录类型
	ThirdOpenId   string  // 第三方唯一标识
	ThirdToken    string  // 第三方帐号密钥
	ThirdExpire   float64 // 第三方过期时间
	NotifyMsg     string  // 消息
}

// APP商城信息,主要是数据的整合
type StShopInfo struct {
	ItemId          int64  // id
	ItemName        string // 名字
	ItemIconName    string // 图标
	ItemPrice       int64  // 价格
	ItemType        int64  // 类型
	ItemRebate      int64  // 折扣
	ItemRebateMoney int64  // 折后价格
	ItemDesc        string // 商品描述
}

// 账单的历史记录的结构体
type StHistoryPayInfo struct {
	PayInfoId  int32  // 数据库中保存的数据库的自增的ID的信息
	SrcRoleUID int64  // 谁赠送的数据UID
	DecRoleUID int64  // 赠送给谁的UID
	CoinNumber int64  // 赠送的金币的数量
	ReasonMsg  string // 赠送的理由
	Instertime string // 赠送的时间
}

// 问答社区的结构体
type StAnswerInfo struct {
	AuthorRoleUID string // 问答社区 作者的的UID
	AuthorName    string // 问答社区 作者的名字
	AuthorIcon    string // 问答社区 作者的ICON
	AnswerContent string // 问答社区 内容数据
	Instertime    string // 问答社区 发表的时间
	Answerid      string // 问答社区 帖子的ID
}

//随机函数
func RandData(iMode int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//Log_Eio.Log(" rand data : %d", r.Intn(iMode))
	return r.Intn(iMode)
}

// md5 加密函数
func MD5ToString() string {
	// md5 加密
	h := md5.New()
	h.Write([]byte(strconv.Itoa(RandData(1000)))) // 需要加密的字符串为 sharejs.com
	return hex.EncodeToString(h.Sum(nil))
}

// md5 加密函数
func MD5ToStringFromString(strdata string) string {
	// md5 加密
	h := md5.New()
	h.Write([]byte(strdata)) // 需要加密的字符串为
	return hex.EncodeToString(h.Sum(nil))
}

// 数据格式转换：
func GBKConvertUTF8(content string) (ret string) {
	dec := mahonia.NewDecoder("gbk")
	ret, ok := dec.ConvertStringOK(content)
	if ok {
		return ret
	} else {
		return ""
	}
}

//--------------------------csv 格式结构定义---------------------------------------------
// 游戏列表
type StGameListInfo struct {
	GameId   uint32 // 游戏的ID
	GameName string // 游戏的名字
	Ip       string // ip地址
	Port     uint32 // 端口
	Type     string // 1：表示深度游戏；2：表示活动游戏
}

// 游戏活动
type StActivitiesInfocsv struct {
	ID        string // 序号
	LoginName string // 商家的登陆名字
	PicUrl    string // 图片地址
	Type      string // 1：全部获奖；2：非quanbu
	JiLV      string
}

// 游戏配置
type StGameConfigInfo struct {
	GameName      string // 游戏名字
	GameID        string // 游戏ID
	PlayerMax     string // 最大人数限制
	PlayerTimeOut string // 超时踢人
	GmaeTime      string // 游戏时长
	PCTimeOut     string // 超时踢掉PC链接
	GetGameJiFen  string // 游戏获取积分
}

// csv配置表
var G_StCard2InfoBaseST map[string]*Card2InfoBase // 卡牌活动结构

// 卡牌活动结构
type Card2InfoBase struct {
	Card2ID       string // 卡牌的ID
	Card2Msg      string // 卡牌的描述
	Card2GameName string // 卡牌的地点
	Card2GameID   string // 策划看到的类型
	PicPath       string //  图片路径
}
