package bundleCmd

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/bundle"
)

type bundleRemoveOptions struct {
	bundleID string
	bookID   int
}

//NewBundleRemoveCmd Remove the book into the bundle
func NewBundleRemoveCmd(db *gorm.DB) *cobra.Command {
	var opts bundleRemoveOptions

	cmd := &cobra.Command{
		Use:   "remove bundleID",
		Short: "Remove the book from the bundle",
		Run: func(cmd *cobra.Command, args []string) {
			opts.bundleID = args[0]
			runRemove(db, opts)
		},
	}
	flags := cmd.Flags()
	flags.IntVar(&opts.bookID, "book-id", 0, "Book ID")
	return cmd
}

func runRemove(db *gorm.DB, opts bundleRemoveOptions) {
	bundle.Remove(db, opts.bundleID, opts.bookID)
}
