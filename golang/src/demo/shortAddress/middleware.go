package main

import (
	"log"
	"net/http"
	"time"
)

// ------------- 中间件

type Middleware struct {
}

// 记录请求所消耗的时间
func (m Middleware) LoggingHandle(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v", r.Method, r.URL.String(), t2.Sub(t1))
	}
	return http.HandlerFunc(f)
}

// 将程序从panic中恢复
func (m Middleware) RecoverHandler(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover form panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}
