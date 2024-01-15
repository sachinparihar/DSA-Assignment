package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	// Create hash sets to store unique elements
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	for _, num := range nums1 {
		set1[num] = true
	}

	result := []int{}
	for _, num := range nums2 {
		if set1[num] && !set2[num] {
			result = append(result, num)
			set2[num] = true
		}
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	result := intersection(nums1, nums2)
	fmt.Println(result)

	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	result = intersection(nums1, nums2)
	fmt.Println(result)
}
