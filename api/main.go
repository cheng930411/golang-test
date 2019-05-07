package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"video_server/api/session"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}
func RegisterHandlers() *httprouter.Router {
	log.Printf("preparing to post request\n")
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	router.GET("/user/:username", GetUserInfo)
	router.POST("/user/:username/videos", AddNewVideo)
	router.GET("/user/:username/videos", ListAllVideos)
	router.DELETE("/user/:username/videos/:vid-id", DeleteVideo)
	router.POST("/videos/:vid-id/comments", PostComment)
	router.GET("/videos/:vid-id/coßßmments", ShowComments)
	return router
}
func Prepare() {
	session.LoadSessionFromDB()
}
func main() {
	Prepare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}

//listen->RegisterHandelrs->handlers
//handler->validation{1.request,2.user}->business logic->reponse
//1.data model 2.error handling
//main->middleware->defs(message,err)->handlers->dbops->response
//http.Handler  a
//ab struct a
//duck type
