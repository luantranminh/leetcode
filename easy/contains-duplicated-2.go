package main

import "fmt"

func main() {

	fmt.Print(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
}

func containsNearbyDuplicate(nums []int, k int) bool {

	slide := make(map[int]bool, k)

	for i := 0; i < len(nums); i++ {

		if len(nums) > k {
			firstPositionInSlide := i - k - 1
			if firstPositionInSlide >= 0 && firstPositionInSlide < len(nums) {
				delete(slide, nums[i-k-1])

			}
		}

		if _, ok := slide[nums[i]]; ok {
			return true
		}

		slide[nums[i]] = true

	}

	return false
}
