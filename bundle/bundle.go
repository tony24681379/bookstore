package bundle

import (
	"fmt"
	"strconv"

	"github.com/tony24681379/bookstore/database"
)

//Add book into the bundle
func Add(bundleID string, bookID int) {
	fmt.Println("add", bundleID, bookID)
	if id, err := strconv.Atoi(bundleID); err == nil {
		db := database.DBConnection()
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
		db.Close()
	}
}

//Create the bundle
func Create(bundleName string, note string) {
	fmt.Println("create", bundleName, note)
	bundle := database.BundleMaster{BundleName: bundleName, Note: note}
	db := database.DBConnection()
	err := db.Create(&bundle).Error
	if err != nil {
		fmt.Println(err)
	}
	db.Close()
}

//Delete the bundle
func Delete(bundleID string) {
	fmt.Println("delete", bundleID)
	if id, err := strconv.Atoi(bundleID); err == nil {
		bundle := database.BundleMaster{ID: uint(id)}
		db := database.DBConnection()
		err := db.Unscoped().Delete(&bundle).Error
		if err != nil {
			fmt.Println(err)
		}
		db.Close()
	} else {
		fmt.Println(err)
	}
}

//Update the bundle
func Update(bundleID string, bundleName string, note string) {
	fmt.Println("update", bundleID, note)
	if id, err := strconv.Atoi(bundleID); err == nil {
		db := database.DBConnection()

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
		db.Close()
	} else {
		fmt.Println(err)
	}
}

//Remove the book from the bundle
func Remove(bundleID string, bookID int) {
	fmt.Println("Remove", bundleID, bookID)
	if id, err := strconv.Atoi(bundleID); err == nil {
		db := database.DBConnection()

		bundleDetail := database.BundleDetail{BundleID: uint(id), BookID: bookID}
		err := db.Unscoped().Where(bundleDetail).Delete(&bundleDetail).Error
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("success")
		}
		db.Close()
	} else {
		fmt.Println(err)
	}
}
