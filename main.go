package main

import (
	"ReconDB/config"
	"ReconDB/database"
	"ReconDB/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"runtime"
	"strings"
)

// Program Info
var (
	version  = "0.7.8"
	build    = "Custom"
	codename = "ReconDB , ReconDB Service."
)

func Version() string {
	return version
}

// VersionStatement returns a list of strings representing the full version info.
func VersionStatement() string {
	return strings.Join([]string{
		// Project name
		"ReconDB ", Version(), " (", codename, ") ", build,
		// Go runtime
		" (", runtime.Version(), " ", runtime.GOOS, "/", runtime.GOARCH, ")",
		// Gin version
		" ", "Gin", " ", gin.Version,
	}, "")
}

func main() {
	fmt.Println(VersionStatement())

	// initial config file
	configuration, err := config.LoadConfig()
	if err != nil {
		log.Fatal("config file not found")
	}

	// register gin engine
	router := config.GinInit(configuration)

	// config gin engine & register routers
	routers.RegisterRouter(router)

	// load mongodb
	client := database.Client(configuration)
	// ping database connection
	database.Ping(client)
	fmt.Println("\u001B[92mConnected to MongoDB", configuration.MongoURI, "\u001B[0m")

	// run gin
	router.Run(configuration.PORT)
}
