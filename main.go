package main

import (
	"fmt"
	"log"

	"github.com/hiphopskynew/mock-api-service/config"
	"github.com/hiphopskynew/mock-api-service/repo"
	"github.com/hiphopskynew/mock-api-service/route"
	"github.com/labstack/echo"
)

func init() {
	config.LoadConfiguration()
	repo.DBSetting = repo.NewMongoDatabase(config.Config.URI, config.Config.DatabaseName, config.Config.CollectionName)
}

func main() {
	e := echo.New()
	route.RegisterRoutes(e)
	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Config.ServicePort)))
}
