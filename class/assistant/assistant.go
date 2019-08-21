package assistant

import (
	"../../db/postgres"
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

	var t request

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println("error: " + err.Error())
		handlers.HandlerError(w, err.Error())
		return
	}

	err = json.Unmarshal(b, &t)
	if err != nil {
		fmt.Println()
		fmt.Println("error: " + err.Error())
		handlers.HandlerError(w, err.Error())
		return
	}

	fmt.Printf("%+v\n", t)

	switch t.QueryResult.Action {
	case "multisensor-1-temp":

		data, err := postgres.GetDataByTypeActually("ebaeecf4-58a1-48f0-8f10-c8c926aaa6c9", 1)
		if err != nil {
			return
		}

		str := fmt.Sprintf("%f", data)
		CreateResponse(w, str, str)

	case "":
		fmt.Println("hm")
	}

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
