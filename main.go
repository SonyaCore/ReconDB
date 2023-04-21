package main

import (
	"ReconDB/config"
	"ReconDB/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
	"strings"
)

// Program Info
var (
	version  = "0.1"
	build    = "Custom"
	codename = "ReconDB , ReconDB Service."
)

func Version() string {
	return version
}

// VersionStatement returns a list of strings representing the full version info.
func VersionStatement() string {
	return strings.Join([]string{
		// project name
		"ReconDB ", Version(), " (", codename, ") ", build,
		// go runtime
		" (", runtime.Version(), " ", runtime.GOOS, "/", runtime.GOARCH, ")",
		// gin version
		" ", "Gin", " ", gin.Version,
	}, "")
}

func main() {
	fmt.Println(VersionStatement())
	router := gin.New()

	config.GinInit(router)
	routers.RegisterRouter(router)

	router.Run(":8080")
}
