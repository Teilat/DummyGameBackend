package db

type User struct {
	Id         uint32 `gorm:"primaryKey"`
	Name       string `gorm:"unique"`
	PwHash     string
	Characters []Character `gorm:"foreignKey:Owner;references:Name"` // one to many
}
type Character struct {
	Id        uint32 `gorm:"primaryKey"`
	Name      string
	MaxHealth float32
	Damage    float32
	Ability   string
	Owner     string // one to many user
}
