/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-11 20:01:00
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-11 20:03:06
 */
package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 = ", res)
}
