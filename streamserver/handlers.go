package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"html/template"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")

	t.Execute(w, nil)
}

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid
	//打开
	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Error when try to open file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	//如果成功
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "file is too big")
		return
	}
	file, _, err := r.FormFile("file") //form  name ="file"
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error :%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")

	}
	filename := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+filename, data, 0666)
	if err != nil {
		log.Printf("Write file  error:%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Upload successfully")
}
