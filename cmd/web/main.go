package main

import (
	"fmt"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/bootstrap"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/config"
	"github.com/sirupsen/logrus"
)

func main() {
	configuration, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("Error loading config: %v", err)
	}

	cfg, err := config.ParseConfig(configuration)
	if err != nil {
		logrus.Fatalf("Error parsing config: %v", err)
	}

	log := bootstrap.NewLogger(cfg)
	db := bootstrap.NewDatabase(cfg, log)
	validate := bootstrap.NewValidator(cfg)
	app := bootstrap.NewFiber(cfg)

	// initial socket io
	socket := bootstrap.NewSocketIOServer(cfg)

	bootstrap.Bootstrap(&bootstrap.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   cfg,
		Socket:   socket,
	})

	webPort := cfg.Server.Port
	err = app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
