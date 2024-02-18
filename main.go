package main

import (

	"github.com/jeannsf/Gopportunities/config"
	"github.com/jeannsf/Gopportunities/router"
)

var (
	logger *config.Logger
)


func main() {
	logger = config.GetLogger("main")
	// Initialize configs
	err := config.Init()
	if err != nil {
	   logger.ErrorF("config initialization error: %v", err)
	   return 
	}


		// Initialize Router

		router.Initialize()
} 