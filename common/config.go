// Imported over
package common

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	SiteInstance string
	SiteId       string
	ApiKey       string
}

func GetConfigFile() Configuration {
	file, _ := os.Open("./resources/conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}

	return configuration
}
