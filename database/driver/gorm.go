package driver

import (
	"context"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/fx"
	"skeleton-code/config"
	"skeleton-code/logger"
	"sync"
	"time"
)

const (
	DefaultMaxOpenConns = 25
	DefaultMaxIdleConns = 25
)

type MariaGormDB struct {
	Db   *gorm.DB
	once sync.Once
}

func NewMariaDB(lifecycle fx.Lifecycle, cnf *config.Config) (*MariaGormDB, error) {
	logger.Info("database connecting..", cnf.Database.Driver, cnf.Database.Source)
	db, err := gorm.Open(cnf.Database.Driver, cnf.Database.Source)
	//db.SetLogger(logger.Root())
	if err != nil {
		return nil, err
	}
	if cnf.Develop {
		db = db.Debug()
	}

	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1").Debug()

	if cnf.Database.MaxOpen <= 0 {
		cnf.Database.MaxOpen = DefaultMaxOpenConns
	}
	if cnf.Database.MaxIdle <= 0 {
		cnf.Database.MaxIdle = DefaultMaxIdleConns
	}
	db.DB().SetMaxOpenConns(cnf.Database.MaxOpen)
	db.DB().SetMaxIdleConns(cnf.Database.MaxIdle)
	db.DB().SetConnMaxLifetime(5 * time.Minute)

	database := &MariaGormDB{
		Db: db,
	}
	lifecycle.Append(
		fx.Hook{
			OnStop: func(ctx context.Context) error {
				database.Close()
				return nil
			},
		},
	)
	return database, nil
}

func (d *MariaGormDB) Get(model interface{}, where ...interface{}) (interface{}, error) {
	result := d.Db.Where(where).First(model)
	return model, result.Error
}

func (d *MariaGormDB) Set(model interface{}, where ...interface{}) error {
	return d.Db.Create(model).Error
}

func (d *MariaGormDB) Update(model interface{}, where ...interface{}) error {
	return d.Db.Model(model).Where(where).Updates(model).Error
}

func (d *MariaGormDB) Delete(model interface{}, where ...interface{}) error {
	return d.Db.Model(model).Where(where).Delete(model).Error
}

func (d *MariaGormDB) Count(model interface{}, where ...interface{}) (int, error) {
	return 0, nil
}

func (db *MariaGormDB) Close() error {
	err := db.Db.Close()
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
