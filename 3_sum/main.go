package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			two := nums[i] + nums[j]
			target := -two
			if target < nums[j] {
				break
			}
			if k := sort.SearchInts(nums[j+1:], target); k < len(nums[j+1:]) && nums[j+1:][k] == target {
				res = append(res, []int{nums[i], nums[j], target})
			}
		}
	}
	return res
}
func main() {
	fmt.Println(threeSum([]int{3, 0, -2, -1, 1, 2}))
}
