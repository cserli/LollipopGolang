package Proto_Data

// 主协议定义
const (
	ProtoINIT = iota // ProtoINIT == 0  初始化协议
	// 客户端与服务器通信 主要协议

	//【1】 官方相关
	Firms_Data_Proto // Firms_Data_Proto == 1  网络相关主协议
	//【2】 玩家相关
	FE_User_Data_Proto // FE_User_Data_Proto == 2  现场用户先关协议
	//【3】 网络相关
	Network_Data_Proto // Network_Data_Proto == 3  网络相关
	//【4】 错误相关
	G_Error_Proto // G_Error_Proto == 4  错误提示

)
