package main

import (
	"fmt"
	"strconv"
)

func main() {
	msg := "masterpiece"
	nom := "nomoretears"
	cd := []int{14, 10, 22, 29, 6, 27, 19, 18, 6, 12, 8}
	cde := []int{15, 17, 14, 17, 19, 7, 21, 7, 2, 20, 20}

	fmt.Println("key for nomoretears is ", MissingKey(nom, cde))
	fmt.Println("key for masterpiece is ", MissingKey(msg, cd))
}

func MissingKey(msg string, code []int) int {
	keySlice := []int{}
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, char := range msg {
		for i, letter := range alphabet {
			if i == -1 && char == letter{
				continue
			}
				num := code[1]
				num1 := num - (i + 1)
				fmt.Println(num1)
				keySlice = append(keySlice, num1)

			}
		}
		keystring := ""

	for _, digit := range keySlice {
		keystring += strconv.Itoa(digit)
	}

	key, err := strconv.Atoi(fmt.Sprint(keySlice))
	if err != nil {
		fmt.Println(err)
	}
	return key
	}
	
