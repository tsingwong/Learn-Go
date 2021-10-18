/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-09 11:37:38
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-09 11:38:45
 */
package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
