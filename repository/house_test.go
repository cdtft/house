package repository

import (
	"log"
	"testing"
)

func TestHouseRepository_SelectByDateString(t *testing.T) {
	type args struct {
		createTime string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				createTime: "2022-12-03",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ho := HouseRepository{}
			house := ho.SelectByDateString(tt.args.createTime)
			for i := range house {
				log.Println(house[i])
			}
		})
	}
}
