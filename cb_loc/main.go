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
	contents := []byte(`Asset Tag,Device Purpose,Location,Notes,Purchased Date,Purchased Price,PO Number,Deployed Date,Warranty Expiration Date,Invoice Number,Vendor,Asset Status`)

	res, err := requests.Post("assets/?$s=500&$o=AssetTag&$filter=(CategoryID%20eq%20'%5B7031883c-a881-e611-80f3-000c29ab80b0%5D')", "")
	if err != nil {
		panic(err)
	}

	schema := models.ItemsResponse[models.Asset]{}
	json.Unmarshal(res, &schema)
	carts := schema.Items

	for _, cart := range carts {
		if cart.AssetID == "" || cart.Location.Name != "Knapp Forest Elementary School" {
			continue
		}
		params := iiq.Parameters{PageSize: 100}
		resp, err := client.LinkedAssets(params, cart.AssetID)
		if err != nil {
			log.Fatalf("%v", err)
		}
		for _, cb := range resp.Body.Items {
			if cb.SerialNumber == "" {
				continue
			}

			tpl := "\n%v,Student,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v"
			row := fmt.Sprintf(tpl, cb.AssetTag, cart.Location.Name, cb.Notes, cb.PurchasedDate, cb.PurchasePrice, cb.PurchasePoNumber,
				cb.DeployedDate, cb.WarrantyExpirationDate, cb.InvoiceNumber, cb.Vendor, cb.Status.Name)
			contents = append(contents, []byte(row)...)
		}
	}

	file, err := os.Create("kf-cb-carts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write(contents)
}
