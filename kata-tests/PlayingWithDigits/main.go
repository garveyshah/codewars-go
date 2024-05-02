package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	n := 695
	p := 2
	fmt.Println(DigPow(n, p))

}
func DigPow(n, p int) int {
	b := strconv.Itoa(n)
	result := 0
	for i, num := range b {
		num2, _ := strconv.Atoi(string(num))
		result += int(math.Pow(float64(num2), float64(p+i)))
	}
	if result%n == 0 {
		return result / n
	} else {
		return -1
	}
}
