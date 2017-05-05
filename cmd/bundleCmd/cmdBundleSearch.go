package bundleCmd

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/bundle"
)

type bundleSearchOptions struct {
	bundleName string
	bookName   string
}

//NewBundleSearchCmd Search the book into the bundle
func NewBundleSearchCmd(db *gorm.DB) *cobra.Command {
	var opts bundleSearchOptions

	cmd := &cobra.Command{
		Use:   "search",
		Short: "Search the book which the bundle",
		Run: func(cmd *cobra.Command, args []string) {
			runSearch(db, opts)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&opts.bookName, "book-name", "", "Search the book name")
	flags.StringVar(&opts.bundleName, "bundle-name", "", "Search the bundle name")
	return cmd
}

func runSearch(db *gorm.DB, opts bundleSearchOptions) {
	bundle.Search(db, opts.bookName, opts.bundleName)
}
