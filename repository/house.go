package repository

type House struct {
	Id       int    //ID
	FloorNum string //楼层
	HouseNo  string //房号
	Floorage string //建筑面积
	Sold     bool   //是否售卖
	Index    int    //顺序编码
}

type HouseRepository struct {
}

func (*HouseRepository) BatchInsert(houses []House) {
	DB.Create(&houses)
}
