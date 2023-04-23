package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"time"
)

var GinModes = []string{
	"debug", "release",
}

// GinInit create an instance of gin.Engine and set the proper mode with gin_mode value in configuration file
func GinInit() *gin.Engine {
	var ginMode = viper.GetString("gin_mode")
	if ginMode == "" || len(ginMode) <= 3 {
		log.Fatalf("\u001B[91mgin_mode is empty.\u001B[0m avaliable modes : %v", GinModes)
	}

	switch ginMode {
	case GinModes[0]:
		gin.SetMode(gin.DebugMode)
	case GinModes[1]:
		gin.SetMode(gin.ReleaseMode)
	default:
		log.Fatal("No gin_mode set in config.json")
	}

	router := gin.New()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	// use recovery to recover panics
	router.Use(gin.Recovery())

	return router
}
