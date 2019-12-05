package main

import (
	"fmt"
	"unicode/utf8"
)
// 寻找最长不含重复字符的字串
func lengthOfNonRepeatSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {

		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	s := "yes中文中文"
	fmt.Println(s)
	for _, b := range []byte(s) {
		fmt.Printf("%X  ", b)
	}

	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}

	fmt.Println()
	fmt.Println("rune count:", utf8.RuneCountInString(s))

	//单个字符转义
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}

	fmt.Println()

	//转成rune数组
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}

	fmt.Println()

	fmt.Println(lengthOfNonRepeatSubStr("kkkkk"))
}
