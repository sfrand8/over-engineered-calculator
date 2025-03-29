package main

import (
	"fmt"
	openapi "github.com/sfrand8/over-engineered-calculator/pkg/http-client"
)

func createTestClient(url string) *openapi.APIClient {
	cfg := openapi.NewConfiguration()
	cfg.Servers = openapi.ServerConfigurations{
		{
			URL: fmt.Sprintf("%s/api/v1", url),
		},
	}
	return openapi.NewAPIClient(cfg)
}
