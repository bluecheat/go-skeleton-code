package database

import (
	"go.uber.org/fx"
	"skeleton-code/config"
	"skeleton-code/database/driver"
	logger "skeleton-code/logger"
)

type Database interface {
	Get(model interface{}, where ...interface{}) (interface{}, error)
	Set(model interface{}, where ...interface{}) error
	Count(model interface{}, where ...interface{}) (int, error)
	Update(model interface{}, where ...interface{}) error
	Delete(model interface{}, where ...interface{}) error
	Close() error
}

func NewDatabase(lifecycle fx.Lifecycle, config *config.Config) Database {
	if config.Database.Driver == "mysql" {
		db, err := driver.NewMariaDB(lifecycle, config)
		if err != nil {
			logger.Error(err)
		}
		return db
	} else {
		return driver.NewMemDB()
	}
}
