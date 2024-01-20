// Given the array of integers nums,
// you will choose two different indices i and j of that array.
// Return the maximum value of (nums[i]-1)*(nums[j]-1).

package main

import (
	"fmt"
	"sort"
)

func maxProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return (nums[n-1] - 1) * (nums[n-2] - 1)
}

func main() {
	nums := []int{3, 4, 5, 2}
	fmt.Println(maxProduct(nums))
}
