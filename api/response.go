package main

import (
	"encoding/json"
	"io"
	"net/http"
	"video_server/api/defs"
)

func sendErrorResponse(w http.ResponseWriter, response defs.ErrorResponse)  {
	w.WriteHeader(response.HttpSC)
	resStr, _ := json.Marshal(&response.Error)
	_, _ = io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int)  {
	w.WriteHeader(sc)
	_, _ = io.WriteString(w, resp)
}
