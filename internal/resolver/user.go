package resolver

import (
	"DummyGameBackend/internal/db"
	"DummyGameBackend/webapi/models"
)

func (r Resolver) GetUserByUsername(userName string) *db.User {
	user := db.User{}
	r.database.Where("user = ?", userName).First(&user)
	return &user
}

func (r Resolver) CreateUser(user models.AddUser) error {
	res := r.database.Create(&db.User{
		Name:   user.Name,
		PwHash: user.Password,
	})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
