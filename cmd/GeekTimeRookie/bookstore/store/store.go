/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-11-01 11:13:46
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-11-03 11:32:03
 */
package store

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

type Book struct {
	Id      string   `json:"id"`      // 图书ID
	Name    string   `json:"name"`    // 图书名
	Authors []string `json:"authors"` // 作者
	Press   string   `json:"press"`   // 出跋社.
}

type Store interface {
	Create(*Book) error
	Update(*Book) error
	Get(string) (Book, error)
	GetAll() ([]Book, error)
	Delete(string) error
}
