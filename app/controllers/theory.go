package controllers

import (
	"github.com/revel/revel"
	"math"
)

type Theory struct {
	*revel.Controller
}

type Point [2]float64

func (c Theory) Index() revel.Result {
	return c.RenderJson([][2]float64{[2]float64{1, 2}})
}

func (c Theory) AwgnBpsk() revel.Result {
	curve := make([]Point, 301)
	for i := 0; i <= 300; i += 1 {
		x := float64(i) / 10.0
		cnr := math.Pow(10.0, x/10.0)
		y := math.Erfc(math.Sqrt(cnr)) / 2.0
		curve[i] = Point{x, y}
	}
	return c.RenderJson(curve)
}

func (c Theory) AwgnQpsk() revel.Result {
	curve := make([]Point, 301)
	for i := 0; i <= 300; i += 1 {
		x := float64(i) / 10.0
		cnr := math.Pow(10.0, x/10.0)
		y := math.Erfc(math.Sqrt(cnr/2.0)) / 2.0
		curve[i] = Point{x, y}
	}
	return c.RenderJson(curve)
}
