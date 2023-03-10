package converters

import (
	"DummyGameBackend/internal/db"
	"DummyGameBackend/webapi/models"
)

func CharactersToApiCharacters(characters []*db.Character) []*models.Character {
	res := make([]*models.Character, len(characters))
	for _, c := range characters {
		res = append(res, &models.Character{
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
