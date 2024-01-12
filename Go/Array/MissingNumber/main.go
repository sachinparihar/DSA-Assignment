package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	n := len(nums)
	sum := (n * (n + 1)) / 2
	for _, num := range nums {
		sum -= num
	}
	return sum

}

func main() {
	nums := []int{3, 0, 1, 2, 5}
	fmt.Println(missingNumber(nums)) // Output: 2
}
