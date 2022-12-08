package chart

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"house/repository"
)

func generateBarItems() ([]opts.BarData, []string) {
	taskMapper := repository.TaskMapper{}
	taskList := taskMapper.FindAll()

	items := make([]opts.BarData, 0)
	axis := make([]string, 0)
	for i := range taskList {
		items = append(items, opts.BarData{Value: taskList[i].Total})
		axis = append(axis, taskList[i].TaskTime)
	}
	return items, axis
}

func generateLindItems() ([]opts.LineData, []string) {
	taskMapper := repository.TaskMapper{}
	taskList := taskMapper.FindAll()
	items := make([]opts.LineData, 0)
	axis := make([]string, 0)
	for i := range taskList {
		items = append(items, opts.LineData{Value: taskList[i].Total})
		axis = append(axis, taskList[i].TaskTime)
	}
	return items, axis
}
