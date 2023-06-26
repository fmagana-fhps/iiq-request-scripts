package requests

import (
	"common"
	"io"
	"net/http"
)

func newRequest(method, params string, payload io.Reader) (*http.Request, error) {
	cfg := common.GetConfigFile()
	url := "https://" + cfg.SiteInstance + params
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("SiteId", cfg.SiteId)
	req.Header.Add("Authorization", "Bearer "+cfg.ApiKey)
	req.Header.Add("Client", "ApiClient")

	return req, nil
}
