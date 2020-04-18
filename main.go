package main

import (
	"log"
	"net/http"
	configuration "order-service/configuration"
	"order-service/infrastructure/persistence"
	interfaces "order-service/interfaces"

	"github.com/joho/godotenv"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
	conf, err := configuration.LoadConfiguration()
	if err != nil {
		panic(err)
	}

	services, err := persistence.NewRepositories()
	if err != nil {
		panic(err)
	}

	defer services.Close()
	services.Automigrate()

	handlers := interfaces.CreateHandlers(services)

	router := interfaces.NewRouter(handlers)
	//Starting the application
	app_port := conf.Api.Port
	if app_port == "" {
		app_port = "8888" //localhost
	}
	log.Fatal(http.ListenAndServe(":"+app_port, router))
}
