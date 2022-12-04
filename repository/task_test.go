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
			if got := ta.AddTask(); got != tt.want {
				t.Errorf("addTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskMapper_UpdateTotalById(t *testing.T) {
	type args struct {
		id    uint
		total int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "UpdateTotalById",
			args: args{
				id:    8,
				total: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := TaskMapper{}
			ta.UpdateTotalById(tt.args.id, tt.args.total)
		})
	}
}
