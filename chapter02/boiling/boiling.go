// boiling 输出水的沸点 声明的例子
package main

import "fmt"

//包级别常量 作用域包内可见
const boilingF = 212.0

func main() {
	//函数内变量 函数内可见
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// 输出：
	// boiling point = 212°F or 100°C
}
