package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Write strings to compare")
		return
	}
	s := os.Args[1:]
	if isAnagram(s[0], s[1]) {
		fmt.Printf("%s and %s are anagrams", s[0], s[1])
	} else {
		fmt.Printf("%s and %s are not anagrams", s[0], s[1])
	}
}

func isAnagram(s1, s2 string) bool {
	m := make(map[rune]int)
	if len(s1) != len(s2) {
		return false
	}
	for _, c := range s1 {
		m[c]++
	}
	for _, c := range s2 {
		m[c]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
