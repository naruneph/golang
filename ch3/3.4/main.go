package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"math"
	"myproj/surface"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "surface.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		//fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		width, err := strconv.ParseInt( r.PostForm.Get("width"), 10, 64); 
		if err != nil {
			width = 600
		}

		height, err := strconv.ParseInt( r.PostForm.Get("height"), 10, 64);
		if err != nil {
			height = 320
		}

		cells, err := strconv.ParseInt( r.PostForm.Get("cells"), 10, 64);
		if err != nil {
			cells = 100
		}

		xyrange, err := strconv.ParseFloat( r.PostForm.Get("xyrange"), 64);
		if err != nil {
			xyrange = 30.0
		}

		xyscale, err := strconv.ParseFloat( r.PostForm.Get("xyscale"), 64);
		if err != nil {
			xyscale = float64(width) / 2 / xyrange 
		}

		zscale, err := strconv.ParseFloat( r.PostForm.Get("zscale"), 64);
		if err != nil {
			zscale = float64(height) * 0.4  
		}

		angle, err := strconv.ParseFloat( r.PostForm.Get("angle"), 64);
		if err != nil {
			angle = math.Pi / 6 
		}


		w.Header().Set("Content-Type", "image/svg+xml")
		surface.Render(w, int(width), int(height), int(cells), xyrange, xyscale, zscale, angle)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

