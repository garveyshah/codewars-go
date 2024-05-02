package main

import "fmt"

func main() {
	randomNums := []int{5,4,3,4,5,4,3,2,1,1,2,2,0,9,8,7,6,6,7,9,9,8}
	duplicates := HighestRank(randomNums)
	fmt.Println("Duplicate numbers:", duplicates)
	
	randomNumbers := []int{3, 5, 2, 7, 5, 8, 2, 3, 9, 7}
    duplicate := HighestRank(randomNumbers)
    fmt.Println("Duplicate numbers:", duplicate)

}

func HighestRank(nums [] int) int {
	numCount := make(map[int]int)

	for _, num := range nums {
		numCount[num]++
	}

	highestRankingValue := 0

	for num, count := range numCount {
		if count > 1 {
			if num > highestRankingValue {
				highestRankingValue = num
			}
		}
	}

	return highestRankingValue
}