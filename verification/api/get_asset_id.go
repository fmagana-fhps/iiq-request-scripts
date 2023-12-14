// Imported over
package api

import (
	"common/models"
	"common/requests"
	"encoding/json"
)

func GetAssetIdFromAssetTag(assetTag string) string {
	body, err := requests.Get("assets/assettag/"+assetTag, "")
	if err != nil {
		panic(err)
	}

	schema := models.MultipleAssets{}
	err = json.Unmarshal(body, &schema)
	if err != nil {
		panic(err)
	}

	return schema.Assets[0].AssetId
}
