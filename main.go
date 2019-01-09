package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"

	"gopkg.in/inconshreveable/log15.v2"
)

//Logging Setup
var log = log15.New()

func main() {

	log.Info("Program starting", "args", os.Args)

	//spin up our classifier
	classifierModel := GetGomlClassifier()

	//Config Settings
	settings := openConfig()

	//Retrieve Reviews from web
	reviews := GetReviews(&http.Client{}, settings.URLToScrape, settings.PagesToScrape)

	//Cacluate the Probabilty and rating
	for i := 0; i < len(reviews); i++ {
		predicted := classifierModel.Predict(reviews[i].Title)
		reviews[i].predictedRating = predicted

		//Generate Prob from just the title.
		_, probability := classifierModel.Probability(reviews[i].Title)
		reviews[i].probabilityOfRating = probability
	}

	//Sort by Rating then Probability
	sort.Sort(ByClassifierResults(reviews))

	//Print Header
	fmt.Printf("| %20s | %10s | %-30s | %-20s | %6s | %-20s | %-26s\n", "Probability", "Rating", "User", "Visit Type", "Score", "Date", "Review")

	//Print Results
	for i := 0; i < settings.ReturnNumber; i++ {
		review := reviews[i]
		fmt.Printf("| %20v | %10v | %-30s | %-20s | %6d | %-20s | %-26s\n", review.probabilityOfRating, review.predictedRating, review.Username, review.VisitType, review.TotalScore, review.VisitDate, review.Title)

	}

}
