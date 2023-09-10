package coinapp

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type HttpClient struct {
	Client *http.Client
}

func NewHTTPClient(timeout time.Duration) (*HttpClient, error) {
	return &HttpClient{
		Client: &http.Client{
			Timeout: timeout,
			Transport: &LoggerRoundTripper{
				next: http.DefaultTransport,
			},
		},
	}, nil
}

func (c HttpClient) GetAssets() Asset {
	req, err := http.NewRequest("GET", "https://api.coincap.io/v2/assets?jane=qwerty", nil)
	resp, err := c.Client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	// It is necessary to close the body of the answer (at the end), because memory is allocated for its maintenance
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var assetResp Asset
	err = json.Unmarshal(body, &assetResp)

	if err != nil {
		log.Fatal(err)
	}
	return assetResp
}
