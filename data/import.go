package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	//mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	f, err := os.Open("/Users/tony/go/src/github.com/tony24681379/bookstore/data/BOOKSTORE_MST.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(bufio.NewReader(f))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	db := dbConnection()
	db.HasTable("users")
	fmt.Print(records)
}

//DB connect the database
func dbConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:bookstore@/bookstore?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	return db
}
