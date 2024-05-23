package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "masterplan"
	key := 1939
	fmt.Println(Encode(str, key))
}

func Encode(str string, key int) []int {
	cipher := []int{}
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	numstr := strconv.Itoa(key)
	keyIndex := 0

	for _, char := range str {
		for i, letter := range alphabet {
			if char == letter {
				num := int(numstr[keyIndex]- '0')
				crypt := num + i + 1
				cipher = append(cipher, crypt)

				keyIndex = (keyIndex + 1) % len(numstr)
				break
			}
		}
	}
	return cipher
}
