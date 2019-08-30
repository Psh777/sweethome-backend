package http_request

import (
	//"../../modules/config"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
)

func POST(endpoint, request, bodystring string) ([]byte, error) {

	fmt.Printf("post: (%v, %v, %v)", endpoint, request, bodystring)

	//client := ProxyHttpClient()
	client := &http.Client{}

	req, err := http.NewRequest("POST", endpoint+"/"+request, bytes.NewBuffer([]byte(bodystring)))

	//req.Header.Set("api-key", apikey)
	req.Header.Set("Content-Type", "application/json")

	gcloudKey, err := exec.Command("./gcloud.sh").Output()
	if err != nil {
		fmt.Println("GCLOUD KEY ERROR:", err)
	}
	fmt.Printf("GCLOUD KEY %s\n", string(gcloudKey))

	req.Header.Set("Authorization", "Bearer " + string(gcloudKey))

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

func ProxyHttpClient() *http.Client {

	//c := config.GetMyConfig()

	proxyUrl, err := url.Parse("socks5h://10.30.0.2:1070")
	if err != nil {
		fmt.Printf("Invalid proxy url %v\n", proxyUrl)
	}

	//dialer, err := proxy.SOCKS5("tcp", "10.30.0.2:1070", nil, proxy.Direct)

	httpTransport := &http.Transport{
		//DialTLS: dialer.Dial,
		Proxy: http.ProxyURL(proxyUrl),
	}

	httpClient := &http.Client{
		Transport: httpTransport,
	}

	return httpClient

}
