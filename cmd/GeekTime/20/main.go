/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-21 17:52:28
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-21 19:30:49
 */
package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n", elem1)
	fmt.Printf("The second element received from channel ch1: %v\n", <-ch1)
	fmt.Printf("The last element received from channel ch1: %v\n", <-ch1)
}
