package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	openapi "github.com/sfrand8/over-engineered-calculator/pkg/http-client"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"over-engineered-calculator/internal"
	"testing"
)

func TestRoutes(t *testing.T) {
	// Set up your router (this is where you call the same function used in main)
	r := internal.SetupHttp()

	// Create a test server with the router
	ts := httptest.NewServer(r)
	defer ts.Close()
	testClient := createTestClient(ts.URL)

	t.Run("Test Swagger Route", func(t *testing.T) {
		// Create the request
		req, err := http.NewRequest(http.MethodGet, ts.URL+"/swagger/index.html", nil)
		assert.NoError(t, err)

		// Execute the request
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		// Assert the status code
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Calculate expression for user returns result", func(t *testing.T) {
		// Create the request
		var (
			userID, _ = uuid.NewUUID()
		)

		calculationResponse, httpResponse, err := testClient.CalculationsAPI.CalculatePost(
			context.Background()).
			Authorization(fmt.Sprintf("Bearer %s", userID.String())).
			Calculation(openapi.CalculateRequest{Expression: openapi.PtrString("2+2")}).
			Execute()

		// Assert the status code
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
		assert.Equal(t, "4", *calculationResponse.Result)
	})

	t.Run("Test API v1 History Route", func(t *testing.T) {
		// Create the request
		var (
			userID, _  = uuid.NewUUID()
			expression = "3*3"
		)
		calculationResponse, _, err := testClient.CalculationsAPI.CalculatePost(
			context.Background()).
			Authorization(fmt.Sprintf("Bearer %s", userID.String())).
			Calculation(openapi.CalculateRequest{Expression: openapi.PtrString(expression)}).
			Execute()

		assert.NoError(t, err)

		historyResult, httpResponse, err := testClient.CalculationsAPI.GetHistoryGet(
			context.Background()).
			Authorization(fmt.Sprintf("Bearer %s", userID.String())).
			Execute()

		// Assert the status code
		assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
		assert.Equal(t, expression, *historyResult.CalculationHistory[0].Expression)
		assert.Equal(t, *calculationResponse.Result, *historyResult.CalculationHistory[0].Result)
	})
}
