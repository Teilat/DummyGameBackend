package models

type Character struct {
	Name      string
	MaxHealth float32
	Damage    float32
	Ability   string
	User      string
}

type UpdateCharacter struct {
}

type AddCharacter struct {
}
