package main

import (
	"fmt"
)

func countBits(b byte) int {
	count := 0
	for count = 0; b > 0; count++ {
		b &= b - 1
	}
	return count
}

func countForString(s string) int {
	count := 0
	for _, b := range s {
		count += countBits(byte(b))
	}
	return count
}

func main() {
	fmt.Println(countForString("abcd")) // 13
}
