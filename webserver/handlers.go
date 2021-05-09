package webserver

import (
	"../class/alisa"
	"../class/assistant"
	"../class/sensor"
	"../modules/telegram"
	"./handlers"
	"net/http"
)

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	handlers.HandlerSuccess(w, "ok")
}

func sensorUploadHandler(w http.ResponseWriter, r *http.Request) {
	sensor.Upload(w, r)
}

func getSensorHandler(w http.ResponseWriter, r *http.Request) {
	sensor.GetSensors(w, r)
}

func sensorGetDataHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	t := r.URL.Query().Get(":type")
	sensor.GetDataSensor(w, r, id, t)
}

func sensorPostDataHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	t := r.URL.Query().Get(":type")
	sensor.PostDataSensor(w, r, id, t)
}

func assistantPostHandler(w http.ResponseWriter, r *http.Request) {
	assistant.ParseJson(w, r)
}

func alisaPostHandler(w http.ResponseWriter, r *http.Request) {
	alisa.ParseJson(w, r)
}

func alisaGetDevicesHandler(w http.ResponseWriter, r *http.Request) {
	alisa.Devices(w, r)
}

func alisaDevicesActionHandler(w http.ResponseWriter, r *http.Request) {
	alisa.Action(w, r)
}

func alisaDevicesStateHandler(w http.ResponseWriter, r *http.Request) {
	alisa.DeviceState(w, r)
}

func securityOnHandler(w http.ResponseWriter, r *http.Request) {
	telegram.SendMsgBot("Security ON")
	handlers.HandlerInterface(w, "ok")
}

func securityOffHandler(w http.ResponseWriter, r *http.Request) {
	telegram.SendMsgBot("Security OFF")
	handlers.HandlerInterface(w, "ok")
}

func securityAlarmHandler(w http.ResponseWriter, r *http.Request) {
	zone := r.URL.Query().Get(":zone")
	sensortype := r.URL.Query().Get(":sensortype")

	telegram.SendMsgBot("ALARM! (" + sensortype + ") Zone: " + zone)
	handlers.HandlerInterface(w, "ok")
}
