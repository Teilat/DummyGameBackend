package converters

import (
	"DummyGameBackend/internal/db"
	"DummyGameBackend/webapi/models"
)

func UserToApiUser(user *db.User) *models.User {
	return &models.User{
		Name: user.Name,
	}
}
