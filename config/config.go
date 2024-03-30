package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func ConfigStart() error {
	return errors.New("Error sim")
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
