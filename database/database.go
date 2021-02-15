package database

import (
	"go.uber.org/fx"
	"skeleton-code/config"
	"skeleton-code/database/driver"
	logger "skeleton-code/logger"
)

type Database interface {
	Get(model interface{}) error
	Set(model interface{}) error
	Update(model interface{}) error
	Delete(model interface{}) error
	Close() error
}

func NewDatabase(lifecycle fx.Lifecycle, config *config.Config) Database {
	if config.Database.Driver == "mysql" {
		db, err := driver.LoadDatabase(lifecycle, config)
		if err != nil {
			logger.Error(err)
		}
		return db
	} else {
		return driver.NewMemDB()
	}
}
