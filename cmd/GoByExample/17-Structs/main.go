/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-12 15:57:30
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-12 16:12:34
 */
package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.name)

	sp.age = 51
	fmt.Println(s.age)
}
