package sales

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	"github.com/lytics/logrus"
)

//List the sales status
func List(db *gorm.DB, recommendID string, bookstoreID string) {
	sql := `select distinct rm.id
        ,recommend_name
        ,rm.bookstore_id
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
        ,day
        ,week
        ,month
        from recommend_masters rm
        left join recommend_bundles rbundle on rm.id = rbundle.recommend_id
        left join recommend_books rbook on rm.id = rbook.recommend_id
        left join bundle_masters bm on bm.id = rbundle.bundle_id
        left join bundle_details bd on bm.id = bd.bundle_id
        left join products p on bd.book_id = p.id or rbook.book_id = p.id
        left join sales_statuses s on s.product_id = p.id and (s.bookstore_id = rm.bookstore_id or rm.bookstore_id = 0)
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
				Day             string
				Week            string
				Month           string
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
				&StockHigh,
				&Day,
				&Week,
				&Month,
			)
			if id != RecommendID {
				id = RecommendID
				fmt.Println("\n---------------------------------------------------------------------")
				fmt.Printf("Recommend ID: %-5s Recommend Name: %-15s Bookstore: %-10s Capacity: %3s\n", RecommendID, RecommendName, BookstoreID, Capacity)
				// fmt.Println("DayLow | DayHigh | WeekLow | WeekHigh | MonthLow | MonthHigh | StackLow | StackHigh | StockLow | StockHigh |")
				// fmt.Printf("%6s |%8s |%8s |%9s |%9s |%10s |%9s |%10s |%9s |%10s |\n")
				fmt.Printf("Product ID | Product Category | %-40s | [%s,%s]| [%s,%s]| [%s,%s]\n", "ProductName", DayLow, DayHigh,
					WeekLow, WeekHigh,
					MonthLow, MonthHigh)
			}
			if ProductID != "" {
				fmt.Printf("%10s | %-16s | %-20s | %s| %s| %s\n", ProductID, ProductCategory, ProductName,
					highlight(Day, DayLow, DayHigh), highlight(Week, WeekLow, WeekHigh), highlight(Month, MonthLow, MonthHigh))
			}
		}
	} else {
		fmt.Println(err)
	}
}

func highlight(value string, low string, high string) string {
	v, _ := strconv.Atoi(value)
	s, _ := strconv.Atoi(low)
	h, _ := strconv.Atoi(high)
	if v < s {
		return color.RedString(value)
	} else if v > h {
		return color.GreenString(value)
	}
	return value
}
