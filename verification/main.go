package main

import (
	"bufio"
	"common"
	"common/models"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"verification/api"

	"github.com/google/uuid"
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
		fmt.Println(asset, assetId)

		avid := uuid.New()

		verification := models.AssetVerification{
			AssetVerificationId:     avid.String(),
			AssetId:                 assetId,
			CreatedDate:             "",
			VerifiedByUserId:        "10747f8a-b90b-ee11-907c-000d3a054ab7",
			AssetVerificationTypeId: "web-manual", // could leave out but it still auto fills the same
			IsSuccessful:            true,
			Comments:                "Verified using the API",
			LocationId:              locationId,
		}

		payload, err := json.Marshal(verification)
		if err != nil {
			panic(err)
		}

		api.SendNewVerification(assetId, string(payload))
	}
}
