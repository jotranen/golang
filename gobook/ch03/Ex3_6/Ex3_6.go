// TODO

package main

import (
	"image"
	"image/png"
	"os"
	"image/color"
	"math/cmplx"
)

/*
 * Exercise 3.6: Supersampling is a technique to reduce the effect of pixelation by computing the color value
 * at several points within each pixel and taking the average. The simplest method is to divide each pixel
 * into four “subpixels.” Implement it.
 */

func main() {
	const(
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	fifo := make([]int, 4)

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z, fifo))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128, fifo []int) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//col :=
			return color.RGBA{ 0x00, 0x00, 0x00 + contrast*n, 0xFF - contrast*n }
		}
	}

	return color.Black
}
