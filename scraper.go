package main

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

//retrieveDocument will retreive a html page based off the client and url  passed into it. 
func retrieveDocument(client *http.Client, url string) *goquery.Document {

	res, err := client.Get(url)
	if err != nil {
		log.Error("Error Downloading Data:", err, nil)
		panic(err)
	}

	//set the body to close after we read it. 
	defer res.Body.Close()
	//check status to maure sure we got a valid response
	if res.StatusCode != 200 {
		log.Error("status code error:", "Code", res.StatusCode)
		panic(err)
	}

	// Load the HTML document from the response we received
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Error("Unable to load document", "Error", err)
		panic(err)
	}
	return doc
}

//RetrieveRawReviews is Download HTML and parse into List of Reviews
func GetReviewsByURL(client *http.Client, url string) []Review {

	//Get HTML
	doc := retrieveDocument(client, url)

	//initialize an array of reviews
	reviews := make([]Review, 0)

	// Find the review items
	doc.Find(".review-entry").Each(func(i int, s *goquery.Selection) {
		reviews = append(reviews, parseRawReview(s))
	})

	return reviews

}

func parseRawReview(rawReview *goquery.Selection) Review {
	review := Review{}

	//Fill in the standard string properties
	review.Title = rawReview.Find("h3").Text()
	review.Username = strings.TrimSpace(rawReview.Find("span.italic.font-18.black.notranslate").Text())
	review.DealerResponse = rawReview.Find(".review-response").Text()
	review.Body = rawReview.Find(".review-content").Text()
	review.VisitDate = strings.TrimSpace(rawReview.Find(".review-date").Children().First().Text())
	review.VisitType = strings.TrimSpace(rawReview.Find(".dealership-rating").Last().Text())

	//Calculate the number of emplyees this customer worked with. 
	review.EmployeesWorkedWith = rawReview.Find(".review-employee").Length()

	//Get the Scorces from the review
	review.TotalScore = getTotalScore(rawReview)
	

	return review
}

//Get the overall Score for the review
func getTotalScore(rawReview *goquery.Selection) int {
	classTag, exists := rawReview.Find(".dealership-rating").Children().First().Attr("class")
	if exists {
		return getScoreFromClass(classTag)
	}
	return 0

}


//getScoreFromClass takes a class attribute from a tag and returns the score by search for a 2 digit value in the name of one of the classes
//tagClass :
//returns: Score (int)
func getScoreFromClass(tagClass string) int {
	re := regexp.MustCompile("\\d\\d")
	score, err := strconv.Atoi(re.FindString(tagClass))
	if err != nil {
		score = -1 //set as negtive one as we did not get a valid score
	}
	return score
}
