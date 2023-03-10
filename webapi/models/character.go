package models

type Character struct {
	Id        uint32
	Name      string
	MaxHealth float32
	Damage    float32
	Ability   string
	User      string
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
