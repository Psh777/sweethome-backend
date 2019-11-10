package security

import (
	"../../modules/config"
	"../../modules/http_request"
	"encoding/json"
	"fmt"
)


// action

func SetOn() (string, error) {

	fmt.Println("GO on")
	c := config.GetMyConfig()
	ansJson, err := http_request.GET(c.Env.SecurityBackend, "security/on")

	fmt.Println(string(ansJson), err)
	if err != nil {
		mess := "Error: " + err.Error()
		var answer Answer
		err = json.Unmarshal(ansJson, &answer)
		str := fmt.Sprint(mess, " / Code: ", answer.Code)
		return str, err
	}

	var answer Answer
	_ = json.Unmarshal(ansJson, &answer)

	return answer.Result, nil

}

func SetOff() (string, error) {

	fmt.Println("GO off")
	c := config.GetMyConfig()
	ansJson, err := http_request.GET(c.Env.SecurityBackend, "security/off")

	fmt.Println(string(ansJson), err)
	if err != nil {
		mess := "Error: " + err.Error()
		var answer Answer
		err = json.Unmarshal(ansJson, &answer)
		str := fmt.Sprint(mess, " / Code: ", answer.Code)
		return str, err
	}

	var answer Answer
	_ = json.Unmarshal(ansJson, &answer)

	return answer.Result, nil

}

type Answer struct {
	Success bool   `json:"success"`
	Result  string `json:"result"`
	Code    string `json:"code"`
}
