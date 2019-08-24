package sonoff

import (
	"../../modules/http_request"
	"encoding/json"
	"fmt"
)

func Switch(state string) {

	data := Data{
		Switch: state,
	}

	req := Request{
		DeviceID: "10008fac66",
		Data:     data,
	}

	reqString, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = http_request.POST("http://93.123.132.168:9001", "zeroconf/switch", string(reqString))

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

type Request struct {
	DeviceID string `json:"deviceid"`
	Data     Data   `json:"data"`
}

type Data struct {
	Switch string `json:"switch"`
}
