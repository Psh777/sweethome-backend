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

	resp := Response{
		Text: "Привет",
		TTS: "Привет",
	}

	handlers.HandlerInterfaceAssistant(w, answer{
		Version: "1.0",
		Sesion:  t.Session,
		Response: resp,

		//RequestID: t.Request.Session.
	})

}

type request struct {
	Request Request `json:"request"`
	Session Session `json:"session"`
}

type Request struct {
	Command string `json:"command"`
}

type Session struct {
	MessageID int    `json:"message_id"`
	New       bool   `json:"new"`
	SessionID string `json:"session_id"`
	SkillID   string `json:"skill_id"`
	UserID    string `json:"user_id"`
}

type answer struct {
	Sesion   Session  `json:"sesion"`
	Response Response `json:"response"`
	Version  string   `json:"version"`
}

type Response struct {
	Text string `json:"text"`
	TTS  string `json:"tts"`
}
