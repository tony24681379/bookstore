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
	db.AutoMigrate(&RecommendMaster{})
	db.AutoMigrate(&RecommendBundle{})
	db.AutoMigrate(&RecommendBook{})
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

//RecommendMaster define
type RecommendMaster struct {
	ID            uint `gorm:"AUTO_INCREMENT"`
	RecommendName string
	BookstoreID   uint
	Capacity      uint
	DayLow        uint
	DayHigh       uint
	WeekLow       uint
	WeekHigh      uint
	MonthLow      uint
	MonthHigh     uint
	StackLow      uint
	StackHigh     uint
	StockLow      uint
	StockHigh     uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

//RecommendBundle define
type RecommendBundle struct {
	RecommendID uint
	BundleID    uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

//RecommendBook define
type RecommendBook struct {
	RecommendID uint
	BookID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

//Product define
type Product struct {
	ID              uint
	ProductCategory string
	ProductName     string
	Price           string
	Good            bool
	LongTerm        bool
	Risk            bool
	Due             time.Time
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
	Stack       uint
	Stock       uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
