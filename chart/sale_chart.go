package chart

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// generate random data for bar chart

func CreateBar() {
	items, axis := generateBarItems()
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "马厂坝青云壹号TOD销售备案信息",
		Subtitle: "It's extremely easy to use, right?",
	}))

	// Put data into instance
	bar.SetXAxis(axis).
		AddSeries("Category A", items)
	// Where the magic happens
	_ = os.Remove("bar.html")
	f, _ := os.Create("bar.html")
	bar.Render(f)
}
