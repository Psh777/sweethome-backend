package sensor

import (
	"../../db/postgres"
	"../../types"
	"../../webserver/handlers"
	"encoding/json"
	"fmt"
	valid "github.com/asaskevich/govalidator"
	"net/http"
)

const (
	TEMP = iota
	HUM
	MMH
	PA
	CO2
	TVOC
)

func Upload(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	t := types.Sensor{}
	err := d.Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
	} else {
		//fmt.Printf("%+v\n", t)

		if !valid.IsUUID(t.SensorID) {
			handlers.HandlerError(w, "sensor ID invalid")
			return
		}

		if !valid.IsUUID(t.RequestID) {
			handlers.HandlerError(w, "request ID invalid")
			return
		}

		err = postgres.NewData(t)
		handlers.HandlerSuccess(w, "ok")
	}
}
