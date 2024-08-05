package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// stdio
func P1() {
	begin := time.Now()
	// var s string
	// for _, str := range os.Args[1:] {
	// 	s += str + " "
	// 	//fmt.Println(idx, " "+str)
	// }
	// fmt.Println(s)
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(begin.Nanosecond() - time.Now().Nanosecond())
}

func count(input *bufio.Scanner) int {
	m := make(map[string]int)

	for input.Scan() {
		m[input.Text()]++
	}

	for _, v := range m {
		if v > 1 {
			return 1
		}
	}
	return 0
}

// stdio
func P2() {
	if len(os.Args) == 1 {
		count(bufio.NewScanner(os.Stdin))
	} else {
		for _, filename := range os.Args[1:] {
			fmt.Println("Openning file: ", filename)
			file, err := os.Open(filename)
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}
			if count(bufio.NewScanner(file)) == 1 {
				fmt.Println(filename)
			}
			file.Close()
		}
	}

}

// gif
func P3(w io.Writer, c int) {
	pallete := []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}
	const (
		WhiteIndex = 0
		BlackIndex = 1
	)

	cycles := c
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	index := BlackIndex
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(index%256))
		}
		phase += 0.1
		pallete = append(pallete, color.RGBA{uint8(rand.Uint32() % 256), uint8(rand.Uint32() % 256), uint8(rand.Uint32() % 256), 0xFF})
		index++
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim)
}

// fetch
func P4() {
	for _, address := range os.Args[1:] {
		if !strings.HasPrefix(address, "http://") {
			address = "http://" + address
		}
		resp, err := http.Get(address)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp.Status)
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}

// goroutines
func P5_1() {
	ch := make(chan string)
	t := time.Now()
	for _, address := range os.Args[1:] {
		go P5_2(address, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Println(time.Since(t).Seconds())
}

func P5_2(address string, ch chan string) {
	t := time.Now()
	resp, err := http.Get(address)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	ch <- fmt.Sprintf("%s %fs\n", resp.Request.URL.Host, time.Since(t).Seconds())
	resp.Body.Close()
}

var mu sync.Mutex
var cnt int

func P6_1_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
}

func P6_1() {
	http.HandleFunc("/", P6_1_handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func P6_2_handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	cnt++
	mu.Unlock()
	fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
}

func P6_2_counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", cnt)
	mu.Unlock()
}

func P6_2_lissajous(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Print()
	val, _ := strconv.Atoi(r.Form["count"][0])
	P3(w, val)
}

func P6_2() {
	http.HandleFunc("/", P6_2_handler)
	http.HandleFunc("/count", P6_2_counter)
	http.HandleFunc("/lissajous", P6_2_lissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func main() {
	// P1()
	// P2()
	// P3()
	// P4()
	// P5_1()
	// P6_1()
	P6_2()
}
