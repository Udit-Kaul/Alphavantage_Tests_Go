package client

import (
	"fmt"
	"net/http"
)

type AlphavantageClient struct {
	BaseURL string
	APIKey  string
}

func NewClient(apiKey string) *AlphavantageClient {
	return &AlphavantageClient{
		BaseURL: "https://www.alphavantage.co/query",
		APIKey:  apiKey,
	}
}

func (c *AlphavantageClient) FetchStock(symbol string) (*http.Response, error) {
	url := fmt.Sprintf("%s?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", c.BaseURL, symbol, c.APIKey) // Fix URL format string
	return http.Get(url)
}
