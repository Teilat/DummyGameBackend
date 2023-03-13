package converters

import (
	"DummyGameBackend/internal/db"
	"DummyGameBackend/webapi/models"
)

func UserToApiUser(user *db.User, token, expire string) *models.LoginResponse {
	return &models.LoginResponse{
		Login:       user.Name,
		AccessToken: token,
		ExpireToken: expire,
	}
}
