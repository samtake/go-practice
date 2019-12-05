package main

import (
	"fmt"
)

func lengthOfNonRepeatSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {

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
	m := map[string]string{
		"name":  "sam",
		"class": "go",
	}

	fmt.Println(m)

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m2, m3)

	//打印
	for k, v := range m {
		fmt.Println(k, v)
	}

	//获取值
	fmt.Println("getting values")
	name := m["name"]
	println(name)

	fmt.Println("getting values")
	if name, ok := m["name"]; ok {
		println(name)
	} else {
		println("key does not exist")
	}

	//删除
	fmt.Println("delete values")
	name1, ok := m["name"]
	println(name1, ok)

	length := lengthOfNonRepeatSubStr("jjjjjlll")
	println(length)

}
