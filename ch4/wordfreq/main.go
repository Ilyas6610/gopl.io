package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordCounter := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		wordCounter[word]++
	}
	for word, count := range wordCounter {
		fmt.Printf("%s: %d\n", word, count)
	}
}
