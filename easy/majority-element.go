package main

import "fmt"

func majorityElement(nums []int) int {
	candidate := nums[0]
	counter := 0

	for _, v := range nums {

		if counter == 0 {
			candidate = v
		}

		if counter == 0 {
			counter++
		} else {
			counter--
		}
	}

	return candidate
}

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
