package main

import (
	"fmt"
	"regexp"
)

const text = `My email is sam@gmail.com
5555@qq.com
5555@163.com
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)

	for _, m := range match {
		fmt.Println(m)
	}
}
