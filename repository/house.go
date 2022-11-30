package repository

import (
	"fmt"
)

type House struct {
	Id       int    //ID
	FloorNum string //楼层
	HouseNo  string //房号
	Floorage string //建筑面积
	Sold     bool   //是否售卖
	Index    int    //顺序编码
}

func (*House) test() {
	var ids []House
	DB.Raw("select * from config_info").Scan(&ids)
	for i := range ids {
		fmt.Println(ids[i])
	}
}
