package bundleCmd

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/bundle"
)

type bundleCreateOptions struct {
	bundleName string
	note       string
}

//NewBundleCreateCmd Create the book into the bundle
func NewBundleCreateCmd(db *gorm.DB) *cobra.Command {
	var opts bundleCreateOptions
	cmd := &cobra.Command{
		Use:   "create bundleName",
		Short: "Create the bundle",
		Run: func(cmd *cobra.Command, args []string) {
			opts.bundleName = args[0]
			runCreate(db, opts)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&opts.note, "note", "", "Add note to bundle")
	return cmd
}

func runCreate(db *gorm.DB, opts bundleCreateOptions) {
	bundle.Create(db, opts.bundleName, opts.note)
}
