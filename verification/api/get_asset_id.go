package api

import (
	"common/models"
	"common/requests"
	"encoding/json"
	"strings"
)

func GetAssetIdFromAssetTag(assetTag string) string {
	body, err := requests.Get("assets/assettag/"+assetTag, strings.NewReader(``))
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
