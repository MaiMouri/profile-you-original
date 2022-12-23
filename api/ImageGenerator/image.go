package imagegenerator

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

func ImageGenerator(keyword string) []Data {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	// token := config.Config.ApiKey
	token := os.Getenv("OPEN_API_KEY")

	requestBody := &RequestBody{
		Prompt: keyword,
		Number: 1,
		Size:   "512x512",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		panic("Error")
	}
	// fmt.Println(string(jsonString))
	endPoint := "https://api.openai.com/v1/images/generations" //APIのエンドポイントを指定(このURLでOK)
	req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonString))
	if err != nil {
		panic("Error")
	}

	bearer := "Bearer " + token

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)

	// // Header information
	// dump, _ := httputil.DumpRequestOut(req, true)
	// fmt.Printf("%s", dump)

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

	// fmt.Printf("%#v\n", string(bytes))

	b := bytes
	var d DallEAPIResponse
	// var t Data
	if err := json.Unmarshal(b, &d); err != nil {
		panic(err)
	}
	// fmt.Println(&d)
	return d.Data
}
