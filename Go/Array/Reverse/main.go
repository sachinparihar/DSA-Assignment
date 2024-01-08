package main

import (
	"fmt"
)

func reverseArray(arr []int) {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

func main() {
	array := []int{1, 2, 3, 4, 5}
	fmt.Println("Original array:", array)

	reverseArray(array)
	fmt.Println("Reversed array:", array)
}
