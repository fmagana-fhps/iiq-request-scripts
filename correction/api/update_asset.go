package api

import (
	"common/requests"
	"strings"
)

func UpdateAsset(assetId string, payloads string) error {
	payload := strings.NewReader(payloads)
	_, err := requests.Post("assets/"+assetId, payload)
	if err != nil {
		return err
	}
	return nil
}
