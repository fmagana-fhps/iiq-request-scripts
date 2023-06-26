package api

import (
	"common/models"
	"common/requests"
	"encoding/json"
	"strings"
)

func SendNewVerification(assetId, payloads string) error {
	body, err := requests.Post("assets/"+assetId+"/verifications/new", strings.NewReader(payloads))
	if err != nil {
		return err
	}

	resVerify := models.AssetVerificationResponse{}
	err = json.Unmarshal(body, &resVerify)

	if err != nil {
		return err
	}

	return nil
}
