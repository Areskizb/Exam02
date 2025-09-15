package student

import "github.com/01-edu/z01"

func RepeatAlpha(s string) string {
	var a int
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			a = int(r) - int('a')
				for i := 0; i <= a; i++ {
					z01.PrintRune(r)
				}
		} else if r >= 'A' && r <= 'Z' {
			a = int(r) - int('A')
				for i := 0; i <= a; i++ {
					z01.PrintRune(r)
				}
		} else {
			z01.PrintRune(r)
		}
	}
	return "\x00"
}
