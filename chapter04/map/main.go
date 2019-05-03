package main

import "fmt"

func main() {
	// map 会被初始化零值
	// 访问不存在的建 会返回零值
	m := make(map[string]map[string]bool)
	if m["hello"] == nil {
		fmt.Println("nil")
	}
	fmt.Println(m)
	fmt.Printf("%v %T", m["hello"],  m["hello"])
}
