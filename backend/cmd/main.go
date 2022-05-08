package main

import (
	"flag"
	"fmt"
	"github/Omar-Radwan/backend/config"
	
	"github/Omar-Radwan/backend/internal/routes"
	"net/http"
)

func main() {
	
	// env file path
	var envFilePath string

	flag.StringVar(
		&envFilePath,
		"env",
		"./config/res/.env-example",
		"Location of the environment file",
	)

	// config file flag
	var configFilePath string
	flag.StringVar(
		&configFilePath,
		"config",
		"./config/res/config.json",
		"Location of the config file",
	)

	flag.Parse()

	config.LoadEnv(envFilePath)
	config := config.LoadConfig(configFilePath)


	router := routes.LoadRoutes()
	router.Use(routes.Middleware)
	http.ListenAndServe(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port), router)
}
