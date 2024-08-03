// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strconv"
)

// !+
const (
	SHA256 = iota
	SHA384
	SHA512
)

func compare(s1, s2 [32]byte) int{
	var res int = 0
	for i := range 32{
		if s1[i] == s2[i]{
			res++
		}
	}
	return res
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n%d\n", c1, c2, c1 == c2, c1, compare(c1, c2))
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
	hashType := SHA256
	strPos := 1
	switch len(os.Args) {
	case 1:
		return
	case 3:
		hashType, _ = strconv.Atoi(os.Args[1])
		strPos = 2
	}
	switch hashType {
	case SHA384:
		fmt.Printf("%x\n", sha512.Sum384([]byte(os.Args[strPos])))
	case SHA512:
		fmt.Printf("%x\n", sha512.Sum512([]byte(os.Args[strPos])))
	default:
		fmt.Printf("%x\n", sha256.Sum256([]byte(os.Args[strPos])))
	}
}

//!-
