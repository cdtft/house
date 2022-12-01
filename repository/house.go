package repository

type House struct {
	ID         uint   `gorm:"primaryKey"`
	Unit       string //单元号
	FloorNum   string //楼层
	HouseNo    string //房号
	Idx        int    //顺序编码
	Floorage   string //建筑面积
	Sold       bool   //是否售卖
	CreateTime string //创建时间
}

func (House) TableName() string {
	return "house"
}

type HouseRepository struct {
}

func (*HouseRepository) BatchInsert(houses []House) {
	DB.Create(&houses)
}
