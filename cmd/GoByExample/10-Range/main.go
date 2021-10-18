/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-11 17:35:40
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-11 17:54:35
 */
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key: ", k)
	}
	str := "go, 汉字会怎样？"
	for i, c := range str {
		fmt.Println(i, c)
	}
	for i := 0; i < len(str); i++ {
		fmt.Println(i, str[i])
	}
}
