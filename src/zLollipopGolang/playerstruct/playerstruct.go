package PlayerData

// 玩家的结构体
type Plyerdata struct {
	PlayerAccountID int64  // 玩家的账号ID 暂时未使用。
	PlayerRoleUID   int64  // 玩家的角色ID
	PlayerName      string // 玩家的名字
	PlayerSex       int8   // 玩家的性别，0表示女性，1表示男性，2表示未知
	PlayerHeadUrl   string // 玩家头像URL 为空则是系统默认
	PlayerLev       int8   // 玩家的普通等级   上限255；根据等级的饿
	PlayerVIPLev    int8   // 玩家的vip等级    上限255 根据充值的付费的情况而定
	PlayerExp       int32  // 玩家的经验值     上限42亿，回答数据给奖励，经验的奖励
	PlayerMoney     int64  // 系统交换的金币
	PlayerPayMoney  int64  // 付费道具；充钱的道具的等级
}
