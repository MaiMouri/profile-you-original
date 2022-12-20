package imagegenerator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"profileyou/api/config"
)

type DallEAPIResponse struct {
	Created int    `json:"created"`
	Data    []Data `json:"data"`
}

//

type Data struct {
	Url string `json:"url"`
}

//

type RequestBody struct {
	Prompt string `json:"prompt"`
	Number int    `json:"n"`
	Size   string `json:"size"`
}

// func (d Data) String() string {
// 	return fmt.Sprintf(d.Url)
// }

func ImageGenerator(keyword string) []Data {
	token := config.Config.ApiKey

	requestBody := &RequestBody{
		Prompt: keyword,
		Number: 1,
		Size:   "512x512",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		panic("Error")
	}
	fmt.Println(string(jsonString))
	endPoint := "https://api.openai.com/v1/images/generations" //APIのエンドポイントを指定(このURLでOK)
	req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonString))
	if err != nil {
		panic("Error")
	}

	bearer := "Bearer " + token

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)

	dump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s", dump)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		panic("Error")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", string(bytes))

	b := bytes
	var d DallEAPIResponse
	// var t Data
	if err := json.Unmarshal(b, &d); err != nil {
		panic(err)
	}
	fmt.Println(&d)
	return d.Data
}
