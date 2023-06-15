package main

import (
	c "common"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func sendGetRequest(apiKey string, assetId string, payloads string) c.Item {

	if apiKey == "" {
		panic("No API Key Provided")
	}

	payload := strings.NewReader(payloads)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://fhps.incidentiq.com/api/v1.0/assets/"+assetId, payload)

	if err != nil {
		panic(err)
	}

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	req.Header.Add("Client", "WebBrowser")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Accept-Encoding", "gzip deflate, br")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("Authorization", apiKey)
	req.Header.Add("SiteId", "bb843135-5c45-4ca2-a1dd-88530abd42f9")
	req.Header.Add("ProductId", "88df910c-91aa-e711-80c2-0004ffa00010")

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	resAsset := c.Response{}
	err = json.Unmarshal(body, &resAsset)
	if err != nil {
		panic(err)
	}

	return resAsset.Item
}

func sendPostRequest(apiKey string, assetId string, payloads string) {

	if apiKey == "" {
		panic("No API Key Provided")
	}

	payload := strings.NewReader(payloads)

	req, err := http.NewRequest("POST", "https://fhps.incidentiq.com/api/v1.0/assets/"+assetId, payload)

	if err != nil {
		panic(err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", apiKey)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}
