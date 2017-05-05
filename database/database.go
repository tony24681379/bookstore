package database

import (
	"github.com/jinzhu/gorm"
	//mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DBConnection connect the database
func DBConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:bookstore@/bookstore?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	initDB(db)
	return db
}

//InitDB initial the database
func initDB(db *gorm.DB) {
	db.AutoMigrate(&BundleMaster{})
	db.AutoMigrate(&BundleDetail{})
}
