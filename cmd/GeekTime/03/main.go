/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-18 11:36:31
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-18 15:04:20
 */
package main

import (
	"flag"
	"fmt"
	"os"
)

var name *string
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	// flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	// flag.CommandLine.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }
	name = cmdLine.String("name", "everyone", "The greeting object")
	// name = flag.String("name", "everyone", "The greeting object")
	// flag.StringVar(&name, "name", "everyone", "The greeting object")
}

func main() {
	// flag.Parse()
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s!\n", *name)
}
