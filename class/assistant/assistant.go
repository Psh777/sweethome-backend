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
		handlers.HandlerError(w, "1)" + err.Error())
		return
	}
	log.Println(string(body))

	decoder := json.NewDecoder(r.Body)
	fmt.Printf("//// %+v //// %+v \n", r.Body, decoder)
	var t request
	err1 := decoder.Decode(&t)
	if err1 != nil {
		//handlers.HandlerError(w, "2)" + err1.Error())
		//return
	}
	fmt.Printf("%+v\n", t.Text)
	handlers.HandlerInterface(w, "ok")
}

type request struct {
	Text string `json:"text"`
}
