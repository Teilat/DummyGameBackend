package models

type Characters struct {
	Characters []*Character
}

type Character struct {
	Id        uint32
	Name      string
	MaxHealth float32
	Damage    float32
	Ability   string
}

type UpdateCharacter struct {
	Id        uint32
	Name      string
	MaxHealth float32
	Damage    float32
	Ability   string
}

type AddCharacter struct {
	Name      string
	MaxHealth float32
	Damage    float32
	Ability   string
}
