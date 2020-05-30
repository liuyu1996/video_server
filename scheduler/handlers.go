package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/scheduler/dbops"
)

func videoDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	vid := p.ByName("vid-id")
	if len(vid) == 0 {
		sendResponse(w, 400, "video id is empty")
		return
	}
	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		sendResponse(w, 500, "internal error")
		return
	}
	sendResponse(w, 200, "")
	return
}
