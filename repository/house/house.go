package house

import (
	"fmt"
	"house/repository"
)

type Id struct {
	Id int
}

func (*Id) test() {
	var ids []Id
	repository.DB.Raw("select * from config_info").Scan(&ids)
	for i := range ids {
		fmt.Println(ids[i])
	}
}
