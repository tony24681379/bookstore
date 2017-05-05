package bundleCmd

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/bundle"
)

type bundleAddOptions struct {
	bundleID string
	bookID   int
}

//NewBundleAddCmd add the book into the bundle
func NewBundleAddCmd(db *gorm.DB) *cobra.Command {
	var opts bundleAddOptions

	cmd := &cobra.Command{
		Use:   "add bundleID",
		Short: "Add the book into the bundle",
		Run: func(cmd *cobra.Command, args []string) {
			opts.bundleID = args[0]
			runAdd(db, opts)
		},
	}
	flags := cmd.Flags()
	flags.IntVar(&opts.bookID, "book-id", 0, "Book ID")
	return cmd
}

func runAdd(db *gorm.DB, opts bundleAddOptions) {
	bundle.Add(db, opts.bundleID, opts.bookID)
}
