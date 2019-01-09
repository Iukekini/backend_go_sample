package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Simple test to check  and make sure we can spin up our classifier and load the training data.
func TestGetGomlClassifier(t *testing.T) {
	classifier := GetGomlClassifier()
	//Make sure we didn't error out
	assert.NotNil(t, classifier)
	//make sure we have more then one classifier
	assert.True(t, len(classifier.Count) > 0)
}

type precisionCheck struct {
	count   float32
	correct float32
	close float32
}

//Simple test to check  and make sure we can spin up our classifier and load the training data.
func TestCheckPresision(t *testing.T) {
	classifierModel := GetGomlClassifier()

	precisionChecks := make([]precisionCheck, 5)

	//Open the file
	csvFile, _ := os.Open("training/testing_data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	//Loop through each record
	for {
		line, error := reader.Read()

		//Check of End of file or error
		if error == io.EOF {
			break
		} else if error != nil {
			log.Error("Error Reading File", "Error MEssage", error)
		}

		classIdx, _ := strconv.Atoi(line[0])
		//drop 1 for indx
		classIdx--

		precisionChecks[classIdx].count++
		predicted, _ := classifierModel.Probability(line[1])


		//Boom Correct classification
		if classIdx == int(predicted){
			precisionChecks[classIdx].correct++
		}

		//since this is a linar classifier. let say that if we are with in 1 star we also got it right. there is a lot of cross over in this. 
		//so we are only really wrong if we are off by more then 2. 
		if classIdx  == int(predicted) -1 || int(predicted) +1 == classIdx{
			precisionChecks[classIdx].correct++
		}

	}

	fmt.Printf("Testing the Classifier with Test Data.\n -----------------------------------------------\n")
	for idx, result := range precisionChecks {
		percentageCorrect := float32(result.correct) / float32(result.count)
		fmt.Printf("%d star reviews: Total->%g  Predicted->%g Precision->%v \n", idx+1, result.count, result.correct, percentageCorrect)

	}
	fmt.Printf("-----------------------------------------------\n")

}
