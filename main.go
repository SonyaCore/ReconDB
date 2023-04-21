package main

import (
	"ReconDB/config"
	"ReconDB/database"
	"ReconDB/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"runtime"
	"strings"
)

// Program Info
var (
	version  = "0.5.0"
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
	config.LoadConfig(".")
	PORT := viper.GetString("port")

	// register gin engine
	router := config.GinInit()

	// config gin engine & register routers
	routers.RegisterRouter(router)

	// load mongodb
	client := database.Client()
	// ping database connection
	database.Ping(client)
	fmt.Println("\u001B[92mConnected to MongoDB", viper.GetString("mongo_uri"), "\u001B[0m")

	// run gin
	router.Run(PORT)
}
