package converters

import (
	"DummyGameBackend/internal/db"
	"DummyGameBackend/webapi/models"
)

func UserToApiUser(user *db.User) *models.User {
	return &models.User{
		Login: user.Name,
	}
}
