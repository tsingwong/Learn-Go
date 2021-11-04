/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-11-02 14:32:44
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-11-02 21:25:35
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
