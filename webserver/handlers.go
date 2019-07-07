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

func getSensorHandler(w http.ResponseWriter, r *http.Request) {
	sensor.GetSensors(w,r)
}

func sensorGetDataHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	t := r.URL.Query().Get(":type")
	sensor.GetDataSensor(w,r, id, t)
}




