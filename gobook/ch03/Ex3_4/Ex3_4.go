package main

import (
	"math"
	"fmt"
	"net/http"
	"log"
	"io"
)

/*
 * Exercise 3.4: Following the approach of the Lissajous example in Section 1.7,
 * construct a web server that computes surfaces and writes SVG data to the client.
 * The server must set the Content-Type header like this:
 * w.Header(). Set(" Content-Type", "image/ svg + xml")
 * (This step was not required in the Lissajous example because the server uses standard heuristics
 * to recognize common formats like PNG from the first 512 bytes of the response, and generates the proper header.)
 * Allow the client to specify values like height, width, and color as HTTP request parameters.
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
	http.HandleFunc("/surface", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	//u, err := url.Parse(r.URL.String())

	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w)
}

func surface(out io.Writer) {
	// TODO: sfmt.Fprintf would have been smarter...
	out.Write([]byte(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width: 0.7' " +
		"width='%d' height='%d'>\n", width, height)))

	for i := 0; i <  cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if !(math.IsInf(ax, 0) ||
				math.IsInf(ax, 0) ||
				math.IsInf(ay, 0) ||
				math.IsInf(bx, 0) ||
				math.IsInf(by, 0) ||
				math.IsInf(cx, 0) ||
				math.IsInf(cy, 0) ||
				math.IsInf(dx, 0) ||
				math.IsInf(dy, 0)) {
				out.Write([]byte(fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)))
			}

		}
	}
	out.Write([]byte(fmt.Sprintln("</svg>")))

}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x,y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)

	return math.Sin(r) / r
}
