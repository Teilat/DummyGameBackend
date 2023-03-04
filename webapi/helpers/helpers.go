package helpers

import (
	"DummyGameBackend/webapi/models"
	"Messenger/database"
	"fmt"
	"gorm.io/gorm"
)

func CheckUserPass(db *gorm.DB, credentials models.Login) bool {
	user := database.User{}
	res := db.Where("username = ?", credentials.Login).First(&user)
	if res.Error != nil {
		fmt.Printf("Cant find user error:%s", res.Error.Error())
	}

	return user.PwHash == credentials.Password
}
