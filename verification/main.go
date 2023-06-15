package main

import (
	"bufio"
	"fmt"
	"os"
	"verification/verify"
)

func main() {
	key, _ := os.ReadFile(os.Args[1])
	assetsFile, err := os.Open(os.Args[2])

	if err != nil {
		panic(err)
	}
	defer assetsFile.Close()

	scanner := bufio.NewScanner(assetsFile)
	for scanner.Scan() {
		asset := scanner.Text()
		assetId := verify.GetAssetIdFromAssetTag(string(key), asset, "")

		v := verify.GetVerifications(string(key), assetId, "")

		fmt.Println(v)
	}
}
