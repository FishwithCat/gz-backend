package models

type Auth struct {
	ID       int `gorm:"primary_key"`
	Username string
	Password string
}

func CheckAuth(username, password string) (bool, error) {
	if username == "" || password == "" {
		return false, nil
	}
	query := db.Where("username = ? AND password = ?", username, password).First(&Auth{})
	// grom will save data to the table which of 'auth'
	return !query.RecordNotFound(), query.Error
}

func AddAuth(username, password string) error {
	return db.Create(&Auth{
		Username: username,
		Password: password,
	}).Error
}
