package store

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/lytics/logrus"
	"github.com/tony24681379/bookstore/tools"
)

//List list the store status
func List(db *gorm.DB, recommendID string, bookstoreID string) {
	sql := `select distinct rm.id
        ,recommend_name
        ,rm.bookstore_id
        ,p.id
        ,p.product_category
        ,p.product_name
        ,capacity
        ,stack_low
        ,stack_high
        ,stock_low
        ,stock_high
        ,stack
        ,stock
        from recommend_masters rm
        left join recommend_bundles rbundle on rm.id = rbundle.recommend_id
        left join recommend_books rbook on rm.id = rbook.recommend_id
        left join bundle_masters bm on bm.id = rbundle.bundle_id
        left join bundle_details bd on bm.id = bd.bundle_id
        left join products p on bd.book_id = p.id or rbook.book_id = p.id
        left join store_statuses s on s.product_id = p.id and (s.bookstore_id = rm.bookstore_id or rm.bookstore_id = 0)
        where 1=1
	`
	if recommendID != "" {
		sql = sql + " and rm.id = " + recommendID
	}
	if bookstoreID != "" {
		sql = sql + " and bookstore_id = " + bookstoreID
	}
	sql = sql + " order by rm.id, p.id"
	logrus.Debug(sql)

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
				StackLow        string
				StackHigh       string
				StockLow        string
				StockHigh       string
				Stack           string
				Stock           string
			)
			rows.Scan(&RecommendID,
				&RecommendName,
				&BookstoreID,
				&ProductID,
				&ProductCategory,
				&ProductName,
				&Capacity,
				&StackLow,
				&StackHigh,
				&StockLow,
				&StockHigh,
				&Stack,
				&Stock,
			)
			if id != RecommendID {
				id = RecommendID
				fmt.Println("\n---------------------------------------------------------------------")
				fmt.Printf("Recommend ID: %-5s Recommend Name: %-15s Bookstore: %-10s Capacity: %3s\n", RecommendID, RecommendName, BookstoreID, Capacity)
				fmt.Printf("Product ID | Product Category | %-40s | [%2s,%2s] | [%2s,%3s]\n", "ProductName", StackLow, StackHigh, StockLow, StockHigh)
			}
			if ProductID != "" {
				fmt.Printf("%10s | %-16s | %-20s | %7s | %7s \n", ProductID, tools.TruncateString(ProductCategory, 16), tools.TruncateString(ProductName, 40),
					tools.Highlight(Stack, StackLow, StackHigh, 7), tools.Highlight(Stock, StockLow, StockHigh, 7))
			}
		}
	} else {
		fmt.Println(err)
	}
}
