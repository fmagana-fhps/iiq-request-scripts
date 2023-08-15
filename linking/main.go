package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	iiq "github.com/fmagana-fhps/incidentiq-api-go"
	"github.com/fmagana-fhps/incidentiq-api-go/models"
	"github.com/joho/godotenv"
)

var (
	client     *iiq.Client
	assetCache = make(map[string]models.Asset)
	links      = make(map[string][]string)
)

func init() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal(err)
	}

	client, err = iiq.NewClient(&iiq.Options{
		SiteId: os.Getenv("SITEID"),
		Token:  os.Getenv("TOKEN"),
		Domain: os.Getenv("DOMAIN"),
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Starting...")

	assetsFile, err := os.Open("Asset_export.csv")
	if err != nil {
		panic(`The file name needs to be: "Asset_export.txt"`)
	}
	defer assetsFile.Close()

	fmt.Println("File Opened...")

	scanner := bufio.NewScanner(assetsFile)
	fmt.Println("Reading file...")
	for scanner.Scan() {
		row := scanner.Text()
		split := strings.Split(row, ",")
		cart := split[0]
		cbs := strings.Split(split[1], ";")

		GetAssets(1, cart)
		GetAssets(len(cbs), cbs...)

		Link(cart, cbs...)
	}

	for k, v := range links {
		client.UnlinkAssets(k, v...)
	}
}

func Link(parent string, children ...string) {
	p := assetCache[parent].AssetID
	for _, child := range children {
		c := assetCache[child].AssetID
		if _, ok := links[p]; !ok {
			links[p] = []string{c}
			continue
		}
		links[p] = append(links[p], c)
	}
}

func GetAssets(length int, assetTags ...string) {
	need := make([]string, 0, length)
	for i := range assetTags {
		if _, ok := assetCache[assetTags[i]]; ok {
			continue
		}
		need = append(need, assetTags[i])
	}

	params := iiq.Parameters{PageSize: 50}
	assets, err := client.AssetsByAssetTag(params, need...)
	if err != nil {
		log.Fatal(err)
	}

	for i := range assets {
		assetCache[assets[i].AssetTag] = assets[i]
	}
}
