package main

import "fmt"

func findHighestRankingDuplicate(nums []int) int {
	highestRankingIndex := -1
	highestRankingValue := 0

	// Iterate through the array
	for i, num := range nums {
		// Check if the current number is equal to any subsequent number in the array
		for j := i + 1; j < len(nums); j++ {
			if num == nums[j] {
				// If a duplicate is found, update the highest ranking index if needed
				if j > highestRankingIndex {
					highestRankingIndex = j
					highestRankingValue = num
				}
			}
		}
	}

	return highestRankingValue
}

func main() {
	// Sample integer array with random numbers
	randomNumbers := []int{155, 242, 252, 14, 71, 191, 191, 148, 246, 118, 19, 28, 28, 16, 143}

	// Find the highest ranking duplicate number
	highestRankingDuplicate := findHighestRankingDuplicate(randomNumbers)

	// Print the highest ranking duplicate number
	if highestRankingDuplicate != 0 {
		fmt.Printf("The highest ranking duplicate number is: %d\n", highestRankingDuplicate)
	} else {
		fmt.Println("No duplicates found.")
	}
}
