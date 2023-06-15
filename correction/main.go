package main

import (
	"bufio"
	"correction/requests"
	"encoding/json"
	"os"
	"strings"
)

func main() {
	key, _ := os.ReadFile(os.Args[1])
	var serials []string

	assetFile, err := os.Open(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer assetFile.Close()

	scanner := bufio.NewScanner(assetFile)
	for scanner.Scan() {
		serials = strings.Split(scanner.Text(), ",")

		asset := requests.GetAsset(string(key), serials[2], "")
		asset.AssetTag = serials[1]
		asset.SerialNumber = serials[0]
		body, _ := json.Marshal(asset)

		requests.UpdateAsset(string(key), serials[2], string(body))
	}
}
