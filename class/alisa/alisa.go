package alisa

import (
	"../../webserver/handlers"
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJson(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		fmt.Println("error: no body")
		handlers.HandlerError(w, "No body")
		return
	}

	var v interface{}
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		fmt.Println(err)
		//return v, err
	}

	fmt.Printf("%+v\n", v)

}