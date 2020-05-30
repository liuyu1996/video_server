package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/api/session"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func RegisterHandlers() *httprouter.Router  {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:username", login)
	router.GET("/user/:username", GetUserInfo)
	router.POST("/user/:username/videos", AddNewVideo)
	router.GET("/user/:username/videos", ListAllVideo)
	router.DELETE("/user/:username/videos/:vid-id", DeleteVideo)
	router.POST("/user/:vid-id/comments", PostComment)
	router.GET("/user/:vid-id/comments", ShowComment)

	return router
}

func Prepare()  {
	session.LoadSessionsFromDB()
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	_ = http.ListenAndServe(":8000", mh)
}
