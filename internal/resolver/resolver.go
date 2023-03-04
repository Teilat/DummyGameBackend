package resolver

import (
	"gorm.io/gorm"
	"log"
	"os"
)

type Resolver struct {
	log      *log.Logger
	database *gorm.DB
}

func NewResolver(database *gorm.DB) *Resolver {
	return &Resolver{
		log:      log.New(os.Stderr, "resolver", log.LstdFlags),
		database: database,
	}
}
