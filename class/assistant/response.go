package assistant

import (
	"../../webserver/handlers"
	"net/http"
)

func CreateResponse(w http.ResponseWriter, responseText string, responseSpeech string) ResponseAssistant {

	simpleResponse := SimpleResponse{
		TextToSpeech: responseSpeech,
		DisplayText:  responseText,
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

	response := ResponseAssistant{
		Payload: payload,
	}

	handlers.HandlerInterfaceAssistant(w, response)

	return response
}

type ResponseAssistant struct {
	Payload Payload `json:"payload"`
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
