package main

import (
	"common/requests"
	"encoding/json"
	"fmt"
	"log"
	"os"

	iiq "github.com/fmagana-fhps/incidentiq-api-go"
	"github.com/fmagana-fhps/incidentiq-api-go/models"
	"github.com/joho/godotenv"
)

var client *iiq.Client

func init() {
	err := godotenv.Load("./resources/.env.local")
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
	f := []byte(`"School Name","Serial Number"`)

	res, err := requests.Post("assets/?$s=500&$o=AssetTag&$filter=(CategoryID%20eq%20'%5B7031883c-a881-e611-80f3-000c29ab80b0%5D')", "")
	if err != nil {
		panic(err)
	}

	schema := models.ItemsResponse[models.Asset]{}
	json.Unmarshal(res, &schema)
	carts := schema.Items

	for _, cart := range carts {
		if cart.AssetID == "" {
			continue
		}
		params := iiq.Parameters{PageSize: 100}
		links, err := client.GetLinkedAssets(params, cart.AssetID)
		if err != nil {
			log.Fatalf("%v", err)
		}
		for _, link := range links {
			if link.SerialNumber == "" {
				continue
			}
			f = append(f, []byte(fmt.Sprintf("\n\"%s\", %s", cart.Location.Name, link.SerialNumber))...)
		}
	}

	file, err := os.Create("asset-locations-test.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write(f)
}
