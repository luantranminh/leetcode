package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	results := make([]bool, len(candies))

	if len(candies) == 0 {
		return []bool{}
	}

	greatest := candies[0]
	for _, v := range candies {
		if v > greatest {
			greatest = v
		}
	}

	for i, v := range candies {
		results[i] = v+extraCandies >= greatest
	}
	return results
}

func main() {
	fmt.Println(kidsWithCandies([]int{1, 1, 5, 1}, 3))
}
