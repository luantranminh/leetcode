package main

import "fmt"

func shuffle(nums []int, n int) []int {
	for i := 0; i < 2*n; i++ {
		var desireIdx int = i
		for nums[i] > 0 {
			if desireIdx >= n {
				desireIdx = 2*(desireIdx-n) + 1
			} else {
				desireIdx = 2 * desireIdx
			}
			nums[i], nums[desireIdx] = nums[desireIdx], -nums[i]

		}
	}

	for i, _ := range nums {
		nums[i] = -nums[i]
	}

	return nums
}

// expected: [0,5,1,6,2,7,3,8,4,9]
// input:    [0,1,2,3,4,5,6,7,8,9]
// 0: [0,1,2,3,4,5,6,7,8,9]
// 0: [0,2,-1,3,4,5,6,7,8,9]
// 0: [0,4,-1,3,-2,5,6,7,8,9]
// 0: [0,8,-1,3,-2,5,6,7,-4,9]
// 0: [0,5,-1,3,-2,-7,6,-8,-4,9]
// 0: [0,-5,-1,-6,-2,-7,-3,-8,-4,-9]

func main() {
	fmt.Println(shuffle([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
}
