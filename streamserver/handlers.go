package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid
	video, err := os.Open(vl)
	if err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	w.Header().Set("Content-type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		log.Println(err)
		sendErrorResponse(w, http.StatusBadRequest,"File is too big")
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	fileName := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR + fileName, data, 0666)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, _ = io.WriteString(w, "upload successfully")
}