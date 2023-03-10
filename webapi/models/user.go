package models

type User struct {
	Login      string
	Characters []Character
}

type AddUser struct {
	Login    string
	Password string
}
