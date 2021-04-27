// Server2 is a minimal "echo" and counter server.
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
	"time"
)

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

var palette = []color.Color{color.White, color.Black}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(cycles_arg float64, out io.Writer) {

	/* 一个常量的枚举一样的东西 */
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	var cycles float64 // number of complete x oscillator revolutions

	if cycles_arg != 0 {
		cycles = cycles_arg
	} else {
		cycles = 5
	}
	/* 随机浮点数 */
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	/* 指定GIF的LoopCount属性为nframes */
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	// for k, v := range r.Header {
	// 	fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	// }
	// fmt.Fprintf(w, "Host = %q\n", r.Host)
	// fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	// err := r.ParseForm()
	// if err != nil {
	// 	log.Print(err)
	// }
	var cycles_arg float64
	for k, v := range r.Form {
		// fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		if k == "cycles" {
			for i := 0; i < len(v); i++ {
				cycles_arg, _ = strconv.ParseFloat(v[i], 64)
			}
		} else {
			cycles_arg = 0.0
		}
	}
	lissajous(cycles_arg, w)
}
