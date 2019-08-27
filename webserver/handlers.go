package webserver

import (
	"../class/security"
	"../class/sensor"

	"../class/assistant"
	"./handlers"
	"net/http"
)

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	handlers.HandlerSuccess(w, "ok")
}

func sensorUploadHandler(w http.ResponseWriter, r *http.Request) {
	sensor.Upload(w,r)
}

func getSensorHandler(w http.ResponseWriter, r *http.Request) {
	sensor.GetSensors(w,r)
}

func sensorGetDataHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	t := r.URL.Query().Get(":type")
	sensor.GetDataSensor(w,r, id, t)
}

func sensorPostDataHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	t := r.URL.Query().Get(":type")
	sensor.PostDataSensor(w,r, id, t)
}

func assistantPostHandler(w http.ResponseWriter, r *http.Request) {
	assistant.ParseJson(w,r)
}

func securityOnHandler(w http.ResponseWriter, r *http.Request) {
	security.On(w,r)
}

func securityOffHandler(w http.ResponseWriter, r *http.Request) {
	security.Off(w,r)
}

func securityAlarmHandler(w http.ResponseWriter, r *http.Request) {
	security.Alarm(w,r)
}

