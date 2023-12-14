// Imported over
package requests

import (
	"io"
	"net/http"
)

func Post(url, payload string) ([]byte, error) {
	req, err := newIIQRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Accept-Encoding", "gzip deflate, br")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return body, nil
}
