package db_connect

import "github.com/jinzhu/gorm"

func GormConnect() *gorm.DB {

	db, err := gorm.Open("postgres", "host=localhost user=test_user dbname=test001 sslmode=disable password=123456")
	if err != nil {
		panic(err.Error())
	}
	return db
}
