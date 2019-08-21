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

	simpleResponse := SimpleResponse{
		TextToSpeech: "hello",
		DisplayText:  "hello",
	}

	item := Items{
		SimpleResponse: simpleResponse,
	}

	items := make([]Items, 0)
	items = append(items, item)

	richResponse := RichResponse{
		Items: items,
	}

	google := Google{
		ExpectUserResponse: true,
		RichResponse:       richResponse,
	}

	payload := Payload{
		Google: google,
	}

	answer := answer{
		Payload:             payload,
		//FulfillmentText:     "hello",
		//FulfillmentMessages: t2,
		Response:            "hello",
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
	Payload             Payload               `json:"payload"`
	FulfillmentText     string                `json:"fulfillmentText"`
	FulfillmentMessages []FulfillmentMessages `json:"fulfillmentMessages"`
	Response            string                `json:"response"`
}

type Payload struct {
	Google Google `json:"google"`
}

type Google struct {
	ExpectUserResponse bool         `json:"expectUserResponse"`
	RichResponse       RichResponse `json:"richResponse"`
}

type RichResponse struct {
	Items []Items `json:"items"`
}

type Items struct {
	SimpleResponse SimpleResponse `json:"simpleResponse"`
}

type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech"`
	DisplayText  string `json:"displayText"`
}
type FulfillmentMessages struct {
	Text []string `json:"text"`
}
