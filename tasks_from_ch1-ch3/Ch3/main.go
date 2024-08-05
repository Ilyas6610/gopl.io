package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func serverRun() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(writeSvg()))
}

func writeSvg() string {
	res := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, errA := corner(i+1, j)
			bx, by, errB := corner(i, j)
			cx, cy, errC := corner(i, j+1)
			dx, dy, errD := corner(i+1, j+1)
			dist := math.Hypot(ax, ay)
			color := int64((float64(dist/600)*0xff))<<16 + int64(float64((600-dist)/600)*0xff)
			if errA != nil || errB != nil || errC != nil || errD != nil {
				continue
			}
			res += fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' fill='#%x' />\n", ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	res += fmt.Sprintf("</svg>")
	return res
}

func corner(i, j int) (float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	if math.IsInf(z, 0) {
		return 0, 0, errors.New("math: inf value")
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func mandelbrotRun() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py += 2 {
		y1 := float64(py)/height*(ymax-ymin) + ymin
		y2 := float64(py+1)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px += 2 {
			x1 := float64(px)/width*(xmax-xmin) + xmin
			x2 := float64(px)/width*(xmax-xmin) + xmin
			z := (complex(x1, y1) + complex(x2, y2) + complex(x1, y2) + complex(x2, y1)) / 4
			// Image point (px, py) represents complex value z.
			zval := mandelbrot(z)
			img.Set(px, py, zval)
			img.Set(px+1, py, zval)
			img.Set(px, py+1, zval)
			img.Set(px+1, py+1, zval)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return newton(z)
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//
//	= z - (z^4 - 1) / (4 * z^3)
//	= z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}

func main() {
	// serverRun()
	mandelbrotRun()
}
