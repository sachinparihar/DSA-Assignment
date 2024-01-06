package main

import "fmt"

func moveZeroes(nums []int) {
	zeroIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[zeroIndex], nums[i] = nums[i], nums[zeroIndex]
			zeroIndex++
		}
	}
}

func main() {
	nums1 := []int{0, 1, 0, 3, 12}
	moveZeroes(nums1)
	fmt.Println(nums1)

	nums2 := []int{0, 3, 2, 0, 3, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	moveZeroes(nums2)
	fmt.Println(nums2)
}
