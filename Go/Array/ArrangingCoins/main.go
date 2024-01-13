package main

import (
	"fmt"
)

func arrangeCoins(n int) int {
	left, right := 0, n

	for left <= right {
		mid := left + (right-left)/2
		totalCoins := (mid * (mid + 1)) / 2

		if totalCoins == n {
			return mid
		} else if totalCoins < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

func main() {
	fmt.Println(arrangeCoins(5))
}
