package resolver

import (
	"DummyGameBackend/internal/db"
	"DummyGameBackend/webapi/models"
	"fmt"
)

func (r *Resolver) GetUserByUsername(userName string) *db.User {
	user := db.User{}
	r.database.Where("name = ?", userName).First(&user)
	return &user
}

func (r *Resolver) CreateUser(user models.AddUser) error {
	var count int64
	r.database.Model(&db.User{}).Where("name = ?", user.Name).Count(&count)
	if count > 0 {
		return fmt.Errorf("user with this name already exist")
	}

	res := r.database.Create(&db.User{
		Name:   user.Name,
		PwHash: user.Password,
	})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
