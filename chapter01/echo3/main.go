package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//使用strings.Join 避免重复做字符串的运算和存取
	fmt.Println(strings.Join(os.Args[1:], " "))
	//可以用下面的方式 调试程序 与上面的输出很像
	//fmt.Println(os.Args[1:])
}

/*
➜  echo3 git:(master) ✗ go run main.go hello test c++
hello test c++
[hello test c++]
➜  echo3 git:(master) ✗
*/
