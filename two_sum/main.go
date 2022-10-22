package main

import (
	"fmt"
	"sort"
)

func values(nums []int, target int) (int, int) {
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		val1 := nums[i]
		second := target - val1
		candidate := sort.SearchInts(nums, target-nums[i])

		if candidate != len(nums) && candidate != i && nums[candidate] == second {
			return nums[i], nums[candidate]
		}
	}
	return -1, -1
}

func twoSum(nums []int, target int) []int {
	n := make([]int, len(nums))
	copy(n, nums)
	first, second := values(n, target)

	a := -1
	b := -1
	for i, num := range nums {
		if num == first && a == -1 {
			a = i
			continue
		}
		if num == second {
			b = i
		}
		if a != -1 && b != -1 {
			break
		}
	}
	return []int{a, b}
}

func main() {
	fmt.Println(twoSum([]int{0, 4, 3, 0}, 0))
}
