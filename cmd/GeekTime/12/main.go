/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-18 16:55:11
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-18 19:11:38
 */
package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}

	_, ok1 := interface{}(container).([]string)
	_, ok2 := interface{}(container).(map[int]string)
	if ok1 || ok2 {
		fmt.Printf("The element is %q.\n", container[1])
	}

}
