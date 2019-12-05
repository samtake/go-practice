package main

import (
	"fmt"
)

func exp() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}      //冒号等于的定义，需要给数组一个初值
	arr3 := [...]int{1, 2, 3, 4} //或者不指定数组个数
	fmt.Println(arr1, arr2, arr3)

	//二维数组定义
	var grid [4][5]int
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	for _, v := range arr3 {
		fmt.Println(v)
	}

	printArray(&arr3)
	fmt.Println(arr3, arr3[0])
}

func printArray(arr *[4]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100 //在外面调用时，改变为100
}

func main() {
	exp()
}
