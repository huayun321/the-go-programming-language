//dup1 输出标准输入种出现次数大于1的行，前面是次数
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	//如何终止input.Scan ctl+d
	//https://stackoverflow.com/questions/36214456/how-do-i-end-scanner-scan-loop-in-console-for-golang
	for input.Scan() {
		counts[input.Text()]++
	}
	//注意：忽略了input.Err()种可能出现的错误
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	//printf verb
	//%v 内置格式的任何值
	//%T 任何值的类型
}
