package main

import (
	"github.com/sedessapi/banking/app"
	"github.com/sedessapi/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
