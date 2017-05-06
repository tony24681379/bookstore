package productCmd

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/product"
)

type productListOptions struct {
	id       string
	bookName string
}

//NewProductListCmd List the book into the product
func NewProductListCmd(db *gorm.DB) *cobra.Command {
	var opts productListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the book which the product",
		Run: func(cmd *cobra.Command, args []string) {
			runList(db, opts)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&opts.id, "book-id", "", "List the book id")
	flags.StringVar(&opts.bookName, "book-name", "", "List the book name")
	return cmd
}

func runList(db *gorm.DB, opts productListOptions) {
	product.List(db, opts.id, opts.bookName)
}
