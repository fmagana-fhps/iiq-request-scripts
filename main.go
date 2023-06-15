package main

import (
	"bufio"
	"encoding/json"
	"fmt"
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

		asset := sendGetRequest(string(key), serials[2], "")
		asset.AssetTag = serials[1]
		asset.SerialNumber = serials[0]
		body, _ := json.Marshal(asset)
		fmt.Println(string(body))

		sendPostRequest(string(key), serials[2], string(body))
		fmt.Println("200")
	}
}
