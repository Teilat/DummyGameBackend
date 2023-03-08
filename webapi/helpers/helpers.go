package helpers

import (
	"DummyGameBackend/internal/db"
	"DummyGameBackend/webapi/models"
	"fmt"
	"gorm.io/gorm"
)

func CheckUserPass(database *gorm.DB, credentials models.Login) bool {
	user := db.User{}
	res := database.Where("name = ?", credentials.Login).First(&user)
	if res.Error != nil {
		fmt.Printf("Cant find user error:%s\n", res.Error.Error())
		return false
	}

	return user.PwHash == credentials.Password
}
