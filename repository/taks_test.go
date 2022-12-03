package repository

import "testing"

func TestTaskMapper_addTask(t *testing.T) {
	var id uint
	DB.Raw("select id from `task` order by id desc limit 1 ").Scan(&id)
	tests := []struct {
		name string
		want uint
	}{
		{"addTask", id + 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := TaskMapper{}
			if got := ta.addTask(); got != tt.want {
				t.Errorf("addTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
