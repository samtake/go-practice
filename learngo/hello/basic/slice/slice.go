package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func printArray(arr []int) {
	arr[1] = 50
	for i, v := range arr {
		fmt.Println(i, v)
	}

}

func sliceExtension() {
	fmt.Println("sliceExtension")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
}

func sliceAppend() {
	fmt.Println("slice append")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]       //2 3 4 5
	s2 := s1[3:5]        //5 6
	s3 := append(s2, 10) //5 6 10
	s4 := append(s3, 11) //5 6 10 11
	s5 := append(s4, 12) //5 6 10 11 12
	fmt.Println("s3=", s3)
	fmt.Println("s4=", s4)
	fmt.Println("s5=", s5)
	fmt.Println("arr=", arr) //0, 1, 2, 3, 4, 5, 6 10  数组长度不变的
}

func printSlice(s []int) {
	// for i, v := range s {
	// 	fmt.Println(i, v)
	// }
	fmt.Printf("%v, len=%d, cap= %d\n", s, len(s), cap(s))
}
func sliceCreate() {
	var s []int
	for i := 0; i < 20; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	println(s)

	s1 := []int{2, 3, 5, 6, 8, 2, 5}
	println(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32) //32是开辟的空间cap
	println("print Slice  s2 s3")
	printSlice(s2)
	printSlice(s3)

	println("Slice copy")
	copy(s2, s1) //将s1复制给s2
	printSlice(s2)

	//[2 3 5 6 8 2 5 0 0 0 0 0 0 0 0 0], len=16, cap= 16
	//如何删除8
	println("Slice delete")
	s2 = append(s2[:4], s2[5:]...)
	printSlice(s2)

	//删除头尾
	fmt.Println("pop from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println("pop from back")
	back := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(front, back)
	printSlice(s2)
}

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//s := arr[2:6]
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	fmt.Println("arr[2:] =", arr[2:])
	fmt.Println("arr[:] =", arr[:])

	s1 := arr[2:]
	s2 := arr[:]

	fmt.Println("after updateSlice(s1")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("after updateSlice(s2")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("printArray")
	printArray(arr[:])

	fmt.Println("reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	sliceExtension()

	sliceAppend()

	sliceCreate()
}
