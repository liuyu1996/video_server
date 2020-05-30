package main

import (
	"io"
	"net/http"
)

func sendResponse(w http.ResponseWriter, sc int, rep string)  {
	w.WriteHeader(sc)
	_, _ = io.WriteString(w, rep)
}