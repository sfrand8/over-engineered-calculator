package calculate

import (
	"bytes"
	"encoding/json"
	http_middleware "github.com/sfrand8/over-engineered-calculator/pkg/http-middleware"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"over-engineered-calculator/internal/database"
	"over-engineered-calculator/internal/helpers"
	"testing"
)

func TestCreateCalculateHandler(t *testing.T) {
	var createHttpRequest = func(request Request, userID string) *http.Request {
		reqBytes, _ := json.Marshal(request)
		req := &http.Request{
			Body: io.NopCloser(bytes.NewBuffer(reqBytes)),
		}
		return req.WithContext(http_middleware.ContextWithUserID(req.Context(), userID))
	}
	t.Run("returns result of calculation when everything succeeds", func(t *testing.T) {
		var (
			expectedResult   = "someResult"
			expressionString = "someExpression"
			userID           = "someUserID"
			calculatorMock   = &calculatorMock{
				PerformCalculationForUserFunc: func(userID string, expressionString string) (database.Calculation, error) {
					return database.Calculation{
						Expression: expressionString,
						Result:     expectedResult,
					}, nil
				},
			}
			request = createHttpRequest(Request{Expression: expressionString}, userID)
			rr      = newResponseWriterMock()
			sut     = createCalculateHandler(calculatorMock)
		)

		// act
		sut(rr, request)

		// assert
		assert.Equal(t, http.StatusOK, rr.statusCode)
		var resp Response
		err := json.NewDecoder(rr.body).Decode(&resp)
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, resp.Result)
	})

	t.Run("returns bad request when calculator returns InvalidExpressionError", func(t *testing.T) {
		var (
			expectedError    = "Invalid expression"
			expressionString = "someExpression"
			userID           = "someUserID"
			calculatorMock   = &calculatorMock{
				PerformCalculationForUserFunc: func(userID string, expressionString string) (database.Calculation, error) {
					return database.Calculation{}, InvalidExpressionError
				},
			}
			request = createHttpRequest(Request{Expression: expressionString}, userID)
			rr      = newResponseWriterMock()
			sut     = createCalculateHandler(calculatorMock)
		)

		// act
		sut(rr, request)

		// assert
		assert.Equal(t, http.StatusBadRequest, rr.statusCode)
		var resp helpers.ErrorResponse
		err := json.NewDecoder(rr.body).Decode(&resp)
		assert.NoError(t, err)
		assert.Equal(t, expectedError, resp.Error)
	})

	t.Run("returns internal server error when calculator returns an error", func(t *testing.T) {
		var (
			expectedError    = "something went wrong"
			expressionString = "someExpression"
			userID           = "someUserID"
			calculatorMock   = &calculatorMock{
				PerformCalculationForUserFunc: func(userID string, expressionString string) (database.Calculation, error) {
					return database.Calculation{}, assert.AnError
				},
			}
			request = createHttpRequest(Request{Expression: expressionString}, userID)
			rr      = newResponseWriterMock()
			sut     = createCalculateHandler(calculatorMock)
		)

		// act
		sut(rr, request)

		// assert
		assert.Equal(t, http.StatusInternalServerError, rr.statusCode)
		var resp helpers.ErrorResponse
		err := json.NewDecoder(rr.body).Decode(&resp)
		assert.NoError(t, err)
		assert.Equal(t, expectedError, resp.Error)
	})
}

type responseWriterMock struct {
	header     http.Header
	body       *bytes.Buffer
	statusCode int
}

func newResponseWriterMock() *responseWriterMock {
	return &responseWriterMock{
		header:     http.Header{},
		body:       new(bytes.Buffer),
		statusCode: http.StatusOK,
	}
}

func (rw *responseWriterMock) Header() http.Header {
	return rw.header
}

func (rw *responseWriterMock) Write(b []byte) (int, error) {
	return rw.body.Write(b)
}

func (rw *responseWriterMock) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
}
