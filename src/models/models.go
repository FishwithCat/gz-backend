package models

import (
	"fmt"

	"github.com/FishwithCat/gz-backend/pkg/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_ket" json:"id"`
	CreatedOn  int `json: "created_on"`
	ModifiedOn int `json: "modified_on"`
}

func SetUp() {
	dbName := "gouzi"
	var err error
	fmt.Printf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DbUser,
		setting.DbPassword,
		setting.DbAddress,
		dbName)
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DbUser,
		setting.DbPassword,
		setting.DbAddress,
		dbName))
	//db connect eg: gorm.Open("mysql","root:root@tcp(localhost:3306)/gorm")
	if err != nil {
		fmt.Println("db error")
		return
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDb() {
	defer db.Close()
}
