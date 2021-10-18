/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-11 20:24:13
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-12 14:25:16
 */
package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
