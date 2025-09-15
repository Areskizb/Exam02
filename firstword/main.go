package main

import (
    "fmt"
    student "piscine"
)

func main() {
    fmt.Print(student.FirstWord("hello there"))
    fmt.Print(student.FirstWord(""))
    fmt.Print(student.FirstWord("hello   .........  bye"))
}
