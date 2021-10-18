/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-13 19:54:01
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-13 20:12:59
 */
package main

import "fmt"

func main() {
	messages := make(chan string)
	singals := make(chan string)
	go func() {
		messages <- "3"
	}()
	select {
	case msg := <-messages:
		fmt.Println("1 received message ", msg)
	default:
		fmt.Println("1 no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("2 sent message", msg)
	default:
		fmt.Println("2 no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("3 received message", msg)
	case sig := <-singals:
		fmt.Println("3 received signal", sig)
	default:
		fmt.Println("3 no activity")
	}

}
