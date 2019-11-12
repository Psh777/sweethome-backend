package alisa

import (
"../../webserver/handlers"
"encoding/json"
"fmt"
"io/ioutil"
"net/http"
)

func DeviceState(w http.ResponseWriter, r *http.Request) {

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

	fmt.Printf("%+v\n", string(b))

	var t actionRequest

	err = json.Unmarshal(b, &t)
	if err != nil {
		fmt.Println()
		fmt.Println("error: " + err.Error())
		handlers.HandlerError(w, err.Error())
		return
	}

	fmt.Printf("%+v\n", t)

	//action

	handlers.HandlerInterfaceAssistant(w, "ok")

}
