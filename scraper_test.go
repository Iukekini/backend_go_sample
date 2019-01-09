package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

//Test to make sure that the retrieval and parsing of the html
//works. Use the test http server instead of a live call. 
func TestRetrieveHtml(t *testing.T) {

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

	result := retrieveDocument(server.Client(), server.URL)

	if result == nil {
		t.Fatal("Invalid Result")
	}
}

func TestGetReviewsByURL(t *testing.T){

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

	result := GetReviewsByURL(server.Client(), server.URL)

	//make sure we got all rows back we expected
	assert.Equal(t, 10, len(result))
}



//Get the sample data and parse the 1st review
//Should return valid fully flushed out review
func TestParseRawReview_Values(t *testing.T) {

	doc := GetDocument()
	val := parseRawReview(doc.Find(".review-entry").First())

	assert.NotNil(t,"\"Good care &amp; Service\"", val.Title)                   
	assert.NotNil(t,"- KenWW", val.Username)
	assert.NotNil(t, "SALES VISIT - NEW",val.VisitType)
	assert.NotNil(t, "October 31, 2018",val.VisitDate)
}


//Get Review with the incorrect div. 
//Should review empty review
func TestParseRawReview_WrongDiv(t *testing.T) {

	doc := GetEmptyDocument()
	val := parseRawReview(doc.Find(".wrongdiv").First())

	//Mark sure we got a review back
	assert.NotNil(t, val)
}

//Get Review with the invalid (empty) node. 
//Should review empty review
func TestParseRawReview_InvalidNode(t *testing.T) {

	doc := GetEmptyDocument()
	val := parseRawReview(doc.Find(".notFound").First())

	//Mark sure we got a review back
	assert.NotNil(t, val)
}





//Get first review from sample to use for test. Review score is 48 Verfiy that we get the
//correct score.
func TestGetTotal(t *testing.T) {
	doc := GetDocument()
	val := getTotalScore(doc.Find(".review-entry").First())
	assert.Equal(t, 48, val)
}

//Pass the wrong div into the getTotalScore Func
//Should return 0 and not panic
func TestGetTotal_WrongDiv(t *testing.T) {
	doc := GetEmptyDocument()
	val := getTotalScore(doc.Find(".wrongdiv").First())
	assert.Equal(t, 0, val)
}

//Search for a div that doesn't exist and pass that in
//Should return 0 and not panic
func TestGetTotal_InvalidDiv(t *testing.T) {
	doc := GetEmptyDocument()
	val := getTotalScore(doc.Find(".NotFound").First())
	assert.Equal(t, 0, val)
}

func TestGetScoreFromClass(t *testing.T) {

	//check to find the numeric value 50
	val := getScoreFromClass("rating-static-indv rating-50 margin-top-none td")
	assert.Equal(t, val, 50)

}

//Send and emapty string to GetScoreFromClass
//Should return -1 as it didn't find anything
func TestGetScoreFromClass_EmptySTring(t *testing.T) {
	val := getScoreFromClass("")
	assert.Equal(t, val, -1)
}

//Send a class attr with no numerical value
//Should return -1 
func TestGetScoreFromClass_NoNumeric(t *testing.T) {
	val := getScoreFromClass("ating-static-indv margin-top-none td")

	assert.Equal(t, val, -1)
}

//Functino GetDcoument is used to get the sample page for testing
//it spins up a test http server and call the Revtrieve document method
//Server is closed after testing
func GetDocument() *goquery.Document {
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

	return retrieveDocument(server.Client(), server.URL)

}

//Functino GetDcoument is used to get the empty test page for testing
//it spins up a test http server and call the Revtrieve document method
//Server is closed after testing
func GetEmptyDocument() *goquery.Document {
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

	return retrieveDocument(server.Client(), server.URL)

}
