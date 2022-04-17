package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const configFilePath = "../config/config.json"

var configParams *config = nil

type config struct {
	Port      int    `json:"port"`
	JwtSecret string `json:"jwtSecret"`
	DbName    string `json:"dbName"`
	DbUsr     string `json:"dbUser"`
	DbPw      string `json:"dbPw"`
	DbHost    string `json:"dbHost"`
}

func GetConfig() config {
	if configParams == nil {
		jsonFile, err := os.Open(configFilePath)
		if err != nil {
			fmt.Println("Error, configuration file not found")
			fmt.Println("Loading default configuration...")
			*configParams = loadDefaultConfig()
			return *configParams
		}
		defer jsonFile.Close()

		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			log.Fatal(err.Error())
		}
		json.Unmarshal(byteValue, &configParams)
	}
	return *configParams
}

func loadDefaultConfig() config {
	return config{Port: 8080, JwtSecret: "secret", DbName: "", DbUsr: "", DbPw: "", DbHost: ""}
}
