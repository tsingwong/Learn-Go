/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-13 20:29:07
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-14 11:21:57
 */
package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
