package Redis_DB

//	"strconv"
//	"zLollipopGolang/playerstruct"

// 更新内存数据库的数据
// 设计原理
// 1 获取内存的数据，保存为临时的变量，更新需要更新的数据的字段信息
// 2 直接更新数据到内存数据库中

// 更新内存数据中的人物的结构的数据信息： 更新内存数据库中的的人物的头像
func UpdateRedisPalyerStructOfHeadUrl(iroleUID int32, iHeadUrl string) bool {

	//	// 先获取内存数据的人物的数据的信息
	//	var PlayerData PlayerData.Plyerdata
	//	// 读取内存的数据
	//	PlayerData = Redis_Read_PlayerInfo(uint64(iroleUID))
	//	// 重新组合人物结构体数据
	//	PlayerData.PlayerHeadUrl = iHeadUrl
	//	// 重新写入到内存数据库中数据
	//	Redis_Write_PlayerInfo(strconv.Itoa(int(iroleUID)), PlayerData)

	return true
}

// 更新内存数据中的人物的名字信息
func UpdateRedisPalyerStructOfName(iroleUID int32, strName string) bool {

	//	// 先获取内存数据的人物的数据的信息
	//	var PlayerData PlayerData.Plyerdata
	//	// 读取内存的数据
	//	PlayerData = Redis_Read_PlayerInfo(uint64(iroleUID))
	//	// 重新组合人物结构体数据
	//	PlayerData.PlayerName = strName
	//	// 重新写入到内存数据库中数据
	//	Redis_Write_PlayerInfo(strconv.Itoa(int(iroleUID)), PlayerData)

	return true
}

// 更新内存数据中的人物的coin信息
func UpdateRedisPalyerStructOfCoin(iroleUID int32, icoin int32) bool {

	//	// 先获取内存数据的人物的数据的信息
	//	var PlayerData PlayerData.Plyerdata
	//	// 读取内存的数据
	//	PlayerData = Redis_Read_PlayerInfo(uint64(iroleUID))
	//	// 重新组合人物结构体数据
	//	PlayerData.PlayerMoney = int64(icoin)
	//	// 重新写入到内存数据库中数据
	//	Redis_Write_PlayerInfo(strconv.Itoa(int(iroleUID)), PlayerData)

	return true
}

// 更新内存数据中的人物的Paycoin信息
func UpdateRedisPalyerStructOfPayCoin(iroleUID int32, ipaycoin int32) bool {

	//	// 先获取内存数据的人物的数据的信息
	//	var PlayerData PlayerData.Plyerdata
	//	// 读取内存的数据
	//	PlayerData = Redis_Read_PlayerInfo(uint64(iroleUID))
	//	// 重新组合人物结构体数据
	//	PlayerData.PlayerPayMoney = int64(ipaycoin)
	//	// 重新写入到内存数据库中数据
	//	Redis_Write_PlayerInfo(strconv.Itoa(int(iroleUID)), PlayerData)

	return true
}
