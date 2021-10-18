/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-18 15:55:55
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-18 15:57:00
 */
package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
	flag.Parse()
	fmt.Printf("Hello, %v!\n", name)
}
