/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-18 15:15:36
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-18 15:32:47
 */
package main

import (
	"flag"
	"learn-go/cmd/GeekTime/05/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)
}
