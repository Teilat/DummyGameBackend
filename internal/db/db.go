package db

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DbProvider struct {
	connString string
	log        *log.Logger
}

func NewDbProvider() *DbProvider {
	return &DbProvider{
		connString: fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			viper.Get("postgresql.user"),
			viper.Get("postgresql.pass"),
			viper.Get("postgresql.host"),
			viper.Get("postgresql.port"),
			viper.Get("postgresql.db")),
		log: log.New(os.Stderr, "db", log.LstdFlags),
	}
}

func (db *DbProvider) StartDatabase() (*gorm.DB, error) {
	// connecting
	database, err := gorm.Open(postgres.Open(db.connString), &gorm.Config{})
	if err != nil {
		return nil, errors.New("error while connecting to database:" + err.Error())
	}
	// creating tables from code models
	err = database.AutoMigrate(&User{}, &Character{})
	if err != nil {
		return nil, errors.New("error migrating database:" + err.Error())
	}
	return database, nil
}
