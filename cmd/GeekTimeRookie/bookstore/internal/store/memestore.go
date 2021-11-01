/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-11-01 11:12:58
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-11-01 19:34:33
 */
package store

import (
	mystore "bookstore/store"
	"bookstore/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}
