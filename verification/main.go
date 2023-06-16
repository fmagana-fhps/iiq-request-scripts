package main

import (
	"bufio"
	"common"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"verification/verify"

	"github.com/google/uuid"
)

func main() {
	key, err := os.ReadFile("notouch.txt")
	if err != nil {
		panic(`DO NOT TOUCH THIS FILE, YOU BROKE IT`)
	}
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
		assetId := verify.GetAssetIdFromAssetTag(string(key), asset)

		avid := uuid.New()

		verification := common.AssetVerification{
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

		v := verify.SendNewVerification(string(key), assetId, string(payload))

		fmt.Println(asset, "audited on", v.AssetVerification.CreatedDate)
	}
}
