package Proto3_Data

// G_Gm_Data_Proto == 14 的主协议

// GM 后台系统协议
const (
	GM_INIT = iota

	// GM增加申请码--手机wap+web都支持
	GM2S_Add_Apply_Code_Proto // GM2S_Add_Apply_Code_Proto == 1
	S2GM_Add_Apply_Code_Proto // S2GM_Add_Apply_Code_Proto == 2

	// GM 禁言/封号/解封 协议
	GM2S_Promissio_Insurata_Proto // GM2S_Promissio_Insurata_Proto == 3
	S2GM_Promissio_Insurata_Proto // S2GM_Promissio_Insurata_Proto == 4

	// GM 审核申请
	GM2S_Auditing_Apply_Proto // GM2S_Auditing_Apply_Proto == 5
	S2GM_Auditing_Apply_Proto // S2GM_Auditing_Apply_Proto == 6

)

//--------------------------------------------------------------------------------

// 【GM 审核申请】 GM2S_Auditing_Apply_Proto
// 手机wap+web 请求
type GM2S_Auditing_Apply struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	Itype     string // GM_A_1:获取申请列表   GM_A_2:审批申请
	ListID    string // 申请列表UID （ Itype == GM_A_2 有效）
	OP_A_ID   string // 操作ID （ Itype == GM_A_2 有效）  OP_A_ID == 服务器和GM系统协商
}

// 【GM 审核申请】 S2GM_Auditing_Apply_Proto
// 服务器回复
type S2GM_Auditing_Apply struct {
	Protocol  uint32                 // 主协议
	Protocol2 uint32                 // 子协议
	ApplyList map[string]*Apply_List // 申请列表数据 （ Itype == GM_A_1 才返回）
	RetMsg    string                 // 数据返回
}

//--------------------------------------------------------------------------------

// 【禁言/封号/解封】 GM2S_Promissio_Insurata_Proto
// GM 禁言/封号/解封 协议
type GM2S_Promissio_Insurata struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	OPUID     string // 系统管理UID
	OpenID    string // 微信用户的UID
	Itype     string //  PI_1 : 禁言（Time为禁言时长）；PI_2 : 封号；PI_3 : 解封；
	Time      string // 时间戳
	MsgID     string // 理由
}

// 【禁言/封号/解封】 S2GM_Promissio_Insurata_Proto
// GM 禁言/封号/解封 协议
type S2GM_Promissio_Insurata struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	RetMSG    string // 返回数据给GM
}

//--------------------------------------------------------------------------------

// 【用户讲师申请】 GM2S_Add_Apply_Code_Proto
// 手机wap+web 请求
type GM2S_Add_Apply_Code struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	OPUID     string // 系统管理UID
}

// 【用户讲师申请】 S2GM_Add_Apply_Code_Proto
// 服务器回复
type S2GM_Add_Apply_Code struct {
	Protocol   uint32 // 主协议
	Protocol2  uint32 // 子协议
	Apply_Code string // 申请码
}

//--------------------------------------------------------------------------------
