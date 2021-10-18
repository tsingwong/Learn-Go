/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-08 19:31:15
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-09 11:02:14
 */
package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
