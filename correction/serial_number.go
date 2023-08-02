package main

import (
	"bufio"
	"correction/api"
	"encoding/json"
	"os"
	"strings"
)

func CorrctSerialNumber() {
	assetFile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer assetFile.Close()

	scanner := bufio.NewScanner(assetFile)
	for scanner.Scan() {
		serials := strings.Split(scanner.Text(), ",")
		asset, err := api.GetAsset(serials[2])
		if err != nil {
			panic(err)
		}

		asset.AssetTag = serials[1]
		asset.SerialNumber = serials[0]
		body, _ := json.Marshal(asset)

		go api.UpdateAsset(serials[2], string(body))
	}
}
