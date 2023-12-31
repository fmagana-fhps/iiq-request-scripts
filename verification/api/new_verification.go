// Imported over
package api

import (
	"common/models"
	"common/requests"
	"encoding/json"
	"fmt"
)

func SendNewVerification(assetId, payload string) error {
	body, err := requests.Post("assets/"+assetId+"/verifications/new", payload)
	if err != nil {
		return err
	}

	resVerify := models.AssetVerificationResponse{}
	err = json.Unmarshal(body, &resVerify)

	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
