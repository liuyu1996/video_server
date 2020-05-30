package utils

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetCurrentTimestampSec() int {
	ts, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1e9, 10))
	return ts
}

func SendDeleteVideoRequest(id string)  {
	addr := "127.0.0.1" + ":9001"
	url := "http://" + addr + "/video-delete-record/" + id
	_, err := http.Get(url)
	if err != nil {
		log.Printf("Send delete video request failed : %s", err)
		return
	}
}
