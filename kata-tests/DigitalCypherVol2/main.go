package main

import (
	"fmt"
	"strconv"
)
func main() {
	code := []int{8, 24, 7, 32, 10, 23}
	key := 1939

	fmt.Println(Decode(code, key))
}

func Decode(code []int, key int) string{
	plainText := []rune{}
	keyIndex := 0

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	keyStr := strconv.Itoa(key)

	for _, num := range code {
		digit := int(keyStr[keyIndex]-'0')
		crypt := num - digit

		keyIndex = (keyIndex + 1) % len(keyStr)
		
	for i, letter := range alphabet {
			if crypt == i+1 {
				plainText = append(plainText, letter)
			}
		}
	}
	return string(plainText)
}
