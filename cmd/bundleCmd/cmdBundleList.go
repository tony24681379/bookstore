package bundleCmd

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/bundle"
)

type bundleListOptions struct {
	bundleName string
	bookName   string
}

//NewBundleListCmd List the book into the bundle
func NewBundleListCmd(db *gorm.DB) *cobra.Command {
	var opts bundleListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the book which the bundle",
		Run: func(cmd *cobra.Command, args []string) {
			runList(db, opts)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&opts.bookName, "book-name", "", "List the book name")
	flags.StringVar(&opts.bundleName, "bundle-name", "", "List the bundle name")
	return cmd
}

func runList(db *gorm.DB, opts bundleListOptions) {
	bundle.List(db, opts.bookName, opts.bundleName)
}
