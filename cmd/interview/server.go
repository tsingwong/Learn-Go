/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-11-22 22:31:40
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-11-22 23:05:37
 */
package main

import (
	"fmt"
	"net/http"
)

type Server interface {
	Route(pattern string, handleFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

func (s *sdkHttpServer) Route(pattern string, handleFunc http.HandlerFunc) {
	http.HandleFunc(pattern, handleFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func NewHTTPServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}

// 在没有 context 抽象的情况下，是长这样的
func SignUpWithoutContext(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}
	ctx := &Context{
		W: w,
		R: r,
	}

	err := ctx.ReadJson(req)
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	resp := &commonResponse{
		Data: 123,
	}
	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		// 写日志
		fmt.Printf("相应失败: %v", err)
	}
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
