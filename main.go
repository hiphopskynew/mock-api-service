package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/hiphopskynew/mock-api-service/repo"
	"github.com/hiphopskynew/mock-api-service/route"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	var (
		mongoURI     = os.Getenv("MONGO_URI")
		mongoDbName  = os.Getenv("MONGO_DATABASE_NAME")
		mongoColName = os.Getenv("MONGO_URI_COLLECTION_NAME")
	)
	if len(strings.TrimSpace(mongoURI)) == 0 {
		mongoURI = "mongodb://admin:password@localhost:27017/admin"
		mongoDbName = "mock-api-service"
		mongoColName = "setting"
	}
	repo.DBSetting = repo.NewMongoDatabase(mongoURI, mongoDbName, mongoColName)
}

func main() {
	port := os.Getenv("PORT")
	if len(strings.TrimSpace(port)) == 0 {
		port = "8080"
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	route.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
