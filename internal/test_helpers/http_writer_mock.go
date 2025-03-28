package test_helpers

import (
	"bytes"
	"net/http"
)

type ResponseWriterMock struct {
	header     http.Header
	Body       *bytes.Buffer
	StatusCode int
}

func NewResponseWriterMock() *ResponseWriterMock {
	return &ResponseWriterMock{
		header:     http.Header{},
		Body:       new(bytes.Buffer),
		StatusCode: http.StatusOK,
	}
}

func (rw *ResponseWriterMock) Header() http.Header {
	return rw.header
}

func (rw *ResponseWriterMock) Write(b []byte) (int, error) {
	return rw.Body.Write(b)
}

func (rw *ResponseWriterMock) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
}
