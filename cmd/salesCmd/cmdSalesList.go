package salesCmd

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/sales"
)

type salesListOptions struct {
	recommendID string
	bookstoreID string
}

//NewSalesListCmd List the book into the sales
func NewSalesListCmd(db *gorm.DB) *cobra.Command {
	var opts salesListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the book which the sales",
		Run: func(cmd *cobra.Command, args []string) {
			runList(db, opts)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&opts.recommendID, "recommend-id", "", "List the recommend ID")
	flags.StringVar(&opts.bookstoreID, "bookstore-id", "", "List the bookstore ID")
	return cmd
}

func runList(db *gorm.DB, opts salesListOptions) {
	sales.List(db, opts.recommendID, opts.bookstoreID)
}
