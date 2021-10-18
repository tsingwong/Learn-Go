/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-13 11:47:44
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-13 11:55:55
 */
package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
