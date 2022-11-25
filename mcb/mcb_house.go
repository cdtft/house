package mcb

import (
	"bytes"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"log"
	"strconv"
	"time"
)

const url = `https://zw.cdzjryb.com/roompricezjw/index.html?param=01B22707BDA6122314D9A47242F86EED9377E62F4C7B9D79884D32930ABC173D7B475137ED7750492A0539B4C0CE3A8D`

type House struct {
	FloorNum string //楼层
	HouseNo  string //房号
	Floorage string //建筑面积
	Sold     bool   //是否售卖
	Index    int    //顺序编码
}

func SaleInfo() {
	saleHouseTotal := 0
	for i := 1; i <= 7; i++ {
		for y := 1; y <= 2; y++ {
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
			fmt.Printf("%d栋，%d单元\n", i, y)
			total, hostList := parse(url, buffer.String(), navItemBuffer.String(), i)
			fmt.Printf("已卖出【%d】套\n", total)
			saleHouseTotal = saleHouseTotal + total
		}
	}
	fmt.Printf("总共卖出【%d】套\n", saleHouseTotal)
}

func parse(url string, element string, navItem string, lowNum int) (int, []*House) {
	total := 0
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`body`),
	)
	if lowNum != 7 {
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
	var allHouse []*House
	err = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		i := 0
		for _, node := range nodes {
			err := dom.RequestChildNodes(node.NodeID).WithDepth(-1).Do(ctx)
			if err != nil {
				return err
			}
			time.Sleep(time.Second)
			value := node.Children[5].Children[0].Children[1].NodeValue
			sold := false
			if value != "可售" {
				sold = true
				total++
			}
			house := &House{
				FloorNum: node.Children[0].Children[0].NodeValue,
				HouseNo:  node.Children[1].Children[0].NodeValue,
				Floorage: node.Children[2].Children[0].NodeValue,
				Sold:     sold,
				Index:    i,
			}
			allHouse = append(allHouse, house)
		}
		return nil
	}))
	if err != nil {
		log.Fatal(err)
	}
	cancel()
	return total, allHouse
}
