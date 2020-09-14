package main

import (
	"os"

	appInit "github.com/wdhafin/eFishery-test/init"
)

func init() {
	// Start pre-requisite app dependencies
	appInit.StartAppInit()
}

func main() {
	// Get PG Conn Instance
	pgDb, err := appInit.ConnectToPGServer()
	if err != nil {
		os.Exit(1)
	}
}
