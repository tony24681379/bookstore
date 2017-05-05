package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	//mysql driver
	"io"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tony24681379/bookstore/database"
)

func main() {
	f, err := os.Open("/Users/tony/go/src/github.com/tony24681379/bookstore/data/PRODUCT_MST.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(bufio.NewReader(f))
	i := 0
	db := database.DBConnection()
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		i++
		if i == 1 {
			continue
		}
		//		fmt.Println(record)

		id, _ := strconv.Atoi(record[0])
		product := database.Product{
			ID:              uint(id),
			ProductCategory: record[1],
			ProductName:     record[2],
			Price:           record[3],
		}

		err = db.Create(&product).Error
		if err != nil {
			fmt.Println(err)
		}
	}

	db.Close()
}
