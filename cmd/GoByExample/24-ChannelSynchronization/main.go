/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-13 14:07:36
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-13 17:51:17
 */
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}
