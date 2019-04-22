// ftoc 输出两个华氏温度-摄氏温度的转化 声明的例子
package main

import "fmt"

func main() {
	//常量声明
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

//函数声明
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
