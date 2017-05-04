package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/bundle"
)

//NewBundleDeleteCmd Delete the book into the bundle
func NewBundleDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete bundleID",
		Short: "Delete the bundle",
		Run: func(cmd *cobra.Command, args []string) {
			bundleID := args[0]
			runDelete(bundleID)
		},
	}
	return cmd
}

func runDelete(bundleID string) {
	bundle.Delete(bundleID)
}
