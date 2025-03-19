package database

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() {
	// dsn := config.GetConfig().GetString("database.connect")
	// _db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println("Error when create connection")
	// }
	// db = _db
}

func GetDB() *gorm.DB {
	return db
}
