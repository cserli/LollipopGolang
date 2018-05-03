package Proto3_Data

//【1】 通信相关 Network_Data_Proto 子协议
const (
	NetworkINIT = iota
	// 心跳检测
	Net_Heart_Beating_Proto // Net_Heart_Beating_Proto == 1 心跳检测

)

// 【心跳协议】 Net_Heart_Beating_Proto
type Net_Heart_Beating struct {
	Protocol  uint32 // 主协议
	Protocol2 uint32 // 子协议
	OpenID    string // 手机的OpenID
}

//--------------------------------------------------------------------------------
