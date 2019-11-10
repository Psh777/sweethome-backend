package assistant

import (
	"../../class/security"
	"../../db/postgres"
	"../../webserver/handlers"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
	case "multisensor":

		intTypeSensor, err := strconv.Atoi(t.QueryResult.Parameters.TypeSensor)
		if err != nil {
			return
		}

		data, err := postgres.GetDataByTypeActually(t.QueryResult.Parameters.Room, intTypeSensor)
		if err != nil {
			return
		}

		var prefix string

		switch intTypeSensor {
		case 1:
			prefix = "градусов"
		case 2:
			prefix = "процентов"
		case 3:
			prefix = "единиц ртутного столба"
		case 5:
			prefix = "ппм"
		}

		str := "" + fmt.Sprintf("%.0f", data) + " " + prefix
		CreateResponse(w, str, str)

	case "light":

		//sonoff.Switch(t.QueryResult.Parameters.SwitchState)
		CreateResponse(w, "Готово", "Готово")

	case "security":

		switch t.QueryResult.Parameters.SwitchState {
		case "on":
			msg, err := security.SetOn()
			if err != nil {
				CreateResponse(w, err.Error(), err.Error())
			} else {
				CreateResponse(w, msg, msg)
			}

		case "off":
			msg, err := security.SetOff()
			if err != nil {
				CreateResponse(w, err.Error(), err.Error())
			} else {
				CreateResponse(w, msg, msg)
			}

		}
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
	TypeSensor  string `json:"type-sensor"`
	Room        string `json:"room"`
	SwitchState string `json:"switch-state"`
}
