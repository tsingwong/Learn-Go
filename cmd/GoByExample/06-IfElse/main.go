/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-09 13:46:10
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-09 16:11:46
 */
package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
