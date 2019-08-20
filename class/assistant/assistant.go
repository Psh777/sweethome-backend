package assistant

import (
	"../../webserver/handlers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseJson(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handlers.HandlerError(w, err.Error())
		return
	}
	log.Println(string(body))

	decoder := json.NewDecoder(r.Body)
	fmt.Printf("//// %+v //// %+v \n", r.Body, decoder)
	var t request
	err = decoder.Decode(&t)
	if err != nil {
		handlers.HandlerError(w, err.Error())
		return
	}
	fmt.Printf("%+v\n", t.Text)
	handlers.HandlerInterface(w, "ok")
}

type request struct {
	Text string `json:"text"`
}
