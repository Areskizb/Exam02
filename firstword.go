package student

import "github.com/01-edu/z01"

func FirstWord(s string) string {
	runes := []rune(s)
	if len(runes) == 0 {
		z01.PrintRune('\n')
	}
	for _, r := range s {
		z01.PrintRune(r)
		if r == ' ' {
			z01.PrintRune('\n')
			break
		}
	}
	return "\x00"
}
