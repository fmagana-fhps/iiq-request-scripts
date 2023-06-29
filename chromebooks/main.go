package main

import (
	"chromebooks/api"
	"fmt"
	"strings"
)

func main() {
	assets := api.GetDellChromebooks()

	for _, asset := range assets {
		if strings.Contains(asset.Name, "Dell") {
			fmt.Println(asset.AssetTag, asset.SerialNumber)
		}
	}
}
