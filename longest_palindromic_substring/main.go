package main

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
