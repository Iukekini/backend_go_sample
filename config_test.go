package main


import (
	"testing"
	"github.com/stretchr/testify/assert"
)


//Test to make sure that we can load the config file
func TestOpenConfig(t *testing.T) {

	settings := openConfig()

	//quick check to see if we got the correct var loaded into the config. 
	assert.Equal(t, 5, settings.PagesToScrape)
}