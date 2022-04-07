package arcaneStorageDataType

type PC_P01_A struct {
	GAME_STATE string `redis:"GameState"`
	// 魔方电池的五面
	PC_P01_I01_A bool `redis:"PC-P01_I01_A"`
	PC_P01_I01_B bool `redis:"PC-P01_I01_B"`
	PC_P01_I01_C bool `redis:"PC-P01_I01_C"`
	PC_P01_I01_D bool `redis:"PC-P01_I01_D"`
	PC_P01_I01_E bool `redis:"PC-P01_I01_E"`
	//符文能量台
	PC_P01_I03_A bool `redis:"PC-P01_I03-A"`
	PC_P01_I03_B bool `redis:"PC-P01_I03-B"`
	PC_P01_I03_C bool `redis:"PC-P01_I03-C"`
	PC_P01_I03_D bool `redis:"PC-P01_I03-D"`
	PC_P01_I03_E bool `redis:"PC-P01_I03-E"`
	//场景符文石
	PC_P01_I02_A bool `redis:"PC-P01_I02-A"`
	PC_P01_I02_B bool `redis:"PC-P01_I02-B"`
	PC_P01_I02_C bool `redis:"PC-P01_I02-C"`
	PC_P01_I02_D bool `redis:"PC-P01_I02-D"`
	PC_P01_I02_E bool `redis:"PC-P01_I02-E"`
}
