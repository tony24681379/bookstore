package recommendCmd

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/recommend"
)

type recommendAddOptions struct {
	recommendID uint
	bundleID    uint
	bookID      uint
}

//NewRecommendAddCmd Add the book into the recommend
func NewRecommendAddCmd(db *gorm.DB) *cobra.Command {
	var opts recommendAddOptions
	cmd := &cobra.Command{
		Use:   "add recommendID",
		Short: "Add the recommended area, if bookstoreID = 0, it's set by head office",
		Run: func(cmd *cobra.Command, args []string) {
			if r, err := strconv.Atoi(args[0]); err == nil {
				opts.recommendID = uint(r)
				runAdd(db, opts)
			}
		},
	}
	flags := cmd.Flags()
	flags.UintVar(&opts.bundleID, "bundle-id", 0, "Add a bundle into recommend area")
	flags.UintVar(&opts.bookID, "book-id", 0, "Add a book into recommend area")

	return cmd
}

func runAdd(db *gorm.DB, opts recommendAddOptions) {
	recommend.Add(db, opts.recommendID, opts.bundleID, opts.bookID)
}
