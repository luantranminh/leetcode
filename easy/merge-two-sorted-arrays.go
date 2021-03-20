// https://leetcode.com/problems/merge-sorted-array/
package main

import "fmt"

func merge(a1, a2 []int) []int {
	result := make([]int, len(a1)+len(a2))

	i1 := 0
	i2 := 0
	ri := 0
	for i1 < len(a1) || i2 < len(a2) {

		if i1 >= len(a1) {
			result[ri] = a2[i2]
			ri++
			i2++
			continue
		}

		if i2 >= len(a2) {
			result[ri] = a1[i1]
			ri++
			i1++
			continue
		}

		if a1[i1] < a2[i2] {
			result[ri] = a1[i1]
			ri++
			i1++
		}
		if a1[i1] >= a2[i2] {
			result[ri] = a2[i2]
			ri++
			i2++
		}

	}

	return result
}

// while a1 < a2, result push a1
// while a2 <= a1, result push a2
// repeat until both run out of range

func main() {
	a1 := []int{1, 1, 1, 1, 1, 3, 4}
	a2 := []int{2, 2, 3}

	merged := merge(a1, a2)

	fmt.Println(merged)
}
