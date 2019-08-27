package http_request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GET(endpoint, request string) ([]byte, error) {

	req, err := http.NewRequest("GET", endpoint + "/" + request, bytes.NewBuffer(nil))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//fmt.Printf("response Status: (%v)", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	if resp.Status[:3] != "200" {
		fmt.Printf("response Status: (%v)", resp.Status)
		fmt.Println("response Body:", string(body))
		var response interface{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return body, err
		}
		return body, errors.New("status " + resp.Status)
	}

	return body, nil
}