package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/joho/godotenv"
)
func LoadConfig(configFilePath string) (ApplicationConfig){
	
	file, err := ioutil.ReadFile(configFilePath)
	if err != nil{
		fmt.Println(err);
	}
	var config ApplicationConfig
	err = json.Unmarshal(file, &config)
	if err!= nil{
		fmt.Println(err)
	}

	return config;
}

func LoadEnv(envPath string){
	// load .env file
	err := godotenv.Load(envPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	
}