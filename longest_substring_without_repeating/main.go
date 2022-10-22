package main

func lengthOfLongestSubstring(s string) int {
	var max = 0
	var current = 0

	var starting = 0

	m := map[int32]int{}

	for index := 0; index < len(s); index++ {
		char := int32(s[index])
		if _, ok := m[char]; !ok {
			current++
			m[char] = index
		} else {
			current = 0
			if starting == m[char] {
				index = m[char]
			} else {
				index = m[char] - 1
			}

			m = map[int32]int{}
			starting = index + 1
		}
		if current > max {
			max = current
		}
	}
	return max
}

func main() {
	println(lengthOfLongestSubstring("dvdf"))
	println(lengthOfLongestSubstring("abcabcbb"))
}
