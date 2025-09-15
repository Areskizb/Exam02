package main

import (
	student "piscine"
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(student.FirstRune("Hello!"))
	z01.PrintRune(student.FirstRune("Salut!"))
	z01.PrintRune(student.FirstRune("Ola!"))
	z01.PrintRune('\n')
}