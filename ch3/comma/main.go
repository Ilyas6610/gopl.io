// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"bytes"
	"fmt"
	"os"
	"unicode"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
		fmt.Printf("  %s\n", nonrecursive_comma(os.Args[i]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//!-

func nonrecursive_comma(s string) string {
	var buf bytes.Buffer
	n, i := 0, 3-len(s)%3
	if i == 3 {
		i = 0
	}
	for n < len(s) {
		if i == 3 {
			buf.Write([]byte{','})
			i = 0
		}
		if unicode.IsDigit(rune(s[n])) {
			i++
		}
		buf.Write([]byte{s[n]})
		n++
	}
	return buf.String()
}
