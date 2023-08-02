package requests

import (
	"common"
	"net/http"
	"strings"
)

func newIIQRequest(method, params, payload string) (*http.Request, error) {
	cfg := common.GetConfigFile()
	url := "https://" + cfg.SiteInstance + params
	req, err := http.NewRequest(method, url, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("SiteId", cfg.SiteId)
	req.Header.Add("Authorization", "Bearer "+cfg.ApiKey)
	req.Header.Add("Client", "ApiClient")

	return req, nil
}

func NewRequest(method, site, params, payload string) (*http.Request, error) {
	url := "https://" + site + params
	req, err := http.NewRequest(method, url, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	return req, nil
}
