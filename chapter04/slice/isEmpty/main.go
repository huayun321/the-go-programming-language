package main

import "fmt"

func main() {
	var s []int // len(s) == 0, s == nil
	fmt.Printf("len==%d, s==%t \n", len(s), s == nil)
	s = nil // len(s) == 0, s == nil
	fmt.Printf("len==%d, s==%t \n", len(s), s == nil)
	s = []int(nil) // len(s) == 0, s == nil
	fmt.Printf("len==%d, s==%t \n", len(s), s == nil)
	s = []int{} // len(s) == 0, s != nil
	fmt.Printf("len==%d, s==%t \n", len(s), s == nil)
	s = make([]int, 3)
	s = s[3:]
	fmt.Printf("len==%d, s==%t \n", len(s), s == nil)

	//slice的零值是nil,值为nil的slice长度和容量都是0
	//但是也有非nil的值和容量是0的情况
	//所以判断一个slice是否为空  应该使用
	if len(s) == 0 {
		fmt.Println("s is empty")
	}

}
