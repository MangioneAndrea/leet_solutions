package main

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
	str string
	num int
}

func (r *roman) remove(big, small int, bigL, smallL rune) *roman {
	for r.num%big != r.num {
		r.num -= big
		r.str += string(bigL)
	}

	if r.num >= big-small {
		r.num -= big - small
		r.str += string(smallL) + string(bigL)
	}

	return r
}

func intToRoman(num int) string {
	return (&roman{"", num}).
		remove(M, C, 'M', 'C').
		remove(D, C, 'D', 'C').
		remove(C, X, 'C', 'X').
		remove(L, X, 'L', 'X').
		remove(X, I, 'X', 'I').
		remove(V, I, 'V', 'I').
		remove(I, 0, 'I', '-').
		str
}

func main() {
	println(intToRoman(3))
	println(intToRoman(1994))
}
