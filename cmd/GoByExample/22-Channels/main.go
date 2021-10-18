/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-13 11:38:47
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-13 11:41:57
 */
package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
