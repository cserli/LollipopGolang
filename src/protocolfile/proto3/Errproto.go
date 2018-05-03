package Proto3_Data

//【4】 错误提示相关 G_Error_Proto 子协议
const (
	ErrorINIT = iota
	// 公用错误提示
	G_Error_All_Proto // G_Error_All_Proto == 1 公用错误提示
)

//--------------------------------------------------------------------------------
// 【错误提示】 G_Error_All_Proto
type G_Error_All struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	ErrCode   string // 错误码 400001 申请码过期  400002 申请码已经使用  400003 申请码输入错误
	ErrMsg    string // 错误说明
}

type Data struct {
	data interface{}
}

//--------------------------------------------------------------------------------
