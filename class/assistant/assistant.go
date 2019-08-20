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
		handlers.HandlerError(w, "1)"+err.Error())
		return
	}
	log.Println(string(body))

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	var t request

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(b, &t)
	//err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		handlers.HandlerError(w, err.Error())
		return
	}

	fmt.Printf("%+v\n", t)
	handlers.HandlerInterface(w, "ok")
}

type request struct {
	ResponseId  string      `json:"responseId"`
	QueryResult QueryResult `json:"queryResult"`
}

type QueryResult struct {
	QueryText  string     `json:"queryText"`
	Action     string     `json:"action"`
	Parameters Parameters `json:"parameters"`
}

type Parameters struct {
}
