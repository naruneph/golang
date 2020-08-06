package surface

import (
	"fmt"
	"io"
	"math"
)

// const (
// 	width, height = 600, 320            // canvas size in pixels
// 	cells         = 100                 // number of grid cells
// 	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
// 	xyscale       = width / 2 / xyrange // pixels per x or y unit
// 	zscale        = height * 0.4        // pixels per z unit
// 	angle         = math.Pi / 6         // angle of x, y axes (=30°)
// )

func Render(w io.Writer, width, height, cells int, xyrange, xyscale, zscale, angle float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height, cells, xyrange, xyscale, zscale, angle)
			bx, by := corner(i, j, width, height, cells, xyrange, xyscale, zscale, angle)
			cx, cy := corner(i, j+1, width, height, cells, xyrange, xyscale, zscale, angle)
			dx, dy := corner(i+1, j+1, width, height, cells, xyrange, xyscale, zscale, angle)

			if checkIfFinite(ax, ay, bx, by, cx, cy, dx, dy) {
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}

		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j, width, height, cells int, xyrange, xyscale, zscale, angle float64) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*math.Cos(angle)*xyscale
	sy := float64(height)/2 + (x+y)*math.Sin(angle)*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r

}

func isFinite(f float64) bool {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return false
	}
	return true
}

func checkIfFinite(floats ...float64) bool {
	for _, f := range floats {
		if !isFinite(f) {
			return false
		}
	}
	return true
}
