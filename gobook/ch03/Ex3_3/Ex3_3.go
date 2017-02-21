package main

import (
	"math"
	"fmt"
)

/*
 * Exercise 3.3: Color each polygon based on its height, so that the peaks are colored red (# ff0000)
 * and the valleys blue (# 0000ff).
 */

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width: 0.7' " +
		"width='%d' height='%d'>\n", width, height)

	for i := 0; i <  cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, color := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)
			if !(math.IsInf(ax, 0) ||
				math.IsInf(ax, 0) ||
				math.IsInf(ay, 0) ||
				math.IsInf(bx, 0) ||
				math.IsInf(by, 0) ||
				math.IsInf(cx, 0) ||
				math.IsInf(cy, 0) ||
				math.IsInf(dx, 0) ||
				math.IsInf(dy, 0)) {
				fmt.Printf("<polygon style='fill:%s' points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
					color, ax, ay, bx, by, cx, cy, dx, dy)
			}

		}
	}
	fmt.Println("</svg>")

}

func corner(i, j int) (float64, float64, string) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x,y)

	color := "00ff00"
	if z > 0 {
		color = "#ff0000"
	}
	if z < 0 {
		color = "#0000ff"
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, color
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)

	return math.Sin(r) / r
}