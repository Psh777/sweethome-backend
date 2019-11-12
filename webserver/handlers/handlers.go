package handlers

import (
	"../../modules/lang"
	"encoding/json"
	"net/http"
)

func Headers(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "content-type,X-Requested-With")
	return w
}

func HandlerError(w http.ResponseWriter, codeError string) {
	w = Headers(w)
	textError := lang.CodeError(codeError)
	err := Error{false, codeError, textError, nil}
	Json, _ := json.Marshal(err)
	w.WriteHeader(400)
	_, _ = w.Write(Json)
}

func HandlerError200(w http.ResponseWriter, codeError string) {
	w = Headers(w)
	textError := lang.CodeError(codeError)
	Error := Error{false, codeError, textError, nil}
	Json, _ := json.Marshal(Error)
	w.WriteHeader(400)
	_, _ = w.Write(Json)
}

func HandlerSuccess(w http.ResponseWriter, successText string) {
	w = Headers(w)
	Success := Success{true, successText}
	Json, _ := json.Marshal(Success)
	w.WriteHeader(200)
	_, _ = w.Write(Json)
}

func HandlerInterface(w http.ResponseWriter, data interface{}) {
	w = Headers(w)
	out := Success{
		Success: true,
		Result:  data,
	}
	Json, _ := json.Marshal(out)
	_, _ = w.Write(Json)
}

func HandlerInterfaceAssistant(w http.ResponseWriter, data interface{}) {
	w = Headers(w)
	Json, _ := json.MarshalIndent(data, "", " ")
	_, _ = w.Write(Json)
}

func HandlerInterfaceError(w http.ResponseWriter, data interface{}) {
	w = Headers(w)
	out := Success{
		Success: false,
		Result:  data,
	}
	Json, _ := json.Marshal(out)
	w.WriteHeader(400)
	_, _ = w.Write(Json)
}

func HandlerPrint(w http.ResponseWriter, data string) {
	w.Header().Set("Content-Type", "Content-Type: text/html; charset=utf-8")
	_, _ = w.Write([]byte(data))
}
