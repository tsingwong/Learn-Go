/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-12 14:27:19
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-12 14:27:19
 */
package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
