package lang

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

var codeError map[string]string

func ReadFileCodeError() error {

	file, err1 := ioutil.ReadFile("./static/lang/message_error.en.json")
	if err1 != nil {
		fmt.Println("error : ", file, err1)
		return err1
	}

	c := make(map[string]string)

	e := json.Unmarshal(file, &c)
	if e != nil {
		return e
	}

	codeError = c

	return nil
}

func CodeError(code string) string {
	return codeError[code]
}
