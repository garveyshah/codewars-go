package main

import "fmt"

func main() {
	input := "abcde"
	input1 := "abcdef"
	input2 := "ab cd de "
	input3 := "a1b2c34"

	fmt.Println(SplitStrings(input))
	fmt.Println(SplitStrings(input1))
	fmt.Println(SplitStrings(input2))
	fmt.Println(SplitStrings(input3))
}

func SplitStrings(str string) []string {
	var final []string
	for i := 0; i < len(str); i += 2 {
		if i+1 < len(str) {
			final = append(final, str[i:i+2])
		} else {
			final = append(final, str[i:i+1]+"_")
		}
	}
	return final
}
