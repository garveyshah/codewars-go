package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Enter a Threshold")
	}

	num, err := CustomAtoi(os.Args[1])
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(HammingNumber(num))

}

func HammingNumber(n int) uint {
	if n <= 0 {
		return 0 // or handle invalid input as needed
	}

	HammingNums := []uint{1}

	i, j, k := 0, 0, 0
	var next1, next2, next3 uint

	for uint(len(HammingNums)) < uint(n) {
		// Calculate the next potential Hamming numbers.
		next1 = HammingNums[i] * 2
		next2 = HammingNums[j] * 3
		next3 = HammingNums[k] * 5

		//Find the minimum of the next potenial Hamming numbers
		nextHamming, _ := Minimum(next1, next2, next3)
		HammingNums = append(HammingNums, nextHamming)

		// Increment the index for the smallest number that was chosen
		if nextHamming == next1 {
			i++
		}
		if nextHamming == next2 {
			j++
		}
		if nextHamming == next3 {
			k++
		}
	}
	return HammingNums[n-1]
}

// Helper function to find the minimum of three numbers.
func Minimum(num ...uint) (uint, error) {
	if len(num) == 0 {
		return 0, fmt.Errorf("no numbers provided")
	}
	min := num[0]
	for _, num := range num[1:] {
		if num < min {
			min = num
		}
	}
	return min, nil
}

// Helper function to convert digit strings to intergers.
func CustomAtoi(s string) (int, error) {
	var result int
	var sign bool

	if s[0] == '-' {
		sign = true
		s = s[1:]
	}

	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, fmt.Errorf("invalid character %v", string(char))
		}
		result = result*10 + int(char-'0')
	}

	if sign {
		result *= -1
	}
	return result, nil
}
