package Proto3_Data

//【1】 通信相关 Network_Data_Proto 子协议
const (
	NetworkINIT = iota
	// 心跳检测
	Net_Heart_Beating_Proto // Net_Heart_Beating_Proto == 1 心跳检测

	// 微信用户授权后的协议，获取自己的数据及讲师的列表
	// 服务器返回的数据是类似与个人中心的数据
	C2S_Net_WinXin_OpenID_BaoMing_Proto // Net_WinXin_OpenID_BaoMing_Proto == 2
	S2C_Net_WinXin_OpenID_BaoMing_Proto // S2C_Net_WinXin_OpenID_BaoMing_Proto == 3

	// 答题者选择---老师
	C2S_WinXinPlayer_Choose_T_Proto // C2S_WinXinPlayer_Choose_T_Proto == 4
	S2C_WinXinPlayer_Choose_T_Proto // S2C_WinXinPlayer_Choose_T_Proto == 5

	// 广播协议--> 更新所有客户端的老师的状态信息；注：老师上线，下线广播，答题中广播忙碌
	Teacher_State_Broadcast_Proto // Teacher_State_Broadcast_Proto == 6

	// 广播协议（针对房间内人员有效）--临时房间交流
	Room_AC_Broadcast_Proto // Room_AC_Broadcast_Proto == 7

	// 用户讲师申请--普通用户申请为老师
	// 需要提供社区资格（申请码）--微店购买价值1000元，目前活动5折
	C2S_WinXinPlayer_Apply_Teacher_Proto // C2S_WinXinPlayer_Apply_Teacher_Proto == 8

	// 提问功能
	C2S_WinXinPlayer_TiWen_Proto // C2S_WinXinPlayer_TiWen_Proto == 9
	S2C_WinXinPlayer_TiWen_Proto // S2C_WinXinPlayer_TiWen_Proto == 10

	// 获取老师列表
	C2S_WinXinPlayer_GetList_T_Proto // C2S_WinXinPlayer_GetList_T_Proto == 11
	S2C_WinXinPlayer_GetList_T_Proto // S2C_WinXinPlayer_GetList_T_Proto == 12

	// 获取问题列表
	C2S_WinXinPlayer_GetList_WT_Proto // C2S_WinXinPlayer_GetList_WT_Proto == 13
	S2C_WinXinPlayer_GetList_WT_Proto // S2C_WinXinPlayer_GetList_WT_Proto == 14

	// 发表评论
	C2S_WinXinPlayer_Comment_WT_Proto // C2S_WinXinPlayer_Comment_WT_Proto == 15
	S2C_WinXinPlayer_Comment_WT_Proto // S2C_WinXinPlayer_Comment_WT_Proto == 16

	// 返回数据给前端--测试
	S2C_WinXinPlayer_Apply_Teacher_Proto // S2C_WinXinPlayer_Apply_Teacher_Proto == 17

	// 获取我的提问的列表
	C2S_WinXinPlayer_My_WT_Proto // C2S_WinXinPlayer_My_WT_Proto == 18
	S2C_WinXinPlayer_My_WT_Proto // S2C_WinXinPlayer_My_WT_Proto == 19

	// 获取提问的评论的列表
	C2S_WinXinPlayer_CommentList_WT_Proto // C2S_WinXinPlayer_CommentList_WT_Proto == 20
	S2C_WinXinPlayer_CommentList_WT_Proto // S2C_WinXinPlayer_CommentList_WT_Proto == 21

	// 获取提问的评论的列表--选择最优的评论
	C2S_WinXinPlayer_CommentList_Of_Best_WT_Proto // C2S_WinXinPlayer_CommentList_Of_Best_WT_Proto == 22
	S2C_WinXinPlayer_CommentList_Of_Best_WT_Proto // S2C_WinXinPlayer_CommentList_Of_Best_WT_Proto == 23

	// 发送数据聊天的
	C2S_WinXinPlayer_Chat_WT_Proto // C2S_WinXinPlayer_Chat_WT_Proto == 24
	S2C_WinXinPlayer_Chat_WT_Proto // S2C_WinXinPlayer_Chat_WT_Proto == 25
)

// --------------------------------------------------------------------------------

// 【发送数据聊天的】 C2S_WinXinPlayer_Chat_WT_Proto
type C2S_WinXinPlayer_Chat_WT struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	TOpenID   string // 老师的openid
	MyOpenID  string // 我的openid
	Data      string // 聊天内容
}

// 服务器返回我的选择最优的评论
// S2C_WinXinPlayer_Chat_WT_Proto
type S2C_WinXinPlayer_Chat_WT struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	TOpenID   string // 老师的openid
	WOpenID   string // 问问题的openid
	WHeadUrl  string // 问问题的URL
	WName     string // 问答者的名字
	Data      string // 聊天内容
}

// --------------------------------------------------------------------------------

//type WX_Player_DataDB struct {
//	ID      string
//	OpenID  string
//	Name    string
//	HeadUrl string
//	XingJi  string
//	Coin    string
//	Type    string
//	Time    string
//}

// 客户端请求
type C2S_WinXinPlayer_GetList_T struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
}

// 获取老师列表
// S2C_WinXinPlayer_GetList_T_Proto
type S2C_WinXinPlayer_GetList_T struct {
	Protocol  uint32                       // 主协议
	Protocol2 uint32                       // 子协议
	Data      map[string]*WX_Player_DataDB // 微信用户绑定的系统的存储数据
}

// --------------------------------------------------------------------------------

// 【获取提问的评论的列表--选择最优的评论】 C2S_WinXinPlayer_CommentList_Of_Best_WT_Proto
type C2S_WinXinPlayer_CommentList_Of_Best_WT struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	WenDaID   string // 问答的ID
	CommentID string // 评论的ID
}

// 服务器返回我的选择最优的评论
// S2C_WinXinPlayer_CommentList_Of_Best_WT_Proto
type S2C_WinXinPlayer_CommentList_Of_Best_WT struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	IS_Succ   bool   // 是否成功
}

// --------------------------------------------------------------------------------

type WenDaPLBak struct {
	Id       string
	WenTiID  string // 问答的id
	OpenID   string // 发布问答的openid
	PlOpenID string // 评论者的openid
	Data     string // 评论的内容
	Type     string // 评论的类型
	Time     string // 评论的时间
}

type WenDaPL struct {
	Id      string
	OpenID  string // 评论者的openid
	Name    string // 评论者的名字
	Headurl string // 评论者的头像
	Data    string // 评论的内容
	Type    string // 是不是 最优评论
	Time    string // 评论的时间
}

// 【获取我的问答列表】 C2S_WinXinPlayer_CommentList_WT_Proto
type C2S_WinXinPlayer_CommentList_WT struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	OpenID    string // 用户的唯一凭证
	WenDaID   string // 问答的ID
	PageNum   string // 第一次获取PageNum == ”0“  ； 后面的页面 PageNum == 每页的最小ID
}

// 服务器返回我的问答列表
// S2C_WinXinPlayer_CommentList_WT_Proto
type S2C_WinXinPlayer_CommentList_WT struct {
	Protocol  uint32              // 主协议
	Protocol2 uint32              // 子协议
	Data      map[string]*WenDaPL // 5个数据
}

// --------------------------------------------------------------------------------

// 【获取我的问答列表】 C2S_WinXinPlayer_My_WT_Proto
type C2S_WinXinPlayer_My_WT struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	OpenID    string // 用户的唯一凭证
	PageNum   string // 第一次获取PageNum == ”0“  ； 后面的页面 PageNum == 每页的最小ID
}

// 服务器返回我的问答列表
// S2C_WinXinPlayer_My_WT_Proto
type S2C_WinXinPlayer_My_WT struct {
	Protocol  uint32              // 主协议
	Protocol2 uint32              // 子协议
	Data      map[string]*WenDaJG // 5个数据
}

// --------------------------------------------------------------------------------

// 【发表评论】 C2S_WinXinPlayer_Comment_WT_Proto
// 用户提问
// 数据需要保存用户的绑定信息
type C2S_WinXinPlayer_Comment_WT struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	WenTiID   string // 问答的ID
	OpenID    string // 提问者的openid
	PLOpenID  string // PLOpenID,评论问题的人的openid
	Data      string // 组合数据发过来--服务器只负责保存---评论的内容
	Type      string // 评论的类型--提问者选择的最佳的评论
}

// S2C_WinXinPlayer_Comment_WT_Proto
type S2C_WinXinPlayer_Comment_WT struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	IS_Succ   bool   // 是否成功
}

// 数据更新操作的说明
// --------------------------------------------------------------------------------

// --------------------------------------------------------------------------------
// 问答结构
type WenDaJG struct {
	Id      string
	OpenID  string
	Name    string
	Headurl string
	Data    string
	PicData string
	Coin    string
	State   string
	Time    string
	PlName  string
	PlData  string
}

// 【获取问答列表】 C2S_WinXinPlayer_GetList_WT_Proto
type C2S_WinXinPlayer_GetList_WT struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	OpenID    string // 用户的唯一凭证
	PageNum   string // 第一次获取PageNum == ”0“  ； 后面的页面 PageNum == 每页的最小ID
}

// 服务器返回
// S2C_WinXinPlayer_GetList_WT_Proto
type S2C_WinXinPlayer_GetList_WT struct {
	Protocol  uint32              // 主协议
	Protocol2 uint32              // 子协议
	Data      map[string]*WenDaJG // 5个数据
}

// --------------------------------------------------------------------------------

// 【提问功能】 C2S_WinXinPlayer_TiWen_Proto
// 用户提问
// 数据需要保存用户的绑定信息
type C2S_WinXinPlayer_TiWen struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	OpenID    string // 用户的唯一凭证
	Data      string // 组合数据发过来--服务器只负责保存
	PicData   string // 图片数据  base64数据；服务器存储
	Coin      string // 悬赏金币---由提问者自己选择多少，且大于1至少，最高100
}

// S2C_WinXinPlayer_TiWen_Proto
type S2C_WinXinPlayer_TiWen struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	IS_Succ   bool   // 是否成功
}

// 数据更新操作的说明
// --------------------------------------------------------------------------------

// 【用户讲师申请】 C2S_WinXinPlayer_Apply_Teacher_Proto
// 用户讲师申请--普通用户申请为老师
// 需要提供社区资格（申请码）--微店购买价值1000元，目活动5折
type C2S_WinXinPlayer_Apply_Teacher struct {
	Protocol   uint32 // 主协议
	Protocol2  uint32 // 子协议
	OpenID     string // 申请人的微信的OpenID 数据--服务器判断用户的唯一
	Apply_Code string // 申请码  GM 系统生成
}

// S2C_WinXinPlayer_Apply_Teacher_Proto
type S2C_WinXinPlayer_Apply_Teacher struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	Code      string // 1:表示邀请码已经使用，2：邀请码不存在，3：申请成功,4:已经申请中
}

//--------------------------------------------------------------------------------

// 【临时房间交流】 Room_AC_Broadcast_Proto
// 广播协议（针对房间内人员有效）--临时房间交流
type Room_AC_Broadcast struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	RoomUID   string // 消息接收的房间UID
	Data      string // 组合数据--手机+服务器按照一定规则定制;同时包括解答问题结束

}

//--------------------------------------------------------------------------------

// 【广播协议】
// 更新所有客户端的老师的状态信息
type Teacher_State_Broadcast struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	T_UID     string // 老师UID信息--服务器唯一ID
	T_State   string // t_1:表示空闲，t_2:忙碌，t_3：不在线，t_4:未知状态
}

//--------------------------------------------------------------------------------

// 【答题者选择---老师】 C2S_WinXinPlayer_Choose_T_Proto
//  手机请求
type C2S_WinXinPlayer_Choose_T struct {
	Protocol     uint32 // 主协议
	Protocol2    uint32 // 子协议
	T_PlayerUID  string // 提问者的UID
	T_TeacherUID string // 选择的讲师的UID信息
	T_Coin       string // 选择悬赏的金币的钱数 -----> 社区币;用户兑换活动费用等 1元=10个，用户答题使用；充值到数据
	T_Problem    string // 问题（问题描述）--->  暂时不支持  代码块提交；迭代版本中再优化
	T_IsPic      bool   // 是否有图片数据;  false:表示无； true :存在
	T_PicBase64  string // 图片base64数据；注：图片限制再90K传输；且只可以是1张
	T_IsZJ       bool   // 如果老师角色此题难度 超过提问者悬赏，当提问者设置可以追加后，老师可以点击付费增加

}

// 【答题者选择---老师】 S2C_WinXinPlayer_Choose_T_Proto
// 服务器返回
// 此协议同样需要广播给对应老师及提问者
type S2C_WinXinPlayer_Choose_T struct {
	Protocol     uint32                  // 主协议
	Protocol2    uint32                  // 子协议
	T_State      string                  // t_1:表示空闲，t_2:忙碌，t_3：不在线，t_4:未知状态  可以去除，但是为了校验可以保留（如果因为网络广播协议没有收到，以此标志位为准）
	TMP_RoomData map[string]*TEMP_RoomST // 服务器建立临时交流房间结构
}

//--------------------------------------------------------------------------------
// 用户在社区系统的结构数据
type WX_Player_Data struct {
	UID       string
	Name      string
	GroupType string // 用户组的权限,1：为问答者；2：为老师；3：为官方监督员（禁言用户权限，封号处理等）
	Lev       string // 用户的级别；包括老师星级
	Coin      string // 社区币;用户兑换活动费用等 1元=10个，用户答题使用；充值到数据
	Time      string // 注册时间
}

type WX_Player_DataDB struct {
	ID      string
	OpenID  string
	Name    string
	HeadUrl string
	XingJi  string
	Coin    string
	Type    string
	Time    string
}

// 【微信用户报名的协议】 S2C_Net_WinXin_OpenID_BaoMing_Proto
type S2C_Net_WinXin_OpenID_BaoMing struct {
	Protocol  uint32                       // 主协议
	Protocol2 uint32                       // 子协议
	Data      map[string]*WX_Player_DataDB // 微信用户绑定的系统的存储数据
}

// 【微信用户报名的协议】 C2S_Net_WinXin_OpenID_BaoMing_Proto
type C2S_Net_WinXin_OpenID_BaoMing struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	Code      string // 微信授权的Code
}

//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------

// 【手机踢人心跳协议】 Net_Heart_Beating_Proto
type Net_Heart_Beating struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	OpenID    string // 手机的OpenID
}

//--------------------------------------------------------------------------------
