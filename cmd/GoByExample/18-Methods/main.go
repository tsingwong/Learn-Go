/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-12 16:15:55
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-12 16:25:29
 */
package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}
func main() {
	r := rect{10, 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim: ", rp.perim())
}
