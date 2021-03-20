package main

import (
	"fmt"
	"math"
)

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		idx := int64(math.Abs(float64(nums[i])) - 1)
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	var disappearedNumbers []int
	for i, v := range nums {
		if v > 0 {
			disappearedNumbers = append(disappearedNumbers, i+1)
		}
	}

	return disappearedNumbers
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
