package webapi

import (
	"github.com/gin-contrib/cors"
	"time"
)

func newCors() cors.Config {
	return cors.Config{
		ExposeHeaders:    []string{"Access-Token", "Expire-Token"},
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PATCH", "DELETE"},
		AllowHeaders:     []string{"jwt", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Access-Control-Allow-Credentials", "Authorization", "Origin", "Accept", "X-Requested-With", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
		AllowWebSockets:  true,
	}
}
