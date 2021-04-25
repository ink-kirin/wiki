package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/justinas/alice"
	"gopkg.in/validator.v2"
)

// App encapsulates Env, Router and middleware
type App struct {
	Router     *mux.Router
	Middleware *Middleware
	Config     *Env
}

// 请求结构体
type shortenReq struct {
	URL                 string `json:"url" validate:"nonzero"`                 // 地址
	ExpirationInMinutes int64  `json:"expiration_in_minutes" validate:"min=0"` // 过期时间
}

// 响应结构体
type shortlinkResp struct {
	Shortlink string `json:"shortlink"` // 短地址
}

// 初始化App结构体
func (a *App) Initialize(e *Env) {
	// set log formatter 定义log的格式 (时间|行号)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	a.Config = e
	a.Router = mux.NewRouter()
	a.Middleware = &Middleware{}
	a.initializeRoutes()
}

// 绑定路由与Handle的关系
func (a *App) initializeRoutes() {
	// a.Router.HandleFunc("/api/shorten", a.createShortlink).Methods("POST")
	// a.Router.HandleFunc("/api/info", a.getShortlinkInfo).Methods("GET")
	// a.Router.HandleFunc("/{shortlink:[a-zA-Z0-9]{1,11}}", a.redirect).Methods("GET")
	m := alice.New(a.Middleware.LoggingHandle, a.Middleware.RecoverHandler)
	a.Router.Handle("/api/shorten", m.ThenFunc(a.createShortlink)).Methods("POST")
	a.Router.Handle("/api/info", m.ThenFunc(a.getShortlinkInfo)).Methods("GET")
	a.Router.Handle("/{shortlink:[a-zA-Z0-9]{1,11}}", m.ThenFunc(a.redirect)).Methods("GET")
}

// 把长地址转换为短地址
func (a *App) createShortlink(w http.ResponseWriter, r *http.Request) {
	var req shortenReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, StatusError{http.StatusBadRequest, fmt.Errorf("parse paraments failed %v", r.Body)})
		return
	}
	if err := validator.Validate(req); err != nil {
		respondWithError(w, StatusError{http.StatusBadRequest, fmt.Errorf("validate paraments failed %v", req)})
		return
	}
	defer r.Body.Close()

	fmt.Printf("%s\n", req)
}

func (a *App) getShortlinkInfo(w http.ResponseWriter, r *http.Request) {
	// 获取请求参数
	vals := r.URL.Query()
	s := vals.Get("shortlink")

	fmt.Printf("%s\n", s)
	panic(s)
}

func (a *App) redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("%s\n", vars["shortlink"])
}

func (a *App) Run(add string) {
	// 监听
	log.Fatal(http.ListenAndServe(add, a.Router))
}

// 定义错误格式
func respondWithError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case Error:
		log.Printf("HTTP %d - %s", e.Status(), e.Error())
		respondWithJSON(w, e.Status(), e.Error())
	default:
		respondWithJSON(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

// 定义JSON格式
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	resp, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "applocation/json")
	w.WriteHeader(code)
	w.Write(resp)
}
