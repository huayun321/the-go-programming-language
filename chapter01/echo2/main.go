package main

import (
	"fmt"
	"os"
)

func main() {
	//相对于echo1程序这个程序使用了另外一种for的形式
	//直接遍历数组

	//函数内声明变量的简短方式
	s, sep := "", ""

	//以下变量的声明方式是等价的
	//s := "" //适合函数内声明 简洁
	//var s string //函数内 或包内 都可以
	//var s = "" //go 的字符串 有空默认值 不需要这样使用
	//var s string = "" //go会自我推导变量 不需要这样使用

	//不要系统参数的第一个 也就是执行程序的命令
	//这样虽然比echo1简洁了许多
	//但是字符串每次重新赋值也是个低效的用法
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
