package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Solution: %+v \n", twoSum([]int{3, 2, 4}, 6))
}
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for r := i + 1; r < len(nums); r++ {
			if target-nums[r]-nums[i] == 0 {
				return []int{i, r}
			}
		}
	}
	panic("invalid solution")
}
