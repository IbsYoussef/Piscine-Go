package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append([]rune{'0' + rune(digit)}, digits...)
		n /= 10
	}

	for _, r := range digits {
		z01.PrintRune(r)
	}
}

func main() {
	points := &point{}

	setPoint(points)

	printStr("x = ")
	printNumber(points.x)
	printStr(", y = ")
	printNumber(points.y)
	z01.PrintRune('\n')
}
