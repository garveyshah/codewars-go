package main

import "fmt"

func main() {
	randomNums := []int{12, 10, 8, 12, 7, 6, 4, 10, 10}
	duplicates := HighestRank(randomNums)
	fmt.Println("Duplicate numbers:", duplicates)

	randomNumbers := []int{3, 5, 2, 7, 5, 8, 2, 3, 9, 7}
	duplicate := HighestRank(randomNumbers)
	fmt.Println("Duplicate numbers:", duplicate)
}

func HighestRank(nums []int) int {
	newNums := []int{}

	// loop over the range of nums and look for frequent integers
	// loop over the frequent and find the highet rank int
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num == nums[j] {
				newNums = append(newNums, num)
				break
			}
		}
	}
	sortNums := []int{}
	for i := 0; i < len(newNums); i++ {
		if len(sortNums) > 0 && newNums[i] < sortNums[len(sortNums)-1] {
			sortNums = sortNums[len(sortNums)-1:]
		} else {
			sortNums = append(sortNums, newNums[i])
		}
	}

	return sortNums[len(sortNums)-1]
}
