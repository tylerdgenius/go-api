package http

import (
	"context"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Env Env
}

func NewDatabase(env Env) (*Database, error) {
	return &Database{
		Env: env,
	}, nil
}

func (d *Database) Connect(ctx context.Context) (*gorm.DB, error) {
	gdb, error := gorm.Open(postgres.Open(d.Env.DSN()), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})

	if error != nil {
		log.Panic("Error opening database connection:", error)
	}

	sqlDB, err := gdb.DB()

	if err != nil {
		log.Panic("Error getting database connection:", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	log.Println("Database connection established")

	return gdb, nil
}
