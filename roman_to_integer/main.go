package main

import "strings"

const (
	I = 1
	V = 5
	X = 10
	L = 50
	C = 100
	D = 500
	M = 1000
)

type roman struct {
	s   string
	val int
}

func (rom roman) subValue(n string, valA, valB int) roman {
	charCount := 0
	reducerCount := 0

	lastPos := strings.LastIndex(rom.s, n)

	if lastPos == -1 {
		return rom
	}

	for i := 0; i <= lastPos; i++ {
		if string(rom.s[i]) == n {
			charCount++
		} else {
			reducerCount++
		}
	}
	rom.val += charCount*valA - reducerCount*valB

	if lastPos+1 >= len(rom.s) {
		rom.s = ""
	} else {
		rom.s = rom.s[lastPos+1:]
	}
	return rom
}

func romanToInt(s string) int {
	return roman{s, 0}.
		subValue("M", M, C).
		subValue("D", D, C).
		subValue("C", C, X).
		subValue("L", L, X).
		subValue("X", X, I).
		subValue("V", V, I).
		subValue("I", I, 0).val
}

func main() {
	println(romanToInt("MCMXCIV"))
}
