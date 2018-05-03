package PlayerData

// 玩家的结构体
type YLGame_Plyerdata struct {
	PlayerAccountID  string                           // 玩家的账号ID;微信用户的OpenID
	PlayerRoleUID    string                           // 玩家的角色ID；服务器数据唯一的ID
	PlayerNickname   string                           // 微信昵称
	PlayerHeadimgurl string                           // 微信头像地址
	PlayerSex        uint32                           // 性别
	PlayerLev        string                           // 玩家的普通等级   上限255；根据等级的饿
	PlayerVIPLev     string                           // 玩家的vip等级    上限255 根据充值的付费的情况而定
	PlayerJiFen      string                           // 玩家获取的积分，等级挂钩
	PlayerItemMoney  string                           // 道具商城金币
	PlayerCardMoney  string                           // 卡卷商城金币
	PlayerDouDiZhu   string                           // 斗地主积分
	GItemST          map[string]*YLGame_Player_ItemST // 道具列表
	Sign_In          string                           // 连续签到天数，7天后默认奖励都是一样的
}

// 玩家的道具的结构
type YLGame_Player_ItemST struct {
	ItemID    string // 道具ID
	IsTimeOut bool   // 是否过期
	GameID    string // 道具所属游戏ID
	Number    string // 道具数量
}

//  换算等级的函数
//	public static int GameLevelToXP(float _level)
//	{
//		return (int)(Mathf.Exp (_level/40)*1000-1000)*10;
//	}
