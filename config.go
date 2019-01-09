package main

import (
	//config
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

//Config Settings
type Config struct {
	URLToScrape                  string  `json:"URLToScrape"`
	PagesToScrape                int     `json:"PagesToScrape"`
	ScoreMultipler               float32 `json:"ScoreMultipler"`
	ReviewLengthMultipler        float32 `json:"ReviewLengthMultipler"`
	RealNameValue                float32 `json:"RealNameValue"`
	EmployeesWorkedWithMultipler float32 `json:"EmployeesWorkedWithMultipler"`
	ReturnNumber                 int     `json:"ReturnNumber"`
}

//Open the config file. Store at config.json
//In a more production enviroment we could choose based off some env vars.
func openConfig() Config {

	conf := config.NewConfig()
	// Load json config file
	conf.Load(file.NewSource(file.WithPath("config.json")))

	//Create a new Config instanace
	settings := Config{}

	//Load settings into struct
	conf.Scan(&settings)
	return settings

}
