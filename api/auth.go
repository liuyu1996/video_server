package main

import (
	"net/http"
	"video_server/api/session"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "x-User-Name"

func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}
	userName, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME, userName)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	userName := r.Header.Get(HEADER_FIELD_UNAME)
	if len(userName) == 0 {
		//sendErrorResponse()
		return false
	}
	return true

}
