package main

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

import (
	"flag"
	"fmt"
	"log"
	"os"

	svg "github.com/ajstarks/svgo/float"

	models "./models"
	views "./views"
)

func main() {
	width := flag.Float64("width", 100.0, "a positive integer representing the output svg width in pixels")
	height := flag.Float64("height", 100.0, "a positive integer representing the output svg height in pixels")
	iters := flag.Int("iters", 400, "a positive integer representing the number of discrete time iterations to take")
	flag.Parse()

	hexmap := models.Hexagon()

	for i := 0; i < *iters; i++ {
		f, err := os.Create(fmt.Sprintf("./output/%06d.svg", i))
		if err != nil {
			log.Fatal(err)
		}

		canvas := svg.New(f)
		canvas.Start(*width, *height)

		hexmap = hexmap.Step(func(hex models.Hex) {
			views.Hex(hex, canvas, *width, *height)
		})

		canvas.End()
	}
}
