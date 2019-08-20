package assistant

import (
	"../../webserver/handlers"
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJson(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	fmt.Printf("%+v\n", r.Body)
	var t request
	err := decoder.Decode(&t)
	if err != nil {
		return
	}
	//fmt.Printf("%+v\n", t)
	handlers.HandlerInterface(w, "ok")
}

type request struct {

}