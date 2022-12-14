package daq

import (
	"bytes"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"house/repository"
	"log"
	"strconv"
	"time"
)

const url = `https://zw.cdzjryb.com/roompricezjw/index.html?param=01B22707BDA6122314D9A47242F86EED9377E62F4C7B9D79884D32930ABC173D7B475137ED7750492A0539B4C0CE3A8D`

func SaleInfo() {
	taskMapper := repository.TaskMapper{}
	taskId := taskMapper.AddTask()
	saleHouseTotal := 0
	ch := make(chan int, 10)
	for i := 1; i <= 7; i++ {
		for y := 1; y <= 2; y++ {
			go doSaleInfo(i, y, taskId, ch)
		}
	}
	taskMapper.UpdateTotalById(taskId, saleHouseTotal)
	for i := 0; i < 14; i++ {
		saleNum := <-ch
		saleHouseTotal = saleHouseTotal + saleNum
	}
	fmt.Printf("总共卖出【%d】套\n", saleHouseTotal)
}

func doSaleInfo(i int, y int, taskId uint, ch chan<- int) {
	var navItemBuffer bytes.Buffer
	navItemBuffer.WriteString(`.rp-nav-item[data-val="`)
	navItemBuffer.WriteString(strconv.Itoa(i))
	navItemBuffer.WriteString(`"]`)

	var buffer bytes.Buffer
	buffer.WriteString(`.rp-subnav-item[data-parentval="`)
	buffer.WriteString(strconv.Itoa(i))
	buffer.WriteString(`"][data-val="`)
	buffer.WriteString(strconv.Itoa(y))
	buffer.WriteString(`"]`)
	total, hostList := parse(url, buffer.String(), navItemBuffer.String(), i, y, taskId)
	houseRepository := repository.HouseRepository{}
	houseRepository.BatchInsert(hostList)
	fmt.Printf("\"%d栋，%d单元 已卖出【%d】套\n", i, y, total)
	ch <- total
}

func parse(url string, element string, navItem string, lowNum int, unit int, taskId uint) (int, []repository.House) {
	total := 0
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`body`),
	)
	if lowNum != 1 {
		err = chromedp.Run(ctx, chromedp.Click(navItem, chromedp.NodeVisible))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = chromedp.Run(ctx,
		chromedp.Sleep(time.Second),
		chromedp.Click(element, chromedp.NodeVisible),
		chromedp.WaitVisible(`tbody`),
		chromedp.Sleep(time.Second),
		chromedp.Nodes("tbody tr", &nodes),
	)
	if err != nil {
		log.Fatal(err)
	}
	var allHouse []repository.House
	err = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		i := 1
		for _, node := range nodes {
			err := dom.RequestChildNodes(node.NodeID).WithDepth(-1).Do(ctx)
			if err != nil {
				return err
			}
			time.Sleep(time.Second * 2)
			value := node.Children[5].Children[0].Children[1].NodeValue
			sold := false
			if value != "可售" {
				sold = true
				total++
			}
			house := repository.House{
				FloorNum:   node.Children[0].Children[0].NodeValue,
				BuildNO:    strconv.Itoa(lowNum),
				Unit:       strconv.Itoa(unit),
				HouseNo:    node.Children[1].Children[0].NodeValue,
				Floorage:   node.Children[2].Children[0].NodeValue,
				Sold:       sold,
				Idx:        i,
				TaskId:     taskId,
				CreateTime: time.Now().Format("2006-01-02 15:04:05"),
			}
			allHouse = append(allHouse, house)
			i++
		}
		return nil
	}))
	if err != nil {
		log.Fatal(err)
	}
	cancel()
	return total, allHouse
}
