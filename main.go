package main

import (
	"github.com/DiegoAndresMarmota/Gorga/config"
	r "github.com/DiegoAndresMarmota/Gorga/router"
)

var logger *config.Logger

func main() {

	logger = config.GetLogger("main")
	err := config.ConfigStart()
	if err != nil {
		logger.Errorf("Error al inicializar, tipo: %v", err)
		return
	}

	r.InitRoute()
}
