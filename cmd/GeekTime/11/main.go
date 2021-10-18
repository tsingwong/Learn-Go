/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-18 16:42:54
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-18 16:43:59
 */
package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}

	fmt.Printf("The element is %q.\n", container[1])
}
