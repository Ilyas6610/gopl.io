package main

import (
	"fmt"
	"math"
	"time"
)

type Fahrenheit float64
type Celcius float64

func (c Celcius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func FToC(f Fahrenheit) Celcius {
	return Celcius((f - 32) * 5 / 9)
}

func CToF(c Celcius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	begin := time.Now()
	res := int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
	fmt.Println(time.Since(begin).Nanoseconds())
	return res
}

func PopCount2(x uint64) int {
	begin := time.Now()
	var res int
	for i := 0; i < 8; i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	fmt.Println(time.Since(begin).Nanoseconds())
	return res
}

func PopCount3(x uint64) int {
	begin := time.Now()
	var res int
	for i := 0; i < 64; i++ {
		res += int((x >> i) & 1)
	}
	fmt.Println(time.Since(begin).Nanoseconds())
	return res
}

func PopCount4(x uint64) int {
	begin := time.Now()
	var res int
	for x != 0 {
		res += 1
		x = x & (x - 1)
	}
	fmt.Println(time.Since(begin).Nanoseconds())
	return res
}

type PopCountFunc func(uint64)

func main() {
	// for _, v := range os.Args[1:] {
	// 	f, err := strconv.ParseFloat(v, 64)
	// 	if err != nil {
	// 		fmt.Printf("%v", err)
	// 	}
	// 	fmt.Printf("Celsius: %s\nCToF: %s\nFahrenheit: %s\nFToC: %s\n", Celcius(f), CToF(Celcius(f)), Fahrenheit(f), FToC(Fahrenheit(f)))
	// }
	fmt.Println(PopCount(math.MaxUint64), PopCount2(math.MaxUint64), PopCount3(math.MaxUint64), PopCount4(math.MaxUint64))
}
