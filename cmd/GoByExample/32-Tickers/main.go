/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-14 11:33:26
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-14 11:34:49
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
