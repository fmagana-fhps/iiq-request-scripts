// Imported over
package api

import (
	"common/models"
	"common/requests"
	"encoding/json"
)

func GetAsset(assetId string) (models.Asset, error) {
	body, err := requests.Get("assets/"+assetId, "")
	if err != nil {
		return models.Asset{}, err
	}

	asset := models.SingleAsset{}
	err = json.Unmarshal([]byte(body), &asset)
	if err != nil {
		panic(err)
	}

	return asset.Asset, nil
}
