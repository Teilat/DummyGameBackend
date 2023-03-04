package models

type User struct {
	Name       string
	Characters []Character
}

type AddUser struct {
	Name     string
	Password string
}
