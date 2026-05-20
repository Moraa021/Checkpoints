package main

import "github.com/01-edu/z01"

func main() {
	firstiteration := true

	for a := '9'; a >= '2'; a-- {
		for b := a - 1; b >= '1'; b-- {
			for c := b - 1; c >= '0'; c-- {
				if !firstiteration {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)

				firstiteration = false
			}
		}
	}
	z01.PrintRune('\n')
}
