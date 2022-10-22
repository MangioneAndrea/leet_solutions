package main

import "sort"

// One could do the naive implementation, loop all letters in the string and iterate left and right to see for which level it is a palindrome
// That solution would run into O(n^2) time complexity

// Let's try to stay at O(n log n) time complexity. The space should stay at O(n)
// The idea is to put all letters into a map pointing to which indexes they are in the string (sorted) time: O(n) space: O(n)
// Then we can just forget about the string. We can just take each array, and start watching for a range which keeps to be palindrome
// If the longest path is a palindrome, there are no possible longer palindromes

// Example "aabbbabc" (Good case)
// map = { a: [0, 1, 5], b: [2, 3, 4, 6], c: [7] }

// check a as extremes:
//    0 - 5. Now to be a palindrome, the following indexes should be in the same arrays. 1 is in a, but 4 is not. This is not a palindrome
//    0 - 1. Left and right touched. This is a palindrome, length 2
//    1 - 5. Check 2 is in b, so is 4. 3 is the middle, so this is a palindrome, length 5
// check b as extremes:
//    2 - 6. The distance is 5, no need to check.
//    all other combinations are smaller than 2 - 6, so no need to check
// check c as extremes:
//    Only 1 element, nothing to check
// Result 5, Good. Very fast

// Example "abcadbcad" (Bad case, no palindromes)
// map = { a: [0, 4, 8], b: [1, 5], c: [2, 6], d: [3, 7] }

// check a as extremes:
//    0 - 8. 1 is in b, but 7 is not. This is not a palindrome
//    0 - 4. 1 is in b, but 3 is not. This is not a palindrome
//    4 - 8. 5 is in b, but 7 is not. This is not a palindrome
// check b as extremes:
//    1 - 5. 2 is in c, but 4 is not. This is not a palindrome
// check c as extremes:
//    2 - 6. 3 is in d, but 5 is not. This is not a palindrome
// check d as extremes:
//    3 - 7. 4 is in b, but 6 is not. This is not a palindrome
// Result 1, as the string is bigger than 0. Very fast

// Example "abcdefgedcba" (Very bad case, no palindrome, hard to catch)
// map = { a: [0, 11], b: [1, 10], c: [2, 9], d: [3, 8], e: [4, 7], f: [5], g: [6] }
// check a as extremes:
//    0 - 11. 1 and 10 are in b, 2 and 9 are in c, 3 and 8 are in d, 4 and 7 are in e, 5 is in f but g is not. This is not a palindrome.
// check b as extremes:
//    1 - 10. 2 and 9 are in c, 3 and 8 are in d, 4 and 7 are in e, 5 is in f but g is not. This is not a palindrome.
// check c as extremes:
//    2 - 9. 3 and 8 are in d, 4 and 7 are in e, 5 is in f but g is not. This is not a palindrome.
// check d as extremes:
//    3 - 8. 4 and 7 are in e, 5 is in f but g is not. This is not a palindrome.
// check e as extremes:
//    4 - 7. 5 is in f but g is not. This is not a palindrome.
// check f as extremes:
//    Only 1 element, nothing to check
// check g as extremes:
//    Only 1 element, nothing to check
// result 1, as the string is bigger than 0. Not fast at all.
// This is why we can implement a cache. This will increase the space complexity by O(n), which still is O(n) in the end.

func longestPalindrome(s string) string {
	var lettersArrays = map[int32][]int{}
	var lettersMap = map[int32]map[int32]bool{}
	var isPalindrome = map[int32]map[int32]bool{}

	for i, char := range s {
		isPalindrome[int32(i)] = map[int32]bool{}
		if _, ok := lettersArrays[char]; ok {
			lettersArrays[char] = append(lettersArrays[char], i)
			lettersMap[char][int32(i)] = true
		} else {
			lettersArrays[char] = []int{i}
			lettersMap[char] = map[int32]bool{int32(i): true}
		}
	}

	for _, ints := range lettersArrays {
		sort.Ints(ints)
	}

	var maxPalindrome = 0
	var palindrome = ""

	for _, ints := range lettersArrays {
		if len(ints) == 1 {
			if maxPalindrome == 0 {
				maxPalindrome = 1
				palindrome = s[ints[0] : ints[0]+1]
			}
			continue
		}

		for leftIndex, left := range ints {
			for rightIndex := len(ints) - 1; rightIndex > leftIndex; rightIndex-- {
				right := ints[rightIndex]
				if res, checked := isPalindrome[int32(left)][int32(right)]; checked {
					if res {
						if right-left+1 > maxPalindrome {
							maxPalindrome = right - left + 1
							palindrome = s[left : right+1]
						}
					}
					break
				}
				if right-left+1 <= maxPalindrome {
					break
				}
				isPal := true
				for i := 1; i <= (right-left)/2; i++ {
					l, r := int32(left+i), int32(right-i)
					if !lettersMap[int32(s[l])][r] {
						isPal = false
						break
					}
				}
				if isPal {
					maxPalindrome = right - left + 1
					palindrome = s[left : right+1]
					break
				}
				isPalindrome[int32(left)][int32(right)] = isPal
			}
		}
	}

	if maxPalindrome == 0 {
		return s[0:1]
	}

	return palindrome
}

func main() {
	println(longestPalindrome("babad"))
	println(longestPalindrome("cbbd"))
	println(longestPalindrome("abcdefgedcba"))
	println(longestPalindrome("abcadbcad"))
	println(longestPalindrome("aabbbabc"))
}
