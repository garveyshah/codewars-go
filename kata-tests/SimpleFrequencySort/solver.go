package simplefrequencysort

func Solve(arr []int) []int {
	// step 1 : Count occurences of each 
	counter := make(map[int]int)

	for _, num := range arr {
		counter[num]++
	}

	sort.Sort()
	
}

// Step 1: Count occurrences of each number
frequencyMap := make(map[int]int)
for _, num := range arr {
	frequencyMap[num]++
}

// Step 2: Implement a bubble sort for custom sorting
for i := 0; i < len(arr); i++ {
	for j := 0; j < len(arr)-i-1; j++ {
		// Compare based on frequency
		if frequencyMap[arr[j]] < frequencyMap[arr[j+1]] ||
			(frequencyMap[arr[j]] == frequencyMap[arr[j+1]] && arr[j] > arr[j+1]) {
			// Swap arr[j] and arr[j+1]
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}

return arr