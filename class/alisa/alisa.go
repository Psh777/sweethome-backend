package alisa

import (
	"../../webserver/handlers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseJson(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		fmt.Println("error: no body")
		handlers.HandlerError(w, "No body")
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println("error: " + err.Error())
		handlers.HandlerError(w, err.Error())
		return
	}

	var t request

	err = json.Unmarshal(b, &t)
	if err != nil {
		fmt.Println()
		fmt.Println("error: " + err.Error())
		handlers.HandlerError(w, err.Error())
		return
	}
	fmt.Printf("%+v\n", t)

	handlers.HandlerInterface(w, answer{
		//RequestID: t.Request.Session.
	})

}

type request struct {
	Request Request `json:"request"`
}

type Request struct {
	Command string  `json:"command"`
	Session Session `json:"session"`
}

type Session struct {
	MessageID int    `json:"message_id"`
	New       bool   `json:"new"`
	SessionID string `json:"session_id"`
	SkillID   string `json:"skill_id"`
	UserID    string `json:"user_id"`
}

type answer struct {
	RequestID string `json:"request_id"`
}
