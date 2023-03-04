package db

type User struct {
	Id         uint32 `gorm:"primaryKey"`
	Name       string
	PwHash     string
	Characters []Character `gorm:"foreignKey:User"` // one to many
}
type Character struct {
	Id        uint32 `gorm:"primaryKey"`
	Name      string
	MaxHealth float32
	Damage    float32
	Ability   string
	User      string // one to many user
}
