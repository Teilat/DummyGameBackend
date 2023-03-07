package db

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type DbProvider struct {
	connString string
	log        *log.Logger
	retryCount int
}

func NewDbProvider(retryCount int) *DbProvider {
	return &DbProvider{
		retryCount: retryCount,
		connString: fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			viper.Get("postgresql.user"),
			viper.Get("postgresql.pass"),
			viper.Get("postgresql.host"),
			viper.Get("postgresql.port"),
			viper.Get("postgresql.db")),
		log: log.New(os.Stderr, "[Database] ", log.LstdFlags),
	}
}

func (db *DbProvider) StartDatabase() (*gorm.DB, error) {
	// connecting
	db.log.Println("Connecting")
	var database *gorm.DB
	var err error
	database, err = gorm.Open(postgres.Open(db.connString), &gorm.Config{})

	for i := 0; err != nil || i > db.retryCount; i++ {
		db.log.Printf("Error while connecting error:%#v \nRetry №%d in %d seconds.\n", err.Error(), i, i*2)
		err = nil

		time.Sleep(time.Second * time.Duration(i*2))
		database, err = gorm.Open(postgres.Open(db.connString), &gorm.Config{})
	}

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
