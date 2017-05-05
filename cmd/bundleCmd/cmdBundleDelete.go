package bundleCmd

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/bundle"
)

//NewBundleDeleteCmd Delete the book into the bundle
func NewBundleDeleteCmd(db *gorm.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete bundleID",
		Short: "Delete the bundle",
		Run: func(cmd *cobra.Command, args []string) {
			bundleID := args[0]
			runDelete(db, bundleID)
		},
	}
	return cmd
}

func runDelete(db *gorm.DB, bundleID string) {
	bundle.Delete(db, bundleID)
}
