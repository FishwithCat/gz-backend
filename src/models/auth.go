package models

type Auth struct {
	ID       int
	Username string
	Password string
}

func CheckAuth(username, password string) bool {
	return true
}
