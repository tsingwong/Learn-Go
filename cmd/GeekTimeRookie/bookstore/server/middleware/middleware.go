/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-11-01 11:13:12
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-11-03 11:18:15
 */
package middleware

import (
	"log"
	"mime"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("rec a %s request from %s", r.Method, r.RemoteAddr)
		next.ServeHTTP(rw, r)
	})
}

func Validating(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		mediatype, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
		}
		if mediatype != "application/json" {
			http.Error(rw, "invalid Content-Type", http.StatusUnsupportedMediaType)
		}
		next.ServeHTTP(rw, r)
	})
}
