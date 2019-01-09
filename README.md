# Backend Coding Assessment


## How to Run the Solution

1. Install Go
	you can find the instructions [here](https://golang.org/doc/install)

	*Install GCC, needed for the testing*

	*Make sure to set your GOROOT Directory [Instructions Here](https://github.com/golang/go/wiki/SettingGOPATH)*

2.	Install Dep
	You can find Instructions [here](https://github.com/golang/dep).

	*Dep is used for the package management in this application. 

3. Download code into the correct directory. 

	Go is a little picky about where code is. It wants to be in your go root directory and the code needs to be in the following path for this project. 
	
	`$GOROOT/src/github.com/Iukekini/backend-coding-assessment-Iukekini-1052`

4. Build the Project and test project.

	I setup a make file to do all the dependency loading / building / testing. 
	
	`make` 

	If you want to run it manually, here are the commands it will run. 

		dep ensure
		go build -o podium-backend-assessment -v
		go test -v ./...

6. Run the application. 
	
	`./podium-backend-assessment`

	*Results Notes*
	The results are laid out in a table wiht the follow columns
	* Probability - The is the probability that the classifier put the review in the right class (1-5). 
	* Rating - This is the class returned by the classifier
	* User - User that authored the review
	* Visit Type - Service / Sales / Used
	* Score - this is the score the user gave the review
	* Date - Date of Visit
	* Review - This is the title of the review. I didn't include the body as it was too long to display nicely. 


If you want to see more reviews or pull more data (parse more pages) you can adjust that from the config.json file. 

## How I determined "Overly positive" reviews

In order to rank the reviews based on their positivity. I setup a Bayes classifier. I used a set of amazon reviews to train the classifier on what a positive review looked like. The classifier has 5 classes based on the 5 stars of an amazon review. After the classifier was trained I checked each of the reviews that I had parsed from the site against the classifier. I took the result and used that to sort the reviews and pick the highest rated 3 reviews to show. 

*notes*

The classifier training data was not perfect for this scenario. Since an amazon review is more love / neutral / hate type of review. The classifier had a harder time picking between a good review and an over the top review. This problem could be solved by creating a set of training data that better represented this problem. 

### Problems / Questions / Frustrations 

Please feel free to give me a call 808-896-8715 or email me at justindaw@gmail.com

### Open Source References

[goml](https://github.com/cdipaolo/goml/tree/master/text) for the classification algorithm

[go-config](github.com/micro/go-config) for the Config loading and management

[Testify](https://github.com/stretchr/testify) Some add ons for the go test suite. Enables assert and panic checks. 

[Log15](gopkg.in/inconshreveable/log15.v2) for the Logging 

[goquery](github.com/PuerkitoBio/goquery) like jquery but for go. Used it for searching parsing the webpages. 

[Training Data](https://drive.google.com/drive/folders/0Bz8a_Dbh9Qhbfll6bVpmNUtUcFdjYmF2SEpmZUZUcVNiMUw1TWN6RDV3a0JHT3kxLVhVR2M) I used the amazon review csv to train the classifier. I only used the first 4k rows

[Dep](https://github.com/golang/dep) for package management