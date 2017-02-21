package main

import (
	"net/http"
	"log"
	"io"
	"image/gif"
	"image"
	"math"
	"image/color"
	"math/rand"
	"net/url"
	"fmt"
	"strconv"
)

/*
 * Exercise 1.12: Modify the Lissajous server to read parameter values from the URL.
 * For example, you might arrange it so that a URL like http:// localhost: 8000/? cycles = 20 sets
 * the number of cycles to 20 instead of the default 5.
 *
 * Use the strconv.Atoi function to convert the string parameter into an integer.
 * You can see its documentation with go doc strconv.Atoi.
 */

func main() {
	http.HandleFunc("/lissajous", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	cycles := 30
	if err != nil {
		fmt.Println(err)
	} else {
		m, _ := url.ParseQuery(u.RawQuery)
		c, _ := strconv.Atoi(m["cycles"][0])
		cycles = c

		}

	lissajous(w, cycles)
}

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	WhiteIndex = 0
	blackIndex = 1
)
func lissajous(out io.Writer, cycles int) {
	const (
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}