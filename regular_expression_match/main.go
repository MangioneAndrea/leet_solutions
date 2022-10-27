package main

import (
	"log"
)

const repeat = '*' // 42

func matchPartial(input, pattern []rune, iIndex, patternIndex int, optional bool) bool {
	// both input and pattern are empty, so everything passed the check
	if patternIndex >= len(pattern) && iIndex >= len(input) {
		return true
	}

	// pattern finished, but the input didn't, so it's wrong
	if patternIndex >= len(pattern) {
		return false
	}

	// input finished, and the pattern is not optional (no way this can be correct), so it's wrong
	if !optional && iIndex >= len(input) {
		return false
	}

	// input and pattern do not match and the pattern is not optional
	if !optional && (pattern[patternIndex] != input[iIndex] && pattern[patternIndex] != '.') {
		return false
	}

	// non-repeating and correct --> go to next char;
	if !optional && (pattern[patternIndex] == input[iIndex] || pattern[patternIndex] == '.') {
		isNextOptional := len(pattern) > patternIndex+2 && pattern[patternIndex+2] == repeat
		return matchPartial(input, pattern, iIndex+1, patternIndex+1, isNextOptional)
	}
	// repeating and wrong --> go to 2 chars ahead
	if iIndex >= len(input) || (pattern[patternIndex] != input[iIndex] && pattern[patternIndex] != '.') {
		isNextOptional := len(pattern) > patternIndex+3 && pattern[patternIndex+3] == repeat
		if matchPartial(input, pattern, iIndex, patternIndex+2, isNextOptional) {
			return true
		}
	} else {
		// repeating character and correct (recursive, as it might look correct but it's not asd == as*d*sd)
		if matchPartial(input, pattern, iIndex+1, patternIndex, true) {
			return true
		}

		isNextOptional := len(pattern) > patternIndex+3 && pattern[patternIndex+3] == repeat
		// it looked correct, but it is wrong, maybe removing the optional pattern might end in a correct solution
		return matchPartial(input, pattern, iIndex, patternIndex+2, isNextOptional)
	}
	return false
}

func isMatch(s string, p string) bool {
	optional := len(p) > 1 && p[1] == repeat
	return matchPartial([]rune(s), []rune(p), 0, 0, optional)
}

func expect(input, pattern string, correct bool) {
	if isMatch(input, pattern) != correct {
		negation := ""
		if !correct {
			negation = "not "
		}

		log.Fatalf("%v should %vmatch %v", input, negation, pattern)
	}
}

func main() {
	expect("asd", "a*", false)
	expect("asd", "as*d*sd", true)
	expect("casd", "c*asd", true)
	expect("asd", "a*d", false)

	expect("asd", "asd", true)
	expect("asd", "a.d", true)
	expect("asd", "as", false)

}
