package product

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lytics/logrus"
	"github.com/tony24681379/bookstore/database"
	"github.com/tony24681379/bookstore/tools"
)

//List the product
func List(db *gorm.DB, id string, bookName string) {
	logrus.Debug("list", id, bookName)
	var products database.Products
	err := db.Find(&products).Error
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println("Product ID | Product Category | Product Name                   |  Good |  Long |  Risk | Due")
	for _, v := range products {
		fmt.Printf("%10s | %16s | %-30s |",
			strconv.Itoa(int(v.ID)),
			tools.TruncateString(v.ProductCategory, 16),
			tools.TruncateString(v.ProductName, 30))
		if v.Good {
			fmt.Printf(" %5v |", v.Good)
		} else {
			fmt.Printf(" %5v |", "")
		}
		if v.LongTerm {
			fmt.Printf(" %5v |", v.LongTerm)
		} else {
			fmt.Printf(" %5v |", "")
		}
		if v.Risk {
			fmt.Printf(" %5v | ", v.Risk)
		} else {
			fmt.Printf(" %5v | ", "")
		}
		if !v.Due.IsZero() {
			fmt.Println(v.Due)
		} else {
			fmt.Println()
		}
	}
}

//Mark the book
func Mark(db *gorm.DB, id string, good string, long string, risk string, Due time.Time) {
	var product database.Product
	i, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error(err)
		return
	}
	product.ID = uint(i)
	if good != "" {
		if good == "true" {
			product.Good = true
		} else {
			product.Good = false
		}
	}
	if long != "" {
		if long == "true" {
			product.LongTerm = true
		} else {
			product.LongTerm = false
		}
	}
	if risk != "" {
		if risk == "true" {
			product.Risk = true
		} else {
			product.Risk = false
		}
	}
	err = db.Model(&product).Update(product).Error
	if err != nil {
		logrus.Error(err)
		return
	}
}

//Stock list the books inventory
func Stock(db *gorm.DB, bookstoreID string, bookID string) {
	sql := `select b.bookstore_id
		,p.id
		,product_category
		,product_name
		,stack
		,stock
		from products p
		join store_statuses s on p.id = s.product_id
		join book_stores b on s.bookstore_id = b.bookstore_id
		where 1=1
	`
	if bookstoreID != "" {
		sql = sql + " and s.bookstore_id = " + bookstoreID
	}
	if bookID != "" {
		sql = sql + " and p.id = " + bookID
	}
	sql = sql + " order by p.id, s.bookstore_id "
	logrus.Debug(sql)

	rows, err := db.Raw(sql).Rows()
	defer rows.Close()

	if err == nil {
		id := ""
		for rows.Next() {
			var (
				BookstoreID     string
				ProductID       string
				ProductCategory string
				ProductName     string
				Stack           string
				Stock           string
			)
			rows.Scan(
				&BookstoreID,
				&ProductID,
				&ProductCategory,
				&ProductName,
				&Stack,
				&Stock,
			)
			if ProductID != "" {
				if id != ProductID {
					id = ProductID
					fmt.Println("\n---------------------------------------------------------------------")
					fmt.Printf("Bookstore ID | Product ID | Product Category | %-30s |   Stack |   Stock \n", "ProductName")
				}
				fmt.Printf("%12s | %10s | %-16s | %-30s | %7s | %7s \n", BookstoreID, ProductID,
					tools.TruncateString(ProductCategory, 16), tools.TruncateString(ProductName, 30),
					tools.Highlight(Stack, "10", "50", 7), tools.Highlight(Stock, "30", "70", 7))
			}
		}
	} else {
		logrus.Error(err)
	}
}
