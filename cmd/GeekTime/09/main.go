/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-18 16:20:18
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-18 16:24:29
 */
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var err error
	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("%d bytes(s) were written.\n", n)
}
