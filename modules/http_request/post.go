package http_request

import (
	//"../../modules/config"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

func POST(endpoint, request, bodystring string) ([]byte, error) {

	fmt.Printf("post: (%v, %v, %v)", endpoint, request, bodystring)

	client := &http.Client{}

	req, err := http.NewRequest("POST", endpoint+"/"+request, bytes.NewBuffer([]byte(bodystring)))

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Printf("response Status: (%v)", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	if resp.Status[:3] == "200" || resp.Status[:3] == "201" { } else {
		fmt.Printf("response Status: (%v)", resp.Status)
		fmt.Println("response Body:", string(body))
		var response interface{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("status " + resp.Status)
	}

	return body, nil
}


func POSTFLOW(endpoint, request, bodystring string) ([]byte, error) {

	fmt.Printf("post: (%v, %v, %v)", endpoint, request, bodystring)

	client := &http.Client{}

	req, err := http.NewRequest("POST", endpoint+"/"+request, bytes.NewBuffer([]byte(bodystring)))

	req.Header.Set("Content-Type", "application/json")

	gcloudKey, err := exec.Command("./../gcloud.sh").Output()
	if err != nil {
		fmt.Println("GCLOUD KEY ERROR:", err)
	}
	gcloudKeyStr := strings.Replace(string(gcloudKey), "\n", "", -1)
	fmt.Printf("GCLOUD KEY %s\n", gcloudKeyStr)

	req.Header.Set("Authorization", "Bearer " + gcloudKeyStr)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Printf("response Status: (%v)", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	if resp.Status[:3] == "200" || resp.Status[:3] == "201" { } else {
		fmt.Printf("response Status: (%v)", resp.Status)
		fmt.Println("response Body:", string(body))
		var response interface{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("status " + resp.Status)
	}

	return body, nil
}

