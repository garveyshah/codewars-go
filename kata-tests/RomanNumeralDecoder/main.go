package main

import "fmt"

func main() {
	arg := "MM"
	arg1 := "MDCLXVI"
	arg2 := "M"
	arg3 := "CD"
	arg4 := "XC"
	arg5 := "XL"
	arg6 := "I"

	fmt.Println(Decode(arg))
	fmt.Println(Decode(arg1))
	fmt.Println(Decode(arg2))
	fmt.Println(Decode(arg3))
	fmt.Println(Decode(arg4))
	fmt.Println(Decode(arg5))
	fmt.Println(Decode(arg6))
}

func Decode(roman string) int {
	romanLibrary := make(map[string] int) 
		romanLibrary["I"] = 1
		romanLibrary["V"] = 5
		romanLibrary["X"] = 10
		romanLibrary["L"] = 50
		romanLibrary["C"] = 100
		romanLibrary["D"] = 500
		romanLibrary["M"] = 1000
	
		for key, value range romanLibrary {}
fmt.Println(romanLibrary)
	return 0
}