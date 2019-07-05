package webserver

import (
	"../class/sensor"
	"./handlers"
	"net/http"
)

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	handlers.HandlerSuccess(w, "ok")
}

func sensorUploadHandler(w http.ResponseWriter, r *http.Request) {

	sensor.Upload(w,r)

}


