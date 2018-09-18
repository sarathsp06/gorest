package core

import (
	"log"
	"time"

	config "github.com/sarathsp06/gorest/config"
	"github.com/sarathsp06/gorest/core/order"
	"github.com/sarathsp06/gorest/core/order/repository"
	"github.com/sarathsp06/gorest/db/crud"
	"github.com/sarathsp06/gorest/db/crud/mongo"

	"github.com/sarathsp06/gorest/utils/routes"
	"github.com/sarathsp06/gorest/utils/routes/google"
)

func initializeCRUD() {
	db, err := mongo.New(
		config.Config.Mongo.Host,
		config.Config.Mongo.Port,
		config.Config.ProcessName,
		time.Duration(config.Config.Mongo.ConnectionTimeout)*time.Second,
	)
	if err != nil {
		log.Fatal("Error initiaizing mongodb", err.Error())
	}

	crud.SetDefault(db)
}

func initializeRoute() {
	client, err := google.New(config.Config.Google.Key)
	if err != nil {
		log.Fatalf("Error initializing google route service ,Err: %s", err)
	}
	routes.SetDefault(client)
}

func initializeBackingServices() {
	initializeCRUD()
	initializeRoute()
}

func initializeResources() {
	order.SetExecuter(repository.New())
}

// Initialize initializes all the core implementations and
// backing services
// upon failure fail execution of process
func Initialize() {
	initializeResources()
	initializeBackingServices()
}
