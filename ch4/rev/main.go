// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"
	reverseArr(&a)
	fmt.Println(a)
	//!-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	rotate(s)
	fmt.Println(s)
	//!-slice

	sl := []string{"1", "2", "2", "2", "2", "4", "5", "6", "6", "6"}
	s1 := removeDupsInRow(sl)
	fmt.Println(s1)

	byteSl := []byte{'1', '2', '2', '2', ' ', ' ', ' ', '4', '4'}
	fmt.Println(string(removeSpacesInRow(byteSl)))
	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		reverse(ints)
		fmt.Printf("%v\n", ints)
	}
	// NOTE: ignoring potential errors from input.Err()
}

// !+rev
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev

const arrSize = 6

func reverseArr(arr *[arrSize]int) {
	for i := 0; i < arrSize/2; i++ {
		arr[i], arr[arrSize-1-i] = arr[arrSize-1-i], arr[i]
	}
}

func rotate(s []int) {
	val := s[0]
	for i := 0; i < len(s)-1; i++ {
		s[i] = s[i+1]
	}
	s[len(s)-1] = val
}

func removeDupsInRow(s []string) []string {
	firstVal := 0
	i := 1
	for {
		firstVal = i
		for i < len(s) && s[i] == s[i-1] {
			i++
		}
		if i == len(s) {
			return s[:firstVal]
		}
		copy(s[firstVal:], s[i:])
		i = firstVal + 1
	}
}

func removeSpacesInRow(s []byte) []byte {

	firstVal := 0
	size := len(s)
	i := 1
	for {
		firstVal = i
		for i < len(s) && s[i] == s[i-1] && unicode.IsSpace(rune(s[i])) {
			i++
		}
		if i == len(s) {
			return s[:size]
		}
		copy(s[firstVal:], s[i:])
		size -= (i - firstVal)
		i = firstVal + 1
	}
}
