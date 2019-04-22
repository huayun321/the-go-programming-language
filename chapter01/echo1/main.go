package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	//i++ i-- 是语句不是表达式 j=i++是不合法的
	//l := 1
	//k := l++

	//for的四种形式
	//for initialization; condition; post {}
	//先初始化 后判断条件
	//如果为真 则执行for块内代码
	//如果为假 则终止

	//for condition {}
	//与while condition {} 相同

	//for {}
	//无限循环
	//使用break return跳出

	//for i, v := range list
	//for k, v := range map
	//迭代list 或 map
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

	//数组的切片
	fmt.Println(len(os.Args))
	fmt.Println(os.Args[:len(os.Args)])
	//注意数组切片：
	//所有引用使用半开区间，即包含第一个索引，不包含最后一个索引
	//slice[m:n] 0<=m<=n<=len(s),包含n-m个元素
	arr := []int{1, 2, 3}
	fmt.Println(arr[0:3])
}

/*
➜  echo1 git:(master) ✗ go build main.go
➜  echo1 git:(master) ✗ ./main hello golang rpc test
hello golang rpc test
5
[./main hello golang rpc test]
[1 2 3]
➜  echo1 git:(master) ✗
*/
