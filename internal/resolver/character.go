package resolver

import (
	"DummyGameBackend/internal/db"
	"DummyGameBackend/webapi/models"
	"fmt"
)

func (r *Resolver) GetAllCharacters(user string) []*db.Character {
	var res []*db.Character
	resp := r.database.Find(&res, "owner = ?", user)
	if resp.Error != nil {
		fmt.Println(resp.Error)
	}

	return res
}

func (r *Resolver) GetCharacter(id string, user string) *db.Character {
	character := db.Character{}
	r.database.First(&character, "id = ? AND owner = ?", id, user)
	return &character
}

func (r *Resolver) UpdateCharacter(character *models.UpdateCharacter, user string) error {
	char := db.Character{
		Id:        character.Id,
		Name:      character.Name,
		MaxHealth: character.MaxHealth,
		Damage:    character.Damage,
		Ability:   character.Ability,
		Owner:     user,
	}

	res := r.database.Save(char)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *Resolver) DeleteCharacter(id string, user string) error {
	res := r.database.Delete(&db.Character{}, id)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *Resolver) CreateCharacter(character *models.AddCharacter, user string) error {
	res := r.database.Create(&db.Character{
		Name:      character.Name,
		MaxHealth: character.MaxHealth,
		Damage:    character.Damage,
		Ability:   character.Ability,
		Owner:     user,
	})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
