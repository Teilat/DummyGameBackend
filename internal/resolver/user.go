package resolver

import (
	"DummyGameBackend/internal/db"
	"DummyGameBackend/webapi/models"
	"fmt"
	"strings"
)

func (r *Resolver) GetUserByUsername(userName string) *db.User {
	user := db.User{}
	r.database.Where("name = ?", userName).First(&user)
	return &user
}

func (r *Resolver) CreateUser(user models.AddUser) error {
	var count int64
	r.database.Model(&db.User{}).Where("name = ?", user.Login).Count(&count)
	if count > 0 {
		return fmt.Errorf("user with this name already exist")
	}

	if strings.TrimSpace(user.Login) == "" {
		return fmt.Errorf("user with empty name not allowed")
	}

	res := r.database.Create(&db.User{
		Name:   user.Login,
		PwHash: user.Password,
	})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
