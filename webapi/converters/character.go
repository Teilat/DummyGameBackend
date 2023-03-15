package converters

import (
	"DummyGameBackend/internal/db"
	"DummyGameBackend/webapi/models"
)

func CharactersToApiCharacters(characters []*db.Character) models.Characters {
	res := models.Characters{}
	for _, c := range characters {
		res.Characters = append(res.Characters, &models.Character{
			Id:        c.Id,
			Name:      c.Name,
			MaxHealth: c.MaxHealth,
			Damage:    c.Damage,
			Ability:   c.Ability,
		})
	}
	return res
}

func CharacterToApiCharacter(character *db.Character) *models.Character {
	return &models.Character{
		Id:        character.Id,
		Name:      character.Name,
		MaxHealth: character.MaxHealth,
		Damage:    character.Damage,
		Ability:   character.Ability,
	}
}
