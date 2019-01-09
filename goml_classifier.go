package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"

	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/text"
)

//GomlClassifier stands up a bayesin classifier suing the goml package and training data from amazon reviews.
func GetGomlClassifier() *text.NaiveBayes {

	// create the channel of data and errors
	stream := make(chan base.TextDatapoint, 100)
	errors := make(chan error)

	model := text.NewNaiveBayes(stream, 5, func(r rune) bool {
		return !(r == ' ' || r == '!' || unicode.IsLetter(r) || unicode.IsDigit(r))
	})

	go model.OnlineLearn(errors)

	trainGoMLClassifier(stream)

	close(stream)
	for {
		err, _ := <-errors
		if err != nil {
			fmt.Printf("Error passed: %v", err)
		} else {
			// training is done!
			break
		}
	}

	return model
}

//Thsi function trains the classifier with the amazon review data.
func trainGoMLClassifier(stream chan base.TextDatapoint) {

	//Open the file
	csvFile, _ := os.Open("training/training_data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	reviewCount := make([]int, 5)

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

		reviewCount[classIdx]++
		stream <- base.TextDatapoint{
			X: line[1],
			Y: uint8(classIdx),
		}
	}

}
