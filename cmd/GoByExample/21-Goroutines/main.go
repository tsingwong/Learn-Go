/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-12 17:29:42
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-13 11:27:12
 */
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// 同步方法
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
