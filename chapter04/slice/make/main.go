package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 3)
	s2 := make([]int, 3, 5)

	fmt.Println(s1, s2)
	fmt.Printf("s1 len %d cap %d \n", len(s1), cap(s1))
	fmt.Printf("s2 len %d cap %d \n", len(s2), cap(s2))
}
