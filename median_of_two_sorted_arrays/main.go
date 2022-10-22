package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	target := len(nums1) + len(nums2)

	currentA := 0
	currentB := 0

	last, secondlast := 0, 0

	for {
		if currentA+currentB > target/2 {
			if float64(target/2) != float64(target)/2 {
				return float64(last)
			} else {
				return float64(last+secondlast) / 2
			}
		}

		if currentA >= len(nums1) {
			last, secondlast = nums2[currentB], last
			currentB++
		} else if currentB >= len(nums2) {
			last, secondlast = nums1[currentA], last
			currentA++
		} else if nums1[currentA] <= nums2[currentB] {
			last, secondlast = nums1[currentA], last
			currentA++
		} else {
			last, secondlast = nums2[currentB], last
			currentB++
		}
	}
}

func main() {
	var res = 0.
	res = findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	println(res)
	res = findMedianSortedArrays([]int{1, 3}, []int{2})
	println(res)
	res = findMedianSortedArrays([]int{}, []int{1})
	println(res)
	res = findMedianSortedArrays([]int{}, []int{2, 3})
	println(res)
}
