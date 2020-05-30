package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/scheduler/taskrunner"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video-delete-record/:vid-id", videoDelRecHandler)

	return router
}

func main()  {
	go taskrunner.Start()
	r := RegisterHandlers()
	_ = http.ListenAndServe(":9001", r)
}



