/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-13 17:51:55
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-13 19:17:02
 */
package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
