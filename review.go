package main

import (
	//"regexp"

	"net/http"
	"strconv"
	"strings"
)

//Review Struct is used to hold the user review data parsed from the review site.
type Review struct {
	Title               string //Title of the Reivew
	Body                string //Main body of Review
	Username            string //User Created
	VisitType           string //Type of Visit, i.e. Service / New / Used
	VisitDate           string // Date of Vist
	DealerResponse      string // The Dealers Response if any
	TotalScore          int //Score user gave dealer
	EmployeesWorkedWith int //Number of Employees the user reviews
	predictedRating     uint8 //rating out of 0-4 on how positive a review is
	probabilityOfRating float64 //probability on the rating. 
}

//funcs to handle the sorting or the array.
type ByClassifierResults []Review

func (o ByClassifierResults) Len() int      { return len(o) }
func (o ByClassifierResults) Swap(i, j int) { o[i], o[j] = o[j], o[i] }

//This funcs allows us to sort by rating then by probability.
func (o ByClassifierResults) Less(i, j int) bool {
	if o[i].predictedRating == o[j].predictedRating {
		return o[j].probabilityOfRating < o[i].probabilityOfRating
	} else {
		return o[j].predictedRating < o[i].predictedRating
	}
}

//GetReviews calls out to the URL in the config based of the URL and requested page number
//It then parses the reviews from the given set of pages and returns an array of reviews.
func GetReviews(client *http.Client, baseUrl string, pages int) []Review {

	//Create our return boject
	reviews := make([]Review, 0)

	//Loop through al lthe pages
	for i := 1; i <= pages; i++ {
		//Generate the URL to the reviews for a given i
		url := strings.Replace(baseUrl, "{PAGENUM}", strconv.Itoa(i), 1)
		//Retreive Reviews
		pageReviews := GetReviewsByURL(client, url)
		//Append reviews to return list
		reviews = append(reviews, pageReviews...)
	}
	//return reviews
	return reviews
}
