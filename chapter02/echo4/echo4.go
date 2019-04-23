// echo4 输出其命令行参数
package main

import (
	"flag"
	"fmt"
	"strings"
)

//返回的都是指针
var n = flag.Bool("n", false, "omit trailing newline") //是否省略最后的换行符
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	//所以这里sep 要使用* 取指针变量sep 的值
	fmt.Print(strings.Join(flag.Args(), *sep))

	//如果没有将n 设为true 则打印换行符
	if !*n {
		fmt.Println()
	}
}

/**
➜  echo4 git:(master) ✗ go run echo4.go -s / -n a b c
a/b/c%
➜  echo4 git:(master) ✗ go run echo4.go -s / a b c
a/b/c
➜  echo4 git:(master) ✗

 */
