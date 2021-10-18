/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-11 20:11:14
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-11 20:17:49
 */
package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
