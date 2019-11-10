package alisa

import (
	"../../modules/config"
	"../../modules/http_request"
	"../../webserver/handlers"
	"../assistant"
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
	//fmt.Printf("%+v\n", t)

	//diaglofFlow
	text := Text{
		Text:         t.Request.Command,
		LanguageCode: "RU-ru",
	}

	q := QueryInout{
		Text: text,
	}

	body := RequestDialogFlow{
		QueryInput: q,
	}

	c := config.GetMyConfig()

	jsonBody, _ := json.Marshal(body)
	ansDialogFlowJson, err := http_request.POSTFLOW("https://dialogflow.googleapis.com/v2beta1", "projects/"+ c.Env.DialogFlowProjectID + "/agent/sessions/123456789:detectIntent", string(jsonBody))
	if err != nil {
		fmt.Println(err)
		handlers.HandlerError(w, err.Error())
		return
	}

	var ansDialogFlow DialogFlowResponse
	err = json.Unmarshal(ansDialogFlowJson, &ansDialogFlow)
	if err != nil {
		fmt.Println(err)
		handlers.HandlerError(w, err.Error())
		return
	}

	var resp Response

	if ansDialogFlow.WebhookStatus.Code > 0 {

		resp = Response{
			Text: ansDialogFlow.QueryResult.FulfillmentText,
		}

	} else {

		resp = Response{
			Text: ansDialogFlow.QueryResult.WebhookPayload.Google.RichResponse.Items[0].SimpleResponse.DisplayText,
			TTS:  ansDialogFlow.QueryResult.WebhookPayload.Google.RichResponse.Items[0].SimpleResponse.TextToSpeech,
		}

	}

	handlers.HandlerInterfaceAssistant(w, answer{
		Version:  "1.0",
		Session:  t.Session,
		Response: resp,
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
	Session  Session  `json:"session"`
	Response Response `json:"response"`
	Version  string   `json:"version"`
}

type Response struct {
	Text string `json:"text"`
	TTS  string `json:"tts"`
}

type RequestDialogFlow struct {
	QueryInput QueryInout `json:"query_input"`
}

type QueryInout struct {
	Text Text `json:"text"`
}

type Text struct {
	Text         string `json:"text"`
	LanguageCode string `json:"language_code"`
}

type DialogFlowResponse struct {
	QueryResult   QueryResult   `json:"queryResult"`
	WebhookStatus WebhookStatus `json:"webhookStatus"`
}

type QueryResult struct {
	WebhookPayload  WebhookPayload `json:"webhookPayload"`
	FulfillmentText string         `json:"fulfillmentText"`
}

type WebhookPayload struct {
	Google assistant.Google `json:"google"`
}

type WebhookStatus struct {
	Code int `json:"code"`
}
