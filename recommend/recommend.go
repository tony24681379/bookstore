package recommend

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/lytics/logrus"
	"github.com/tony24681379/bookstore/database"
)

//Add bundle or book in a recommend area
func Add(db *gorm.DB, recommendID uint, bundleID uint, bookID uint) {
	fmt.Println("add", recommendID, bundleID, bookID)
	var err error
	if bundleID != 0 {
		recommend := database.RecommendBundle{
			RecommendID: recommendID,
			BundleID:    bundleID,
		}
		err = db.Create(&recommend).Error
		logrus.Infof("Recommend area: %+v", recommend)
	} else if bookID != 0 {
		recommend := database.RecommendBook{
			RecommendID: recommendID,
			BookID:      bookID,
		}
		err = db.Create(&recommend).Error
		logrus.Infof("Recommend area: %+v", recommend)
	}
	if err != nil {
		logrus.Error(err)
	}
}

//Create a recommend area in a branch
func Create(db *gorm.DB, recommendName string, bookstoreID uint, capacity uint,
	dayLow uint, dayHigh uint,
	weekLow uint, weekHigh uint,
	monthLow uint, monthHigh uint,
	stackLow uint, stackHigh uint,
	stockLow uint, stockHigh uint) {
	logrus.Debug("create", recommendName, bookstoreID)
	recommend := database.RecommendMaster{
		RecommendName: recommendName,
		BookstoreID:   bookstoreID,
		Capacity:      capacity,
		DayLow:        dayLow,
		DayHigh:       dayHigh,
		WeekLow:       weekLow,
		WeekHigh:      weekHigh,
		MonthLow:      monthLow,
		MonthHigh:     monthHigh,
		StackLow:      stackLow,
		StackHigh:     stackHigh,
		StockLow:      stockLow,
		StockHigh:     stockHigh,
	}
	err := db.Create(&recommend).Error
	if err != nil {
		logrus.Error(err)
	}
	logrus.Infof("Recommend area: %+v", recommend)
}

//Delete a recommend area in a branch
func Delete(db *gorm.DB, recommendID uint, bookstoreID uint) {
	logrus.Debug("delete", recommendID, bookstoreID)
	recommend := database.RecommendMaster{
		ID:          recommendID,
		BookstoreID: bookstoreID,
	}
	err := db.Unscoped().Delete(&recommend).Error
	if err != nil {
		logrus.Error(err)
	}
}

//Remove bundle or book in a recommend area
func Remove(db *gorm.DB, bundleID uint, bookID uint) {

}

//List the recommend areas
func List(db *gorm.DB, recommendID string, bookstoreID string) {
	sql := `
	select distinct rm.id
	,recommend_name
	,bookstore_id
	,p.id
	,p.product_category
	,p.product_name
	,capacity
	,day_low
	,day_high
	,week_low
	,week_high
	,month_low
	,month_high
	,stack_low
	,stack_high
	,stock_low
	,stock_high
	from recommend_masters rm
	left join recommend_bundles rbundle on rm.id = rbundle.recommend_id
	left join recommend_books rbook on rm.id = rbook.recommend_id
	left join bundle_masters bm on bm.id = rbundle.bundle_id
	left join bundle_details bd on bm.id = bd.bundle_id
	left join products p on bd.book_id = p.id or rbook.book_id = p.id
	where 1=1
	`
	if recommendID != "" {
		sql = sql + " and rm.id = " + recommendID
	}
	if bookstoreID != "" {
		sql = sql + " and bookstore_id = " + bookstoreID
	}
	sql = sql + " order by rm.id, p.id"
	logrus.Info(sql)

	rows, err := db.Raw(sql).Rows()
	defer rows.Close()

	if err == nil {
		id := ""
		for rows.Next() {
			var (
				RecommendID     string
				RecommendName   string
				BookstoreID     string
				ProductID       string
				ProductCategory string
				ProductName     string
				Capacity        string
				DayLow          string
				DayHigh         string
				WeekLow         string
				WeekHigh        string
				MonthLow        string
				MonthHigh       string
				StackLow        string
				StackHigh       string
				StockLow        string
				StockHigh       string
			)
			rows.Scan(&RecommendID,
				&RecommendName,
				&BookstoreID,
				&ProductID,
				&ProductCategory,
				&ProductName,
				&Capacity,
				&DayLow,
				&DayHigh,
				&WeekLow,
				&WeekHigh,
				&MonthLow,
				&MonthHigh,
				&StackLow,
				&StackHigh,
				&StockLow,
				&StockHigh)
			if id != RecommendID {
				id = RecommendID
				fmt.Println("\n---------------------------------------------------------------------")
				fmt.Printf("Recommend ID: %-5s Recommend Name: %-15s Bookstore: %-10s Capacity: %3s\n", RecommendID, RecommendName, BookstoreID, Capacity)
				fmt.Println("DayLow | DayHigh | WeekLow | WeekHigh | MonthLow | MonthHigh | StackLow | StackHigh | StockLow | StockHigh |")
				fmt.Printf("%6s |%8s |%8s |%9s |%9s |%10s |%9s |%10s |%9s |%10s |\n", DayLow, DayHigh,
					WeekLow, WeekHigh,
					MonthLow, MonthHigh,
					StackLow, StackHigh,
					StockLow, StockHigh)
				fmt.Println("Product ID | Product Category |  ProductName")
			}
			if ProductID != "" {
				fmt.Printf("%10s | %-16s |  %-40s\n", ProductID, ProductCategory, ProductName)
			}
		}
	} else {
		fmt.Println(err)
	}
}
