package main

import (
	"bytes"
	"encoding/base64"

	"github.com/wcharczuk/go-chart"
)

var (
	widthPerPoint = 2
)

// Chart holds an io.Writer which represents a png of a sine wave
type Chart struct {
	b *bytes.Buffer
}

// NewChart takes a series of points, as generated from Sine{}.Graph()
// and creates a png representing these points plotted as a graph
func NewChart(points []float64) (c Chart, err error) {
	position := make([]float64, len(points))
	for i := range position {
		position[i] = float64(i)
	}

	ts := chart.ContinuousSeries{
		Name: "Wave",
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.GetDefaultColor(2),
		},
		XValues: position,
		YValues: points,
	}

	empty := chart.StyleShow()
	empty.Show = false

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      "Position",
			NameStyle: chart.StyleShow(),
			Style:     empty,
		},

		YAxis: chart.YAxis{
			Name:      "Phase",
			NameStyle: empty,
			Style:     chart.StyleShow(),
		},

		Series: []chart.Series{
			ts,
		},

		Width: len(points) * widthPerPoint,
	}

	c.b = bytes.NewBuffer([]byte{})

	err = graph.Render(chart.PNG, c.b)

	return
}

// Base64 returns a base64 representation of this chart
func (c Chart) Base64() string {
	return base64.StdEncoding.EncodeToString(c.b.Bytes())
}
