package api

import (
	"common/models"
	"common/requests"
	"encoding/json"
)

func GetDellChromebooks() []models.Asset {
	res, err := requests.Post("assets/?$s=12000&$o=AssetTag&$d=Ascending&$filter=(CategoryId%20eq%20'%5B36C6BE98-4F39-E611-BF3B-005056BB000E%5D')", "")
	if err != nil {
		panic(err)
	}

	schema := models.MultipleAssets{}
	err = json.Unmarshal(res, &schema)
	if err != nil {
		panic(err)
	}
	return schema.Assets
}
