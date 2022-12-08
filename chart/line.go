package chart

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func lineSplitLine() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "马厂坝青云壹号TOD销售备案信息",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			SplitLine: &opts.SplitLine{
				Show: true,
			},
		}),
	)

	items, axis := generateLindItems()

	line.SetXAxis(axis).AddSeries("Category A", items,
		charts.WithLabelOpts(
			opts.Label{Show: true},
		))
	return line
}

type LineExamples struct{}

func (LineExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		lineSplitLine(),
	)
	os.Remove("line.html")
	f, err := os.Create("line.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
