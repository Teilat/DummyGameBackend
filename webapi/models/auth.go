package models

type Login struct {
	Login    string
	Password string
}

type LoginResponse struct {
	Login       string
	AccessToken string
	ExpireToken string
}

type Register struct {
	Login    string
	Password string
}
