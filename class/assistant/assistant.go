package assistant

import (
	"../../webserver/handlers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseJson(w http.ResponseWriter, r *http.Request) {
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	handlers.HandlerError(w, "1)"+err.Error())
	//	return
	//}
	//log.Println(string(body))

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

	t1 := make([]string, 0)
	t1 = append(t1, "text")
	t2 := make([]FulfillmentMessages, 0)
	var f1 FulfillmentMessages
	f1.Text = t1
	t2 = append(t2, f1)

	answer := answer{
		FulfillmentText:     "hello",
		FulfillmentMessages: t2,
	}

	handlers.HandlerInterfaceAssistant(w, answer)
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

type answer struct {
	FulfillmentText     string                `json:"fulfillment_text"`
	FulfillmentMessages []FulfillmentMessages `json:"fulfillmentMessages"`
}

type FulfillmentMessages struct {
	Text []string `json:"text"`
}
