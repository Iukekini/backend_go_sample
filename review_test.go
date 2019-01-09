package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReviews(t *testing.T) {

	//Read Sample data into local var
	dat, _ := ioutil.ReadFile("sample_html/sample.html")
	SampleHTML := string(dat)

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		// Send response to be tested
		rw.Write([]byte(SampleHTML))
	}))
	// Close the server when test finishes
	defer server.Close()

	result := GetReviews(server.Client(), server.URL, 1)

	//make sure we got all rows back we expected
	assert.Equal(t, 10, len(result))
}

func TestGetReviews_WrongPage(t *testing.T) {

	//Read Sample data into local var
	dat, _ := ioutil.ReadFile("sample_html/Empty.html")
	SampleHTML := string(dat)

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		// Send response to be tested
		rw.Write([]byte(SampleHTML))
	}))
	// Close the server when test finishes
	defer server.Close()

	result := GetReviews(server.Client(), server.URL, 1)

	//make sure we got all rows back we expected
	assert.Equal(t, 0, len(result))
}
