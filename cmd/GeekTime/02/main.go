/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-18 11:36:31
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-18 14:37:56
 */
package main

import (
	"flag"
	"fmt"
)

var name *string

func init() {
	name = flag.String("name", "everyone", "The greeting object")
	// flag.StringVar(&name, "name", "everyone", "The greeting object")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", *name)
}
