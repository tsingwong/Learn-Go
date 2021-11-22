/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-11-22 22:51:39
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-11-23 07:41:21
 */
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IContext interface {
	ReadJson(obj interface{}) error
	WriteJson(code int, resp interface{}) error
	OKJson(resp interface{}) error
	SystemErrorJson(resp interface{}) error
}

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) ReadJson(req interface{}) error {
	r := c.R
	w := c.W
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		// 要返回掉，不然就会继续执行后面的代码
		return err
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		fmt.Fprintf(w, "deserialized failed: %v", err)
		// 要返回掉，不然就会继续执行后面的代码
		return err
	}
	return nil
}
func (c *Context) WriteJson(code int, resp interface{}) error {
	w := c.W
	w.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		fmt.Fprintf(w, "Marshal failed: %v", err)
		return err
	}
	_, err = w.Write(respJson)
	return err
}
func (c *Context) OKJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}

func (c *Context) SystemErrorJson(resp interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, resp)
}
