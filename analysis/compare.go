package analysis

import "house/repository"

// 比较上一次的任务信息
func compareLatestTaskSaleInfo() {
	taskMapper := repository.TaskMapper{}
	houseRepository := repository.HouseRepository{}
	ids := taskMapper.FindLatestTaskIds()

}
