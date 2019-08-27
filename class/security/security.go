package security

import (
	"../../modules/telegram"
	"../../webserver/handlers"
	"net/http"
)

func On(w http.ResponseWriter, _ *http.Request) {
	telegram.SendMsgBot("Security ON")
	handlers.HandlerInterface(w, "ok")
}

func Off(w http.ResponseWriter, _ *http.Request) {
	telegram.SendMsgBot("Security OFF")
	handlers.HandlerInterface(w, "ok")

}

func Alarm(w http.ResponseWriter, _ *http.Request) {
	telegram.SendMsgBot("ALARM!!! ALARM!!! ALARM!!!")
	handlers.HandlerInterface(w, "ok")
}
