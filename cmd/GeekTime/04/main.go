/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-18 15:06:24
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-18 15:08:43
 */
package main

import "flag"

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	hello(name)
}
