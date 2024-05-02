package main

import "fmt"

func main() {
	a := []string{"NORTH", "SOUTH", "EAST", "WEST"}
	arr := []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}
	arr1 := []string{"NORTH", "WEST", "SOUTH", "EAST"}
	arr2 := []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}

	fmt.Println(DirReduc(a))
	fmt.Println(DirReduc(arr))
	fmt.Println(DirReduc(arr1))
	fmt.Println(DirReduc(arr2))
}

func DirReduc(arr []string) []string {
	if len(arr) == 0 {
		return []string{}
	}
	library := map[string]string{"NORTH": "SOUTH", "SOUTH": "NORTH", "EAST": "WEST", "WEST": "EAST", "East": "West", "West": "East", "North": "South", "South": "North"}
	result := []string{}

	for _, direc := range arr {
		if len(result) > 0 && library[direc] == result[len(result)-1] {
			result = result[:len(result)-1]
		} else {
			result = append(result, direc)
		}
	}

	return result
}
