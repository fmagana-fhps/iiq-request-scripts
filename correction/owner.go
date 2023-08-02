package main

import (
	"common/models"
	"common/requests"
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

func ChangePreviousToCurrentOwner() {
	res, err := requests.Post("assets/?$s=20000&$o=AssetTag&$d=Ascending&$filter=(ManufacturerId%20eq%20'%5B518000c0-4dff-e511-a789-005056bb000e%5D')", "")
	if err != nil {
		panic(err)
	}

	schema := models.MultipleAssets{}
	json.Unmarshal(res, &schema)
	assets := schema.Assets

	var updatedAssets []models.Asset

	for idx, asset := range assets {

		if asset.PreviousOwnerID != "" && asset.OwnerID == "" {
			assets[idx].OwnerID = asset.PreviousOwnerID
			assets[idx].Owner = asset.PreviousOwner

			assets[idx].PreviousOwnerID = ""
			assets[idx].PreviousOwner = models.Owner{}

			updatedAssets = append(updatedAssets, assets[idx])
		}
	}

	fmt.Println("updating:", len(updatedAssets), "assets")

	var wg2 sync.WaitGroup
	wg2.Add(len(updatedAssets))

	for _, asset := range updatedAssets {

		body, err := json.Marshal(asset)
		if err != nil {
			panic(err)
		}

		go func(a models.Asset, data string) {
			defer wg2.Done()

			_, err := requests.Post("assets/"+a.AssetId, data)

			if err != nil {
				log.Println(err.Error())
				return
			}

			fmt.Println("updated:", a.AssetTag)
		}(asset, string(body))
	}

	wg2.Wait()

	fmt.Println("done")
}
