package bundleCmd

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/bundle"
)

type bundleUpdateOptions struct {
	bundleID   string
	bundleName string
	Note       string
}

//NewBundleUpdateCmd Update the book into the bundle
func NewBundleUpdateCmd(db *gorm.DB) *cobra.Command {
	var opts bundleUpdateOptions
	cmd := &cobra.Command{
		Use:   "update bundleID",
		Short: "Update the bundle",
		Run: func(cmd *cobra.Command, args []string) {
			opts.bundleID = args[0]
			runUpdate(db, opts)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&opts.bundleName, "bundle-name", "", "Update the bundle name")
	flags.StringVar(&opts.Note, "note", "", "Update the bundle note")
	return cmd
}

func runUpdate(db *gorm.DB, opts bundleUpdateOptions) {
	bundle.Update(db, opts.bundleID, opts.bundleName, opts.Note)
}
