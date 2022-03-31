package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const ConfigFilePath = "../config/config.json"

var ConfigParams *Config = ReadConfig()

type Config struct {
	Port      int    `json:"port"`
	JwtSecret string `json:"jwtSecret"`
	DbName    string `json:"dbName"`
	DbUsr     string `json:"dbUser"`
	DbPw      string `json:"dbPw"`
	DbHost    string `json:"dbHost"`
}

func ReadConfig() *Config {

	jsonFile, err := os.Open(ConfigFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	var config *Config
	json.Unmarshal(byteValue, &config)

	return config
}
