package bootstrap

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/config"
	"github.com/sirupsen/logrus"
)

func NewLogger(config *config.Config) *logrus.Logger {
	log := logrus.New()

	log.SetLevel(logrus.Level(config.Log.Level))
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}
