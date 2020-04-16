package main

import (
	"fmt"

	"github.com/hiphopskynew/mock-api-service/config"
	"github.com/hiphopskynew/mock-api-service/repo"
	"github.com/hiphopskynew/mock-api-service/route"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	config.LoadConfiguration()
	repo.DBSetting = repo.NewMongoDatabase(config.Config.URI, config.Config.DatabaseName, config.Config.CollectionName)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	route.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Config.ServicePort)))
}
