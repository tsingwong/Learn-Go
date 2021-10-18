/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-14 19:43:17
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-14 19:45:44
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

// 必须指针传递
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}
