/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-25 20:51:26
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-29 11:43:55
 */
package main

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	logger.Info("hello, go module", zap.ByteString("uri", ctx.RequestURI()))
}

func main() {
	fasthttp.ListenAndServe(":8081", fastHTTPHandler)
}
