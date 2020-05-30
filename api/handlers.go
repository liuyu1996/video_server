package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/session"
	"video_server/api/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	res, _ := ioutil.ReadAll(r.Body)
	user := &defs.Users{}
	if err := json.Unmarshal(res, user); err != nil{
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	if err := dbops.AddUserCredential(user.Username, user.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
	}
	id := session.GenerateNewSessionId(user.Username)
	su := &defs.SignedUp{Success:true, SessionId:id}
	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalError)
		return
	}else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func login (w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	res, _ := ioutil.ReadAll(r.Body)
	log.Printf("%s", res)
	ubody := &defs.Users{}
	if err := json.Unmarshal(res, ubody); err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	uname := p.ByName("username")
	if uname != ubody.Username {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	log.Printf("%s", ubody.Username)
	hashpwd, err := dbops.GetUserCredential(ubody.Username)
	ok, err := utils.PasswordVerify(hashpwd, ubody.Pwd)
	if err != nil || !ok {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	id := session.GenerateNewSessionId(ubody.Username)
	si := &defs.SignedUp{Success:true, SessionId:id}
	if res, err := json.Marshal(si); err != nil {
		sendErrorResponse(w, defs.ErrorInternalError)
	}else {
		sendNormalResponse(w, string(res), 200)
	}
}

func GetUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	if !ValidateUser(w, r) {
		log.Printf("Unauthorized user \n")
		return
	}
	uname := p.ByName("username")
	userInfo, err := dbops.GetUser(uname)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	user := &defs.Users{Id:userInfo.Id}
	if resp, err := json.Marshal(user); err != nil {
		sendErrorResponse(w, defs.ErrorInternalError)
		return
	}else {
		sendNormalResponse(w, string(resp), 200)
	}
}

func AddNewVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	if !ValidateUser(w, r) {
		log.Printf("Unauthorized user \n")
		return
	}
	res, _ := ioutil.ReadAll(r.Body)
	video := &defs.VideoInfo{}
	if err := json.Unmarshal(res, video); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	v, err := dbops.AddVideo(video.AuthorId, video.Name)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	if resp, err := json.Marshal(v); err !=nil {
		sendErrorResponse(w, defs.ErrorInternalError)
		return
	}else {
		sendNormalResponse(w, string(resp), 200)
	}
}

func ListAllVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	if !ValidateUser(w, r) {
		return
	}

}

func DeleteVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	if !ValidateUser(w, r) {
		return
	}
	vid := p.ByName("vid-id")
	err := dbops.DeleteVideo(vid)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	go utils.SendDeleteVideoRequest(vid)
	sendNormalResponse(w, "", 204)
}

func PostComment(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	if !ValidateUser(w, r) {
		return
	}
	req, _ := ioutil.ReadAll(r.Body)
	comment := &defs.Comments{}
	if err := json.Unmarshal(req, comment); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	vid := p.ByName("vid-id")
	if err := dbops.AddNewComments(vid, comment.AuthorId, comment.Content); err != nil{
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	sendNormalResponse(w, "ok", 201)
}

func ShowComment(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	if !ValidateUser(w, r) {
		return
	}
	vid := p.ByName("vid-id")
	commentList, err := dbops.ListComments(vid, 0, utils.GetCurrentTimestampSec())
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	if resp, err := json.Marshal(commentList); err != nil {
		sendErrorResponse(w, defs.ErrorInternalError)
	}else {
		sendNormalResponse(w, string(resp), 200)
	}
}