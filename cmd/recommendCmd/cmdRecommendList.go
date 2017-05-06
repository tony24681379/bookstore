package recommendCmd

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/recommend"
)

type recommendListOptions struct {
	recommendID string
	bookstoreID string
}

//NewRecommendListCmd List the book into the Recommend
func NewRecommendListCmd(db *gorm.DB) *cobra.Command {
	var opts recommendListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the book which the Recommend",
		Run: func(cmd *cobra.Command, args []string) {
			runList(db, opts)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&opts.recommendID, "recommend-id", "", "List the recommend ID")
	flags.StringVar(&opts.bookstoreID, "bookstore-id", "", "List the bookstore ID")
	return cmd
}

func runList(db *gorm.DB, opts recommendListOptions) {
	recommend.List(db, opts.recommendID, opts.bookstoreID)
}
