package main

import (
	"bufio"
	"common"
	"common/models"
	"encoding/json"
	"os"
	"strings"
	"verification/api"
)

func main() {
	assetsFile, err := os.Open("assets.txt")
	if err != nil {
		panic(`The file name needs to be: "assets.txt"`)
	}
	defer assetsFile.Close()

	locationId, present := common.GetLocationIds()[strings.ToUpper(os.Args[1])]
	if !present {
		panic("You need to provide a location using it's abbreviation")
	}

	scanner := bufio.NewScanner(assetsFile)
	for scanner.Scan() {
		asset := scanner.Text()
		assetId := api.GetAssetIdFromAssetTag(asset)

		verification := models.AssetVerification{
			VerifiedByUserId: "10747f8a-b90b-ee11-907c-000d3a054ab7",
			IsSuccessful:     true,
			LocationId:       locationId,
		}

		payload, err := json.Marshal(verification)
		if err != nil {
			panic(err)
		}

		api.SendNewVerification(assetId, string(payload))
	}
}
