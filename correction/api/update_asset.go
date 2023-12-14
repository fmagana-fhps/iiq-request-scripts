// Imported over
package api

import (
	"common/requests"
)

func UpdateAsset(assetId string, payload string) error {
	_, err := requests.Post("assets/"+assetId, payload)
	if err != nil {
		return err
	}
	return nil
}
