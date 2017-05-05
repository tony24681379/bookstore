package database

import (
	"time"

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
	db.AutoMigrate(&BookStore{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&SalesStatus{})
	db.AutoMigrate(&StoreStatus{})
}

//BundleMaster define
type BundleMaster struct {
	ID         uint `gorm:"AUTO_INCREMENT"`
	BundleName string
	Note       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

//BundleDetail define
type BundleDetail struct {
	BundleID  uint
	BookID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

//BookStore define
type BookStore struct {
	BookstoreID uint
	Location    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

//RecommedMaster define
type RecommedMaster struct {
	ID           uint `gorm:"AUTO_INCREMENT"`
	RecommedName string
	Capacity     uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

//RecommedBundle define
type RecommedBundle struct {
	ID        uint
	BundleID  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

//RecommedBook define
type RecommedBook struct {
	ID        uint
	BookID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

//Product define
type Product struct {
	ID              uint
	ProductCategory string
	ProductName     string
	Price           string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

//SalesStatus define
type SalesStatus struct {
	ProductID   uint
	BookstoreID uint
	Day         uint
	Week        uint
	Month       uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

//StoreStatus define
type StoreStatus struct {
	ProductID   uint
	BookstoreID uint
	stack       uint
	stock       uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
