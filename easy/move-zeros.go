package main

import "fmt"

func moveZeroes(nums []int) {
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
}

func main() {
	a := &([]int{0, 1, 0, 3, 12})
	moveZeroes(*a)
	fmt.Println(a)
}
