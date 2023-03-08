package resolver

import (
	"DummyGameBackend/internal/db"
	"DummyGameBackend/webapi/models"
)

func (r *Resolver) GetAllCharacters() []*db.Character {
	res := make([]*db.Character, 0)
	return res
}

func (r *Resolver) GetCharacter(id uint32) *db.Character {
	character := db.Character{}
	r.database.Where("id = ?", id).First(&character)
	return &character
}

func (r *Resolver) UpdateCharacter(character *db.Character) error {
	res := r.database.Save(character)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *Resolver) DeleteCharacter(id uint32) error {
	res := r.database.Delete(&db.Character{}, id)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *Resolver) CreateCharacter(character *models.Character) error {
	res := r.database.Create(&db.Character{
		Name:      character.Name,
		MaxHealth: character.MaxHealth,
		Damage:    character.Damage,
		Ability:   character.Ability,
		User:      character.User,
	})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
