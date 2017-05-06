package storeCmd

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/store"
)

type storeListOptions struct {
	recommendID string
	bookstoreID string
}

//NewStoreListCmd List the book into the store
func NewStoreListCmd(db *gorm.DB) *cobra.Command {
	var opts storeListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the book which the store",
		Run: func(cmd *cobra.Command, args []string) {
			runList(db, opts)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&opts.recommendID, "recommend-id", "", "List the recommend ID")
	flags.StringVar(&opts.bookstoreID, "bookstore-id", "", "List the bookstore ID")
	return cmd
}

func runList(db *gorm.DB, opts storeListOptions) {
	store.List(db, opts.recommendID, opts.bookstoreID)
}
