package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, left := range nums {
		for j, right := range nums {
			if left+right == target && i != j {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)

	fmt.Println("Indices:", result)
}
