package requests

import (
	"io"
	"net/http"
)

func Get(url string, payload io.Reader) ([]byte, error) {
	req, err := newRequest("GET", url, payload)
	if err != nil {
		return nil, err
	}

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
