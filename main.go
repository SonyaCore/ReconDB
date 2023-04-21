package main

import (
	"fmt"
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
		"ReconDB ", Version(), " (", codename, ") ", build, " (", runtime.Version(), " ", runtime.GOOS, "/", runtime.GOARCH, ")",
	}, "")
}

func main() {
	fmt.Println(VersionStatement())
}
