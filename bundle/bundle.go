package bundle

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/tony24681379/bookstore/database"
)

//Add book into the bundle
func Add(db *gorm.DB, bundleID string, bookID int) {
	fmt.Println("add", bundleID, bookID)
	if id, err := strconv.Atoi(bundleID); err == nil {
		var bundleMaster database.BundleMaster

		if db.Where(database.BundleMaster{ID: uint(id)}).First(&bundleMaster).RecordNotFound() {
			fmt.Println("bundle ID does not exsit")
		} else {
			bundleDetail := database.BundleDetail{BundleID: uint(id), BookID: bookID}
			err := db.Create(&bundleDetail).Error
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

//Create the bundle
func Create(db *gorm.DB, bundleName string, note string) {
	fmt.Println("create", bundleName, note)
	bundle := database.BundleMaster{BundleName: bundleName, Note: note}
	err := db.Create(&bundle).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Bundle ID:", bundle.ID)
}

//Delete the bundle
func Delete(db *gorm.DB, bundleID string) {
	fmt.Println("delete", bundleID)
	if id, err := strconv.Atoi(bundleID); err == nil {
		bundle := database.BundleMaster{ID: uint(id)}

		err := db.Unscoped().Delete(&bundle).Error
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

//Update the bundle
func Update(db *gorm.DB, bundleID string, bundleName string, note string) {
	fmt.Println("update", bundleID, note)
	if id, err := strconv.Atoi(bundleID); err == nil {
		var bundle database.BundleMaster
		var err error
		if bundleName != "" && note != "" {
			err = db.Model(&bundle).Where(database.BundleMaster{ID: uint(id)}).
				Update(database.BundleMaster{BundleName: bundleName, Note: note}).Error
		} else if bundleName != "" {
			err = db.Model(&bundle).Where(database.BundleMaster{ID: uint(id)}).
				Update(database.BundleMaster{BundleName: bundleName}).Error

		} else if note != "" {
			err = db.Model(&bundle).Where(database.BundleMaster{ID: uint(id)}).
				Update(database.BundleMaster{Note: note}).Error
		}
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

//Remove the book from the bundle
func Remove(db *gorm.DB, bundleID string, bookID int) {
	fmt.Println("Remove", bundleID, bookID)
	if id, err := strconv.Atoi(bundleID); err == nil {
		bundleDetail := database.BundleDetail{BundleID: uint(id), BookID: bookID}
		err := db.Unscoped().Where(bundleDetail).Delete(&bundleDetail).Error
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("success")
		}
	} else {
		fmt.Println(err)
	}
}

//List the book or bundle
func List(db *gorm.DB, bookName string, bundleName string) {
	fmt.Println("Remove", bookName, bundleName)
	sql := `select m.id as ID
		,m.bundle_name as BundleName 
		,m.note as Note
		,p.id as ProductID 
		,p.product_category as ProductCategory
		,p.product_name as ProductName
		from bundle_masters m
		left join bundle_details d on m.id = d.bundle_id
		left join products p on d.book_id = p.id
		where 1=1`
	if bookName != "" {
		sql = sql + " and product_name like '%" + bookName + "%'"
	}
	if bundleName != "" {
		sql = sql + " and bundle_name like '%" + bundleName + "%'"
	}
	sql = sql + " order by m.id, p.id"
	//	fmt.Println(sql)

	rows, err := db.Raw(sql).Rows()
	defer rows.Close()

	if err == nil {
		var id string
		for rows.Next() {
			var (
				ID              string
				BundleName      string
				Note            string
				ProductID       string
				ProductCategory string
				ProductName     string
			)
			rows.Scan(&ID, &BundleName, &Note, &ProductID, &ProductCategory, &ProductName)
			if id != ID {
				id = ID
				fmt.Println("\n---------------------------------------------------------------------")
				fmt.Printf("Bundle ID: %-6s Bundle Name: %-20s Note: %s\n", ID, BundleName, Note)
				fmt.Printf("Product ID | Product Category |  ProductName\n")
			} else {
				fmt.Printf("%10s | %-16s |  %-40s\n", ProductID, ProductCategory, ProductName)
			}
		}
	} else {
		fmt.Println(err)
	}
}
