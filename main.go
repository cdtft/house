package main

import (
	"house/daq"
)

func main() {
	//c := cron.New(cron.WithParser(cron.NewParser(
	//	cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
	//_, err := c.AddFunc("0 51 16 * * *", daq.SaleInfo)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//c.Run()
	daq.SaleInfo()
}
