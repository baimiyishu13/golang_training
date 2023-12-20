package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G🤔", 6)
	// 传入到 banner的实际上是两个指针

	s := "G🤔"
	fmt.Println("len:", len(s)) // = 6 unicode

	for i, r := range s {
		fmt.Println(i, r)
		if i == 1 {
			fmt.Printf("%c of type %T\n", r, r)
		}
	}

	// byte uint8
	// rune int32
	b := s[0]
	fmt.Printf("%c of type %T\n", b, b) //G of type uint8%

	x, y := 1, "1"
	fmt.Printf("x=%v,y=%v\n", x, y)
	fmt.Printf("x=%#v,y=%#v\n", x, y) // Use #v in debu/log

	fmt.Printf("%s %20s!", s, s)
}

// isPalindrome("g") -> true
// isPalindrome("go") -> false
// isPalindrome("gog") -> true
// isPalindrome("gogo") -> false
func isPalindrome(s string) bool {
	// todo: you code goer here

	return false
}

func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text)) / 2 // BUG: len is in bytes
	//padding := (width - len(text)) / 2   // BUG: len is in bytes
	// padding 需要大于 0
	println(padding)
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
