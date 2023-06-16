package verify

import (
	"common"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func SendNewVerification(apiKey, assetId, payloads string) common.AssetVerificationResponse {
	if apiKey == "" {
		panic("No API Key Provided")
	}

	req, err := http.NewRequest("POST", "https://fhps.incidentiq.com/api/v1.0/assets/"+assetId+"/verifications/new", strings.NewReader(payloads))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Accept-Encoding", "gzip deflate, br")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	resVerify := common.AssetVerificationResponse{}
	err = json.Unmarshal(body, &resVerify)

	if err != nil {
		panic(err)
	}

	return resVerify
}
