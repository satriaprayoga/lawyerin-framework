package db

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Driver struct {
	DbType string
	DB     *gorm.DB
}

func OpenDB(dbType string, connString string) (*gorm.DB, error) {
	var (
		conn     *gorm.DB
		err      error
		logLevel logger.LogLevel = logger.Silent
		gormCfg  *gorm.Config
	)
	if os.Getenv("DEBUG") == "true" {
		logLevel = logger.Info
	}
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logLevel,    // Log level
			Colorful:      false,       // Disable color
		},
	)

	gormCfg = &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   os.Getenv("DB_TABLE_PREFIX"),
			SingularTable: true,
		},
		Logger: dbLogger,
	}

	if dbType == "postgres" || dbType == "postgresql" {
		conn, err = gorm.Open(postgres.Open(connString), gormCfg)
		if err != nil {
			return nil, err
		}
	} else if dbType == "mysql" || dbType == "mariadb" {
		conn, err = gorm.Open(mysql.Open(connString), gormCfg)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("unsupported database type")
	}

	sqlDB, err := conn.DB()
	if err != nil {
		return nil, err

	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return conn, nil
}

func Close(db *gorm.DB) {
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDb.Close()
}

func BuildConnDB() string {
	var conn string

	switch os.Getenv("DB_TYPE") {
	case "postgres", "postgresql":
		conn = fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_SSLMODE"),
			os.Getenv("DB_TIMEZONE"))
		if os.Getenv("DB_PASSWORD") != "" {
			conn = fmt.Sprintf("%s password=%s", conn, os.Getenv("DB_PASSWORD"))
		}
	case "mysql":
		conn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"))
	default:
	}
	return conn
}
