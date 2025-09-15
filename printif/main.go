package main

import (
	"fmt"
	student "piscine"
)

func main() {
	fmt.Print(student.PrintIf("abcdefz"))
	fmt.Print(student.PrintIf("abc"))
	fmt.Print(student.PrintIf(""))
	fmt.Print(student.PrintIf("14"))
}