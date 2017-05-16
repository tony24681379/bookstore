package productCmd

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/product"
)

type productStockOptions struct {
	bookstoreID string
	bookID      string
}

//NewProductStockCmd list the stock
func NewProductStockCmd(db *gorm.DB) *cobra.Command {
	var opts productStockOptions

	cmd := &cobra.Command{
		Use:   "stock",
		Short: "Stock the book which the product",
		Run: func(cmd *cobra.Command, args []string) {
			runStock(db, opts)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&opts.bookstoreID, "bookstore-id", "", "list the bookstore product")
	flags.StringVar(&opts.bookID, "book-id", "", "list the book")
	return cmd
}

func runStock(db *gorm.DB, opts productStockOptions) {
	product.Stock(db, opts.bookstoreID, opts.bookID)
}
