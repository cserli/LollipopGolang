package Proto3_Data

// 临时房间结构

type TEMP_RoomST struct {
	RoomUID  string // 房间UID 时间戳微秒级别都可以
	RoomTime uint32 // 房间的时间限制，如果10分钟没有交流，自动关闭
}

// 申请列表结构
type Apply_List struct {
	UID    string // 列表唯一ID
	OPenID string // 微信用户申请UID
	Code   string // 申请码
	MsgID  string // 申请理由  --- wap 与服务器协商一致即可
	Time   string // 申请时间
}
