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

	fmt.Printf("%s %20s!\n", s, s)

	fmt.Println("g", isPalindrome("g"))
	fmt.Println("go", isPalindrome("go"))
	fmt.Println("gog", isPalindrome("gog"))
	fmt.Println("gogo", isPalindrome("gogo"))
	fmt.Println("g🤔g", isPalindrome("g🤔g"))

}

// isPalindrome("g") -> true
// isPalindrome("go") -> false
// isPalindrome("gog") -> true
// isPalindrome("gogo") -> false
// isPalindrome("g🤔g") -> true
// 比较字符串的首尾字符、第二个字符和倒数第二个字符，以此类推，来确定字符串是否对称
func isPalindrome(s string) bool {
	// todo: you code goer here
	rs := []rune(s)
	//fmt.Println(rs)
	for i := 0; i < len(rs)/2; i++ {
		if rs[i] != rs[len(rs)-i-1] {
			return false
		}
	}
	return true
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
