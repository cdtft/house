package analysis

import (
	"house/repository"
	"log"
	"strconv"
)

// 比较上一次的任务信息
func compareLatestTaskSaleInfo() {
	taskMapper := repository.TaskMapper{}
	houseRepository := repository.HouseRepository{}
	ids := taskMapper.FindLatestTaskIds()
	fistTaskId := ids[0]
	secondTaskId := ids[1]
	firstHouse := houseRepository.FindByTaskId(fistTaskId)
	secondHouse := houseRepository.FindByTaskId(secondTaskId)
	secondHouseMap := make(map[string]bool, len(secondHouse))
	for i := range secondHouse {
		house := secondHouse[i]
		key := house.BuildNO + "-" + house.Unit + "-" + house.HouseNo + "-" + strconv.Itoa(house.Idx)
		log.Printf("%d, %s\n", i, key)
		secondHouseMap[key] = true
		log.Println(len(secondHouseMap))
	}
	log.Println(len(secondHouseMap))
	for i := range firstHouse {
		house := firstHouse[i]
		key := house.BuildNO + "-" + house.Unit + "-" + house.HouseNo + "-" + strconv.Itoa(house.Idx)
		b := secondHouseMap[key]
		if b == false {
			log.Println(house)
		}
	}
}
