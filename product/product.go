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
	logrus.Info("add", id, bookName)
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
func Mark(db *gorm.DB, id string, good bool, long bool, risk bool, Due time.Time) {
	var product database.Product
	// err := db.Where("id = ?", id).Find(&product).Error
	// if err != nil {
	// 	logrus.Error(err)
	// 	return
	// }
	i, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error(err)
		return
	}
	product.ID = uint(i)
	if good {
		product.Good = true
	}
	if long {
		product.LongTerm = true
	}
	if risk {
		product.Risk = true
	}
	err = db.Model(&product).Update(product).Error
	if err != nil {
		logrus.Error(err)
		return
	}
}
