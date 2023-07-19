package main

import (
	"go-grow-events/config"
	handler "go-grow-events/delivery/http"
	"go-grow-events/repository"
	"go-grow-events/router"
	"go-grow-events/usecase"
	"log"
)

func main() {
    config.GetEnvConfig()

    /*
    Config for JSON

    config.GetJSONConfig()
    g := viper.Get("server.port")
    fmt.Println(g)
    
    */

    databaseConnection, err := config.ConnectDB()
    if err != nil {
        log.Fatalf("could not initialize database connection: %s", err)
    }

    baseRepository := repository.NewBaseRepository(databaseConnection)

	eventUsecase := usecase.NewEventUsecase(baseRepository)
	eventHandler := handler.NewEventHandler(eventUsecase)

    router.InitRouter(eventHandler)
    router.Start()
}
