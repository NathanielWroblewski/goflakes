package views

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

import (
	"fmt"

	models "../models"
	utils "../utils"

	svg "github.com/ajstarks/svgo/float"
)

var orientation = models.Pointy()
var size = models.Vector{3.0, 3.0}
var xscache = map[string][]float64{}
var yscache = map[string][]float64{}

// Hex draws a hexagon onto an SVG context
func Hex(hex models.Hex, canvas *svg.SVG, width float64, height float64) {
	var xs, ys []float64

	attributes := fmt.Sprintf("line-width: 1; fill: %s; stroke: #A7BBC9;", utils.Colorize(hex.Value))

	xs, ok := xscache[hex.Key()]
	if ok {
		ys = yscache[hex.Key()]
	} else {
		origin := models.Vector{width / 2, height / 2}

		for _, vector := range hex.Corners(orientation, size, origin) {
			xs = append(xs, vector.X())
			ys = append(ys, vector.Y())
		}

		xs = append(xs, xs[0])
		ys = append(ys, ys[0])

		xscache[hex.Key()] = xs
		yscache[hex.Key()] = ys
	}

	canvas.Polyline(xs, ys, attributes)
}
