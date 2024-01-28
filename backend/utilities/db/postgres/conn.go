package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Opts struct {
	Dsn         string
	EnablePool  bool
	MaxIdlConn  int
	MaxOpenConn int
	MaxLifetime time.Duration
}

func New(opts Opts) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(opts.Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if !opts.EnablePool {
		return db, nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(opts.MaxIdlConn)
	sqlDB.SetMaxOpenConns(opts.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(opts.MaxLifetime)

	return db, nil
}
