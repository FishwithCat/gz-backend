package models

type Auth struct {
	ID       int `gorm:"primary_key"`
	Username string
	Password string
}

func CheckAuth(username, password string) (bool, error) {
	query := db.Where("username = ? AND password = ?", username, password).First(&Auth{})
	return !query.RecordNotFound(), query.Error
}

func AddAuth(username, password string) error {
	return db.Create(&Auth{
		Username: username,
		Password: password,
	}).Error
}
