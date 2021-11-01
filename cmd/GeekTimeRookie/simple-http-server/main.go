/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-11-01 10:13:14
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-11-01 10:17:30
 */
package main

import "net/http"

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello world"))
	})
	http.ListenAndServe(":8080", nil)
}
