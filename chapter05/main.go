package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		//指数退避策略的简单实现
		//睡眠时间为 time.Second * 2^0  time.Second * 2^1  time.Second * 2^2 ...
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond affter %s", url, timeout)
}

func main()  {
	fmt.Println(time.Second << 2)
}