/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-18 16:31:06
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-18 16:31:06
 */
package main

import "fmt"

var block = "package "

func main() {
	block := "function"
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)

}
