// измените сервер с фигурами Лиссажу так, чтобы значения параметров считывались из URL.
// Например, URL вида http://localhost:8000/?cycles=20 устанавливает количество циклов равным 20 вместо значения по умолчанию, раного 5.
// Используйте функцию strconv.Atoi для преобразования строгового параметра в целое число.
// Просмотреть документацию по данной функции можно с помощью команды go doc strconv.Atoi.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := map[string]int{
		"cycles":  5,
		"size":    100,
		"nframes": 64,
		"delay":   8,
	}
	for param := range params {
		if r.FormValue(param) != "" {
			params[param], _ = strconv.Atoi(r.FormValue(param))
		}
	}
	lissajous(w, params)
}

func lissajous(out io.Writer, params map[string]int) {
	cycles := params["cycles"]
	res := 0.001
	size := params["size"]
	nframes := params["nframes"]
	delay := params["delay"]
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
