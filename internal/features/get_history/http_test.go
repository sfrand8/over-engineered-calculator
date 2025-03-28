﻿package get_history

import (
	"bytes"
	"encoding/json"
	http_middleware "github.com/sfrand8/over-engineered-calculator/pkg/http-middleware"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"over-engineered-calculator/internal/database"
	"over-engineered-calculator/internal/helpers"
	"testing"
)

func TestCreateHistoryHandler(t *testing.T) {
	var createHttpRequest = func(userID string) *http.Request {
		req := &http.Request{
			Body: io.NopCloser(bytes.NewBuffer(nil)),
		}
		return req.WithContext(http_middleware.ContextWithUserID(req.Context(), userID))
	}

	t.Run("returns calculation history when everything succeeds", func(t *testing.T) {
		var (
			userID          = "someUserID"
			expectedHistory = []database.Calculation{
				{Expression: "2+2", Result: "4"},
				{Expression: "3*3", Result: "9"},
			}
			historyRetrieverMock = &calculationHistoryRetrieverMock{
				GetHistoryFunc: func(userID string) ([]database.Calculation, error) {
					return expectedHistory, nil
				},
			}
			request          = createHttpRequest(userID)
			rr               = httptest.NewRecorder()
			sut              = createHistoryHandler(historyRetrieverMock)
			expectedResponse = Response{
				CalculationHistory: []HistoryEntry{
					{Expression: expectedHistory[0].Expression, Result: expectedHistory[0].Result},
					{Expression: expectedHistory[1].Expression, Result: expectedHistory[1].Result},
				},
			}
		)

		// act
		sut(rr, request)

		// assert
		assert.Equal(t, http.StatusOK, rr.Code)
		var resp Response
		err := json.NewDecoder(rr.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, resp)
		assert.Equal(t, rr.Header()["Content-Type"][0], "application/json")
	})

	t.Run("returns internal server error when history retriever fails", func(t *testing.T) {
		var (
			userID               = "someUserID"
			expectedError        = "something went wrong"
			historyRetrieverMock = &calculationHistoryRetrieverMock{
				GetHistoryFunc: func(userID string) ([]database.Calculation, error) {
					return nil, assert.AnError
				},
			}
			request = createHttpRequest(userID)
			rr      = httptest.NewRecorder()
			sut     = createHistoryHandler(historyRetrieverMock)
		)

		// act
		sut(rr, request)

		// assert
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		var resp helpers.ErrorResponse
		err := json.NewDecoder(rr.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.Equal(t, expectedError, resp.Error)
		assert.Equal(t, rr.Header()["Content-Type"][0], "application/json")
		assert.Len(t, historyRetrieverMock.GetHistoryCalls(), 1)
		assert.Equal(t, historyRetrieverMock.GetHistoryCalls()[0].UserID, userID)
	})
}
