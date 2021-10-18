/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-12 14:39:21
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-12 14:41:24
 */
package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	fmt.Println("pointer: ", &i)
}
